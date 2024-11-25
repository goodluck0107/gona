package algorithm

import (
	"fmt"
	"math"
	"math/rand"

	"gona/boost"
	"gona/systemx"
)

func init() {
	systemx.SeriousPanic(func() {
		rand.Seed(boost.Now().UnixNano())
	})
}

func Floor(value float64, floor int64) float64 {
	v := math.Floor(value/float64(floor)) * float64(floor)
	return v
}

// RandPR rands yes or no by probability
func RandPR(pr float64) bool {
	return rand.Float64() <= pr
}

// RandWeight returns int64 for RandTypeWeight
func RandWeight(v interface{}) int64 {
	r, err := RandTypeWeight(v)
	if err != nil {
		panic(err)
	}
	n, ok := r.(int64)
	if !ok {
		panic(fmt.Errorf("rand weight get invalid type"))
	}
	return n
}

func RandFloatWeight(v interface{}) float64 {
	r, err := RandTypeWeight(v)
	if err != nil {
		panic(err)
	}
	n, ok := r.(float64)
	if !ok {
		panic(fmt.Errorf("rand float weight get invalid type"))
	}
	return n
}

func FastFindIndex(n float64, s []float64) int {
	var (
		min    = 0
		max    = len(s) - 1
		middle int
	)

	for {
		middle = (min + max) / 2
		if max == middle { // finish it: case min == max
			return middle
		} else if min == middle { // finish it: case min + 1 == max
			if n <= s[max] && s[max-1] < n {
				return max
			} else {
				return min
			}
		} else {
			if n <= s[middle-1] {
				max = middle - 1
			} else if s[middle] < n {
				min = middle + 1
			} else {
				return middle
			}
		}
	}
}

// RandTypeWeight rands type by weight
func RandTypeWeight(v interface{}) (interface{}, error) {
	switch weights := v.(type) {
	case []int64:
		var weightSum float64
		for _, r := range weights {
			weightSum += float64(r)
		}
		n := rand.Float64() * weightSum
		var lastKey int64
		for i, r := range weights {
			nReach := float64(r)
			if n <= nReach {
				return int64(i), nil
			}
			n -= nReach
			lastKey = int64(i)
		}
		return lastKey, nil
	case []float64:
		var weightSum float64
		for _, r := range weights {
			weightSum += r
		}
		n := rand.Float64() * weightSum
		var lastKey int64
		for i, r := range weights {
			nReach := r
			if n <= nReach {
				return int64(i), nil
			}
			n -= nReach
			lastKey = int64(i)
		}
		return lastKey, nil
	case [][]int64:
		var weightSum float64
		for _, r := range weights {
			weightSum += float64(r[1])
		}
		n := rand.Float64() * weightSum
		var lastKey int64
		for _, r := range weights {
			nReach := float64(r[1])
			if n <= nReach {
				return r[0], nil
			}
			n -= nReach
			lastKey = r[0]
		}
		return lastKey, nil
	case map[int64]int64:
		var weightSum float64
		for _, r := range weights {
			weightSum += float64(r)
		}
		n := rand.Float64() * weightSum
		var lastKey int64
		for k, r := range weights {
			nReach := float64(r)
			if n <= nReach {
				return k, nil
			}
			n -= nReach
			lastKey = k
		}
		return lastKey, nil
	case map[int64]float64:
		var weightSum float64
		for _, r := range weights {
			weightSum += r
		}
		n := rand.Float64() * weightSum
		var lastKey int64
		for k, r := range weights {
			nReach := r
			if n <= nReach {
				return k, nil
			}
			n -= nReach
			lastKey = k
		}
		return lastKey, nil
	case [][]float64:
		var weightSum float64
		for _, r := range weights {
			weightSum += r[1]
		}
		n := rand.Float64() * weightSum
		var lastKey float64
		for _, r := range weights {
			nReach := r[1]
			if n <= nReach {
				return r[0], nil
			}
			n -= nReach
			lastKey = r[0]
		}
		return lastKey, nil
	case map[float64]float64:
		var weightSum float64
		for _, r := range weights {
			weightSum += r
		}
		n := rand.Float64() * weightSum
		var lastKey float64
		for k, r := range weights {
			nReach := r
			if n <= nReach {
				return k, nil
			}
			n -= nReach
			lastKey = k
		}
		return lastKey, nil
	default:
		return 0, fmt.Errorf("rand type weight get invalid type")
	}
}

