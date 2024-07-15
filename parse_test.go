package gohtml

import (
	"reflect"
	"testing"
)

func Test_parseActionNodeField(t *testing.T) {
	type args struct {
		path []string
		pipe string
	}
	tests := []struct {
		name string
		args args
		want StructField
	}{
		{
			name: "basic item",
			args: args{
				path: []string{"Hello"},
				pipe: ".Name",
			},
			want: StructField{
				Path: []string{"Hello"},
				Name: "Name",
				Type: "any",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseActionNodeField(tt.args.path, tt.args.pipe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseActionNodeField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRangeNodeField(t *testing.T) {
	type args struct {
		path []string
		pipe string
	}
	tests := []struct {
		name string
		args args
		want StructField
	}{
		{
			name: "basic loop item",
			args: args{
				path: []string{"Loop"},
				pipe: "Item := .Options",
			},
			want: StructField{
				Path: []string{"Loop"},
				Name: "Options",
				Type: "[]LoopOptionsItem",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRangeNodeField(tt.args.path, tt.args.pipe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRangeNodeField()\ngot: %+#v\nwant: %+#v", got, tt.want)
			}
		})
	}
}
