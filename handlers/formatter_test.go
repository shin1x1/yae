package handlers

import (
	"bytes"
	"reflect"
	"testing"
)

func TestYamlFormatter_Run(t *testing.T) {
	type fields struct {
		filepath string
	}
	tests := []struct {
		name   string
		fields fields
		output string
	}{
		{
			"test1",
			fields{"./test1_actual.yaml"},
			"./test1_expected.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y1 := &YamlFormatter{
				filepath: tt.fields.filepath,
			}
			buf1 := &bytes.Buffer{}
			y1.Run(buf1)

			y2 := &YamlFormatter{
				filepath: tt.output,
			}
			buf2 := &bytes.Buffer{}
			y2.Run(buf2)

			if !reflect.DeepEqual(buf1, buf2) {
				t.Error("test error\n", "\n=== actual:\n"+buf1.String(), "\n=== expected:\n"+buf2.String())
			}
		})
	}
}
