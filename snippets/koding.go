
type Storage interface {
	Get(...) (Kites, error)
	Add(...) error
	Update(...) error
	Delete(...) error
	Upsert(...) error
}
