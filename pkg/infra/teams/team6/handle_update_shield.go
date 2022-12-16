package team6

import (
	"infra/game/agent"
	"infra/game/decision"
)

func (r *Perry) HandleUpdateShield(base agent.BaseAgent) decision.ItemIdx {
	// shields := b.AgentState().Shields
	// return decision.ItemIdx(rand.Intn(shields.Len() + 1))
	// TODO
	return decision.ItemIdx(0)
}
