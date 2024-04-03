# CouponCodeCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueCode** | **string** | This is a unique string that will be or is assigned to each customer/profile and is associated with a coupon. | 
**ExpiresAt** | Pointer to **NullableTime** | The datetime when this coupon code will expire. If not specified or set to null, it will be automatically set to 1 year. | [optional] 

## Methods

### NewCouponCodeCreateQueryResourceObjectAttributes

`func NewCouponCodeCreateQueryResourceObjectAttributes(uniqueCode string, ) *CouponCodeCreateQueryResourceObjectAttributes`

NewCouponCodeCreateQueryResourceObjectAttributes instantiates a new CouponCodeCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCodeCreateQueryResourceObjectAttributesWithDefaults

`func NewCouponCodeCreateQueryResourceObjectAttributesWithDefaults() *CouponCodeCreateQueryResourceObjectAttributes`

NewCouponCodeCreateQueryResourceObjectAttributesWithDefaults instantiates a new CouponCodeCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueCode

`func (o *CouponCodeCreateQueryResourceObjectAttributes) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *CouponCodeCreateQueryResourceObjectAttributes) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *CouponCodeCreateQueryResourceObjectAttributes) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.


### GetExpiresAt

`func (o *CouponCodeCreateQueryResourceObjectAttributes) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CouponCodeCreateQueryResourceObjectAttributes) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CouponCodeCreateQueryResourceObjectAttributes) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CouponCodeCreateQueryResourceObjectAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CouponCodeCreateQueryResourceObjectAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CouponCodeCreateQueryResourceObjectAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


