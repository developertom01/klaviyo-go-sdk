# EmailSubscriptionParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Marketing** | [**MarketingSubscriptionParameters**](MarketingSubscriptionParameters.md) |  | 

## Methods

### NewEmailSubscriptionParameters

`func NewEmailSubscriptionParameters(marketing MarketingSubscriptionParameters, ) *EmailSubscriptionParameters`

NewEmailSubscriptionParameters instantiates a new EmailSubscriptionParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSubscriptionParametersWithDefaults

`func NewEmailSubscriptionParametersWithDefaults() *EmailSubscriptionParameters`

NewEmailSubscriptionParametersWithDefaults instantiates a new EmailSubscriptionParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketing

`func (o *EmailSubscriptionParameters) GetMarketing() MarketingSubscriptionParameters`

GetMarketing returns the Marketing field if non-nil, zero value otherwise.

### GetMarketingOk

`func (o *EmailSubscriptionParameters) GetMarketingOk() (*MarketingSubscriptionParameters, bool)`

GetMarketingOk returns a tuple with the Marketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketing

`func (o *EmailSubscriptionParameters) SetMarketing(v MarketingSubscriptionParameters)`

SetMarketing sets Marketing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


