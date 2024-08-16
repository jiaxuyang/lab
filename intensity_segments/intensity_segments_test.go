package intensity_segments

import (
	"reflect"
	"testing"
)

func TestAddHelper(t *testing.T) {
	tests := []struct {
		name    string
		inputs  [][]int
		outputs []string
	}{
		{name: "example1", inputs: [][]int{
			{10, 30, 1}, {20, 40, 1}, {10, 40, -2},
		}, outputs: []string{
			"[[10,1],[30,0]]", "[[10,1],[20,2],[30,1],[40,0]]", "[[10,-1],[20,0],[30,-1],[40,0]]",
		}},
		{name: "example2", inputs: [][]int{
			{10, 30, 1}, {20, 40, 1}, {10, 40, -1}, {10, 40, -1},
		}, outputs: []string{
			"[[10,1],[30,0]]", "[[10,1],[20,2],[30,1],[40,0]]", "[[20,1],[30,0]]", "[[10,-1],[20,0],[30,-1],[40,0]]",
		}},
		{name: "corner", inputs: [][]int{
			{10, 30, 1}, {0, 0, 100}, {20, 10, 100}, {-999, 999, 0},
		}, outputs: []string{
			"[[10,1],[30,0]]", "[[10,1],[30,0]]", "[[10,1],[30,0]]", "[[10,1],[30,0]]",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if outputs := AddHelper(tt.inputs); !reflect.DeepEqual(outputs, tt.outputs) {
				t.Errorf("TestAddHelper() = %+v, want %+v", outputs, tt.outputs)
			}
		})
	}
}

func TestSetHelper(t *testing.T) {
	tests := []struct {
		name    string
		inputs  [][]int
		outputs []string
	}{
		{name: "example1", inputs: [][]int{
			{10, 30, 1}, {20, 40, 1}, {10, 40, -2},
		}, outputs: []string{
			"[[10,1],[30,0]]", "[[10,1],[40,0]]", "[[10,-2],[40,0]]",
		}},
		{name: "example2", inputs: [][]int{
			{10, 30, 1}, {20, 40, 1}, {10, 40, -1}, {10, 40, -1},
		}, outputs: []string{
			"[[10,1],[30,0]]", "[[10,1],[40,0]]", "[[10,-1],[40,0]]", "[[10,-1],[40,0]]",
		}},
		{name: "corner", inputs: [][]int{
			{10, 30, 1}, {0, 0, 100}, {20, 10, 100}, {-999, 999, 0},
		}, outputs: []string{
			"[[10,1],[30,0]]", "[[10,1],[30,0]]", "[[10,1],[30,0]]", "[]",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if outputs := SetHelper(tt.inputs); !reflect.DeepEqual(outputs, tt.outputs) {
				t.Errorf("TestAddHelper() = %+v, want %+v", outputs, tt.outputs)
			}
		})
	}
}
