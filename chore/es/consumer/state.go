package consumer

import "encoding/json"

type State struct {
	LastEventID uint64
	UserId      uint64
}

func (s *State) GetEventType() uint64 {
	return s.LastEventID
}

func (s *State) SetEventType(id uint64) {
	s.LastEventID = id
}

func (s *State) GetUserId() uint64 {
	return s.UserId
}

func (s *State) SetUserId(id uint64) {
	s.UserId = id
}

func (s *State) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *State) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	return nil
}
