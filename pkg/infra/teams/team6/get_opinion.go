package team6

func (r *Perry) getOpinion(id string) *PerryOpinion {
	opinion, ok := r.opinions[id]
	if !ok {
		opinion = newOpinion()
		r.opinions[id] = opinion
	}
	return opinion
}
