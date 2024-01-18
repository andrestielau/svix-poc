package call

import (
	callrouterapi "svix-poc/app/router/api/call"
	callroutergrpc "svix-poc/app/router/grpc/call"
	callroutertopic "svix-poc/app/router/topic/call"
	callwebhookapi "svix-poc/app/webhook/api/call"
	callwebhookgrpc "svix-poc/app/webhook/grpc/call"
	callwebhooktopic "svix-poc/app/webhook/topic/call"
	"svix-poc/cmd/call/mock"
	"svix-poc/cmd/call/svix"
	"svix-poc/cmd/call/topic"
	"svix-poc/lib/app/cmd"
)

var Root = cmd.New("call",
	cmd.Add(topic.Root, svix.Root, mock.Root,
		cmd.New("api", cmd.Alias("a"), cmd.Add(callwebhookapi.Root, callrouterapi.Root)),
		cmd.New("grpc", cmd.Alias("g"), cmd.Add(callwebhookgrpc.Root, callroutergrpc.Root)),
		cmd.New("publish", cmd.Alias("p"), cmd.Add(callwebhooktopic.Root, callroutertopic.Root)),
	),
)
