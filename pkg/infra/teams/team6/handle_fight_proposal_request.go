package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message"

	"github.com/benbjohnson/immutable"
)

func (r *Team6Agent) HandleFightProposalRequest(
	msg *message.FightProposalMessage,
	base agent.BaseAgent,
	my *immutable.Map[commons.ID, decision.FightAction],
) bool {
	if msg.Proposal() == *my {
		println("My Prop!")
	}
	return true
}
