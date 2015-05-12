type Driver interface {
	Create() error
	GetIP() (string, error)
	Start() error
	Stop() error
	...
}