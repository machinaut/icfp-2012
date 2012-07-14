// map.go - Map definition and functions to read in a map
// parsing should be pretty robust, even if errors aren't very verbose
package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

const (
	// Map symbols
	NONE   byte = 0 // invalid symbol
	ROBOT       = 'R'
	ROCK        = '*'
	CLOSED      = 'L'
	OPEN        = 'O'
	EARTH       = '.'
	WALL        = '#'
	LAMBDA      = '\\'
	EMPTY       = ' '
	BEARD       = 'W'
	RAZOR       = '!'

	// Robot Moves
	LEFT  = 'L'
	RIGHT = 'R'
	UP    = 'U'
	DOWN  = 'D'
	WAIT  = 'W'
	ABORT = 'A'

	// Completion States
	INCOMPLETE = 0
	WIN        = 1
	ABORTED    = 2
	LOSE       = 3
)

type Point struct {
	x, y int
}

type Map struct {
	m          [][]byte
	Robot      Point
	Lift       Point
	Lambdas    int
	Steps      int
	Complete   byte
	Water      int
	Flooding   int
	Waterproof int
	Trampoline map[byte]byte
	Tramp      map[byte]Point
	Target     map[byte]Point
	Beard      map[Point]int
	Growth     int
	Razors     int
}

// Read an arbitrarily long line
func ReadLine(b *bufio.Reader) ([]byte, error) {
	l := make([]byte, 0, 4096)
	for { // The joys of handling unconstrained input...
		line, isPrefix, err := b.ReadLine()

		if err != nil && err != io.EOF {
			return nil, err
		}

		l = append(l, line...)

		if err == io.EOF {
			return l, err
		}
		if !isPrefix {
			break
		}
	}
	return l, nil
}

// Read a whole map in, with optional metadata
func ReadMap(r io.Reader) (*Map, error) {
	m := new(Map)
	b := bufio.NewReaderSize(r, 4096)

	m.Trampoline = make(map[byte]byte, 9)
	m.Tramp = make(map[byte]Point, 9)
	m.Target = make(map[byte]Point, 9)
	m.Beard = make(map[Point]int, 9)

	m.m = make([][]byte, 0, 1024)
	// Parse map
	for {
		l, err := ReadLine(b)

		if err != nil && err != io.EOF {
			return nil, err
		}

		if len(l) == 0 {
			break
		}

		m.m = append(m.m, l)

		if err == io.EOF {
			break
		}
	}

	// Parse metadata
	for {
		l, err := ReadLine(b)
		if err != nil && err != io.EOF {
			return nil, err
		}

		words := strings.Fields(string(l))
		if len(words) > 1 {
			switch words[0] {
			case "Water":
				if m.Water, err = strconv.Atoi(words[1]); err != nil {
					return nil, err
				}
			case "Flooding":
				if m.Flooding, err = strconv.Atoi(words[1]); err != nil {
					return nil, err
				}
			case "Waterproof":
				if m.Waterproof, err = strconv.Atoi(words[1]); err != nil {
					return nil, err
				}
			case "Trampoline":
				if len(words) > 4 && len(words[1]) > 1 && len(words[3]) > 1 {
					from, to := words[1][0], words[3][0]
					m.Trampoline[from] = to
				}
			case "Growth":
				if m.Growth, err = strconv.Atoi(words[1]); err != nil {
					return nil, err
				}
			case "Razors":
				if m.Razors, err = strconv.Atoi(words[1]); err != nil {
					return nil, err
				}
			}
		}
		if err == io.EOF {
			break
		}
	}

	// Reverse rows (we read the file upside down)
	n := len(m.m)
	for i := 0; i < n/2; i++ {
		m.m[i], m.m[n-1-i] = m.m[n-1-i], m.m[i]
	}

	m.Init() // just init it here
	return m, nil
}
