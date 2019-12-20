package main

import (
	"fmt"
	"io/ioutil"
	"math"
	// "sort"
	"strconv"
	"strings"
)

type Navigation struct {
	Θ string
	R int
}

func (n Navigation) toXy() [2]int {
	xy := [2]int{0, 0}
	switch {
	case n.Θ == "U":
		xy[1] = 1
	case n.Θ == "R":
		xy[0] = 1
	case n.Θ == "D":
		xy[1] = -1
	case n.Θ == "L":
		xy[0] = -1
	}
	xy[0] *= n.R
	xy[1] *= n.R
	return xy
}

type Point struct {
	W1      bool
	W2      bool
	W1steps int
	W2steps int
}

func emptyPoint() Point {
	return Point{false, false, -1, -1}
}

func (p *Point) setW1(w1 bool) Point {
	p.W1 = w1
	return *p
}

func (p *Point) setW2(w2 bool) Point {
	p.W2 = w2
	return *p
}

func (p *Point) setStepsW1(steps int) Point {
	// only set it once
	if p.W1steps == 0 {
		p.W1steps = steps
	}
	return *p
}

func (p *Point) setStepsW2(steps int) Point {
	// only set it once
	if p.W2steps == 0 {
		p.W2steps = steps
	}
	return *p
}

type Grid struct {
	State []Point
	X, Y  int
}

func makeGrid(x int, y int) Grid {
	emptyGrid := make([]Point, x*y)
	// THIS BIT IS CRAZY slow
	// for i, _ := range emptyGrid {
	// 	emptyGrid[i] = emptyPoint()
	// }
	return Grid{emptyGrid, x, y}
}

func (g Grid) xyToi(x, y int) int {
	if x > g.X || x < 0 {
		panic("x out of bounds")
	}
	if y > g.Y || y < 0 {
		panic("y out of bounds")
	}
	return x + (y * g.Y)
}

func (g Grid) update(w, x, y, steps int) Grid {
	i := g.xyToi(x, y)
	p := g.State[i]
	if w == 1 {
		p = p.setW1(true)
		p = p.setStepsW1(steps)
	} else {
		p = p.setW2(true)
		p = p.setStepsW2(steps)
	}
	g.State[i] = p // TODO abstract me
	return g
}

func iToXy(i, gridX, gridY int) [2]int {
	xy := [2]int{0, 0}
	y := i / gridY
	x := i - (y * gridY)
	xy[0] = x
	xy[1] = y
	return xy
}

func abs(i int) int {
	return int(math.Abs(float64(i)))
}

func (g Grid) plotLine(w, x, y int, xyVec [2]int, stepsPtr *StepsCounter) {
	if xyVec[0] != 0 {
		// mod x
		i := 1
		xEnd := xyVec[0]
		for i <= abs(xyVec[0]) {
			stepsPtr.inc()
			if xEnd < 0 {
				g = g.update(w, x-i, y, stepsPtr.c)
			} else {
				g = g.update(w, x+i, y, stepsPtr.c)
			}
			i += 1
		}
	} else {
		// mod y
		i := 1
		yEnd := xyVec[1]
		for i <= abs(xyVec[1]) {
			stepsPtr.inc()
			if yEnd < 0 {
				g = g.update(w, x, y-i, stepsPtr.c)
			} else {
				g = g.update(w, x, y+i, stepsPtr.c)
			}
			i += 1
		}
	}
}

type StepsCounter struct {
	c int
}

func (sc *StepsCounter) inc() {
	sc.c += 1
}

func (g Grid) threadWireOntoGrid(origin [2]int, wire []Navigation, w int) Grid {
	x, y := origin[0], origin[1]
	stepsPtr := StepsCounter{0}
	for _, nav := range wire {
		xy := nav.toXy()
		g.plotLine(w, x, y, xy, &stepsPtr)
		x += xy[0]
		y += xy[1]
	}
	return g
}

type Input struct {
	Wire1 []Navigation
	Wire2 []Navigation
}

func parseNavs(line string) []Navigation {
	instructions := strings.Split(line, ",")
	var navs = make([]Navigation, len(instructions))
	for i, nav := range instructions {
		θ := string(nav[0])
		r, _ := strconv.Atoi(string(nav[1:]))
		navs[i] = Navigation{θ, r}
	}
	return navs
}

func readInputs(filename string) Input {
	if dat, err := ioutil.ReadFile(filename); err == nil {
		lines := strings.Split(string(dat), "\n")
		wire1 := parseNavs(lines[0])
		wire2 := parseNavs(lines[1])
		return Input{wire1, wire2}
	} else {
		empty := make([]Navigation, 0)
		return Input{empty, empty}
	}
}

func distance(xy1, xy2 [2]int) int {
	x := xy2[0] - xy1[0]
	y := xy2[1] - xy1[1]
	return abs(x) + abs(y)
}

type Result struct {
	Distance int
	Xy       [2]int
	StepsSum int
}

func main() {
	inputs := readInputs("/tmp/aoc3")
	// inputs := readInputs("/tmp/aoc3_test")
	w1 := inputs.Wire1
	w2 := inputs.Wire2
	g := makeGrid(25999, 25999)
	const midpoint int = 13000
	// g := makeGrid(3, 3)
	// const midpoint int = 1
	origin := [2]int{midpoint, midpoint}
	g = g.threadWireOntoGrid(origin, w1, 1)
	g = g.threadWireOntoGrid(origin, w2, 2)
	var intersections []Result
	for i, p := range g.State {
		if p.W1 && p.W2 { // intersection
			xy := iToXy(i, g.X, g.Y)
			stepsSum := p.W1steps + p.W2steps
			fmt.Println(xy, stepsSum)
			intersections = append(intersections, Result{distance(origin, xy), xy, stepsSum})
		}
	}
	fmt.Println(intersections)
}

/*** GRID ****
* 11x10
*
* ...........
* ...........
* ...........
* ....+----+.
* ....|....|.
* ....|....|.
* ....|....|.
* .........|.
* .o-------+.
* ...........
*
*************/
