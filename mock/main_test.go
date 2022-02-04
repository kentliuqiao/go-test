package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

var (
	want = `3
2
1
Go!
`
	wantOps = []Op{OpSleep, OpWrite, OpSleep, OpWrite, OpSleep, OpWrite, OpSleep, OpWrite}
)

var (
	spy = &SpyOperations{Out: &bytes.Buffer{}}
	out = &bytes.Buffer{}
)

func TestCountDown(t *testing.T) {
	type args struct {
		sleeper Sleeper
	}
	tests := []struct {
		name    string
		args    args
		out     io.Writer
		wantOut string
		wantOps []Op
	}{
		{
			name: "count down",
			args: args{
				sleeper: &SpySleep{},
			},
			out:     out,
			wantOut: want,
			wantOps: nil,
		},
		{
			name:    "spy operation",
			args:    args{sleeper: spy},
			wantOut: want,
			out:     spy.Out,
			wantOps: wantOps,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountDown(tt.out, tt.args.sleeper)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("CountDown() = %v, want %v", gotOut, tt.wantOut)
			}
			if !reflect.DeepEqual(tt.wantOps, spy.Calls) {
				t.Errorf("ops want %v, bugt got %v", wantOps, spy.Calls)
			}
		})
	}
}
