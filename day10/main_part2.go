package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Position struct {
	y int
	x int
}

type Tile struct {
	content  string
	pos      Position
	distance int
	up       bool
	down     bool
	left     bool
	right    bool
	loop     bool
}

/**
 * Day 10: Pipe Maze - Part 2
 * url: https://adventofcode.com/2023/day/10
 */
func main() {
	grid, start := parse()
	ctx := Context{grid, start, nil, []Position{*start}}

	// dumb DFS
	for {
		current := ctx.queue[0]
		ctx.queue = ctx.queue[1:]

		propagate(current, grid, &ctx)
		if len(ctx.queue) == 0 {
			break
		}
	}

	// for each non loop tile, check what's on the left, and count line crossing
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !grid[i][j].loop && inLoop(grid, &grid[i][j]) {
				count++
			}
		}
	}

	fmt.Println("Part 2 =", count)
}

func inLoop(grid [][]Tile, tile *Tile) bool {
	countCrossinePath := 0
	countUp := 0
	countDown := 0
	for i := 0; i < tile.pos.x; i++ {
		if !grid[tile.pos.y][i].loop {
			continue
		}

		// we ignore horizontal tile (they dont impact the result)
		if grid[tile.pos.y][i].content == "-" {
			continue
		}

		// count |
		if grid[tile.pos.y][i].content == "|" {
			countCrossinePath++
			continue
		}

		// count zig zag
		if grid[tile.pos.y][i].up {
			if countDown == 1 {
				countCrossinePath++
			}
			countUp = (countUp + 1) % 2
		}

		if grid[tile.pos.y][i].down {
			if countUp == 1 {
				countCrossinePath++
			}
			countDown = (countDown + 1) % 2
		}
	}

	// if we have an odd number of crossing path, we are in the loop
	return countCrossinePath%2 == 1
}

type Context struct {
	grid  [][]Tile
	start *Position
	far   *Tile
	queue []Position
}

func (t *Tile) nextMoves(grid [][]Tile) []*Tile {
	var next []*Tile
	if t.up && grid[t.pos.y-1][t.pos.x].distance == 0 {
		next = append(next, &grid[t.pos.y-1][t.pos.x])
	}
	if t.down && grid[t.pos.y+1][t.pos.x].distance == 0 {
		next = append(next, &grid[t.pos.y+1][t.pos.x])
	}
	if t.left && grid[t.pos.y][t.pos.x-1].distance == 0 {
		next = append(next, &grid[t.pos.y][t.pos.x-1])
	}
	if t.right && grid[t.pos.y][t.pos.x+1].distance == 0 {
		next = append(next, &grid[t.pos.y][t.pos.x+1])
	}
	return next
}

func propagate(pos Position, grid [][]Tile, ctx *Context) {
	origin := grid[pos.y][pos.x]
	for _, tile := range grid[pos.y][pos.x].nextMoves(grid) {
		if tile.content == "S" {
			continue
		}
		tile.distance = origin.distance + 1
		tile.loop = true
		if ctx.far == nil || tile.distance > ctx.far.distance {
			ctx.far = tile
		}
		ctx.queue = append(ctx.queue, tile.pos)
	}
}

func parse() ([][]Tile, *Position) {
	var start *Position
	var tiles [][]Tile

	// transform input into 2D array of struct
	for i, s := range strings.Split(strings.TrimSpace(input), "\n") {
		var row []Tile
		for j, c := range s {
			switch c {
			case '.':
				row = append(row, Tile{string(c), Position{i, j}, 0, false, false, false, false, false})
			case '|':
				row = append(row, Tile{string(c), Position{i, j}, 0, true, true, false, false, false})
			case '-':
				row = append(row, Tile{string(c), Position{i, j}, 0, false, false, true, true, false})
			case 'L':
				row = append(row, Tile{string(c), Position{i, j}, 0, true, false, false, true, false})
			case 'J':
				row = append(row, Tile{string(c), Position{i, j}, 0, true, false, true, false, false})
			case '7':
				row = append(row, Tile{string(c), Position{i, j}, 0, false, true, true, false, false})
			case 'F':
				row = append(row, Tile{string(c), Position{i, j}, 0, false, true, false, true, false})
			case 'S':
				tile := Tile{string(c), Position{i, j}, 0, true, true, true, true, true}
				row = append(row, tile)
				start = &tile.pos
			}
		}
		tiles = append(tiles, row)
	}

	// optimization, clean dead end
	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles[i]); j++ {
			if tiles[i][j].up && (i <= 0 || !tiles[i-1][j].down) {
				tiles[i][j].up = false
			}

			if tiles[i][j].left && (j <= 0 || !tiles[i][j-1].right) {
				tiles[i][j].left = false
			}

			if tiles[i][j].down && (i >= len(tiles)-1 || !tiles[i+1][j].up) {
				tiles[i][j].down = false
			}

			if tiles[i][j].right && (j >= len(tiles[i])-1 || !tiles[i][j+1].left) {
				tiles[i][j].right = false
			}
		}
	}

	return tiles, start
}
