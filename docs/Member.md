# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**Roles** | [**[]OrganizationRole**](OrganizationRole.md) |  | 
**InvitePending** | **bool** |  | 

## Methods

### NewMember

`func NewMember(name string, email string, roles []OrganizationRole, invitePending bool, ) *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Member) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Member) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Member) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *Member) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Member) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Member) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoles

`func (o *Member) GetRoles() []OrganizationRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Member) GetRolesOk() (*[]OrganizationRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Member) SetRoles(v []OrganizationRole)`

SetRoles sets Roles field to given value.


### GetInvitePending

`func (o *Member) GetInvitePending() bool`

GetInvitePending returns the InvitePending field if non-nil, zero value otherwise.

### GetInvitePendingOk

`func (o *Member) GetInvitePendingOk() (*bool, bool)`

GetInvitePendingOk returns a tuple with the InvitePending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitePending

`func (o *Member) SetInvitePending(v bool)`

SetInvitePending sets InvitePending field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


