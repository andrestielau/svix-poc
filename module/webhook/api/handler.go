package webhookapi

import (
	"net/http"
	eventsv1 "svix-poc/module/router/api/v1"
	"svix-poc/x"

	svix "github.com/svix/svix-webhooks/go"
)

type WebHooks struct {
	*svix.Svix
}

// Create

// CreateWebhookEndpoints implements eventsv1.ServerInterface.
func (h *WebHooks) CreateWebhookEndpoints(w http.ResponseWriter, r *http.Request, params eventsv1.CreateWebhookEndpointsParams) {
	if req, err := x.JsonReq[eventsv1.NewWebHookEndpoints](w, r); err == nil {
		x.JsonRes(w, WebHookEndpointResultToJson(Batch(NewWebHookEndpointsFromJson(req), func(in *svix.EndpointIn) (*svix.EndpointOut, error) {
			return h.Endpoint.Create(r.Context(), params.TenantId, in)
		})))
	}
}

// CreateWebhookMessages implements eventsv1.ServerInterface.
func (h *WebHooks) CreateWebhookMessages(w http.ResponseWriter, r *http.Request, params eventsv1.CreateWebhookMessagesParams) {
	if req, err := x.JsonReq[eventsv1.NewWebHookMessages](w, r); err == nil {
		x.JsonRes(w, WebHookMessageResultToJson(Batch(NewWebHookMessagesFromJson(req), func(in *svix.MessageIn) (*svix.MessageOut, error) {
			return h.Message.Create(r.Context(), params.TenantId, in)
		})))
	}
}

// Delete

// DeleteWebhookEndpoint implements eventsv1.ServerInterface.
func (h *WebHooks) DeleteWebhookEndpoint(w http.ResponseWriter, r *http.Request, endpointId string, params eventsv1.DeleteWebhookEndpointParams) {
	if SetStatus(w, h.Endpoint.Delete(r.Context(), params.TenantId, endpointId)) {
		w.WriteHeader(http.StatusNoContent)
	}
}

// ExpungeWebhookMessage implements eventsv1.ServerInterface.
func (h *WebHooks) ExpungeWebhookMessage(w http.ResponseWriter, r *http.Request, msgId string, attemptId string, params eventsv1.ExpungeWebhookMessageParams) {
	if SetStatus(w, h.MessageAttempt.ExpungeContent(r.Context(), params.TenantId, msgId, params.TenantId)) {
		w.WriteHeader(http.StatusNoContent)
	}
}

// Get

// GetWebhookEndpoint implements eventsv1.ServerInterface.
func (h *WebHooks) GetWebhookEndpoint(w http.ResponseWriter, r *http.Request, endpointId string, params eventsv1.GetWebhookEndpointParams) {
	if res, err := h.Endpoint.Get(r.Context(), params.TenantId, endpointId); SetStatus(w, err) {
		x.JsonRes(w, WebHookEndpointToJson(*res))
	}
}

// GetWebhookMessage implements eventsv1.ServerInterface.
func (h *WebHooks) GetWebhookMessage(w http.ResponseWriter, r *http.Request, msgId string, params eventsv1.GetWebhookMessageParams) {
	if res, err := h.Message.Get(r.Context(), params.TenantId, msgId); SetStatus(w, err) {
		x.JsonRes(w, WebHookMessageToJson(*res))
	}
}

// GetWebhookMessageAttempt implements eventsv1.ServerInterface.
func (h *WebHooks) GetWebhookMessageAttempt(w http.ResponseWriter, r *http.Request, msgId string, attemptId string, params eventsv1.GetWebhookMessageAttemptParams) {
	if res, err := h.MessageAttempt.Get(r.Context(), params.TenantId, msgId, attemptId); SetStatus(w, err) {
		x.JsonRes(w, WebHookAttemptToJson(*res))
	}
}

// List

// ListWebhookEndpointAttempts implements eventsv1.ServerInterface.
func (h *WebHooks) ListWebhookEndpointAttempts(w http.ResponseWriter, r *http.Request, endpointId string, params eventsv1.ListWebhookEndpointAttemptsParams) {
	if res, err := h.MessageAttempt.ListByEndpoint(r.Context(), params.TenantId, endpointId, &svix.MessageAttemptListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		x.JsonRes(w, ListWebHookAttemptsToJson(res))
	}
}

// ListWebhookEndpointMessageAttempts implements eventsv1.ServerInterface.
func (h *WebHooks) ListWebhookEndpointMessageAttempts(w http.ResponseWriter, r *http.Request, endpointId string, msgId string, params eventsv1.ListWebhookEndpointMessageAttemptsParams) {
	if res, err := h.MessageAttempt.ListAttemptsForEndpoint(r.Context(), params.TenantId, msgId, endpointId, &svix.MessageAttemptListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		x.JsonRes(w, ListWebHookEndpointAttemptsToJson(res))
	}
}

// ListWebhookEndpointMessages implements eventsv1.ServerInterface.
func (h *WebHooks) ListWebhookEndpointMessages(w http.ResponseWriter, r *http.Request, endpointId string, params eventsv1.ListWebhookEndpointMessagesParams) {
	if res, err := h.MessageAttempt.ListAttemptedMessages(r.Context(), params.TenantId, endpointId, &svix.MessageAttemptListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		x.JsonRes(w, ListWebHookEndpointMessagesToJson(res))
	}
}

// ListWebhookEndpoints implements eventsv1.ServerInterface.
func (h *WebHooks) ListWebhookEndpoints(w http.ResponseWriter, r *http.Request, params eventsv1.ListWebhookEndpointsParams) {
	if res, err := h.Endpoint.List(r.Context(), params.TenantId, &svix.EndpointListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		x.JsonRes(w, ListWebHookEndpointsToJson(res))
	}
}

// ListWebhookMessageAttempts implements eventsv1.ServerInterface.
func (h *WebHooks) ListWebhookMessageAttempts(w http.ResponseWriter, r *http.Request, msgId string, params eventsv1.ListWebhookMessageAttemptsParams) {
	if res, err := h.MessageAttempt.ListByMsg(r.Context(), params.TenantId, msgId, &svix.MessageAttemptListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		x.JsonRes(w, ListWebHookAttemptsToJson(res))
	}
}

// ListWebhookMessages implements eventsv1.ServerInterface.
func (h *WebHooks) ListWebhookMessages(w http.ResponseWriter, r *http.Request, params eventsv1.ListWebhookMessagesParams) {
	if res, err := h.Message.List(r.Context(), params.TenantId, &svix.MessageListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		x.JsonRes(w, ListWebHookMessagesToJson(res))
	}
}

// ListWebhookTypes implements eventsv1.ServerInterface.
func (h *WebHooks) ListWebhookTypes(w http.ResponseWriter, r *http.Request, params eventsv1.ListWebhookTypesParams) {
	if res, err := h.EventType.List(r.Context(), &svix.EventTypeListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		x.JsonRes(w, ListWebHookTypesToJson(res))
	}
}
