import (
	"slices"
	"cmp"
)

func carFleet(target int, position, speed []int) int {
	// Combine position and speed, then sort
	type Car struct {
		position int
		speed    int
	}
	cars := make([]Car, len(position))
	for i, position := range position {
		cars[i] = Car{position, speed[i]}
	}
	slices.SortFunc(cars, func(a, b Car) int {
		return cmp.Compare(b.position, a.position)
	})

	// Iterate over cars
	stack := []float32{}
	for _, car := range cars {
		// Calculate ETA
		time := float32(target-car.position) / float32(car.speed)
		if (len(stack) > 0 && time > stack[len(stack)-1]) || len(stack) == 0 {
			// If time is greater than car in front, then add to stack
			stack = append(stack, time)
		}
	}
	return len(stack)
}



