// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package repository

import (
	"context"
	"flag"
	"log"
	"sort"
	"strconv"
	"sync"

	"github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/golang-api-server/pkg/service/project/name"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	defaultPageSize = 20
	maxPageSize     = 10000
	maxBatchSize    = 1000
)

type MemoryConfig struct {
	Enabled bool `yaml:"enabled"`
}

func (cfg *MemoryConfig) RegisterFlagsWithPrefix(prefix string, fs *flag.FlagSet) {
	fs.BoolVar(&cfg.Enabled, prefix+"memory.enabled", true, "Enables memory Repository")
}

func (cfg *MemoryConfig) RegisterFlags(fs *flag.FlagSet) {
	cfg.RegisterFlagsWithPrefix("", fs)
}

func (cfg *MemoryConfig) Validate() error {
	if !cfg.Enabled {
		return nil
	}

	return nil
}

// MemoryRepo fulfills the Repository interface
// All objects are managed in an in-memory non-persistent store.
//
// MemoryRepo is used to implement ProjectServiceServer.
type MemoryRepo struct {
	mu sync.Mutex // global mutex to synchronize service access

	projects map[string]*projectpb.Project
}

// NewMemoryRepo is a factory function to generate a new repository
func NewMemoryRepo() (*MemoryRepo, error) {
	mr := &MemoryRepo{
		projects: map[string]*projectpb.Project{},
	}
	return mr, nil
}

// validatePageSize returns the default page size if the specified page size is 0, otherwise it
// validates the specified page size.
func validatePageSize(ps int32) (int32, error) {
	switch {
	case ps == 0:
		return defaultPageSize, nil
	case ps > maxPageSize:
		return 0, status.Errorf(codes.InvalidArgument, "page size %d cannot be large than max page size %d", ps, maxPageSize)
	case ps < 0:
		return 0, status.Errorf(codes.InvalidArgument, "page size %d cannot be negative", ps)
	}

	return ps, nil
}

// Parses the page token to an int. Returns defaultValue if parsing fails
func parsePageToken(pageToken string, defaultValue int) int {
	if pageToken == "" {
		return defaultValue
	}
	parsed, err := strconv.Atoi(pageToken)
	if err != nil {
		return defaultValue
	}
	return parsed
}

// nextPageToken returns the next page token (the next item index or empty if not more items are left)
func nextPageToken(lastPage, total int) string {
	if lastPage == total {
		return ""
	}
	return strconv.Itoa(lastPage)
}

func (mr *MemoryRepo) CreateProject(_ context.Context, req *projectpb.CreateProjectRequest) (*projectpb.Project, error) {
	proj := req.Project
	if proj == nil {
		log.Print("ProjectServer must not be empty.")
		return nil, status.Errorf(codes.InvalidArgument, "ProjectServer must not be empty")
	}

	if proj.Name == "" {
		log.Printf("ProjectServer name must not be empty: %v", proj.Name)
		return nil, status.Errorf(codes.InvalidArgument, "ProjectServer name must not be empty")
	}

	pID, err := name.ParseProject(proj.Name)
	if err != nil {
		log.Printf("Invalid project name: %v", proj.Name)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid project name")
	}

	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, ok := mr.projects[pID]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "ProjectServer with name %q already exists", pID)
	}

	mr.projects[pID] = proj
	return mr.projects[pID], nil
}

func (mr *MemoryRepo) GetProject(_ context.Context, req *projectpb.GetProjectRequest) (*projectpb.Project, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	pID, err := name.ParseProject(req.Name)
	if err != nil {
		log.Printf("Invalid project name: %v", req.Name)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid project name")
	}

	p, ok := mr.projects[pID]
	if !ok {
		return nil, status.Errorf(codes.AlreadyExists, "ProjectServer with name %q already exists", pID)
	}

	return p, nil
}

// ListProjects returns up to pageSize number of projects beginning at pageToken, or from
// start if pageToken is the empty string.
func (mr *MemoryRepo) ListProjects(_ context.Context, req *projectpb.ListProjectsRequest) (*projectpb.ListProjectsResponse, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	_ = maxBatchSize
	size, _ := validatePageSize(req.PageSize)

	projects := make([]*projectpb.Project, len(mr.projects))
	i := 0
	for k := range mr.projects {
		projects[i] = mr.projects[k]
		i++
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Name < projects[j].Name
	})

	startPos := parsePageToken(req.PageToken, 0)

	endPos := minx(startPos+int(size), len(projects))

	resp := projectpb.ListProjectsResponse{
		Projects:      projects[startPos:endPos],
		NextPageToken: nextPageToken(endPos, len(projects)),
	}

	return &resp, nil

}

// Returns the smallest of a and b
func minx(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (mr *MemoryRepo) DeleteProject(_ context.Context, req *projectpb.DeleteProjectRequest) (*emptypb.Empty, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	pID, err := name.ParseProject(req.Name)
	if err != nil {
		log.Printf("Invalid project name: %v", req.Name)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid project name")
	}
	_, ok := mr.projects[pID]
	if !ok {
		return nil, status.Errorf(codes.AlreadyExists, "ProjectServer with name %q already exists", pID)
	}

	delete(mr.projects, pID)
	return &emptypb.Empty{}, nil
}
