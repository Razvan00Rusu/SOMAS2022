package discussion

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message/proposal"
	"infra/game/state"
	"infra/game/tally"

	"github.com/benbjohnson/immutable"
)

func ResolveFightDiscussion(gs state.State, agentMap map[commons.ID]agent.Agent, currentLeader agent.Agent, manifesto decision.Manifesto, tally *tally.Tally[decision.FightAction]) decision.FightResult {
	fightActions := make(map[commons.ID]decision.FightAction)
	// todo: cleanup the nil check that acts to check if the leader died in combat
	var prop commons.ImmutableList[proposal.Rule[decision.FightAction]]
	if manifesto.FightImposition() && currentLeader.Strategy != nil {
		prop = currentLeader.Strategy.FightResolution(*currentLeader.BaseAgent)
	} else {
		// get proposal with most votes
		prop = tally.GetMax().Rules()
	}

	predicate := proposal.ToSinglePredicate(prop)
	if predicate == nil {
		for id, a := range agentMap {
			fightActions[id] = a.FightAction(*a.BaseAgent)
		}
	} else {
		for id, a := range agentMap {
			fightActions[id] = predicate(gs, a.AgentState())
		}
	}

	return decision.FightResult{
		Choices:         fightActions,
		AttackingAgents: nil,
		ShieldingAgents: nil,
		CoweringAgents:  nil,
		AttackSum:       0,
		ShieldSum:       0,
	}
}

func ResolveLootDiscussion(gs state.State, agentMap map[commons.ID]agent.Agent, pool *state.LootPool, leader agent.Agent, manifesto decision.Manifesto, tally *tally.Tally[decision.LootAction]) immutable.Map[commons.ID, immutable.List[commons.ItemID]] {
	if manifesto.LootImposition() {
		// todo: eliminate double allocations here
		return leader.Strategy.LootAllocation(*leader.BaseAgent)
	} else {
		predicate := proposal.ToMultiPredicate(tally.GetMax().Rules())
		getsWeapon := make([]commons.ID, 0)
		getsShield := make([]commons.ID, 0)
		getsHealthPotion := make([]commons.ID, 0)
		getsStaminaPotion := make([]commons.ID, 0)
		for id := range agentMap {
			actions := predicate(gs, gs.AgentState[id])
			if _, ok := actions[decision.Weapon]; ok {
				getsWeapon = append(getsWeapon, id)
			}
			if _, ok := actions[decision.Shield]; ok {
				getsShield = append(getsShield, id)
			}
			if _, ok := actions[decision.HealthPotion]; ok {
				getsHealthPotion = append(getsHealthPotion, id)
			}
			if _, ok := actions[decision.StaminaPotion]; ok {
				getsStaminaPotion = append(getsStaminaPotion, id)
			}
		}
		m := make(map[commons.ID][]commons.ItemID)
		buildAllocation(pool.Weapons(), getsWeapon, m)
		buildAllocation(pool.Shields(), getsShield, m)
		buildAllocation(pool.HpPotions(), getsHealthPotion, m)
		buildAllocation(pool.StaminaPotions(), getsStaminaPotion, m)
		mMapped := make(map[commons.ID]immutable.List[commons.ItemID])
		for id, itemIDS := range m {
			mMapped[id] = commons.ListToImmutable(itemIDS)
		}
		return commons.MapToImmutable(mMapped)
	}
}

func buildAllocation(pool *commons.ImmutableList[state.Item], proposedLooters []commons.ID, allocation map[commons.ID][]commons.ItemID) {
	idx := 0
	iterator := pool.Iterator()
	for !iterator.Done() {
		if idx >= len(proposedLooters) {
			break
		}
		next, _ := iterator.Next()
		if l, ok := allocation[proposedLooters[idx]]; ok {
			l = append(l, next.Id())
			allocation[proposedLooters[idx]] = l
		} else {
			l := make([]commons.ItemID, 0)
			l = append(l, next.Id())
			allocation[proposedLooters[idx]] = l
		}
	}
}
