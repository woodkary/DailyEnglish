package utils

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(st time.Time, machineID int64) (err error, node *sf.Node) {
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return nil, node
}

func GenerateID() int64 {
	st := time.Now()
	_, node := Init(st, 1145141919810)
	return node.Generate().Int64()
}
