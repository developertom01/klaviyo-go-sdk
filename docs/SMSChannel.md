# SMSChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Marketing** | Pointer to [**SMSMarketing**](SMSMarketing.md) |  | [optional] 

## Methods

### NewSMSChannel

`func NewSMSChannel() *SMSChannel`

NewSMSChannel instantiates a new SMSChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSChannelWithDefaults

`func NewSMSChannelWithDefaults() *SMSChannel`

NewSMSChannelWithDefaults instantiates a new SMSChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketing

`func (o *SMSChannel) GetMarketing() SMSMarketing`

GetMarketing returns the Marketing field if non-nil, zero value otherwise.

### GetMarketingOk

`func (o *SMSChannel) GetMarketingOk() (*SMSMarketing, bool)`

GetMarketingOk returns a tuple with the Marketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketing

`func (o *SMSChannel) SetMarketing(v SMSMarketing)`

SetMarketing sets Marketing field to given value.

### HasMarketing

`func (o *SMSChannel) HasMarketing() bool`

HasMarketing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


