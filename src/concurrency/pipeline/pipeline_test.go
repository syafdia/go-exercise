package pipeline

import (
	"testing"
)

func Test_pipeline_Pipe(t *testing.T) {
	type fields struct {
		dataC     chan interface{}
		errC      chan error
		executors []Executor
	}
	type args struct {
		executor Executor
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		check  func(f fields) bool
	}{
		{
			name: "should success adding executor to executors",
			fields: fields{
				executors: []Executor{},
			},
			args: args{
				executor: func(in interface{}) (interface{}, error) {
					return 1, nil
				},
			},
			check: func(f fields) bool {
				return len(f.executors) == 1
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pipeline{
				dataC:     tt.fields.dataC,
				errC:      tt.fields.errC,
				executors: tt.fields.executors,
			}
			p.Pipe(tt.args.executor)
			if !tt.check(fields{p.dataC, p.errC, p.executors}) {
				t.Errorf("pipeline.Pipe() not run as expected")
			}
		})
	}
}
