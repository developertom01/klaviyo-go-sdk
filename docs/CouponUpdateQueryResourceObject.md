# CouponUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CouponEnum**](CouponEnum.md) |  | 
**Id** | **string** | The internal id of a Coupon is equivalent to its external id stored within an integration. | 
**Attributes** | [**CouponUpdateQueryResourceObjectAttributes**](CouponUpdateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewCouponUpdateQueryResourceObject

`func NewCouponUpdateQueryResourceObject(type_ CouponEnum, id string, attributes CouponUpdateQueryResourceObjectAttributes, ) *CouponUpdateQueryResourceObject`

NewCouponUpdateQueryResourceObject instantiates a new CouponUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponUpdateQueryResourceObjectWithDefaults

`func NewCouponUpdateQueryResourceObjectWithDefaults() *CouponUpdateQueryResourceObject`

NewCouponUpdateQueryResourceObjectWithDefaults instantiates a new CouponUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponUpdateQueryResourceObject) GetType() CouponEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponUpdateQueryResourceObject) GetTypeOk() (*CouponEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponUpdateQueryResourceObject) SetType(v CouponEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CouponUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CouponUpdateQueryResourceObject) GetAttributes() CouponUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponUpdateQueryResourceObject) GetAttributesOk() (*CouponUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponUpdateQueryResourceObject) SetAttributes(v CouponUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


