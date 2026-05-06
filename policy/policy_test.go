package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 90, Capacity: 73, Latency: 8, Risk: 12, Weight: 13}, wantScore: 216, wantDecision: "accept"},
		{name: "case_2", signal: Signal{Demand: 68, Capacity: 75, Latency: 16, Risk: 25, Weight: 13}, wantScore: 93, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 95, Capacity: 82, Latency: 16, Risk: 7, Weight: 12}, wantScore: 241, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
