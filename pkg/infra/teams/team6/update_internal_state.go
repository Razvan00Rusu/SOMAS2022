package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"

	"github.com/benbjohnson/immutable"
)

func (t6 *Team6Agent) UpdateInternalState(baseAgent agent.BaseAgent, fightResult *commons.ImmutableList[decision.ImmutableFightResult], voteResult *immutable.Map[decision.Intent, uint]) {

}
