// map.go - Map definition and functions to read in a map
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

	// Robot Moves
	LEFT  = 'L'
	RIGHT = 'R'
	UP    = 'U'
	DOWN  = 'D'
	WAIT  = 'W'
	ABORT = 'A'

	// Completion States
	INCOMPLETE = iota
	WIN
	ABORTED
	LOSE
)

type Point struct {
	x, y int
}

type Map struct {
	m          [][]byte
	Water      int
	Flooding   int
	Waterproof int
	Robot      Point
	Lift       Point
	Lambdas    int
	Steps      int
	Complete   byte
}

// Read a 'Word #' line (e.g. "Waterproof 10")
func ReadWord(line []byte) (string, int) {
	fields := strings.Fields(string(line))
	if len(fields) < 2 {
		return "", -1
	}
	i, err := strconv.Atoi(fields[1])
	if err != nil {
		return "", -2
	}
	return fields[0], i
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

	m.m = make([][]byte, 0, 1024)
	for {
		l, err := ReadLine(b)

		if err != nil && err != io.EOF {
			return nil, err
		}

		if len(l) > 0 { // ignore empty lines
			m.m = append(m.m, l)
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

	// Parse metadata
	for i := 0; i < 3; i++ {
		word, val := ReadWord(m.m[0])
		switch word {
		case "Water":
			m.Water = val
		case "Flooding":
			m.Flooding = val
		case "Waterproof":
			m.Waterproof = val
		default:
			goto done
		}
		m.m = m.m[1:] // remove the line we just parsed
	}
done:
	m.Init() // just init it here
	return m, nil
}
