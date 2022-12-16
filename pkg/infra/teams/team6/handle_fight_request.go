package team6

import (
	"infra/game/decision"
	"infra/game/message"

	"github.com/benbjohnson/immutable"
)

func (r *Perry) HandleFightRequest(
	msg message.TaggedRequestMessage[message.FightRequest],
	log *immutable.Map[string, decision.FightAction],
) message.FightInform {
	// What does this do?
	return nil
}
