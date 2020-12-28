package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntQueue_Subset(t *testing.T) {
	tests := []struct {
		name string
		qty  int
		want *IntQueue
	}{
		{"a", 3, &IntQueue{
			store:    []int{1, 2, 3, 4, 0},
			capacity: 5,
			head:     0,
			tail:     3,
		},
		},
		{"b", 2, &IntQueue{
			store:    []int{1, 2, 3, 4, 0},
			capacity: 5,
			head:     0,
			tail:     2,
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := newIntQueue(5)
			for i := 1; i < 5; i++ {
				q.Enqueue(i)
			}
			if got := q.Subset(tt.qty); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got.Size(), got.capacity, got.head, got.tail, got.store)
				fmt.Println(tt.want.Size(), tt.want.capacity, tt.want.head, tt.want.tail, tt.want.store)
				t.Errorf("IntQueue.Subset() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}
