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
	goImage   = "golang:1.22.1"            // use golang:1.22.1 container as builder
	runImage  = "gcr.io/distroless/static" // use gcr.io/distroless/static container as runtime
	imageRepo = "docker.io"                // the container registry for the app image
	appImage  = "qclaogui/gaip:latest"     // the app image
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
		WithMountedDirectory("/workspace", source).
		WithWorkdir("/workspace")

	platformVariants := make([]*dagger.Container, 0, len(platforms))
	for _, platform := range platforms {
		goos := platformFormat.MustParse(string(platform)).OS
		goarch := platformFormat.MustParse(string(platform)).Architecture

		path := fmt.Sprintf("bin/gaip_%s_%s", goos, goarch)

		goContainer = goContainer.
			WithEnvVariable("CGO_ENABLED", "0").
			WithEnvVariable("GOOS", goos).     // setup platform GOOS
			WithEnvVariable("GOARCH", goarch). // setup platform  GOARCH
			WithExec([]string{"go", "build", "-o", path, "cmd/main.go"})

		// // install dependencies tools
		// if _, err = goContainer.WithExec([]string{"make", "install-build-deps"}).Sync(ctx); err != nil {
		// 	panic(err)
		// }
		// // Running lint
		// if _, err = goContainer.WithExec([]string{"make", "lint"}).Sync(ctx); err != nil {
		// 	panic(err)
		// }
		// // Running tests
		// if _, err = goContainer.WithExec([]string{"make", "test"}).Sync(ctx); err != nil {
		// 	panic(err)
		// }
		// // Running build
		// builder, err := goContainer.WithExec([]string{"make", "build"}).Sync(ctx)
		// if err != nil {
		// 	panic(err)
		// }

		// copy binary file from builder
		app := client.Container(dagger.ContainerOpts{Platform: platform}).
			From(runImage).
			WithFile("/bin/gaip", goContainer.File(path)).
			WithEntrypoint([]string{"/bin/gaip"})

		platformVariants = append(platformVariants, app)
	}

	// // generate uuid for ttl.sh publish
	// _, _ = username, password
	// imageDigest, err := client.Container().
	// 	Publish(ctx, fmt.Sprintf("ttl.sh/gaip-%s:1h", uuid.New().String()), dagger.ContainerPublishOpts{PlatformVariants: platformVariants})

	// publish image
	imageDigest, err := client.Container().WithRegistryAuth(imageRepo, username, password).
		Publish(ctx, appImage, dagger.ContainerPublishOpts{PlatformVariants: platformVariants})
	if err != nil {
		panic(err)
	}

	fmt.Println("Pushed multi-platform image w/ digest: ", imageDigest)
}
