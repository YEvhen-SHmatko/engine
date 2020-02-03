package app

import (
	"github.com/webitel/engine/auth_manager"
	"github.com/webitel/engine/model"
)

func (a *App) GetCalendarsPage(domainId int64, search *model.SearchCalendar) ([]*model.Calendar, bool, *model.AppError) {
	list, err := a.Store.Calendar().GetAllPage(domainId, search)
	if err != nil {
		return nil, false, err
	}
	search.RemoveLastElemIfNeed(&list)
	return list, search.EndOfList(), nil
}

func (a *App) GetCalendarPageByGroups(domainId int64, groups []int, search *model.SearchCalendar) ([]*model.Calendar, bool, *model.AppError) {
	list, err := a.Store.Calendar().GetAllPageByGroups(domainId, groups, search)
	if err != nil {
		return nil, false, err
	}
	search.RemoveLastElemIfNeed(&list)
	return list, search.EndOfList(), nil
}

func (a *App) CreateCalendar(calendar *model.Calendar) (*model.Calendar, *model.AppError) {
	return a.Store.Calendar().Create(calendar)
}

func (a *App) UpdateCalendar(calendar *model.Calendar) (*model.Calendar, *model.AppError) {
	oldCalendar, err := a.GetCalendarById(calendar.DomainId, calendar.Id)
	if err != nil {
		return nil, err
	}

	oldCalendar.Timezone.Id = calendar.Timezone.Id
	oldCalendar.Description = calendar.Description
	oldCalendar.Name = calendar.Name
	oldCalendar.StartAt = calendar.StartAt
	oldCalendar.EndAt = calendar.EndAt
	oldCalendar.UpdatedAt = calendar.UpdatedAt
	oldCalendar.UpdatedBy.Id = calendar.UpdatedBy.Id
	oldCalendar.Accepts = calendar.Accepts
	oldCalendar.Excepts = calendar.Excepts

	oldCalendar, err = a.Store.Calendar().Update(oldCalendar)
	if err != nil {
		return nil, err
	}

	return oldCalendar, nil
}

func (a *App) CalendarCheckAccess(domainId, id int64, groups []int, access auth_manager.PermissionAccess) (bool, *model.AppError) {
	return a.Store.Calendar().CheckAccess(domainId, id, groups, access)
}

func (a *App) GetCalendarById(domainId, id int64) (*model.Calendar, *model.AppError) {
	return a.Store.Calendar().Get(domainId, id)
}

func (a *App) RemoveCalendar(domainId, id int64) (*model.Calendar, *model.AppError) {
	calendar, err := a.Store.Calendar().Get(domainId, id)

	if err != nil {
		return nil, err
	}

	err = a.Store.Calendar().Delete(domainId, id)
	if err != nil {
		return nil, err
	}
	return calendar, nil
}

func (a *App) GetCalendarTimezoneAllPage(search *model.SearchTimezone) ([]*model.Timezone, bool, *model.AppError) {
	list, err := a.Store.Calendar().GetTimezoneAllPage(search)
	if err != nil {
		return nil, false, err
	}
	search.RemoveLastElemIfNeed(&list)
	return list, search.EndOfList(), nil
}
