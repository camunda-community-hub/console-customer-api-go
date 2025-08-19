# GenerationUpgradeForClusterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**GenerationUpgradeForClusterDtoCluster**](GenerationUpgradeForClusterDtoCluster.md) |  | 
**NewGeneration** | [**GenerationUpgradeForClusterDtoCluster**](GenerationUpgradeForClusterDtoCluster.md) |  | 
**OldGeneration** | [**GenerationUpgradeForClusterDtoCluster**](GenerationUpgradeForClusterDtoCluster.md) |  | 
**OrgId** | **string** |  | 

## Methods

### NewGenerationUpgradeForClusterDto

`func NewGenerationUpgradeForClusterDto(cluster GenerationUpgradeForClusterDtoCluster, newGeneration GenerationUpgradeForClusterDtoCluster, oldGeneration GenerationUpgradeForClusterDtoCluster, orgId string, ) *GenerationUpgradeForClusterDto`

NewGenerationUpgradeForClusterDto instantiates a new GenerationUpgradeForClusterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerationUpgradeForClusterDtoWithDefaults

`func NewGenerationUpgradeForClusterDtoWithDefaults() *GenerationUpgradeForClusterDto`

NewGenerationUpgradeForClusterDtoWithDefaults instantiates a new GenerationUpgradeForClusterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *GenerationUpgradeForClusterDto) GetCluster() GenerationUpgradeForClusterDtoCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GenerationUpgradeForClusterDto) GetClusterOk() (*GenerationUpgradeForClusterDtoCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GenerationUpgradeForClusterDto) SetCluster(v GenerationUpgradeForClusterDtoCluster)`

SetCluster sets Cluster field to given value.


### GetNewGeneration

`func (o *GenerationUpgradeForClusterDto) GetNewGeneration() GenerationUpgradeForClusterDtoCluster`

GetNewGeneration returns the NewGeneration field if non-nil, zero value otherwise.

### GetNewGenerationOk

`func (o *GenerationUpgradeForClusterDto) GetNewGenerationOk() (*GenerationUpgradeForClusterDtoCluster, bool)`

GetNewGenerationOk returns a tuple with the NewGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewGeneration

`func (o *GenerationUpgradeForClusterDto) SetNewGeneration(v GenerationUpgradeForClusterDtoCluster)`

SetNewGeneration sets NewGeneration field to given value.


### GetOldGeneration

`func (o *GenerationUpgradeForClusterDto) GetOldGeneration() GenerationUpgradeForClusterDtoCluster`

GetOldGeneration returns the OldGeneration field if non-nil, zero value otherwise.

### GetOldGenerationOk

`func (o *GenerationUpgradeForClusterDto) GetOldGenerationOk() (*GenerationUpgradeForClusterDtoCluster, bool)`

GetOldGenerationOk returns a tuple with the OldGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldGeneration

`func (o *GenerationUpgradeForClusterDto) SetOldGeneration(v GenerationUpgradeForClusterDtoCluster)`

SetOldGeneration sets OldGeneration field to given value.


### GetOrgId

`func (o *GenerationUpgradeForClusterDto) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *GenerationUpgradeForClusterDto) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *GenerationUpgradeForClusterDto) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


