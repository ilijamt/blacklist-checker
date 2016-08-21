package queue

// Job queue interface
type Job interface {
	Start(finished chan<- bool)
}
