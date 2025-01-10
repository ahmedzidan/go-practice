package main

import (
	"sort"
	"testing"
)

func main() {
	//trace.Start(os.Stdout)
	//defer trace.Stop()
	// Open a file to store the CPU profile
	//profileFile, err := os.Create("p.out")
	//if err != nil {
	//	fmt.Println("Error creating profile file:", err)
	//	return
	//}
	//defer profileFile.Close()
	//
	//err = pprof.StartCPUProfile(profileFile)
	//if err != nil {
	//	return
	//}
	//
	//defer pprof.StopCPUProfile()
	//fmt.Println(BenchmarkGroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
func BenchmarkGroupAnagrams(b *testing.B) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	for i := 0; i < b.N; i++ {
		groupAnagrams(input)
	}
}

func groupAnagrams(strs []string) [][]string {

	groups := make(map[string][]string)

	for _, str := range strs {
		groups[sortString(str)] = append(groups[sortString(str)], str)
	}
	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
