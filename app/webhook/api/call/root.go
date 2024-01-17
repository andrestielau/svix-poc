package call

import "svix-poc/lib/app/cmd"

var Root = cmd.New("webhook", cmd.Add(EventType, App, Endpoint, Message, Attempt))
