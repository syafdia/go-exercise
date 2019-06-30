package linkedlist

import (
	"reflect"
	"testing"

	"github.com/syafdia/go-exercise/src/datastructure/types"
)

func Test_node_Head(t *testing.T) {
	type fields struct {
		head types.T
		tail LinkedList
	}
	tests := []struct {
		name   string
		fields fields
		want   types.T
	}{
		{
			name: "success return head",
			fields: fields{
				head: 1,
				tail: nil,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := node{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := n.Head(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Head() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Tail(t *testing.T) {
	type fields struct {
		head types.T
		tail LinkedList
	}
	tests := []struct {
		name   string
		fields fields
		want   LinkedList
	}{
		{
			name: "success return non nil tail",
			fields: fields{
				head: 1,
				tail: node{
					head: 2,
					tail: nil,
				},
			},
			want: node{
				head: 2,
				tail: nil,
			},
		},
		{
			name: "success return nil tail",
			fields: fields{
				head: 1,
				tail: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := node{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := n.Tail(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Tail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Map(t *testing.T) {
	type fields struct {
		head types.T
		tail LinkedList
	}
	type args struct {
		mapper func(t types.T) types.U
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   LinkedList
	}{
		// {
		// 	name: "success mapping data",
		// 	fields: fields{
		// 		head: 1,
		// 		tail: node{
		// 			head: 2,
		// 			tail: nil,
		// 		},
		// 	},
		// 	args: args{
		// 		mapper: func(v types.T) types.U {
		// 			return v.(int) * 10
		// 		},
		// 	},
		// 	want: node{
		// 		head: 10,
		// 		tail: node{
		// 			head: 20,
		// 			tail: nil,
		// 		},
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := node{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := n.Map(tt.args.mapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
