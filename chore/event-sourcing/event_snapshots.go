package event_sourcing

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

const snapshotKey = "snapshot"

type Snapshot interface {
	Snapshot(t interface{}) error
	Restore(t interface{}) error
}

func NewRedisSnapshot(
	client *redis.Client,
) Snapshot {
	return &redisSnapshot{client: client}
}

type redisSnapshot struct {
	client *redis.Client
}

func (r *redisSnapshot) Snapshot(t interface{}) error {
	_, err := r.client.Set(ctx, snapshotKey, t, time.Hour*1).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisSnapshot) Restore(t interface{}) error {
	_, err := r.client.Get(ctx, snapshotKey).Result()
	if err != nil {
		return err
	}
	return nil
}
