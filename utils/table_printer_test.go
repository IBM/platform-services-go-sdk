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

package utils_test

import (
	"github.com/IBM/platform-services-go-sdk/utils"
	"reflect"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/testhelpers/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Test suite
var _ = Describe("Table Printer", func() {

	AfterEach(func() {
		utils.TableHeaderOrder = []string{}
	})

	Describe("CreateTable", func() {
		It("Should return nil if input is nil", func() {
			ui := terminal.NewFakeUI()
			var data *utils.TableData = nil
			table := utils.CreateTable(data, ui.Writer())
			Expect(table).To(BeNil())
		})

		It("Should return nil if there are no headers", func() {
			ui := terminal.NewFakeUI()
			data := new(utils.TableData)
			table := utils.CreateTable(data, ui.Writer())
			Expect(table).To(BeNil())
		})

		It("Should still return table if there are no values", func() {
			ui := terminal.NewFakeUI()
			data := new(utils.TableData)
			data.Headers = []string{"header"}
			table := utils.CreateTable(data, ui.Writer())
			Expect(table).NotTo(BeNil())
		})

		It("Should return table when provided values", func() {
			ui := terminal.NewFakeUI()
			data := new(utils.TableData)
			data.Headers = []string{"header"}
			data.Values = [][]string{
				[]string{"1"},
				[]string{"2"},
			}
			table := utils.CreateTable(data, ui.Writer())
			Expect(table).NotTo(BeNil())
		})
	})

	Describe("FormatTableData", func() {
		It("Should return nil if given a nil value", func() {
			var x interface{} = nil
			data := utils.FormatTableData(x, "")
			Expect(data).To(BeNil())
		})

		It("Should return a table for a slice of string pointers", func() {
			str1 := "foo"
			str2 := "bar"

			// slice value
			x := []*string{&str1, &str2}

			// slice values should be explicitly handled in the formatter
			Expect(reflect.ValueOf(x).Kind()).To(Equal(reflect.Slice))

			data := utils.FormatTableData(x, "")

			// for a single array, there should only be one column
			// i.e. only one header
			Expect(len(data.Headers)).To(Equal(1))
			Expect(data.Headers[0]).To(Equal("values")) // the default column header

			// there should be a row for each value in the array, so two
			// each row should only have one item
			Expect(len(data.Values)).To(Equal(2))
			Expect(len(data.Values[0])).To(Equal(1))
			Expect(len(data.Values[1])).To(Equal(1))

			Expect(data.Values[0][0]).To(Equal("foo"))
			Expect(data.Values[1][0]).To(Equal("bar"))
		})

		It("Should return a table for an array of numbers", func() {
			// array value
			// note: the go sdk only produces integers of type int64
			x := [2]int64{1, 2}

			// array values should be explicitly handled in the formatter
			Expect(reflect.ValueOf(x).Kind()).To(Equal(reflect.Array))

			// pass in a mock jmesquery
			data := utils.FormatTableData(x, "resources[0].list_prop")

			// for a single array, there should only be one column
			// i.e. only one header
			Expect(len(data.Headers)).To(Equal(1))
			Expect(data.Headers[0]).To(Equal("list_prop")) // the last query segment

			// there should be a row for each value in the array, so two
			// each row should only have one item
			Expect(len(data.Values)).To(Equal(2))
			Expect(len(data.Values[0])).To(Equal(1))
			Expect(len(data.Values[1])).To(Equal(1))

			Expect(data.Values[0][0]).To(Equal("1"))
			Expect(data.Values[1][0]).To(Equal("2"))
		})

		It("Should return a table for a list of maps", func() {
			// slice of maps
			x := []map[string]int64{
				map[string]int64{"foo": 1, "bar": 2},
				map[string]int64{"foo": 3, "bar": 4},
			}

			utils.TableHeaderOrder = []string{"foo", "bar"}

			data := utils.FormatTableData(x, "")

			// there should be two columns and two rows
			Expect(len(data.Headers)).To(Equal(2))
			Expect(data.Headers[0]).To(Equal("foo"))
			Expect(data.Headers[1]).To(Equal("bar"))

			Expect(len(data.Values)).To(Equal(2))
			Expect(len(data.Values[0])).To(Equal(2))
			Expect(len(data.Values[1])).To(Equal(2))

			Expect(data.Values[0][0]).To(Equal("1"))
			Expect(data.Values[0][1]).To(Equal("2"))
			Expect(data.Values[1][0]).To(Equal("3"))
			Expect(data.Values[1][1]).To(Equal("4"))
		})

		It("Should include all columns in the final table, even if they aren't present in all maps", func() {
			// slice of maps
			x := []map[string]int64{
				map[string]int64{"foo": 1},
				map[string]int64{"bar": 4},
			}

			utils.TableHeaderOrder = []string{"foo", "bar"}

			data := utils.FormatTableData(x, "")

			// there should be two columns and two rows
			Expect(len(data.Headers)).To(Equal(2))
			Expect(data.Headers[0]).To(Equal("foo"))
			Expect(data.Headers[1]).To(Equal("bar"))

			Expect(len(data.Values)).To(Equal(2))
			Expect(len(data.Values[0])).To(Equal(2))
			Expect(len(data.Values[1])).To(Equal(2))

			Expect(data.Values[0][0]).To(Equal("1"))
			Expect(data.Values[0][1]).To(Equal("-"))
			Expect(data.Values[1][0]).To(Equal("-"))
			Expect(data.Values[1][1]).To(Equal("4"))
		})

		It("Should skip url fields for a list of maps", func() {
			// slice of maps
			x := []map[string]int64{
				map[string]int64{"url": 1, "bar": 2},
				map[string]int64{"foo": 3, "bar": 4},
			}

			data := utils.FormatTableData(x, "")

			// there should be two columns and no headers should be "url"
			Expect(len(data.Headers)).To(Equal(2))
			for _, header := range data.Headers {
				Expect(header).NotTo(Equal("url"))
			}
		})

		It("Should return a table for a map with one map array property", func() {
			// map with one array property
			x := map[string]interface{}{
				"array_values": []map[string]int64{
					map[string]int64{"black": 1, "green": 2},
					map[string]int64{"black": 3, "green": 4},
				},
				"foo": "foo_prop",
				"bar": int64(42),
			}

			utils.TableHeaderOrder = []string{"foo", "bar", "black", "green"}

			data := utils.FormatTableData(x, "")

			// there should be four columns and two rows
			Expect(len(data.Headers)).To(Equal(4))
			Expect(data.Headers[0]).To(Equal("foo"))
			Expect(data.Headers[1]).To(Equal("bar"))
			Expect(data.Headers[2]).To(Equal("black"))
			Expect(data.Headers[3]).To(Equal("green"))

			Expect(len(data.Values)).To(Equal(2))
			Expect(len(data.Values[0])).To(Equal(4))
			Expect(len(data.Values[1])).To(Equal(4))

			Expect(data.Values[0][0]).To(Equal("foo_prop"))
			Expect(data.Values[0][1]).To(Equal("42"))
			Expect(data.Values[0][2]).To(Equal("1"))
			Expect(data.Values[0][3]).To(Equal("2"))

			Expect(data.Values[1][0]).To(Equal("foo_prop"))
			Expect(data.Values[1][1]).To(Equal("42"))
			Expect(data.Values[1][2]).To(Equal("3"))
			Expect(data.Values[1][3]).To(Equal("4"))
		})

		It("Should include all headers for the single array map property", func() {
			// map with one array property
			x := map[string]interface{}{
				"array_values": []map[string]int64{
					map[string]int64{"black": 1},
					map[string]int64{"green": 4},
				},
				"foo": "foo_prop",
				"bar": int64(42),
			}

			utils.TableHeaderOrder = []string{"foo", "bar", "black", "green"}

			data := utils.FormatTableData(x, "")

			// there should be four columns and two rows
			Expect(len(data.Headers)).To(Equal(4))
			Expect(data.Headers[0]).To(Equal("foo"))
			Expect(data.Headers[1]).To(Equal("bar"))
			Expect(data.Headers[2]).To(Equal("black"))
			Expect(data.Headers[3]).To(Equal("green"))

			Expect(len(data.Values)).To(Equal(2))
			Expect(len(data.Values[0])).To(Equal(4))
			Expect(len(data.Values[1])).To(Equal(4))

			Expect(data.Values[0][0]).To(Equal("foo_prop"))
			Expect(data.Values[0][1]).To(Equal("42"))
			Expect(data.Values[0][2]).To(Equal("1"))
			Expect(data.Values[0][3]).To(Equal("-"))

			Expect(data.Values[1][0]).To(Equal("foo_prop"))
			Expect(data.Values[1][1]).To(Equal("42"))
			Expect(data.Values[1][2]).To(Equal("-"))
			Expect(data.Values[1][3]).To(Equal("4"))
		})

		It("Should skip url fields for a single map array property", func() {
			// map with one array property
			x := map[string]interface{}{
				"array_values": []map[string]int64{
					map[string]int64{"url": 1, "green": 2},
					map[string]int64{"black": 3, "green": 4},
				},
				"foo": "foo_prop",
				"bar": int64(42),
			}

			data := utils.FormatTableData(x, "")

			// there should be four columns and no headers should be "url"
			Expect(len(data.Headers)).To(Equal(4))
			for _, header := range data.Headers {
				Expect(header).NotTo(Equal("url"))
			}
		})

		It("Should return a table for a map with one primitive array property", func() {
			// map with one array property
			x := map[string]interface{}{
				"array_values": []int64{1, 2, 3, 4},
				"foo":          "foo_prop",
				"bar":          int64(42),
			}

			utils.TableHeaderOrder = []string{"foo", "bar", "array_values"}

			data := utils.FormatTableData(x, "")

			// there should be three columns and four rows
			Expect(len(data.Headers)).To(Equal(3))
			Expect(data.Headers[0]).To(Equal("foo"))
			Expect(data.Headers[1]).To(Equal("bar"))
			Expect(data.Headers[2]).To(Equal("array_values"))

			Expect(len(data.Values)).To(Equal(4))
			Expect(len(data.Values[0])).To(Equal(3))
			Expect(len(data.Values[1])).To(Equal(3))
			Expect(len(data.Values[2])).To(Equal(3))
			Expect(len(data.Values[3])).To(Equal(3))

			Expect(data.Values[0][0]).To(Equal("foo_prop"))
			Expect(data.Values[0][1]).To(Equal("42"))
			Expect(data.Values[0][2]).To(Equal("1"))

			Expect(data.Values[1][0]).To(Equal("foo_prop"))
			Expect(data.Values[1][1]).To(Equal("42"))
			Expect(data.Values[1][2]).To(Equal("2"))

			Expect(data.Values[2][0]).To(Equal("foo_prop"))
			Expect(data.Values[2][1]).To(Equal("42"))
			Expect(data.Values[2][2]).To(Equal("3"))

			Expect(data.Values[3][0]).To(Equal("foo_prop"))
			Expect(data.Values[3][1]).To(Equal("42"))
			Expect(data.Values[3][2]).To(Equal("4"))
		})

		It("Should skip top-level url field for a map with one array property", func() {
			// map with one array property
			x := map[string]interface{}{
				"array_values": []int64{1, 2, 3, 4},
				"url":          "url_prop",
				"bar":          int64(42),
			}

			data := utils.FormatTableData(x, "")

			// there should be two columns and no headers should be "url"
			Expect(len(data.Headers)).To(Equal(2))
			for _, header := range data.Headers {
				Expect(header).NotTo(Equal("url"))
			}
		})

		It("Should return a table for a map with two array properties", func() {
			// map with two array properties
			x := map[string]interface{}{
				"array_values":      []int64{1, 2, 3, 4},
				"more_array_values": []int64{5, 6},
				"foo":               "foo_prop",
				"bar":               int64(42),
			}

			utils.TableHeaderOrder = []string{"array_values", "more_array_values", "foo", "bar"}

			data := utils.FormatTableData(x, "")

			// there should be two columns because output is transposed
			Expect(len(data.Headers)).To(Equal(2))

			Expect(data.Headers[0]).To(Equal(""))
			Expect(data.Headers[1]).To(Equal(""))

			Expect(len(data.Values)).To(Equal(4))
			Expect(len(data.Values[0])).To(Equal(2))

			Expect(data.Values[0][0]).To(Equal("array_values"))
			Expect(data.Values[0][1]).To(Equal("<Array>"))
			Expect(data.Values[1][0]).To(Equal("more_array_values"))
			Expect(data.Values[1][1]).To(Equal("<Array>"))
			Expect(data.Values[2][0]).To(Equal("foo"))
			Expect(data.Values[2][1]).To(Equal("foo_prop"))
			Expect(data.Values[3][0]).To(Equal("bar"))
			Expect(data.Values[3][1]).To(Equal("42"))
		})

		It("Should return a table for a map with a single, empty map array property", func() {
			// map with one array property that is empty
			x := map[string]interface{}{
				"array_values": []map[string]string{},
				"foo":          "foo_prop",
				"bar":          int64(42),
			}

			utils.TableHeaderOrder = []string{"foo", "bar", "array_values"}

			data := utils.FormatTableData(x, "")

			// there should be two columns and one row
			Expect(len(data.Headers)).To(Equal(3))

			Expect(data.Headers[0]).To(Equal("foo"))
			Expect(data.Headers[1]).To(Equal("bar"))
			Expect(data.Headers[2]).To(Equal("array_values"))

			Expect(len(data.Values)).To(Equal(1))
			Expect(len(data.Values[0])).To(Equal(3))

			Expect(data.Values[0][0]).To(Equal("foo_prop"))
			Expect(data.Values[0][1]).To(Equal("42"))
			Expect(data.Values[0][2]).To(Equal("-"))
		})

		It("Should return a table for a map with a single, empty primitive array property", func() {
			// map with one array property that is empty
			x := map[string]interface{}{
				"array_values": []int{},
				"foo":          "foo_prop",
				"bar":          int64(42),
			}

			utils.TableHeaderOrder = []string{"foo", "bar", "array_values"}

			data := utils.FormatTableData(x, "")

			// there should be two columns and one row
			Expect(len(data.Headers)).To(Equal(3))
			Expect(data.Headers[0]).To(Equal("foo"))
			Expect(data.Headers[1]).To(Equal("bar"))
			Expect(data.Headers[2]).To(Equal("array_values"))

			Expect(len(data.Values)).To(Equal(1))
			Expect(len(data.Values[0])).To(Equal(3))

			Expect(data.Values[0][0]).To(Equal("foo_prop"))
			Expect(data.Values[0][1]).To(Equal("42"))
			Expect(data.Values[0][2]).To(Equal("-"))
		})

		It("Should return a table for a map with no array properties", func() {
			// map with all non-arrays
			x := map[string]interface{}{
				"foo": "foo_prop",
				"bar": int64(42),
			}

			utils.TableHeaderOrder = []string{"foo", "bar"}

			data := utils.FormatTableData(x, "")

			// there should be two columns and two rows
			Expect(len(data.Headers)).To(Equal(2))
			Expect(data.Headers[0]).To(Equal(""))
			Expect(data.Headers[1]).To(Equal(""))

			Expect(len(data.Values)).To(Equal(2))
			Expect(len(data.Values[0])).To(Equal(2))

			Expect(data.Values[0][0]).To(Equal("foo"))
			Expect(data.Values[0][1]).To(Equal("foo_prop"))
			Expect(data.Values[1][0]).To(Equal("bar"))
			Expect(data.Values[1][1]).To(Equal("42"))
		})

		It("Should skip url fields in a flat map", func() {
			// map with all non-arrays
			x := map[string]interface{}{
				"foo": "foo_prop",
				"bar": int64(42),
				"url": "www.ibm.com/skipthisurlitissuperlong",
			}

			data := utils.FormatTableData(x, "")

			// there should be two columns and no headers should be "url"
			Expect(len(data.Headers)).To(Equal(2))
			for _, header := range data.Headers {
				Expect(header).NotTo(Equal("url"))
			}
		})

		It("Should return a table for an empty map", func() {
			x := map[string]interface{}{}
			data := utils.FormatTableData(x, "")

			// there should be no column headers
			// this is how the table printer knows not to print
			// an empty table
			Expect(len(data.Headers)).To(Equal(0))
			Expect(len(data.Values)).To(Equal(1))
		})

		It("Should return a table for a single value", func() {
			x := int64(4)
			data := utils.FormatTableData(x, "resource.number_prop")

			// there should be one column and one row
			Expect(len(data.Headers)).To(Equal(1))
			Expect(data.Headers[0]).To(Equal("number_prop"))

			Expect(len(data.Values)).To(Equal(1))
			Expect(len(data.Values[0])).To(Equal(1))

			Expect(data.Values[0][0]).To(Equal("4"))
		})
	})
})
