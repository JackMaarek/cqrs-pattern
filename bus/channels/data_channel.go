package channels

import "github.com/JackMaarek/cqrsPattern/structs"

type DataChannel chan structs.Command

var C = make(DataChannel)

type DataChannelSlice []DataChannel
