package team6

import (
	"infra/game/decision"
)

func (r *Team6Agent) CurrentAction() decision.FightAction {
	action := r.nextAction
	return action.highestRatedAction()
}
