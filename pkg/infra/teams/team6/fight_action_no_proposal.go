package team6

import (
	"infra/game/agent"
	"infra/game/decision"
)

// TODO
func (r *Team6Agent) FightActionNoProposal(baseAgent agent.BaseAgent) decision.FightAction {
	return decision.Cower
}
