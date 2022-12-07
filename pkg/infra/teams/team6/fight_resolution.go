package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message/proposal"

	"github.com/benbjohnson/immutable"
)

func (r *Team6Agent) FightResolution(
	baseAgent agent.BaseAgent,
	prop commons.ImmutableList[proposal.Rule[decision.FightAction]],
) immutable.Map[commons.ID, decision.FightAction] {

	actions := make(map[commons.ID]decision.FightAction)
	ratings := make(map[commons.ID]ActionDecision)
	view := baseAgent.View()
	agentState := view.AgentState()
	itrRatings := agentState.Iterator()
	for !itrRatings.Done() {
		id, _, ok := itrRatings.Next()
		if !ok {
			break
		}
		ratings[id] = r.rateActions(id, baseAgent)
	}
	r.nextAction = ratings[baseAgent.ID()]

	// TODO: Implement a form of adjustment based on HP pool, monster health, etc.

	itrActions := agentState.Iterator()
	for !itrActions.Done() {
		id, _, ok := itrActions.Next()
		if !ok {
			break
		}
		action := ratings[id]
		actions[id] = action.highestRatedAction()
	}

	return commons.MapToImmutable(actions)
}
