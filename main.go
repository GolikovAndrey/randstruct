package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

type Circle struct {
	x, y, r float64
}

func crosses(left, right Circle) bool {

	return math.Sqrt(
		math.Pow(left.x-right.x, 2)+math.Pow(left.y-right.y, 2),
	) <= left.r+right.r
}

func check(item Circle, items []Circle) bool {
	for _, one := range items {
		if crosses(item, one) {
			return false
		}
	}
	return true
}

func main() {
	i := 1
	j := 2
	errors := 0
	maxRange := 150

	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Circles X coordinate")
	f.SetCellValue("Sheet1", "B1", "Circles Y coordinate")
	f.SetCellValue("Sheet1", "C1", "Circles radius value")
	f.SetCellValue("Sheet1", "E1", "Dots X coordinate")
	f.SetCellValue("Sheet1", "F1", "Dots Y coordinate")
	f.SetCellValue("Sheet1", "G1", "Max. range")
	f.SetCellValue("Sheet1", "H1", maxRange)

	circles := make([]Circle, 0, 20)
	circle := Circle{}
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i < maxRange {
		if errors == 100 {
			break
		}
		circle.r = 5 + 1.5*r1.Float64()
		circle.x = 2*circle.r + (float64(maxRange)-4)*r1.Float64()
		circle.y = 2*circle.r + (float64(maxRange)-4)*r1.Float64()

		if !check(circle, circles) {
			errors++
			continue
		}
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+1), circle.x)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+1), circle.y)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+1), circle.r)
		circles = append(circles, circle)
		angle := 45.0

		// Square figure
		// for k := 0; k < 4; k++ {
		// 	alpha := (angle * math.Pi) / 180
		// 	f.SetCellValue("Sheet1", "E"+strconv.Itoa(j), circle.x+circle.r*math.Cos(alpha))
		// 	f.SetCellValue("Sheet1", "F"+strconv.Itoa(j), circle.y+circle.r*math.Sin(alpha))
		// 	angle += 90
		// 	j++
		// }

		// Triangle figure
		// for k := 0; k < 3; k++ {
		// 	alpha := (angle * math.Pi) / 180
		// 	f.SetCellValue("Sheet1", "E"+strconv.Itoa(j), circle.x+circle.r*math.Cos(alpha))
		// 	f.SetCellValue("Sheet1", "F"+strconv.Itoa(j), circle.y+circle.r*math.Sin(alpha))
		// 	angle += 60
		// 	j++
		// }

		// Circle figure
		for k := 0; k < 18; k++ {
			alpha := (angle * math.Pi) / 180
			f.SetCellValue("Sheet1", "E"+strconv.Itoa(j), circle.x+circle.r*math.Cos(alpha))
			f.SetCellValue("Sheet1", "F"+strconv.Itoa(j), circle.y+circle.r*math.Sin(alpha))
			angle += 20
			j++
		}
		i++
	}
	fmt.Println(errors)

	if err := f.SaveAs("ура.xlsx"); err != nil {
		fmt.Println(err)
	}
}
