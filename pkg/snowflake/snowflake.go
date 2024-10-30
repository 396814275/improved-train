package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)

	return
}
func GenID() int64 {
	return node.Generate().Int64()
}

//func main() {
//	if err := Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
//		fmt.Println(err)
//		fmt.Printf("startime:%s", settings.Conf.StartTime)
//	}
//	id := GenID()
//	fmt.Println(id)
//}
