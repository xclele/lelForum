package snowflake

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func Init(machineId uint16) (err error) {
	sonyMachineID = machineId
	//start time of sonyflake calculator
	t, _ := time.Parse("2007-03-02", "2020-07-01")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GetID returns the generated snowflake user ID
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}

	id, err = sonyFlake.NextID()
	return
}
