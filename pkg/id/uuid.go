package id

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"time"
)

var node *snowflake.Node

func GenId() int64 {
	return node.Generate().Int64()
}

func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	fmt.Println("st:", st)
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)
	return err
}

func GenUUID() uuid.UUID {
	// 生成一个 Version 4 UUID（随机生成）
	newUUID := uuid.New()
	return newUUID
}
