package main

import (
	"context"
	"errors"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/qclaogui/golang-api-server/api/gen/proto/routeguide/v1"
	"github.com/qclaogui/golang-api-server/cmd/client-grpc/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:9095", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {

		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func() { _ = conn.Close() }()

	client := pb.NewRouteGuideServiceClient(conn)

	// Looking for a valid feature
	printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})

	// test Recovery
	//printFeatures(client, nil)

	printFeatures(client, &pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
	})

	//// RecordRoute
	//runRecordRoute(client)
	//
	//// RouteChat
	//runRouteChat(client)

}

// printFeature gets the feature for the given point.
func printFeature(client pb.RouteGuideServiceClient, point *pb.Point) {
	log.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.GetFeature(ctx, &pb.GetFeatureRequest{Point: point})
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}
	log.Println(resp.Feature)
}

// printFeatures lists all the features within the given bounding Rectangle.
func printFeatures(client pb.RouteGuideServiceClient, rect *pb.Rectangle) {
	log.Printf("Looking for features within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ListFeatures(ctx, &pb.ListFeaturesRequest{Rectangle: rect})
	if err != nil {
		log.Fatalf("client.ListFeatures failed: %v", err)
	}
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("client.ListFeatures failed: %v", err)
		}
		feature := resp.GetFeature()
		log.Printf("Feature: name: %q, point:(%v, %v)", feature.GetName(),
			feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
	}
}
