package routerapi

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"
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
	l net.Listener
	*app.BaseActor
}

// Lifecycle

func (h *Handler) Start(ctx context.Context) (first bool, err error) {
	if first, err = h.BaseActor.Start(ctx); !first || err != nil {
		return first, err
	}
	host := string(ProvideHost())
	if h.l, err = net.Listen("tcp", host); err != nil {
		return true, err
	}
	go func() {
		if strings.HasPrefix(host, ":") {
			host = "localhost" + host
		}
		log.Println("router api listening on http://" + host)
		http.Serve(h.l, eventsv1.Handler(h))
	}()
	return true, nil
}
func (h *Handler) Stop(ctx context.Context) (last bool, err error) {
	if last, err = h.BaseActor.Stop(ctx); !last || err != nil {
		return last, err
	}
	return true, h.l.Close()
}

const DefaultLimit = 100

// CreateSubscriptions implements eventsv1.ServerInterface.
func (h *Handler) CreateSubscriptions(w http.ResponseWriter, r *http.Request) {
	if req, err := utils.JsonReq[eventsv1.NewSubscriptions](w, r); err == nil {
	} else if res, err := h.Provider.CreateSubscriptions(r.Context(), NewSubscriptionsFromJson(req)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.JsonRes(w, SubscriptionResultsToJson(res, nil))
	}
}

// DeleteSubscription implements eventsv1.ServerInterface.
func (h *Handler) DeleteSubscription(w http.ResponseWriter, r *http.Request, subscriptionId string) {
	if res, err := h.Provider.DeleteSubscription(r.Context(), []string{subscriptionId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if res.RowsAffected() == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// GetProvider implements eventsv1.ServerInterface.
func (h *Handler) GetProvider(w http.ResponseWriter, r *http.Request, providerId string, params eventsv1.GetProviderParams) {
	if res, err := h.Provider.GetProviders(r.Context(), []string{providerId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if len(res) == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	} else {
		utils.JsonRes(w, ProviderToJson(res[0]))
	}
}

// GetSubscription implements eventsv1.ServerInterface.
func (h *Handler) GetSubscription(w http.ResponseWriter, r *http.Request, subscriptionId string, params eventsv1.GetSubscriptionParams) {
	if res, err := h.Provider.GetSubscriptions(r.Context(), []string{subscriptionId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if len(res) == 0 {
		http.Error(w, "not found", http.StatusNotFound)
	} else {
		utils.JsonRes(w, SubscriptionToJson(res[0]))
	}
}

// ListEventTypes implements eventsv1.ServerInterface.
func (h *Handler) ListEventTypes(w http.ResponseWriter, r *http.Request, params eventsv1.ListEventTypesParams) {
	if res, err := h.Provider.ListNotificationTypes(r.Context(), time.Time{}, int(lo.FromPtrOr(params.Limit, DefaultLimit))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.JsonRes(w, ListEventTypesToJson(res))
	}
}

// ListProviders implements eventsv1.ServerInterface.
func (h *Handler) ListProviders(w http.ResponseWriter, r *http.Request, params eventsv1.ListProvidersParams) {
	if res, err := h.Provider.ListProviders(r.Context(), time.Time{}, int(lo.FromPtrOr(params.Limit, DefaultLimit))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.JsonRes(w, ListProvidersToJson(res))
	}
}

// ListSubscriptions implements eventsv1.ServerInterface.
func (h *Handler) ListSubscriptions(w http.ResponseWriter, r *http.Request, params eventsv1.ListSubscriptionsParams) {
	if res, err := h.Provider.ListSubscriptions(r.Context(), time.Time{}, int(lo.FromPtrOr(params.Limit, DefaultLimit))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		utils.JsonRes(w, ListSubscriptionsToJson(res))
	}
}
