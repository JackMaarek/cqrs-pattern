package events

import "encoding/json"

type TopUp struct {
	*Event
	UserID uint64
	Amount uint64
}

func (t *TopUp) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *TopUp) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	return nil
}
