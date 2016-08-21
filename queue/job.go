package queue

// Job queue interface
type Job interface {
	Start(response chan<- interface{})
}
