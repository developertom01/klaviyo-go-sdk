# ProfileCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Individual&#39;s email address | [optional] 
**PhoneNumber** | Pointer to **NullableString** | Individual&#39;s phone number in E.164 format | [optional] 
**ExternalId** | Pointer to **NullableString** | A unique identifier used by customers to associate Klaviyo profiles with profiles in an external system, such as a point-of-sale system. Format varies based on the external system. | [optional] 
**FirstName** | Pointer to **NullableString** | Individual&#39;s first name | [optional] 
**LastName** | Pointer to **NullableString** | Individual&#39;s last name | [optional] 
**Organization** | Pointer to **NullableString** | Name of the company or organization within the company for whom the individual works | [optional] 
**Title** | Pointer to **NullableString** | Individual&#39;s job title | [optional] 
**Image** | Pointer to **NullableString** | URL pointing to the location of a profile image | [optional] 
**Location** | Pointer to [**ProfileLocation**](ProfileLocation.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | An object containing key/value pairs for any custom properties assigned to this profile | [optional] 

## Methods

### NewProfileCreateQueryResourceObjectAttributes

`func NewProfileCreateQueryResourceObjectAttributes() *ProfileCreateQueryResourceObjectAttributes`

NewProfileCreateQueryResourceObjectAttributes instantiates a new ProfileCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileCreateQueryResourceObjectAttributesWithDefaults

`func NewProfileCreateQueryResourceObjectAttributesWithDefaults() *ProfileCreateQueryResourceObjectAttributes`

NewProfileCreateQueryResourceObjectAttributesWithDefaults instantiates a new ProfileCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ProfileCreateQueryResourceObjectAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProfileCreateQueryResourceObjectAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProfileCreateQueryResourceObjectAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProfileCreateQueryResourceObjectAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ProfileCreateQueryResourceObjectAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ProfileCreateQueryResourceObjectAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *ProfileCreateQueryResourceObjectAttributes) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ProfileCreateQueryResourceObjectAttributes) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ProfileCreateQueryResourceObjectAttributes) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ProfileCreateQueryResourceObjectAttributes) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ProfileCreateQueryResourceObjectAttributes) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ProfileCreateQueryResourceObjectAttributes) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetExternalId

`func (o *ProfileCreateQueryResourceObjectAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ProfileCreateQueryResourceObjectAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ProfileCreateQueryResourceObjectAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ProfileCreateQueryResourceObjectAttributes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ProfileCreateQueryResourceObjectAttributes) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ProfileCreateQueryResourceObjectAttributes) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetFirstName

`func (o *ProfileCreateQueryResourceObjectAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ProfileCreateQueryResourceObjectAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ProfileCreateQueryResourceObjectAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ProfileCreateQueryResourceObjectAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ProfileCreateQueryResourceObjectAttributes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ProfileCreateQueryResourceObjectAttributes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ProfileCreateQueryResourceObjectAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ProfileCreateQueryResourceObjectAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ProfileCreateQueryResourceObjectAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ProfileCreateQueryResourceObjectAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ProfileCreateQueryResourceObjectAttributes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ProfileCreateQueryResourceObjectAttributes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetOrganization

`func (o *ProfileCreateQueryResourceObjectAttributes) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProfileCreateQueryResourceObjectAttributes) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProfileCreateQueryResourceObjectAttributes) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProfileCreateQueryResourceObjectAttributes) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ProfileCreateQueryResourceObjectAttributes) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ProfileCreateQueryResourceObjectAttributes) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetTitle

`func (o *ProfileCreateQueryResourceObjectAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProfileCreateQueryResourceObjectAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProfileCreateQueryResourceObjectAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProfileCreateQueryResourceObjectAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ProfileCreateQueryResourceObjectAttributes) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ProfileCreateQueryResourceObjectAttributes) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetImage

`func (o *ProfileCreateQueryResourceObjectAttributes) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProfileCreateQueryResourceObjectAttributes) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProfileCreateQueryResourceObjectAttributes) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ProfileCreateQueryResourceObjectAttributes) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *ProfileCreateQueryResourceObjectAttributes) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *ProfileCreateQueryResourceObjectAttributes) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetLocation

`func (o *ProfileCreateQueryResourceObjectAttributes) GetLocation() ProfileLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProfileCreateQueryResourceObjectAttributes) GetLocationOk() (*ProfileLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProfileCreateQueryResourceObjectAttributes) SetLocation(v ProfileLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProfileCreateQueryResourceObjectAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProperties

`func (o *ProfileCreateQueryResourceObjectAttributes) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProfileCreateQueryResourceObjectAttributes) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProfileCreateQueryResourceObjectAttributes) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ProfileCreateQueryResourceObjectAttributes) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ProfileCreateQueryResourceObjectAttributes) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ProfileCreateQueryResourceObjectAttributes) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


