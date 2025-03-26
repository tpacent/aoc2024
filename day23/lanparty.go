package day23

import (
	"iter"
	"slices"
	"strings"
)

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
	}
}

type Graph struct {
	nodes map[string]*Node
}

func (g *Graph) Add(idA, idB string) {
	g.ensureNodes(idA, idB)
	g.nodes[idA].Links[idB] = struct{}{}
	g.nodes[idB].Links[idA] = struct{}{}
}

func (g *Graph) Iter() iter.Seq2[string, *Node] {
	return func(yield func(string, *Node) bool) {
		for id, node := range g.nodes {
			if ok := yield(id, node); !ok {
				return
			}
		}
	}
}

func (g *Graph) ensureNodes(ids ...string) {
	for _, id := range ids {
		if _, ok := g.nodes[id]; ok {
			continue
		}

		g.nodes[id] = &Node{
			Links: map[string]struct{}{},
		}
	}
}

type Node struct {
	Links map[string]struct{}
}

func Walk(graph *Graph, loopLen int) map[string]struct{} {
	out := make(map[string]struct{})

	for id := range graph.Iter() {
		path := make([]string, 0, loopLen)
		path = append(path, id)
		for _, path := range findLoops(graph, path, loopLen) {
			slices.Sort(path)
			out[strings.Join(path, ",")] = struct{}{}
		}
	}

	return out
}

func findLoops(graph *Graph, path []string, steps int) (out [][]string) {
	if steps == 0 {
		return nil
	}

	nodeID := path[len(path)-1]
	node := graph.nodes[nodeID]

	for nextID := range node.Links {
		if steps == 1 && nextID == path[0] {
			out = append(out, path)
			break
		}

		if slices.Contains(path, nextID) {
			continue
		}

		newPath := append(path, nextID)

		for _, path := range findLoops(graph, newPath, steps-1) {
			out = append(out, slices.Clone(path))
		}
	}

	return out
}

func GetPassword(graph *Graph) string {
	maxClique := FindClusters(graph)

	ids := make([]string, 0, len(maxClique))
	for k := range maxClique {
		ids = append(ids, k)
	}

	slices.Sort(ids)
	return strings.Join(ids, ",")
}

func CountLoops(graph *Graph, predicate func(string) bool) (count int) {
	for path := range Walk(graph, 3) {
		if slices.ContainsFunc(strings.Split(path, ","), predicate) {
			count++
		}
	}

	return
}

func FindClusters(graph *Graph) map[string]struct{} {
	vertices := map[string]struct{}{}
	for k := range graph.nodes {
		vertices[k] = struct{}{}
	}

	return BronKerbosch(
		map[string]struct{}{},
		vertices,
		map[string]struct{}{},
		graph,
	)
}

func BronKerbosch(r, p, x map[string]struct{}, graph *Graph) (out map[string]struct{}) {
	if len(p) == 0 && len(x) == 0 {
		return r
	}

	vertices := make([]string, 0, len(p))
	for k := range p {
		vertices = append(vertices, k)
	}

	for _, v := range vertices {
		neightborSet := graph.nodes[v].Links

		if set := BronKerbosch(
			union(r, map[string]struct{}{v: {}}),
			intersect(p, neightborSet),
			intersect(x, neightborSet),
			graph,
		); len(set) > len(out) {
			out = set
		}

		x[v] = struct{}{}
		delete(p, v)
	}

	return
}

func union(a, b map[string]struct{}) map[string]struct{} {
	out := map[string]struct{}{}

	for k := range a {
		out[k] = struct{}{}
	}

	for k := range b {
		out[k] = struct{}{}
	}

	return out
}

func intersect(a, b map[string]struct{}) map[string]struct{} {
	out := map[string]struct{}{}

	for k := range a {
		if _, ok := b[k]; ok {
			out[k] = struct{}{}
		}
	}

	return out
}
