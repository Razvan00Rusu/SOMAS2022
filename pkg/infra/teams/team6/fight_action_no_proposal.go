package team6

import (
	"infra/game/agent"
	"infra/game/decision"
)

func (r *Perry) FightActionNoProposal(base agent.BaseAgent) decision.FightAction {
	action := r.nextAction
	return action.highestRatedAction()
}
