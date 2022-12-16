package team6

import (
	"infra/game/decision"
	"math/rand"
)

func (action *ActionDecision) highestRatedAction() decision.FightAction {

	if action.attack > action.cower && (action.attack > action.defend) {
		return decision.Attack
	}

	if action.cower > action.attack && (action.cower > action.defend) {
		return decision.Cower
	}

	if action.defend > action.cower && (action.defend > action.attack) {
		return decision.Defend
	}

	if action.defend > action.cower {
		if rand.Intn(2) == 0 {
			return decision.Attack
		}
		return decision.Defend
	}

	if action.defend > action.attack {
		if rand.Intn(2) == 0 {
			return decision.Cower
		}
		return decision.Defend
	}

	if action.attack > action.defend {
		if rand.Intn(2) == 0 {
			return decision.Attack
		}
		return decision.Cower
	}

	if action.attack == 0 && action.defend == 0 && action.cower == 0 {
		action.attack++
		action.defend++
		action.cower++
	}
	max := int(action.attack) + int(action.cower) + int(action.defend)
	choice := rand.Intn(max)
	if choice <= int(action.attack) {
		return decision.Attack
	}
	if choice <= int(action.attack)+int(action.cower) {
		return decision.Cower
	}
	return decision.Defend
}
