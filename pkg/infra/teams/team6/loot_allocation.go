package team6

import (
	"infra/game/agent"
	"infra/game/commons"
	"infra/game/state"
	"math/rand"

	"github.com/benbjohnson/immutable"
)

// TODO
func (t6 *Team6Agent) LootAllocation(ba agent.BaseAgent) immutable.Map[commons.ID, immutable.SortedMap[commons.ItemID, struct{}]] {
	lootAllocation := make(map[commons.ID][]commons.ItemID)
	view := ba.View()
	ids := commons.ImmutableMapKeys(view.AgentState())
	iterator := ba.Loot().Weapons().Iterator()
	allocateRandomly(iterator, ids, lootAllocation)
	iterator = ba.Loot().Shields().Iterator()
	allocateRandomly(iterator, ids, lootAllocation)
	iterator = ba.Loot().HpPotions().Iterator()
	allocateRandomly(iterator, ids, lootAllocation)
	iterator = ba.Loot().StaminaPotions().Iterator()
	allocateRandomly(iterator, ids, lootAllocation)
	mMapped := make(map[commons.ID]immutable.SortedMap[commons.ItemID, struct{}])
	for id, itemIDS := range lootAllocation {
		mMapped[id] = commons.ListToImmutableSortedSet(itemIDS)
	}
	return commons.MapToImmutable(mMapped)
}

func allocateRandomly(iterator commons.Iterator[state.Item], ids []commons.ID, lootAllocation map[commons.ID][]commons.ItemID) {
	for !iterator.Done() {
		next, _ := iterator.Next()
		toBeAllocated := ids[rand.Intn(len(ids))]
		if l, ok := lootAllocation[toBeAllocated]; ok {
			l = append(l, next.Id())
			lootAllocation[toBeAllocated] = l
		} else {
			l := make([]commons.ItemID, 0)
			l = append(l, next.Id())
			lootAllocation[toBeAllocated] = l
		}
	}
}
