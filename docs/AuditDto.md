# AuditDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | **string** |  | 
**AuditType** | **string** |  | 
**Entity** | **string** |  | 
**EntityAttribute** | Pointer to **string** |  | [optional] 
**EntityAttributeValueFrom** | Pointer to **string** |  | [optional] 
**EntityAttributeValueTo** | Pointer to **string** |  | [optional] 
**EntityId** | **string** |  | 
**OrgId** | **string** |  | 
**ParentEntity** | Pointer to **string** |  | [optional] 
**ParentEntityId** | Pointer to **string** |  | [optional] 
**Service** | **string** |  | 
**Timestamp** | Pointer to **float64** |  | [optional] 
**UserId** | **string** |  | 

## Methods

### NewAuditDto

`func NewAuditDto(audit string, auditType string, entity string, entityId string, orgId string, service string, userId string, ) *AuditDto`

NewAuditDto instantiates a new AuditDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditDtoWithDefaults

`func NewAuditDtoWithDefaults() *AuditDto`

NewAuditDtoWithDefaults instantiates a new AuditDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *AuditDto) GetAudit() string`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *AuditDto) GetAuditOk() (*string, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *AuditDto) SetAudit(v string)`

SetAudit sets Audit field to given value.


### GetAuditType

`func (o *AuditDto) GetAuditType() string`

GetAuditType returns the AuditType field if non-nil, zero value otherwise.

### GetAuditTypeOk

`func (o *AuditDto) GetAuditTypeOk() (*string, bool)`

GetAuditTypeOk returns a tuple with the AuditType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditType

`func (o *AuditDto) SetAuditType(v string)`

SetAuditType sets AuditType field to given value.


### GetEntity

`func (o *AuditDto) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *AuditDto) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *AuditDto) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetEntityAttribute

`func (o *AuditDto) GetEntityAttribute() string`

GetEntityAttribute returns the EntityAttribute field if non-nil, zero value otherwise.

### GetEntityAttributeOk

`func (o *AuditDto) GetEntityAttributeOk() (*string, bool)`

GetEntityAttributeOk returns a tuple with the EntityAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAttribute

`func (o *AuditDto) SetEntityAttribute(v string)`

SetEntityAttribute sets EntityAttribute field to given value.

### HasEntityAttribute

`func (o *AuditDto) HasEntityAttribute() bool`

HasEntityAttribute returns a boolean if a field has been set.

### GetEntityAttributeValueFrom

`func (o *AuditDto) GetEntityAttributeValueFrom() string`

GetEntityAttributeValueFrom returns the EntityAttributeValueFrom field if non-nil, zero value otherwise.

### GetEntityAttributeValueFromOk

`func (o *AuditDto) GetEntityAttributeValueFromOk() (*string, bool)`

GetEntityAttributeValueFromOk returns a tuple with the EntityAttributeValueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAttributeValueFrom

`func (o *AuditDto) SetEntityAttributeValueFrom(v string)`

SetEntityAttributeValueFrom sets EntityAttributeValueFrom field to given value.

### HasEntityAttributeValueFrom

`func (o *AuditDto) HasEntityAttributeValueFrom() bool`

HasEntityAttributeValueFrom returns a boolean if a field has been set.

### GetEntityAttributeValueTo

`func (o *AuditDto) GetEntityAttributeValueTo() string`

GetEntityAttributeValueTo returns the EntityAttributeValueTo field if non-nil, zero value otherwise.

### GetEntityAttributeValueToOk

`func (o *AuditDto) GetEntityAttributeValueToOk() (*string, bool)`

GetEntityAttributeValueToOk returns a tuple with the EntityAttributeValueTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAttributeValueTo

`func (o *AuditDto) SetEntityAttributeValueTo(v string)`

SetEntityAttributeValueTo sets EntityAttributeValueTo field to given value.

### HasEntityAttributeValueTo

`func (o *AuditDto) HasEntityAttributeValueTo() bool`

HasEntityAttributeValueTo returns a boolean if a field has been set.

### GetEntityId

`func (o *AuditDto) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AuditDto) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AuditDto) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetOrgId

`func (o *AuditDto) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuditDto) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuditDto) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetParentEntity

`func (o *AuditDto) GetParentEntity() string`

GetParentEntity returns the ParentEntity field if non-nil, zero value otherwise.

### GetParentEntityOk

`func (o *AuditDto) GetParentEntityOk() (*string, bool)`

GetParentEntityOk returns a tuple with the ParentEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEntity

`func (o *AuditDto) SetParentEntity(v string)`

SetParentEntity sets ParentEntity field to given value.

### HasParentEntity

`func (o *AuditDto) HasParentEntity() bool`

HasParentEntity returns a boolean if a field has been set.

### GetParentEntityId

`func (o *AuditDto) GetParentEntityId() string`

GetParentEntityId returns the ParentEntityId field if non-nil, zero value otherwise.

### GetParentEntityIdOk

`func (o *AuditDto) GetParentEntityIdOk() (*string, bool)`

GetParentEntityIdOk returns a tuple with the ParentEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEntityId

`func (o *AuditDto) SetParentEntityId(v string)`

SetParentEntityId sets ParentEntityId field to given value.

### HasParentEntityId

`func (o *AuditDto) HasParentEntityId() bool`

HasParentEntityId returns a boolean if a field has been set.

### GetService

`func (o *AuditDto) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AuditDto) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AuditDto) SetService(v string)`

SetService sets Service field to given value.


### GetTimestamp

`func (o *AuditDto) GetTimestamp() float64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditDto) GetTimestampOk() (*float64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditDto) SetTimestamp(v float64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AuditDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUserId

`func (o *AuditDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditDto) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


