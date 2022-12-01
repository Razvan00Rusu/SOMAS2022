package team6

type ActionDecision struct {
	cower  uint8
	defend uint8
	attack uint8
}

func NewActionDecision() ActionDecision {
	return ActionDecision{
		cower:  uint8(127),
		defend: uint8(127),
		attack: uint8(127),
	}
}
