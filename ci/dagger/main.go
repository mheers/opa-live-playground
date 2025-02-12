// A generated module for OpaLivePlayground functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/opa-live-playground/internal/dagger"
)

const (
	username    = "mheers"
	buildImage  = "golang:1.23-alpine"
	baseImage   = "alpine"
	targetImage = "docker.io/mheers/opa-live-playground:v0.1.0"
)

type OpaLivePlayground struct{}

func (m *OpaLivePlayground) BuildAndPushImage(ctx context.Context, src *dagger.Directory, registryToken *dagger.Secret) string {
	buildContainer := dag.Container().From(buildImage).
		WithExec([]string{"apk", "update"}).
		WithExec([]string{"apk", "add", "git", "wget"}).
		WithDirectory("/src", src, dagger.ContainerWithDirectoryOpts{
			Include: []string{"go.mod", "go.sum"},
		}).
		WithWorkdir("/src").
		WithExec([]string{"go", "mod", "download"}).
		WithDirectory("/src", src, dagger.ContainerWithDirectoryOpts{
			Exclude: []string{"node_modules", "js/dist", "js/node_modules", "go.work", "go.work.sum", ".idea", "__htmgo"},
		}).
		WithExec([]string{"mkdir", "-p", "/src/__htmgo"}).
		WithExec([]string{"wget", "-q", "-O", "/src/__htmgo/tailwind", "https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-linux-x64"}).
		WithExec([]string{"chmod", "+x", "/src/__htmgo/tailwind"}).
		WithExec([]string{"go", "run", "github.com/maddalax/htmgo/cli/htmgo@latest", "build"})

	targetContainer := dag.Container().From(baseImage).
		WithFile("/opa-live-playground", buildContainer.File("/src/dist/opa-live-playground")).
		WithEntrypoint([]string{"/opa-live-playground"})

	imageDigest, err := targetContainer.
		WithRegistryAuth(targetImage, username, registryToken).
		Publish(ctx, targetImage)

	if err != nil {
		panic(err)
	}
	return imageDigest
}
