package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"dagger.io/dagger"

	platformFormat "github.com/containerd/containerd/platforms"
	"golang.org/x/exp/slog"
)

// use golang:1.20.6-alpine container as builder
const goImage = "golang:1.20.6"

// the container registry for the app image
const imageRepo = "qclaogui/golang-api-server:latest"

var platforms = []dagger.Platform{"linux/amd64", "linux/arm64"}

func main() {
	println("Dagger is a programmable CI/CD engine that runs your pipelines in containers.")

	lvl := slog.LevelInfo
	// LOG_LEVEL is set, let's default to the desired level
	if lvlEnv, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if err := lvl.UnmarshalText([]byte(lvlEnv)); err != nil {
			slog.Error("unknown log level specified, choises are [DEBUG, INFO, WARN, ERROR]", "error", errors.New(lvlEnv))
			os.Exit(-1)
		}
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: lvl})))

	// check for Docker Hub registry credentials in host environment
	vars := []string{"DOCKERHUB_USERNAME", "DOCKERHUB_PASSWORD"}
	for _, v := range vars {
		if os.Getenv(v) == "" {
			slog.Warn(fmt.Sprintf("Environment variable %s is not set", v))
		}
	}

	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// set registry password as secret for Dagger pipeline
	password := client.SetSecret("password", os.Getenv("DOCKERHUB_PASSWORD"))
	username := os.Getenv("DOCKERHUB_USERNAME")

	// source code directory
	source := client.Host().Directory(".")

	platformVariants := make([]*dagger.Container, 0, len(platforms))

	for _, platform := range platforms {

		// use golang container as base
		goContainer := client.Container().From(goImage).
			WithMountedCache("/go/pkg/mod", client.CacheVolume("go-mod")).
			WithMountedCache("/root/.cache/go-build", client.CacheVolume("go-build")).
			WithMountedDirectory("/app", source).
			WithWorkdir("/app").
			WithEnvVariable("GOCACHE", "/root/.cache/go-build"). // set GOCACHE explicitly to point to our mounted cache
			WithEnvVariable("GOOS", platformFormat.MustParse(string(platform)).OS).
			WithEnvVariable("GOARCH", platformFormat.MustParse(string(platform)).Architecture)

		// Running build
		builder, err := goContainer.WithExec([]string{"make", "build"}).Sync(ctx)
		if err != nil {
			slog.Error("Executing the tests failed", "error", err)
			os.Exit(-1)
		}

		// Running tests
		if _, err := goContainer.WithExec([]string{"make", "test"}).Sync(ctx); err != nil {
			slog.Error("Executing the tests failed", "error", err)
			os.Exit(-1)
		}

		// copy binary file from builder
		app := client.Container(dagger.ContainerOpts{Platform: platform}).From("gcr.io/distroless/static").
			WithFile("/bin/main", builder.File("bin/golang-api-server")).
			WithEntrypoint([]string{"main"})

		platformVariants = append(platformVariants, app)
	}

	imageDigest, err := client.Container().WithRegistryAuth("docker.io", username, password).
		Publish(ctx, imageRepo, dagger.ContainerPublishOpts{PlatformVariants: platformVariants})
	if err != nil {
		slog.Error("Executing the tests failed", err)
		os.Exit(-1)
	}

	fmt.Println("Pushed multi-platform image w/ digest: ", imageDigest)
}
