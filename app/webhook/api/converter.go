package webhookapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	webhookv1 "svix-poc/app/webhook/api/v1"

	"github.com/samber/lo"
	svix "github.com/svix/svix-webhooks/go"
)

// Types
// TODO: Map schemas

func WebhookTypeToJson(in svix.EventTypeOut) webhookv1.EventType {
	return webhookv1.EventType{
		Name:        in.Name,
		Description: in.Description,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
}
func WebhookTypesToJson(in []svix.EventTypeOut) []webhookv1.EventType {
	return lo.Map(in, func(item svix.EventTypeOut, _ int) webhookv1.EventType {
		return WebhookTypeToJson(item)
	})
}
func ListTypesToJson(in *svix.ListResponseEventTypeOut) webhookv1.EventTypeList {
	return webhookv1.EventTypeList{
		Data:   WebhookTypesToJson(in.Data),
		Cursor: PageInfo(in),
	}
}

// Endpoints
// TODO: Map metadata

func NewEndpointFromJson(in webhookv1.NewEndpoint) *svix.EndpointIn {
	res := &svix.EndpointIn{
		Url:         in.Url,
		Disabled:    in.Disabled,
		Description: in.Description,
		Channels:    lo.FromPtr(in.Channels),
		FilterTypes: lo.FromPtr(in.FilterTypes),
	}
	if in.Uid != nil {
		res.SetUid(*in.Uid)
	}
	if in.Secret != nil {
		res.SetSecret(*in.Secret)
	}
	if in.RateLimit != nil {
		res.SetRateLimit(*in.RateLimit)
	}
	return res
}
func NewEndpointsFromJson(in []webhookv1.NewEndpoint) []*svix.EndpointIn {
	return lo.Map(in, func(in webhookv1.NewEndpoint, _ int) *svix.EndpointIn {
		return NewEndpointFromJson(in)
	})
}
func EndpointToJson(in svix.EndpointOut) webhookv1.Endpoint {
	return webhookv1.Endpoint{
		Id:          in.Id,
		Url:         in.Url,
		Disabled:    in.Disabled,
		Channels:    in.Channels,
		CreatedAt:   in.CreatedAt,
		Description: in.Description,
		FilterTypes: in.FilterTypes,
		RateLimit:   in.RateLimit.Get(),
		Uid:         in.Uid.Get(),
	}
}
func EndpointsToJson(in []svix.EndpointOut) []webhookv1.Endpoint {
	return lo.Map(in, func(item svix.EndpointOut, _ int) webhookv1.Endpoint {
		return EndpointToJson(item)
	})
}
func ListEndpointsToJson(in *svix.ListResponseEndpointOut) webhookv1.EndpointList {
	return webhookv1.EndpointList{
		Data:   EndpointsToJson(in.Data),
		Cursor: PageInfo(in),
	}
}
func EndpointResultToJson(in []svix.EndpointOut, errs []webhookv1.Error) webhookv1.EndpointResult {
	return webhookv1.EndpointResult{
		Data:   EndpointsToJson(in),
		Errors: lo.EmptyableToPtr(errs),
	}
}

// Messages

func NewMessageFromJson(in webhookv1.NewMessage) *svix.MessageIn {
	var payload map[string]any
	if in.Payload != "" {
		json.Unmarshal([]byte(in.Payload), &payload)
	}
	ret := &svix.MessageIn{
		Tags:      lo.FromPtr(in.Tags),
		Channels:  lo.FromPtr(in.Channels),
		EventType: in.EventType,
		Payload:   payload,
	}
	if in.EventId != nil {
		ret.SetEventId(*in.EventId)
	}
	if in.RetentionPeriod != nil {
		ret.SetPayloadRetentionPeriod(*in.RetentionPeriod)
	}
	return ret
}
func NewMessagesFromJson(in []webhookv1.NewMessage) []*svix.MessageIn {
	return lo.Map(in, func(in webhookv1.NewMessage, _ int) *svix.MessageIn {
		return NewMessageFromJson(in)
	})
}
func MessageToJson(in svix.MessageOut) webhookv1.Message {
	var payload []byte
	if len(in.Payload) != 0 {
		payload, _ = json.Marshal(in.Payload)
	}
	return webhookv1.Message{
		Channels:  in.Channels,
		CreatedAt: in.Timestamp,
		EventId:   in.EventId.Get(),
		EventType: in.EventType,
		Id:        &in.Id,
		Payload:   string(payload),
	}
}
func MessagesToJson(in []svix.MessageOut) []webhookv1.Message {
	return lo.Map(in, func(in svix.MessageOut, _ int) webhookv1.Message {
		return MessageToJson(in)
	})
}

func ListMessagesToJson(in *svix.ListResponseMessageOut) webhookv1.MessageList {
	return webhookv1.MessageList{
		Data:   MessagesToJson(in.Data),
		Cursor: PageInfo(in),
	}
}
func MessageResultToJson(in []svix.MessageOut, errs []webhookv1.Error) webhookv1.MessageResult {
	return webhookv1.MessageResult{
		Data:   MessagesToJson(in),
		Errors: lo.EmptyableToPtr(errs),
	}
}

func EndpointMessageToJson(in svix.EndpointMessageOut) webhookv1.Message {
	return webhookv1.Message{
		Channels:  in.Channels,
		CreatedAt: in.Timestamp,
		EventId:   in.EventId.Get(),
		EventType: in.EventType,
		Id:        &in.Id,
	}
}
func EndpointMessagesToJson(in []svix.EndpointMessageOut) []webhookv1.Message {
	return lo.Map(in, func(in svix.EndpointMessageOut, _ int) webhookv1.Message {
		return EndpointMessageToJson(in)
	})
}
func ListEndpointMessagesToJson(in *svix.ListResponseEndpointMessageOut) webhookv1.MessageList {
	return webhookv1.MessageList{
		Data:   EndpointMessagesToJson(in.Data),
		Cursor: PageInfo(in),
	}
}

// Attempts
// TODO: map trigger types and status

func TriggerTypeToJson[T ~int32](in T) string {
	switch in {
	case 0:
	}
	return ""
}

func StatusToJson[T ~int32](in T) string {
	switch in {
	case 0:
	}
	return ""
}
func AttemptToJson(in svix.MessageAttemptOut) webhookv1.Attempt {
	return webhookv1.Attempt{
		Id:                 in.Id,
		Url:                in.Url,
		MsgId:              in.MsgId,
		EndpointId:         in.EndpointId,
		Response:           in.Response,
		Timestamp:          in.Timestamp,
		ResponseStatusCode: in.ResponseStatusCode,
		Status:             StatusToJson(in.Status),
		TriggerType:        TriggerTypeToJson(in.TriggerType),
	}
}
func AttemptsToJson(in []svix.MessageAttemptOut) []webhookv1.Attempt {
	return lo.Map(in, func(in svix.MessageAttemptOut, _ int) webhookv1.Attempt {
		return AttemptToJson(in)
	})
}
func ListAttemptsToJson(in *svix.ListResponseMessageAttemptOut) webhookv1.AttemptList {
	return webhookv1.AttemptList{
		Data:   AttemptsToJson(in.Data),
		Cursor: PageInfo(in),
	}
}
func EndpointAttemptToJson(in svix.MessageAttemptEndpointOut) webhookv1.Attempt {
	return webhookv1.Attempt{
		Id:                 in.Id,
		Url:                in.Url,
		MsgId:              in.MsgId,
		EndpointId:         in.EndpointId,
		Response:           in.Response,
		Timestamp:          in.Timestamp,
		ResponseStatusCode: in.ResponseStatusCode,
		Status:             StatusToJson(in.Status),
		TriggerType:        TriggerTypeToJson(in.TriggerType),
	}
}
func EndpointAttemptsToJson(in []svix.MessageAttemptEndpointOut) []webhookv1.Attempt {
	return lo.Map(in, func(in svix.MessageAttemptEndpointOut, _ int) webhookv1.Attempt {
		return EndpointAttemptToJson(in)
	})
}
func ListEndpointAttemptsToJson(in *svix.ListResponseMessageAttemptEndpointOut) webhookv1.AttemptList {
	return webhookv1.AttemptList{
		Data:   EndpointAttemptsToJson(in.Data),
		Cursor: PageInfo(in),
	}
}

// Pagination

type SvixPageInfo interface {
	GetDone() bool
	GetIterator() string
	GetPrevIterator() string
}

func PageInfo(in SvixPageInfo) webhookv1.PageInfo {
	return webhookv1.PageInfo{
		Done: in.GetDone(),
		Next: lo.EmptyableToPtr(in.GetIterator()),
		Prev: lo.EmptyableToPtr(in.GetPrevIterator()),
	}
}

type SvixCursor interface {
	SetIterator(string)
	SetLimit(int32)
}

// TODO: improve error handling
func Batch[F, T any](t []F, fn func(F) (*T, error)) (ret []T, errs []webhookv1.Error) {
	for i := range t {
		if r, err := fn(t[i]); err != nil {
			if serr, ok := ErrAs[svix.Error](err); ok {
				errs = append(errs, webhookv1.Error{
					Index:  lo.ToPtr(strconv.Itoa(i)),
					Reason: lo.ToPtr(serr.Error()),
					Code:   lo.ToPtr(serr.Status()),
				})
			} else {
				errs = append(errs, webhookv1.Error{
					Index:  lo.ToPtr(strconv.Itoa(i)),
					Reason: lo.ToPtr(err.Error()),
					Code:   lo.ToPtr(500),
				})
			}
		} else if r != nil {
			ret = append(ret, *r)
		}
	}
	return
}

func SetStatus(w http.ResponseWriter, err error) bool {
	if err == nil {
		return true
	}
	if serr, ok := ErrAs[svix.Error](err); ok {
		http.Error(w, serr.Error(), serr.Status())
		return false
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return false
}

func ErrAs[T any](err error) (T, bool) {
	var serr T
	if errors.As(err, &serr) {
		return serr, true
	}
	return serr, false
}
