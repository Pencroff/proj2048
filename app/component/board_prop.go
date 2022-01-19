package component

import (
	"github.com/pencroff/ai-agents/proj2048/agent"
	"github.com/pencroff/ai-agents/proj2048/common"
	"github.com/pencroff/ai-agents/proj2048/helper"
	"github.com/pencroff/ai-agents/proj2048/resources"
	"image"
	"math/rand"
)

type BoardProp struct {
	Name        string
	Description string
	Score       int
	Step        int
	NoMove      bool
	Speed       int
	State       common.BoardState
	Color       common.ColorPair
	Size        image.Point
	BoardRect   image.Rectangle
	Direction   common.Direction
	FieldProps  *FieldProps
	List        []*TileProp
	Agent       agent.Agent
	IsFinished  bool
	Mode        common.PlayMode
}

func (p *BoardProp) NewTile() (idx int, prop TileProp) {
	lst := p.SerializeList()
	idxLst := filterEmptyIndexes(lst)
	valueLst := []int{2, 4}
	if len(idxLst) == 0 {
		return -1, TileProp{}
	}
	idx = rand.Int() % len(idxLst)
	idx = idxLst[idx]

	valueIdx := rand.Int() % len(valueLst)
	value := valueLst[valueIdx]
	pos := p.ResolveScreenPositionByIndex(idx)

	prop = TileProp{
		Value:    value,
		Position: pos,
		IsMoving: false,
		Sprite:   resources.TileMap[value],
	}
	return
}

func (p *BoardProp) SerializeList() (res []int) {
	res = make([]int, p.Size.X*p.Size.Y)
	for idx, tile := range p.List {
		if tile != nil {
			res[idx] = tile.Value
		}
	}
	return
}

func (p *BoardProp) MoveUp() {
	//log.Printf("Up")
	sign := -1
	idxList := []int{0, 1, 2, 3}
	for _, x := range idxList {
		fullCell := 0
		emptyCell := 0
		for yIdx, y := range idxList {
			idx := p.ResolveIndexByFieldCoordinates(x, y)
			tile := p.List[idx]
			if tile != nil {
				fullCell += 1
				if yIdx >= fullCell {
					y += emptyCell * sign
					p.List[idx] = nil
					idx = p.ResolveIndexByFieldCoordinates(x, y)
					pos := p.ResolveScreenPositionByIndex(idx)

					tile.IsMoving = true
					tile.EndPosition = pos
					tile.Speed = image.Point{Y: p.Speed * sign}
					p.List[idx] = tile
				}
			} else {
				emptyCell += 1
			}

		}
	}
}

func (p *BoardProp) MoveDown() {
	//log.Printf("Down")
	idxList := []int{3, 2, 1, 0}
	for _, x := range idxList {
		fullCell := 0
		emptyCell := 0
		for yIdx, y := range idxList {
			idx := p.ResolveIndexByFieldCoordinates(x, y)
			tile := p.List[idx]
			if tile != nil {
				fullCell += 1
				if yIdx >= fullCell {
					y += emptyCell
					p.List[idx] = nil
					idx = p.ResolveIndexByFieldCoordinates(x, y)
					pos := p.ResolveScreenPositionByIndex(idx)

					tile.IsMoving = true
					tile.EndPosition = pos
					tile.Speed = image.Point{Y: p.Speed}
					p.List[idx] = tile
				}
			} else {
				emptyCell += 1
			}
		}
	}
}

func (p *BoardProp) MoveRight() {
	//log.Printf("Right")
	idxList := []int{3, 2, 1, 0}
	for _, y := range idxList {
		fullCell := 0
		emptyCell := 0
		for xIdx, x := range idxList {
			idx := p.ResolveIndexByFieldCoordinates(x, y)
			tile := p.List[idx]
			if tile != nil {
				fullCell += 1
				if xIdx >= fullCell {
					x += emptyCell
					p.List[idx] = nil
					idx = p.ResolveIndexByFieldCoordinates(x, y)
					pos := p.ResolveScreenPositionByIndex(idx)

					tile.IsMoving = true
					tile.EndPosition = pos
					tile.Speed = image.Point{X: p.Speed}
					p.List[idx] = tile
				}
			} else {
				emptyCell += 1
			}
		}
	}
}

func (p *BoardProp) MoveLeft() {
	//log.Printf("Left")
	idxList := []int{0, 1, 2, 3}
	for _, y := range idxList {
		fullCell := 0
		emptyCell := 0
		for xIdx, x := range idxList {
			idx := p.ResolveIndexByFieldCoordinates(x, y)
			tile := p.List[idx]
			if tile != nil {
				fullCell += 1
				if xIdx >= fullCell {
					x -= emptyCell
					p.List[idx] = nil
					idx = p.ResolveIndexByFieldCoordinates(x, y)
					pos := p.ResolveScreenPositionByIndex(idx)

					tile.IsMoving = true
					tile.EndPosition = pos
					tile.Speed = image.Point{X: -p.Speed}
					p.List[idx] = tile
				}
			} else {
				emptyCell += 1
			}
		}
	}

}

