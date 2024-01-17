# 0 Preparation
- Run `docker-compose up -d` and get Svix token from `token` pod's logs
- Run `export SVIX_AUTH_TOKEN={token}`
- Run `go run . svix event-type create -n TestEvent`
- Run `go run . svix app create -n Test` and get App's Id
- Run `export SVIX_APP_ID={app-id}`
- Run `go run . svix endpoint create -a $SVIX_APP_ID -n Test`

# 1 - CLI Sync WebHook Request
- Run `go run . svix message create -a $SVIX_APP_ID -p '{ "foo": "bar" }'`
Run service for webhooks and call its grpc

# 2 - Async WebHook Request
Run service for webhooks
- Run `go run . svix topic publish webhook -m app=$SVIX_APP_ID -p '{ "foo": "bar" }'`

# 3 - WebHook Request on Subscribed Event
Run service for events and webhooks