package event

import (
	"github.com/JackMaarek/cqrsPattern/application/conf"
	"github.com/JackMaarek/cqrsPattern/application/events"
)

func CreateEventSnapshot(topup *events.TopUp) error {
	snapshot := events.NewRedisSnapshot(conf.RedisClient)
	binVal, err := topup.MarshalBinary()
	err = snapshot.Snapshot(binVal)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveEventSnapshot(topup *events.TopUp) error {
	snapshot := events.NewRedisSnapshot(conf.RedisClient)
	err := snapshot.Restore(topup)
	if err != nil {
		return err
	}
	return nil
}
