# Distribution List API v1 - Go SDK

## Overview

The Distribution List API v1 is an IBM Cloud Platform Service that enables management of notification distribution lists for IBM Cloud accounts. This API allows you to configure and manage destinations where notifications and alerts can be sent, primarily through IBM Cloud Event Notifications service.

## Purpose

This API provides a centralized way to:
- **Manage notification destinations** for IBM Cloud accounts
- **Configure Event Notifications** integration for automated alerts
- **Test notification delivery** to ensure proper configuration
- **Support enterprise-wide notifications** across multiple accounts within the same enterprise

## Key Features

### 1. Distribution List Management
- Add, retrieve, update, and delete notification destinations
- Support for multiple destination types (currently Event Notifications)
- Maximum of 10 destination entries per destination type
- Cross-account support for enterprise accounts

### 2. Event Notifications Integration
- Seamless integration with IBM Cloud Event Notifications service
- Configure destinations using Event Notifications instance GUIDs
- Support for destinations from different accounts within the same enterprise

### 3. Testing & Validation
- Send test notifications to verify destination configuration
- Validate connectivity and proper setup before production use

## Installation

```bash
go get -u github.com/IBM/platform-services-go-sdk/distributionlistapiv1
```

## Authentication

The SDK uses IBM Cloud IAM authentication. Configure authentication using one of these methods:

### Environment Variables
```bash
export DISTRIBUTION_LIST_API_AUTH_TYPE=iam
export DISTRIBUTION_LIST_API_APIKEY=<your-api-key>
export DISTRIBUTION_LIST_API_URL=<service-url>
```

### Programmatic Configuration
```go
import (
    "github.com/IBM/go-sdk-core/v5/core"
    "github.com/IBM/platform-services-go-sdk/distributionlistapiv1"
)

authenticator := &core.IamAuthenticator{
    ApiKey: "<your-api-key>",
}

service, err := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
    Authenticator: authenticator,
    URL:           "<service-url>",
})
```

## API Methods

### 1. GetAllDestinationEntries
Retrieve all destinations in the distribution list for a specified account.

**Endpoint:** `GET /notification-api/v1/distribution_lists/{account_id}/destinations`

**Example:**
```go
options := service.NewGetAllDestinationEntriesOptions(accountID)
result, response, err := service.GetAllDestinationEntries(options)
if err != nil {
    panic(err)
}
```

### 2. AddDestinationEntry
Add a new destination entry to the distribution list.

**Endpoint:** `POST /notification-api/v1/distribution_lists/{account_id}/destinations`

**Constraints:**
- Maximum 10 destination entries per destination type
- For enterprise accounts, can use Event Notifications from different accounts within the same enterprise

**Example:**
```go
import "github.com/go-openapi/strfmt"

// Create Event Notifications destination
destinationRequest := &distributionlistapiv1.AddDestinationEntryRequestEventNotificationDestination{
    ID:              core.UUIDPtr(strfmt.UUID("your-event-notifications-instance-guid")),
    DestinationType: core.StringPtr(distributionlistapiv1.AddDestinationEntryRequest_DestinationType_EventNotifications),
}

options := service.NewAddDestinationEntryOptions(accountID, destinationRequest)
result, response, err := service.AddDestinationEntry(options)
if err != nil {
    panic(err)
}
```

### 3. GetDestinationEntry
Retrieve a specific destination from the distribution list.

**Endpoint:** `GET /notification-api/v1/distribution_lists/{account_id}/destinations/{destination_id}`

**Example:**
```go
destinationID := strfmt.UUID("destination-guid")
options := service.NewGetDestinationEntryOptions(accountID, &destinationID)
result, response, err := service.GetDestinationEntry(options)
if err != nil {
    panic(err)
}
```

### 4. DeleteDestinationEntry
Remove a destination entry from the distribution list.

**Endpoint:** `DELETE /notification-api/v1/distribution_lists/{account_id}/destinations/{destination_id}`

**Example:**
```go
destinationID := strfmt.UUID("destination-guid")
options := service.NewDeleteDestinationEntryOptions(accountID, &destinationID)
response, err := service.DeleteDestinationEntry(options)
if err != nil {
    panic(err)
}
```

### 5. TestDestinationEntry
Send a test notification to verify destination configuration.

**Endpoint:** `POST /notification-api/v1/distribution_lists/{account_id}/destinations/{destination_id}/test`

