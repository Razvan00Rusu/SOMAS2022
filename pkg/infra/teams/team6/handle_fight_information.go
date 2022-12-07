package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message"
	"infra/game/message/proposal"
	"infra/logging"

	"github.com/benbjohnson/immutable"
)

// TODO
func (t6 *Team6Agent) HandleFightInformation(
	msg message.TaggedInformMessage[message.FightInform],
	baseAgent agent.BaseAgent,
	_ *immutable.Map[commons.ID, decision.FightAction],
) {
	baseAgent.Log(
		logging.Trace,
		logging.LogField{"bravery": t6.bravery, "hp": baseAgent.AgentState().Hp},
		"Cowering")

	rules := *commons.NewImmutableList(make([]proposal.Rule[decision.FightAction], 0))
	_ = baseAgent.SendFightProposalToLeader(rules)
}
