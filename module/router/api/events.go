package api

import (
	"net/http"
	eventsv1 "svix-poc/module/router/api/v1"
	"svix-poc/module/router/repo"
	"svix-poc/x"
	"time"

	"github.com/samber/lo"
)

type Events struct {
	*repo.Repository
}

const DefaultLimit = 100

// CreateSubscriptions implements eventsv1.ServerInterface.
func (h *Events) CreateSubscriptions(w http.ResponseWriter, r *http.Request) {
	if req, err := x.JsonReq[eventsv1.NewSubscriptions](w, r); err == nil {
	} else if res, err := h.Repository.CreateSubscriptions(r.Context(), NewSubscriptionsFromJson(req)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		x.JsonRes(w, SubscriptionResultsToJson(res, nil))
	}
}

// DeleteSubscription implements eventsv1.ServerInterface.
func (h *Events) DeleteSubscription(w http.ResponseWriter, r *http.Request, subscriptionId string) {
	if res, err := h.Repository.DeleteSubscription(r.Context(), []string{subscriptionId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if res.RowsAffected() == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// GetProvider implements eventsv1.ServerInterface.
func (h *Events) GetProvider(w http.ResponseWriter, r *http.Request, providerId string, params eventsv1.GetProviderParams) {
	if res, err := h.Repository.GetProviders(r.Context(), []string{providerId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if len(res) == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	} else {
		x.JsonRes(w, ProviderToJson(res[0]))
	}
}

// GetSubscription implements eventsv1.ServerInterface.
func (h *Events) GetSubscription(w http.ResponseWriter, r *http.Request, subscriptionId string, params eventsv1.GetSubscriptionParams) {
	if res, err := h.Repository.GetSubscriptions(r.Context(), []string{subscriptionId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if len(res) == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	} else {
		x.JsonRes(w, SubscriptionToJson(res[0]))
	}
}

// ListNotificationTypes implements eventsv1.ServerInterface.
func (h *Events) ListNotificationTypes(w http.ResponseWriter, r *http.Request, params eventsv1.ListNotificationTypesParams) {
	if res, err := h.Repository.ListNotificationTypes(r.Context(), time.Time{}, int(lo.FromPtrOr(params.Limit, DefaultLimit))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		x.JsonRes(w, ListNotificationTypesToJson(res))
	}
}

// ListProviders implements eventsv1.ServerInterface.
func (h *Events) ListProviders(w http.ResponseWriter, r *http.Request, params eventsv1.ListProvidersParams) {
	if res, err := h.Repository.ListProviders(r.Context(), time.Time{}, int(lo.FromPtrOr(params.Limit, DefaultLimit))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		x.JsonRes(w, ListProvidersToJson(res))
	}
}

// ListSubscriptions implements eventsv1.ServerInterface.
func (h *Events) ListSubscriptions(w http.ResponseWriter, r *http.Request, params eventsv1.ListSubscriptionsParams) {
	if res, err := h.Repository.ListSubscriptions(r.Context(), time.Time{}, int(lo.FromPtrOr(params.Limit, DefaultLimit))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		x.JsonRes(w, ListSubscriptionsToJson(res))
	}
}
