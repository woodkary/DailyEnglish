package utils

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

func GenerateID(st time.Time, machineID int64) int64 {

	sf.Epoch = st.UnixNano() / 1000000
	node, _ := sf.NewNode(machineID)
	return node.Generate().Int64()
}
