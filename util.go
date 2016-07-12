package eval

import "sort"

func _range(o, n int) []int {
    v := make([]int, 0)
    for i := o; i < n; i++ {
        v = append(v, i)
    }
    return v
}

func zip(a, b []int) []int {
    n := min(len(a), len(b))
    c := make([][2]int, n)
    for i := 0; i < n; i++ {
        c[i] = [2]int{a[i], b[i]}
    }
    return c
}

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func difference(a, b []int) []int {
    m := make(map[int]bool)
    for _, v := range a {
        m[v] = true
    }
    for _, v := range b {
        delete(m, v)
    }
    vals := make([]int, 0)
    for _, v := range m {
        vals = append(vals, v)
    }
    sort.Ints(vals)
    return vals
}

