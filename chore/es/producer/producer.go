package producer

import (
	"context"
	"fmt"
	"github.com/JackMaarek/cqrsPattern/application/conf"
	"github.com/JackMaarek/cqrsPattern/chore/es"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

var ctx = context.Background()

const snapshotKey = "snapshot"

type Snapshot interface {
	StreamEvent(e *es.Event) error
	SnapshotEvent(e *es.Event) error
	RestoreEvent() error
}

func NewRedisClient(
	client *redis.Client,
) Snapshot {
	return &redisClient{client: client}
}

type redisClient struct {
	client *redis.Client
}

func (r *redisClient) StreamEvent(e *es.Event) error {
	json, err := e.MarshalBinary()
	if err != nil {
		fmt.Println("cannot marshal event %v", err)
		return err
	}
	strCMD := r.client.XAdd(conf.Ctx,
		&redis.XAddArgs{
			Stream: strings.ReplaceAll(string(e.Type), " ", ""),
			Values: map[string]interface{}{
				"data": json,
			},
		})
	newId, err := strCMD.Result()
	if err != nil {
		fmt.Println("event error: %v", err)
		return err
	} else {
		fmt.Println("event streamed: %v", newId)
		return nil
	}
}


func (r *redisClient) SnapshotEvent(e *es.Event) error {
	json, err:= e.MarshalBinary()
	_, err = r.client.Set(ctx, snapshotKey, json, time.Hour*1).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisClient) RestoreEvent() error {
	var e es.Event
	snap, err := r.client.Get(ctx, snapshotKey).Result()
	if err != nil {
		return err
	}
	err  = e.UnmarshalBinary([]byte(snap))
	if err != nil {
		return err
	}
	fmt.Println(e.Data)
	return nil
}
