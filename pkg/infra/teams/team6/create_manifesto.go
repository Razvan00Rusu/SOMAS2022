package team6

import (
	"infra/game/agent"
	"infra/game/decision"
)

func (r *Perry) CreateManifesto(_ agent.BaseAgent) *decision.Manifesto {
	manifesto := decision.NewManifesto(true, false, 10, 100)
	return manifesto
}
