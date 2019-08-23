package setting

import (
	"util/common"
)
// 默认配置
var (
	// log conf
	LogFlag = 3333

	// gate conf
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 4096
	HTTPTimeout            = 10 * common.Second
	LenMsgLen              = 2
	LittleEndian           = false

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000

	LenStackBuf = 4096

	// log
	LogLevel ="debug"
	LogPath  =""

	// console
	ConsolePort   int
	ConsolePrompt string = "Xgame# "
	ProfilePath   string

	// cluster
	ListenAddr      string
	ConnAddrs       []string
)

