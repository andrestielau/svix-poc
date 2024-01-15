package webhookapi

import (
	"errors"
	"net/http"
	eventsv1 "svix-poc/module/router/api/v1"

	"github.com/samber/lo"
	svix "github.com/svix/svix-webhooks/go"
)

// Types
// TODO: Map schemas

func WebhookTypeToJson(in svix.EventTypeOut) eventsv1.WebHookType {
	return eventsv1.WebHookType{
		Name:        in.Name,
		Description: in.Description,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
}
func WebhookTypesToJson(in []svix.EventTypeOut) []eventsv1.WebHookType {
	return lo.Map(in, func(item svix.EventTypeOut, _ int) eventsv1.WebHookType {
		return WebhookTypeToJson(item)
	})
}
func ListWebHookTypesToJson(in *svix.ListResponseEventTypeOut) eventsv1.WebHookTypeList {
	return eventsv1.WebHookTypeList{
		Data:   WebhookTypesToJson(in.Data),
		Cursor: PageInfo(in),
	}
}

// Endpoints
// TODO: Map metadata

func NewWebHookEndpointFromJson(in eventsv1.NewWebHookEndpoint) *svix.EndpointIn {
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
func NewWebHookEndpointsFromJson(in []eventsv1.NewWebHookEndpoint) []*svix.EndpointIn {
	return lo.Map(in, func(in eventsv1.NewWebHookEndpoint, _ int) *svix.EndpointIn {
		return NewWebHookEndpointFromJson(in)
	})
}
func WebHookEndpointToJson(in svix.EndpointOut) eventsv1.WebHookEndpoint {
	return eventsv1.WebHookEndpoint{
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
func WebHookEndpointsToJson(in []svix.EndpointOut) []eventsv1.WebHookEndpoint {
	return lo.Map(in, func(item svix.EndpointOut, _ int) eventsv1.WebHookEndpoint {
		return WebHookEndpointToJson(item)
	})
}
func ListWebHookEndpointsToJson(in *svix.ListResponseEndpointOut) eventsv1.WebHookEndpointList {
	return eventsv1.WebHookEndpointList{
		Data:   WebHookEndpointsToJson(in.Data),
		Cursor: PageInfo(in),
	}
}
func WebHookEndpointResultToJson(in []svix.EndpointOut, errs []eventsv1.Error) eventsv1.WebHookEndpointResult {
	return eventsv1.WebHookEndpointResult{
		Data:   WebHookEndpointsToJson(in),
		Errors: lo.EmptyableToPtr(errs),
	}
}

// Messages
// TODO: Map payload

func NewWebHookMessageFromJson(in eventsv1.NewWebHookMessage) *svix.MessageIn {
	ret := &svix.MessageIn{
		Tags:      *in.Tags,
		Channels:  *in.Channels,
		EventType: in.EventType,
	}
	if in.EventId != nil {
		ret.SetEventId(*in.EventId)
	}
	if in.RetentionPeriod != nil {
		ret.SetPayloadRetentionPeriod(*in.RetentionPeriod)
	}
	return ret
}
func NewWebHookMessagesFromJson(in []eventsv1.NewWebHookMessage) []*svix.MessageIn {
	return lo.Map(in, func(in eventsv1.NewWebHookMessage, _ int) *svix.MessageIn {
		return NewWebHookMessageFromJson(in)
	})
}
func WebHookMessageToJson(in svix.MessageOut) eventsv1.WebHookMessage {
	return eventsv1.WebHookMessage{
		Channels:  in.Channels,
		CreatedAt: in.Timestamp,
		EventId:   in.EventId.Get(),
		EventType: in.EventType,
		Id:        &in.Id,
	}
}
func WebHookMessagesToJson(in []svix.MessageOut) []eventsv1.WebHookMessage {
	return lo.Map(in, func(in svix.MessageOut, _ int) eventsv1.WebHookMessage {
		return WebHookMessageToJson(in)
	})
}

func ListWebHookMessagesToJson(in *svix.ListResponseMessageOut) eventsv1.WebHookMessageList {
	return eventsv1.WebHookMessageList{
		Data:   WebHookMessagesToJson(in.Data),
		Cursor: PageInfo(in),
	}
}
func WebHookMessageResultToJson(in []svix.MessageOut, errs []eventsv1.Error) eventsv1.WebHookMessageResult {
	return eventsv1.WebHookMessageResult{
		Data:   WebHookMessagesToJson(in),
		Errors: lo.EmptyableToPtr(errs),
	}
}

func WebHookEndpointMessageToJson(in svix.EndpointMessageOut) eventsv1.WebHookMessage {
	return eventsv1.WebHookMessage{
		Channels:  in.Channels,
		CreatedAt: in.Timestamp,
		EventId:   in.EventId.Get(),
		EventType: in.EventType,
		Id:        &in.Id,
	}
}
func WebHookEndpointMessagesToJson(in []svix.EndpointMessageOut) []eventsv1.WebHookMessage {
	return lo.Map(in, func(in svix.EndpointMessageOut, _ int) eventsv1.WebHookMessage {
		return WebHookEndpointMessageToJson(in)
	})
}
func ListWebHookEndpointMessagesToJson(in *svix.ListResponseEndpointMessageOut) eventsv1.WebHookMessageList {
	return eventsv1.WebHookMessageList{
		Data:   WebHookEndpointMessagesToJson(in.Data),
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

func WebHookStatusToJson[T ~int32](in T) string {
	switch in {
	case 0:
	}
	return ""
}
func WebHookAttemptToJson(in svix.MessageAttemptOut) eventsv1.WebHookAttempt {
	return eventsv1.WebHookAttempt{
		Id:                 in.Id,
		Url:                in.Url,
		MsgId:              in.MsgId,
		EndpointId:         in.EndpointId,
		Response:           in.Response,
		Timestamp:          in.Timestamp,
		ResponseStatusCode: in.ResponseStatusCode,
		Status:             WebHookStatusToJson(in.Status),
		TriggerType:        TriggerTypeToJson(in.TriggerType),
	}
}
func WebHookAttemptsToJson(in []svix.MessageAttemptOut) []eventsv1.WebHookAttempt {
	return lo.Map(in, func(in svix.MessageAttemptOut, _ int) eventsv1.WebHookAttempt {
		return WebHookAttemptToJson(in)
	})
}
func ListWebHookAttemptsToJson(in *svix.ListResponseMessageAttemptOut) eventsv1.WebHookAttemptList {
	return eventsv1.WebHookAttemptList{
		Data:   WebHookAttemptsToJson(in.Data),
		Cursor: PageInfo(in),
	}
}
func WebHookEndpointAttemptToJson(in svix.MessageAttemptEndpointOut) eventsv1.WebHookAttempt {
	return eventsv1.WebHookAttempt{
		Id:                 in.Id,
		Url:                in.Url,
		MsgId:              in.MsgId,
		EndpointId:         in.EndpointId,
		Response:           in.Response,
		Timestamp:          in.Timestamp,
		ResponseStatusCode: in.ResponseStatusCode,
		Status:             WebHookStatusToJson(in.Status),
		TriggerType:        TriggerTypeToJson(in.TriggerType),
	}
}
func WebHookEndpointAttemptsToJson(in []svix.MessageAttemptEndpointOut) []eventsv1.WebHookAttempt {
	return lo.Map(in, func(in svix.MessageAttemptEndpointOut, _ int) eventsv1.WebHookAttempt {
		return WebHookEndpointAttemptToJson(in)
	})
}
func ListWebHookEndpointAttemptsToJson(in *svix.ListResponseMessageAttemptEndpointOut) eventsv1.WebHookAttemptList {
	return eventsv1.WebHookAttemptList{
		Data:   WebHookEndpointAttemptsToJson(in.Data),
		Cursor: PageInfo(in),
	}
}

// Pagination

type SvixPageInfo interface {
	GetDone() bool
	GetIterator() string
	GetPrevIterator() string
}

func PageInfo(in SvixPageInfo) eventsv1.PageInfo {
	return eventsv1.PageInfo{
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
func Batch[F, T any](t []F, fn func(F) (*T, error)) (ret []T, errs []eventsv1.Error) {
	for i := range t {
		if r, err := fn(t[i]); err != nil {
			errs = append(errs, eventsv1.Error{})
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
	var serr svix.Error
	if errors.As(err, &serr) {
		http.Error(w, serr.Error(), serr.Status())
		return false
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return false
}