**Example:**
```go
destinationID := strfmt.UUID("destination-guid")

testRequest := &distributionlistapiv1.TestDestinationEntryRequestEventNotificationDestination{
    DestinationType: core.StringPtr(distributionlistapiv1.TestDestinationEntryRequest_DestinationType_EventNotifications),
}

options := service.NewTestDestinationEntryOptions(accountID, &destinationID, testRequest)
result, response, err := service.TestDestinationEntry(options)
if err != nil {
    panic(err)
}
```

## Data Models

### Destination Types

Currently supported destination type:
- **event_notifications**: IBM Cloud Event Notifications service

### Key Structures

#### AddDestinationEntryRequest
```go
type AddDestinationEntryRequestEventNotificationDestination struct {
    ID              *strfmt.UUID  // Event Notifications instance GUID
    DestinationType *string       // "event_notifications"
}
```

#### DestinationListItem
```go
type DestinationListItemEventNotificationDestination struct {
    ID              *strfmt.UUID  // Destination GUID
    DestinationType *string       // Destination type
}
```

## Complete Usage Example

```go
package main

import (
    "fmt"
    "github.com/IBM/go-sdk-core/v5/core"
    "github.com/IBM/platform-services-go-sdk/distributionlistapiv1"
    "github.com/go-openapi/strfmt"
)

func main() {
    // Initialize service
    authenticator := &core.IamAuthenticator{
        ApiKey: "your-api-key",
    }

    service, err := distributionlistapiv1.NewDistributionListApiV1(
        &distributionlistapiv1.DistributionListApiV1Options{
            Authenticator: authenticator,
            URL:           "https://api.example.com",
        },
    )
    if err != nil {
        panic(err)
    }

    accountID := "your-account-id"
    eventNotificationsGUID := strfmt.UUID("your-en-instance-guid")

    // Add a destination
    addRequest := &distributionlistapiv1.AddDestinationEntryRequestEventNotificationDestination{
        ID:              &eventNotificationsGUID,
        DestinationType: core.StringPtr("event_notifications"),
    }

    addOptions := service.NewAddDestinationEntryOptions(accountID, addRequest)
    addResult, _, err := service.AddDestinationEntry(addOptions)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Added destination: %v\n", addResult)

    // List all destinations
    listOptions := service.NewGetAllDestinationEntriesOptions(accountID)
    destinations, _, err := service.GetAllDestinationEntries(listOptions)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Total destinations: %d\n", len(destinations))

    // Test the destination
    testRequest := &distributionlistapiv1.TestDestinationEntryRequestEventNotificationDestination{
        DestinationType: core.StringPtr("event_notifications"),
    }
    testOptions := service.NewTestDestinationEntryOptions(accountID, &eventNotificationsGUID, testRequest)
    testResult, _, err := service.TestDestinationEntry(testOptions)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Test result: %v\n", testResult)
}
```

## Error Handling

The SDK uses the standard Go error handling pattern. All API methods return an error as the last return value:

```go
result, response, err := service.GetAllDestinationEntries(options)
if err != nil {
    // Handle error
    fmt.Printf("Error: %v\n", err)
    if response != nil {
        fmt.Printf("Status Code: %d\n", response.StatusCode)
    }
    return
}
```

## Service Configuration

### Setting Service URL
```go
err := service.SetServiceURL("https://custom-url.example.com")
```

### Enable Gzip Compression
```go
service.SetEnableGzipCompression(true)
```

### Enable Retries
```go
service.EnableRetries(3, 30*time.Second)
```

### Custom Headers
```go
headers := http.Header{}
headers.Add("Custom-Header", "value")
service.SetDefaultHeaders(headers)
```

## Testing

Run the unit tests:
```bash
cd distributionlistapiv1
go test -v
```

Run with coverage:
```bash
go test -v -cover
```

## API Version

- **Version:** 1.0.0
- **Generated by:** IBM OpenAPI SDK Code Generator Version 3.108.0

## Related IBM Cloud Services

- **IBM Cloud Event Notifications**: The primary destination type for distribution lists
- **IBM Cloud IAM**: Authentication and authorization
- **IBM Cloud Enterprise**: Multi-account support for enterprise customers

## Support & Documentation

- [IBM Cloud Documentation](https://cloud.ibm.com/docs)
- [IBM Cloud Event Notifications](https://cloud.ibm.com/docs/event-notifications)
- [Go SDK Core Documentation](https://github.com/IBM/go-sdk-core)

## License

This SDK is released under the Apache 2.0 license. See the LICENSE file for more information.

## Contributing

See CONTRIBUTING.md in the repository root for contribution guidelines.