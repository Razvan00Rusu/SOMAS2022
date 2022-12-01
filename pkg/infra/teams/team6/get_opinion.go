package team6

func (r *Team6Agent) getOpinion(id string) *Team6AgentOpinion {
	opinion, ok := r.opinions[id]
	if !ok {
		opinion = newOpinion()
		r.opinions[id] = opinion
	}
	return opinion
}
