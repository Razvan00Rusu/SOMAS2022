package team6

type ActionDecision struct {
	cower  uint16
	defend uint16
	attack uint16
}

func NewActionDecision() ActionDecision {
	return ActionDecision{
		cower:  uint16(32767),
		defend: uint16(32767),
		attack: uint16(32767),
	}
}
