# PostCouponResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CouponEnum**](CouponEnum.md) |  | 
**Id** | **string** | The internal id of a Coupon is equivalent to its external id stored within an integration. | 
**Attributes** | [**CouponResponseObjectResourceAttributes**](CouponResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostCouponResponseData

`func NewPostCouponResponseData(type_ CouponEnum, id string, attributes CouponResponseObjectResourceAttributes, links ObjectLinks, ) *PostCouponResponseData`

NewPostCouponResponseData instantiates a new PostCouponResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCouponResponseDataWithDefaults

`func NewPostCouponResponseDataWithDefaults() *PostCouponResponseData`

NewPostCouponResponseDataWithDefaults instantiates a new PostCouponResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostCouponResponseData) GetType() CouponEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCouponResponseData) GetTypeOk() (*CouponEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCouponResponseData) SetType(v CouponEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostCouponResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCouponResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCouponResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PostCouponResponseData) GetAttributes() CouponResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostCouponResponseData) GetAttributesOk() (*CouponResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostCouponResponseData) SetAttributes(v CouponResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *PostCouponResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostCouponResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostCouponResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


