# EmailMarketingSuppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** | The reason the profile was suppressed from the list. | 
**Timestamp** | **time.Time** | The timestamp when the profile was suppressed from the list, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | 

## Methods

### NewEmailMarketingSuppression

`func NewEmailMarketingSuppression(reason string, timestamp time.Time, ) *EmailMarketingSuppression`

NewEmailMarketingSuppression instantiates a new EmailMarketingSuppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMarketingSuppressionWithDefaults

`func NewEmailMarketingSuppressionWithDefaults() *EmailMarketingSuppression`

NewEmailMarketingSuppressionWithDefaults instantiates a new EmailMarketingSuppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *EmailMarketingSuppression) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailMarketingSuppression) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailMarketingSuppression) SetReason(v string)`

SetReason sets Reason field to given value.


### GetTimestamp

`func (o *EmailMarketingSuppression) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmailMarketingSuppression) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmailMarketingSuppression) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


