package app

import "github.com/webitel/engine/model"

func (app *App) CreateRoutingOutboundCall(routing *model.RoutingOutboundCall) (*model.RoutingOutboundCall, *model.AppError) {
	return app.Store.RoutingOutboundCall().Create(routing)
}

func (app *App) GetRoutingOutboundCallPage(domainId int64, search *model.SearchRoutingOutboundCall) ([]*model.RoutingOutboundCall, bool, *model.AppError) {
	list, err := app.Store.RoutingOutboundCall().GetAllPage(domainId, search)
	if err != nil {
		return nil, false, err
	}
	search.RemoveLastElemIfNeed(&list)
	return list, search.EndOfList(), nil
}

func (app *App) GetRoutingOutboundCallById(domainId, id int64) (*model.RoutingOutboundCall, *model.AppError) {
	return app.Store.RoutingOutboundCall().Get(domainId, id)
}

func (app *App) UpdateRoutingOutboundCall(routing *model.RoutingOutboundCall) (*model.RoutingOutboundCall, *model.AppError) {
	oldRouting, err := app.GetRoutingOutboundCallById(routing.DomainId, routing.Id)
	if err != nil {
		return nil, err
	}

	oldRouting.Name = routing.Name
	oldRouting.Description = routing.Description
	oldRouting.Pattern = routing.Pattern
	oldRouting.Disabled = routing.Disabled
	oldRouting.UpdatedAt = routing.UpdatedAt

	if routing.GetSchemaId() != nil {
		oldRouting.Schema.Id = routing.Schema.Id
	}

	oldRouting, err = app.Store.RoutingOutboundCall().Update(oldRouting)
	if err != nil {
		return nil, err
	}

	return oldRouting, nil
}

func (app *App) ChangePositionOutboundCall(domainId, fromId, toId int64) *model.AppError {
	return app.Store.RoutingOutboundCall().ChangePosition(domainId, fromId, toId)
}

func (a *App) RemoveRoutingOutboundCall(domainId, id int64) (*model.RoutingOutboundCall, *model.AppError) {
	routing, err := a.Store.RoutingOutboundCall().Get(domainId, id)

	if err != nil {
		return nil, err
	}

	err = a.Store.RoutingOutboundCall().Delete(domainId, id)
	if err != nil {
		return nil, err
	}
	return routing, nil
}

func (a *App) PatchRoutingOutboundCall(domainId, id int64, patch *model.RoutingOutboundCallPatch) (*model.RoutingOutboundCall, *model.AppError) {
	old, err := a.GetRoutingOutboundCallById(domainId, id)
	if err != nil {
		return nil, err
	}
	old.Patch(patch)

	old.UpdatedAt = model.GetMillis()
	old.UpdatedBy.Id = patch.UpdatedById

	if err = old.IsValid(); err != nil {
		return nil, err
	}

	old, err = a.Store.RoutingOutboundCall().Update(old)
	if err != nil {
		return nil, err
	}

	return old, nil
}
