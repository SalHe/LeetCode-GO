package main

import (
	"fmt"
	"time"

	"github.com/buger/goterm"
)

// 没来得及提交

func main() {
	// 16*16
	// D: (0,0)
	// W: (15,15)
	playerA := &player{0, 0, getInstructions(), RIGHT, RIGHT, false}
	playerB := &player{boardSize - 1, boardSize - 1, getInstructions(), LEFT, RIGHT, false}
	steps, result := solve(playerA, playerB)
	println(steps)
	println(result)
}

const (
	boardSize = 16
	maxSteps  = 256
)

func solve(a *player, b *player) (int, string) {
	taken := make([][]*player, boardSize)
	for i := 0; i < boardSize; i++ {
		taken[i] = make([]*player, boardSize)
	}

	var bullets []*bullet
	maxS := len(a.instructions)
	if maxS < len(b.instructions) {
		maxS = len(b.instructions)
	}
	for i := 0; i < maxS; i++ {
		for _, bull := range bullets {
			bull.shot(a)
			bull.shot(b)
		}

		show(taken,
			map[*player]int{
				a: goterm.RED,
				b: goterm.GREEN,
			})

		bullets = []*bullet{}
		if a.dead && b.dead || a.x == b.x && a.y == b.y {
			return i, "P"
		}
		if a.dead {
			return i, "W"
		}
		if b.dead {
			return i, "D"
		}

		control(a, b, i, &bullets, taken)
		control(b, a, i, &bullets, taken)
	}

	aTaken := 0
	bTaken := 0
	for _, players := range taken {
		for _, p := range players {
			if p == a {
				aTaken++
			} else {
				bTaken++
			}
		}
	}
	if aTaken == bTaken {
		return maxS, "P"
	}
	if aTaken > bTaken {
		return maxS, "D"
	}
	return maxS, "W"
}

func control(a, b *player, i int, bullets *[]*bullet, taken [][]*player) {
	// var isMove bool
	var dx, dy int
	var x, y int

	a.shoot = a.instructions[i]
	if a.shoot != SHOOT {
		a.dir = a.shoot
	}
	bullet_ := a.shot()
	if bullet_ != nil {
		*bullets = append(*bullets, bullet_)
	} else {
		dx, dy = a.instructions[i].move()
		x = a.x + dx
		y = a.y + dy
		if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
			// switch a.dir {
			// case UP:
			// 	a.dir = DOWN
			// case DOWN:
			// 	a.dir = UP
			// case LEFT:
			// 	a.dir = LEFT
			// case RIGHT:
			// 	a.dir = RIGHT
			// }
		} else {
			if taken[y][x] != b {
				a.x = x
				a.y = y
				taken[y][x] = a
			}
		}
	}
}

func getInstructions() (instructions []instruction) {
	var s string
	fmt.Scanf("%s\n", &s)
	instructions = make([]instruction, len(s))
	for i, instr := range s {
		instructions[i] = instruction(instr)
	}
	return
}

type player struct {
	x, y         int
	instructions []instruction
	dir          instruction
	shoot        instruction
	dead         bool
}

func (p *player) shot() *bullet {
	return p.shoot.shoot(p)
}

type bullet struct {
	sender *player
	x, y   int
	dir    instruction
}

func (b *bullet) canShot(p *player) bool {
	if p == b.sender {
		return false
	}
	return b.dir == UP && p.x == b.x && p.y <= b.y ||
		b.dir == DOWN && p.x == b.x && p.y >= b.y ||
		b.dir == LEFT && p.y == b.y && p.x <= b.x ||
		b.dir == RIGHT && p.y == b.y && p.x >= b.x
}

func (b *bullet) shot(p *player) {
	p.dead = b.canShot(p)
}

type instruction byte

const (
	UP    = instruction('U')
	DOWN  = instruction('D')
	LEFT  = instruction('L')
	RIGHT = instruction('R')
	SHOOT = instruction('F')
)

var moves = map[instruction][]int{
	RIGHT: {1, 0},
	DOWN:  {0, 1},
	LEFT:  {-1, 0},
	UP:    {0, -1},
}

func (i instruction) shoot(sender *player) *bullet {
	if i == SHOOT {
		return &bullet{
			sender: sender,
			x:      sender.x,
			y:      sender.y,
			dir:    sender.dir,
		}
	}
	return nil
}

func (i instruction) move() (dx, dy int) {
	if i == SHOOT {
		return 0, 0
	}
	return moves[i][0], moves[i][1]
}

// region DEBUG

var shows = map[instruction]string{
	LEFT:  "←",
	RIGHT: "→",
	UP:    "↑",
	DOWN:  "↓",
	SHOOT: "*",
}

func show(taken [][]*player, colors map[*player]int) {
	goterm.Clear()
	for y, players := range taken {
		for x, p := range players {
			if p != nil {
				goterm.MoveCursor(x+1, y+1)
				s := "-"
				if p.x == x && p.y == y {
					if p.dead {
						s = "×"
					} else {
						s = shows[p.dir]
					}
					if p.shoot == SHOOT {
						s = goterm.Background(s, goterm.CYAN)
					}
				}
				s = goterm.Color(s, colors[p])
				goterm.Print(s)

			}
		}
	}
	goterm.Println()
	goterm.Flush()
	time.Sleep(time.Second / 5)
}

// endregion
