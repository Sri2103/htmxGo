package sessionStore

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type UserSession struct {
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `json:"user_id"`
	Email     string    `json:"email"`
}

type redisStore struct {
	Client *redis.Client
}

func CustomRedisStore() *redisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_url"),
		Password: os.Getenv("redis_password"),
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to ping Redis: %v", err)
	}

	return &redisStore{
		Client: client,
	}
}

// Implement the SessionStore interface for our custom store
type SerializableStore interface {
	Get(string) (UserSession, error)
	Set(string, string) error
}

func (r redisStore) Get(id string) (UserSession, error) {
	sessionString, err := r.Client.Get(id).Result()
	if err == redis.Nil {
		return UserSession{}, redis.Nil
	} else if err != nil {
		return UserSession{}, errors.Wrap(err, "Error retrieving user session")
	}
	var session UserSession
	err = json.Unmarshal([]byte(sessionString), &session)
	return session, err
}

func (r redisStore) Set(key string, value string) error {
	j, _ := json.Marshal(value)
	return r.Client.Set(key, j, 0).Err()
}
