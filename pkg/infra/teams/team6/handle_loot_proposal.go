package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"
)

const ABSTAIN_THRESHOLD_LOOT = 0.5
const SIMILARITY_THRESHOLD_LOOT = 0.75

func (r *Perry) HandleLootProposal(
	prop message.Proposal[decision.LootAction],
	base agent.BaseAgent,
) decision.Intent {
	rules := prop.Rules()
	similarity := ProposalSimilarity[decision.LootAction](rules, r.nextLoot)
	if similarity >= SIMILARITY_THRESHOLD_LOOT {
		return decision.Positive
	} else if similarity >= ABSTAIN_THRESHOLD_LOOT {
		return decision.Abstain
	} else {
		return decision.Negative
	}
}
