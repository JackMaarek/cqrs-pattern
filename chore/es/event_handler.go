package es

import (
	"fmt"
	"github.com/JackMaarek/cqrsPattern/application/conf"
	"github.com/JackMaarek/cqrsPattern/chore/es/producer"
)

func NewUserCreatedEvent(d interface{}) error {
	var err error
	var e *Event
	if e , err = NewEvent(UserCreated); err != nil {
		return err
	} else {
		e.SetData(d)
		fmt.Println(e.GetData())
		snapshot := producer.NewRedisClient(conf.RedisClient)
		//err = snapshot.SnapshotEvent(e)
		//err = snapshot.RestoreEvent()
		err = snapshot.StreamEvent(e)
		fmt.Println(err)
	}
	return nil
}

func NewUserUpdatedEvent(d interface{}) error {
	var err error
	var e *Event
	if e , err = NewEvent(UserUpdated); err != nil {
		return err
	} else {
		e.SetData(d)
		snapshot := producer.NewRedisClient(conf.RedisClient)
		//err = snapshot.SnapshotEvent(e)
		//err = snapshot.RestoreEvent()
		err = snapshot.StreamEvent(e)
		fmt.Println(err)
	}
	return nil
}
