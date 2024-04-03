# ContactInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSenderName** | **string** | This field is used to auto-populate the default sender name on flow and campaign emails. | 
**DefaultSenderEmail** | **string** | This field is used to auto-populate the default sender email address on flow and campaign emails. | 
**WebsiteUrl** | **string** |  | 
**OrganizationName** | **string** |  | 
**StreetAddress** | [**StreetAddress**](StreetAddress.md) |  | 

## Methods

### NewContactInformation

`func NewContactInformation(defaultSenderName string, defaultSenderEmail string, websiteUrl string, organizationName string, streetAddress StreetAddress, ) *ContactInformation`

NewContactInformation instantiates a new ContactInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactInformationWithDefaults

`func NewContactInformationWithDefaults() *ContactInformation`

NewContactInformationWithDefaults instantiates a new ContactInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSenderName

`func (o *ContactInformation) GetDefaultSenderName() string`

GetDefaultSenderName returns the DefaultSenderName field if non-nil, zero value otherwise.

### GetDefaultSenderNameOk

`func (o *ContactInformation) GetDefaultSenderNameOk() (*string, bool)`

GetDefaultSenderNameOk returns a tuple with the DefaultSenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSenderName

`func (o *ContactInformation) SetDefaultSenderName(v string)`

SetDefaultSenderName sets DefaultSenderName field to given value.


### GetDefaultSenderEmail

`func (o *ContactInformation) GetDefaultSenderEmail() string`

GetDefaultSenderEmail returns the DefaultSenderEmail field if non-nil, zero value otherwise.

### GetDefaultSenderEmailOk

`func (o *ContactInformation) GetDefaultSenderEmailOk() (*string, bool)`

GetDefaultSenderEmailOk returns a tuple with the DefaultSenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSenderEmail

`func (o *ContactInformation) SetDefaultSenderEmail(v string)`

SetDefaultSenderEmail sets DefaultSenderEmail field to given value.


### GetWebsiteUrl

`func (o *ContactInformation) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *ContactInformation) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *ContactInformation) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.


### GetOrganizationName

`func (o *ContactInformation) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ContactInformation) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ContactInformation) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetStreetAddress

`func (o *ContactInformation) GetStreetAddress() StreetAddress`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *ContactInformation) GetStreetAddressOk() (*StreetAddress, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *ContactInformation) SetStreetAddress(v StreetAddress)`

SetStreetAddress sets StreetAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


