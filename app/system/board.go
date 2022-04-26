package system

import (
	"fmt"
	log "github.com/go-pkgz/lgr"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/pencroff/proj2048/app/agent"
	"github.com/pencroff/proj2048/app/common"
	"github.com/pencroff/proj2048/app/component"
	"github.com/pencroff/proj2048/app/control"
	"github.com/pencroff/proj2048/app/entity"
	"github.com/pencroff/proj2048/app/helper"
	"github.com/pencroff/proj2048/app/resources"
	"github.com/sedyh/mizu/pkg/engine"
	"image"
	"strconv"
	"time"
)

type Board struct {
	*component.BoardProp
}

var state = common.Nope
var spacePressed = 0

func (b *Board) Update(w engine.World) {
	if state != b.State {
		state = b.State
		//log.Printf("[INFO] State: %s", state)
	}
	if spacePressed > 0 {
		spacePressed -= 1
	}

	if b.State == common.AgentAction && ebiten.IsKeyPressed(ebiten.KeySpace) && spacePressed == 0 {
		spacePressed = 8
		if b.Mode == common.AutoPlay {
			b.Mode = common.Manual
			b.Agent = resources.HumanAgentInstance
			b.Speed = 20
		} else if b.Mode == common.Manual {
			b.Mode = common.AutoPlay
			b.Agent = resources.PoolAgentInstance
			b.Speed = 500
		}
		b.Description = b.Agent.GetName()
		b.State = common.StartGame
	}

	if b.State == common.StartGame {
		b.Description = fmt.Sprintf("%s (%d)", b.Agent.GetName(), b.Agent.GetGameId())
		b.Flk.Seed(b.Agent.GetGameSeed())
		for idx, el := range b.List {
			if el != nil {
				el.Removed = true
				b.List[idx] = nil
			}
		}
		b.IsFinished = false
		b.Score = 0
		b.Step = 0
		b.State = common.FillRand
	}
	if b.State == common.FillRand {
		var tileCnt int
		if b.Step == 0 {
			tileCnt = 2
		} else {
			tileCnt = 1
		}

		for n := 0; n < tileCnt; n++ {
			b.addTile(w)
		}
		b.State = common.AgentAction
		if b.Step == 0 {
			err := b.Agent.LogStep(b.Step, b.Score, b.NoMove, b.SerializeList(), common.NoDirection)
			if err != nil {
				log.Printf("[ERROR] Can't log step: %v", err)
			}
		}
		return
	}
	if b.State == common.AgentAction {
		direction := b.Agent.MakeMove(b.Step, b.Score, b.NoMove, b.SerializeList())
		if direction != common.NoDirection {
			b.Direction = direction
			b.State = common.CalcMove
		}
		return
	}
	if b.State == common.CalcMove {
		b.calculateMove()
		if b.isAnimationFinished() {
			b.State = common.EvaluateScore
			b.NoMove = true
		} else {
			b.State = common.Animate
			b.NoMove = false
		}
		return
	}
	if b.State == common.Animate {
		if b.isAnimationFinished() {
			b.State = common.EvaluateScore
		}
		return
	}
	if b.State == common.EvaluateScore {
		score := b.calculateScore()
		lst := b.SerializeList()
		if b.CanMakeMove(lst) {
			if score == 0 && b.NoMove {
				b.State = common.AgentAction
			} else {
				b.calculateMove()
				b.Score += score
				b.Step += 1
				if b.isAnimationFinished() {
					b.State = common.FillRand
				} else {
					b.State = common.PostAnimate
				}
			}
		} else {
			b.State = common.GameEnd
		}
		return
	}
	if b.State == common.PostAnimate {
		if b.isAnimationFinished() {
			b.State = common.FillRand
		}
		return
	}
	if b.State == common.GameEnd {
		b.IsFinished = true
		if !b.Agent.IsManual() {
			time.Sleep(1 * time.Second)
			b.State = common.RestartGame
		}
		return
	}
	if b.State == common.RestartGame {
		b.Agent.GameFinished(b.Step, b.Score, b.NoMove, b.SerializeList(), b.Direction)
		b.State = common.StartGame
	}
}

