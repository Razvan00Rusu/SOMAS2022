package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message"

	"github.com/benbjohnson/immutable"
)

// TODO
func (r *Team6Agent) HandleFightProposalRequest(
	proposal message.Proposal[decision.FightAction],
	base agent.BaseAgent,
	my *immutable.Map[commons.ID, decision.FightAction],
) bool {

	return true
}
