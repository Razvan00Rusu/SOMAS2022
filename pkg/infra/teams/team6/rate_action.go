package team6

import (
	"infra/game/agent"
	"math"
)

func bound(f float64) float64 {
	return math.Max(math.Min(f, 255.0), 0.0)
}

func (r *Team6Agent) rateActions(
	id string,
	base agent.BaseAgent,
) ActionDecision {

	attack := uint8(127)
	cower := uint8(127)
	defend := uint8(127)

	view := base.View()
	agentStates := view.AgentState()
	agentState, ok := agentStates.Get(id)

	if ok {
		hpPercentage :=
			float64(agentState.Hp) / float64(r.config.StartingHealthPoints)
		cower = uint8(bound((1.0 - hpPercentage) * 255.0))

		attackIncrease := agentState.Attack - r.config.StartingAttackStrength
		attackPercentage :=
			float64(attackIncrease) / float64(r.config.StartingAttackStrength)
		attack = uint8(bound(attackPercentage*255.0 + 128.0))

		defendIncrease := agentState.Defense - r.config.StartingShieldStrength
		defendPercentage :=
			float64(defendIncrease) / float64(r.config.StartingShieldStrength)
		defend = uint8(bound(defendPercentage*255.0 + 128.0))

		//staminaIncrease := uint(agentState.Stamina) - r.config.Stamina
	}

	return ActionDecision{
		attack: attack,
		cower:  cower,
		defend: defend,
	}
}
