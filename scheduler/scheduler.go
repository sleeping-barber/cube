package scheduler

type Scheduler interface {
	Pick()
	Score()
	SelectCandidate()
}
