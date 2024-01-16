package routerapi

import (
	"net/http"
	"svix-poc/app/router"
	eventsv1 "svix-poc/app/router/api/v1"
	"svix-poc/package/app"
	"svix-poc/package/utils"
	"time"

	"github.com/samber/lo"
)

var _ eventsv1.ServerInterface = &Handler{}

type Handler struct {
	router.Dependencies
	server *http.Server
	*app.BaseActor
}

const DefaultLimit = 100

// CreateSubscriptions implements eventsv1.ServerInterface.
func (h *Handler) CreateSubscriptions(w http.ResponseWriter, r *http.Request) {
	if req, err := utils.JsonReq[eventsv1.NewSubscriptions](w, r); err == nil {
	} else if res, err := h.Repository.CreateSubscriptions(r.Context(), NewSubscriptionsFromJson(req)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.JsonRes(w, SubscriptionResultsToJson(res, nil))
	}
}

// DeleteSubscription implements eventsv1.ServerInterface.
func (h *Handler) DeleteSubscription(w http.ResponseWriter, r *http.Request, subscriptionId string) {
	if res, err := h.Repository.DeleteSubscription(r.Context(), []string{subscriptionId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if res.RowsAffected() == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// GetProvider implements eventsv1.ServerInterface.
func (h *Handler) GetProvider(w http.ResponseWriter, r *http.Request, providerId string, params eventsv1.GetProviderParams) {
	if res, err := h.Repository.GetProviders(r.Context(), []string{providerId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if len(res) == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	} else {
		utils.JsonRes(w, ProviderToJson(res[0]))
	}
}

// GetSubscription implements eventsv1.ServerInterface.
func (h *Handler) GetSubscription(w http.ResponseWriter, r *http.Request, subscriptionId string, params eventsv1.GetSubscriptionParams) {
	if res, err := h.Repository.GetSubscriptions(r.Context(), []string{subscriptionId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if len(res) == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	} else {
		utils.JsonRes(w, SubscriptionToJson(res[0]))
	}
}

// ListEventTypes implements eventsv1.ServerInterface.
func (h *Handler) ListEventTypes(w http.ResponseWriter, r *http.Request, params eventsv1.ListEventTypesParams) {
	if res, err := h.Repository.ListNotificationTypes(r.Context(), time.Time{}, int(lo.FromPtrOr(params.Limit, DefaultLimit))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.JsonRes(w, ListEventTypesToJson(res))
	}
}

// ListProviders implements eventsv1.ServerInterface.
func (h *Handler) ListProviders(w http.ResponseWriter, r *http.Request, params eventsv1.ListProvidersParams) {
	if res, err := h.Repository.ListProviders(r.Context(), time.Time{}, int(lo.FromPtrOr(params.Limit, DefaultLimit))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.JsonRes(w, ListProvidersToJson(res))
	}
}

// ListSubscriptions implements eventsv1.ServerInterface.
func (h *Handler) ListSubscriptions(w http.ResponseWriter, r *http.Request, params eventsv1.ListSubscriptionsParams) {
	if res, err := h.Repository.ListSubscriptions(r.Context(), time.Time{}, int(lo.FromPtrOr(params.Limit, DefaultLimit))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.JsonRes(w, ListSubscriptionsToJson(res))
	}
}
