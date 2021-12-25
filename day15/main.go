package main

import (
	"fmt"
	"strconv"

	"github.com/stscoundrel/advent-of-code-2021/day15/dijkstra"
	"github.com/stscoundrel/advent-of-code-2021/day15/grid"
	"github.com/stscoundrel/advent-of-code-2021/day15/reader"
)

func GetMostEfficientPath(fileName string) int {
	graph := dijkstra.NewGraph()
	griddata := reader.ReadGridFromFile(fileName)

	for x, row := range griddata {
		for y, _ := range row {
			point := grid.NewPoint(x, y, griddata[x][y])
			for _, pointrelation := range grid.GetPointNeighbours(point, griddata) {
				graph.AddEdge(pointrelation.Origin, pointrelation.Destination, pointrelation.Value)
			}
		}
	}

	// fmt.Println(graph)

	endPoint := strconv.Itoa(len(griddata[0])-1) + string("-") + strconv.Itoa(len(griddata)-1)
	fmt.Println("end is ", endPoint)
	risk, path := graph.GetPath("0-0", endPoint)

	fmt.Println(risk)
	fmt.Println(path)

	return risk
}
