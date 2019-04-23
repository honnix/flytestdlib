// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package logger

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeSlice_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_show-source", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vBool, err := cmdFlags.GetBool("show-source"); err == nil {
				assert.Equal(t, bool(defaultConfig.IncludeSourceCode), vBool)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("show-source", testValue)
			if vBool, err := cmdFlags.GetBool("show-source"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.IncludeSourceCode)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_mute", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vBool, err := cmdFlags.GetBool("mute"); err == nil {
				assert.Equal(t, bool(defaultConfig.Mute), vBool)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("mute", testValue)
			if vBool, err := cmdFlags.GetBool("mute"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.Mute)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_level", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("level"); err == nil {
				assert.Equal(t, int(defaultConfig.Level), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("level", testValue)
			if vInt, err := cmdFlags.GetInt("level"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Level)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_formatter.type", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("formatter.type"); err == nil {
				assert.Equal(t, string(defaultConfig.Formatter.Type), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("formatter.type", testValue)
			if vString, err := cmdFlags.GetString("formatter.type"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Formatter.Type)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}

func TestConfig_elemValueOrNil(t *testing.T) {
	type fields struct {
		IncludeSourceCode bool
		Mute              bool
		Level             Level
		Formatter         FormatterConfig
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				IncludeSourceCode: tt.fields.IncludeSourceCode,
				Mute:              tt.fields.Mute,
				Level:             tt.fields.Level,
				Formatter:         tt.fields.Formatter,
			}
			if got := c.elemValueOrNil(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.elemValueOrNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
