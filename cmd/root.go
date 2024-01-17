package cmd

import (
	"svix-poc/cmd/call"
	"svix-poc/cmd/demo"
	"svix-poc/cmd/serve"
	"svix-poc/lib/app/cmd"
)

var Root = cmd.New("svix-poc",
	cmd.Add(serve.Root, call.Root, demo.Root, demo.Mock))
