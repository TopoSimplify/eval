package main

import (
    "fmt"
    . "github.com/intdxdt/geom"
    . "../"
)

func main() {
    owkt := "LINESTRING ( 142 408, 146 503, 197 548, 279 538, 333 511, 345 468, 320 411, 355 379, 419 376, 473 320, 532 330, 544 376, 527 431, 573 436, 558 489 )"
    swkt := "LINESTRING ( 142 408, 146 503, 279 538, 345 468, 355 379, 473 320, 544 376, 558 489 )"

    oline := ReadGeometry(owkt).AsLinear()[0].Coordinates()
    sline := ReadGeometry(swkt).AsLinear()[0].Coordinates()
    gen := []int {0, 1, 3, 5, 7, 9, 11, 14}

    pcangle := PCAngle(oline, sline)
    pccs := PCCS(oline, sline)
    tlvd := TLVD(oline, gen)
    tapd := TAPD(oline, gen)

    fmt.Println(pcangle)
    fmt.Println(pccs)
    fmt.Println(tlvd)
    fmt.Println(tapd)
}
