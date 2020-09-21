package idgenerator

import (
	"github.com/leiqD/go-socket5/pkg/snowflake"
)

var snowWoker *snowflake.Worker

func init() {
	var snowflakeWorkerId int64 = 101 // TODO:  雪花算法WorkerID，当前写死
	snowWoker, _ = snowflake.NewWorker(snowflakeWorkerId)
}

func GetId() int64 {
	return snowWoker.GetId()
}
