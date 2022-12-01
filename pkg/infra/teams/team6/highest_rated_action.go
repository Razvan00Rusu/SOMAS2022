package team6

import (
	"infra/game/decision"
	"math/rand"
)

func (action *ActionDecision) highestRatedAction() decision.FightAction {
	max := int(action.attack) + int(action.cower) + int(action.defend)
	choice := rand.Intn(max)
	if choice < int(action.attack) {
		return decision.Attack
	}
	if choice < int(action.attack)+int(action.cower) {
		return decision.Cower
	}
	return decision.Defend
}