func (p BoardProp) CalculateUp() (score int) {
	valueList := p.SerializeList()
	idxList := []int{0, 1, 2, 3}
	for _, x := range idxList {
		for _, y := range idxList[:3] {
			idx1 := p.ResolveIndexByFieldCoordinates(x, y)
			idx2 := p.ResolveIndexByFieldCoordinates(x, y+1)
			val1 := valueList[idx1]
			val2 := valueList[idx2]
			if val1 == val2 && val1 > 0 {
				sum := valueList[idx1] + valueList[idx2]
				score += sum
				p.UpdateMergedTiles(idx1, idx2, sum, valueList)
			}
		}
	}
	return
}

func (p BoardProp) CalculateDown() (score int) {
	valueList := p.SerializeList()
	idxList := []int{3, 2, 1, 0}
	for _, x := range idxList {
		for _, y := range idxList[:3] {
			idx1 := p.ResolveIndexByFieldCoordinates(x, y)
			idx2 := p.ResolveIndexByFieldCoordinates(x, y-1)
			val1 := valueList[idx1]
			val2 := valueList[idx2]
			if val1 == val2 && val1 > 0 {
				sum := valueList[idx1] + valueList[idx2]
				score += sum
				p.UpdateMergedTiles(idx1, idx2, sum, valueList)
			}
		}
	}
	return
}

func (p BoardProp) CalculateRight() (score int) {
	valueList := p.SerializeList()
	idxList := []int{3, 2, 1, 0}
	for _, y := range idxList {
		for _, x := range idxList[:3] {
			idx1 := p.ResolveIndexByFieldCoordinates(x, y)
			idx2 := p.ResolveIndexByFieldCoordinates(x-1, y)
			val1 := valueList[idx1]
			val2 := valueList[idx2]
			if val1 == val2 && val1 > 0 {
				sum := valueList[idx1] + valueList[idx2]
				score += sum
				p.UpdateMergedTiles(idx1, idx2, sum, valueList)
			}
		}
	}
	return
}

func (p BoardProp) CalculateLeft() (score int) {
	valueList := p.SerializeList()
	idxList := []int{0, 1, 2, 3}
	for _, y := range idxList {
		for _, x := range idxList[:3] {
			idx1 := p.ResolveIndexByFieldCoordinates(x, y)
			idx2 := p.ResolveIndexByFieldCoordinates(x+1, y)
			val1 := valueList[idx1]
			val2 := valueList[idx2]
			if val1 == val2 && val1 > 0 {
				sum := valueList[idx1] + valueList[idx2]
				score += sum
				p.UpdateMergedTiles(idx1, idx2, sum, valueList)
			}
		}
	}
	return
}

func (p BoardProp) UpdateMergedTiles(idx1 int, idx2 int, sum int, lst []int) {
	tile := p.List[idx1]
	tile.Value = sum
	tile.Sprite = resources.TileMap[sum]
	lst[idx1] = sum
	tile = p.List[idx2]
	tile.Removed = true
	lst[idx2] = 0
	p.List[idx2] = nil
}

func (p *BoardProp) ResolveIndexByFieldCoordinates(x, y int) (idx int) {
	idx = p.Size.Y*y + x
	return
}

func (p *BoardProp) ResolveScreenPositionByIndex(idx int) (pos image.Point) {
	x := idx % p.Size.X
	y := idx / p.Size.Y
	xSpace := helper.RoundToInt(float64(resources.Config.FieldRect.Dx()) * resources.Config.FieldSpace)
	ySpace := helper.RoundToInt(float64(resources.Config.FieldRect.Dy()) * resources.Config.FieldSpace)
	tileW, tileH := resources.Config.Tile.X, resources.Config.Tile.Y
	pos.X = xSpace + x*(xSpace+tileW)
	pos.Y = ySpace + y*(ySpace+tileH) + resources.Config.HeaderRect.Dy()
	return
}

func (p BoardProp) CanMakeMove(lst []int) bool {
	idxList := []int{0, 1, 2, 3}
	var x, y int
	for _, v := range lst {
		if v == 0 {
			return true
		}
	}
	for _, p1 := range idxList {
		for _, p2 := range idxList[:3] {
			// v check
			x = p1
			y = p2
			idx1 := p.ResolveIndexByFieldCoordinates(x, y)
			idx2 := p.ResolveIndexByFieldCoordinates(x, y+1)
			val1 := lst[idx1]
			val2 := lst[idx2]
			if val1 == val2 && val1 > 0 {
				return true
			}
			// h check
			x = p2
			y = p1
			idx1 = p.ResolveIndexByFieldCoordinates(x, y)
			idx2 = p.ResolveIndexByFieldCoordinates(x+1, y)
			val1 = lst[idx1]
			val2 = lst[idx2]
			if val1 == val2 && val1 > 0 {
				return true
			}
		}
	}
	return false
}

func filterEmptyIndexes(values []int) (lst []int) {
	for idx, val := range values {
		if val == 0 {
			lst = append(lst, idx)
		}
	}
	return
}
