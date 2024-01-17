package call

import (
	"svix-poc/cmd/call/svix"
	"svix-poc/cmd/call/topic"
	"svix-poc/package/app/cmd"
)

var Root = cmd.New("call",
	cmd.Add(topic.Root, svix.Root),
)
