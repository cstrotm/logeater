package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

func main() {
	errormap := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		i := strings.Index(txt, "updating")
		if (i > 0) {
			msg := txt[i+13:]
			errormap[msg]++
		}
	}
	for _, res := range sortedKeys(errormap) {
		fmt.Println(errormap[res], ":", res)
	}
}
