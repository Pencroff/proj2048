package storage

import (
	"fmt"
	"github.com/linkedin/goavro/v2"
	"testing"
)

func Test_avro_e2e(t *testing.T) {
	codec, err := goavro.NewCodec(`
        {
          "type": "record",
          "name": "LongList",
          "fields" : [
            {"name": "next", "type": ["null", "LongList"], "default": null}
          ]
        }`)
	if err != nil {
		fmt.Println(err)
	}

	// NOTE: May omit fields when using default value
	textual := []byte(`{"next":{"LongList":{}}}`)

	// Convert textual Avro data (in Avro JSON format) to native Go form
	native, _, err := codec.NativeFromTextual(textual)
	if err != nil {
		fmt.Println(err)
	}

	// Convert native Go form to binary Avro data
	binary, err := codec.BinaryFromNative(nil, native)
	if err != nil {
		fmt.Println(err)
	}

	// Convert binary Avro data back to native Go form
	native, _, err = codec.NativeFromBinary(binary)
	if err != nil {
		fmt.Println(err)
	}

	// Convert native Go form to textual Avro data
	textual, err = codec.TextualFromNative(nil, native)
	if err != nil {
		fmt.Println(err)
	}

	// NOTE: Textual encoding will show all fields, even those with values that
	// match their default values
	fmt.Println(string(textual))
}
