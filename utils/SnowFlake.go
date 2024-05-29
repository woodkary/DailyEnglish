package utils

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

func GenerateID(st time.Time, machineID int64) int64 {

	sf.Epoch = st.UnixNano() / 1000000
	node, _ := sf.NewNode(machineID)
	time.Sleep(3 * time.Millisecond)
	return node.Generate().Int64()
}

func TestGenerateID() {
	st := time.Now()
	var machineID int64 = 1
	id1 := GenerateID(st, machineID)
	id2 := GenerateID(st, machineID)
	id3 := GenerateID(st, machineID)
	println(id1, id2, id3)
	id4 := GenerateID(st, machineID)
	println(id4)
}
