# SendOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseSmartSending** | **bool** |  | 
**IsTransactional** | **bool** |  | 

## Methods

### NewSendOptions

`func NewSendOptions(useSmartSending bool, isTransactional bool, ) *SendOptions`

NewSendOptions instantiates a new SendOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendOptionsWithDefaults

`func NewSendOptionsWithDefaults() *SendOptions`

NewSendOptionsWithDefaults instantiates a new SendOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseSmartSending

`func (o *SendOptions) GetUseSmartSending() bool`

GetUseSmartSending returns the UseSmartSending field if non-nil, zero value otherwise.

### GetUseSmartSendingOk

`func (o *SendOptions) GetUseSmartSendingOk() (*bool, bool)`

GetUseSmartSendingOk returns a tuple with the UseSmartSending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSmartSending

`func (o *SendOptions) SetUseSmartSending(v bool)`

SetUseSmartSending sets UseSmartSending field to given value.


### GetIsTransactional

`func (o *SendOptions) GetIsTransactional() bool`

GetIsTransactional returns the IsTransactional field if non-nil, zero value otherwise.

### GetIsTransactionalOk

`func (o *SendOptions) GetIsTransactionalOk() (*bool, bool)`

GetIsTransactionalOk returns a tuple with the IsTransactional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTransactional

`func (o *SendOptions) SetIsTransactional(v bool)`

SetIsTransactional sets IsTransactional field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


