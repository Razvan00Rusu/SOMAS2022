package team6

import (
	"infra/game/commons"
	"infra/game/decision"
	"infra/game/message"

	"github.com/benbjohnson/immutable"
)

func (r *Team6Agent) HandleFightRequest(
	_ message.TaggedMessage,
	_ *immutable.Map[commons.ID, decision.FightAction],
) message.Payload {
	return nil
}
