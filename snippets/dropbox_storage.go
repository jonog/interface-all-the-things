// START 1 OMIT
type Storage interface {
    Get(...) (interface{}, error)
    Set(...) error
    Delete(...) error
    ...
}
// END 1 OMIT

// structs
type CacheOnStorage
type GenericStorage
type LocalMapStorage
type RateLimitedStorage

// START 2 OMIT
func NewCacheOnStorage(cache, storage Storage) Storage
func NewRateLimitedStorage(storage Storage, maxConcurrency int) Storage
// END 2 OMIT
