# SMSTrackingOptionsSubObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAddUtm** | Pointer to **NullableBool** | Whether the campaign needs UTM parameters. If set to False, UTM params will not be used. | [optional] 
**UtmParams** | Pointer to [**[]UTMParamsSubObject**](UTMParamsSubObject.md) | A list of UTM parameters. If an empty list is given and is_add_utm is True, uses company defaults. | [optional] 

## Methods

### NewSMSTrackingOptionsSubObject

`func NewSMSTrackingOptionsSubObject() *SMSTrackingOptionsSubObject`

NewSMSTrackingOptionsSubObject instantiates a new SMSTrackingOptionsSubObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSTrackingOptionsSubObjectWithDefaults

`func NewSMSTrackingOptionsSubObjectWithDefaults() *SMSTrackingOptionsSubObject`

NewSMSTrackingOptionsSubObjectWithDefaults instantiates a new SMSTrackingOptionsSubObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAddUtm

`func (o *SMSTrackingOptionsSubObject) GetIsAddUtm() bool`

GetIsAddUtm returns the IsAddUtm field if non-nil, zero value otherwise.

### GetIsAddUtmOk

`func (o *SMSTrackingOptionsSubObject) GetIsAddUtmOk() (*bool, bool)`

GetIsAddUtmOk returns a tuple with the IsAddUtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAddUtm

`func (o *SMSTrackingOptionsSubObject) SetIsAddUtm(v bool)`

SetIsAddUtm sets IsAddUtm field to given value.

### HasIsAddUtm

`func (o *SMSTrackingOptionsSubObject) HasIsAddUtm() bool`

HasIsAddUtm returns a boolean if a field has been set.

### SetIsAddUtmNil

`func (o *SMSTrackingOptionsSubObject) SetIsAddUtmNil(b bool)`

 SetIsAddUtmNil sets the value for IsAddUtm to be an explicit nil

### UnsetIsAddUtm
`func (o *SMSTrackingOptionsSubObject) UnsetIsAddUtm()`

UnsetIsAddUtm ensures that no value is present for IsAddUtm, not even an explicit nil
### GetUtmParams

`func (o *SMSTrackingOptionsSubObject) GetUtmParams() []UTMParamsSubObject`

GetUtmParams returns the UtmParams field if non-nil, zero value otherwise.

### GetUtmParamsOk

`func (o *SMSTrackingOptionsSubObject) GetUtmParamsOk() (*[]UTMParamsSubObject, bool)`

GetUtmParamsOk returns a tuple with the UtmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmParams

`func (o *SMSTrackingOptionsSubObject) SetUtmParams(v []UTMParamsSubObject)`

SetUtmParams sets UtmParams field to given value.

### HasUtmParams

`func (o *SMSTrackingOptionsSubObject) HasUtmParams() bool`

HasUtmParams returns a boolean if a field has been set.

### SetUtmParamsNil

`func (o *SMSTrackingOptionsSubObject) SetUtmParamsNil(b bool)`

 SetUtmParamsNil sets the value for UtmParams to be an explicit nil

### UnsetUtmParams
`func (o *SMSTrackingOptionsSubObject) UnsetUtmParams()`

UnsetUtmParams ensures that no value is present for UtmParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


