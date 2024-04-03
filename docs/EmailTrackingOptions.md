# EmailTrackingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddUtm** | **bool** |  | 
**UtmParams** | [**[]UtmParamInfo**](UtmParamInfo.md) |  | 
**IsTrackingOpens** | **bool** |  | 
**IsTrackingClicks** | **bool** |  | 

## Methods

### NewEmailTrackingOptions

`func NewEmailTrackingOptions(addUtm bool, utmParams []UtmParamInfo, isTrackingOpens bool, isTrackingClicks bool, ) *EmailTrackingOptions`

NewEmailTrackingOptions instantiates a new EmailTrackingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTrackingOptionsWithDefaults

`func NewEmailTrackingOptionsWithDefaults() *EmailTrackingOptions`

NewEmailTrackingOptionsWithDefaults instantiates a new EmailTrackingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddUtm

`func (o *EmailTrackingOptions) GetAddUtm() bool`

GetAddUtm returns the AddUtm field if non-nil, zero value otherwise.

### GetAddUtmOk

`func (o *EmailTrackingOptions) GetAddUtmOk() (*bool, bool)`

GetAddUtmOk returns a tuple with the AddUtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUtm

`func (o *EmailTrackingOptions) SetAddUtm(v bool)`

SetAddUtm sets AddUtm field to given value.


### GetUtmParams

`func (o *EmailTrackingOptions) GetUtmParams() []UtmParamInfo`

GetUtmParams returns the UtmParams field if non-nil, zero value otherwise.

### GetUtmParamsOk

`func (o *EmailTrackingOptions) GetUtmParamsOk() (*[]UtmParamInfo, bool)`

GetUtmParamsOk returns a tuple with the UtmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmParams

`func (o *EmailTrackingOptions) SetUtmParams(v []UtmParamInfo)`

SetUtmParams sets UtmParams field to given value.


### GetIsTrackingOpens

`func (o *EmailTrackingOptions) GetIsTrackingOpens() bool`

GetIsTrackingOpens returns the IsTrackingOpens field if non-nil, zero value otherwise.

### GetIsTrackingOpensOk

`func (o *EmailTrackingOptions) GetIsTrackingOpensOk() (*bool, bool)`

GetIsTrackingOpensOk returns a tuple with the IsTrackingOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOpens

`func (o *EmailTrackingOptions) SetIsTrackingOpens(v bool)`

SetIsTrackingOpens sets IsTrackingOpens field to given value.


### GetIsTrackingClicks

`func (o *EmailTrackingOptions) GetIsTrackingClicks() bool`

GetIsTrackingClicks returns the IsTrackingClicks field if non-nil, zero value otherwise.

### GetIsTrackingClicksOk

`func (o *EmailTrackingOptions) GetIsTrackingClicksOk() (*bool, bool)`

GetIsTrackingClicksOk returns a tuple with the IsTrackingClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingClicks

`func (o *EmailTrackingOptions) SetIsTrackingClicks(v bool)`

SetIsTrackingClicks sets IsTrackingClicks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