func (b *Board) Draw(_ engine.World, screen *ebiten.Image) {
	padding := 20
	zeroPoint := image.Point{X: padding, Y: padding}

	img := ebiten.NewImage(b.BoardRect.Dx(), b.BoardRect.Dy())
	img.Fill(b.Color.Paper)

	headerFont := resources.HeaderFF
	titleFont := resources.TitleFF
	textFont := resources.TextFF

	_, headerZP := helper.BoundStringImage(resources.HeaderFF, b.Name)
	headerZP = headerZP.Add(zeroPoint)
	text.Draw(img, b.Name, headerFont, headerZP.X, headerZP.Y, b.Color.Ink)

	_, descriptionZP := helper.BoundStringImage(resources.TextFF, b.Description)
	descriptionZP = descriptionZP.Add(zeroPoint)
	descriptionZP = descriptionZP.Add(image.Point{Y: int(resources.HeaderFontSize * 1.5)})
	text.Draw(img, b.Description, textFont, descriptionZP.X, descriptionZP.Y, b.Color.Ink)

	badgeSize := image.Pt(75, 55)

	badge := control.Badge{
		Control: control.Control{
			BorderRadius: 4,
			Padding:      image.Pt(8, 8),
			Size:         badgeSize,
			Ink:          b.Color.Paper,
			Paper:        b.Color.Ink,
		},
		AlignH:        control.Center,
		Label:         "Score",
		LabelFontFace: titleFont,
		Value:         strconv.Itoa(b.Score),
		ValueFontFace: textFont,
	}

	scoreImg := ebiten.NewImageFromImage(badge.Draw())

	op := &ebiten.DrawImageOptions{}
	offsetX := float64(b.BoardRect.Dx() - (padding + badgeSize.X))
	offsetY := float64(padding)
	op.GeoM.Translate(offsetX, offsetY)
	img.DrawImage(scoreImg, op)
	op.GeoM.Reset()

	badge.Label = "Step"
	badge.Value = strconv.Itoa(b.Step)
	stepImg := ebiten.NewImageFromImage(badge.Draw())
	offsetX = float64(b.BoardRect.Dx() - (padding+badgeSize.X)*2)
	offsetY = float64(padding)
	op.GeoM.Translate(offsetX, offsetY)
	img.DrawImage(stepImg, op)
	op.GeoM.Reset()

	screen.DrawImage(img, op)

	pool, ok := b.Agent.(*agent.PoolAgent)
	msg := fmt.Sprintf("Press SPACE to switch between manual and auto-play (%s)\n", b.Mode)
	if ok {
		msg += pool.GetStatStr()
		ebitenutil.DebugPrintAt(screen, msg, 10, 607)
	} else {
		ebitenutil.DebugPrintAt(screen, msg, 10, 607)
	}

}

func (b *Board) addTile(w engine.World) {
	idx, tileProp := b.NewTile()
	if idx == -1 {
		b.State = common.EvaluateScore
		return
	}
	ptr := &tileProp
	tile := &entity.Tile{
		TilePropWrap: component.TilePropWrap{Ptr: ptr},
	}
	if b.List[idx] == nil {
		b.List[idx] = ptr
		w.AddEntities(tile)
	} else {
		log.Printf("[WARN] Index %v already occupied", idx)
	}
}

func (b *Board) calculateMove() {
	if b.Direction == common.Up {
		b.MoveUp()
	}
	if b.Direction == common.Right {
		b.MoveRight()
	}
	if b.Direction == common.Down {
		b.MoveDown()
	}
	if b.Direction == common.Left {
		b.MoveLeft()
	}
}

func (b *Board) calculateScore() (score int) {
	if b.Direction == common.Up {
		score = b.CalculateUp()
	}
	if b.Direction == common.Down {
		score = b.CalculateDown()
	}
	if b.Direction == common.Right {
		score = b.CalculateRight()
	}
	if b.Direction == common.Left {
		score = b.CalculateLeft()
	}
	return
}

func (b *Board) isAnimationFinished() bool {
	for _, tile := range b.List {
		if tile != nil && tile.IsMoving {
			return false
		}
	}
	return true
}
