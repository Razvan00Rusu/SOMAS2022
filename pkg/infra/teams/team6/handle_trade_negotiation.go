package team6

import (
	"infra/game/agent"
	"infra/game/message"
)

func (r *Perry) HandleTradeNegotiation(
	base agent.BaseAgent,
	info message.TradeInfo,
) message.TradeMessage {
	// TODO
	return message.TradeReject{}
}
