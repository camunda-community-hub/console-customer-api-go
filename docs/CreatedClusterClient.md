# CreatedClusterClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**Links** | Pointer to [**CreatedClusterClientLinks**](CreatedClusterClientLinks.md) |  | [optional] 
**Name** | **string** |  | 
**Permissions** | **[]string** |  | 
**Uuid** | **string** |  | 

## Methods

### NewCreatedClusterClient

`func NewCreatedClusterClient(clientId string, clientSecret string, name string, permissions []string, uuid string, ) *CreatedClusterClient`

NewCreatedClusterClient instantiates a new CreatedClusterClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedClusterClientWithDefaults

`func NewCreatedClusterClientWithDefaults() *CreatedClusterClient`

NewCreatedClusterClientWithDefaults instantiates a new CreatedClusterClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CreatedClusterClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreatedClusterClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreatedClusterClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CreatedClusterClient) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreatedClusterClient) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreatedClusterClient) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetLinks

`func (o *CreatedClusterClient) GetLinks() CreatedClusterClientLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreatedClusterClient) GetLinksOk() (*CreatedClusterClientLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreatedClusterClient) SetLinks(v CreatedClusterClientLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreatedClusterClient) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *CreatedClusterClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedClusterClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedClusterClient) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *CreatedClusterClient) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreatedClusterClient) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreatedClusterClient) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetUuid

`func (o *CreatedClusterClient) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CreatedClusterClient) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CreatedClusterClient) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


