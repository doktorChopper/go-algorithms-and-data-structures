package main

import "fmt"

func main() {

    a := []int{30, 29, 84, 8, 92, 78, 2, 1, 9, 44, 17, 19}

    m := BuildHeap(a)

    for _, v := range m.arr {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    // m := NewMinHeap(10)
    // m.Add(3)
    // m.Add(6)
    // m.Add(4)
    // m.Add(9)
    // m.Add(8)
    // m.Add(12)
    // m.Add(7)
    // m.Add(11)
    // m.Add(9)
    // // m.Add(5)
    //
    // fmt.Println(m.ExtractMin())

    fmt.Println("Hello, World!")
}
