package sort

func BubbleSort(a []int) {
    for i := 0; i < len(a); i++ {
        for j := i; j < len(a); j++ {
            if a[i] > a[j] {
                tmp := a[i]
                a[i] = a[j]
                a[j] = tmp
            }
        }
    } 
}

func InsertionSort(a []int) {
    for i := 1; i < len(a); i++ {
        v := a[i]
        j := i
        for  j > 0 && a[j - 1] > v {
            a[j] = a[j - 1]
            j--
        }
        a[j] = v
    }
}

func SelectionSort(a []int) {
    for i := 0; i < len(a) - 1; i++ {
        minElem := i
        for j := i + 1; j < len(a); j++ {
            if a[j] < a[minElem] {
                minElem = j
            }
        }
        tmp := a[minElem] 
        a[minElem] = a[i]
        a[i] = tmp
     }
}
