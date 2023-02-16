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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

// Test suite
var _ = Describe("Table Printer Utils", func() {

	Describe("GetLastQuerySegment", func() {
		It("Should return default if query is empty", func() {
			seg := utils.GetLastQuerySegment("")
			Expect(seg).To(Equal("values"))
		})

		It("Should return last segment in query, split by period", func() {
			seg := utils.GetLastQuerySegment("foo.bar.resources")
			Expect(seg).To(Equal("resources"))
		})

		It("Should return whole query if it contains no periods", func() {
			seg := utils.GetLastQuerySegment("resources")
			Expect(seg).To(Equal("resources"))
		})
	})

	Describe("HasExactlyOneArrayProperty", func() {
		It("Should return true if map with one array property", func() {
			x := map[string]interface{}{
				"a": []string{"foo", "bar"},
				"b": "one",
				"c": 2,
			}
			Expect(utils.HasExactlyOneArrayProperty(reflect.ValueOf(x))).To(Equal(true))
		})

		It("Should return false if map with two array properties", func() {
			x := map[string][]string{
				"a": []string{"foo", "bar"},
				"b": []string{"one", "two"},
			}
			Expect(utils.HasExactlyOneArrayProperty(reflect.ValueOf(x))).To(Equal(false))
		})

		It("Should return false if map with no array properties", func() {
			x := map[string]int{
				"foo": 1,
				"bar": 2,
			}
			Expect(utils.HasExactlyOneArrayProperty(reflect.ValueOf(x))).To(Equal(false))
		})

		It("Should return false if not a map", func() {
			x := "string"
			Expect(utils.HasExactlyOneArrayProperty(reflect.ValueOf(x))).To(Equal(false))
		})
	})

	Describe("DerefValue", func() {
		It("Should return value if not pointer or interface", func() {
			x := "foo"
			value := reflect.ValueOf(x)
			Expect(value.Kind()).To(Equal(reflect.String))

			Expect(utils.DerefValue(value).Kind()).To(Equal(reflect.String))
		})

		It("Should return element if value is pointer", func() {
			x := "foo"
			value := reflect.ValueOf(&x)
			Expect(value.Kind()).To(Equal(reflect.Ptr))

			Expect(utils.DerefValue(value).Kind()).To(Equal(reflect.String))
		})

		It("Should return element if value is interface", func() {
			x := map[string]interface{}{
				"foo": "bar",
			}
			value := reflect.ValueOf(x["foo"])
			Expect(utils.DerefValue(value).Kind()).To(Equal(reflect.String))
		})
	})

	Describe("GetStringValue", func() {
		It("Should return a hyphen when given a nil value", func() {
			data := reflect.ValueOf(nil)
			str := utils.GetStringValue(data)
			Expect(str).To(Equal("-"))
		})

		It("Should handle a string", func() {
			data := reflect.ValueOf("test")
			str := utils.GetStringValue(data)
			Expect(str).To(Equal("test"))
		})

		It("Should handle a boolean", func() {
			data := reflect.ValueOf(true)
			str := utils.GetStringValue(data)
			Expect(str).To(Equal("true"))
		})

		It("Should handle an int64", func() {
			var x int64 = 42
			data := reflect.ValueOf(x)
			str := utils.GetStringValue(data)
			Expect(str).To(Equal("42"))
		})

		It("Should handle a float32", func() {
			var x float32 = 3.14
			data := reflect.ValueOf(x)
			str := utils.GetStringValue(data)
			Expect(str).To(Equal("3.14"))
		})

		It("Should handle float64", func() {
			var x float64 = 0.3333
			data := reflect.ValueOf(x)
			str := utils.GetStringValue(data)
			Expect(str).To(Equal("0.3333"))
		})

		It("Should handle a map", func() {
			data := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
			str := utils.GetStringValue(data)
			Expect(str).To(Equal("<Nested Object>"))
		})

		It("Should handle a slice with elements", func() {
			data := reflect.ValueOf([]int{1, 2, 3})
			str := utils.GetStringValue(data)
			Expect(str).To(Equal("<Array>"))
		})

		It("Should handle a slice with no elements", func() {
			data := reflect.ValueOf([]string{})
			str := utils.GetStringValue(data)
			Expect(str).To(Equal("-"))
		})

		It("Should return a hyphen for an unsupported type (like struct)", func() {
			type TestStruct struct {
				Foo string
			}

			data := reflect.ValueOf(TestStruct{Foo: "bar"})
			str := utils.GetStringValue(data)
			Expect(str).To(Equal("-"))
		})
	})

	Describe("GetArrayElementType", func() {
		It("Should return the type of element in the array", func() {
			data := []string{"foo", "bar"}
			value := reflect.ValueOf(data)
			kind := utils.GetArrayElementType(value)

			Expect(kind).To(Equal(reflect.String))
		})

		It("Should determine type of array element when wrapped in interface", func() {
			var foo interface{}
			var bar interface{}

			foo = "foo"
			bar = "bar"

			data := []interface{}{foo, bar}

			value := reflect.ValueOf(data)
			kind := utils.GetArrayElementType(value)

			Expect(kind).To(Equal(reflect.String))
		})
	})

	Describe("IsArrayType", func() {
		It("Should return true if kind is a slice", func() {
			x := []string{"foo", "bar"}

			kind := reflect.ValueOf(x).Kind()
			Expect(kind).To(Equal(reflect.Slice))

			Expect(utils.IsArrayType(kind)).To(Equal(true))
		})

		It("Should return true if kind is an array", func() {
			x := [2]int64{1, 2}

			kind := reflect.ValueOf(x).Kind()
			Expect(kind).To(Equal(reflect.Array))

			Expect(utils.IsArrayType(kind)).To(Equal(true))
		})

		It("Should return false if kind is not array or slice", func() {
			x := map[string]string{"foo": "one", "bar": "two"}

			kind := reflect.ValueOf(x).Kind()
			Expect(kind).To(Equal(reflect.Map))

			Expect(utils.IsArrayType(kind)).To(Equal(false))
		})
	})

	Describe("IsValidTableData", func() {
		It("Should return true if table exists and has headers", func() {
			data := new(utils.TableData)
			data.Headers = []string{"header"}
			data.Values = [][]string{[]string{"value"}}

			Expect(utils.IsValidTableData(data)).To(Equal(true))
		})

		It("Should return false if data is nil", func() {
			var data *utils.TableData

			Expect(utils.IsValidTableData(data)).To(Equal(false))
		})

		It("Should return false if data has no headers", func() {
			data := new(utils.TableData)
			data.Values = [][]string{[]string{"value"}}

			Expect(utils.IsValidTableData(data)).To(Equal(false))
		})
	})
})
