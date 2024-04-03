# CouponCodeUpdateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **NullableString** | The API status of our coupon codes. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | The datetime when this coupon code will expire. If not specified or set to null, it will be automatically set to 1 year. | [optional] 

## Methods

### NewCouponCodeUpdateQueryResourceObjectAttributes

`func NewCouponCodeUpdateQueryResourceObjectAttributes() *CouponCodeUpdateQueryResourceObjectAttributes`

NewCouponCodeUpdateQueryResourceObjectAttributes instantiates a new CouponCodeUpdateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCodeUpdateQueryResourceObjectAttributesWithDefaults

`func NewCouponCodeUpdateQueryResourceObjectAttributesWithDefaults() *CouponCodeUpdateQueryResourceObjectAttributes`

NewCouponCodeUpdateQueryResourceObjectAttributesWithDefaults instantiates a new CouponCodeUpdateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CouponCodeUpdateQueryResourceObjectAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponCodeUpdateQueryResourceObjectAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CouponCodeUpdateQueryResourceObjectAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CouponCodeUpdateQueryResourceObjectAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CouponCodeUpdateQueryResourceObjectAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CouponCodeUpdateQueryResourceObjectAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetExpiresAt

`func (o *CouponCodeUpdateQueryResourceObjectAttributes) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CouponCodeUpdateQueryResourceObjectAttributes) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CouponCodeUpdateQueryResourceObjectAttributes) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CouponCodeUpdateQueryResourceObjectAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CouponCodeUpdateQueryResourceObjectAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CouponCodeUpdateQueryResourceObjectAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


