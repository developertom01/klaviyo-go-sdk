# EmailMarketing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanReceiveEmailMarketing** | **bool** | Whether or not this profile has implicit consent to receive email marketing. True if it does profile does not have any global suppressions. | 
**Consent** | **string** | The consent status for email marketing. | 
**ConsentTimestamp** | Pointer to **NullableTime** | The timestamp when consent was recorded or updated for email marketing, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 
**LastUpdated** | Pointer to **NullableTime** | The timestamp when a field on the email marketing object was last modified. | [optional] 
**Method** | Pointer to **NullableString** | The method by which the profile was subscribed to email marketing. | [optional] 
**MethodDetail** | Pointer to **NullableString** | Additional details about the method by which the profile was subscribed to email marketing. This may be empty if no details were provided. | [optional] [default to ""]
**CustomMethodDetail** | Pointer to **NullableString** | Additional detail provided by the caller when the profile was subscribed. This may be empty if no details were provided. | [optional] 
**DoubleOptin** | Pointer to **NullableBool** | Whether the profile was subscribed to email marketing using a double opt-in. | [optional] 
**Suppression** | Pointer to [**[]EmailMarketingSuppression**](EmailMarketingSuppression.md) | The global email marketing suppression for this profile. | [optional] 
**ListSuppressions** | Pointer to [**[]EmailMarketingListSuppression**](EmailMarketingListSuppression.md) | The list suppressions for this profile. | [optional] 

## Methods

### NewEmailMarketing

`func NewEmailMarketing(canReceiveEmailMarketing bool, consent string, ) *EmailMarketing`

NewEmailMarketing instantiates a new EmailMarketing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMarketingWithDefaults

`func NewEmailMarketingWithDefaults() *EmailMarketing`

NewEmailMarketingWithDefaults instantiates a new EmailMarketing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanReceiveEmailMarketing

`func (o *EmailMarketing) GetCanReceiveEmailMarketing() bool`

GetCanReceiveEmailMarketing returns the CanReceiveEmailMarketing field if non-nil, zero value otherwise.

### GetCanReceiveEmailMarketingOk

`func (o *EmailMarketing) GetCanReceiveEmailMarketingOk() (*bool, bool)`

GetCanReceiveEmailMarketingOk returns a tuple with the CanReceiveEmailMarketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReceiveEmailMarketing

`func (o *EmailMarketing) SetCanReceiveEmailMarketing(v bool)`

SetCanReceiveEmailMarketing sets CanReceiveEmailMarketing field to given value.


### GetConsent

`func (o *EmailMarketing) GetConsent() string`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *EmailMarketing) GetConsentOk() (*string, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *EmailMarketing) SetConsent(v string)`

SetConsent sets Consent field to given value.


### GetConsentTimestamp

`func (o *EmailMarketing) GetConsentTimestamp() time.Time`

GetConsentTimestamp returns the ConsentTimestamp field if non-nil, zero value otherwise.

### GetConsentTimestampOk

`func (o *EmailMarketing) GetConsentTimestampOk() (*time.Time, bool)`

GetConsentTimestampOk returns a tuple with the ConsentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentTimestamp

`func (o *EmailMarketing) SetConsentTimestamp(v time.Time)`

SetConsentTimestamp sets ConsentTimestamp field to given value.

### HasConsentTimestamp

`func (o *EmailMarketing) HasConsentTimestamp() bool`

HasConsentTimestamp returns a boolean if a field has been set.

### SetConsentTimestampNil

`func (o *EmailMarketing) SetConsentTimestampNil(b bool)`

 SetConsentTimestampNil sets the value for ConsentTimestamp to be an explicit nil

### UnsetConsentTimestamp
`func (o *EmailMarketing) UnsetConsentTimestamp()`

UnsetConsentTimestamp ensures that no value is present for ConsentTimestamp, not even an explicit nil
### GetLastUpdated

`func (o *EmailMarketing) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EmailMarketing) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EmailMarketing) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EmailMarketing) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *EmailMarketing) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *EmailMarketing) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetMethod

