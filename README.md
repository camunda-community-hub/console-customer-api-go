# Go API client for Camunda Console

Manage Camunda Cloud Clusters and API Clients via API.

* [Official documentation](https://docs.camunda.io/docs/apis-clients/console-api-reference/)
* [Swagger documentation](https://console.cloud.camunda.io/customer-api/openapi/docs/)


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.3.3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/sijoma/console-customer-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://api.cloud.camunda.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClientsApi* | [**CreateClient**](docs/ClientsApi.md#createclient) | **Post** /clusters/{clusterUuid}/clients | 
*ClientsApi* | [**DeleteClient**](docs/ClientsApi.md#deleteclient) | **Delete** /clusters/{clusterUuid}/clients/{clientId} | 
*ClientsApi* | [**GetClient**](docs/ClientsApi.md#getclient) | **Get** /clusters/{clusterUuid}/clients/{clientId} | 
*ClientsApi* | [**GetClients**](docs/ClientsApi.md#getclients) | **Get** /clusters/{clusterUuid}/clients | 
*ClustersApi* | [**CreateCluster**](docs/ClustersApi.md#createcluster) | **Post** /clusters | 
*ClustersApi* | [**DeleteCluster**](docs/ClustersApi.md#deletecluster) | **Delete** /clusters/{clusterUuid} | 
*ClustersApi* | [**GetCluster**](docs/ClustersApi.md#getcluster) | **Get** /clusters/{clusterUuid} | 
*ClustersApi* | [**GetClusters**](docs/ClustersApi.md#getclusters) | **Get** /clusters | 
*ClustersApi* | [**GetParameters**](docs/ClustersApi.md#getparameters) | **Get** /clusters/parameters | 


## Documentation For Models

 - [Cluster](docs/Cluster.md)
 - [ClusterChannel](docs/ClusterChannel.md)
 - [ClusterClient](docs/ClusterClient.md)
 - [ClusterClientConnectionDetails](docs/ClusterClientConnectionDetails.md)
 - [ClusterGeneration](docs/ClusterGeneration.md)
 - [ClusterLinks](docs/ClusterLinks.md)
 - [ClusterPlanType](docs/ClusterPlanType.md)
 - [ClusterRegion](docs/ClusterRegion.md)
 - [ClusterStatus](docs/ClusterStatus.md)
 - [ClusterStatus2](docs/ClusterStatus2.md)
 - [CreateClusterBody](docs/CreateClusterBody.md)
 - [CreateClusterClientBody](docs/CreateClusterClientBody.md)
 - [CreatedClusterClient](docs/CreatedClusterClient.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [Parameters](docs/Parameters.md)
 - [ParametersAllowedGenerations](docs/ParametersAllowedGenerations.md)
 - [ParametersChannels](docs/ParametersChannels.md)
 - [ParametersClusterPlanTypes](docs/ParametersClusterPlanTypes.md)


## Documentation For Authorization



### bearer

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


### oAuthConsole


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **get:clusters**: get clusters
 - **create:clusters**: create clusters
 - **delete:clusters**: delete clusters
 - **get:clients**: get cluster clients
 - **create:clients**: create cluster clients
 - **delete:clients**: delete cluster clients

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



