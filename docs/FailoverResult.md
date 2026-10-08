# FailoverResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewClusterId** | **string** |  | 
**Results** | [**FailoverResultResults**](FailoverResultResults.md) |  | 

## Methods

### NewFailoverResult

`func NewFailoverResult(newClusterId string, results FailoverResultResults, ) *FailoverResult`

NewFailoverResult instantiates a new FailoverResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverResultWithDefaults

`func NewFailoverResultWithDefaults() *FailoverResult`

NewFailoverResultWithDefaults instantiates a new FailoverResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewClusterId

`func (o *FailoverResult) GetNewClusterId() string`

GetNewClusterId returns the NewClusterId field if non-nil, zero value otherwise.

### GetNewClusterIdOk

`func (o *FailoverResult) GetNewClusterIdOk() (*string, bool)`

GetNewClusterIdOk returns a tuple with the NewClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClusterId

`func (o *FailoverResult) SetNewClusterId(v string)`

SetNewClusterId sets NewClusterId field to given value.


### GetResults

`func (o *FailoverResult) GetResults() FailoverResultResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FailoverResult) GetResultsOk() (*FailoverResultResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FailoverResult) SetResults(v FailoverResultResults)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