`func (o *EmailMarketing) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *EmailMarketing) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *EmailMarketing) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *EmailMarketing) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *EmailMarketing) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *EmailMarketing) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetMethodDetail

`func (o *EmailMarketing) GetMethodDetail() string`

GetMethodDetail returns the MethodDetail field if non-nil, zero value otherwise.

### GetMethodDetailOk

`func (o *EmailMarketing) GetMethodDetailOk() (*string, bool)`

GetMethodDetailOk returns a tuple with the MethodDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodDetail

`func (o *EmailMarketing) SetMethodDetail(v string)`

SetMethodDetail sets MethodDetail field to given value.

### HasMethodDetail

`func (o *EmailMarketing) HasMethodDetail() bool`

HasMethodDetail returns a boolean if a field has been set.

### SetMethodDetailNil

`func (o *EmailMarketing) SetMethodDetailNil(b bool)`

 SetMethodDetailNil sets the value for MethodDetail to be an explicit nil

### UnsetMethodDetail
`func (o *EmailMarketing) UnsetMethodDetail()`

UnsetMethodDetail ensures that no value is present for MethodDetail, not even an explicit nil
### GetCustomMethodDetail

`func (o *EmailMarketing) GetCustomMethodDetail() string`

GetCustomMethodDetail returns the CustomMethodDetail field if non-nil, zero value otherwise.

### GetCustomMethodDetailOk

`func (o *EmailMarketing) GetCustomMethodDetailOk() (*string, bool)`

GetCustomMethodDetailOk returns a tuple with the CustomMethodDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMethodDetail

`func (o *EmailMarketing) SetCustomMethodDetail(v string)`

SetCustomMethodDetail sets CustomMethodDetail field to given value.

### HasCustomMethodDetail

`func (o *EmailMarketing) HasCustomMethodDetail() bool`

HasCustomMethodDetail returns a boolean if a field has been set.

### SetCustomMethodDetailNil

`func (o *EmailMarketing) SetCustomMethodDetailNil(b bool)`

 SetCustomMethodDetailNil sets the value for CustomMethodDetail to be an explicit nil

### UnsetCustomMethodDetail
`func (o *EmailMarketing) UnsetCustomMethodDetail()`

UnsetCustomMethodDetail ensures that no value is present for CustomMethodDetail, not even an explicit nil
### GetDoubleOptin

`func (o *EmailMarketing) GetDoubleOptin() bool`

GetDoubleOptin returns the DoubleOptin field if non-nil, zero value otherwise.

### GetDoubleOptinOk

`func (o *EmailMarketing) GetDoubleOptinOk() (*bool, bool)`

GetDoubleOptinOk returns a tuple with the DoubleOptin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleOptin

`func (o *EmailMarketing) SetDoubleOptin(v bool)`

SetDoubleOptin sets DoubleOptin field to given value.

### HasDoubleOptin

`func (o *EmailMarketing) HasDoubleOptin() bool`

HasDoubleOptin returns a boolean if a field has been set.

### SetDoubleOptinNil

`func (o *EmailMarketing) SetDoubleOptinNil(b bool)`

 SetDoubleOptinNil sets the value for DoubleOptin to be an explicit nil

### UnsetDoubleOptin
`func (o *EmailMarketing) UnsetDoubleOptin()`

UnsetDoubleOptin ensures that no value is present for DoubleOptin, not even an explicit nil
### GetSuppression

`func (o *EmailMarketing) GetSuppression() []EmailMarketingSuppression`

GetSuppression returns the Suppression field if non-nil, zero value otherwise.

### GetSuppressionOk

`func (o *EmailMarketing) GetSuppressionOk() (*[]EmailMarketingSuppression, bool)`

GetSuppressionOk returns a tuple with the Suppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppression

`func (o *EmailMarketing) SetSuppression(v []EmailMarketingSuppression)`

SetSuppression sets Suppression field to given value.

### HasSuppression

`func (o *EmailMarketing) HasSuppression() bool`

HasSuppression returns a boolean if a field has been set.

### SetSuppressionNil

`func (o *EmailMarketing) SetSuppressionNil(b bool)`

 SetSuppressionNil sets the value for Suppression to be an explicit nil

### UnsetSuppression
`func (o *EmailMarketing) UnsetSuppression()`

UnsetSuppression ensures that no value is present for Suppression, not even an explicit nil
### GetListSuppressions

`func (o *EmailMarketing) GetListSuppressions() []EmailMarketingListSuppression`

GetListSuppressions returns the ListSuppressions field if non-nil, zero value otherwise.

### GetListSuppressionsOk

`func (o *EmailMarketing) GetListSuppressionsOk() (*[]EmailMarketingListSuppression, bool)`

GetListSuppressionsOk returns a tuple with the ListSuppressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListSuppressions

`func (o *EmailMarketing) SetListSuppressions(v []EmailMarketingListSuppression)`

SetListSuppressions sets ListSuppressions field to given value.

### HasListSuppressions

`func (o *EmailMarketing) HasListSuppressions() bool`

HasListSuppressions returns a boolean if a field has been set.

### SetListSuppressionsNil

`func (o *EmailMarketing) SetListSuppressionsNil(b bool)`

 SetListSuppressionsNil sets the value for ListSuppressions to be an explicit nil

### UnsetListSuppressions
`func (o *EmailMarketing) UnsetListSuppressions()`

UnsetListSuppressions ensures that no value is present for ListSuppressions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


