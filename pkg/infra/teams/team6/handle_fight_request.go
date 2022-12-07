package team6

import (
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message"

	"github.com/benbjohnson/immutable"
)

// TODO
func (r *Team6Agent) HandleFightRequest(m message.TaggedRequestMessage[message.FightRequest], log *immutable.Map[commons.ID, decision.FightAction]) message.FightInform {
	return nil
}
