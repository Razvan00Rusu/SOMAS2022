package team6

const GENEROSITY_BIAS = 0.05
const BRAVERY_BIAS = 0.05
const SIMILARITY_BIAS = 0.25
const LEADERSHIP_BIAS = 0.65

func getVotingScore(opinion *PerryOpinion) float64 {
	generosity := float64(opinion.generosity) * GENEROSITY_BIAS
	bravery := float64(opinion.bravery) * BRAVERY_BIAS
	similarity := float64(opinion.similarity) * SIMILARITY_BIAS
	leadership := float64(opinion.leadership) * LEADERSHIP_BIAS

	return generosity + similarity + bravery + leadership
}
