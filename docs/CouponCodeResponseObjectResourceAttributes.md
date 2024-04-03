# CouponCodeResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueCode** | Pointer to **NullableString** | This is a unique string that will be or is assigned to each customer/profile and is associated with a coupon. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | The datetime when this coupon code will expire. If not specified or set to null, it will be automatically set to 1 year. | [optional] 
**Status** | Pointer to **NullableString** | The current status of the coupon code. | [optional] 

## Methods

### NewCouponCodeResponseObjectResourceAttributes

`func NewCouponCodeResponseObjectResourceAttributes() *CouponCodeResponseObjectResourceAttributes`

NewCouponCodeResponseObjectResourceAttributes instantiates a new CouponCodeResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCodeResponseObjectResourceAttributesWithDefaults

`func NewCouponCodeResponseObjectResourceAttributesWithDefaults() *CouponCodeResponseObjectResourceAttributes`

NewCouponCodeResponseObjectResourceAttributesWithDefaults instantiates a new CouponCodeResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueCode

`func (o *CouponCodeResponseObjectResourceAttributes) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *CouponCodeResponseObjectResourceAttributes) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *CouponCodeResponseObjectResourceAttributes) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.

### HasUniqueCode

`func (o *CouponCodeResponseObjectResourceAttributes) HasUniqueCode() bool`

HasUniqueCode returns a boolean if a field has been set.

### SetUniqueCodeNil

`func (o *CouponCodeResponseObjectResourceAttributes) SetUniqueCodeNil(b bool)`

 SetUniqueCodeNil sets the value for UniqueCode to be an explicit nil

### UnsetUniqueCode
`func (o *CouponCodeResponseObjectResourceAttributes) UnsetUniqueCode()`

UnsetUniqueCode ensures that no value is present for UniqueCode, not even an explicit nil
### GetExpiresAt

`func (o *CouponCodeResponseObjectResourceAttributes) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CouponCodeResponseObjectResourceAttributes) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CouponCodeResponseObjectResourceAttributes) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CouponCodeResponseObjectResourceAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CouponCodeResponseObjectResourceAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CouponCodeResponseObjectResourceAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetStatus

`func (o *CouponCodeResponseObjectResourceAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponCodeResponseObjectResourceAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CouponCodeResponseObjectResourceAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CouponCodeResponseObjectResourceAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CouponCodeResponseObjectResourceAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CouponCodeResponseObjectResourceAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


