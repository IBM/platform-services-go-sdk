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
	"github.com/IBM/go-sdk-core/v5/core"
	translation "github.com/IBM/platform-services-go-sdk/i18n"
	"os"
	"reflect"
)

// Deserialize a single value into the target variable. This is most often used for types like Date and DateTime.
// If the value provided by the user is not already valid JSON, wrap it in double quotes before unmarshalling.
func SingleValue(flag string, flagName string, targetDataType string, result interface{}) (err error, msg string) {
	// convert the value to JSON, if needed
	flag = toJsonValue(flag)
	return deserializeJson([]byte(flag), flagName, targetDataType, result)
}

// Deserialize a list of values into the target variable. If the value provided by the user is not already a valid JSON
// array for the given type, convert it to the proper format before unmarshalling.
func List(flag string, flagName string, targetDataType string, result interface{}) (err error, msg string) {
	// convert the string to a JSON array, if needed
	flag = toJsonArray(flag, result)
	return deserializeJson([]byte(flag), flagName, targetDataType, result)
}

// Deserialize a JSON string value into the target variable. JSON can also be provided through a file if the string
// starts with '@'. The string is assumed to be valid JSON.
func JSON(flag string, flagName string, targetDataType string, result interface{}) (error, string) {
	valueAsBytes, err, msg := getJsonStringAsBytes(flag)
	if err != nil {
		return err, msg
	}

	return deserializeJson(valueAsBytes, flagName, targetDataType, result)
}

// Read the contents of a file into the target variable. The type of the target variable
// should be `io.ReadCloser` and the user-provided value should be a path to a valid file.
func File(flag string, result interface{}) (error, string) {
	file, err := os.Open(flag)

	if err != nil {
		msg := translation.T("file-opening-error", map[string]interface{}{
			"FILENAME": flag,
		})

		return err, msg
	}

	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(file))

	return nil, ""
}

// Deserialize a JSON object string value into the target variable, the type of which is a struct representing a model.
func Model(flag string, flagName string, modelName string, unmarshalFunc core.ModelUnmarshaller, result interface{}) (error, string) {
	return deserializeModel(flag, flagName, modelName, unmarshalFunc, false, result)
}

// Deserialize a JSON array string value into the target variable, the type of which is a struct slice representing a list of models.
func ModelSlice(flag string, flagName string, modelName string, unmarshalFunc core.ModelUnmarshaller, result interface{}) (error, string) {
	return deserializeModel(flag, flagName, modelName, unmarshalFunc, true, result)
}
