package alc_lang

import "testing"

func TestTagFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Tag
	}{
		{name: "1", args: args{s: "cn"}, want: Cn},
		{name: "2", args: args{s: "en"}, want: En},
		{name: "3", args: args{s: ""}, want: None},
		{name: "3", args: args{s: "test"}, want: None},
		{name: "3", args: args{s: "en_"}, want: None},
		{name: "3", args: args{s: "en-"}, want: None},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TagFromString(tt.args.s); got != tt.want {
				t.Errorf("TagFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
