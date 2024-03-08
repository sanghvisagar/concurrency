package main

import (
	"fmt"
	"sort"
)

type FruitFrequency struct {
	Name      string
	Frequency int
}

func main() {
	inputStr := []string{"Apple", "Mango", "Guava", "Apple", "Guava", "Green", "Green"}

	frquencyMap := make(map[string]int)
	firstOccurenceMap := make(map[string]int)

	for idx, fruit := range inputStr {
		if _, ok := frquencyMap[fruit]; ok {
			frquencyMap[fruit] += 1
		} else {
			frquencyMap[fruit] = 1
			firstOccurenceMap[fruit] = idx
		}
	}

	fruencyObjectList := []FruitFrequency{}
	for key, val := range frquencyMap {
		fruencyObjectList = append(fruencyObjectList, FruitFrequency{
			Name:      key,
			Frequency: val,
		})
	}

	// fruencyObjectList.sort
	sort.Slice(fruencyObjectList, func(i int, j int) bool {
		if fruencyObjectList[i].Frequency == fruencyObjectList[j].Frequency {
			return getAlphabeticalOrder(fruencyObjectList[i].Name, fruencyObjectListp[j].Name)
			// return firstOccurenceMap[fruencyObjectList[i].Name] < firstOccurenceMap[fruencyObjectList[j].Name]
		}
		return fruencyObjectList[i].Frequency > fruencyObjectList[j].Frequency
	})

	for _, obj := range fruencyObjectList {
		fmt.Println(obj.Name, " ", obj.Frequency)
	}

}

func getAlphabeticalOrder(str1 string, str2 string) bool {
	n := len(str1)
	m := len(str2)

	for i:=0; i<min(n,m); i++ {
		if str1[i] == str2[i] {
			continue
		}

		if str1[i] < str2[i] {
			return true
		} else {
			return false
		}
	}

	if n < m {
		return true
	}

	return false
}


// Green
// Greeneeee