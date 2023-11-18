# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpdate** | **bool** | The cluster will automatically be updated by Camunda when a new C8 version is released If the cluster you want to create is running on a non-stable Channel passing this property will be ignored! | 
**Channel** | [**ClusterChannel**](ClusterChannel.md) |  | 
**Created** | **time.Time** |  | 
**Generation** | [**ClusterGeneration**](ClusterGeneration.md) |  | 
**Ipwhitelist** | Pointer to [**[]ClusterIpwhitelistInner**](ClusterIpwhitelistInner.md) | the IP Whitelist rules for your cluster - will only be returned if your organization has the feature enabled and the client you are using has the permission to see it. | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Links** | [**ClusterLinks**](ClusterLinks.md) |  | 
**Name** | **string** |  | 
**OwnerId** | **string** |  | 
**PlanType** | [**ClusterPlanType**](ClusterPlanType.md) |  | 
**Region** | [**ClusterRegion**](ClusterRegion.md) |  | 
**Status** | [**ClusterComponentStatus**](ClusterStatus.md) |  | 
**Uuid** | **string** | The ID used in all further API operations referencing your cluster. | 

## Methods

### NewCluster

`func NewCluster(autoUpdate bool, channel ClusterChannel, created time.Time, generation ClusterGeneration, links ClusterLinks, name string, ownerId string, planType ClusterPlanType, region ClusterRegion, status ClusterComponentStatus, uuid string, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpdate

`func (o *Cluster) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *Cluster) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *Cluster) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.


### GetChannel

`func (o *Cluster) GetChannel() ClusterChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Cluster) GetChannelOk() (*ClusterChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Cluster) SetChannel(v ClusterChannel)`

SetChannel sets Channel field to given value.


### GetCreated

`func (o *Cluster) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Cluster) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Cluster) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetGeneration

`func (o *Cluster) GetGeneration() ClusterGeneration`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *Cluster) GetGenerationOk() (*ClusterGeneration, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *Cluster) SetGeneration(v ClusterGeneration)`

SetGeneration sets Generation field to given value.


### GetIpwhitelist

`func (o *Cluster) GetIpwhitelist() []ClusterIpwhitelistInner`

GetIpwhitelist returns the Ipwhitelist field if non-nil, zero value otherwise.

### GetIpwhitelistOk

`func (o *Cluster) GetIpwhitelistOk() (*[]ClusterIpwhitelistInner, bool)`

GetIpwhitelistOk returns a tuple with the Ipwhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpwhitelist

`func (o *Cluster) SetIpwhitelist(v []ClusterIpwhitelistInner)`

SetIpwhitelist sets Ipwhitelist field to given value.

### HasIpwhitelist

`func (o *Cluster) HasIpwhitelist() bool`

HasIpwhitelist returns a boolean if a field has been set.

### GetLabels

`func (o *Cluster) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Cluster) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Cluster) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Cluster) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLinks

`func (o *Cluster) GetLinks() ClusterLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Cluster) GetLinksOk() (*ClusterLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Cluster) SetLinks(v ClusterLinks)`

SetLinks sets Links field to given value.


### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *Cluster) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Cluster) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Cluster) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetPlanType

`func (o *Cluster) GetPlanType() ClusterPlanType`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *Cluster) GetPlanTypeOk() (*ClusterPlanType, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *Cluster) SetPlanType(v ClusterPlanType)`

SetPlanType sets PlanType field to given value.


### GetRegion

`func (o *Cluster) GetRegion() ClusterRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Cluster) GetRegionOk() (*ClusterRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Cluster) SetRegion(v ClusterRegion)`

SetRegion sets Region field to given value.


### GetStatus

`func (o *Cluster) GetStatus() ClusterComponentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*ClusterComponentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v ClusterComponentStatus)`

SetStatus sets Status field to given value.


### GetUuid

`func (o *Cluster) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Cluster) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Cluster) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


