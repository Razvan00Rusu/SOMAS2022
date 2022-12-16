package team6

import (
	"infra/game/agent"
	"math"
)

func bound(f float64) uint16 {
	return uint16(math.Max(math.Min(f*65535.0, 65535.0), 0.0))
}

func (r *Team6Agent) rateActions(
	id string,
	base agent.BaseAgent,
) ActionDecision {

	attack := uint16(32767)
	cower := uint16(32767)
	defend := uint16(32767)

	view := base.View()
	agentStates := view.AgentState()
	agentState, ok := agentStates.Get(id)

	if ok {
		hpPercentage :=
			math.Max(float64(agentState.Hp), 1.0) / float64(r.config.StartingHealthPoints)
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

		monsterHealth := math.Abs(float64(view.MonsterHealth()))
		monsterAttack := math.Abs(float64(view.MonsterAttack()))

		// Bias for attack on low hp monster
		attBias := 1.0 / normalise(monsterAttack/(monsterHealth*float64(attackIncrease)))
		defBias := 1.0 / normalise(monsterHealth/(monsterHealth*float64(defendIncrease)))
		//println(attBias, defBias)

		attack = bound(attackPercentage * hpPercentage * staminaPercentage * attBias)
		if int(agentState.Stamina)-int(agentState.BonusAttack) < 0 {
			attack = 0
		}
		defend = bound(defendPercentage * staminaPercentage * hpPercentage * defBias)
		if int(agentState.Stamina)-int(agentState.BonusDefense) < 0 {
			defend = 0
		}
	}

	return ActionDecision{
		attack: attack,
		cower:  cower,
		defend: defend,
	}
}
