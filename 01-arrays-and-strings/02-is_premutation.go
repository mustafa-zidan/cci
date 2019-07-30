package main

import(
	"fmt"
	"sort"
)

type SortedRunes []rune

func (s SortedRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s SortedRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s SortedRunes) Len() int {
    return len(s)
}

func Sort(s string) string {
    r := []rune(s)
    sort.Sort(SortedRunes(r))
    return string(r)
}


// sort and compare
// func isPremutation(src, target string) bool {
// 	if len(src) != len(target) {
// 		return false
// 	}
// 	return Sort(src) == Sort(target)
// }



// Counter
func isPremutation(src, target string) bool {
	if len(src) != len(target) {
		return false
	}
	m := make(map[rune]int, 0)
	for _, c := range src {
		if _,ok := m[c]; !ok {
			m[c] = 0
		}
		m[c]++
	}

	for _, c:= range target {
		if _, ok := m[c]; !ok {
			return false
		}
		m[c]--
		if v, _ := m[c]; v== 0 {
			delete(m, c)
		}
	}
	return len(m) == 0
}

func main() {
	fmt.Println("isPremutation ", "hello",  "hello my friend" , isPremutation("hello", "hello my friend") )
	fmt.Println("isPremutation ", "elloh",  "hello" , isPremutation("elloh", "hello") )
}