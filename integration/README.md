# Integration tests

This document provides an outline of how to run and add new integration tests to the project.

The purpose of these tests is to verify simple, happy-path pipelines to catch issues between the gaip and external dependencies.

The external dependencies are launched as Docker containers.

## Running tests(Makefile)

Once the Docker image is built, you can run integration tests:

```shell
go test -v -tags=requires_docker ./integration/tests/...

```

If you want to run a single test you can use a filter. For example, to only run `TestEcho`:

```shell
go test -v -tags=requires_docker ./integration/tests/... -run "^TestEcho"

```

When running all integration tests, the test process may time out before the tests complete. If this happens, you can increase `go test`'s default 10 minute timeout to something longer with the `-timeout` flag, for example:

```shell
go test -v -tags=requires_docker -timeout=20m ./integration/tests/...

```

## Running tests

Execute the integration tests using the following command:

`go run .`

### Flags

* `--test`: Specifies a particular directory within the tests directory to run (default: runs all tests)

## Adding new tests

Follow these steps to add a new integration test to the project:

- Create a new directory under the tests directory to house the files for the new test.
- Create a `_test.go` file within the new test directory.
