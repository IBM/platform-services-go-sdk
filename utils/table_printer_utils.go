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

package utils

import (
	"reflect"
	"strconv"
	"strings"
)

// If there was a jmespath query against the data, this function will
// extract the final segment as it may need to be used as a column header.
// If no query was given, use a default of "values".
func GetLastQuerySegment(query string) string {
	if query == "" {
		return "values"
	}

	queryArr := strings.Split(query, ".")
	return queryArr[len(queryArr)-1]
}

// Returns true if the value is a map that
// has exactly one property that is an array.
func HasExactlyOneArrayProperty(thing reflect.Value) bool {
	// two potential paths here:
	// 1. its a map
	// 2. it is a type that can't have properties, so return false
	numArrayProps := 0
	if thing.Kind() == reflect.Map {
		iter := thing.MapRange()
		for iter.Next() {
			// note that all values in the map are likely interfaces
			// so we need to deref them
			kind := DerefValue(iter.Value()).Kind()
			if kind == reflect.Slice {
				numArrayProps += 1
			}
		}
	}

	return numArrayProps == 1
}

// Return the dereferenced value if a pointer or interface,
// hand the value back if not.
func DerefValue(thing reflect.Value) reflect.Value {
	if thing.Kind() == reflect.Interface {
		// interface elements can be pointers
		return DerefValue(thing.Elem())
	} else if thing.Kind() == reflect.Ptr {
		return thing.Elem()
	} else {
		return thing
	}
}

// Takes the final value that is to be written to the table
// and formats it as a string if possible.
func GetStringValue(thing reflect.Value) string {
	var result string

	// don't bother with invalid values
	if !thing.IsValid() {
		return "-"
	}

	actualValue := thing.Interface()
	switch thing.Kind() {
	case reflect.String:
		result = thing.String()

	case reflect.Bool:
		result = strconv.FormatBool(actualValue.(bool))

	case reflect.Int64:
		result = strconv.FormatInt(actualValue.(int64), 10)

	case reflect.Float32:
		// FormatFloat must take a float64 as its first value, so typecast is needed
		result = strconv.FormatFloat(float64(actualValue.(float32)), 'g', -1, 32)

	case reflect.Float64:
		result = strconv.FormatFloat(actualValue.(float64), 'g', -1, 64)

	case reflect.Map:
		result = "<Nested Object>"

	case reflect.Slice:
		// print something if an array was returned but is hidden
		// to indicate that there is data there
		if thing.Len() > 0 {
			result = "<Array>"
		} else {
			result = "-"
		}

	default:
		// fmt.Println("Type not yet supported: " + thing.Kind().String())
		result = "-"
	}

	return result
}

// for an array value, get the "Kind" of its individual elements
func GetArrayElementType(value reflect.Value) reflect.Kind {
	arrayElementType := value.Type().Elem().Kind()

	if arrayElementType == reflect.Interface {
		// base the underlying types on the first value
		firstInterface := DerefValue(value.Index(0))
		arrayElementType = firstInterface.Kind()
	}

	return arrayElementType
}

// returns true if kind is Slice or Array
// returns false otherwise
func IsArrayType(kind reflect.Kind) bool {
	return kind == reflect.Slice || kind == reflect.Array
}

// validates the data to ensure a table can be printed
// returns false if the table doesn't have sufficient data
func IsValidTableData(data *TableData) bool {
	return data != nil && len(data.Headers) > 0
}
