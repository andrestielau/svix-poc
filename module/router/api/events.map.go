package api

import (
	eventsv1 "svix-poc/module/router/api/v1"
	"svix-poc/module/router/repo"

	"github.com/samber/lo"
)

// TODO: finish mapping

func SubscriptionResultsToJson(in []repo.CreateSubscriptionsRow, errs []eventsv1.Error) eventsv1.SubscriptionResult {
	return eventsv1.SubscriptionResult{
		Data: lo.Map(in, func(item repo.CreateSubscriptionsRow, _ int) eventsv1.Subscription {
			return eventsv1.Subscription{
				CreatedAt: item.CreatedAt,
				Id:        item.Uid,
			}
		}),
		Errors: &errs,
	}
}

func NewSubscriptionsFromJson(in []eventsv1.NewSubscription) []repo.NewSubscription {
	return lo.Map(in, func(item eventsv1.NewSubscription, _ int) repo.NewSubscription {
		return repo.NewSubscription{
			NotificationTypeUid: item.NotificationTypeId,
		}
	})
}

func ProviderToJson(in repo.GetProvidersRow) eventsv1.Provider {
	return eventsv1.Provider{
		Id:   in.Uid,
		Name: in.Name,
	}
}

func SubscriptionToJson(in repo.GetSubscriptionsRow) eventsv1.Subscription {
	return eventsv1.Subscription{
		CreatedAt: in.CreatedAt,
		Id:        in.Uid,
	}
}

func ListNotificationTypesToJson(in []repo.ListNotificationTypesRow) eventsv1.NotificationTypeList {
	return eventsv1.NotificationTypeList{
		Data: lo.Map(in, func(item repo.ListNotificationTypesRow, _ int) eventsv1.NotificationType {
			return eventsv1.NotificationType{
				Id:   item.Uid,
				Name: item.Name,
			}
		}),
		Cursor: eventsv1.PageInfo{},
	}
}
func ListProvidersToJson(in []repo.ListProvidersRow) eventsv1.ProviderList {
	return eventsv1.ProviderList{
		Data: lo.Map(in, func(item repo.ListProvidersRow, _ int) eventsv1.Provider {
			return eventsv1.Provider{
				Id:   item.Uid,
				Name: item.Name,
			}
		}),
		Cursor: eventsv1.PageInfo{},
	}
}
func ListSubscriptionsToJson(in []repo.ListSubscriptionsRow) eventsv1.SubscriptionList {
	return eventsv1.SubscriptionList{
		Data: lo.Map(in, func(item repo.ListSubscriptionsRow, _ int) eventsv1.Subscription {
			return eventsv1.Subscription{
				CreatedAt: item.CreatedAt,
				Id:        item.Uid,
			}
		}),
		Cursor: eventsv1.PageInfo{},
	}
}
