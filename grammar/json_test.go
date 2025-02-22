package grammar

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func PrintPoint(point *Point) {
	fmt.Println(strconv.Itoa(point.X) + ":" + strconv.Itoa(point.Y))
}

func TestJson(t *testing.T) {
	point := Point{X: 10, Y: 10}
	jsonStr, err := json.Marshal(point)
	if err != nil {
		err.Error()
	} else {
		fmt.Println(string(jsonStr))
	}

	var point1 Point
	_ = json.Unmarshal(jsonStr, &point1)
	fmt.Print(point1.X)
}
