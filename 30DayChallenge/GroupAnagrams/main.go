package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Printf("res:%v", res)

}

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for i := range strs {
		temp := strings.Split(strs[i], "")
		sort.Strings(temp)
		res := strings.Join(temp, "")
		m_temp := m[res]
		m_temp = append(m_temp, strs[i])
		m[res] = m_temp
	}

	for k, _ := range m {
		key := m[k]
		res = append(res, key)
	}

	return res
}
