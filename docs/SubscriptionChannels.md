# SubscriptionChannels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**EmailSubscriptionParameters**](EmailSubscriptionParameters.md) |  | [optional] 
**Sms** | Pointer to [**SMSSubscriptionParameters**](SMSSubscriptionParameters.md) |  | [optional] 

## Methods

### NewSubscriptionChannels

`func NewSubscriptionChannels() *SubscriptionChannels`

NewSubscriptionChannels instantiates a new SubscriptionChannels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionChannelsWithDefaults

`func NewSubscriptionChannelsWithDefaults() *SubscriptionChannels`

NewSubscriptionChannelsWithDefaults instantiates a new SubscriptionChannels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SubscriptionChannels) GetEmail() EmailSubscriptionParameters`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubscriptionChannels) GetEmailOk() (*EmailSubscriptionParameters, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubscriptionChannels) SetEmail(v EmailSubscriptionParameters)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SubscriptionChannels) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSms

`func (o *SubscriptionChannels) GetSms() SMSSubscriptionParameters`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *SubscriptionChannels) GetSmsOk() (*SMSSubscriptionParameters, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *SubscriptionChannels) SetSms(v SMSSubscriptionParameters)`

SetSms sets Sms field to given value.

### HasSms

`func (o *SubscriptionChannels) HasSms() bool`

HasSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


