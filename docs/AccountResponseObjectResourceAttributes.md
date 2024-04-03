# AccountResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestAccount** | **bool** | Indicates if the account is a test account. Test accounts are not a separate testing engineering environment. Test accounts use the same production environment as normal Klaviyo accounts. This feature is primarily UI based to reduce human errors | 
**ContactInformation** | [**ContactInformation**](ContactInformation.md) |  | 
**Industry** | **string** | The kind of business and/or types of goods that the business sells. This is leveraged in Klaviyo analytics and guidance. | 
**Timezone** | **string** | The account&#39;s timezone is used when displaying dates and times. This is an IANA timezone. See [the full list here ](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). | 
**PreferredCurrency** | **string** | The preferred currency for the account. This is the currency used for currency-based metrics in dashboards, analytics, coupons, and templates. | 
**PublicApiKey** | **string** | The Public API Key can be used for client-side API calls. [More info here](https://developers.klaviyo.com/en/docs/retrieve_api_credentials). | 

## Methods

### NewAccountResponseObjectResourceAttributes

`func NewAccountResponseObjectResourceAttributes(testAccount bool, contactInformation ContactInformation, industry string, timezone string, preferredCurrency string, publicApiKey string, ) *AccountResponseObjectResourceAttributes`

NewAccountResponseObjectResourceAttributes instantiates a new AccountResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseObjectResourceAttributesWithDefaults

`func NewAccountResponseObjectResourceAttributesWithDefaults() *AccountResponseObjectResourceAttributes`

NewAccountResponseObjectResourceAttributesWithDefaults instantiates a new AccountResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestAccount

`func (o *AccountResponseObjectResourceAttributes) GetTestAccount() bool`

GetTestAccount returns the TestAccount field if non-nil, zero value otherwise.

### GetTestAccountOk

`func (o *AccountResponseObjectResourceAttributes) GetTestAccountOk() (*bool, bool)`

GetTestAccountOk returns a tuple with the TestAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccount

`func (o *AccountResponseObjectResourceAttributes) SetTestAccount(v bool)`

SetTestAccount sets TestAccount field to given value.


### GetContactInformation

`func (o *AccountResponseObjectResourceAttributes) GetContactInformation() ContactInformation`

GetContactInformation returns the ContactInformation field if non-nil, zero value otherwise.

### GetContactInformationOk

`func (o *AccountResponseObjectResourceAttributes) GetContactInformationOk() (*ContactInformation, bool)`

GetContactInformationOk returns a tuple with the ContactInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInformation

`func (o *AccountResponseObjectResourceAttributes) SetContactInformation(v ContactInformation)`

SetContactInformation sets ContactInformation field to given value.


### GetIndustry

`func (o *AccountResponseObjectResourceAttributes) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *AccountResponseObjectResourceAttributes) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *AccountResponseObjectResourceAttributes) SetIndustry(v string)`

SetIndustry sets Industry field to given value.


### GetTimezone

`func (o *AccountResponseObjectResourceAttributes) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AccountResponseObjectResourceAttributes) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AccountResponseObjectResourceAttributes) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetPreferredCurrency

`func (o *AccountResponseObjectResourceAttributes) GetPreferredCurrency() string`

GetPreferredCurrency returns the PreferredCurrency field if non-nil, zero value otherwise.

### GetPreferredCurrencyOk

`func (o *AccountResponseObjectResourceAttributes) GetPreferredCurrencyOk() (*string, bool)`

GetPreferredCurrencyOk returns a tuple with the PreferredCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredCurrency

`func (o *AccountResponseObjectResourceAttributes) SetPreferredCurrency(v string)`

SetPreferredCurrency sets PreferredCurrency field to given value.


### GetPublicApiKey

`func (o *AccountResponseObjectResourceAttributes) GetPublicApiKey() string`

GetPublicApiKey returns the PublicApiKey field if non-nil, zero value otherwise.

### GetPublicApiKeyOk

`func (o *AccountResponseObjectResourceAttributes) GetPublicApiKeyOk() (*string, bool)`

GetPublicApiKeyOk returns a tuple with the PublicApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicApiKey

`func (o *AccountResponseObjectResourceAttributes) SetPublicApiKey(v string)`

SetPublicApiKey sets PublicApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


