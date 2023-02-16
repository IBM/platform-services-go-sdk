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

package deserialize_test

import (
	"bytes"
	"encoding/json"
	"github.com/IBM/go-sdk-core/v5/core"
	d "github.com/IBM/platform-services-go-sdk/utils/deserialize"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"os"
	"reflect"
)

// set up a temporary directory to hold test files
var _ = BeforeSuite(func() {
	// create a temp mock file
	dirErr := os.Mkdir("tempdir", 0755)
	if dirErr != nil {
		Fail(dirErr.Error())
	}

	message := []byte(`{"string_prop": "test string", "int_prop": 25}`)
	fileErr := os.WriteFile("tempdir/test.json", message, 0644)
	if fileErr != nil {
		Fail(fileErr.Error())
	}
})

// cleanup test file directory
var _ = AfterSuite(func() {
	err := os.RemoveAll("tempdir")
	if err != nil {
		Fail(err.Error())
	}
})

// Test suite
var _ = Describe("Deserizalize", func() {
	Describe("SingleValue", func() {
		It("Should deserialize single value out of JSON form", func() {
			var target string
			// strings need to be wrapped in double quotes to be parsed as JSON
			// the code should do that for us
			err, msg := d.SingleValue("test", "foo", "string", &target)
			Expect(err).To(BeNil())
			Expect(msg).To(Equal(""))
			Expect(target).To(Equal("test"))
		})

		It("Should deserialize single value already in JSON form", func() {
			var target int64
			// integers need to NOT be wrapped in double quotes
			// the code should not do that in this case
			err, msg := d.SingleValue("10", "foo", "number", &target)
			Expect(err).To(BeNil())
			Expect(msg).To(Equal(""))
			Expect(target).To(Equal(int64(10)))
		})
	})

	Describe("List", func() {
		It("Should deserialize list of strings in CSV form", func() {
			var target []string
			err, msg := d.List("monday,tuesday,wednesday", "foo", "string array", &target)
			Expect(err).To(BeNil())
			Expect(msg).To(Equal(""))
			Expect(target).To(Equal([]string{"monday", "tuesday", "wednesday"}))
		})

		It("Should deserialize list of integers in CSV form", func() {
			var target []int64
			err, msg := d.List("323,207,511", "foo", "number array", &target)
			Expect(err).To(BeNil())
			Expect(msg).To(Equal(""))
			Expect(target).To(Equal([]int64{323, 207, 511}))
		})

		It("Should deserialize list in JSON form", func() {
			var target []string
			err, msg := d.List(`["monday", "tuesday", "wednesday"]`, "foo", "string array", &target)
			Expect(err).To(BeNil())
			Expect(msg).To(Equal(""))
			Expect(target).To(Equal([]string{"monday", "tuesday", "wednesday"}))
		})
	})

	Describe("JSON", func() {
		It("Should deserialize JSON from string", func() {
			var target map[string]interface{}
			value := `{"string_prop": "test string", "int_prop": 25}`
			err, msg := d.JSON(value, "foo", "any object", &target)
			Expect(err).To(BeNil())
			Expect(msg).To(Equal(""))
			Expect(target["string_prop"]).To(Equal("test string"))
		})

		It("Should deserialize JSON from file", func() {
			var target map[string]interface{}

			err, msg := d.JSON("@tempdir/test.json", "foo", "any object", &target)
			Expect(err).To(BeNil())
			Expect(msg).To(Equal(""))
			Expect(target["string_prop"]).To(Equal("test string"))
		})

		It("Should return an error when JSON is invalid", func() {
			var target map[string]interface{}
			value := `{"foo: "bar"}`
			err, msg := d.JSON(value, "bad", "any object", &target)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("invalid character 'b' after object key"))
			Expect(msg).To(Equal("Error parsing flag 'bad' as 'any object'"))
		})
	})

	Describe("File", func() {
		It("Should read the contents of the given filename into the target", func() {
			var target io.ReadCloser
			err, msg := d.File("tempdir/test.json", &target)
			Expect(err).To(BeNil())
			Expect(msg).To(Equal(""))

			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(target)
			Expect(err).To(BeNil())
			Expect(buf.String()).To(Equal(`{"string_prop": "test string", "int_prop": 25}`))
		})
	})

	Describe("Model", func() {
		It("Should deserialize model from JSON string", func() {
			var target *MyModel
			value := `{"foo": "test string", "bar": 25}`
			err, msg := d.Model(
				value,
				"model",
				"MyModel",
				UnmarshalMyModel,
				&target,
			)

			Expect(err).To(BeNil())
			Expect(msg).To(Equal(""))
			Expect(*target.Foo).To(Equal("test string"))
			Expect(*target.Bar).To(Equal(int64(25)))
		})
	})

	Describe("ModelSlice", func() {
		It("Should deserialize model slice from JSON string", func() {
			var target []MyModel
			value := `[{"foo": "test string", "bar": 25}]`
			err, msg := d.ModelSlice(
				value,
				"model",
				"MyModel",
				UnmarshalMyModel,
				&target,
			)

			Expect(err).To(BeNil())
			Expect(msg).To(Equal(""))
			Expect(len(target)).To(Equal(1))

			model := target[0]
			Expect(*model.Foo).To(Equal("test string"))
			Expect(*model.Bar).To(Equal(int64(25)))
		})
	})
})

// for the model tests:
type MyModel struct {
	Foo *string `json:"foo" validate:"required"`
	Bar *int64  `json:"bar,omitempty"`
}

func UnmarshalMyModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MyModel)
	err = core.UnmarshalPrimitive(m, "foo", &obj.Foo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bar", &obj.Bar)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
