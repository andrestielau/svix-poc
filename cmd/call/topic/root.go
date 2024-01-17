package topic

import (
	"errors"
	"log"
	"os"
	"svix-poc/lib/app/cmd"
	"svix-poc/lib/app/flag"

	"cloud.google.com/go/pubsub"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

var metadata message.Metadata
var client *pubsub.Client
var projectId string
var payload string
var Root = cmd.New("topic",
	cmd.Alias("a"),
	cmd.Add(
		cmd.New("list", cmd.Alias("l"), cmd.Run(runList)),
		cmd.New("create", cmd.Alias("c"), cmd.Run(runCreate)),
		cmd.New("publish", cmd.Alias("p"), cmd.Run(runPublish), cmd.Flags(
			flag.StringP(&payload, "data", "", "d", "data"),
			flag.StringToStringP(&metadata, "metadata", nil, "m", "metadata"),
		)),
	),
	cmd.PFlags(
		flag.StringP(&projectId, "project", "demo", "p", "Project Id"),
	),
	cmd.PPreRun(func(cmd *cobra.Command, _ []string) {
		os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085") // TODO: google black magic requires this or an actual env
		client = lo.Must(pubsub.NewClient(cmd.Context(), projectId))
	}),
	cmd.PPostRun(func(_ *cobra.Command, _ []string) {
		if client != nil {
			lo.Must0(client.Close())
		}
	}),
)

func runList(cmd *cobra.Command, _ []string) {
	it := client.Topics(cmd.Context())
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		lo.Must0(err)
		log.Println(topic)
	}
}

func runCreate(cmd *cobra.Command, topics []string) {
	var errs []error
	for _, topic := range topics {
		if _, err := client.CreateTopic(cmd.Context(), topic); err != nil {
			errs = append(errs, err)
		}
	}
	lo.Must0(errors.Join(errs...))
}

func runPublish(cmd *cobra.Command, topics []string) {
	pub := lo.Must(googlecloud.NewPublisher(googlecloud.PublisherConfig{
		ProjectID: projectId,
	}, watermill.NewStdLogger(true, false)))
	defer pub.Close()
	for _, topic := range topics {
		lo.Must0(pub.Publish(topic, &message.Message{
			UUID:     uuid.NewString(),
			Payload:  []byte(payload),
			Metadata: metadata,
		}))
	}
}
