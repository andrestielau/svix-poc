package webhookapi

import (
	"net/http"
	"svix-poc/app/webhook"
	webhookv1 "svix-poc/app/webhook/api/v1"
	"svix-poc/lib/utils"

	svix "github.com/svix/svix-webhooks/go"
)

type Handler struct {
	webhook.Dependencies
}

// Create

// CreateEndpoints implements webhookv1.ServerInterface.
func (h *Handler) CreateEndpoints(w http.ResponseWriter, r *http.Request, params webhookv1.CreateEndpointsParams) {
	if req, err := utils.JsonReq[webhookv1.NewEndpoints](w, r); err == nil {
		utils.JsonRes(w, EndpointResultToJson(Batch(NewEndpointsFromJson(req), func(in *svix.EndpointIn) (*svix.EndpointOut, error) {
			return h.Endpoint.Create(r.Context(), params.TenantId, in)
		})))
	}
}

// CreateMessages implements webhookv1.ServerInterface.
func (h *Handler) CreateMessages(w http.ResponseWriter, r *http.Request, params webhookv1.CreateMessagesParams) {
	if req, err := utils.JsonReq[webhookv1.NewMessages](w, r); err == nil {
		utils.JsonRes(w, MessageResultToJson(Batch(NewMessagesFromJson(req), func(in *svix.MessageIn) (*svix.MessageOut, error) {
			return h.Message.Create(r.Context(), params.TenantId, in)
		})))
	}
}

// Delete

// DeleteEndpoint implements webhookv1.ServerInterface.
func (h *Handler) DeleteEndpoint(w http.ResponseWriter, r *http.Request, endpointId string, params webhookv1.DeleteEndpointParams) {
	if SetStatus(w, h.Endpoint.Delete(r.Context(), params.TenantId, endpointId)) {
		w.WriteHeader(http.StatusNoContent)
	}
}

// ExpungeMessage implements webhookv1.ServerInterface.
func (h *Handler) ExpungeMessage(w http.ResponseWriter, r *http.Request, msgId string, attemptId string, params webhookv1.ExpungeMessageParams) {
	if SetStatus(w, h.MessageAttempt.ExpungeContent(r.Context(), params.TenantId, msgId, params.TenantId)) {
		w.WriteHeader(http.StatusNoContent)
	}
}

// Get

// GetEndpoint implements webhookv1.ServerInterface.
func (h *Handler) GetEndpoint(w http.ResponseWriter, r *http.Request, endpointId string, params webhookv1.GetEndpointParams) {
	if res, err := h.Endpoint.Get(r.Context(), params.TenantId, endpointId); SetStatus(w, err) {
		utils.JsonRes(w, EndpointToJson(*res))
	}
}

// GetMessage implements webhookv1.ServerInterface.
func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request, msgId string, params webhookv1.GetMessageParams) {
	if res, err := h.Message.Get(r.Context(), params.TenantId, msgId); SetStatus(w, err) {
		utils.JsonRes(w, MessageToJson(*res))
	}
}

// GetMessageAttempt implements webhookv1.ServerInterface.
func (h *Handler) GetMessageAttempt(w http.ResponseWriter, r *http.Request, msgId string, attemptId string, params webhookv1.GetMessageAttemptParams) {
	if res, err := h.MessageAttempt.Get(r.Context(), params.TenantId, msgId, attemptId); SetStatus(w, err) {
		utils.JsonRes(w, AttemptToJson(*res))
	}
}

// List

// ListEndpointAttempts implements webhookv1.ServerInterface.
func (h *Handler) ListEndpointAttempts(w http.ResponseWriter, r *http.Request, endpointId string, params webhookv1.ListEndpointAttemptsParams) {
	if res, err := h.MessageAttempt.ListByEndpoint(r.Context(), params.TenantId, endpointId, &svix.MessageAttemptListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		utils.JsonRes(w, ListAttemptsToJson(res))
	}
}

// ListEndpointMessageAttempts implements webhookv1.ServerInterface.
func (h *Handler) ListEndpointMessageAttempts(w http.ResponseWriter, r *http.Request, endpointId string, msgId string, params webhookv1.ListEndpointMessageAttemptsParams) {
	if res, err := h.MessageAttempt.ListAttemptsForEndpoint(r.Context(), params.TenantId, msgId, endpointId, &svix.MessageAttemptListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		utils.JsonRes(w, ListEndpointAttemptsToJson(res))
	}
}

// ListEndpointMessages implements webhookv1.ServerInterface.
func (h *Handler) ListEndpointMessages(w http.ResponseWriter, r *http.Request, endpointId string, params webhookv1.ListEndpointMessagesParams) {
	if res, err := h.MessageAttempt.ListAttemptedMessages(r.Context(), params.TenantId, endpointId, &svix.MessageAttemptListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		utils.JsonRes(w, ListEndpointMessagesToJson(res))
	}
}

// ListEndpoints implements webhookv1.ServerInterface.
func (h *Handler) ListEndpoints(w http.ResponseWriter, r *http.Request, params webhookv1.ListEndpointsParams) {
	if res, err := h.Endpoint.List(r.Context(), params.TenantId, &svix.EndpointListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		utils.JsonRes(w, ListEndpointsToJson(res))
	}
}

// ListMessageAttempts implements webhookv1.ServerInterface.
func (h *Handler) ListMessageAttempts(w http.ResponseWriter, r *http.Request, msgId string, params webhookv1.ListMessageAttemptsParams) {
	if res, err := h.MessageAttempt.ListByMsg(r.Context(), params.TenantId, msgId, &svix.MessageAttemptListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		utils.JsonRes(w, ListAttemptsToJson(res))
	}
}

// ListMessages implements webhookv1.ServerInterface.
func (h *Handler) ListMessages(w http.ResponseWriter, r *http.Request, params webhookv1.ListMessagesParams) {
	if res, err := h.Message.List(r.Context(), params.TenantId, &svix.MessageListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		utils.JsonRes(w, ListMessagesToJson(res))
	}
}

// ListTypes implements webhookv1.ServerInterface.
func (h *Handler) ListTypes(w http.ResponseWriter, r *http.Request, params webhookv1.ListTypesParams) {
	if res, err := h.EventType.List(r.Context(), &svix.EventTypeListOptions{
		Iterator: params.Cursor,
		Limit:    params.Limit,
	}); SetStatus(w, err) {
		utils.JsonRes(w, ListTypesToJson(res))
	}
}
