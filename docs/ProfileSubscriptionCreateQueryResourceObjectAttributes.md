# ProfileSubscriptionCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | The email address to subscribe or to set on the profile if &#x60;channels&#x60; is specified and the email channel is omitted. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number to subscribe or to set on the profile if &#x60;channels&#x60; is specified and the SMS channel is omitted. This must be in E.164 format. | [optional] 
**Subscriptions** | Pointer to [**SubscriptionChannels**](SubscriptionChannels.md) |  | [optional] 

## Methods

### NewProfileSubscriptionCreateQueryResourceObjectAttributes

`func NewProfileSubscriptionCreateQueryResourceObjectAttributes() *ProfileSubscriptionCreateQueryResourceObjectAttributes`

NewProfileSubscriptionCreateQueryResourceObjectAttributes instantiates a new ProfileSubscriptionCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileSubscriptionCreateQueryResourceObjectAttributesWithDefaults

`func NewProfileSubscriptionCreateQueryResourceObjectAttributesWithDefaults() *ProfileSubscriptionCreateQueryResourceObjectAttributes`

NewProfileSubscriptionCreateQueryResourceObjectAttributesWithDefaults instantiates a new ProfileSubscriptionCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetSubscriptions

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) GetSubscriptions() SubscriptionChannels`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) GetSubscriptionsOk() (*SubscriptionChannels, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) SetSubscriptions(v SubscriptionChannels)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *ProfileSubscriptionCreateQueryResourceObjectAttributes) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


