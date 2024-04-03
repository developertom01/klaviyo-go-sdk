# CouponResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CouponEnum**](CouponEnum.md) |  | 
**Id** | **string** | The internal id of a Coupon is equivalent to its external id stored within an integration. | 
**Attributes** | [**CouponResponseObjectResourceAttributes**](CouponResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewCouponResponseObjectResource

`func NewCouponResponseObjectResource(type_ CouponEnum, id string, attributes CouponResponseObjectResourceAttributes, links ObjectLinks, ) *CouponResponseObjectResource`

NewCouponResponseObjectResource instantiates a new CouponResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponResponseObjectResourceWithDefaults

`func NewCouponResponseObjectResourceWithDefaults() *CouponResponseObjectResource`

NewCouponResponseObjectResourceWithDefaults instantiates a new CouponResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponResponseObjectResource) GetType() CouponEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponResponseObjectResource) GetTypeOk() (*CouponEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponResponseObjectResource) SetType(v CouponEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CouponResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CouponResponseObjectResource) GetAttributes() CouponResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponResponseObjectResource) GetAttributesOk() (*CouponResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponResponseObjectResource) SetAttributes(v CouponResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *CouponResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CouponResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CouponResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


