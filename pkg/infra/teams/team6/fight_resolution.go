package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message/proposal"
	"math"

	"github.com/benbjohnson/immutable"
)

func normalise(f float64) float64 {
	return (math.Log10(f+0.1) + 1.0) / 2.1
}

func (r *Perry) FightResolution(
	baseAgent agent.BaseAgent,
	prop commons.ImmutableList[proposal.Rule[decision.FightAction]],
	proposedActions immutable.Map[string, decision.FightAction],
) immutable.Map[commons.ID, decision.FightAction] {
	actions := make(map[commons.ID]decision.FightAction)
	ratings := make(map[commons.ID]ActionDecision)
	view := baseAgent.View()
	agentState := view.AgentState()
	itrRatings := agentState.Iterator()

	numAgents := 0

	totalAttack := 0
	totalCower := 0
	totalDefense := 0

	for !itrRatings.Done() {
		id, _, ok := itrRatings.Next()
		if !ok {
			break
		}
		numAgents++
		ratings[id] = r.rateActions(id, baseAgent)
		if ratings[id].attack >= ratings[id].cower && ratings[id].attack >= ratings[id].defend {
			totalAttack++
		}
		if ratings[id].cower >= ratings[id].attack && ratings[id].cower >= ratings[id].defend {
			totalCower++
		}
		if ratings[id].defend >= ratings[id].attack && ratings[id].defend >= ratings[id].cower {
			totalDefense++
		}
	}
	r.nextAction = ratings[baseAgent.ID()]

	adjustAttack := float64(numAgents) / (3.0 * float64(totalAttack))
	adjustCower := float64(numAgents) / (3.0 * float64(totalCower))
	if totalCower == 0 {
		adjustCower = 0
	}
	adjustDefense := float64(numAgents) / (3.0 * float64(totalDefense))
	//println("Attack:", adjustAttack, "Cower:", adjustCower, "Defense:", adjustDefense)

	// TODO: Implement a form of adjustment based on HP pool, monster health, etc.

	itrActions := agentState.Iterator()
	for !itrActions.Done() {
		id, _, ok := itrActions.Next()
		if !ok {
			break
		}
		action := ratings[id]
		action.attack = bound((float64(action.attack) / 65536.0) * adjustAttack)
		action.cower = bound((float64(action.cower) / 65536.0) * adjustCower)
		action.defend = bound((float64(action.defend) / 65536.0) * adjustDefense)
		//println(action.attack, action.cower, action.defend)
		actions[id] = action.highestRatedAction()
	}
	return commons.MapToImmutable(actions)
}
