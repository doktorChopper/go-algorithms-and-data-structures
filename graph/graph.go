package graph

import (
	"fmt"
)

type Graph struct {
    vertices map[int]*vertex
}

type vertex struct {
    data        int
    edges       map[int]*edge
    color       byte
}

type edge struct {
    weight       int
    vertex       *vertex
}

func NewGraph() *Graph {
    return &Graph {
        vertices: map[int]*vertex{},
    }
} 

func (g *Graph) AddVertex(k, v int) {
    g.vertices[k] = &vertex {
        data:       v,
        edges:      map[int]*edge{},
        color:      'w',
    }
}

func (g *Graph) AddEdge(srcK, destK, w int) bool {
    
    if _, ok := g.vertices[srcK]; !ok {
        return false
    }

    if _, ok := g.vertices[destK]; !ok {
        return false
    }

    g.vertices[srcK].edges[destK] = &edge {
        weight: w,
        vertex: g.vertices[destK],
    }

    return true
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

func (g *Graph) Neighbors(k int) []int {

    r := []int{}
    for _, edge := range g.vertices[k].edges {
        r = append(r, edge.vertex.data)
    }

    return r
}

func (g *Graph) BFS(s int) {

    for _, v := range g.vertices {
        v.color = 'w'
    }

    g.vertices[s].color = 'g'

    q := newQueueBFS()
    q.enqueue(g.vertices[s])

    for !q.isEmpty() {

        u := q.dequeue().data
        fmt.Printf("%d ", u.data)

        for _, edge := range u.edges {
            if edge.vertex.color == 'w' {
                edge.vertex.color = 'g'
                q.enqueue(edge.vertex)
            }
        }

        u.color = 'b'
    }

    fmt.Println()
}

func (g *Graph) DFS(s int) {

    for _, v := range g.vertices {
        v.color = 'w'
    }

    g.vertices[s].color = 'g'

    g.dfsHelper(s)

    for k, v := range g.vertices {
        if v.color == 'w' {
            g.dfsHelper(k)
        }
    }

}

func (g *Graph) dfsHelper(s int) {

    st := newStackDFS()
    st.push(g.vertices[s])

    for !st.isEmpty() {

        u := st.pop().data
        fmt.Printf("%d ", u.data)

        u.color = 'b'

        for _, edge := range u.edges {
            if edge.vertex.color == 'w' {
                if edge.vertex == u {
                    continue
                }
                edge.vertex.color = 'g'
                st.push(edge.vertex)
            }
        }

    }
}

type node struct {
    data *vertex
    next *node
}

type stackDFS struct {
    top *node
}

func newStackDFS() *stackDFS {
    return &stackDFS {
        top: nil,
    }
}

func (s *stackDFS) isEmpty() bool {
    return s.top == nil
}

func (s *stackDFS) push(v *vertex) {
    n := &node {
        data: v,
        next: s.top,
    }

    s.top = n
}

func (s *stackDFS) pop() *node {

    if s.isEmpty() {
        return nil
    }

    ret := s.top

    s.top = s.top.next
    return ret
}

type queueBFS struct {
    head *node
    tail *node
}

func newQueueBFS() *queueBFS {
    return &queueBFS {
        head: nil,
        tail: nil,
    }
}

func (q *queueBFS) isEmpty() bool {
    return q.head == nil
}

func (q *queueBFS) enqueue(v *vertex) {

    n := &node {
        data: v,
        next: nil,
    }

    if !q.isEmpty() {
        q.tail.next = n
    } else {
        q.head = n
    }

    q.tail = n
}

func (q *queueBFS) dequeue() *node {

    if q.isEmpty() {
        return nil
    }

    ret := q.head
    q.head = q.head.next

    if q.head == nil {
        q.tail = nil
    }

    return ret
}