// Sum sums values in slice
func Sum(v interface{}) interface{} {
	switch ns := v.(type) {
	case []int:
		var sum int
		for _, n := range ns {
			sum += n
		}
		return sum
	case []int8:
		var sum int8
		for _, n := range ns {
			sum += n
		}
		return sum
	case []int16:
		var sum int16
		for _, n := range ns {
			sum += n
		}
		return sum
	case []int32:
		var sum int32
		for _, n := range ns {
			sum += n
		}
		return sum
	case []int64:
		var sum int64
		for _, n := range ns {
			sum += n
		}
		return sum
	case []uint:
		var sum uint
		for _, n := range ns {
			sum += n
		}
		return sum
	case []uint8:
		var sum uint8
		for _, n := range ns {
			sum += n
		}
		return sum
	case []uint16:
		var sum uint16
		for _, n := range ns {
			sum += n
		}
		return sum
	case []uint32:
		var sum uint32
		for _, n := range ns {
			sum += n
		}
		return sum
	case []uint64:
		var sum uint64
		for _, n := range ns {
			sum += n
		}
		return sum
	case []float32:
		var sum float32
		for _, n := range ns {
			sum += n
		}
		return sum
	case []float64:
		var sum float64
		for _, n := range ns {
			sum += n
		}
		return sum
	default:
		panic(fmt.Errorf("unsuppored slice type for boost.Sum"))
	}
}

// In is to get if the number is in slice
func In(v interface{}, s interface{}) bool {
	switch base := v.(type) {
	case int64:
		for _, n := range s.([]int64) {
			if base == n {
				return true
			}
		}
	case float64:
		for _, n := range s.([]float64) {
			if base == n {
				return true
			}
		}
	case bool:
		for _, n := range s.([]bool) {
			if base == n {
				return true
			}
		}
	case string:
		for _, n := range s.([]string) {
			if base == n {
				return true
			}
		}
	}

	return false
}

// Position position element in a slice or map
func Position(v interface{}, s interface{}) ([]int64, int64) {
	pos := make([]int64, 0)
	switch s := s.(type) {
	case []int64:
		v := v.(int64)
		for p, n := range s {
			if n == v {
				pos = append(pos, int64(p))
			}
		}
	case []float64:
		v := v.(float64)
		for p, n := range s {
			if n == v {
				pos = append(pos, int64(p))
			}
		}
	}
	return pos, int64(len(pos))
}

// Count counts element in a slice or map
func Count(v interface{}, s interface{}) int64 {
	_, count := Position(v, s)
	return count
}

// MaxContinuousCount returns max continuous count in slice  [start end) index
func MaxContinuousCount(v int64, s []int64) (int64, int64, int64) {
	maxCount := 0
	start := 0
	end := 0
	for index := 0; index < len(s); {
		if s[index] == v {
			tmp := index + 1
			for tmp < len(s) && s[tmp] == v {
				tmp++
			}
			if tmp-index > maxCount {
				maxCount = tmp - index
				start = index
				end = tmp
			}
			index = tmp
		} else {
			index++
		}
	}
	return int64(maxCount), int64(start), int64(end)
}

// ShuffleArray shuffle array
func ShuffleArray(v interface{}) interface{} {
	switch array := v.(type) {
	case []int64:
		length := int64(len(array))
		for i := length; i > 0; i-- {
			pos := rand.Int63n(i)
			array[length-1], array[pos] = array[pos], array[length-1]
		}
		return array
	case []float64:
		length := int64(len(array))
		for i := length; i > 0; i-- {
			pos := rand.Int63n(i)
			array[length-1], array[pos] = array[pos], array[length-1]
		}
		return array
	}
	return v
}

// Replace replace element form arry
func Replace(v interface{}, o interface{}, n interface{}) interface{} {
	switch array := v.(type) {
	case []int64:
		length := len(array)
		for i := 0; i < length; i++ {
			if array[i] == o {
				array[i] = n.(int64)
			}
		}
		return array
	case []float64:
		length := len(array)
		for i := 0; i < length; i++ {
			if array[i] == o {
				array[i] = n.(float64)
			}
		}
		return array
	}
	return v
}

// RandUnRepeatedMany 以s为权重随机count个不重复的 返回索引的切片
func RandUnRepeatedMany(s interface{}, count int64) ([]int64, error) {
	switch s := s.(type) {
	case []int64:
		input := s
		if count <= 0 || count > int64(len(input)) {
			return nil, fmt.Errorf("parameter invalid")
		}
		scopy := make([]int64, len(input))
		copy(scopy, input)

		result := make([]int64, 0)
		for int64(len(result)) < count {
			var index int64
			for index = RandWeight(scopy); In(index, result); {
				index = RandWeight(scopy)
			}
			result = append(result, index)
		}
		return result, nil
	case []float64:
		input := s
		if count <= 0 || count > int64(len(input)) {
			return nil, fmt.Errorf("parameter invalid")
		}
		scopy := make([]float64, len(input))
		copy(scopy, input)

		result := make([]int64, 0)
		for int64(len(result)) < count {
			var index int64
			for index = RandWeight(scopy); In(index, result); {
				index = RandWeight(scopy)
			}
			result = append(result, index)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("parameter type invalid")
	}
}

func StringMod(s string, n int) int {
	var sum int
	for _, b := range []byte(s) {
		sum += int(b)
	}
	return sum % n
}

// RandIntn Picks a random value in the specified object
func RandIntn(s int, e int) int {
	return rand.Intn(e-s+1) + s
}
