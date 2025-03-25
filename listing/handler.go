package listing

import (
	"encoding/json"
	"github.com/nasermirzaei89/problem"
	"github.com/nasermirzaei89/respond"
	"github.com/pkg/errors"
	"net/http"
)

type Handler struct {
	mux        *http.ServeMux
	listingSvc Service
}

var _ http.Handler = (*Handler)(nil)

func NewHandler(listingSvc Service) *Handler {
	h := &Handler{
		mux:        http.NewServeMux(),
		listingSvc: listingSvc,
	}

	h.registerRoutes()

	return h
}

func (h *Handler) registerRoutes() {
	h.mux.Handle("GET /locations", h.listLocationsHandler())
	h.mux.Handle("GET /published-items", h.listPublishedItemsHandler())
	h.mux.Handle("GET /my-items", h.listMyItemsHandler())
	h.mux.Handle("GET /items/{itemId}", h.getItemHandler())
	h.mux.Handle("GET /published-items/{itemId}", h.getPublishedItemHandler())
	h.mux.Handle("PUT /items/{itemId}/send-for-publish", h.sendItemForPublishHandler())
	h.mux.Handle("PUT /items/{itemId}/publish", h.publishItemHandler())
	h.mux.Handle("PUT /items/{itemId}/archive", h.archiveItemHandler())
	h.mux.Handle("DELETE /items/{itemId}", h.deleteItemHandler())
	h.mux.Handle("POST /items", h.createItemHandler())
	h.mux.Handle("PUT /items/{itemId}", h.updateItemHandler())
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// listLocationsHandler
//
//	@Summary	List locations
//	@Tags		listing
//	@Return		application/json
//	@Success	200	{object}	ListLocationsResponse
//	@Router		/listing/locations [get]
func (h *Handler) listLocationsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := h.listingSvc.ListLocations(r.Context())
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on listing locations")))

			return
		}

		respond.Done(w, r, res)
	})
}

// listPublishedItemsHandler
//
//	@Summary	List published items
//	@Tags		listing
//	@Return		application/json
//	@Success	200	{object}	ListItemsResponse
//	@Router		/listing/published-items [get]
func (h *Handler) listPublishedItemsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")

		res, err := h.listingSvc.ListPublishedItems(r.Context(), q)
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on listing published items")))

			return
		}

		respond.Done(w, r, res)
	})
}

// listMyItemsHandler
//
//	@Summary	List my items
//	@Tags		listing
//	@Security	OAuth2Implicit
//	@Return		application/json
//	@Success	200	{object}	ListItemsResponse
//	@Router		/listing/my-items [get]
func (h *Handler) listMyItemsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")

		res, err := h.listingSvc.ListMyItems(r.Context(), q)
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on listing my items")))

			return
		}

		respond.Done(w, r, res)
	})
}

// getItemHandler
//
//	@Summary	Get item
//	@Tags		listing
//	@Security	OAuth2Implicit
//	@Param		itemId	path	string	true	"Item id"
//	@Return		application/json
//	@Success	200	{object}	Item
//	@Router		/listing/items/{itemId} [get]
func (h *Handler) getItemHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemID := r.PathValue("itemId")

		item, err := h.listingSvc.GetItem(r.Context(), itemID)
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on getting item")))

			return
		}

		respond.Done(w, r, item)
	})
}

// getPublishedItemHandler
//
//	@Summary	Get published item
//	@Tags		listing
//	@Param		itemId	path	string	true	"Item id"
//	@Return		application/json
//	@Success	200	{object}	Item
//	@Router		/listing/published-items/{itemId} [get]
func (h *Handler) getPublishedItemHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemID := r.PathValue("itemId")

		item, err := h.listingSvc.GetPublishedItem(r.Context(), itemID)
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on getting published item")))

			return
		}

		respond.Done(w, r, item)
	})
}

// sendItemForPublishHandler
//
//	@Summary	Send item for publish
//	@Tags		listing
//	@Security	OAuth2Implicit
//	@Param		itemId	path	string	true	"Item id"
//	@Success	204
//	@Router		/listing/items/{itemId}/send-for-publish [put]
func (h *Handler) sendItemForPublishHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemID := r.PathValue("itemId")

		err := h.listingSvc.SendItemForPublish(r.Context(), itemID)
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on sending item for publish")))

			return
		}

		respond.Done(w, r, nil)
	})
}

// publishItemHandler
//
//	@Summary	Publish item
//	@Tags		listing
//	@Security	OAuth2Implicit
//	@Param		itemId	path	string	true	"Item id"
//	@Success	204
//	@Router		/listing/items/{itemId}/publish [put]
func (h *Handler) publishItemHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemID := r.PathValue("itemId")

		err := h.listingSvc.PublishItem(r.Context(), itemID)
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on publishing item")))

			return
		}

		respond.Done(w, r, nil)
	})
}

// archiveItemHandler
//
//	@Summary	Publish item
//	@Tags		listing
//	@Security	OAuth2Implicit
//	@Param		itemId	path	string	true	"Item id"
//	@Success	204
//	@Router		/listing/items/{itemId}/archive [put]
func (h *Handler) archiveItemHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemID := r.PathValue("itemId")

		err := h.listingSvc.ArchiveItem(r.Context(), itemID)
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on archiving item")))

			return
		}

		respond.Done(w, r, nil)
	})
}

// deleteItemHandler
//
//	@Summary	Delete item
//	@Tags		listing
//	@Security	OAuth2Implicit
//	@Param		itemId	path	string	true	"Item id"
//	@Success	204
//	@Router		/listing/items/{itemId} [delete]
func (h *Handler) deleteItemHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemID := r.PathValue("itemId")

		err := h.listingSvc.DeleteItem(r.Context(), itemID)
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on deleting item")))

			return
		}

		respond.Done(w, r, nil)
	})
}

// createItemHandler
//
//	@Summary	Create new item
//	@Tags		listing
//	@Security	OAuth2Implicit
//	@Param		req	body		CreateItemRequest	true	"Request body"
//	@Success	201	{object}	Item
//	@Router		/listing/items [post]
func (h *Handler) createItemHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req CreateItemRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			respond.Done(w, r, problem.BadRequest("invalid request"))

			return
		}

		item, err := h.listingSvc.CreateItem(r.Context(), &req)
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on creating item")))

			return
		}

		res := struct {
			respond.WithStatusCreated
			*Item
		}{
			WithStatusCreated: respond.WithStatusCreated{},
			Item:              item,
		}

		respond.Done(w, r, res)
	})
}

// updateItemHandler
//
//	@Summary	Update item
//	@Tags		listing
//	@Security	OAuth2Implicit
//	@Param		itemId	path	string				true	"Item id"
//	@Param		req		body	UpdateItemRequest	true	"Request body"
//	@Success	204
//	@Router		/listing/items/{itemId} [put]
func (h *Handler) updateItemHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemID := r.PathValue("itemId")

		var req UpdateItemRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			respond.Done(w, r, problem.BadRequest("invalid request"))

			return
		}

		err = h.listingSvc.UpdateItem(r.Context(), itemID, &req)
		if err != nil {
			respond.Done(w, r, problem.InternalServerError(errors.Wrap(err, "error on updating item")))

			return
		}

		respond.Done(w, r, nil)
	})
}
