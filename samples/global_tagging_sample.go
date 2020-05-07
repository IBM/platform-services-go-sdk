/**
 * (C) Copyright IBM Corp. 2020.
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

// Sample file to demonstrate how to use the Global Tagging V1 SDK.

package samples

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
)

// Create a GlobalTaggingV1 service to be able to call global tagging functions
var service *globaltaggingv1.GlobalTaggingV1
var serviceErr error

func init() {

	// Initialization

	// Load environment variables (should be stored in .env file)
	envVarErr := godotenv.Load("../.env")
	if envVarErr != nil {
		log.Fatal("Error loading .env file")
	}

	// Create the IAM authenticator.
	authenticator := &core.IamAuthenticator{
	    ApiKey: os.Getenv("IAMAPIKEY"),
	}

	service, serviceErr = globaltaggingv1.
		NewGlobalTaggingV1(&globaltaggingv1.GlobalTaggingV1Options{
			Authenticator: authenticator,
		})

	if serviceErr != nil {
		fmt.Println("Error: " + serviceErr.Error())
	}

	// Call any of the functions below to use them.

	// getTagsSample()
	// attachTagsSample()
	// detachTagsSample()
	// deleteTagsSample()

}

func GetTagsSample() {

	// List tags in alphabetical ascending ('asc') or descending ('desc') order - default is ascending
	orderByName := "asc"

	// Create a ListTagsOptions object that serves as the ListTags function parameters.
	getTagsOptions := globaltaggingv1.ListTagsOptions{

		// Specify whether you would like to retrieve GhoST tags, IMS tags, or both.
		Providers:   []string{"ghost", "ims"},
		OrderByName: &orderByName,
	}

	// Function call to ListTags using the ListTagsOptions struct generated earlier
	_, detailedResponse, _ := service.ListTags(&getTagsOptions)

	fmt.Println(detailedResponse)

}

func AttachTagsSample() {

	// Set the CRN of the object to be tagged here.
	crn := "crn-here"

	// The 'resource' is the instance in an IBM account that is to be tagged.
	resource := globaltaggingv1.Resource{
		ResourceID: &crn,
	}

	options := globaltaggingv1.AttachTagOptions{
		Resources: []globaltaggingv1.Resource{resource},

		// Insert the desired tag names here.
		TagNames: []string{"sample", "tag-names", "here"},
	}

	// Call the AttachTag method.
	_, detailedResponse, _ := service.AttachTag(&options)

	fmt.Println(detailedResponse)

}

func DetachTagsSample() {

	// Set the CRN of the object whose tag(s) are being detached here.
	crn := "crn-here"

	// The 'resource' is the instance in an IBM account that is to have its tag removed.
	resource := globaltaggingv1.Resource{
		ResourceID: &crn,
	}

	options := globaltaggingv1.DetachTagOptions{
		Resources: []globaltaggingv1.Resource{resource},

		// Insert the names of the tags to be removed here.
		TagNames: []string{"sample", "tag-names", "here"},
	}

	// Call the AttachTag method.
	_, detailedResponse, _ := service.DetachTag(&options)

	fmt.Println(detailedResponse)

}

func DeleteTagsSample() {

	// Tags can only be deleted if no instance has that tag attached to it

	// The name of the tag to be deleted from an IBM Cloud account.
	tagName := "d-wade"

	// Necessary to populate the header with Content-Type for a delete request.
	headers := make(map[string]string)
	headers["Content-Type"] = "application-json"

	delTagsOptions := globaltaggingv1.DeleteTagOptions{
		TagName: &tagName,
		// Tells the SDK where to look (i.e. GhoST, IMS, or both)
		Providers: []string{"ghost", "ims"},
		Headers:   headers,
	}

	_, detailedResponse, _ := service.DeleteTag(&delTagsOptions)

	fmt.Println(detailedResponse)

}
