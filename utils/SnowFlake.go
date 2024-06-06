package utils

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	// 设置 Snowflake 算法的起始时间为某个固定的时间点，例如：2020-01-01 00:00:00 UTC
	var st time.Time
	st, _ = time.Parse("2006-01-02 15:04:05", "2020-01-01 00:00:00")
	snowflake.Epoch = st.UnixNano() / 1000000

	// 假设机器ID为1，这里可以根据实际情况进行设置
	var machineID int64 = 1
	var err error
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		panic(err)
	}
}

func GenerateID() int64 {
	// 直接使用预先创建的node生成ID
	return node.Generate().Int64()
}
