package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"

	"github.com/benbjohnson/immutable"
)

func (r *Perry) HandleFightProposalRequest(
	prop message.Proposal[decision.FightAction],
	base agent.BaseAgent,
	log *immutable.Map[string, decision.FightAction],
) bool {
	// What should this do?
	return true
}
