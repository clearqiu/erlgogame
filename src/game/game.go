package game

import (
	"fmt"
	"setting"
	"util/log"
	"util/common"
)

func Start() {
	fmt.Println("Hello Game", common.Microsecond)
	// logger
	logger, err := log.New(setting.LogLevel, setting.LogPath, setting.LogFlag)
	if err != nil {
		panic(err)
	}
	log.Export(logger)
	defer logger.Close()
	log.Release("xgame %v starting up", "sssss")
}