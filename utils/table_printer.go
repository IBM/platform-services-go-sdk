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
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"io"
	"reflect"
	"sort"
	"strings"
)

var lastQuerySegment string

type TableData struct {
	Headers []string
	Values  [][]string
}

func CreateTable(data *TableData, tableWriter io.Writer) terminal.Table {
	if !IsValidTableData(data) {
		return nil
	}

	table := terminal.NewTable(tableWriter, data.Headers)

	for _, row := range data.Values {
		table.Add(row...)
	}

	return table
}

func FormatTableData(result interface{}, jmesQuery string) *TableData {
	var table *TableData

	// get last segment of jmes query in case it needs to be
	// used as a column header
	lastQuerySegment = GetLastQuerySegment(jmesQuery)

	resultValue := DerefValue(reflect.ValueOf(result))

	// if nothing is passed in, there's nothing to do
	if !resultValue.IsValid() {
		return nil
	}

	kind := resultValue.Kind()

	if IsArrayType(kind) {
		// data is in the form of an array
		table = getTableForArray(resultValue)

	} else if HasExactlyOneArrayProperty(resultValue) {
		// data is in the form of a map with a single array property
		// i.e. other properties are not arrays
		table = getTableForMapWithSingleArrayProperty(resultValue)

	} else if resultValue.Kind() == reflect.Map {
		// data is in the form of a map
		table = getTableForMap(resultValue)
	} else {
		// data is almost certainly in the form of a single, primitive value
		table = getTableForSingleValue(resultValue)
	}

	return table
}

// Sorts tableHeaders to match the order of TableHeaderOrder. All headers
// not in TableHeaderOrder are appended to the end in alphabetical order.
func getSortedTableHeaders(tableHeaders []string) []string {
	orderedHeaders := []string{}
	addedSet := make(map[string]bool)

	// sort given tableHeaders so we may perform binary search
	sort.Strings(tableHeaders)
	// add headers from TableHeaderOrder to the result
	for _, header := range TableHeaderOrder {
		index := sort.SearchStrings(tableHeaders, header)
		if index < len(tableHeaders) && tableHeaders[index] == header {
			// header is in the list of given tableHeaders
			orderedHeaders = append(orderedHeaders, header)
			addedSet[header] = true
		}
	}

	// tableHeaders already sorted alphabetically
	// add remaining headers in alphabetical order
	for _, header := range tableHeaders {
		if !addedSet[header] {
			orderedHeaders = append(orderedHeaders, header)
		}
	}

	return orderedHeaders
}

// takes a map of and makes a sorted Table where the keys in the map
// are headers, and the values are the first row.
func makeSingleRowTableFromMap(tableMap map[string]string) *TableData {
	table := new(TableData)

	// sort the keys in tableMap
	tableHeaders := make([]string, 0, len(tableMap))
	for header := range tableMap {
		tableHeaders = append(tableHeaders, header)
	}
	sortedTableHeaders := getSortedTableHeaders(tableHeaders)

	// map has single row
	row := make([]string, 0, len(sortedTableHeaders))
	for _, header := range sortedTableHeaders {
		row = append(row, tableMap[header])
	}

	table.Headers = sortedTableHeaders
	table.Values = [][]string{row}

	return table
}

// for an array of maps, it may be the case that not all maps share the same keys
// this function will iterate through all of the maps to assemble a superset of unique
// keys to use when creating the table
func getAllKeysForMapArray(mapArray reflect.Value) (map[string]reflect.Value, []string) {
	// maintain a superset of all unique keys across the maps
	masterKeyList := make(map[string]reflect.Value)
	// maintain an ordered set of keys to help enforce alignment between headers and values
	orderedKeys := make([]string, 0)

	for i := 0; i < mapArray.Len(); i++ {
		mapElement := DerefValue(mapArray.Index(i))

		iter := mapElement.MapRange()
		for iter.Next() {
			keyAsString := iter.Key().String()

			// check for uniqueness
			_, exists := masterKeyList[keyAsString]
			// url fields clutter the table and dont provide much value, so they are skipped
			if !exists && strings.ToLower(keyAsString) != "url" {
				masterKeyList[keyAsString] = iter.Key()
				orderedKeys = append(orderedKeys, keyAsString)
			}
		}
	}

	return masterKeyList, getSortedTableHeaders(orderedKeys)
}

