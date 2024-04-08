package snowflake

import (
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	startTime = time.Now().Format("2006-01-02")
	st, err = time.Parse("2006-01-02", startTime)
	fmt.Println(st)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return err
}

func GenID() int64 {
	return node.Generate().Int64()
}
