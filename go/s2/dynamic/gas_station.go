package dynamic

// problem:
// There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].
// You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.
// Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique

// insight:
// if fuel becomes negative at a station then start any start before that point will always fail,
// because, we cannot enter a station with negative gas, entry gas is either positive or zero and yet
// we are failing. Meaning, the cost is too high
func canCompleteCircuit(gas []int, cost []int) int {
	f, gf, startAt := 0, 0, 0

	for i, g := range gas {
		nc := g - cost[i]
		f += nc
		gf += nc

		if f < 0 {
			startAt = i + 1
			f = 0
		}
	}

	if gf < 0 {
		return -1
	}

	return startAt
}

// takes too long
func canCompleteCircuit_1(gas []int, cost []int) int {

	l := len(gas)
	mark := make([]int, l)
	backC := 0
	for i := 0; i < l; i++ {

		if gas[i] > 0 && gas[i]-cost[i] >= 0 && mark[i] == 0 {
			g := 0
			ends := true
			for j := i; j < l; j++ {
				nc := gas[j] - cost[j]
				if nc < 0 {
					mark[i] = 1
				}

				g += nc
				if g < 0 {
					ends = false
					break
				}
			}

			if ends {
				gg := g + backC
				if gg < 0 {
					return -1
				} else {
					return i
				}
			}
		}

		backC += gas[i] - cost[i]
	}

	return -1
}
