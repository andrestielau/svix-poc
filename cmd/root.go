package cmd

import (
	"svix-poc/cmd/demo"
	"svix-poc/cmd/serve"
	"svix-poc/cmd/topic"
	"svix-poc/package/app/cmd"
)

var Root = cmd.New("svix-poc",
	cmd.Add(serve.Root, topic.Root,
		demo.Root, demo.Mock),
)
