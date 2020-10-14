package producer

import (
	"fmt"
	"github.com/JackMaarek/cqrsPattern/application/conf"
	"github.com/JackMaarek/cqrsPattern/chore/es/event_store"
	"github.com/go-redis/redis/v8"
	"time"
)

const snapshotKey = "snapshot"

type Snapshot interface {
	ProduceEvent(e *estore.Event) (*estore.Event, error)
	ConsumeEvent() (interface{}, error)
	SnapshotEvent(e *estore.Event) error
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

func (r *redisClient) ProduceEvent(e *estore.Event) (*estore.Event, error) {
	json, err := e.MarshalBinary()
	strCMD := r.client.XAdd(conf.Ctx,
		&redis.XAddArgs{
			Stream: "events",
			Values: map[string]interface{}{
				"type": json,
			},
		})
	newId, err := strCMD.Result()
	fmt.Println(newId)
	if err != nil {
		fmt.Println("event error: %v", err)
		return nil, err
	} else {
		fmt.Println("event streamed: %v", newId)
		e.SetID(newId)
		return e, nil
	}
}

func (r *redisClient) ConsumeEvent() (interface{}, error){
	//_, err := r.client.XGroupCreate(conf.Ctx, "events", "test", "$").Result()
	var el []*estore.Event
	status, _ := r.client.XRange(conf.Ctx, "events", "-", "+").Result()
	for _, st := range status{
		var e estore.Event
		e.SetID(st.ID)
		//e.SetType(st.Values)
		el = append(el, &e)
	}
	fmt.Println(el)
	return nil, nil
}


func (r *redisClient) SnapshotEvent(e *estore.Event) error {
	json, err:= e.MarshalBinary()
	snap, err := r.client.Set(conf.Ctx, snapshotKey, json, time.Hour*1).Result()
	fmt.Println(snap)
	if err != nil {
		return err
	}
	return nil
}

func (r *redisClient) RestoreEvent() error {
	var e estore.Event
	snap, err := r.client.Get(conf.Ctx, snapshotKey).Result()
	if err != nil {
		return err
	}
	err  = e.UnmarshalBinary([]byte(snap))
	if err != nil {
		return err
	}
	return nil
}
