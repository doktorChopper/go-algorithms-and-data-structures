package matrix

import "fmt"

type Matrix struct {
    rows    int
    cols    int
    mat     [][]float64
}

func NewMatrix() *Matrix {
    return &Matrix {
        rows:   0,
        cols:   0,
        mat:    nil,
    }
}

func NewMatrixNM(n, m int) *Matrix {

    r := make([][]float64, n)
    for i := range r {
        r[i] = make([]float64, m)
    }

    return &Matrix {
        rows:   n,
        cols:   m,
        mat:    r,
    }
}

func NewMatrixFromSlice(s [][]float64) *Matrix {

    return &Matrix {
        rows:   len(s),
        cols:   len(s[0]),
        mat:    s,
    }
}

func (m *Matrix) Transposition() *Matrix {

    mT := make([][]float64, m.rows)
    for i := range mT {
        mT[i] = make([]float64, m.cols)
    }

    // TODO
    return nil
}

func (m *Matrix) Display() {

    for _, r := range m.mat {

        for _, c := range r {
            fmt.Printf("%.2f  ", c)
        }
        fmt.Println()
    }

    fmt.Println()
}
