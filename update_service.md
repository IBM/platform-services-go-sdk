# How to update a service
This document describes the steps needed to update a service that is already contained in this SDK project.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i update_service.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Steps to update a service](#steps-to-update-a-service)
  * [Validate the API definition](#validate-the-api-definition)
  * [Create new feature branch](#create-new-feature-branch)
  * [Re-generate the SDK code](#re-generate-the-sdk-code)
  * [Run unit tests](#run-unit-tests)
  * [Inspect new generated SDK code](#inspect-new-generated-sdk-code)
  * [Modify integration tests and examples as needed](#modify-integration-tests-and-examples-as-needed)
  * [Open PR with your changes](#open-pr-with-your-changes)
- [References](#references)

<!-- tocstop -->

## Overview 
It is a good practice to keep the SDK code for a particular service updated so that it is in sync
with the most recent production version of its API definition.
So, when a service's API definition is changed, the SDK code for the service should be updated in
each SDK project in which it exists.
This could be a change such as editorial changes made to operation or property/parameter descriptions, adding
a new parameter to an existing operation or adding one or more new operations.

## Steps to update a service

### Validate the API definition
Prior to re-generating the SDK code for your service, be sure to validate the updated version of the API definition
using the [IBM OpenAPI Validator](https://github.com/IBM/openapi-validator).
Example:
```sh
lint-openapi -s example-service.yaml
```
This command will display a list of errors and warnings found in the API definition
as well as a summary at the end.
It's not required that you fix all errors and warnings before trying to use the SDK generator, but
this step should identify any critical errors that will need to be fixed prior to the generation step.


### Create new feature branch
After validating the API definition, you're ready to generate new SDK code for your service.
However, before you do that, you should probably create a new feature branch in which to deliver your updates:

```sh
cd <project-root>

git checkout -b update-example-service
```


### Re-generate the SDK code
Next, run the [IBM OpenAPI SDK Generator](https://github.ibm.com/CloudEngineering/openapi-sdkgen) to process your API definition and generate new service and unit test code
for the service:
```sh
cd <project-root>

openapi-sdkgen.sh generate -g ibm-go -i example-service.json -o .
```
The generated service and unit test code is written to the service's package directory within the SDK project.


### Run unit tests
After re-generating the service and unit test code, the next step would be
to run the unit tests.
You can run the unit tests for all the services like this:

```
cd <project-root>

make all
```

or you can run the unit tests for your particular service like this:

```
cd <project-root>/exampleservicev1

go test
```

The unit tests should run clean.  If not, then any test failures should be diagnosed and resolved
before proceeding.


### Inspect new generated SDK code
Next, it is recommended that you inspect the differences in the new and previous generated code to
get an overall view of the changes caused by the re-generation step. The changes that you see in the
generated service code should be in line with the API definition changes that have occurred since you last
generated the SDK code.
Example:
```
git diff
```


### Modify integration tests and examples as needed
After ensuring that the unit tests for your service run clean, the next step would be to modify
your service's integration tests and working examples code to reflect the updated version of
your API definition.

At a mininum, the integration tests and examples for your service should run clean after you
re-generate the service and unit test code.

However, there are situations in which you should also update your service's integration tests
and/or working examples, such as when a new parameter is added to an operation or an entirely new
operation is added to the API.  Keep in mind that the integration tests are used to verify that the
generated SDK code interacts correctly with the service implementation, so any non-trivial changes
made to the API definition (and hence the generated service code) should probably result in updates
to the integration tests.

While modifying the integration tests, also consider if you should make any changes to the service's
working examples code.  We want the working examples to provide a good example for users
to follow when writing their own application code which uses your service, so consider whether or not
the examples code should be updated to reflect the changes made to the API.

After modifying integration tests and working examples, you should run them to make sure they run
clean along with the updated service and unit test code.
Example:
```sh
cd <project-root>/exampleservicev1

go test -tags=integration

go test -tags=examples
```

Note: to successfully run the integration tests and working examples code for a particular service,
you'll need to define the configuration properties required for your service.  For help on this,
contact the project maintainer.

### Open PR with your changes
After completing the previous steps to update the service, unit test, integration test, and working examples
code, commit your changes. Example:
```sh

cd <project-root>

git commit -m "feat(Example Service): re-gen service after recent API changes"

git push
```

Finally, open a pull request (PR) and tag the project maintainer for approval.


## References
- [IBM OpenAPI Validator](https://github.com/IBM/openapi-validator)
- [IBM OpenAPI SDK Generator](https://github.ibm.com/CloudEngineering/openapi-sdkgen)
- [Effective Go - The Go Programming Language](https://golang.org/doc/effective_go)
