package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAdapter(t *testing.T) {
	tests := []struct {
		name string
		want *Adapter
	}{
		{name: "create new adapter", want: &Adapter{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdapter(); !assert.Equal(t, got, tt.want) {
				t.Errorf("NewAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapter_Addition(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name    string
		arith   *Adapter
		args    args
		want    int32
		wantErr bool
	}{
		{name: "addition of 1 and 2", arith: NewAdapter(), args: args{a: 1, b: 2}, want: 3, wantErr: false},
	}
	assert.Greater(t, len(tests), 0, "no test cases found")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.arith.Addition(tt.args.a, tt.args.b)
			assert.Nil(t, err, "error should be nil")
			assert.Equal(t, tt.want, got, "values should be equal")
		})
	}
}

func TestAdapter_Subtraction(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name    string
		arith   *Adapter
		args    args
		want    int32
		wantErr bool
	}{
		{name: "positive subtraction", arith: NewAdapter(), args: args{a: 5, b: 2}, want: 5 - 2, wantErr: false},
		{name: "negative subtraction", arith: NewAdapter(), args: args{a: 2, b: 5}, want: 2 - 5, wantErr: false},
	}

	assert.Greater(t, len(tests), 0, "no test cases found")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.arith.Subtraction(tt.args.a, tt.args.b)
			assert.Nil(t, err, "error should be nil")
			assert.Equal(t, tt.want, got, "values should be equal")
		})
	}
}

func TestAdapter_Multiplication(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name    string
		arith   *Adapter
		args    args
		want    int32
		wantErr bool
	}{
		{name: "multiplication of 1 and 2", arith: NewAdapter(), args: args{a: 1, b: 2}, want: 1 * 2, wantErr: false},
	}
	assert.Greater(t, len(tests), 0, "no test cases found")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.arith.Multiplication(tt.args.a, tt.args.b)
			assert.Nil(t, err, "error should be nil")
			assert.Equal(t, tt.want, got, "values should be equal")
		})
	}
}

func TestAdapter_Division(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name    string
		arith   *Adapter
		args    args
		want    int32
		wantErr bool
	}{
		{name: "integer division", arith: NewAdapter(), args: args{a: 6, b: 2}, want: 3, wantErr: false},
		{name: "division by zero", arith: NewAdapter(), args: args{a: 6, b: 0}, want: 0, wantErr: true},
		{name: "division by negative number", arith: NewAdapter(), args: args{a: 6, b: -2}, want: -3, wantErr: false},
		{name: "non-integer division", arith: NewAdapter(), args: args{a: 5, b: 2}, want: 2, wantErr: false},
	}
	assert.Greater(t, len(tests), 0, "no test cases found")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.arith.Division(tt.args.a, tt.args.b)
			if tt.wantErr {
				assert.NotNil(t, err, "error should not be nil")
			} else {
				assert.Nil(t, err, "error should be nil")
				assert.Equal(t, tt.want, got, "values should be equal")
			}
		})
	}
}
