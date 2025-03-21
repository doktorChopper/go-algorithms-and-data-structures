package graph

import "fmt"

type Graph struct {
    vertices map[int]*vertex
}

type vertex struct {
    v       int
    edges   map[int]*edge
}

type edge struct {
    w       int
    v       *vertex
}

func NewGraph() *Graph {
    return &Graph {
        vertices: map[int]*vertex{},
    }
} 

func (g *Graph) AddVertex(k, v int) {
    g.vertices[k] = &vertex {
        v:      v,
        edges:  map[int]*edge{},
    }
}

func (g *Graph) AddEdge(srcK, destK, w int) {
    
    if _, ok := g.vertices[srcK]; !ok {
        return
    }

    if _, ok := g.vertices[destK]; !ok {
        return
    }

    g.vertices[srcK].edges[destK] = &edge {
        w: w,
        v: g.vertices[destK],
    }
}

func (g *Graph) Display() {

    for k, i := range g.vertices {
        fmt.Printf("%d: ", k)
        for l := range i.edges {
            fmt.Printf("%d ", l)
        } 
        fmt.Println()
    }
}
