package team6

import (
	"infra/game/agent"
	"infra/game/decision"
)

// proportion of agents we would rather have as leader to vote yes
const VONC_THRESHOLD = 0.5

func (r *Team6Agent) HandleConfidencePoll(
	base agent.BaseAgent,
) decision.Intent {
	view := base.View()
	leader := view.CurrentLeader()
	leaderScore := getVotingScore(r.getOpinion(leader))
	numBetter := 0
	numEq := 0
	agents := view.AgentState()
	it := agents.Iterator()
	for !it.Done() {
		id, _, ok := it.Next()
		if ok {
			score := getVotingScore(r.getOpinion(id))
			if score > leaderScore {
				numBetter++
			}
			if score >= leaderScore {
				numEq++
			}
		}
	}
	if float64(numBetter)/float64(agents.Len()) >= VONC_THRESHOLD {
		return decision.Positive
	}
	if float64(numEq)/float64(agents.Len()) >= VONC_THRESHOLD {
		return decision.Abstain
	}
	return decision.Negative
}
