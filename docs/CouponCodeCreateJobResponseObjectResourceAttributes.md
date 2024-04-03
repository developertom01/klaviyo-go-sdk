# CouponCodeCreateJobResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the asynchronous job. | 
**CreatedAt** | **time.Time** | The date and time the job was created in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | 
**TotalCount** | **int32** | The total number of operations to be processed by the job. See &#x60;completed_count&#x60; for the job&#39;s current progress. | 
**CompletedCount** | Pointer to **NullableInt32** | The total number of operations that have been completed by the job. | [optional] [default to 0]
**FailedCount** | Pointer to **NullableInt32** | The total number of operations that have failed as part of the job. | [optional] [default to 0]
**CompletedAt** | Pointer to **NullableTime** | Date and time the job was completed in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 
**Errors** | Pointer to [**[]APIJobErrorPayload**](APIJobErrorPayload.md) | Array of errors encountered during the processing of the job. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | Date and time the job expires in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 

## Methods

### NewCouponCodeCreateJobResponseObjectResourceAttributes

`func NewCouponCodeCreateJobResponseObjectResourceAttributes(status string, createdAt time.Time, totalCount int32, ) *CouponCodeCreateJobResponseObjectResourceAttributes`

NewCouponCodeCreateJobResponseObjectResourceAttributes instantiates a new CouponCodeCreateJobResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCodeCreateJobResponseObjectResourceAttributesWithDefaults

`func NewCouponCodeCreateJobResponseObjectResourceAttributesWithDefaults() *CouponCodeCreateJobResponseObjectResourceAttributes`

NewCouponCodeCreateJobResponseObjectResourceAttributesWithDefaults instantiates a new CouponCodeCreateJobResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTotalCount

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetCompletedCount

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCompletedCount() int32`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCompletedCountOk() (*int32, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetCompletedCount(v int32)`

SetCompletedCount sets CompletedCount field to given value.

### HasCompletedCount

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) HasCompletedCount() bool`

HasCompletedCount returns a boolean if a field has been set.

### SetCompletedCountNil

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetCompletedCountNil(b bool)`

 SetCompletedCountNil sets the value for CompletedCount to be an explicit nil

### UnsetCompletedCount
`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) UnsetCompletedCount()`

UnsetCompletedCount ensures that no value is present for CompletedCount, not even an explicit nil
### GetFailedCount

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### SetFailedCountNil

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetFailedCountNil(b bool)`

 SetFailedCountNil sets the value for FailedCount to be an explicit nil

### UnsetFailedCount
`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) UnsetFailedCount()`

UnsetFailedCount ensures that no value is present for FailedCount, not even an explicit nil
### GetCompletedAt

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetErrors

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetErrors() []APIJobErrorPayload`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetErrorsOk() (*[]APIJobErrorPayload, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetErrors(v []APIJobErrorPayload)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetExpiresAt

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CouponCodeCreateJobResponseObjectResourceAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


