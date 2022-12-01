package team6

type Team6AgentOpinion struct {
	generosity int8 // Likeliness to accept trades beneficial to us
	bravery    int8 // Likeliness to fight
	similarity int8 // Likeliness to fight when we would fight
	leadership int8 // Success rating if the agent has been a leader
}

func newOpinion() *Team6AgentOpinion {
	opinion := Team6AgentOpinion{}
	return &opinion
}

func (r *Team6AgentOpinion) AddGenerosity(i int) {
	generosity := int(r.generosity) + i
	if generosity > 127 {
		r.generosity = 127
	} else if generosity < -128 {
		r.generosity = -128
	} else {
		r.generosity = int8(generosity)
	}
}

func (r *Team6AgentOpinion) AddBravery(i int) {
	bravery := int(r.bravery) + i
	if bravery > 127 {
		r.bravery = 127
	} else if bravery < -128 {
		r.bravery = -128
	} else {
		r.bravery = int8(bravery)
	}
}

func (r *Team6AgentOpinion) AddSimilarity(i int) {
	similarity := int(r.similarity) + i
	if similarity > 127 {
		r.similarity = 127
	} else if similarity < -128 {
		r.similarity = -128
	} else {
		r.similarity = int8(similarity)
	}
}

func (r *Team6AgentOpinion) AddLeadership(i int) {
	leadership := int(r.leadership) + i
	if leadership > 127 {
		r.leadership = 127
	} else if leadership < -128 {
		r.leadership = -128
	} else {
		r.leadership = int8(leadership)
	}
}
