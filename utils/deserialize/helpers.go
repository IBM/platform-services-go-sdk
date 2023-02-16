/**
 * (C) Copyright IBM Corp. 2023.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package deserialize

import (
	"encoding/json"
	"github.com/IBM/go-sdk-core/v5/core"
	translation "github.com/IBM/platform-services-go-sdk/i18n"
	"os"
	"reflect"
	"strings"
)

func getJsonStringAsBytes(json string) (stringAsBytes []byte, err error, msg string) {
	if readAsFile(json) {
		// read the json from a file
		// the [1:] removes the @ symbol from the string, used to designate a file
		fileContents, fileErr := os.ReadFile(json[1:])
		if fileErr != nil {
			err = fileErr
			msg = translation.T("file-reading-error", map[string]interface{}{
				"FILENAME": json[1:],
			})
			return
		}
		stringAsBytes = fileContents
	} else {
		stringAsBytes = []byte(json)
	}

	return
}

func readAsFile(userInput string) bool {
	return strings.HasPrefix(userInput, "@")
}

func deserializeJson(contents []byte, flagName string, dataType string, result interface{}) (err error, msg string) {
	err = json.Unmarshal(contents, result)
	if err != nil {
		msg = translation.T("parsing-error", map[string]interface{}{
			"FLAG_NAME": flagName,
			"TYPE":      dataType,
		})
	}

	return
}

func deserializeModel(flag string, flagName string, modelName string, unmarshalFunc core.ModelUnmarshaller, slice bool, result interface{}) (err error, msg string) {
	var valueAsBytes []byte
	valueAsBytes, err, msg = getJsonStringAsBytes(flag)
	if err != nil {
		return
	}

	// the only difference between single models and slices of models is the type
	// of the generic value we pass to the unmarshaller. this bit of code prevents
	// duplication by using the right type for the right scenario
	var raw interface{}
	if slice {
		var generic []json.RawMessage
		err = json.Unmarshal(valueAsBytes, &generic)
		raw = generic
	} else {
		var generic map[string]json.RawMessage
		err = json.Unmarshal(valueAsBytes, &generic)
		raw = generic
	}

	if err != nil {
		msg = translation.T("parsing-error", map[string]interface{}{
			"FLAG_NAME": flagName,
			"TYPE":      "JSON",
		})
		return
	}

	err = core.UnmarshalModel(raw, "", result, unmarshalFunc)
	if err != nil {
		msg = translation.T("parsing-error", map[string]interface{}{
			"FLAG_NAME": flagName,
			"TYPE":      "model " + modelName,
		})
		return
	}

	return
}

func toJsonValue(s string) string {
	// if the value is not already JSON, assume it just needs to be wrapped in double quotes
	if !json.Valid([]byte(s)) {
		s = `"` + s + `"`
	}

	return s
}

func toJsonArray(s string, r interface{}) string {
	// if the value is not already JSON, assume it is a comma separated string
	// note: single values are valid json, so make sure it is already an array as well
	if !json.Valid([]byte(s)) || !strings.HasPrefix(s, "[") {
		t := reflect.TypeOf(r)
		switch getKind(t) {
		case reflect.Bool, reflect.Int64, reflect.Float32, reflect.Float64:
			// booleans and numbers don't need to be wrapped in double quotes
			s = `[` + s + `]`
		default:
			// this should be strings and types like Dates/Datetimes
			// escape existing quotes before wrapping values in quotes
			s = strings.ReplaceAll(s, `"`, `\"`)
			s = `["` + strings.ReplaceAll(s, `,`, `","`) + `"]`
		}
	}

	return s
}

// recursively get the "Kind" from pointers and slices
func getKind(t reflect.Type) reflect.Kind {
	k := t.Kind()
	switch k {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		return getKind(t.Elem())
	default:
		return k
	}
}
