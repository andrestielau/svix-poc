package webhookgrpc

import (
	webhooksv1 "svix-poc/app/webhook/grpc/v1"

	"github.com/samber/lo"
	svix "github.com/svix/svix-webhooks/go"
)

func AppFromProto(v *webhooksv1.App) *svix.ApplicationIn {
	if v == nil {
		return nil
	}
	m := &svix.ApplicationIn{
		Name: v.Name,
	}
	if v.Uid != "" {
		m.SetUid(v.Uid)
	}
	return m
}
func AppToProto(v *svix.ApplicationOut) *webhooksv1.App {
	if v == nil {
		return nil
	}
	return &webhooksv1.App{
		Uid:  lo.FromPtr(v.Uid.Get()),
		Name: v.Name,
		Id:   v.Id,
	}
}

func EndpointFromProto(v *webhooksv1.Endpoint) *svix.EndpointIn {
	if v == nil {
		return nil
	}
	m := &svix.EndpointIn{
		Url:         v.Url,
		Description: lo.EmptyableToPtr(""),
		Disabled:    lo.ToPtr(false),
		Metadata:    lo.EmptyableToPtr(map[string]string{}),
	}
	if v.Uid != "" {
		m.SetUid(v.Uid)
	}
	return m
}
func EndpointToProto(v *svix.EndpointOut) *webhooksv1.Endpoint {
	if v == nil {
		return nil
	}
	return &webhooksv1.Endpoint{
		Id:  v.Id,
		Uid: v.GetUid(),
		Url: v.Url,
	}
}

func EventTypeFromProto(v *webhooksv1.EventType) *svix.EventTypeIn {
	if v == nil {
		return nil
	}
	return &svix.EventTypeIn{
		Description: "",
		Name:        v.Name,
		Archived:    lo.ToPtr(false),
		Schemas:     map[string]map[string]any{},
	}
}
func EventTypeToProto(v *svix.EventTypeOut) *webhooksv1.EventType {
	if v == nil {
		return nil
	}
	return &webhooksv1.EventType{}
}

func MessageFromProto(v *webhooksv1.Message) *svix.MessageIn {
	if v == nil {
		return nil
	}
	return &svix.MessageIn{
		Channels:               []string{},
		EventType:              "",
		Payload:                map[string]any{},
		PayloadRetentionPeriod: lo.EmptyableToPtr[int64](0),
		Tags:                   []string{},
	}
}
func MessageToProto(v *svix.MessageOut) *webhooksv1.Message {
	if v == nil {
		return nil
	}
	return &webhooksv1.Message{
		Id: v.Id,
	}
}
