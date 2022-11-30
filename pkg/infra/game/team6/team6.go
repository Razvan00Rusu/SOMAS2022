package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message"
	"infra/game/state"
	"infra/logging"
	"math/rand"

	"github.com/benbjohnson/immutable"
)

type Team6Agent struct {
	bravery      uint
	fightUtility MapDefault // rank of agents 0-100
}

func (t6 *Team6Agent) CreateManifesto(view *state.View, baseAgent agent.BaseAgent) *decision.Manifesto {
	manifesto := decision.NewManifesto(true, false, 10, 50)
	return manifesto
}

func (t6 *Team6Agent) HandleConfidencePoll(view *state.View, baseAgent agent.BaseAgent) decision.Intent {
	switch rand.Intn(3) {
	case 0:
		return decision.Abstain
	case 1:
		return decision.Negative
	default:
		return decision.Positive
	}
}

func (t6 *Team6Agent) HandleFightInformation(message message.TaggedMessage, view *state.View, agent agent.BaseAgent, log *immutable.Map[commons.ID, decision.FightAction]) {
	senderPreviousAction, ok := log.Get(message.Sender())
	if !ok {
		agent.Log(logging.Debug, logging.LogField{"senderID": message.Sender()}, "Message sender not in log")
	}

	agentState := view.AgentState()
	senderState, ok2 := agentState.Get(message.Sender())
	if !ok2 {
		agent.Log(logging.Debug, logging.LogField{"senderID": message.Sender()}, "Message sender not in agentState view")
	}
	var fightUtilityMultiplier float32

	switch senderState.Hp {
	case state.HealthRange(state.LowHealth):
		switch senderPreviousAction {
		case decision.Attack:
			fightUtilityMultiplier = 1.25
		case decision.Defend:
			fightUtilityMultiplier = 1.20
		case decision.Cower:
			fightUtilityMultiplier = 1.0 // their expected action
		}
	case state.HealthRange(state.MidHealth):
		switch senderPreviousAction {
		case decision.Attack:
			fightUtilityMultiplier = 1.05
		case decision.Defend:
			fightUtilityMultiplier = 1.05
		case decision.Cower:
			fightUtilityMultiplier = 0.95
		}
	case state.HealthRange(state.HighHealth):
		switch senderPreviousAction {
		case decision.Attack:
			fightUtilityMultiplier = 1.05
		case decision.Defend:
			fightUtilityMultiplier = 1.05
		case decision.Cower:
			fightUtilityMultiplier = 0.85
		}
	}
	newUtility := uint(float32(t6.fightUtility.GetValOrDefault(message.Sender(), 70)) * fightUtilityMultiplier)
	t6.fightUtility[message.Sender()] = newUtility

	agent.Log(logging.Trace, logging.LogField{"bravery": t6.bravery, "hp": agent.ViewState().Hp, "senderID": message.Sender(), "senderUtility": newUtility}, "Cowering")
}

func (t6 *Team6Agent) HandleFightRequest(_ message.TaggedMessage, _ *state.View, _ *immutable.Map[commons.ID, decision.FightAction]) message.Payload {
	return nil
}

func (t6 *Team6Agent) CurrentAction() decision.FightAction {
	fight := rand.Intn(3)
	switch fight {
	case 0:
		return decision.Cower
	case 1:
		return decision.Attack
	default:
		return decision.Defend
	}
}

func (t6 *Team6Agent) HandleElectionBallot(view *state.View, _ agent.BaseAgent, _ *decision.ElectionParams) decision.Ballot {
	// Extract ID of alive agents
	agentState := view.AgentState()
	aliveAgentIds := make([]string, agentState.Len())
	i := 0
	itr := agentState.Iterator()
	for !itr.Done() {
		id, a, ok := itr.Next()
		if ok && a.Hp > 0 {
			aliveAgentIds[i] = id
			i++
		}
	}

	// Randomly fill the ballot
	var ballot decision.Ballot
	numAliveAgents := len(aliveAgentIds)
	numCandidate := 2
	for i := 0; i < numCandidate; i++ {
		randomIdx := rand.Intn(numAliveAgents)
		randomCandidate := aliveAgentIds[uint(randomIdx)]
		ballot = append(ballot, randomCandidate)
	}

	return ballot
}

func NewTeam6Agent() agent.Strategy {
	return &Team6Agent{bravery: 0, fightUtility: make(map[commons.ID]uint)}
}
