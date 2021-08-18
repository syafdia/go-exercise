package calc

import "testing"

func TestFactorial(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "N is 5",
			args: args{n: 5},
			want: 120,
		},
		{
			name: "N is 0",
			args: args{n: 0},
			want: 1,
		},
		{
			name: "N is 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "N is 3",
			args: args{n: 3},
			want: 6,
		},
		{
			name: "N is 10",
			args: args{n: 10},
			want: 3628800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
