package event

import (
	"github.com/JackMaarek/cqrsPattern/application/conf"
	"github.com/JackMaarek/cqrsPattern/application/events"
	es "github.com/JackMaarek/cqrsPattern/chore/event-sourcing"
)

func CreateEventSnapshot(topup *es.TopUp) error {
	snapshot := es.NewRedisSnapshot(conf.RedisClient)
	binVal, err := topup.MarshalBinary()
	err = snapshot.Snapshot(binVal)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveEventSnapshot(topup *es.TopUp) error {
	snapshot := es.NewRedisSnapshot(conf.RedisClient)
	err := snapshot.Restore(topup)
	if err != nil {
		return err
	}
	return nil
}
