package data

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"

)

var ctx = context.Background()

var Datasource Repository

type KeyValue struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

func init() {
	db := redis.NewClient(&redis.Options{
		//Addr:     "redis-master:6379",
		Addr: "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	Datasource= &DataAccess{
		Client: db,
	}
}


type Repository interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
}

type DataAccess struct {
	Client *redis.Client
}

// Set attaches the redis repository and set the data
func (r *DataAccess) Set(key string, value interface{}, exp time.Duration) error {
	return r.Client.Set(ctx,key, value, exp).Err()
}

// Get attaches the redis repository and get the data
func (r *DataAccess) Get(key string) (string, error) {
	get := r.Client.Get(ctx,key)
	return get.Result()
}