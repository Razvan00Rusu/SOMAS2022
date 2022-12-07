package team6

import (
	"infra/game/agent"
	"infra/game/decision"
)

func (t6 *Team6Agent) FightAction(baseAgent agent.BaseAgent, proposedAction decision.FightAction) decision.FightAction {
	return t6.FightActionNoProposal(baseAgent)
}
