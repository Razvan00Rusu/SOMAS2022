package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"math/rand"
)

func (r *Team6Agent) HandleConfidencePoll(_ agent.BaseAgent) decision.Intent {
	switch rand.Intn(3) {
	case 0:
		return decision.Abstain
	case 1:
		return decision.Negative
	default:
		return decision.Positive
	}
}
