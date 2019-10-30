package unit

import "testing"

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	excepted := [...]int{1, 4, 9}

	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != excepted[i] {
			t.Errorf("input is %d,the excepted is %d,the actual is %d", inputs[i], excepted[i], ret)
		}
	}
}
