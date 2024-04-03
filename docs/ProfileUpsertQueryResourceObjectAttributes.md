# ProfileUpsertQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Individual&#39;s email address | [optional] 
**PhoneNumber** | Pointer to **NullableString** | Individual&#39;s phone number in E.164 format | [optional] 
**ExternalId** | Pointer to **NullableString** | A unique identifier used by customers to associate Klaviyo profiles with profiles in an external system, such as a point-of-sale system. Format varies based on the external system. | [optional] 
**AnonymousId** | Pointer to **NullableString** | Id that can be used to identify a profile when other identifiers are not available | [optional] 
**Kx** | Pointer to **NullableString** | Also known as the &#x60;exchange_id&#x60;, this is an encrypted identifier used for identifying a profile by Klaviyo&#39;s web tracking.  You can use this field as a filter when retrieving profiles via the Get Profiles endpoint. | [optional] 
**FirstName** | Pointer to **NullableString** | Individual&#39;s first name | [optional] 
**LastName** | Pointer to **NullableString** | Individual&#39;s last name | [optional] 
**Organization** | Pointer to **NullableString** | Name of the company or organization within the company for whom the individual works | [optional] 
**Title** | Pointer to **NullableString** | Individual&#39;s job title | [optional] 
**Image** | Pointer to **NullableString** | URL pointing to the location of a profile image | [optional] 
**Location** | Pointer to [**ProfileLocation**](ProfileLocation.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | An object containing key/value pairs for any custom properties assigned to this profile | [optional] 

## Methods

### NewProfileUpsertQueryResourceObjectAttributes

`func NewProfileUpsertQueryResourceObjectAttributes() *ProfileUpsertQueryResourceObjectAttributes`

NewProfileUpsertQueryResourceObjectAttributes instantiates a new ProfileUpsertQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileUpsertQueryResourceObjectAttributesWithDefaults

`func NewProfileUpsertQueryResourceObjectAttributesWithDefaults() *ProfileUpsertQueryResourceObjectAttributes`

NewProfileUpsertQueryResourceObjectAttributesWithDefaults instantiates a new ProfileUpsertQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetExternalId

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetAnonymousId

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetAnonymousId() string`

GetAnonymousId returns the AnonymousId field if non-nil, zero value otherwise.

### GetAnonymousIdOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetAnonymousIdOk() (*string, bool)`

GetAnonymousIdOk returns a tuple with the AnonymousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousId

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetAnonymousId(v string)`

SetAnonymousId sets AnonymousId field to given value.

### HasAnonymousId

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasAnonymousId() bool`

HasAnonymousId returns a boolean if a field has been set.

### SetAnonymousIdNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetAnonymousIdNil(b bool)`

 SetAnonymousIdNil sets the value for AnonymousId to be an explicit nil

### UnsetAnonymousId
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetAnonymousId()`

UnsetAnonymousId ensures that no value is present for AnonymousId, not even an explicit nil
### GetKx

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetKx() string`

GetKx returns the Kx field if non-nil, zero value otherwise.

### GetKxOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetKxOk() (*string, bool)`

GetKxOk returns a tuple with the Kx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKx

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetKx(v string)`

SetKx sets Kx field to given value.

### HasKx

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasKx() bool`

HasKx returns a boolean if a field has been set.

### SetKxNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetKxNil(b bool)`

 SetKxNil sets the value for Kx to be an explicit nil

### UnsetKx
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetKx()`

UnsetKx ensures that no value is present for Kx, not even an explicit nil
### GetFirstName

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetOrganization

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetTitle

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetImage

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetLocation

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetLocation() ProfileLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetLocationOk() (*ProfileLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetLocation(v ProfileLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProperties

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProfileUpsertQueryResourceObjectAttributes) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ProfileUpsertQueryResourceObjectAttributes) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ProfileUpsertQueryResourceObjectAttributes) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ProfileUpsertQueryResourceObjectAttributes) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


