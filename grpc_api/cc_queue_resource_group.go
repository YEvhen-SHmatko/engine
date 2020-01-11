package grpc_api

import (
	"context"
	"github.com/webitel/engine/app"
	"github.com/webitel/engine/grpc_api/engine"
	"github.com/webitel/engine/model"
)

type queueResource struct {
	app *app.App
}

func NewQueueResourceApi(app *app.App) *queueResource {
	return &queueResource{app: app}
}

func (api *queueResource) CreateQueueResourceGroup(ctx context.Context, in *engine.CreateQueueResourceGroupRequest) (*engine.QueueResourceGroup, error) {
	session, err := api.app.GetSessionFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	permission := session.GetPermission(model.PERMISSION_SCOPE_CC_QUEUE)
	if !permission.CanRead() {
		return nil, api.app.MakePermissionError(session, permission, model.PERMISSION_ACCESS_READ)
	}

	if !permission.CanUpdate() {
		return nil, api.app.MakePermissionError(session, permission, model.PERMISSION_ACCESS_UPDATE)
	}

	if permission.Rbac {
		var perm bool
		if perm, err = api.app.QueueCheckAccess(session.Domain(in.GetDomainId()), in.GetQueueId(), session.RoleIds, model.PERMISSION_ACCESS_UPDATE); err != nil {
			return nil, err
		} else if !perm {
			return nil, api.app.MakeResourcePermissionError(session, in.GetQueueId(), permission, model.PERMISSION_ACCESS_UPDATE)
		}
	}

	var queueResourceGroup = &model.QueueResourceGroup{
		QueueId: in.GetQueueId(),
		ResourceGroup: model.Lookup{
			Id: int(in.GetResourceGroup().GetId()),
		},
	}

	if err = queueResourceGroup.IsValid(); err != nil {
		return nil, err
	}

	queueResourceGroup, err = api.app.CreateQueueResourceGroup(queueResourceGroup)
	if err != nil {
		return nil, err
	}

	return toEngineQueueResourceGroup(queueResourceGroup), nil
}

func (api *queueResource) SearchQueueResourceGroup(ctx context.Context, in *engine.SearchQueueResourceGroupRequest) (*engine.ListQueueResourceGroup, error) {
	session, err := api.app.GetSessionFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	permission := session.GetPermission(model.PERMISSION_SCOPE_CC_QUEUE)
	if !permission.CanRead() {
		return nil, api.app.MakePermissionError(session, permission, model.PERMISSION_ACCESS_READ)
	}

	if permission.Rbac {
		var perm bool
		if perm, err = api.app.QueueCheckAccess(session.Domain(in.GetDomainId()), in.GetQueueId(), session.RoleIds, model.PERMISSION_ACCESS_READ); err != nil {
			return nil, err
		} else if !perm {
			return nil, api.app.MakeResourcePermissionError(session, in.GetQueueId(), permission, model.PERMISSION_ACCESS_READ)
		}
	}

	var list []*model.QueueResourceGroup
	list, err = api.app.GetQueueResourceGroupPage(session.Domain(in.DomainId), in.GetQueueId(), int(in.Page), int(in.Size))
	if err != nil {
		return nil, err
	}

	items := make([]*engine.QueueResourceGroup, 0, len(list))
	for _, v := range list {
		items = append(items, toEngineQueueResourceGroup(v))
	}
	return &engine.ListQueueResourceGroup{
		Items: items,
	}, nil
}

