package day12

import (
	"fmt"
	"io"
	"strings"

	"github.com/comfortablynull/aoc2021go/internal/reader"
)

type path interface {
	add(n *Node) (path, bool)
}

type p1 map[string]struct{}

func (p p1) add(n *Node) (path, bool) {
	if p == nil {
		return p1(map[string]struct{}{n.Value: {}}), true
	}
	if _, ok := p[n.Value]; ok && n.Lower {
		return nil, false
	}
	pn := p1(map[string]struct{}{n.Value: {}})
	for k, v := range p {
		pn[k] = v
	}
	return pn, true
}

type p2 struct {
	visited map[string]struct{}
	lv, lc  int
}

func (p p2) add(n *Node) (path, bool) {
	_, ok := p.visited[n.Value]
	if (ok && n.Lower && (p.lc+1)-p.lv > 1) || (ok && n.Value == "end") {
		return nil, false
	}
	pn := p2{visited: map[string]struct{}{n.Value: {}}, lv: p.lv, lc: p.lc}
	if n.Lower {
		if !ok {
			pn.lv += 1
		}
		pn.lc += 1
	}
	for k, v := range p.visited {
		pn.visited[k] = v
	}
	return pn, true
}

type Node struct {
	Value       string
	Connections []*Node
	Lower       bool
}

func (n *Node) Path(p path) int {
	pn, ok := p.add(n)
	if !ok {
		return 0
	}
	if n.Value == "start" {
		return 1
	}
	c := 0
	for _, n := range n.Connections {
		c += n.Path(pn)
	}
	return c
}

type Connection struct {
	From, To string
}

func connectionDec(s string) (Connection, error) {
	split := strings.Split(s, "-")
	if len(split) != 2 {
		return Connection{}, fmt.Errorf("bad line: %s", s)
	}
	return Connection{From: split[0], To: split[1]}, nil
}

func Run(r io.ReadSeeker) (int, int, error) {
	var first, second int
	scanner := reader.NewScanner[Connection](r, reader.NewDecoder(reader.NewBasicParser(connectionDec)))
	connections := map[string]*Node{
		"start": {Value: "start", Lower: true},
		"end":   {Value: "end", Lower: true},
	}
	for scanner.Scan() {
		con, err := scanner.Result()
		if err != nil {
			return first, second, err
		}
		if _, ok := connections[con.From]; !ok {
			connections[con.From] = &Node{Value: con.From, Lower: 'a' <= con.From[0] && con.From[0] <= 'z'}
		}
		if _, ok := connections[con.To]; !ok {
			connections[con.To] = &Node{Value: con.To, Lower: 'a' <= con.To[0] && con.To[0] <= 'z'}

		}
		connections[con.From].Connections = append(connections[con.From].Connections, connections[con.To])
		connections[con.To].Connections = append(connections[con.To].Connections, connections[con.From])
	}
	first = connections["end"].Path(p1(nil))
	second = connections["end"].Path(p2{})
	return first, second, nil
}
