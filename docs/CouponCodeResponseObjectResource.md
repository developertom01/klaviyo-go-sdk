# CouponCodeResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CouponCodeEnum**](CouponCodeEnum.md) |  | 
**Id** | **string** | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with. | 
**Attributes** | [**CouponCodeResponseObjectResourceAttributes**](CouponCodeResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewCouponCodeResponseObjectResource

`func NewCouponCodeResponseObjectResource(type_ CouponCodeEnum, id string, attributes CouponCodeResponseObjectResourceAttributes, links ObjectLinks, ) *CouponCodeResponseObjectResource`

NewCouponCodeResponseObjectResource instantiates a new CouponCodeResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCodeResponseObjectResourceWithDefaults

`func NewCouponCodeResponseObjectResourceWithDefaults() *CouponCodeResponseObjectResource`

NewCouponCodeResponseObjectResourceWithDefaults instantiates a new CouponCodeResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponCodeResponseObjectResource) GetType() CouponCodeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponCodeResponseObjectResource) GetTypeOk() (*CouponCodeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponCodeResponseObjectResource) SetType(v CouponCodeEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CouponCodeResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponCodeResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponCodeResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CouponCodeResponseObjectResource) GetAttributes() CouponCodeResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponCodeResponseObjectResource) GetAttributesOk() (*CouponCodeResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponCodeResponseObjectResource) SetAttributes(v CouponCodeResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *CouponCodeResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CouponCodeResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CouponCodeResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


