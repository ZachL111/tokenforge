package policy

const (
	Threshold      = 169
	RiskPenalty   = 5
	LatencyPenalty = 2
	WeightBonus   = 3
)

type Signal struct {
	Demand   int
	Capacity int
	Latency  int
	Risk     int
	Weight   int
}

func Score(signal Signal) int {
	return signal.Demand*2 + signal.Capacity + signal.Weight*WeightBonus - signal.Latency*LatencyPenalty - signal.Risk*RiskPenalty
}

func Classify(signal Signal) string {
	if Score(signal) >= Threshold {
		return "accept"
	}
	return "review"
}
