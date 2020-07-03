package worker

// Job ...
type Job interface {
	Exec() string
	Name() string
}
