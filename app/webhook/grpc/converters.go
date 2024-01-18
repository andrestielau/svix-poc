package webhookgrpc

import (
	"encoding/json"
	"log"
	webhooksv1 "svix-poc/app/webhook/grpc/v1"

	"github.com/samber/lo"
	svix "github.com/svix/svix-webhooks/go"
)

func AppFromProto(v *webhooksv1.App) *svix.ApplicationIn {
	if v == nil {
		return nil
	}
	m := &svix.ApplicationIn{
		Metadata: lo.EmptyableToPtr(v.Metadata),
		Name:     v.Name,
	}
	if v.RateLimit > 0 {
		m.SetRateLimit(v.RateLimit)
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
		Uid:       lo.FromPtr(v.Uid.Get()),
		Name:      v.Name,
		Id:        v.Id,
		Metadata:  v.Metadata,
		RateLimit: v.GetRateLimit(),
	}
}

func EndpointFromProto(v *webhooksv1.Endpoint) *svix.EndpointIn {
	if v == nil {
		return nil
	}
	m := &svix.EndpointIn{
		Url:         v.Url,
		Channels:    v.Channels,
		FilterTypes: v.FilterTypes,
		Description: lo.EmptyableToPtr(v.Description),
		Disabled:    lo.ToPtr(v.Disabled),
		Metadata:    lo.EmptyableToPtr(v.Metadata),
	}
	if v.Uid != "" {
		m.SetUid(v.Uid)
	}
	if v.Secret != "" {
		m.SetSecret(v.Secret)
	}
	if v.Version > 0 {
		m.SetVersion(v.Version)
	}
	if v.RateLimit > 0 {
		m.SetRateLimit(v.RateLimit)
	}
	return m
}
func EndpointToProto(v *svix.EndpointOut) *webhooksv1.Endpoint {
	if v == nil {
		return nil
	}
	return &webhooksv1.Endpoint{
		Id:          v.Id,
		Uid:         v.GetUid(),
		Url:         v.Url,
		FilterTypes: v.FilterTypes,
		Description: v.Description,
		Disabled:    v.GetDisabled(),
		Version:     v.Version,
		Metadata:    v.Metadata,
		RateLimit:   v.GetRateLimit(),
		Channels:    v.Channels,
	}
}

func EventTypeFromProto(v *webhooksv1.EventType) *svix.EventTypeIn {
	if v == nil {
		return nil
	}
	m := &svix.EventTypeIn{
		Name:        v.Name,
		Description: v.Description,
		Archived:    lo.ToPtr(v.Archived),
		Schemas: lo.MapValues(v.Schemas, func(schema []byte, k string) (res map[string]any) {
			if err := json.Unmarshal(schema, &res); err != nil {
				log.Println(k + ": invalid schema: " + string(schema))
			}
			return res
		}),
	}
	if v.FeatureFlag != "" {
		m.SetFeatureFlag(v.FeatureFlag)
	}
	return m
}
func EventTypeToProto(v *svix.EventTypeOut) *webhooksv1.EventType {
	if v == nil {
		return nil
	}
	return &webhooksv1.EventType{
		Name:        v.Name,
		Description: v.Description,
		FeatureFlag: v.GetFeatureFlag(),
		Archived:    v.GetArchived(),
		Schemas: lo.MapValues(v.Schemas, func(schema map[string]any, k string) []byte {
			res, err := json.Marshal(schema)
			if err != nil {
				log.Println(k+": invalid schema: ", schema)
			}
			return res
		}),
	}
}

func MessageFromProto(v *webhooksv1.Message) *svix.MessageIn {
	if v == nil {
		return nil
	}
	var payload map[string]any
	if len(v.Payload) > 0 {
		if err := json.Unmarshal(v.Payload, &payload); err != nil {
			log.Println(err)
		}
	}
	m := &svix.MessageIn{
		EventType:              v.EventType,
		Payload:                payload,
		Tags:                   v.Tags,
		Channels:               v.Channels,
		PayloadRetentionPeriod: lo.EmptyableToPtr[int64](0),
	}
	if v.EventId != "" {
		m.SetEventId(v.EventId)
	}
	return m
}
func MessageToProto(v *svix.MessageOut) *webhooksv1.Message {
	if v == nil {
		return nil
	}
	payload, err := json.Marshal(v.Payload)
	if err != nil {
		log.Println(err)
	}
	return &webhooksv1.Message{
		Id:        v.Id,
		EventType: v.EventType,
		Payload:   payload,
		EventId:   v.GetEventId(),
		Channels:  v.Channels,
		Timestamp: v.Timestamp.String(),
	}
}
