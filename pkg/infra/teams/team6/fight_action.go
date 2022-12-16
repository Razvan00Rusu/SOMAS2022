package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"
)

const LEADERSHIP_THRESHOLD = 100

func (r *Perry) FightAction(
	base agent.BaseAgent,
	prop decision.FightAction,
	_ message.Proposal[decision.FightAction],
) decision.FightAction {
	view := base.View()
	leader := view.CurrentLeader()
	if r.opinions[leader].leadership > LEADERSHIP_THRESHOLD {
		return prop
	} else {
		return r.FightActionNoProposal(base)
	}

}
