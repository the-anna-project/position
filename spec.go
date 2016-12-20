package position

type Service interface {
	Default() (string, error)
}
