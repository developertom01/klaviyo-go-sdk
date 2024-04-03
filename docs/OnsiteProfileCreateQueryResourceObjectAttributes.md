# OnsiteProfileCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Individual&#39;s email address | [optional] 
**PhoneNumber** | Pointer to **NullableString** | Individual&#39;s phone number in E.164 format | [optional] 
**ExternalId** | Pointer to **NullableString** | A unique identifier used by customers to associate Klaviyo profiles with profiles in an external system, such as a point-of-sale system. Format varies based on the external system. | [optional] 
**AnonymousId** | Pointer to **NullableString** |  | [optional] 
**Kx** | Pointer to **NullableString** | Also known as the &#x60;exchange_id&#x60;, this is an encrypted identifier used for identifying a profile by Klaviyo&#39;s web tracking.  You can use this field as a filter when retrieving profiles via the Get Profiles endpoint. | [optional] 
**FirstName** | Pointer to **NullableString** | Individual&#39;s first name | [optional] 
**LastName** | Pointer to **NullableString** | Individual&#39;s last name | [optional] 
**Organization** | Pointer to **NullableString** | Name of the company or organization within the company for whom the individual works | [optional] 
**Title** | Pointer to **NullableString** | Individual&#39;s job title | [optional] 
**Image** | Pointer to **NullableString** | URL pointing to the location of a profile image | [optional] 
**Location** | Pointer to [**ProfileLocation**](ProfileLocation.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | An object containing key/value pairs for any custom properties assigned to this profile | [optional] 

## Methods

### NewOnsiteProfileCreateQueryResourceObjectAttributes

`func NewOnsiteProfileCreateQueryResourceObjectAttributes() *OnsiteProfileCreateQueryResourceObjectAttributes`

NewOnsiteProfileCreateQueryResourceObjectAttributes instantiates a new OnsiteProfileCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnsiteProfileCreateQueryResourceObjectAttributesWithDefaults

`func NewOnsiteProfileCreateQueryResourceObjectAttributesWithDefaults() *OnsiteProfileCreateQueryResourceObjectAttributes`

NewOnsiteProfileCreateQueryResourceObjectAttributesWithDefaults instantiates a new OnsiteProfileCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetExternalId

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetAnonymousId

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetAnonymousId() string`

GetAnonymousId returns the AnonymousId field if non-nil, zero value otherwise.

### GetAnonymousIdOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetAnonymousIdOk() (*string, bool)`

GetAnonymousIdOk returns a tuple with the AnonymousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousId

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetAnonymousId(v string)`

SetAnonymousId sets AnonymousId field to given value.

### HasAnonymousId

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasAnonymousId() bool`

HasAnonymousId returns a boolean if a field has been set.

### SetAnonymousIdNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetAnonymousIdNil(b bool)`

 SetAnonymousIdNil sets the value for AnonymousId to be an explicit nil

### UnsetAnonymousId
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetAnonymousId()`

UnsetAnonymousId ensures that no value is present for AnonymousId, not even an explicit nil
### GetKx

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetKx() string`

GetKx returns the Kx field if non-nil, zero value otherwise.

### GetKxOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetKxOk() (*string, bool)`

GetKxOk returns a tuple with the Kx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKx

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetKx(v string)`

SetKx sets Kx field to given value.

### HasKx

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasKx() bool`

HasKx returns a boolean if a field has been set.

### SetKxNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetKxNil(b bool)`

 SetKxNil sets the value for Kx to be an explicit nil

### UnsetKx
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetKx()`

UnsetKx ensures that no value is present for Kx, not even an explicit nil
### GetFirstName

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetOrganization

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetTitle

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetImage

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetLocation

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetLocation() ProfileLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetLocationOk() (*ProfileLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetLocation(v ProfileLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProperties

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *OnsiteProfileCreateQueryResourceObjectAttributes) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


