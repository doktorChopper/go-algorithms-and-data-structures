package main

const defaultStartArrCapacity = 16

type MinHeap struct {
    arr     []int
    size    int
}

func leftChild(i int) int {
    return i * 2 + 1
}

func rightChild(i int) int {
    return i * 2 + 2
}

func parent(i int) int {
    return (i - 1) / 2
}

func NewDefaultMinHeap() *MinHeap {
    return &MinHeap {
        arr:    make([]int, 0, defaultStartArrCapacity),
        size:   0,
    }
}

func NewMinHeap(n int) *MinHeap {
    return &MinHeap{
        arr:    make([]int, 0, n),
        size:   0,
    }
}

func (m *MinHeap) siftUp(i int) {

    for i > 0 && m.arr[parent(i)] > m.arr[i] {
        swp := m.arr[parent(i)]
        m.arr[parent(i)] = m.arr[i]
        m.arr[i] = swp

        i = parent(i)
    }
}

func (m *MinHeap) siftDown(i int) {

    for i * 2 + 1 < m.size {

        l := i * 2 + 1
        r := i * 2 + 2
        j := l
        
        if r < m.size && m.arr[r] < m.arr[l] {
            j = r
        }

        if m.arr[i] <= m.arr[j] {
            break
        }

        swp := m.arr[i]
        m.arr[i] = m.arr[j]
        m.arr[j] = swp

        i = j
    }
    
}

func BuildHeap(a []int) *MinHeap {

    r := &MinHeap {
        arr:    a,
        size:   len(a),
    }

    for i := len(r.arr) / 2; i >= 0; i-- {
        r.siftDown(i)
    }

    return r    
}

func (m *MinHeap) ExtractMin() int {
    r := m.arr[0] 
    m.arr[0] = m.arr[m.size - 1]
    m.size--
    m.siftDown(0)
    return r
}

func (m *MinHeap) Add(v int){
    m.size++
    m.arr = append(m.arr, v)
    m.siftUp(m.size - 1)
}
