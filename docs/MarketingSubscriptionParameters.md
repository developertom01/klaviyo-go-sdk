# MarketingSubscriptionParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consent** | **string** | The Consent status to subscribe to for the \&quot;Marketing\&quot; type. Currently supports \&quot;SUBSCRIBED\&quot;. | 
**ConsentedAt** | Pointer to **NullableTime** | The timestamp of when the profile&#39;s consent was gathered | [optional] 

## Methods

### NewMarketingSubscriptionParameters

`func NewMarketingSubscriptionParameters(consent string, ) *MarketingSubscriptionParameters`

NewMarketingSubscriptionParameters instantiates a new MarketingSubscriptionParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSubscriptionParametersWithDefaults

`func NewMarketingSubscriptionParametersWithDefaults() *MarketingSubscriptionParameters`

NewMarketingSubscriptionParametersWithDefaults instantiates a new MarketingSubscriptionParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsent

`func (o *MarketingSubscriptionParameters) GetConsent() string`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *MarketingSubscriptionParameters) GetConsentOk() (*string, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *MarketingSubscriptionParameters) SetConsent(v string)`

SetConsent sets Consent field to given value.


### GetConsentedAt

`func (o *MarketingSubscriptionParameters) GetConsentedAt() time.Time`

GetConsentedAt returns the ConsentedAt field if non-nil, zero value otherwise.

### GetConsentedAtOk

`func (o *MarketingSubscriptionParameters) GetConsentedAtOk() (*time.Time, bool)`

GetConsentedAtOk returns a tuple with the ConsentedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentedAt

`func (o *MarketingSubscriptionParameters) SetConsentedAt(v time.Time)`

SetConsentedAt sets ConsentedAt field to given value.

### HasConsentedAt

`func (o *MarketingSubscriptionParameters) HasConsentedAt() bool`

HasConsentedAt returns a boolean if a field has been set.

### SetConsentedAtNil

`func (o *MarketingSubscriptionParameters) SetConsentedAtNil(b bool)`

 SetConsentedAtNil sets the value for ConsentedAt to be an explicit nil

### UnsetConsentedAt
`func (o *MarketingSubscriptionParameters) UnsetConsentedAt()`

UnsetConsentedAt ensures that no value is present for ConsentedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


