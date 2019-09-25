package sqlstore

import (
	"fmt"
	"github.com/webitel/engine/model"
	"github.com/webitel/engine/store"
	"net/http"
)

type SqlResourceTeamStore struct {
	SqlStore
}

func NewSqlResourceTeamStore(sqlStore SqlStore) store.ResourceTeamStore {
	us := &SqlResourceTeamStore{sqlStore}
	return us
}

func (s SqlResourceTeamStore) Create(in *model.ResourceInTeam) (*model.ResourceInTeam, *model.AppError) {
	var out *model.ResourceInTeam
	err := s.GetMaster().SelectOne(&out, `with i as (
    insert into cc_agent_in_team (team_id, agent_id, skill_id, lvl, min_capacity, max_capacity)
        values (:TeamId, :AgentId, :SkillId, :Lvl, :MinCapacity, :MaxCapacity)
        returning *
)
select i.id,
       i.team_id,
       cc_get_lookup(a.id, u.name) as agent,
       cc_get_lookup(s.id, s.name) as skill,
       i.lvl,
       i.min_capacity,
       i.max_capacity
from i
         left join cc_agent a on a.id = i.agent_id
         left join directory.wbt_user u on u.id = a.user_id
         left join cc_skill s on s.id = i.skill_id`, map[string]interface{}{
		"TeamId":      in.TeamId,
		"AgentId":     in.AgentId(),
		"SkillId":     in.SkillId(),
		"Lvl":         in.Lvl,
		"MinCapacity": in.MinCapacity,
		"MaxCapacity": in.MaxCapacity,
	})

	if err != nil {
		return nil, model.NewAppError("SqlResourceTeamStore.Save", "store.sql_resource_team.save.app_error", nil,
			fmt.Sprintf("TeamId=%v, %v", in.TeamId, err.Error()), extractCodeFromErr(err))
	}

	return out, nil
}

func (s SqlResourceTeamStore) Get(domainId, teamId int64, id int64) (*model.ResourceInTeam, *model.AppError) {
	var resource *model.ResourceInTeam
	if err := s.GetReplica().SelectOne(&resource, `select i.id,
       i.team_id,
       cc_get_lookup(a.id, u.name) as agent,
       cc_get_lookup(s.id, s.name) as skill,
       i.lvl,
       i.min_capacity,
       i.max_capacity
from cc_agent_in_team i
         left join cc_agent a on a.id = i.agent_id
         left join directory.wbt_user u on u.id = a.user_id
         left join cc_skill s on s.id = i.skill_id
where i.id = :Id and i.team_id = :TeamId and exists (select 1 from cc_team t where t.id = :TeamId and t.domain_id = :DomainId)`,
		map[string]interface{}{
			"Id":       id,
			"TeamId":   teamId,
			"DomainId": domainId,
		}); err != nil {
		return nil, model.NewAppError("SqlResourceTeamStore.Get", "store.sql_resource_team.get.app_error", nil,
			fmt.Sprintf("Id=%v, %s", id, err.Error()), extractCodeFromErr(err))
	} else {
		return resource, nil
	}
}

func (s SqlResourceTeamStore) GetAllPage(domainId, teamId int64, offset, limit int) ([]*model.ResourceInTeam, *model.AppError) {
	var resources []*model.ResourceInTeam

	if _, err := s.GetReplica().Select(&resources,
		`select i.id,
       i.team_id,
       cc_get_lookup(a.id, u.name) as agent,
       cc_get_lookup(s.id, s.name) as skill,
       i.lvl,
       i.min_capacity,
       i.max_capacity
from cc_agent_in_team i
         left join cc_agent a on a.id = i.agent_id
         left join directory.wbt_user u on u.id = a.user_id
         left join cc_skill s on s.id = i.skill_id
where i.team_id = :TeamId and exists (select 1 from cc_team t where t.id = :TeamId and t.domain_id = :DomainId)
order by i.id
limit :Limit
offset :Offset`, map[string]interface{}{
			"DomainId": domainId,
			"TeamId":   teamId,
			"Limit":    limit,
			"Offset":   offset,
		}); err != nil {
		return nil, model.NewAppError("SqlResourceTeamStore.GetAllPage", "store.sql_resource_team.get_all.app_error", nil, err.Error(), http.StatusInternalServerError)
	} else {
		return resources, nil
	}
}

func (s SqlResourceTeamStore) Update(resource *model.ResourceInTeam) (*model.ResourceInTeam, *model.AppError) {
	err := s.GetMaster().SelectOne(&resource, `with i as (
    update cc_agent_in_team at
    set agent_id = :AgentId,
        skill_id = :SkillId,
        lvl = :Lvl,
        min_capacity = :MinCapacity,
        max_capacity = :MaxCapacity
    where at.id = :Id and at.team_id = :TeamId
    returning *
)
select i.id,
       i.team_id,
       cc_get_lookup(a.id, u.name) as agent,
       cc_get_lookup(s.id, s.name) as skill,
       i.lvl,
       i.min_capacity,
       i.max_capacity
from i
         left join cc_agent a on a.id = i.agent_id
         left join directory.wbt_user u on u.id = a.user_id
         left join cc_skill s on s.id = i.skill_id`, map[string]interface{}{
		"AgentId":     resource.AgentId(),
		"SkillId":     resource.SkillId(),
		"Lvl":         resource.Lvl,
		"MinCapacity": resource.MinCapacity,
		"MaxCapacity": resource.MaxCapacity,
		"Id":          resource.Id,
		"TeamId":      resource.TeamId,
	})
	if err != nil {
		return nil, model.NewAppError("SqlResourceTeamStore.Update", "store.sql_resource_team.update.app_error", nil,
			fmt.Sprintf("Id=%v, %s", resource.Id, err.Error()), extractCodeFromErr(err))
	}
	return resource, nil
}

func (s SqlResourceTeamStore) Delete(domainId, teamId int64, id int64) *model.AppError {
	if _, err := s.GetMaster().Exec(`delete from cc_agent_in_team at
where at.team_id = :TeamId and at.id = :Id and exists(select 1 from cc_team t where t.id = :TeamId and t.domain_id = :DomainId)`,
		map[string]interface{}{"Id": id, "DomainId": domainId, "TeamId": teamId}); err != nil {
		return model.NewAppError("SqlResourceTeamStore.Delete", "store.sql_resource_team.delete.app_error", nil,
			fmt.Sprintf("Id=%v, %s", id, err.Error()), http.StatusInternalServerError)
	}
	return nil
}
