package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message"
	"infra/logging"

	"github.com/benbjohnson/immutable"
)

// TODO
func (r *Team6Agent) HandleFightInformation(
	msg message.TaggedInformMessage[message.FightInform],
	baseAgent agent.BaseAgent,
	_ *immutable.Map[commons.ID, decision.FightAction],
) {
	baseAgent.Log(
		logging.Trace,
		logging.LogField{"bravery": r.bravery, "hp": baseAgent.AgentState().Hp},
		"Cowering")

	prop := r.FightResolution(baseAgent)
	_ = baseAgent.SendFightProposalToLeader(prop)
}
