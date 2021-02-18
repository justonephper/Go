# 常用排序算法

## 1.冒泡排序
```
func bubbleSort(arr []int) {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
			   arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
```

## 2.选择排序
```
func selectSort(arr []int) {
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			//比较，取较小
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}
```

## 3.插入排序
```
func insertSort(arr []int) {
	for i := 1; i < length; i++ {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
	fmt.Println(arr)
}
```

## 4.归并排序
```
func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	var tmparr = []int{}
	var s1, s2 = start, mid + 1
	for s1 <= mid && s2 <= end {
		if arr[s1] <= arr[s2] {
			tmparr = append(tmparr, arr[s1])
			s1++
		} else {
			tmparr = append(tmparr, arr[s2])
			s2++
		}
	}

	if s1 <= mid {
		tmparr = append(tmparr, arr[s1:mid+1]...)
	}
	if s2 <= end {
		tmparr = append(tmparr, arr[s2:end+1]...)
	}

	for pos, item := range tmparr {
		arr[start+pos] = item
	}
}
```

## 5.快速排序
```
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	spliteData := arr[0]
	low := make([]int, 0, 0)
	hight := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, spliteData)

	for i := 1; i < len(arr); i++ {
		if arr[i] < spliteData {
			low = append(low, arr[i])
		} else if arr[i] > spliteData {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = quickSort(low), quickSort(hight)
	newArr := append(append(low, mid...), hight...)
	return newArr
}

```
