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

// NewProject is a factory function to generate a new repository
func NewProject() (projectpb.ProjectServiceServer, error) {
	mr := &projectImpl{
		projects: map[string]*projectpb.Project{},
	}
	return mr, nil
}

// projectImpl fulfills the Repository projectImpl interface
// All objects are managed in an in-memory non-persistent store.
//
// projectImpl is used to implement projectpb.ProjectServiceServer.
type projectImpl struct {
	projectpb.UnimplementedProjectServiceServer

	mu sync.Mutex // global mutex to synchronize service access

	projects map[string]*projectpb.Project
}

func (p *projectImpl) CreateProject(_ context.Context, req *projectpb.CreateProjectRequest) (*projectpb.Project, error) {
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

	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.projects[pID]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "ProjectSrv with name %q already exists", pID)
	}

	p.projects[pID] = proj
	return p.projects[pID], nil
}

func (p *projectImpl) GetProject(_ context.Context, req *projectpb.GetProjectRequest) (*projectpb.Project, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	pID, err := name.ParseProject(req.Name)
	if err != nil {
		log.Printf("Invalid project name: %v", req.Name)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid project name")
	}

	proj, ok := p.projects[pID]
	if !ok {
		return nil, status.Errorf(codes.AlreadyExists, "ProjectSrv with name %q already exists", pID)
	}

	return proj, nil
}

// ListProjects returns up to pageSize number of projects beginning at pageToken, or from
// start if pageToken is the empty string.
func (p *projectImpl) ListProjects(_ context.Context, req *projectpb.ListProjectsRequest) (*projectpb.ListProjectsResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	_ = maxBatchSize
	size, _ := validatePageSize(req.PageSize)

	projects := make([]*projectpb.Project, len(p.projects))
	i := 0
	for k := range p.projects {
		projects[i] = p.projects[k]
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

func (p *projectImpl) DeleteProject(_ context.Context, req *projectpb.DeleteProjectRequest) (*emptypb.Empty, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	pID, err := name.ParseProject(req.Name)
	if err != nil {
		log.Printf("Invalid project name: %v", req.Name)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid project name")
	}
	_, ok := p.projects[pID]
	if !ok {
		return nil, status.Errorf(codes.AlreadyExists, "ProjectSrv with name %q already exists", pID)
	}

	delete(p.projects, pID)
	return &emptypb.Empty{}, nil
}
