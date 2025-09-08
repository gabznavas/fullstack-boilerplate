package cache

import (
	"backend/internal/todo/models"
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type TodoCache interface {
	GetAllTodos(ctx context.Context) ([]*models.Todo, error)
	PushOneTodo(ctx context.Context, value *models.Todo) error
	PushAllTodos(ctx context.Context, value []*models.Todo) error
}

type RedisCacheImpl struct {
	allTodosKey string
	expiration  time.Duration
	redis       *redis.Client
}

func NewRedisCache(redisClient *redis.Client, allTodosKey string, expiration time.Duration) TodoCache {
	return &RedisCacheImpl{redis: redisClient, allTodosKey: allTodosKey, expiration: expiration}
}

func (c *RedisCacheImpl) GetAllTodos(ctx context.Context) ([]*models.Todo, error) {
	todos := []*models.Todo{}
	data := c.redis.Get(ctx, c.allTodosKey)
	if data.Err() != nil {
		if data.Err() == redis.Nil {
			return []*models.Todo{}, nil
		}
		return nil, data.Err()
	}
	json.Unmarshal([]byte(data.Val()), &todos)
	return todos, nil
}

func (c *RedisCacheImpl) PushOneTodo(ctx context.Context, todo *models.Todo) error {
	todos := []*models.Todo{}
	dataSet := []byte{}
	err := error(nil)

	dataGet := c.redis.Get(ctx, c.allTodosKey)
	if dataGet.Err() != nil {
		if dataGet.Err() == redis.Nil {
			todos = append(todos, todo)
			dataSet, err = json.Marshal(todos)
			if err != nil {
				return err
			}
		} else {
			return dataGet.Err()
		}
	}

	json.Unmarshal([]byte(dataGet.Val()), &todos)
	todos = append(todos, todo)
	dataSet, err = json.Marshal(todos)

	return c.redis.Set(ctx, c.allTodosKey, dataSet, c.expiration).Err()
}

func (c *RedisCacheImpl) PushAllTodos(ctx context.Context, todos []*models.Todo) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return c.redis.Set(ctx, c.allTodosKey, data, c.expiration).Err()
}
