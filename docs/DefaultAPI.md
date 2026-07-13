# \DefaultAPI

All URIs are relative to *https://api.cloud.camunda.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateMonitoring**](DefaultAPI.md#ActivateMonitoring) | **Post** /clusters/{clusterUuid}/monitoring | 
[**ActivateSecureConnectivity**](DefaultAPI.md#ActivateSecureConnectivity) | **Post** /clusters/{clusterUuid}/secure-connectivity | 
[**CreateBackup**](DefaultAPI.md#CreateBackup) | **Post** /clusters/{clusterUuid}/backups | 
[**CreateClient**](DefaultAPI.md#CreateClient) | **Post** /clusters/{clusterUuid}/clients | 
[**CreateCluster**](DefaultAPI.md#CreateCluster) | **Post** /clusters | 
[**CreateMonitoringClient**](DefaultAPI.md#CreateMonitoringClient) | **Post** /clusters/{clusterUuid}/monitoring/clients | 
[**CreateSecret**](DefaultAPI.md#CreateSecret) | **Post** /clusters/{clusterUuid}/secrets | 
[**DeactivateMonitoring**](DefaultAPI.md#DeactivateMonitoring) | **Delete** /clusters/{clusterUuid}/monitoring | 
[**DeactivateSecureConnectivity**](DefaultAPI.md#DeactivateSecureConnectivity) | **Delete** /clusters/{clusterUuid}/secure-connectivity | 
[**DeleteBackup**](DefaultAPI.md#DeleteBackup) | **Delete** /clusters/{clusterUuid}/backups/{backupId} | 
[**DeleteClient**](DefaultAPI.md#DeleteClient) | **Delete** /clusters/{clusterUuid}/clients/{clientId} | 
[**DeleteCluster**](DefaultAPI.md#DeleteCluster) | **Delete** /clusters/{clusterUuid} | 
[**DeleteMember**](DefaultAPI.md#DeleteMember) | **Delete** /members/{email} | 
[**DeleteMonitoringClient**](DefaultAPI.md#DeleteMonitoringClient) | **Delete** /clusters/{clusterUuid}/monitoring/clients/{clientUuid} | 
[**DeleteSecret**](DefaultAPI.md#DeleteSecret) | **Delete** /clusters/{clusterUuid}/secrets/{secretName} | 
[**GetBackups**](DefaultAPI.md#GetBackups) | **Get** /clusters/{clusterUuid}/backups | 
[**GetClient**](DefaultAPI.md#GetClient) | **Get** /clusters/{clusterUuid}/clients/{clientId} | 
[**GetClients**](DefaultAPI.md#GetClients) | **Get** /clusters/{clusterUuid}/clients | 
[**GetCluster**](DefaultAPI.md#GetCluster) | **Get** /clusters/{clusterUuid} | 
[**GetClusters**](DefaultAPI.md#GetClusters) | **Get** /clusters | 
[**GetCsv**](DefaultAPI.md#GetCsv) | **Get** /activity/csv | 
[**GetJson**](DefaultAPI.md#GetJson) | **Get** /activity/json | 
[**GetMembers**](DefaultAPI.md#GetMembers) | **Get** /members | 
[**GetMeta**](DefaultAPI.md#GetMeta) | **Get** /meta/ip-ranges | 
[**GetMonitoringClients**](DefaultAPI.md#GetMonitoringClients) | **Get** /clusters/{clusterUuid}/monitoring/clients | 
[**GetParameters**](DefaultAPI.md#GetParameters) | **Get** /clusters/parameters | 
[**GetRestore**](DefaultAPI.md#GetRestore) | **Get** /clusters/{clusterUuid}/backups/{backupId}/restore | 
[**GetSecrets**](DefaultAPI.md#GetSecrets) | **Get** /clusters/{clusterUuid}/secrets | 
[**GetSecureConnectivityStatus**](DefaultAPI.md#GetSecureConnectivityStatus) | **Get** /clusters/{clusterUuid}/secure-connectivity | 
[**RestoreFromBackup**](DefaultAPI.md#RestoreFromBackup) | **Post** /clusters/{clusterUuid}/backups/{backupId}/restore | 
[**RotateMonitoringClientPassword**](DefaultAPI.md#RotateMonitoringClientPassword) | **Post** /clusters/{clusterUuid}/monitoring/clients/{clientUuid}/rotate | 
[**UpdateCluster**](DefaultAPI.md#UpdateCluster) | **Patch** /clusters/{clusterUuid} | 
[**UpdateClusterEncryption**](DefaultAPI.md#UpdateClusterEncryption) | **Put** /clusters/{clusterUuid}/encryption | 
[**UpdateIpAllowlist**](DefaultAPI.md#UpdateIpAllowlist) | **Put** /clusters/{clusterUuid}/ipallowlist | 
[**UpdateIpWhitelist**](DefaultAPI.md#UpdateIpWhitelist) | **Put** /clusters/{clusterUuid}/ipwhitelist | 
[**UpdateMembers**](DefaultAPI.md#UpdateMembers) | **Post** /members/{email} | 
[**UpdateSecret**](DefaultAPI.md#UpdateSecret) | **Put** /clusters/{clusterUuid}/secrets/{secretName} | 
[**UpgradeCluster**](DefaultAPI.md#UpgradeCluster) | **Put** /clusters/{clusterUuid}/upgrade | 
[**Wake**](DefaultAPI.md#Wake) | **Put** /clusters/{clusterUuid}/wake | 



## ActivateMonitoring

> ActivateMonitoring(ctx, clusterUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ActivateMonitoring(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ActivateMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivateSecureConnectivity

> ActivateSecureConnectivity(ctx, clusterUuid).ActivateSecureConnectivityRequest(activateSecureConnectivityRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	activateSecureConnectivityRequest := *openapiclient.NewActivateSecureConnectivityRequest([]string{"AllowedPrincipals_example"}, []string{"AllowedRegions_example"}) // ActivateSecureConnectivityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ActivateSecureConnectivity(context.Background(), clusterUuid).ActivateSecureConnectivityRequest(activateSecureConnectivityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ActivateSecureConnectivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSecureConnectivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activateSecureConnectivityRequest** | [**ActivateSecureConnectivityRequest**](ActivateSecureConnectivityRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBackup

> BackupDto CreateBackup(ctx, clusterUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateBackup(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackup`: BackupDto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupDto**](BackupDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClient

> CreatedClusterClient CreateClient(ctx, clusterUuid).CreateClusterClientBody(createClusterClientBody).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	createClusterClientBody := *openapiclient.NewCreateClusterClientBody("ClientName_example") // CreateClusterClientBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateClient(context.Background(), clusterUuid).CreateClusterClientBody(createClusterClientBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClient`: CreatedClusterClient
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createClusterClientBody** | [**CreateClusterClientBody**](CreateClusterClientBody.md) |  | 

### Return type

[**CreatedClusterClient**](CreatedClusterClient.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCluster

> CreateCluster200Response CreateCluster(ctx).CreateClusterRequest(createClusterRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	createClusterRequest := *openapiclient.NewCreateClusterRequest("ChannelId_example", "GenerationId_example", "Name_example", "PlanTypeId_example", "RegionId_example") // CreateClusterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateCluster(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCluster`: CreateCluster200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClusterRequest** | [**CreateClusterRequest**](CreateClusterRequest.md) |  | 

### Return type

[**CreateCluster200Response**](CreateCluster200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitoringClient

> CreateMonitoringClient200Response CreateMonitoringClient(ctx, clusterUuid).CreateMonitoringClientRequest(createMonitoringClientRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	createMonitoringClientRequest := *openapiclient.NewCreateMonitoringClientRequest("Username_example") // CreateMonitoringClientRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateMonitoringClient(context.Background(), clusterUuid).CreateMonitoringClientRequest(createMonitoringClientRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateMonitoringClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMonitoringClient`: CreateMonitoringClient200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateMonitoringClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitoringClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMonitoringClientRequest** | [**CreateMonitoringClientRequest**](CreateMonitoringClientRequest.md) |  | 

### Return type

[**CreateMonitoringClient200Response**](CreateMonitoringClient200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecret

> CreateSecret(ctx, clusterUuid).CreateSecretBody(createSecretBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	createSecretBody := *openapiclient.NewCreateSecretBody("SecretName_example", "SecretValue_example") // CreateSecretBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.CreateSecret(context.Background(), clusterUuid).CreateSecretBody(createSecretBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSecretBody** | [**CreateSecretBody**](CreateSecretBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateMonitoring

> DeactivateMonitoring(ctx, clusterUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeactivateMonitoring(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeactivateMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateSecureConnectivity

> DeactivateSecureConnectivity(ctx, clusterUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeactivateSecureConnectivity(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeactivateSecureConnectivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateSecureConnectivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackup

> BackupDto DeleteBackup(ctx, clusterUuid, backupId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	backupId := "backupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteBackup(context.Background(), clusterUuid, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackup`: BackupDto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BackupDto**](BackupDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClient

> DeleteClient(ctx, clusterUuid, clientId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteClient(context.Background(), clusterUuid, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> DeleteCluster(ctx, clusterUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteCluster(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMember

> DeleteMember(ctx, email).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	email := "email_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteMember(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitoringClient

> DeleteMonitoringClient(ctx, clusterUuid, clientUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	clientUuid := "clientUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteMonitoringClient(context.Background(), clusterUuid, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteMonitoringClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 
**clientUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitoringClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecret

> DeleteSecret(ctx, clusterUuid, secretName).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	secretName := "secretName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteSecret(context.Background(), clusterUuid, secretName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 
**secretName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackups

> []BackupDto GetBackups(ctx, clusterUuid).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetBackups(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackups`: []BackupDto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BackupDto**](BackupDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClient

> ClusterClientConnectionDetails GetClient(ctx, clusterUuid, clientId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClient(context.Background(), clusterUuid, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClient`: ClusterClientConnectionDetails
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterClientConnectionDetails**](ClusterClientConnectionDetails.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClients

> []ClusterClient GetClients(ctx, clusterUuid).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClients(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClients`: []ClusterClient
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ClusterClient**](ClusterClient.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> Cluster GetCluster(ctx, clusterUuid).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCluster(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCluster`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cluster**](Cluster.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusters

> []Cluster GetClusters(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusters`: []Cluster
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersRequest struct via the builder pattern


### Return type

[**[]Cluster**](Cluster.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCsv

> string GetCsv(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCsv(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCsv`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCsv`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCsvRequest struct via the builder pattern


### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJson

> []AuditDto GetJson(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetJson(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJson`: []AuditDto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsonRequest struct via the builder pattern


### Return type

[**[]AuditDto**](AuditDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembers

> []Member GetMembers(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMembers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMembers`: []Member
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMembers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMembersRequest struct via the builder pattern


### Return type

[**[]Member**](Member.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeta

> MetaDto GetMeta(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMeta(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeta`: MetaDto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMeta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaRequest struct via the builder pattern


### Return type

[**MetaDto**](MetaDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitoringClients

> GetMonitoringClients200Response GetMonitoringClients(ctx, clusterUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMonitoringClients(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMonitoringClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonitoringClients`: GetMonitoringClients200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMonitoringClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMonitoringClients200Response**](GetMonitoringClients200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParameters

> Parameters GetParameters(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetParameters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParameters`: Parameters
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetParameters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetParametersRequest struct via the builder pattern


### Return type

[**Parameters**](Parameters.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestore

> RestoreDto GetRestore(ctx, clusterUuid, backupId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	backupId := "backupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRestore(context.Background(), clusterUuid, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestore`: RestoreDto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestoreDto**](RestoreDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecrets

> map[string]string GetSecrets(ctx, clusterUuid).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSecrets(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecrets`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecureConnectivityStatus

> GetSecureConnectivityStatus200Response GetSecureConnectivityStatus(ctx, clusterUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSecureConnectivityStatus(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSecureConnectivityStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecureConnectivityStatus`: GetSecureConnectivityStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSecureConnectivityStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecureConnectivityStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSecureConnectivityStatus200Response**](GetSecureConnectivityStatus200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreFromBackup

> RestoreDto RestoreFromBackup(ctx, clusterUuid, backupId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	backupId := "backupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RestoreFromBackup(context.Background(), clusterUuid, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RestoreFromBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreFromBackup`: RestoreDto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RestoreFromBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFromBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestoreDto**](RestoreDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateMonitoringClientPassword

> CreateMonitoringClient200Response RotateMonitoringClientPassword(ctx, clusterUuid, clientUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	clientUuid := "clientUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RotateMonitoringClientPassword(context.Background(), clusterUuid, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RotateMonitoringClientPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateMonitoringClientPassword`: CreateMonitoringClient200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RotateMonitoringClientPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 
**clientUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateMonitoringClientPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateMonitoringClient200Response**](CreateMonitoringClient200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> UpdateCluster(ctx, clusterUuid).UpdateClusterBody(updateClusterBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	updateClusterBody := *openapiclient.NewUpdateClusterBody() // UpdateClusterBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateCluster(context.Background(), clusterUuid).UpdateClusterBody(updateClusterBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateClusterBody** | [**UpdateClusterBody**](UpdateClusterBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterEncryption

> Cluster UpdateClusterEncryption(ctx, clusterUuid).UpdateClusterEncryptionBody(updateClusterEncryptionBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	updateClusterEncryptionBody := *openapiclient.NewUpdateClusterEncryptionBody("PrimaryEncryptionKeyId_example") // UpdateClusterEncryptionBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateClusterEncryption(context.Background(), clusterUuid).UpdateClusterEncryptionBody(updateClusterEncryptionBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateClusterEncryption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClusterEncryption`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateClusterEncryption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterEncryptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateClusterEncryptionBody** | [**UpdateClusterEncryptionBody**](UpdateClusterEncryptionBody.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIpAllowlist

> UpdateIpAllowlist(ctx, clusterUuid).IpAllowListBody(ipAllowListBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	ipAllowListBody := *openapiclient.NewIpAllowListBody([]openapiclient.ClusterIpallowlistInner{*openapiclient.NewClusterIpallowlistInner("Description_example", "Ip_example")}) // IpAllowListBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateIpAllowlist(context.Background(), clusterUuid).IpAllowListBody(ipAllowListBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateIpAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipAllowListBody** | [**IpAllowListBody**](IpAllowListBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIpWhitelist

> UpdateIpWhitelist(ctx, clusterUuid).IpWhiteListBody(ipWhiteListBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	ipWhiteListBody := *openapiclient.NewIpWhiteListBody([]openapiclient.ClusterIpallowlistInner{*openapiclient.NewClusterIpallowlistInner("Description_example", "Ip_example")}) // IpWhiteListBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateIpWhitelist(context.Background(), clusterUuid).IpWhiteListBody(ipWhiteListBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateIpWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipWhiteListBody** | [**IpWhiteListBody**](IpWhiteListBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMembers

> UpdateMembers(ctx, email).PostMemberBody(postMemberBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	email := "email_example" // string | 
	postMemberBody := *openapiclient.NewPostMemberBody([]openapiclient.AssignableOrganizationRoleType{openapiclient.AssignableOrganizationRoleType("admin")}) // PostMemberBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateMembers(context.Background(), email).PostMemberBody(postMemberBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postMemberBody** | [**PostMemberBody**](PostMemberBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecret

> UpdateSecret(ctx, clusterUuid, secretName).UpdateSecretBody(updateSecretBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 
	secretName := "secretName_example" // string | 
	updateSecretBody := *openapiclient.NewUpdateSecretBody("SecretValue_example") // UpdateSecretBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.UpdateSecret(context.Background(), clusterUuid, secretName).UpdateSecretBody(updateSecretBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 
**secretName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSecretBody** | [**UpdateSecretBody**](UpdateSecretBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeCluster

> GenerationUpgradeForClusterDto UpgradeCluster(ctx, clusterUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpgradeCluster(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpgradeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeCluster`: GenerationUpgradeForClusterDto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpgradeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenerationUpgradeForClusterDto**](GenerationUpgradeForClusterDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Wake

> Wake(ctx, clusterUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/camunda-community-hub/console-customer-api-go"
)

func main() {
	clusterUuid := "clusterUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.Wake(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Wake``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWakeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

