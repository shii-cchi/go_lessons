package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	///// можно выполнять какие-то действия с данными, как и в отдельном тесте
	res := m.Run()
	/////
	os.Exit(res)
}

//func TestMultiple(t *testing.T) { // один тест в одной функции
//	var x, y, result = 2, 2, 4
//
//	realResult := Multiple(x, y)
//
//	if realResult != result {
//		t.Errorf("%d != %d", result, realResult)
//	}
//}

func TestMultiple1(t *testing.T) { // много тестов в одной
	t.Cleanup(func() {
		fmt.Println("cleenup") // выполниться в конце всех проверок, если будет паника, то все равно выполниться
	})

	t.Run("test_1", func(t *testing.T) {
		var x, y, result = 2, 2, 4

		realResult := Multiple(x, y)

		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}
	})

	t.Run("test_2", func(t *testing.T) {
		panic("something goes wrong")
	})

	t.Run("test_3", func(t *testing.T) {
		var x, y, result = -1, 1, -1

		realResult := Multiple(x, y)

		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}
	})
}

//func TestMultiple2(t *testing.T) { // много тестов в одной параллельно
//	t.Run("test_1", func(t *testing.T) {
//		t.Parallel()
//		t.Log("test_1")
//		var x, y, result = 2, 2, 4
//
//		realResult := Multiple(x, y)
//
//		if realResult != result {
//			t.Errorf("%d != %d", result, realResult)
//		}
//	})
//
//	t.Run("test_2", func(t *testing.T) {
//		t.Parallel()
//		t.Log("test_2")
//		var x, y, result = 2, 0, 0
//
//		realResult := Multiple(x, y)
//
//		if realResult != result {
//			t.Errorf("%d != %d", result, realResult)
//		}
//	})
//
//	t.Run("test_3", func(t *testing.T) {
//		t.Parallel()
//		t.Log("test_3")
//		var x, y, result = -1, 1, -1
//
//		realResult := Multiple(x, y)
//
//		if realResult != result {
//			t.Errorf("%d != %d", result, realResult)
//		}
//	})
//}

//func TestMultiple3(t *testing.T) { // много тестов в одной параллельно
//	t.Run("group1", func(t *testing.T) {
//		t.Log("1")
//		t.Run("test1", func(t *testing.T) {
//			t.Parallel()
//			t.Log("test_1")
//			var x, y, result = 2, 2, 4
//
//			realResult := Multiple(x, y)
//
//			if realResult != result {
//				t.Errorf("%d != %d", result, realResult)
//			}
//		})
//		t.Run("test2", func(t *testing.T) {
//			t.Parallel()
//			t.Log("test_2")
//			var x, y, result = 2, 0, 0
//
//			realResult := Multiple(x, y)
//
//			if realResult != result {
//				t.Errorf("%d != %d", result, realResult)
//			}
//		})
//		t.Run("test3", func(t *testing.T) {
//			t.Parallel()
//			t.Log("test_3")
//			var x, y, result = -1, 1, -1
//
//			realResult := Multiple(x, y)
//
//			if realResult != result {
//				t.Errorf("%d != %d", result, realResult)
//			}
//		})
//	})
//
//	t.Run("group2", func(t *testing.T) {
//		t.Log("2")
//		t.Run("test1", func(t *testing.T) {
//			t.Parallel()
//			t.Log("test_1")
//			var x, y, result = 2, 2, 4
//
//			realResult := Multiple(x, y)
//
//			if realResult != result {
//				t.Errorf("%d != %d", result, realResult)
//			}
//		})
//		t.Run("test2", func(t *testing.T) {
//			t.Parallel()
//			t.Log("test_2")
//			var x, y, result = 2, 0, 0
//
//			realResult := Multiple(x, y)
//
//			if realResult != result {
//				t.Errorf("%d != %d", result, realResult)
//			}
//		})
//		t.Run("test3", func(t *testing.T) {
//			t.Parallel()
//			t.Log("test_3")
//			var x, y, result = -1, 1, -1
//
//			realResult := Multiple(x, y)
//
//			if realResult != result {
//				t.Errorf("%d != %d", result, realResult)
//			}
//		})
//	})
//
//	t.Run("group3", func(t *testing.T) {
//		t.Log("3")
//		t.Run("test1", func(t *testing.T) {
//			t.Parallel()
//			t.Log("test_1")
//			var x, y, result = 2, 2, 4
//
//			realResult := Multiple(x, y)
//
//			if realResult != result {
//				t.Errorf("%d != %d", result, realResult)
//			}
//		})
//		t.Run("test2", func(t *testing.T) {
//			t.Parallel()
//			t.Log("test_2")
//			var x, y, result = 2, 0, 0
//
//			realResult := Multiple(x, y)
//
//			if realResult != result {
//				t.Errorf("%d != %d", result, realResult)
//			}
//		})
//		t.Run("test3", func(t *testing.T) {
//			t.Parallel()
//			t.Log("test_3")
//			var x, y, result = -1, 1, -1
//
//			realResult := Multiple(x, y)
//
//			if realResult != result {
//				t.Errorf("%d != %d", result, realResult)
//			}
//		})
//	})
//}
