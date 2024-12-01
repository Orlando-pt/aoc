package utils

import "strconv"

func StrToInt(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}

func IntAbs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}
