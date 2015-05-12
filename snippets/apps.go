// START 1 OMIT
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Infof(args ...interface{})
	Warnf(args ...interface{})
	Errorf(args ...interface{})
}

// END 1 OMIT

// START 2 OMIT

type TokenStore interface {
	Generate(id int64) (string, error)
	Revoke(token string) error
	Get(token string) (int64, error)
}

type Publisher interface {
	Publish(queueName string, payload interface{}) error
}

type Mailer interface {
	Mail(interface{}) error
}

// END 2 OMIT