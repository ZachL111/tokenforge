package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 90, Capacity: 73, Latency: 8, Risk: 12, Weight: 13}
	if got := Score(signal); got != 216 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 68, Capacity: 75, Latency: 16, Risk: 25, Weight: 13}
	if got := Score(signal); got != 93 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 95, Capacity: 82, Latency: 16, Risk: 7, Weight: 12}
	if got := Score(signal); got != 241 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
