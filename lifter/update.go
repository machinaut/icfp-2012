// update.go - methods on Map to initialize and update
package main

import "fmt"

// get a byte, checking ranges
func (m *Map) G(x, y int) byte { // uppercase = safe
	switch {
	case y < 0 || y > len(m.m):
		return NONE
	case x < 0 || x > len(m.m[y]):
		return NONE
	}
	return m.m[y][x]
}

// get a byte, NOT CHECKING RANGES
func (m *Map) g(x, y int) byte { // lowercase = unsafe
	return m.m[y][x]
}

// set a byte, NOT CHECKING RANGES
func (m *Map) s(x, y int, c byte) { // lowercase = unsafe
	m.m[y][x] = c
}

// Initialize Map: find robot & lift, count lambdas
func (m *Map) Init() {
	for y := 0; y < len(m.m); y++ {
		for x := 0; x < len(m.m[y]); x++ {
			switch m.G(x, y) {
			case OPEN, CLOSED:
				m.Lift = Point{x, y}
			case ROBOT:
				m.Robot = Point{x, y}
			case LAMBDA:
				m.Lambdas += 1
			}
		}
	}
}

// Robot Movement Step
func (m *Map) Move(command byte) {
	x, y := m.Robot.x, m.Robot.y
	m.s(x, y, EMPTY)
	switch command {
	case LEFT:
		x = m.Robot.x - 1
	case RIGHT:
		x = m.Robot.x + 1
	case UP:
		y = m.Robot.y + 1
	case DOWN:
		y = m.Robot.y - 1
	case ABORT:
		m.Complete = ABORTED
	}
	c := m.G(x, y)
	switch c {
	case EMPTY, EARTH:
		m.Robot.x, m.Robot.y = x, y
	case LAMBDA:
		m.Robot.x, m.Robot.y = x, y
		m.Lambdas -= 1
	case OPEN:
		m.Robot.x, m.Robot.y = x, y
		m.Complete = WIN
	case ROCK:
		switch {
		case command == LEFT && m.G(x-1, y) == EMPTY:
			m.Robot.x = x
			m.s(x-1, y, ROCK)
		case command == RIGHT && m.G(x+1, y) == EMPTY:
			m.Robot.x = x
			m.s(x+1, y, ROCK)
		}
	}
	m.s(m.Robot.x, m.Robot.y, ROBOT)
}

// Drop a Rock, see if it kills us
func (m *Map) Rock(x, y int) {
	m.s(x, y, ROCK)
	if x == m.Robot.x && y == m.Robot.y+1 {
		m.Complete = LOSE
	}

}

// Update Map step
func (m *Map) Update() {
	for y := 0; y < len(m.m); y++ { // oh god the nesting
		for x := 0; x < len(m.m[y]); x++ {
			if m.g(x, y) == ROCK {
				switch m.G(x, y-1) {
				case EMPTY:
					m.s(x, y, EMPTY)
					m.Rock(x, y-1)
				case ROCK:
					switch {
					case m.G(x+1, y) == EMPTY && m.G(x+1, y-1) == EMPTY:
						m.s(x, y, EMPTY)
						m.Rock(x+1, y-1)
					case m.G(x-1, y) == EMPTY && m.G(x-1, y-1) == EMPTY:
						m.s(x, y, EMPTY)
						m.Rock(x-1, y-1)
					}
				case LAMBDA:
					if m.G(x+1, y) == EMPTY && m.G(x+1, y-1) == EMPTY {
						m.s(x, y, EMPTY)
						m.Rock(x+1, y-1)
					}
				}
			}
		}
	}
	if m.Lambdas == 0 {
		m.s(m.Lift.x, m.Lift.y, OPEN)
	}
	fmt.Println("nope")
}

// Whole Step (Robot Move, Update, Exit Check) Returns true if game over
func (m *Map) Step(command byte) bool {
	m.Steps += 1
	m.Move(command)
	if m.Complete != INCOMPLETE {
		return true
	}
	m.Update()
	if m.Complete != INCOMPLETE {
		return true
	}
	return false
}
