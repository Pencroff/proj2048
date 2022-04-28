package common

import "image/color"

type PlayMode string

const (
	Manual     PlayMode = "manual"
	AutoPlay            = "auto_play"
	Simulation          = "simulation"
)

type Direction string

const (
	NoDirection Direction = "no_direction"
	Up                    = "up"
	Right                 = "right"
	Down                  = "down"
	Left                  = "left"
)

type BoardState string

const (
	Nope          BoardState = "nope"
	StartGame                = "start_game"
	FillRand                 = "fill_rand"
	AgentAction              = "agent_action"
	CalcMove                 = "calc_move"
	Animate                  = "animate"
	EvaluateScore            = "evaluate_score"
	PostAnimate              = "post_animate"
	GameEnd                  = "game_end"
	RestartGame              = "restart_game"
)

type Pair struct {
	A, B interface{}
}

type ColorStrPair struct {
	Ink, Paper string
}

type ColorPair struct {
	Ink, Paper color.Color
}
