package topic

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type Transformer interface {
	Transform(*message.Message) ([]*message.Message, error)
}
