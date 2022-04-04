package main

type Input struct {
	Shortcuts []int
}

func CalculatePathEnergy(in *Input) []int {
	nodes := make([][]int, len(in.Shortcuts))
	for i := 0; i < len(in.Shortcuts); i++ {
		nodes[i] = make([]int, 0)
		if i != 0 {
			nodes[i] = append(nodes[i], i-1)
		}
	}
	for i, s := range in.Shortcuts {
		connectedNodeId := s - 1
		if connectedNodeId == i+1 || connectedNodeId == i {
			continue
		}
		nodes[connectedNodeId] = append(nodes[connectedNodeId], i)
	}
	results := make(map[int]int)
	for i := 0; i < len(nodes); i++ {
		results[i] = findPathLenToZero(i, nodes, results, 0)
	}
	resultsSlice := make([]int, len(results))
	for id, pathLen := range results {
		resultsSlice[id] = pathLen
	}
	return resultsSlice
}

func findPathLenToZero(nId int, nodes [][]int, results map[int]int, acc int) int {
	currNode := nodes[nId]
	if len(currNode) == 0 {
		return acc
	}
	if r, ok := results[nId]; ok {
		return acc + r
	}
	acc += 1
	min := -1
	for _, c := range currNode {
		res := findPathLenToZero(c, nodes, results, acc)
		if min == -1 {
			min = res
			continue
		}
		if min > res {
			min = res
			continue
		}
	}
	return min
}
