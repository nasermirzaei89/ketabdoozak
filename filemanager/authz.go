package filemanager

import (
	"context"
	"github.com/nasermirzaei89/ketabdoozak/sharedcontext"
	"github.com/nasermirzaei89/services/authorization"
	"github.com/pkg/errors"
	"io"
)

const (
	ActionUploadFile       = "uploadFile"
	ActionGetPublicFileURL = "getPublicFileUrl"
	ActionDeleteFile       = "deleteFile"
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

func (mw *AuthorizationMiddleware) addPolicy(ctx context.Context, object, action string) error {
	err := mw.authzSvc.AddPolicy(ctx, authorization.AddPolicyRequest{
		Subject: sharedcontext.GetSubject(ctx),
		Domain:  ServiceName,
		Object:  object,
		Action:  action,
	})
	if err != nil {
		return errors.Wrap(err, "error on add policy")
	}

	return nil
}

func (mw *AuthorizationMiddleware) removePolicy(ctx context.Context, object, action string) error {
	err := mw.authzSvc.RemovePolicy(ctx, authorization.RemovePolicyRequest{
		Subject: sharedcontext.GetSubject(ctx),
		Domain:  ServiceName,
		Object:  object,
		Action:  action,
	})
	if err != nil {
		return errors.Wrap(err, "error on remove policy")
	}

	return nil
}

func (mw *AuthorizationMiddleware) UploadFile(ctx context.Context, filename string, reader io.Reader, fileSize int64, contentType string) (*File, error) {
	err := mw.checkAccess(ctx, "", ActionUploadFile)
	if err != nil {
		return nil, errors.Wrap(err, "error on check permission")
	}

	file, err := mw.next.UploadFile(ctx, filename, reader, fileSize, contentType)
	if err != nil {
		return nil, errors.Wrap(err, "error on upload file")
	}

	err = mw.addPolicy(ctx, file.Filename, ActionDeleteFile)
	if err != nil {
		return nil, errors.Wrap(err, "error on add delete file policy")
	}

	return file, nil
}

func (mw *AuthorizationMiddleware) GetPublicFileURL(ctx context.Context, filename string) (string, error) {
	err := mw.checkAccess(ctx, filename, ActionGetPublicFileURL)
	if err != nil {
		return "", errors.Wrap(err, "error on check permission")
	}

	fileURL, err := mw.next.GetPublicFileURL(ctx, filename)
	if err != nil {
		return "", errors.Wrap(err, "error on get public file url")
	}

	return fileURL, nil
}

func (mw *AuthorizationMiddleware) DeleteFile(ctx context.Context, filename string) error {
	err := mw.checkAccess(ctx, filename, ActionGetPublicFileURL)
	if err != nil {
		return errors.Wrap(err, "error on check permission")
	}

	err = mw.next.DeleteFile(ctx, filename)
	if err != nil {
		return errors.Wrap(err, "error on delete file")
	}

	err = mw.removePolicy(ctx, filename, ActionDeleteFile)
	if err != nil {
		return errors.Wrap(err, "error on remove delete file policy")
	}

	return nil
}
