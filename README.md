# Binary Search Tree server
Manages integers as BST. Provides a REST API for basic tree actions

## Installation
Install [Golang 1.12](https://golang.org/doc/install) or higher.

Run `go get github.com/TriAnMan/alltoolstest/...` and `go build github.com/TriAnMan/alltoolstest/cmd/bst-server`

## Run
Download `https://github.com/TriAnMan/alltoolstest/blob/master/init.sample.json` into your working directory.

Run `./bst-server -init-file="./init.sample.json"`

## Design decisions
1. Use Red-Black Tree as of BST to achieve O(log n) time for a search operation (no need to rebalance a tree).
2. App exploits a fail-fast methodology.
3. Simple DDD pattern is used to prevent circular dependencies (https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/).
4. Non over-engineered architecture is implemented to reduce development time.
5. /insert and /delete REST methods are idempotent.
6. REST replies absent body but have consistent HTTP codes.

## Possible future improvements
1. Use some REST API framework with a SWAGGER docs generator.
2. Improve godoc.
3. Cover REST API methods with functional tests.
