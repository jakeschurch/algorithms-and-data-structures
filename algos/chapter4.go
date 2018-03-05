package algos

import "fmt"

func towerofHanoi() {
	moves := 0
	tower1 := newStack(&stackVal{3}, &stackVal{2}, &stackVal{1})
	tower2 := newStack()
	tower3 := newStack()
}

// func hanoiMove(val *stackVal, toStack *stack) {
// 	if val.Value <=
// }

type stack struct {
	count  int
	Values []*stackVal
}

// func newStack(values ...*stackVal) (stack *stack) {
// 	if len(values) != 0 {
// 		for _, v := range values {
// 			stack.push(v)
// 		}
// 	}
// 	return &stack
// }

func (stack *stack) initFirstTower() {
	stack.push(&stackVal{3})
	stack.push(&stackVal{2})
	stack.push(&stackVal{1})
}

type stackVal struct {
	Value int
}

func (s *stack) push(v *stackVal) []*stackVal {
	s.count++
	s.Values = append(s.Values, v)
	return s.Values
}

func (s *stack) pop() (val *stackVal) {
	if s.count == 0 {
		return nil
	}
	val = s.Values[s.count]
	s.count--
	s.Values = s.Values[:s.count]
	return val
}

var memo = make(map[int]int)

// MemoFib ...
func MemoFib(n int) int {
	var f int
	var ok bool

	if f, ok = memo[n]; ok {
		return f
	}
	if n == 0 {
		f = 1
	} else if n <= 2 {
		f = 1
	} else {
		f = MemoFib(n-1) + MemoFib(n-2)
	}
	memo[n] = f
	return memo[n]
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

// Make array to store elements
func q415(A []int) (maxLeft, maxRight, sum int) {
	i := 0
	j := 0
	runningSum := 0

	for i < len(A) {

		if j == len(A) {
			i++
			runningSum = 0
			continue
		}
		runningSum += A[j]
		if runningSum > sum {
			maxLeft = i
			maxRight = j
			sum = runningSum
		}
		fmt.Println(A[i:j])
		if j != len(A) {
			j++
		}

	}
	return maxLeft, maxRight, sum
}

// Only looking at input as prices
func bruteForceStockPrices(priceList []int) (buyDay, sellDay int) {
	bestProfit := 0

	for i, buy := range priceList {
		for j, sell := range priceList[i:] {
			fmt.Println(priceList[i+1:])
			if p := sell - buy; p > bestProfit {
				bestProfit = p
				buyDay, sellDay = i, j+i
			}
		}
	}
	return buyDay, sellDay
}

func bruteFindMaxCrossSubArray(A []int) (maxLeft, maxRight, sum int) {
	var iSum int

	for i := range A {
		iSum = 0
		for j, val := range A[i:] {
			iSum += val

			if iSum >= sum {
				sum = iSum
				maxLeft = i
				maxRight = j
			}
		}
	}
	return maxLeft, maxRight, sum
}

func findMaxCrossingSubArray(A []int, low, high int) (maxLeft, maxRight, sum int) {
	mid := len(A) / 2
	var leftSum, rightSum int
	sum = 0
	for i := mid; i >= low; i-- {
		leftSum += A[i]
		if leftSum > sum {
			sum = leftSum
			maxLeft = i
		}
	}
	sum = 0
	for j := mid + 1; j <= high; j++ {
		rightSum += A[j]
		if rightSum > sum {
			sum = rightSum
			maxRight = j
		}
	}
	sum = leftSum + rightSum
	return maxLeft, maxRight, sum
}

func divideConquerMaxSubArray(A []int, low, high int) (maxLeft, maxRight, sum int) {
	if high == low {
		return low, high, A[low]
	}
	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := findMaxCrossingSubArray(A, low, mid)
	rightLow, rightHigh, rightSum := findMaxCrossingSubArray(A, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossingSubArray(A, low, high)

	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && leftSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}
