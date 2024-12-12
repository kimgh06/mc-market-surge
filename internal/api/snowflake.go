package api

import (
	"fmt"
	"github.com/godruoyi/go-snowflake"
	"github.com/sirupsen/logrus"
	"surge/internal/conf"
	"time"
)

var SnowflakeStartTime = time.Date(2024, 11, 0, 0, 0, 0, 0, time.UTC)

func InitSnowflake(config *conf.SurgeSnowflakeConfigurations) {
	startTime := SnowflakeStartTime
	if config.StartTime != "" {
		parsed, err := time.Parse(time.RFC3339, config.StartTime)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse SNOWFLAKE_START_TIME: %+v ", err))
		} else {
			startTime = parsed
		}
	}

	machineID := snowflake.PrivateIPToMachineID()
	if config.MachineID != -1 {
		machineID = uint16(config.MachineID)
	}

	logrus.WithField("core", "snowflake").Infoln("Using StartTime: ", startTime)
	logrus.WithField("core", "snowflake").Infoln("Using MachineID: ", machineID)

	snowflake.SetStartTime(startTime)
	snowflake.SetMachineID(machineID)
}
