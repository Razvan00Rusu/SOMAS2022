package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"sort"
)

func (r *Perry) HandleElectionBallot(
	b agent.BaseAgent,
	_ *decision.ElectionParams,
) decision.Ballot {
	// Extract ID of alive agents
	view := b.View()
	agentState := view.AgentState()
	aliveAgentIDs := make([]string, agentState.Len())
	i := 0
	itr := agentState.Iterator()
	for !itr.Done() {
		id, a, ok := itr.Next()
		if ok && a.Hp > 0 {
			aliveAgentIDs[i] = id
			i++
		}
	}

	// Randomly fill the ballot
	var ballot decision.Ballot
	var scores []float64
	numAliveAgents := len(aliveAgentIDs)
	for i := 0; i < numAliveAgents; i++ {
		id := aliveAgentIDs[i]
		opinion := r.getOpinion(id)
		votingScore := getVotingScore(opinion)
		ballot = append(ballot, id)
		scores = append(scores, votingScore)
	}

	sort.Slice(ballot, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	return ballot
}
