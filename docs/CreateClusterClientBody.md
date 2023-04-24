# CreateClusterClientBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | **string** |  | 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateClusterClientBody

`func NewCreateClusterClientBody(clientName string, ) *CreateClusterClientBody`

NewCreateClusterClientBody instantiates a new CreateClusterClientBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterClientBodyWithDefaults

`func NewCreateClusterClientBodyWithDefaults() *CreateClusterClientBody`

NewCreateClusterClientBodyWithDefaults instantiates a new CreateClusterClientBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *CreateClusterClientBody) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *CreateClusterClientBody) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *CreateClusterClientBody) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetPermissions

`func (o *CreateClusterClientBody) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateClusterClientBody) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateClusterClientBody) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateClusterClientBody) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


