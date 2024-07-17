package main

import "testing"

func TestMultiple(t *testing.T) {
	// setup
	// insert test data in db

	t.Run("groupA", func(t *testing.T) {
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("expected result %d != %d", result, realResult)
			}
			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 222, 222, 49284
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("expected result %d != %d", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("expected result %d != %d", result, realResult)
			}
		})
	})

	t.Run("groupB", func(t *testing.T) {
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("expected result %d != %d", result, realResult)
			}
			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 222, 222, 49284
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("expected result %d != %d", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("expected result %d != %d", result, realResult)
			}
		})
	})

	// t.Run("simple", func(t *testing.T) {
	// 	t.Parallel()
	// 	t.Log("simple")
	// 	var x, y, result = 2, 2, 4
	// 	realResult := Multiple(x, y)
	// 	if realResult != result {
	// 		t.Errorf("expected result %d != %d", result, realResult)
	// 	}
	// 	t.Run("1", func(t *testing.T) {
	// 		r := Multiple(1, 1)
	// 		if r != 1 {
	// 			t.Errorf("failed")
	// 		}
	// 	})
	// })
	// t.Run("medium", func(t *testing.T) {
	// 	t.Parallel()
	// 	t.Log("medium")
	// 	var x, y, result = 222, 222, 49284
	// 	realResult := Multiple(x, y)
	// 	if realResult != result {
	// 		t.Errorf("expected result %d != %d", result, realResult)
	// 	}
	// })

	// t.Run("negative", func(t *testing.T) {
	// 	t.Parallel()
	// 	t.Log("negative")
	// 	var x, y, result = -2, 4, -8
	// 	realResult := Multiple(x, y)
	// 	if realResult != result {
	// 		t.Errorf("expected result %d != %d", result, realResult)
	// 	}
	// })
	// tearDown
	// delete test data in db
}

// go test -v
// go test -v -json
// go test -v -run TestMultiple/simple
// go test -v -run /simple/1

func TestAdd(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		var x, y, result = 2, 2, 4
		realResult := Add(x, y)
		if realResult != result {
			t.Errorf("expected result %d != %d", result, realResult)
		}
		t.Run("1", func(t *testing.T) {
			r := Add(1, 1)
			if r != 2 {
				t.Errorf("failed")
			}
		})
	})
}
