package policy

import "testing"

func TestDomainReviewLane(t *testing.T) {
	item := DomainReview{Signal: 57, Slack: 46, Drag: 13, Confidence: 93}
	if got := DomainReviewScore(item); got != 214 {
		t.Fatalf("score = %d", got)
	}
	if got := DomainReviewLane(item); got != "ship" {
		t.Fatalf("lane = %s", got)
	}
}
