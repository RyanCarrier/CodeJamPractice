package main

import (
	"fmt"
	"strconv"

	ioh "github.com/RyanCarrier/CodeJamPractice/iohelper"
)

//T test cases
var T int

type group struct {
	E int
	H int
	S []int
	D int
}

//var io = iohelper
func main() {
	T = ioh.GetInt()
	for test := 1; test <= T; test++ {
		fmt.Printf("Case #%d: ", test)
		solve()
	}
}

func solve() {
	E := ioh.GetInt()
	N := ioh.GetInt()
	S := ioh.GetIntsN(N)
	H := 0
	G := group{E, H, S, 0}
	fmt.Println(strconv.Itoa(loop(G)))
}

func loop(G group) int {
	//spew.Dump(G)
	if len(G.S) == 0 {
		return G.H
	}
	if G.dance() || G.delay() || G.recruit() || G.truce() {
		return loop(G)
	}
	return G.H
}

func (g *group) dance() bool {
	if g.E <= g.S[0] {
		return false
	}
	g.E = g.E - g.S[0]
	g.H++
	g.S = g.S[1:]
	g.D = 0
	return true
}

func (g *group) delay() bool {
	if g.D >= len(g.S) || len(g.S) == 1 {
		return false
	}
	g.S = append(g.S[1:], g.S[0])
	g.D++
	return true
}

func (g *group) truce() bool {
	g.S = g.S[1:]
	g.D = 0
	return true
}

func (g *group) recruit() bool {
	if g.H == 0 {
		return false
	}
	g.E += g.S[0]
	g.S = g.S[1:]
	g.H--
	g.D = 0
	return true
}
