package resources

import (
	"github.com/pencroff/ai-agents/proj2048/agent"
	"github.com/pencroff/ai-agents/proj2048/helper"
	"time"
)

var gameId = int64(helper.Extract(uint64(time.Now().UTC().UnixNano()), 0, 15))
var PoolAgentInstance = agent.NewPoolAgent(gameId)
var HumanAgentInstance = agent.NewHumanAgent(gameId)
