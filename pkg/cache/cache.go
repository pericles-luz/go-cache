package cache

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string, durationInMinutes int) error
	Del(key string) error
	Ping() error
}
