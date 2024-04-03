# SMSTrackingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddUtm** | **bool** |  | 
**UtmParams** | [**[]UtmParamInfo**](UtmParamInfo.md) |  | 

## Methods

### NewSMSTrackingOptions

`func NewSMSTrackingOptions(addUtm bool, utmParams []UtmParamInfo, ) *SMSTrackingOptions`

NewSMSTrackingOptions instantiates a new SMSTrackingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSTrackingOptionsWithDefaults

`func NewSMSTrackingOptionsWithDefaults() *SMSTrackingOptions`

NewSMSTrackingOptionsWithDefaults instantiates a new SMSTrackingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddUtm

`func (o *SMSTrackingOptions) GetAddUtm() bool`

GetAddUtm returns the AddUtm field if non-nil, zero value otherwise.

### GetAddUtmOk

`func (o *SMSTrackingOptions) GetAddUtmOk() (*bool, bool)`

GetAddUtmOk returns a tuple with the AddUtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUtm

`func (o *SMSTrackingOptions) SetAddUtm(v bool)`

SetAddUtm sets AddUtm field to given value.


### GetUtmParams

`func (o *SMSTrackingOptions) GetUtmParams() []UtmParamInfo`

GetUtmParams returns the UtmParams field if non-nil, zero value otherwise.

### GetUtmParamsOk

`func (o *SMSTrackingOptions) GetUtmParamsOk() (*[]UtmParamInfo, bool)`

GetUtmParamsOk returns a tuple with the UtmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmParams

`func (o *SMSTrackingOptions) SetUtmParams(v []UtmParamInfo)`

SetUtmParams sets UtmParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


