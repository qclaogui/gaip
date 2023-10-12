// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
	platformFormat "github.com/containerd/containerd/platforms"
)

const (
	goImage   = "golang:1.21.3"                     // use golang:1.21.3 container as builder
	runImage  = "gcr.io/distroless/static"          // use gcr.io/distroless/static container as runtime
	imageRepo = "docker.io"                         // the container registry for the app image
	appImage  = "qclaogui/golang-api-server:latest" // the app image
)

var platforms = []dagger.Platform{"linux/amd64", "linux/arm64"}

func main() {
	println("Dagger is a programmable CI/CD engine that runs your pipelines in containers.")

	// check for Docker Hub registry credentials in host environment
	vars := []string{"DOCKERHUB_USERNAME", "DOCKERHUB_PASSWORD"}
	for _, v := range vars {
		if os.Getenv(v) == "" {
			panic(fmt.Sprintf("Environment variable %s is not set", v))
		}
	}

	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	// set registry password as secret for Dagger pipeline
	password := client.SetSecret("password", os.Getenv("DOCKERHUB_PASSWORD"))
	username := os.Getenv("DOCKERHUB_USERNAME")

	// source code directory
	source := client.Host().Directory(".")

	// use golang container as base
	goContainer := client.Container().
		From(goImage).
		WithMountedCache("/go/pkg/mod", client.CacheVolume("go-mod")).
		WithMountedCache("/root/.cache/go-build", client.CacheVolume("go-build")).
		WithEnvVariable("GOCACHE", "/root/.cache/go-build").
		WithMountedDirectory("/app", source).
		WithWorkdir("/app")

	platformVariants := make([]*dagger.Container, 0, len(platforms))
	for _, platform := range platforms {
		goContainer = goContainer.
			WithEnvVariable("GOOS", platformFormat.MustParse(string(platform)).OS).             // setup platform GOOS
			WithEnvVariable("GOARCH", platformFormat.MustParse(string(platform)).Architecture). // setup platform  GOARCH
			WithExec([]string{"make", "install-build-deps"})                                    // install dependencies tools

		// Running lint
		if _, err = goContainer.WithExec([]string{"make", "lint"}).Sync(ctx); err != nil {
			panic(err)
		}

		// Running tests
		if _, err = goContainer.WithExec([]string{"make", "test"}).Sync(ctx); err != nil {
			panic(err)
		}

		// Running build
		builder, err := goContainer.WithExec([]string{"make", "build"}).Sync(ctx)
		if err != nil {
			panic(err)
		}

		// copy binary file from builder
		app := client.Container(dagger.ContainerOpts{Platform: platform}).
			From(runImage).
			WithFile("/bin/main", builder.File("bin/golang-api-server")).
			WithEntrypoint([]string{"main"})

		platformVariants = append(platformVariants, app)
	}

	// CD: publish image to RepoCfg
	imageDigest, err := client.Container().WithRegistryAuth(imageRepo, username, password).
		Publish(ctx, appImage, dagger.ContainerPublishOpts{PlatformVariants: platformVariants})
	if err != nil {
		panic(err)
	}

	fmt.Println("Pushed multi-platform image w/ digest: ", imageDigest)
}
