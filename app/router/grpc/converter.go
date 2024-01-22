package routergrpc

import (
	"strconv"
	eventsv1 "svix-poc/app/router/grpc/v1"
	"svix-poc/app/router/repo"
	"time"

	"github.com/jackc/pgtype"
	"github.com/samber/lo"
)

func EventTypesFromProto(in []*eventsv1.EventType) []repo.NewEventType {
	return lo.Map(in, func(item *eventsv1.EventType, _ int) repo.NewEventType {
		return repo.NewEventType{
			ID: item.Id,
		}
	})
}
func NewEventTypesToProto(in []repo.CreateEventTypesRow) []*eventsv1.EventType {
	return lo.Map(in, func(item repo.CreateEventTypesRow, _ int) *eventsv1.EventType {
		return &eventsv1.EventType{
			Id: item.ID,
		}
	})
}
func NotificationTypesFromProto(in []*eventsv1.NotificationType) []repo.NewNotificationType {
	return lo.Map(in, func(item *eventsv1.NotificationType, _ int) repo.NewNotificationType {
		return repo.NewNotificationType{
			ID: item.Id,
		}
	})
}
func NewNotificationTypesToProto(in []repo.CreateNotificationTypesRow) []*eventsv1.NotificationType {
	return lo.Map(in, func(item repo.CreateNotificationTypesRow, _ int) *eventsv1.NotificationType {
		return &eventsv1.NotificationType{
			Id: item.ID,
		}
	})
}
func SubscriptionsFromProto(in []*eventsv1.Subscription) []repo.NewSubscription {
	return lo.Map(in, func(item *eventsv1.Subscription, _ int) repo.NewSubscription {
		return repo.NewSubscription{
			NotificationType: item.Notification,
			TenantID:         item.TenantId,
			Metadata: pgtype.JSONB{
				Status: pgtype.Present,
				Bytes:  item.Metadata,
			},
		}
	})
}
func NewSubscriptionsToProto(in []repo.CreateSubscriptionsRow) []*eventsv1.Subscription {
	return lo.Map(in, func(item repo.CreateSubscriptionsRow, _ int) *eventsv1.Subscription {
		return &eventsv1.Subscription{
			Id:       item.Uid,
			TenantId: item.TenantID,
			Metadata: item.Metadata.Bytes,
		}
	})
}
func GotEventTypesToProto(in []repo.GetEventTypesRow) map[string]*eventsv1.EventType {
	return lo.SliceToMap(in, func(item repo.GetEventTypesRow) (string, *eventsv1.EventType) {
		return item.ID, &eventsv1.EventType{
			Id: item.ID,
		}
	})
}
func GotNotificationTypesToProto(in []repo.GetNotificationTypesRow) map[string]*eventsv1.NotificationType {
	return lo.SliceToMap(in, func(item repo.GetNotificationTypesRow) (string, *eventsv1.NotificationType) {
		return item.ID, &eventsv1.NotificationType{
			Id: item.ID,
		}
	})
}
func GotProvidersToProto(in []string) map[string]*eventsv1.Provider {
	return lo.SliceToMap(in, func(item string) (string, *eventsv1.Provider) {
		return item, &eventsv1.Provider{
			Id: item,
		}
	})
}
func GotSubscriptionsToProto(in []repo.GetSubscriptionsRow) map[string]*eventsv1.Subscription {
	return lo.SliceToMap(in, func(item repo.GetSubscriptionsRow) (string, *eventsv1.Subscription) {
		return item.Uid, &eventsv1.Subscription{
			Id:           item.Uid,
			TenantId:     item.TenantID,
			Notification: item.NotificationType,
			Metadata:     item.Metadata.Bytes,
		}
	})
}
func EventTypeListToProto(in []repo.ListEventTypesRow) []*eventsv1.EventType {
	return lo.Map(in, func(item repo.ListEventTypesRow, _ int) *eventsv1.EventType {
		return &eventsv1.EventType{
			Id: item.ID,
			// TODO: convert times
		}
	})
}
func NotificationTypeListToProto(in []repo.ListNotificationTypesRow) []*eventsv1.NotificationType {
	return lo.Map(in, func(item repo.ListNotificationTypesRow, _ int) *eventsv1.NotificationType {
		return &eventsv1.NotificationType{
			Id: item.ID,
			// TODO: convert times
		}
	})
}
func ProviderListToProto(in []repo.ListProvidersRow) []*eventsv1.Provider {
	return lo.Map(in, func(item repo.ListProvidersRow, _ int) *eventsv1.Provider {
		return &eventsv1.Provider{
			Id: item.ID,
		}
	})
}
func SubscriptionListToProto(in []repo.ListSubscriptionsRow) []*eventsv1.Subscription {
	return lo.Map(in, func(item repo.ListSubscriptionsRow, _ int) *eventsv1.Subscription {
		return &eventsv1.Subscription{
			Id: item.Uid,
		}
	})
}

func Cursor(p *eventsv1.PageRequest) (t time.Time) {
	if p != nil {
		cursor, _ := strconv.Atoi(p.Cursor)
		t = t.Add(time.Duration(cursor) * time.Nanosecond)
	}
	return t
}
