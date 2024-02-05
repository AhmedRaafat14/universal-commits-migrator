package interfaces

// Destination is an interface that all destination control providers implement
type Destination interface {
	PushCommits([]string) error
}
