# EmailMarketingListSuppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListId** | **string** | The ID of list to which the suppression applies. | 
**Reason** | **string** | The reason the profile was suppressed from the list. | 
**Timestamp** | **time.Time** | The timestamp when the profile was suppressed from the list, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | 

## Methods

### NewEmailMarketingListSuppression

`func NewEmailMarketingListSuppression(listId string, reason string, timestamp time.Time, ) *EmailMarketingListSuppression`

NewEmailMarketingListSuppression instantiates a new EmailMarketingListSuppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMarketingListSuppressionWithDefaults

`func NewEmailMarketingListSuppressionWithDefaults() *EmailMarketingListSuppression`

NewEmailMarketingListSuppressionWithDefaults instantiates a new EmailMarketingListSuppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListId

`func (o *EmailMarketingListSuppression) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *EmailMarketingListSuppression) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *EmailMarketingListSuppression) SetListId(v string)`

SetListId sets ListId field to given value.


### GetReason

`func (o *EmailMarketingListSuppression) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailMarketingListSuppression) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailMarketingListSuppression) SetReason(v string)`

SetReason sets Reason field to given value.


### GetTimestamp

`func (o *EmailMarketingListSuppression) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EmailMarketingListSuppression) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EmailMarketingListSuppression) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


