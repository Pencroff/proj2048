package agent

import (
	"encoding/binary"
	"github.com/OneOfOne/xxhash"
	log "github.com/go-pkgz/lgr"
	"github.com/pencroff/ai-agents/proj2048/common"
	"hash"
)

type Agent interface {
	GetId() string
	GetName() string
	IsManual() bool
	GetGameId() int64
	GetGameSeed() int64

	MakeMove(step int, score int, noMove bool, valueList []int) common.Direction
	LogStep(step int, score int, noMove bool, valueList []int, d common.Direction) error

	GameFinished(step int, score int, noMove bool, valueList []int, d common.Direction)
}

type GenericAgent struct {
	id            string
	name          string
	isManual      bool
	noMoveCounter int
	gameId        int64
	hasher        hash.Hash64
}

func NewGenericAgent(id string, name string, isManual bool, startGameId int64) GenericAgent {
	return GenericAgent{
		id:            id,
		name:          name,
		isManual:      isManual,
		noMoveCounter: 0,
		gameId:        startGameId,
		hasher:        xxhash.New64(),
	}
}

func (a *GenericAgent) GetId() string {
	return a.id
}

func (a *GenericAgent) GetName() string {
	return a.name
}

func (a *GenericAgent) IsManual() bool {
	return a.isManual
}

func (a *GenericAgent) GetGameId() int64 {
	return a.gameId
}

func (a *GenericAgent) GetGameSeed() int64 {
	arr := make([]byte, 8)
	binary.LittleEndian.PutUint64(arr, uint64(a.gameId))
	a.hasher.Reset()
	_, err := a.hasher.Write(arr)
	if err != nil {
		log.Printf("[ERROR] Cant write to hasher %v", err)
		return 0
	}
	return int64(a.hasher.Sum64())
}

func (a *GenericAgent) LogStep(step int, score int, noMove bool, lst []int, d common.Direction) error {
	log.Printf("=============")
	log.Printf("%s - %d", a.GetName(), a.gameId)
	log.Printf("%d %d %t", step, score, noMove)
	log.Printf("%v", lst[0:4])
	log.Printf("%v", lst[4:8])
	log.Printf("%v", lst[8:12])
	log.Printf("%v", lst[12:])
	log.Printf("Move: %s", d)
	log.Printf("=============")
	return nil
}

func (a *GenericAgent) GameFinished(step int, score int, noMove bool, lst []int, d common.Direction) {
	a.gameId += 1
	err := a.LogStep(step, score, noMove, lst, d)
	if err != nil {
		log.Printf("[ERROR] GameFinished - Can't log step: %v", err)
	}
}