// rotates a table so that the headers are in the first
// column and the values are in the 2...N columns
// however, right now, this only occurs when there is only
// one row to transpose (for a flat map)
func transposeTable(data *TableData) *TableData {
	if !IsValidTableData(data) {
		return data
	}

	// collect all values in a single array of arrays
	// we will create the new table from this
	newTableValues := make([][]string, 0)

	// move the headers to their new spots - as the first elements
	// of each new array
	for _, header := range data.Headers {
		newRow := make([]string, 0)
		newRow = append(newRow, header)
		newTableValues = append(newTableValues, newRow)
	}

	// for each row in the original table, process the values and move
	// them to the new table
	for _, rowToProcess := range data.Values {
		for i, valueToMove := range rowToProcess {
			newTableValues[i] = append(newTableValues[i], valueToMove)
		}
	}

	// create the new table from newTableValues
	table := new(TableData)

	// the table printing code in the terminal package distinguishes the first row
	// (normally the headers) by putting them in bold. it doesn't make sense to
	// arbitrarily bold the first row, so fill the header values with empty strings
	table.Headers = make([]string, len(newTableValues[0]))
	for i := range table.Headers {
		table.Headers[i] = ""
	}

	// use the newly transposed table as the values
	table.Values = newTableValues

	return table
}

/**********	ARRAY **********/
// These methods deal with data in the form of an array
// by printing each value in its own row in the table.

// Determine the type of the elements in the array and return
// the appropriate table.
func getTableForArray(resultValue reflect.Value) *TableData {
	if resultValue.Len() == 0 {
		return nil
	}

	var table *TableData

	arrayElementType := GetArrayElementType(resultValue)

	switch arrayElementType {
	case reflect.Map:
		table = getTableForMapArray(resultValue)

	default:
		// should be array of primitives
		table = getTableForPrimitivesArray(resultValue)
	}

	return table
}

// The property names, or keys, from all of the maps are used as the
// column headers for the table and the values are used to fill each row.
func getTableForMapArray(resultValue reflect.Value) *TableData {
	table := new(TableData)
	tableHeaders := make([]string, 0)
	tableValues := make([][]string, 0)

	masterKeyList, orderedKeys := getAllKeysForMapArray(resultValue)
	tableHeaders = append(tableHeaders, orderedKeys...)

	// cycle through all of the maps in the array
	for i := 0; i < resultValue.Len(); i++ {
		mapElement := DerefValue(resultValue.Index(i))
		rowValues := make([]string, 0)

		// cycle through the keys and pull out the values
		for _, key := range tableHeaders {
			field := DerefValue(mapElement.MapIndex(masterKeyList[key]))
			rowValues = append(rowValues, GetStringValue(field))
		}

		// add row to table before moving to next map
		tableValues = append(tableValues, rowValues)
	}

	table.Headers = tableHeaders
	table.Values = tableValues

	return table
}

// This logic assumes that the result is a list of primitive types. This typically
// results from the use of a JMESPath query, so the final segment of the query is used
// as the column header.
func getTableForPrimitivesArray(resultValue reflect.Value) *TableData {
	table := new(TableData)
	tableValues := make([][]string, 0)

	for i := 0; i < resultValue.Len(); i++ {
		listElement := DerefValue(resultValue.Index(i))
		tableValues = append(tableValues, []string{GetStringValue(listElement)})
	}

	table.Headers = []string{lastQuerySegment}
	table.Values = tableValues

	return table
}

/**********	MAP WITH SINGLE ARRAY PROPERTY	**********/
// This method deals with data in the form of a map which may
// have a number of properties but has exactly one property that
// is an array. This is the common pattern for list operations.

