package routerapi

import (
	eventsv1 "svix-poc/app/router/api/v1"
	"svix-poc/app/router/repo"

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
			NotificationType: item.EventTypeId,
		}
	})
}
func ProviderToJson(in string) eventsv1.Provider {
	return eventsv1.Provider{
		Id: in,
	}
}
func SubscriptionToJson(in repo.GetSubscriptionsRow) eventsv1.Subscription {
	return eventsv1.Subscription{
		CreatedAt: in.CreatedAt,
		Id:        in.Uid,
	}
}
func ListEventTypesToJson(in []repo.ListNotificationTypesRow) eventsv1.EventTypeList {
	return eventsv1.EventTypeList{
		Data: lo.Map(in, func(item repo.ListNotificationTypesRow, _ int) eventsv1.EventType {
			return eventsv1.EventType{
				Id: item.ID,
			}
		}),
		Cursor: eventsv1.PageInfo{},
	}
}
func ListProvidersToJson(in []repo.ListProvidersRow) eventsv1.ProviderList {
	return eventsv1.ProviderList{
		Data: lo.Map(in, func(item repo.ListProvidersRow, _ int) eventsv1.Provider {
			return eventsv1.Provider{
				Id: item.ID,
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
