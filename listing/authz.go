package listing

import (
	"context"
	"github.com/nasermirzaei89/ketabdoozak/sharedcontext"
	"github.com/nasermirzaei89/services/authorization"
	"github.com/pkg/errors"
)

const (
	ActionListPublishedItems = "listPublishedItems"
	ActionListMyItems        = "listMyItems"
	ActionGetItem            = "getItem"
	ActionGetPublishedItem   = "getPublishedItem"
	ActionSendItemForPublish = "sendItemForPublish"
	ActionPublishItem        = "publishItem"
	ActionArchiveItem        = "archiveItem"
	ActionDeleteItem         = "deleteItem"
	ActionCreateItem         = "createItem"
	ActionUpdateItem         = "updateItem"
)

type AuthorizationMiddleware struct {
	next     Service
	authzSvc *authorization.Service
}

var _ Service = (*AuthorizationMiddleware)(nil)

func NewAuthorizationMiddleware(next Service, authzSvc *authorization.Service) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{
		next:     next,
		authzSvc: authzSvc,
	}
}

func (mw *AuthorizationMiddleware) checkAccess(ctx context.Context, object, action string) error {
	err := mw.authzSvc.CheckAccess(ctx, authorization.CheckAccessRequest{
		Subject: sharedcontext.GetSubject(ctx),
		Domain:  ServiceName,
		Object:  object,
		Action:  action,
	})
	if err != nil {
		return errors.Wrap(err, "error on check permission")
	}

	return nil
}

func (mw *AuthorizationMiddleware) addPolicy(ctx context.Context, object string, action ...string) error {
	reqs := make([]authorization.AddPolicyRequest, 0, len(action))

	sub := sharedcontext.GetSubject(ctx)

	for i := range action {
		reqs = append(reqs, authorization.AddPolicyRequest{
			Subject: sub,
			Domain:  ServiceName,
			Object:  object,
			Action:  action[i],
		})
	}

	err := mw.authzSvc.AddPolicy(ctx, reqs...)
	if err != nil {
		return errors.Wrap(err, "error on add policy")
	}

	return nil
}

func (mw *AuthorizationMiddleware) removePolicy(ctx context.Context, object string, action ...string) error {
	reqs := make([]authorization.RemovePolicyRequest, 0, len(action))

	sub := sharedcontext.GetSubject(ctx)

	for i := range action {
		reqs = append(reqs, authorization.RemovePolicyRequest{
			Subject: sub,
			Domain:  ServiceName,
			Object:  object,
			Action:  action[i],
		})
	}

	err := mw.authzSvc.RemovePolicy(ctx, reqs...)
	if err != nil {
		return errors.Wrap(err, "error on remove policy")
	}

	return nil
}

func (mw *AuthorizationMiddleware) ListLocations(ctx context.Context) (*ListLocationsResponse, error) {
	//TODO: add policy
	rsp, err := mw.next.ListLocations(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error on list locations")
	}

	return rsp, nil
}

func (mw *AuthorizationMiddleware) ListPublishedItems(ctx context.Context, q string) (*ListItemsResponse, error) {
	err := mw.checkAccess(ctx, "", ActionListPublishedItems)
	if err != nil {
		return nil, errors.Wrap(err, "error on check permission")
	}

	rsp, err := mw.next.ListPublishedItems(ctx, q)
	if err != nil {
		return nil, errors.Wrap(err, "error on list published items")
	}

	return rsp, nil
}

func (mw *AuthorizationMiddleware) ListMyItems(ctx context.Context, q string) (*ListItemsResponse, error) {
	err := mw.checkAccess(ctx, "", ActionListMyItems)
	if err != nil {
		return nil, errors.Wrap(err, "error on check permission")
	}

	rsp, err := mw.next.ListMyItems(ctx, q)
	if err != nil {
		return nil, errors.Wrap(err, "error on list my items")
	}

	return rsp, nil
}

func (mw *AuthorizationMiddleware) GetItem(ctx context.Context, itemID string) (*Item, error) {
	err := mw.checkAccess(ctx, itemID, ActionGetItem)
	if err != nil {
		return nil, errors.Wrap(err, "error on check permission")
	}

	item, err := mw.next.GetItem(ctx, itemID)
	if err != nil {
		return nil, errors.Wrap(err, "error on get item")
	}

	return item, nil
}

func (mw *AuthorizationMiddleware) GetPublishedItem(ctx context.Context, itemID string) (*Item, error) {
	err := mw.checkAccess(ctx, itemID, ActionGetPublishedItem)
	if err != nil {
		return nil, errors.Wrap(err, "error on check permission")
	}

	item, err := mw.next.GetPublishedItem(ctx, itemID)
	if err != nil {
		return nil, errors.Wrap(err, "error on get published item")
	}

	return item, nil
}

func (mw *AuthorizationMiddleware) SendItemForPublish(ctx context.Context, itemID string) error {
	err := mw.checkAccess(ctx, itemID, ActionSendItemForPublish)
	if err != nil {
		return errors.Wrap(err, "error on check permission")
	}

	err = mw.next.SendItemForPublish(ctx, itemID)
	if err != nil {
		return errors.Wrap(err, "error on send item for publish")
	}

	return nil
}

func (mw *AuthorizationMiddleware) PublishItem(ctx context.Context, itemID string) error {
	err := mw.checkAccess(ctx, itemID, ActionPublishItem)
	if err != nil {
		return errors.Wrap(err, "error on check permission")
	}

	err = mw.next.PublishItem(ctx, itemID)
	if err != nil {
		return errors.Wrap(err, "error on publish item")
	}

	return nil
}

func (mw *AuthorizationMiddleware) ArchiveItem(ctx context.Context, itemID string) error {
	err := mw.checkAccess(ctx, itemID, ActionArchiveItem)
	if err != nil {
		return errors.Wrap(err, "error on check permission")
	}

	err = mw.next.ArchiveItem(ctx, itemID)
	if err != nil {
		return errors.Wrap(err, "error on archive item")
	}

	return nil
}

func (mw *AuthorizationMiddleware) DeleteItem(ctx context.Context, itemID string) error {
	err := mw.checkAccess(ctx, itemID, ActionDeleteItem)
	if err != nil {
		return errors.Wrap(err, "error on check permission")
	}

	err = mw.next.DeleteItem(ctx, itemID)
	if err != nil {
		return errors.Wrap(err, "error on delete item")
	}

	return nil
}

func (mw *AuthorizationMiddleware) CreateItem(ctx context.Context, req *CreateItemRequest) (*Item, error) {
	err := mw.checkAccess(ctx, "", ActionCreateItem)
	if err != nil {
		return nil, errors.Wrap(err, "error on check permission")
	}

	item, err := mw.next.CreateItem(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error on create item")
	}

	err = mw.addPolicy(ctx, item.ID,
		ActionGetItem,
		ActionSendItemForPublish,
		ActionArchiveItem,
		ActionDeleteItem,
		ActionUpdateItem,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error on add item policies")
	}

	return item, nil
}

func (mw *AuthorizationMiddleware) UpdateItem(ctx context.Context, itemID string, req *UpdateItemRequest) error {
	err := mw.checkAccess(ctx, itemID, ActionUpdateItem)
	if err != nil {
		return errors.Wrap(err, "error on check permission")
	}

	err = mw.next.UpdateItem(ctx, itemID, req)
	if err != nil {
		return errors.Wrap(err, "error on update item")
	}

	return nil
}
