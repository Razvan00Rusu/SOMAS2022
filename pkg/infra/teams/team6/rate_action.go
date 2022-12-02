package team6

import (
	"infra/game/agent"
	"math"
)

func bound(f float64) uint8 {
	return uint8(math.Max(math.Min(f*255.0, 255.0), 0.0))
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
		lowHpPercentage := 1.0 - hpPercentage
		cower = bound(lowHpPercentage)

		attackIncrease := agentState.BonusAttack + agentState.Attack
		attackPercentage :=
			float64(attackIncrease) / float64(r.config.StartingAttackStrength)
		// Don't attack if low hp

		defendIncrease := agentState.BonusDefense + agentState.Defense
		defendPercentage :=
			float64(defendIncrease) / float64(r.config.StartingShieldStrength)

		staminaPercentage :=
			float64(agentState.Stamina) / float64(r.config.Stamina)

		monsterHealth := view.MonsterHealth()

		// Bias for attack on low hp monster
		attBias := float64(agentState.Hp) / float64(monsterHealth)

		defBias := 1.0 / attBias
		attack = bound(attackPercentage * hpPercentage * staminaPercentage * attBias)
		defend = bound(defendPercentage * staminaPercentage * hpPercentage * defBias)
	}

	return ActionDecision{
		attack: attack,
		cower:  cower,
		defend: defend,
	}
}
