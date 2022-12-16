package team6

import (
	"infra/game/agent"
	"infra/game/message"
)

func (r *Perry) HandleLootInformation(
	m message.TaggedInformMessage[message.LootInform],
	base agent.BaseAgent,
) {
}
