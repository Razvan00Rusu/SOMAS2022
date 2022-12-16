package team6

import (
	"infra/game/agent"
	"infra/game/decision"
)

func (r *Perry) HandleUpdateWeapon(base agent.BaseAgent) decision.ItemIdx {
	// weapons := b.AgentState().Weapons
	// return decision.ItemIdx(rand.Intn(weapons.Len() + 1))

	// 0th weapon has greatest attack points
	// TODO
	return decision.ItemIdx(0)
}
