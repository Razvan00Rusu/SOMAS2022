package team6

import (
	"infra/config"
	"infra/game/agent"
	"infra/game/stage/initialise"
)

type Team6Agent struct {
	bravery    int
	opinions   map[string](*Team6AgentOpinion)
	nextAction ActionDecision
	config     config.GameConfig
}

func NewTeam6Agent() agent.Strategy {
	return &Team6Agent{
		bravery:    0,
		opinions:   make(map[string](*Team6AgentOpinion)),
		nextAction: NewActionDecision(),
		config:     initialise.InitGameConfig(),
	}
}
