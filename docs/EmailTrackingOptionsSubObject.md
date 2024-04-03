# EmailTrackingOptionsSubObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAddUtm** | Pointer to **NullableBool** | Whether the campaign needs UTM parameters. If set to False, UTM params will not be used. | [optional] 
**UtmParams** | Pointer to [**[]UTMParamsSubObject**](UTMParamsSubObject.md) | A list of UTM parameters. If an empty list is given and is_add_utm is True, uses company defaults. | [optional] 
**IsTrackingClicks** | Pointer to **NullableBool** | Whether the campaign is tracking click events. If not specified, uses company defaults. | [optional] 
**IsTrackingOpens** | Pointer to **NullableBool** | Whether the campaign is tracking open events. If not specified, uses company defaults. | [optional] 

## Methods

### NewEmailTrackingOptionsSubObject

`func NewEmailTrackingOptionsSubObject() *EmailTrackingOptionsSubObject`

NewEmailTrackingOptionsSubObject instantiates a new EmailTrackingOptionsSubObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTrackingOptionsSubObjectWithDefaults

`func NewEmailTrackingOptionsSubObjectWithDefaults() *EmailTrackingOptionsSubObject`

NewEmailTrackingOptionsSubObjectWithDefaults instantiates a new EmailTrackingOptionsSubObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAddUtm

`func (o *EmailTrackingOptionsSubObject) GetIsAddUtm() bool`

GetIsAddUtm returns the IsAddUtm field if non-nil, zero value otherwise.

### GetIsAddUtmOk

`func (o *EmailTrackingOptionsSubObject) GetIsAddUtmOk() (*bool, bool)`

GetIsAddUtmOk returns a tuple with the IsAddUtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAddUtm

`func (o *EmailTrackingOptionsSubObject) SetIsAddUtm(v bool)`

SetIsAddUtm sets IsAddUtm field to given value.

### HasIsAddUtm

`func (o *EmailTrackingOptionsSubObject) HasIsAddUtm() bool`

HasIsAddUtm returns a boolean if a field has been set.

### SetIsAddUtmNil

`func (o *EmailTrackingOptionsSubObject) SetIsAddUtmNil(b bool)`

 SetIsAddUtmNil sets the value for IsAddUtm to be an explicit nil

### UnsetIsAddUtm
`func (o *EmailTrackingOptionsSubObject) UnsetIsAddUtm()`

UnsetIsAddUtm ensures that no value is present for IsAddUtm, not even an explicit nil
### GetUtmParams

`func (o *EmailTrackingOptionsSubObject) GetUtmParams() []UTMParamsSubObject`

GetUtmParams returns the UtmParams field if non-nil, zero value otherwise.

### GetUtmParamsOk

`func (o *EmailTrackingOptionsSubObject) GetUtmParamsOk() (*[]UTMParamsSubObject, bool)`

GetUtmParamsOk returns a tuple with the UtmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmParams

`func (o *EmailTrackingOptionsSubObject) SetUtmParams(v []UTMParamsSubObject)`

SetUtmParams sets UtmParams field to given value.

### HasUtmParams

`func (o *EmailTrackingOptionsSubObject) HasUtmParams() bool`

HasUtmParams returns a boolean if a field has been set.

### SetUtmParamsNil

`func (o *EmailTrackingOptionsSubObject) SetUtmParamsNil(b bool)`

 SetUtmParamsNil sets the value for UtmParams to be an explicit nil

### UnsetUtmParams
`func (o *EmailTrackingOptionsSubObject) UnsetUtmParams()`

UnsetUtmParams ensures that no value is present for UtmParams, not even an explicit nil
### GetIsTrackingClicks

`func (o *EmailTrackingOptionsSubObject) GetIsTrackingClicks() bool`

GetIsTrackingClicks returns the IsTrackingClicks field if non-nil, zero value otherwise.

### GetIsTrackingClicksOk

`func (o *EmailTrackingOptionsSubObject) GetIsTrackingClicksOk() (*bool, bool)`

GetIsTrackingClicksOk returns a tuple with the IsTrackingClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingClicks

`func (o *EmailTrackingOptionsSubObject) SetIsTrackingClicks(v bool)`

SetIsTrackingClicks sets IsTrackingClicks field to given value.

### HasIsTrackingClicks

`func (o *EmailTrackingOptionsSubObject) HasIsTrackingClicks() bool`

HasIsTrackingClicks returns a boolean if a field has been set.

### SetIsTrackingClicksNil

`func (o *EmailTrackingOptionsSubObject) SetIsTrackingClicksNil(b bool)`

 SetIsTrackingClicksNil sets the value for IsTrackingClicks to be an explicit nil

### UnsetIsTrackingClicks
`func (o *EmailTrackingOptionsSubObject) UnsetIsTrackingClicks()`

UnsetIsTrackingClicks ensures that no value is present for IsTrackingClicks, not even an explicit nil
### GetIsTrackingOpens

`func (o *EmailTrackingOptionsSubObject) GetIsTrackingOpens() bool`

GetIsTrackingOpens returns the IsTrackingOpens field if non-nil, zero value otherwise.

### GetIsTrackingOpensOk

`func (o *EmailTrackingOptionsSubObject) GetIsTrackingOpensOk() (*bool, bool)`

GetIsTrackingOpensOk returns a tuple with the IsTrackingOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOpens

`func (o *EmailTrackingOptionsSubObject) SetIsTrackingOpens(v bool)`

SetIsTrackingOpens sets IsTrackingOpens field to given value.

### HasIsTrackingOpens

`func (o *EmailTrackingOptionsSubObject) HasIsTrackingOpens() bool`

HasIsTrackingOpens returns a boolean if a field has been set.

### SetIsTrackingOpensNil

`func (o *EmailTrackingOptionsSubObject) SetIsTrackingOpensNil(b bool)`

 SetIsTrackingOpensNil sets the value for IsTrackingOpens to be an explicit nil

### UnsetIsTrackingOpens
`func (o *EmailTrackingOptionsSubObject) UnsetIsTrackingOpens()`

UnsetIsTrackingOpens ensures that no value is present for IsTrackingOpens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


