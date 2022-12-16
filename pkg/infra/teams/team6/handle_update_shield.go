package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/state"
)

func (r *Team6Agent) HandleUpdateShield(view *state.View, b agent.BaseAgent) decision.ItemIdx {
	// shields := b.AgentState().Shields
	// return decision.ItemIdx(rand.Intn(shields.Len() + 1))
	return decision.ItemIdx(0)
}
