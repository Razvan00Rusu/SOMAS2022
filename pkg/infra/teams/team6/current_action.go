package team6

import (
	"infra/game/decision"
)

func (r *Perry) CurrentAction() decision.FightAction {
	action := r.nextAction
	return action.highestRatedAction()
}
