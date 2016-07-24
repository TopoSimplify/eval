package eval

import "sort"

func _range(o, n int) []int {
    v := make([]int, 0)
    for i := o; i < n; i++ {
        v = append(v, i)
    }
    return v
}

func zip(a, b []int) [][2]int {
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
    m := make(map[int]struct{})
    for _, v := range a {
        m[v] = struct{}{}
    }
    for _, v := range b {
        delete(m, v)
    }
    vals := make([]int, 0)
    for k := range m {
        vals = append(vals, k)
    }
    sort.Ints(vals)
    return vals
}


//count change in deflection angles
func num_dflns(angles []float64) int {
    var k, sign, cursign int
    for i := 0; i < len(angles); i++ {
        a := angles[i]
        if i == 0 {
            k += 1
            sign = _sign(a)
        } else {
            cursign = _sign(a)//current sign
            if cursign != sign {
                //if not equal to prev sign -> new curve
                k += 1
                sign = cursign
            }
        }
    }
    return k
}

//sign ruturns  1  if f >= 0  else -1
func _sign(f float64) int {
    if f >= 0 {
        return 1
    }
    return -1
}

