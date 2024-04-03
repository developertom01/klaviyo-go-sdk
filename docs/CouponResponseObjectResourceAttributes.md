# CouponResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | This is the id that is stored in an integration such as Shopify or Magento. | 
**Description** | Pointer to **NullableString** | A description of the coupon. | [optional] 

## Methods

### NewCouponResponseObjectResourceAttributes

`func NewCouponResponseObjectResourceAttributes(externalId string, ) *CouponResponseObjectResourceAttributes`

NewCouponResponseObjectResourceAttributes instantiates a new CouponResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponResponseObjectResourceAttributesWithDefaults

`func NewCouponResponseObjectResourceAttributesWithDefaults() *CouponResponseObjectResourceAttributes`

NewCouponResponseObjectResourceAttributesWithDefaults instantiates a new CouponResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CouponResponseObjectResourceAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CouponResponseObjectResourceAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CouponResponseObjectResourceAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetDescription

`func (o *CouponResponseObjectResourceAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CouponResponseObjectResourceAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CouponResponseObjectResourceAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CouponResponseObjectResourceAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CouponResponseObjectResourceAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CouponResponseObjectResourceAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


