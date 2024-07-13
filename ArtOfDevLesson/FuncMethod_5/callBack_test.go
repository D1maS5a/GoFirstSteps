package main

import (
	"reflect"
	"testing"
)

// ПКМ по функции GO:Show All Commands... -> Go: Generate Unit test for ...

func Test_totalPrice(t *testing.T) {
	type args struct {
		initPrice int
	}
	tests := []struct {
		name string
		args args
		want func(int) int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalPrice(tt.args.initPrice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("totalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
