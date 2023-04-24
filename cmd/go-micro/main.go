package main

import (
	"github.com/kioku-project/go-micro-cli/cmd"

	// register commands
	_ "github.com/kioku-project/go-micro-cli/cmd/call"
	_ "github.com/kioku-project/go-micro-cli/cmd/completion"
	_ "github.com/kioku-project/go-micro-cli/cmd/describe"
	_ "github.com/kioku-project/go-micro-cli/cmd/generate"
	_ "github.com/kioku-project/go-micro-cli/cmd/new"
	_ "github.com/kioku-project/go-micro-cli/cmd/run"
	_ "github.com/kioku-project/go-micro-cli/cmd/services"
	_ "github.com/kioku-project/go-micro-cli/cmd/stream"

	// plugins
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
)

func main() {
	cmd.Run()
}
