package wrq

type Job interface {
	Execute() error
	Name() string
}
