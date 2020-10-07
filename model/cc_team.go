package model

type AgentTeam struct {
	DomainRecord
	Name              string `json:"name" db:"name"`
	Description       string `json:"description" db:"description"`
	Strategy          string `json:"strategy" db:"strategy"`
	MaxNoAnswer       int16  `json:"max_no_answer" db:"max_no_answer"`
	WrapUpTime        int16  `json:"wrap_up_time" db:"wrap_up_time"`
	NoAnswerDelayTime int16  `json:"no_answer_delay_time" db:"no_answer_delay_time"`
	CallTimeout       int16  `json:"call_timeout" db:"call_timeout"`
	PostProcessing    bool   `json:"post_processing" db:"post_processing"`
}

func (team AgentTeam) DefaultOrder() string {
	return "id"
}

func (team AgentTeam) AllowFields() []string {
	return team.DefaultFields()
}

func (team AgentTeam) DefaultFields() []string {
	return []string{"id", "name", "description", "strategy", "max_no_answer", "wrap_up_time",
		"no_answer_delay_time", "call_timeout", "updated_at", "post_processing"}
}

func (team AgentTeam) EntityName() string {
	return "cc_team"
}

type SearchAgentTeam struct {
	ListRequest
	Ids []string
}

func (team *AgentTeam) IsValid() *AppError {
	return nil
}