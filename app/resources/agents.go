package resources

import (
	"github.com/pencroff/proj2048/app/agent"
	"github.com/pencroff/proj2048/app/helper"
	"time"
)

var gameId = int64(helper.Extract(uint64(time.Now().UTC().UnixNano()), 0, 15))
var PoolAgentInstance = agent.NewPoolAgent(gameId)
var HumanAgentInstance = agent.NewHumanAgent(gameId)
