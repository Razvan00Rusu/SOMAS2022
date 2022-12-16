package team6

import (
	"infra/config"
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message/proposal"
	"infra/game/stage/initialise"
)

type Perry struct {
	bravery    int
	opinions   map[string](*PerryOpinion)
	nextAction ActionDecision
	nextFights commons.ImmutableList[proposal.Rule[decision.FightAction]]
	nextLoot   commons.ImmutableList[proposal.Rule[decision.LootAction]]
	config     config.GameConfig
}

func NewPerry() agent.Strategy {
	var fights []proposal.Rule[decision.FightAction]
	var loot []proposal.Rule[decision.LootAction]
	return &Perry{
		bravery:    0,
		opinions:   make(map[string](*PerryOpinion)),
		nextAction: NewActionDecision(),
		nextFights: *commons.NewImmutableList[proposal.Rule[decision.FightAction]](fights),
		nextLoot:   *commons.NewImmutableList[proposal.Rule[decision.LootAction]](loot),
		config:     initialise.InitGameConfig(),
	}
}
