# ProfileImportJobResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the asynchronous job. | 
**CreatedAt** | **time.Time** | The date and time the job was created in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | 
**TotalCount** | **int32** | The total number of operations to be processed by the job. See &#x60;completed_count&#x60; for the job&#39;s current progress. | 
**CompletedCount** | Pointer to **NullableInt32** | The total number of operations that have been completed by the job. | [optional] [default to 0]
**FailedCount** | Pointer to **NullableInt32** | The total number of operations that have failed as part of the job. | [optional] [default to 0]
**CompletedAt** | Pointer to **NullableTime** | Date and time the job was completed in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | Date and time the job expires in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 
**StartedAt** | Pointer to **NullableTime** | Date and time the job started processing in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 

## Methods

### NewProfileImportJobResponseObjectResourceAttributes

`func NewProfileImportJobResponseObjectResourceAttributes(status string, createdAt time.Time, totalCount int32, ) *ProfileImportJobResponseObjectResourceAttributes`

NewProfileImportJobResponseObjectResourceAttributes instantiates a new ProfileImportJobResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileImportJobResponseObjectResourceAttributesWithDefaults

`func NewProfileImportJobResponseObjectResourceAttributesWithDefaults() *ProfileImportJobResponseObjectResourceAttributes`

NewProfileImportJobResponseObjectResourceAttributesWithDefaults instantiates a new ProfileImportJobResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTotalCount

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetCompletedCount

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetCompletedCount() int32`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetCompletedCountOk() (*int32, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetCompletedCount(v int32)`

SetCompletedCount sets CompletedCount field to given value.

### HasCompletedCount

`func (o *ProfileImportJobResponseObjectResourceAttributes) HasCompletedCount() bool`

HasCompletedCount returns a boolean if a field has been set.

### SetCompletedCountNil

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetCompletedCountNil(b bool)`

 SetCompletedCountNil sets the value for CompletedCount to be an explicit nil

### UnsetCompletedCount
`func (o *ProfileImportJobResponseObjectResourceAttributes) UnsetCompletedCount()`

UnsetCompletedCount ensures that no value is present for CompletedCount, not even an explicit nil
### GetFailedCount

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *ProfileImportJobResponseObjectResourceAttributes) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### SetFailedCountNil

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetFailedCountNil(b bool)`

 SetFailedCountNil sets the value for FailedCount to be an explicit nil

### UnsetFailedCount
`func (o *ProfileImportJobResponseObjectResourceAttributes) UnsetFailedCount()`

UnsetFailedCount ensures that no value is present for FailedCount, not even an explicit nil
### GetCompletedAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *ProfileImportJobResponseObjectResourceAttributes) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetExpiresAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ProfileImportJobResponseObjectResourceAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetStartedAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ProfileImportJobResponseObjectResourceAttributes) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ProfileImportJobResponseObjectResourceAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *ProfileImportJobResponseObjectResourceAttributes) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ProfileImportJobResponseObjectResourceAttributes) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


