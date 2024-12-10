// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package routeguide_test

import (
	"context"
	"errors"
	"io"
	"log"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	routeguide "github.com/qclaogui/gaip/genproto/routeguide/apiv1"
	pb "github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"google.golang.org/protobuf/proto"
)

// Clients are initialized in main_test.go.
var (
	routeguideGRPC *routeguide.Client
	// routeguideREST *routeguide.Client
)

func TestGetFeature(t *testing.T) {
	req := &pb.GetFeatureRequest{
		Point: &pb.Point{
			Latitude:  409146138,
			Longitude: -746188906,
		},
	}

	for typ, client := range map[string]*routeguide.Client{"grpc": routeguideGRPC} {
		resp, err := client.GetFeature(context.Background(), req)
		if err != nil {
			t.Fatalf("client.GetFeature() failed: %v", err)
		}

		// The point where the feature is detected.
		point := resp.GetFeature().GetLocation()
		if diff := cmp.Diff(point, req.GetPoint(), cmp.Comparer(proto.Equal)); diff != "" {
			t.Errorf("%s client.GetFeature() got=-, want=+:%s", typ, diff)
		}
	}
}

func TestListFeatures(t *testing.T) {
	rect := &pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
	}
	req := &pb.ListFeaturesRequest{
		Rectangle: rect,
	}

	for typ, client := range map[string]*routeguide.Client{"grpc": routeguideGRPC} {
		stream, err := client.ListFeatures(context.Background(), req)
		if err != nil {
			t.Fatalf("client.ListFeatures() failed: %v", err)
		}

		for {
			resp, err2 := stream.Recv()
			if errors.Is(err2, io.EOF) {
				break
			}
			if err2 != nil {
				t.Fatalf("client.ListFeatures() failed: %v", err2)
			}

			feature := resp.GetFeature()
			if !inRange(feature.Location, rect) {
				t.Errorf("%s client.ListFeatures() got Feature: name: %q, point:(%v, %v) not inRange.", typ, feature.GetName(),
					feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
			}
		}
	}
}

// TestRecordRoute runRecordRoute sends a sequence of points to server and expects to get a RouteSummary from server.
func TestRecordRoute(t *testing.T) {
	// Create a random number of random points
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2 // Traverse at least two points
	var points []*pb.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}

	for typ, client := range map[string]*routeguide.Client{"grpc": routeguideGRPC} {
		stream, err := client.RecordRoute(context.Background())
		if err != nil {
			t.Fatalf("client.RecordRoute() failed: %v", err)
		}

		//  Send points
		for _, point := range points {
			if err := stream.Send(&pb.RecordRouteRequest{Point: point}); err != nil {
				t.Fatalf("client.RecordRoute(): stream.Send(%v) failed: %v", point, err)
			}
		}
		//  recv reply
		reply, err := stream.CloseAndRecv()
		if err != nil {
			t.Fatalf("client.RecordRoute() failed: %v", err)
		}

		summary := reply.GetRouteSummary()
		if diff := cmp.Diff(int(summary.PointCount), len(points)); diff != "" {
			t.Errorf("%s client.RecordRoute() got=-, want=+:%s", typ, diff)
		}
	}
}

func TestRouteChat(t *testing.T) {
	notes := []*pb.RouteNote{
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "First message"},  // recvNoteCount + 1
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Second message"}, // recvNoteCount + 1
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},  // recvNoteCount + 1
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"}, // recvNoteCount + 2
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},  // recvNoteCount + 2
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},  // recvNoteCount + 2
	}

	for typ, client := range map[string]*routeguide.Client{"grpc": routeguideGRPC} {
		stream, err := client.RouteChat(context.Background())
		if err != nil {
			t.Fatalf("client.RouteChat() failed: %v", err)
		}

		waitc := make(chan struct{})
		recvNoteCount := 0
		go func() {
			for {
				in, err := stream.Recv()
				if errors.Is(err, io.EOF) {
					close(waitc)
					return
				}
				if err != nil {
					log.Fatalf("client.RouteChat() failed: %v", err)
				}
				routeNote := in.GetRouteNote()
				log.Printf("Got message %s at point(%d, %d)",
					routeNote.GetMessage(),
					routeNote.GetLocation().Latitude,
					routeNote.GetLocation().Longitude)
				recvNoteCount++
			}
		}()

		for _, note := range notes {
			if err := stream.Send(&pb.RouteChatRequest{RouteNote: note}); err != nil {
				t.Fatalf("client.RouteChat(): stream.Send(%v) failed: %v", note, err)
			}
		}
		if err = stream.CloseSend(); err != nil {
			t.Fatalf("client.RouteChat(): stream.CloseSend() failed: %v", err)
		}
		<-waitc

		if diff := cmp.Diff(recvNoteCount, 9); diff != "" {
			t.Errorf("%s client.RouteChat() got=-, want=+:%s", typ, diff)
		}
	}
}

func randomPoint(r *rand.Rand) *pb.Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &pb.Point{Latitude: lat, Longitude: long}
}

func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))

	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left && float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom && float64(point.Latitude) <= top {
		return true
	}
	return false
}
