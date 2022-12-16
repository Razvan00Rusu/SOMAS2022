package team6

type PerryOpinion struct {
	generosity int16 // Likeliness to accept trades beneficial to us
	bravery    int16 // Likeliness to fight
	similarity int16 // Likeliness to fight when we would fight
	leadership int16 // Success rating if the agent has been a leader
}

func newOpinion() *PerryOpinion {
	opinion := PerryOpinion{}
	return &opinion
}

func (r *PerryOpinion) AddGenerosity(i int) {
	generosity := int(r.generosity) + i
	if generosity > 32767 {
		r.generosity = 32767
	} else if generosity < -32768 {
		r.generosity = -32768
	} else {
		r.generosity = int16(generosity)
	}
}

func (r *PerryOpinion) AddBravery(i int) {
	bravery := int(r.bravery) + i
	if bravery > 32767 {
		r.bravery = 32767
	} else if bravery < -32768 {
		r.bravery = -32768
	} else {
		r.bravery = int16(bravery)
	}
}

func (r *PerryOpinion) AddSimilarity(i int) {
	similarity := int(r.similarity) + i
	if similarity > 32767 {
		r.similarity = 32767
	} else if similarity < -32768 {
		r.similarity = -32768
	} else {
		r.similarity = int16(similarity)
	}
}

func (r *PerryOpinion) AddLeadership(i int) {
	leadership := int(r.leadership) + i
	if leadership > 32767 {
		r.leadership = 32767
	} else if leadership < -32768 {
		r.leadership = -32768
	} else {
		r.leadership = int16(leadership)
	}
}