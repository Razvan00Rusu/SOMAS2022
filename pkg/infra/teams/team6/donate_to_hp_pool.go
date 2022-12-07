package team6

import (
	"infra/game/agent"
)

func (r *Team6Agent) DonateToHpPool(baseAgent agent.BaseAgent) uint {
	hp := baseAgent.AgentState().Hp
	if hp > uint(0.5*float32(r.config.StartingHealthPoints)) {
		return uint((hp - r.config.StartingHealthPoints) / 3)
	} else {
		return 0
	}
}
