package agent

import (
	"fmt"
	"github.com/pencroff/ai-agents/proj2048/common"
	"math"
	"strings"
)

type AgentStatRecord struct {
	Score  int
	Step   int
	GameId int64
}

type AgentAvgRecord struct {
	Score int
	Cnt   int
}

type AgentStat struct {
	Count    int
	Min      AgentStatRecord
	AvgScore float64
	Max      AgentStatRecord
}

// PoolAgent Represent a collection of agents which is executed one by one
type PoolAgent struct {
	currentIdx int
	agent      Agent
	pool       []Agent
	statMap    map[string]AgentStat
	statStr    string
}

func (p *PoolAgent) GetId() string {
	return p.agent.GetId()
}

func (p *PoolAgent) GetName() string {
	return p.agent.GetName()
}

func (p *PoolAgent) IsManual() bool {
	return p.agent.IsManual()
}

func (p *PoolAgent) GetGameId() int64 {
	return p.agent.GetGameId()
}

func (p *PoolAgent) GetGameSeed() int64 {
	return p.agent.GetGameSeed()
}

func (p *PoolAgent) MakeMove(step int, score int, noMove bool, lst []int) common.Direction {
	return p.agent.MakeMove(step, score, noMove, lst)
}

func (p *PoolAgent) LogStep(step int, score int, noMove bool, lst []int, d common.Direction) error {
	return p.agent.LogStep(step, score, noMove, lst, d)
}

func (p *PoolAgent) GameFinished(step int, score int, noMove bool, lst []int, d common.Direction) {
	agent := p.agent
	p.countStat(step, score)
	p.statStr = statToStr(p.statMap)
	p.currentIdx += 1
	p.agent = p.pool[p.currentIdx%len(p.pool)]
	agent.GameFinished(step, score, noMove, lst, d)
}

func (p *PoolAgent) GetStatStr() string {
	return p.statStr

}

func (p *PoolAgent) countStat(step int, score int) {
	agentId := p.agent.GetId()
	gameId := p.agent.GetGameId()
	stat, ok := p.statMap[agentId]

	if !ok {
		stat = createAgentStat()
	}

	stat.Count += 1

	// Min
	if stat.Count == 1 || score < stat.Min.Score {
		stat.Min.Step = step
		stat.Min.Score = score
		stat.Min.GameId = gameId
	}
	// Max
	if stat.Count == 1 || score > stat.Max.Score {
		stat.Max.Step = step
		stat.Max.Score = score
		stat.Max.GameId = gameId
	}
	// Avg Score
	currentAvg := stat.AvgScore
	stat.AvgScore = currentAvg + (float64(score)-currentAvg)/float64(stat.Count)
	// to str
	p.statMap[agentId] = stat
}

func statToStr(kvs map[string]AgentStat) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%24s | %12s | %5s | %12s |\n", "", "Min", "Avg", "Max"))
	for k, stat := range kvs {
		sb.WriteString(fmt.Sprintf("%24s | %5d (%4d) | %5.0f | %5d (%4d) |\n",
			k, stat.Min.Score, stat.Min.Step, math.Round(stat.AvgScore), stat.Max.Score, stat.Max.Step))
		sb.WriteString(fmt.Sprintf("%24s | %12d | %5s | %12d |\n",
			"", stat.Min.GameId, "", stat.Max.GameId))
	}
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%24s | %12s | %5s | %12s |\n", "Agenda:", "score", "score", "score"))
	sb.WriteString(fmt.Sprintf("%24s | %12s | %5s | %12s |\n", "", "(in steps)", "", "(in steps)"))
	sb.WriteString(fmt.Sprintf("%24s | %12s | %5s | %12s |\n", "", "game id", "", "game id"))
	return sb.String()
}

func createAgentStat() AgentStat {
	return AgentStat{
		Count: 0,
		Min: AgentStatRecord{
			Score:  0,
			Step:   0,
			GameId: 0,
		},
		AvgScore: 0,
		Max: AgentStatRecord{
			Score:  0,
			Step:   0,
			GameId: 0,
		},
	}
}

func NewPoolAgent(gameId int64) Agent {
	agent1 := NewClockwiseAgent(gameId)
	return &PoolAgent{
		currentIdx: 0,
		agent:      agent1,
		pool: []Agent{
			agent1,
			NewAnticlockwiseAgent(gameId),
		},
		statMap: make(map[string]AgentStat),
	}
}
