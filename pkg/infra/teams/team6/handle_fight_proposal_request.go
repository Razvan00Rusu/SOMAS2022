package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message"
	"math/rand"

	"github.com/benbjohnson/immutable"
)

func (r *Team6Agent) HandleFightProposalRequest(
	_ *message.FightProposalMessage,
	_ agent.BaseAgent,
	_ *immutable.Map[commons.ID, decision.FightAction],
) bool {
	switch rand.Intn(2) {
	case 0:
		return true
	default:
		return false
	}
}
