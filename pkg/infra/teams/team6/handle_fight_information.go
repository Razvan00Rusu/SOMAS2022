package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"
	"infra/logging"

	"github.com/benbjohnson/immutable"
)

func (r *Perry) HandleFightInformation(
	m message.TaggedInformMessage[message.FightInform],
	baseAgent agent.BaseAgent,
	log *immutable.Map[string, decision.FightAction],
) {
	baseAgent.Log(
		logging.Trace,
		logging.LogField{"bravery": r.bravery, "hp": baseAgent.AgentState().Hp},
		"Cowering")

	prop := r.CreateFightProposal(baseAgent)
	_ = baseAgent.SendFightProposalToLeader(prop)
}
