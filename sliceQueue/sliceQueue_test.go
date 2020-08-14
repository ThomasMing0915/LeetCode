package sliceQueue

import (
	"fmt"
	"testing"
)

func TestSliceQueue_DeQueue(t *testing.T) {
	sq := SliceQueue{array: make([]int, 0), front: 0, rear: 0}
	sq.EnQueue(1)
	sq.EnQueue(2)
	fmt.Println(sq.Size())
	sq.DeQueue()
	fmt.Println(sq.Size())
}

func TestSliceQueue_EnQueue(t *testing.T) {
	type fields struct {
		array []int
		front int
		rear  int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SliceQueue{
				array: tt.fields.array,
				front: tt.fields.front,
				rear:  tt.fields.rear,
			}
		})
	}
}

func TestSliceQueue_GetFront(t *testing.T) {
	type fields struct {
		array []int
		front int
		rear  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SliceQueue{
				array: tt.fields.array,
				front: tt.fields.front,
				rear:  tt.fields.rear,
			}
			if got := s.GetFront(); got != tt.want {
				t.Errorf("GetFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceQueue_GetRear(t *testing.T) {
	type fields struct {
		array []int
		front int
		rear  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SliceQueue{
				array: tt.fields.array,
				front: tt.fields.front,
				rear:  tt.fields.rear,
			}
			if got := s.GetRear(); got != tt.want {
				t.Errorf("GetRear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceQueue_IsEmpty(t *testing.T) {
	type fields struct {
		array []int
		front int
		rear  int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SliceQueue{
				array: tt.fields.array,
				front: tt.fields.front,
				rear:  tt.fields.rear,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceQueue_Size(t *testing.T) {
	type fields struct {
		array []int
		front int
		rear  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SliceQueue{
				array: tt.fields.array,
				front: tt.fields.front,
				rear:  tt.fields.rear,
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
