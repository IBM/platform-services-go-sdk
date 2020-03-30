<!--
[![Build Status](https://travis-ci.com/IBM/platform-services-go-sdk.svg?token=eW5FVD71iyte6tTby8gr&branch=master)](https://travis-ci.com/IBM/platform-services-go-sdk)
-->

# IBM Cloud Platform Services Go SDK Version 0.0.0

Go client library to interact with various
[IBM Cloud Platform Service APIs](https://cloud.ibm.com/apidocs?category=platform_services).

Disclaimer: this SDK is being released initially as a **pre-release** version.
Changes might occur which impact applications that use this SDK.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
    + [`go get` command](#go-get-command)
    + [Go modules](#go-modules)
    + [`dep` dependency manager](#dep-dependency-manager)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Platform Services Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Global Search](https://cloud.ibm.com/apidocs/search) | globalsearchv2
[Global Tagging](https://cloud.ibm.com/apidocs/tagging) | globaltaggingv1
[IAM Access Groups](https://cloud.ibm.com/apidocs/iam-access-groups) | iamaccessgroupsv2

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one
[here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
The current version of this SDK: 0.0.0

There are a few different ways to download and install the Platform Services Go SDK project for use by your
Go application:

#### `go get` command  
Use this command to download and install the Platform Services Go SDK project to allow your Go application to
use it:

```
go get -u github.com/IBM/platform-services-go-sdk
```

#### Go modules  
If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

#### `dep` dependency manager  
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.com/IBM/platform-services-go-sdk/globalsearchv2"
  version = "0.0.0"

```

then run `dep ensure`.

## Using the SDK
For general SDK usage information, please see
[this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md)

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/platform-services-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

The IBM Cloud Platform Services Go SDK is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
