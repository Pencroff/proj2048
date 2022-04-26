package stats

import "github.com/jmoiron/sqlx"

func NewStatRecorder(db *sqlx.DB) *StatRecorder {
	return &StatRecorder{db}
}

type StatRecorder struct {
	*sqlx.DB
}

func (r *StatRecorder) StartGame() {

}

func (r *StatRecorder) FinishGame() {

}

func (r *StatRecorder) AddStep() {

}
