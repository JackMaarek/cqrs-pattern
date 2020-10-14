package es

import (
	"fmt"
	"github.com/JackMaarek/cqrsPattern/application/conf"
	"github.com/JackMaarek/cqrsPattern/chore/es/event_store"
	"github.com/JackMaarek/cqrsPattern/chore/es/producer"
)

func NewOrderCreatedEvent(d interface{}) error {
	var err error
	var e *estore.Event
	if e, err = NewEvent(OrderCreated); err != nil {
		return err
	} else {
		snapshot := producer.NewRedisClient(conf.RedisClient)
		e, err = snapshot.ProduceEvent(e)
		if err = snapshot.SnapshotEvent(e); err != nil {
			fmt.Println(err)
			return err
		}
		_, err = snapshot.ConsumeEvent()
		return nil
	}
}

/*func NewUserUpdatedEvent() error {
	var err error
	var e estore.Event
	if e , err = NewEvent(UserUpdated); err != nil {
		return err
	} else {
		snapshot := producer.NewRedisClient(conf.RedisClient)
		err = snapshot.ProduceEvent(&e)
		fmt.Println(err)
	}
	return nil
}*/
