package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"
)

const ABSTAIN_THRESHOLD = 0.5
const SIMILARITY_THRESHOLD = 0.75

func (r *Perry) HandleFightProposal(
	prop message.Proposal[decision.FightAction],
	base agent.BaseAgent,
) decision.Intent {
	rules := prop.Rules()
	similarity := ProposalSimilarity[decision.FightAction](rules, r.nextFights)
	if similarity >= SIMILARITY_THRESHOLD {
		return decision.Positive
	} else if similarity >= ABSTAIN_THRESHOLD {
		return decision.Abstain
	} else {
		return decision.Negative
	}
}
