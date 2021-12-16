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
