# Subscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**EmailChannel**](EmailChannel.md) |  | [optional] 
**Sms** | Pointer to [**SMSChannel**](SMSChannel.md) |  | [optional] 

## Methods

### NewSubscriptions

`func NewSubscriptions() *Subscriptions`

NewSubscriptions instantiates a new Subscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsWithDefaults

`func NewSubscriptionsWithDefaults() *Subscriptions`

NewSubscriptionsWithDefaults instantiates a new Subscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Subscriptions) GetEmail() EmailChannel`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Subscriptions) GetEmailOk() (*EmailChannel, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Subscriptions) SetEmail(v EmailChannel)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Subscriptions) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSms

`func (o *Subscriptions) GetSms() SMSChannel`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *Subscriptions) GetSmsOk() (*SMSChannel, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *Subscriptions) SetSms(v SMSChannel)`

SetSms sets Sms field to given value.

### HasSms

`func (o *Subscriptions) HasSms() bool`

HasSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


