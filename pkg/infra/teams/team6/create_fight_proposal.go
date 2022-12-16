package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message/proposal"
)

func (r *Perry) CreateFightProposal(base agent.BaseAgent) commons.ImmutableList[proposal.Rule[decision.FightAction]] {
	// TODO
	var rtn []proposal.Rule[decision.FightAction]

	return *commons.NewImmutableList[proposal.Rule[decision.FightAction]](rtn)
}
