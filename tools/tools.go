//go:build tools
// +build tools

package tools

import (
	_ "github.com/daixiang0/gci"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
