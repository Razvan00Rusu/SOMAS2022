package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message/proposal"
)

func (r *Team6Agent) FightResolution(baseAgent agent.BaseAgent) commons.ImmutableList[proposal.Rule[decision.FightAction]] {
	// TODO: Implement a set of rules based on HP pool, monster health, etc.

	/*
		Example rule:

		*proposal.NewRule[decision.FightAction](decision.Defend,
			proposal.NewComparativeCondition(proposal.TotalDefence, proposal.GreaterThan, 1000),
		)
	*/

	rules := make([]proposal.Rule[decision.FightAction], 0)
	return *commons.NewImmutableList(rules)
}
