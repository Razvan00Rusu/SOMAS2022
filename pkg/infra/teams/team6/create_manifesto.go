package team6

import (
	"infra/game/agent"
	"infra/game/decision"
)

func (r *Team6Agent) CreateManifesto(_ agent.BaseAgent) *decision.Manifesto {
	manifesto := decision.NewManifesto(true, false, 10, 50)
	return manifesto
}