func (api *queueResource) ReadQueueResourceGroup(ctx context.Context, in *engine.ReadQueueResourceGroupRequest) (*engine.QueueResourceGroup, error) {
	session, err := api.app.GetSessionFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	permission := session.GetPermission(model.PERMISSION_SCOPE_CC_QUEUE)
	if !permission.CanRead() {
		return nil, api.app.MakePermissionError(session, permission, model.PERMISSION_ACCESS_READ)
	}

	if permission.Rbac {
		var perm bool
		if perm, err = api.app.QueueCheckAccess(session.Domain(in.GetDomainId()), in.GetQueueId(), session.RoleIds, model.PERMISSION_ACCESS_READ); err != nil {
			return nil, err
		} else if !perm {
			return nil, api.app.MakeResourcePermissionError(session, in.GetQueueId(), permission, model.PERMISSION_ACCESS_READ)
		}
	}

	var out *model.QueueResourceGroup
	out, err = api.app.GetQueueResourceGroup(session.Domain(in.GetDomainId()), in.GetQueueId(), in.GetId())
	if err != nil {
		return nil, err
	}
	return toEngineQueueResourceGroup(out), nil
}

func (api *queueResource) UpdateQueueResourceGroup(ctx context.Context, in *engine.UpdateQueueResourceGroupRequest) (*engine.QueueResourceGroup, error) {
	session, err := api.app.GetSessionFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	permission := session.GetPermission(model.PERMISSION_SCOPE_CC_QUEUE)
	if !permission.CanRead() {
		return nil, api.app.MakePermissionError(session, permission, model.PERMISSION_ACCESS_READ)
	}

	if !permission.CanUpdate() {
		return nil, api.app.MakePermissionError(session, permission, model.PERMISSION_ACCESS_UPDATE)
	}

	if permission.Rbac {
		var perm bool
		if perm, err = api.app.QueueCheckAccess(session.Domain(in.GetDomainId()), in.GetQueueId(), session.RoleIds, model.PERMISSION_ACCESS_UPDATE); err != nil {
			return nil, err
		} else if !perm {
			return nil, api.app.MakeResourcePermissionError(session, in.GetQueueId(), permission, model.PERMISSION_ACCESS_UPDATE)
		}
	}

	qr := &model.QueueResourceGroup{
		Id:      in.GetId(),
		QueueId: in.GetQueueId(),
		ResourceGroup: model.Lookup{
			Id: int(in.GetResourceGroup().GetId()),
		},
	}

	if err = qr.IsValid(); err != nil {
		return nil, err
	}

	qr, err = api.app.UpdateQueueResourceGroup(session.Domain(in.GetDomainId()), qr)
	if err != nil {
		return nil, err
	}

	return toEngineQueueResourceGroup(qr), nil
}

func (api *queueResource) DeleteQueueResourceGroup(ctx context.Context, in *engine.DeleteQueueResourceGroupRequest) (*engine.QueueResourceGroup, error) {
	session, err := api.app.GetSessionFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	permission := session.GetPermission(model.PERMISSION_SCOPE_CC_QUEUE)
	if !permission.CanRead() {
		return nil, api.app.MakePermissionError(session, permission, model.PERMISSION_ACCESS_READ)
	}

	if !permission.CanUpdate() {
		return nil, api.app.MakePermissionError(session, permission, model.PERMISSION_ACCESS_UPDATE)
	}

	if permission.Rbac {
		var perm bool
		if perm, err = api.app.QueueCheckAccess(session.Domain(in.GetDomainId()), in.GetQueueId(), session.RoleIds, model.PERMISSION_ACCESS_UPDATE); err != nil {
			return nil, err
		} else if !perm {
			return nil, api.app.MakeResourcePermissionError(session, in.GetQueueId(), permission, model.PERMISSION_ACCESS_UPDATE)
		}
	}

	var qr *model.QueueResourceGroup
	qr, err = api.app.RemoveQueueResourceGroup(session.Domain(in.GetDomainId()), in.GetQueueId(), in.GetId())
	if err != nil {
		return nil, err
	}
	return toEngineQueueResourceGroup(qr), nil
}

func toEngineQueueResourceGroup(src *model.QueueResourceGroup) *engine.QueueResourceGroup {
	return &engine.QueueResourceGroup{
		Id: src.Id,
		ResourceGroup: &engine.Lookup{
			Id:   int64(src.ResourceGroup.Id),
			Name: src.ResourceGroup.Name,
		},
	}
}