package genfuncs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type PersonName struct {
	First string
	Last  string
}

func TestAny(t *testing.T) {
	type args struct {
		slice     []string
		predicate Predicate[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty",
			args: args{
				slice:     []string{},
				predicate: func(s string) bool { return s == "a" },
			},
			want: false,
		},
		{
			name: "Not Found",
			args: args{
				slice:     []string{"b", "c"},
				predicate: func(s string) bool { return s == "a" },
			},
			want: false,
		},
		{
			name: "Found",
			args: args{
				slice:     []string{"b", "a", "c"},
				predicate: func(s string) bool { return s == "a" },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Any(tt.args.slice, tt.args.predicate)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestAssociateBy(t *testing.T) {
	var fName KeySelector[PersonName, string] = func(p PersonName) string { return p.First }
	type args struct {
		slice       []PersonName
		keySelector KeySelector[PersonName, string]
	}
	tests := []struct {
		name     string
		args     args
		wantSize int
		contains []string
	}{
		{
			name: "Empty",
			args: args{
				slice:       []PersonName{},
				keySelector: fName,
			},
			wantSize: 0,
		},
		{
			name: "Two Unique",
			args: args{
				slice: []PersonName{
					{
						First: "fred",
						Last:  "flintstone",
					},
					{
						First: "barney",
						Last:  "rubble",
					},
				},
				keySelector: fName,
			},
			wantSize: 2,
			contains: []string{"fred", "baarney"},
		},
		{
			name: "Duplicate",
			args: args{
				slice: []PersonName{
					{
						First: "fred",
						Last:  "flintstone",
					},
					{
						First: "fred",
						Last:  "astaire",
					},
				},
				keySelector: fName,
			},
			wantSize: 1,
			contains: []string{"fred"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fNameMap := AssociateBy(tt.args.slice, tt.args.keySelector)
			assert.Equal(t, tt.wantSize, len(fNameMap))
			for k, _ := range fNameMap {
				_, ok := fNameMap[k]
				assert.True(t, ok)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		slice     []int
		predicate Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty",
			args: args{
				slice:     []int{},
				predicate: func(i int) bool { return true },
			},
			want: nil,
		},
		{
			name: "Evens",
			args: args{
				slice:     []int{1, 2, 3, 4},
				predicate: func(i int) bool { return i%2 == 0 },
			},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Filter(tt.args.slice, tt.args.predicate)
			assert.Equal(t, len(tt.want), len(result))
			for _, v := range result {
				assert.True(t, Any(result, func(i int) bool { return i == v }))
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		slice     []float32
		predicate Predicate[float32]
	}
	tests := []struct {
		name      string
		args      args
		want      float32
		wantFound bool
	}{
		{
			name: "Empty",
			args: args{
				slice:     []float32{},
				predicate: func(f float32) bool { return false },
			},
			wantFound: false,
		},
		{
			name: "Not Found",
			args: args{
				slice:     []float32{1.0, 2.0, 3.0},
				predicate: func(f float32) bool { return false },
			},
			wantFound: false,
		},
		{
			name: "Found",
			args: args{
				slice:     []float32{1.0, 2.0, 3.0},
				predicate: func(f float32) bool { return f > 1.0 },
			},
			want: 2.0,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Find(tt.args.slice, tt.args.predicate)
			if !tt.wantFound {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFindLast(t *testing.T) {
	type args struct {
		slice     []float32
		predicate Predicate[float32]
	}
	tests := []struct {
		name      string
		args      args
		want      float32
		wantFound bool
	}{
		{
			name: "Empty",
			args: args{
				slice:     []float32{},
				predicate: func(f float32) bool { return false },
			},
			wantFound: false,
		},
		{
			name: "Not Found",
			args: args{
				slice:     []float32{1.0, 2.0, 3.0},
				predicate: func(f float32) bool { return f  < 0.0 },
			},
			wantFound: false,
		},
		{
			name: "Found",
			args: args{
				slice:     []float32{1.0, 2.0, 3.0},
				predicate: func(f float32) bool { return f > 1.0 },
			},
			want: 3.0,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := FindLast(tt.args.slice, tt.args.predicate)
			if !tt.wantFound {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFold(t *testing.T) {
	si := []int{1, 2, 3}
	sum := Fold(si, 10, func(r int, i int) int { return r + i  })
	assert.Equal(t, 16, sum)
}
