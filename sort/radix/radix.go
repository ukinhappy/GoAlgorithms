//  基数
package radix

//	"fmt"

func Radix(arr []int, nLength int) {
	if nil == arr || nLength < 0 {
		return
	}

	// 找到整个数组的最大值
	max := findMax(arr, nLength)

	// 查找基数的最大值
	maxRadix := findMaxRadix(max)

	// 以此按照各级基数 存放数据入桶，然后存入下一阶桶

	maparr := make([][]int, 10)
	for i := 0; i < 10; i++ {
		maparr = append(maparr, make([]int, 0))
	}

	for radix := 1; radix <= maxRadix; radix++ {
		for _, value := range arr {
			index := value / GetBei(radix) % (10)
			maparr[index] = append(maparr[index], value)
		}

		i := 0
		for _, indexvalue := range maparr {
			for _, value := range indexvalue {
				arr[i] = value
				i++
			}
		}
		maparr = make([][]int, 10)
		for i := 0; i < 10; i++ {
			maparr = append(maparr, make([]int, 0))
		}

	}
}

func GetBei(radix int) int {
	var value int = 1
	for {
		if radix > 1 {
			value *= 10
			radix--
		} else {
			break
		}

	}
	return value
}

func findMax(arr []int, nLength int) int {
	var max int
	max = arr[0]
	for i := 1; i < nLength; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func findMaxRadix(max int) int {
	var maxRadix int
	maxRadix = 1
	for {
		if max/10 == 0 {
			break
		}
		maxRadix++
		max = max / 10
	}

	return maxRadix
}
