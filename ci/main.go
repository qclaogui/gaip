package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"dagger.io/dagger"

	"golang.org/x/exp/slog"
)

// use golang:1.20.6-alpine container as builder
const goImage = "golang:1.20.6-alpine"

// the container registry for the app image
const imageRepo = "qclaogui/golang-api-server:latest"

func main() {
	println("Building with Dagger")

	lvl := slog.LevelInfo
	// LOG_LEVEL is set, let's default to the desired level
	if lvlEnv, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if err := lvl.UnmarshalText([]byte(lvlEnv)); err != nil {
			slog.Error("unknown log level specified, choises are [DEBUG, INFO, WARN, ERROR]", errors.New(lvlEnv))
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

	// create a cache volume for Go downloads
	goModCache := client.CacheVolume("gomodcache")

	// set registry password as secret for Dagger pipeline
	password := client.SetSecret("password", os.Getenv("DOCKERHUB_PASSWORD"))
	username := os.Getenv("DOCKERHUB_USERNAME")

	// get reference to source code directory
	source := client.Host().Directory(".")

	goContainer := client.Container().From(goImage).
		WithMountedCache("/go/pkg/mod", goModCache).
		WithMountedDirectory("/app", source).
		WithWorkdir("/app")

	// Running build
	builder, err := goContainer.WithExec([]string{"make", "build"}).Sync(ctx)
	if err != nil {
		slog.Error("Executing the tests failed", err)
		os.Exit(-1)
	}

	// Running tests
	if _, err := goContainer.WithExec([]string{"make", "test"}).Sync(ctx); err != nil {
		slog.Error("Executing the tests failed", err)
		os.Exit(-1)
	}

	// use gcr.io/distroless/static container as base
	// copy binary file from builder
	appImage := client.Container().From("gcr.io/distroless/static").
		WithFile("/bin/main", builder.File("bin/golang-api-server")).
		WithEntrypoint([]string{"main"})

	// Running deploy
	// publishing the final image
	imageDigest, err := appImage.WithRegistryAuth("docker.io", username, password).
		Publish(ctx, imageRepo)
	if err != nil {
		slog.Error("Executing the tests failed", err)
		os.Exit(-1)
	}
	fmt.Println("Image published at: ", imageDigest)
}