// This logic extracts all of the map property names to be used as column headers.
// If the array elements are maps, it is assumed that they all have the same schema and
// the property names of the array element maps are also used as column headers. If the
// array elements are not maps, they are assumed to be primitives and the property name
// of the array property is used as a column header. The values for non-array properties
// are repeated in each row, while the array property values change with each row.
// The number of rows will be equal to the amount of items in the single array, unless
// there are no items, in which case the other properties will still be printed in the table
// on one row.
func getTableForMapWithSingleArrayProperty(resultValue reflect.Value) *TableData {
	table := new(TableData)
	tableValues := make([][]string, 0)

	var theArrayKey reflect.Value

	headers := resultValue.MapKeys()

	headerValueMap := make(map[string]string)

	// iterate through the map, pulling out the keys and values of
	// the non-array properties. the keys will be used as table headers
	// and the values will be present in every row of the table
	for _, key := range headers {
		// url fields clutter the table and dont provide much value, so they are skipped
		if strings.ToLower(key.String()) == "url" {
			continue
		}

		field := DerefValue(resultValue.MapIndex(key))

		if IsArrayType(field.Kind()) {
			// save the key to the single array property
			// so it can be accessed later
			// this should only happen once
			theArrayKey = key

		} else {
			headerValueMap[key.String()] = GetStringValue(field)
		}
	}

	sortedTable := makeSingleRowTableFromMap(headerValueMap)
	tableHeaders := sortedTable.Headers
	nonArrayValues := sortedTable.Values[0]

	arrayProperty := DerefValue(resultValue.MapIndex(theArrayKey))

	// handle the case that the array has no elements
	if arrayProperty.Len() == 0 {
		// if the array is empty, add the key as a header and a hyphen
		// as the value to indicate emptiness
		tableHeaders = append(tableHeaders, theArrayKey.String())
		nonArrayValues = append(nonArrayValues, "-")
		tableValues = append(tableValues, nonArrayValues)
	} else {
		arrayElementType := GetArrayElementType(arrayProperty)

		// if the array property is an array of maps, bring up the map
		// keys to be used as additional table headers and create rows
		// from the values
		if arrayElementType == reflect.Map {
			masterKeyList, mapElementKeys := getAllKeysForMapArray(arrayProperty)
			tableHeaders = append(tableHeaders, mapElementKeys...)

			// cycle through all of the maps and create rows using the
			// collected non-array property values and the map values
			for i := 0; i < arrayProperty.Len(); i++ {
				mapElement := DerefValue(arrayProperty.Index(i))
				mapValues := make([]string, 0)
				for _, key := range mapElementKeys {
					mapField := DerefValue(mapElement.MapIndex(masterKeyList[key]))
					mapValues = append(mapValues, GetStringValue(mapField))
				}

				row := append(nonArrayValues, mapValues...)
				tableValues = append(tableValues, row)
			}
		} else {
			// assuming this is a primitive
			tableHeaders = append(tableHeaders, theArrayKey.String())
			var row []string
			for i := 0; i < arrayProperty.Len(); i++ {
				row = append(nonArrayValues, GetStringValue(DerefValue(arrayProperty.Index(i))))
				tableValues = append(tableValues, row)
			}
		}
	}

	table.Headers = tableHeaders
	table.Values = tableValues

	return table
}

/**********	MAP	**********/
// This method deals with data in the form of a single map.
// The keys are used as column headers and the values are printed
// along a single row.

// Iterates through a map and collects the keys and values
// together, ensuring that they are aligned properly.
func getTableForMap(resultValue reflect.Value) *TableData {
	tableMap := make(map[string]string)

	// loop through map fields
	iter := resultValue.MapRange()
	for iter.Next() {
		// url fields clutter the table and dont provide much value, so they are skipped
		if strings.ToLower(iter.Key().String()) != "url" {
			tableMap[iter.Key().String()] = GetStringValue(DerefValue(iter.Value()))
		}
	}

	return transposeTable(makeSingleRowTableFromMap(tableMap))
}

/**********	SINGLE VALUE	**********/
// This method deals with data in the form of a single, primitive value.
// This is assumed to occur because of a JMESPath query, so the final
// segment of the query is used for the header of the single column. The value
// is printed as the single row.

func getTableForSingleValue(resultValue reflect.Value) *TableData {
	table := new(TableData)
	tableValues := make([][]string, 0)

	singleValue := GetStringValue(DerefValue(resultValue))
	tableValues = append(tableValues, []string{singleValue})

	table.Headers = []string{lastQuerySegment}
	table.Values = tableValues

	return table
}
