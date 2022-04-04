package main

type Input struct {
	Shortcuts []int
}

type Node struct {
	Id          int
	ChildrenIds map[int]struct{}
}

func CalculatePathEnergy(in *Input) []int {
	nodes := make([]*Node, len(in.Shortcuts))
	for i := 0; i < len(in.Shortcuts); i++ {
		nodes[i] = &Node{
			Id:          i + 1,
			ChildrenIds: make(map[int]struct{}),
		}
		if i != 0 {
			nodes[i].ChildrenIds[nodes[i-1].Id] = struct{}{}
		}
	}
	for i := 0; i < len(in.Shortcuts); i++ {
		shortcutNodeId := in.Shortcuts[i]
		currNode := nodes[shortcutNodeId-1]
		childNode := nodes[i]
		if currNode.Id == childNode.Id {
			continue
		}
		currNode.ChildrenIds[childNode.Id] = struct{}{}
	}
	results := make(map[int]int)
	for i := 0; i < len(nodes); i++ {
		res := findPathLenToZero(nodes[i], nodes, results, 0)
		results[i+1] = res
	}
	resultsSlice := make([]int, len(results))
	for id, pathLen := range results {
		resultsSlice[id-1] = pathLen
	}
	return resultsSlice
}

func findPathLenToZero(n *Node, nodes []*Node, results map[int]int, acc int) int {
	if len(n.ChildrenIds) == 0 {
		return acc
	}
	if r, ok := results[n.Id]; ok {
		return acc + r
	}
	acc += 1
	min := -1
	for c := range n.ChildrenIds {
		if c == n.Id {
			continue
		}
		res := findPathLenToZero(nodes[c-1], nodes, results, acc)
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
