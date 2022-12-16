package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/logging"

	"github.com/benbjohnson/immutable"
)

func (r *Perry) UpdateInternalState(
	base agent.BaseAgent,
	fightResult *commons.ImmutableList[decision.ImmutableFightResult],
	_ *immutable.Map[decision.Intent, uint],
	_ chan<- logging.AgentLog,
) {
	panic("Update State :: Not Implemented")
}
