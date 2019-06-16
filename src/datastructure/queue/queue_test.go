package queue

import (
	"reflect"
	"testing"

	"github.com/syafdia/go-exercise/src/types"
)

func Test_queue_Enqueue(t *testing.T) {
	type fields struct {
		values []types.T
	}
	type args struct {
		v types.T
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []types.T
	}{
		{
			name: "success adding to queue",
			fields: fields{
				values: []types.T{1, 2, 3},
			},
			args: args{
				v: 4,
			},
			want: []types.T{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := queue{
				values: tt.fields.values,
			}
			q.Enqueue(tt.args.v)

			if !reflect.DeepEqual(q.values, tt.want) {
				t.Errorf("q.values = %v, want %v", q.values, tt.want)
			}
		})
	}
}

func Test_queue_Dequeue(t *testing.T) {
	type fields struct {
		values []types.T
	}
	tests := []struct {
		name   string
		fields fields
		want   types.T
		want2  []types.T
	}{
		{
			name: "success dequeueing from queue",
			fields: fields{
				values: []types.T{1, 2, 3},
			},
			want:  1,
			want2: []types.T{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				values: tt.fields.values,
			}
			if got := q.Dequeue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queue.Dequeue() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(q.values, tt.want2) {
				t.Errorf("q.values = %v, want %v", q.values, tt.want2)
			}
		})
	}
}

func Test_queue_Size(t *testing.T) {
	type fields struct {
		values []types.T
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "success counting queue",
			fields: fields{
				values: []types.T{1, 2, 3, 4, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				values: tt.fields.values,
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("queue.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
