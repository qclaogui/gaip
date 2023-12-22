// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"context"
	"log"
	"sort"
	"sync"

	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/pkg/service/project/name"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Project fulfills the Repository Project interface
// All objects are managed in an in-memory non-persistent store.
//
// Project is used to implement ProjectServiceServer.
type Project struct {
	mu sync.Mutex // global mutex to synchronize service access

	projects map[string]*projectpb.Project
}

// NewProject is a factory function to generate a new repository
func NewProject() (*Project, error) {
	mr := &Project{
		projects: map[string]*projectpb.Project{},
	}
	return mr, nil
}

func (mr *Project) CreateProject(_ context.Context, req *projectpb.CreateProjectRequest) (*projectpb.Project, error) {
	proj := req.Project
	if proj == nil {
		log.Print("ProjectSrv must not be empty.")
		return nil, status.Errorf(codes.InvalidArgument, "ProjectSrv must not be empty")
	}

	if proj.Name == "" {
		log.Printf("ProjectSrv name must not be empty: %v", proj.Name)
		return nil, status.Errorf(codes.InvalidArgument, "ProjectSrv name must not be empty")
	}

	pID, err := name.ParseProject(proj.Name)
	if err != nil {
		log.Printf("Invalid project name: %v", proj.Name)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid project name")
	}

	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, ok := mr.projects[pID]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "ProjectSrv with name %q already exists", pID)
	}

	mr.projects[pID] = proj
	return mr.projects[pID], nil
}

func (mr *Project) GetProject(_ context.Context, req *projectpb.GetProjectRequest) (*projectpb.Project, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	pID, err := name.ParseProject(req.Name)
	if err != nil {
		log.Printf("Invalid project name: %v", req.Name)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid project name")
	}

	p, ok := mr.projects[pID]
	if !ok {
		return nil, status.Errorf(codes.AlreadyExists, "ProjectSrv with name %q already exists", pID)
	}

	return p, nil
}

// ListProjects returns up to pageSize number of projects beginning at pageToken, or from
// start if pageToken is the empty string.
func (mr *Project) ListProjects(_ context.Context, req *projectpb.ListProjectsRequest) (*projectpb.ListProjectsResponse, error) {
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

func (mr *Project) DeleteProject(_ context.Context, req *projectpb.DeleteProjectRequest) (*emptypb.Empty, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	pID, err := name.ParseProject(req.Name)
	if err != nil {
		log.Printf("Invalid project name: %v", req.Name)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid project name")
	}
	_, ok := mr.projects[pID]
	if !ok {
		return nil, status.Errorf(codes.AlreadyExists, "ProjectSrv with name %q already exists", pID)
	}

	delete(mr.projects, pID)
	return &emptypb.Empty{}, nil
}