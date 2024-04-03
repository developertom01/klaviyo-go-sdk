# CouponCodeUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CouponCodeEnum**](CouponCodeEnum.md) |  | 
**Id** | **string** | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with. | 
**Attributes** | [**CouponCodeUpdateQueryResourceObjectAttributes**](CouponCodeUpdateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewCouponCodeUpdateQueryResourceObject

`func NewCouponCodeUpdateQueryResourceObject(type_ CouponCodeEnum, id string, attributes CouponCodeUpdateQueryResourceObjectAttributes, ) *CouponCodeUpdateQueryResourceObject`

NewCouponCodeUpdateQueryResourceObject instantiates a new CouponCodeUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCodeUpdateQueryResourceObjectWithDefaults

`func NewCouponCodeUpdateQueryResourceObjectWithDefaults() *CouponCodeUpdateQueryResourceObject`

NewCouponCodeUpdateQueryResourceObjectWithDefaults instantiates a new CouponCodeUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponCodeUpdateQueryResourceObject) GetType() CouponCodeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponCodeUpdateQueryResourceObject) GetTypeOk() (*CouponCodeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponCodeUpdateQueryResourceObject) SetType(v CouponCodeEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CouponCodeUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponCodeUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponCodeUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CouponCodeUpdateQueryResourceObject) GetAttributes() CouponCodeUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponCodeUpdateQueryResourceObject) GetAttributesOk() (*CouponCodeUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponCodeUpdateQueryResourceObject) SetAttributes(v CouponCodeUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


