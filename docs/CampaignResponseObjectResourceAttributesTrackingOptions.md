# CampaignResponseObjectResourceAttributesTrackingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAddUtm** | Pointer to **NullableBool** | Whether the campaign needs UTM parameters. If set to False, UTM params will not be used. | [optional] 
**UtmParams** | Pointer to [**[]UTMParamsSubObject**](UTMParamsSubObject.md) | A list of UTM parameters. If an empty list is given and is_add_utm is True, uses company defaults. | [optional] 
**IsTrackingClicks** | Pointer to **NullableBool** | Whether the campaign is tracking click events. If not specified, uses company defaults. | [optional] 
**IsTrackingOpens** | Pointer to **NullableBool** | Whether the campaign is tracking open events. If not specified, uses company defaults. | [optional] 

## Methods

### NewCampaignResponseObjectResourceAttributesTrackingOptions

`func NewCampaignResponseObjectResourceAttributesTrackingOptions() *CampaignResponseObjectResourceAttributesTrackingOptions`

NewCampaignResponseObjectResourceAttributesTrackingOptions instantiates a new CampaignResponseObjectResourceAttributesTrackingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignResponseObjectResourceAttributesTrackingOptionsWithDefaults

`func NewCampaignResponseObjectResourceAttributesTrackingOptionsWithDefaults() *CampaignResponseObjectResourceAttributesTrackingOptions`

NewCampaignResponseObjectResourceAttributesTrackingOptionsWithDefaults instantiates a new CampaignResponseObjectResourceAttributesTrackingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAddUtm

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) GetIsAddUtm() bool`

GetIsAddUtm returns the IsAddUtm field if non-nil, zero value otherwise.

### GetIsAddUtmOk

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) GetIsAddUtmOk() (*bool, bool)`

GetIsAddUtmOk returns a tuple with the IsAddUtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAddUtm

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) SetIsAddUtm(v bool)`

SetIsAddUtm sets IsAddUtm field to given value.

### HasIsAddUtm

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) HasIsAddUtm() bool`

HasIsAddUtm returns a boolean if a field has been set.

### SetIsAddUtmNil

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) SetIsAddUtmNil(b bool)`

 SetIsAddUtmNil sets the value for IsAddUtm to be an explicit nil

### UnsetIsAddUtm
`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) UnsetIsAddUtm()`

UnsetIsAddUtm ensures that no value is present for IsAddUtm, not even an explicit nil
### GetUtmParams

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) GetUtmParams() []UTMParamsSubObject`

GetUtmParams returns the UtmParams field if non-nil, zero value otherwise.

### GetUtmParamsOk

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) GetUtmParamsOk() (*[]UTMParamsSubObject, bool)`

GetUtmParamsOk returns a tuple with the UtmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmParams

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) SetUtmParams(v []UTMParamsSubObject)`

SetUtmParams sets UtmParams field to given value.

### HasUtmParams

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) HasUtmParams() bool`

HasUtmParams returns a boolean if a field has been set.

### SetUtmParamsNil

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) SetUtmParamsNil(b bool)`

 SetUtmParamsNil sets the value for UtmParams to be an explicit nil

### UnsetUtmParams
`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) UnsetUtmParams()`

UnsetUtmParams ensures that no value is present for UtmParams, not even an explicit nil
### GetIsTrackingClicks

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) GetIsTrackingClicks() bool`

GetIsTrackingClicks returns the IsTrackingClicks field if non-nil, zero value otherwise.

### GetIsTrackingClicksOk

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) GetIsTrackingClicksOk() (*bool, bool)`

GetIsTrackingClicksOk returns a tuple with the IsTrackingClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingClicks

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) SetIsTrackingClicks(v bool)`

SetIsTrackingClicks sets IsTrackingClicks field to given value.

### HasIsTrackingClicks

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) HasIsTrackingClicks() bool`

HasIsTrackingClicks returns a boolean if a field has been set.

### SetIsTrackingClicksNil

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) SetIsTrackingClicksNil(b bool)`

 SetIsTrackingClicksNil sets the value for IsTrackingClicks to be an explicit nil

### UnsetIsTrackingClicks
`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) UnsetIsTrackingClicks()`

UnsetIsTrackingClicks ensures that no value is present for IsTrackingClicks, not even an explicit nil
### GetIsTrackingOpens

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) GetIsTrackingOpens() bool`

GetIsTrackingOpens returns the IsTrackingOpens field if non-nil, zero value otherwise.

### GetIsTrackingOpensOk

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) GetIsTrackingOpensOk() (*bool, bool)`

GetIsTrackingOpensOk returns a tuple with the IsTrackingOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrackingOpens

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) SetIsTrackingOpens(v bool)`

SetIsTrackingOpens sets IsTrackingOpens field to given value.

### HasIsTrackingOpens

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) HasIsTrackingOpens() bool`

HasIsTrackingOpens returns a boolean if a field has been set.

### SetIsTrackingOpensNil

`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) SetIsTrackingOpensNil(b bool)`

 SetIsTrackingOpensNil sets the value for IsTrackingOpens to be an explicit nil

### UnsetIsTrackingOpens
`func (o *CampaignResponseObjectResourceAttributesTrackingOptions) UnsetIsTrackingOpens()`

UnsetIsTrackingOpens ensures that no value is present for IsTrackingOpens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


