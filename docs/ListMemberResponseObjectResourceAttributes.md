# ListMemberResponseObjectResourceAttributes

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
**Created** | Pointer to **NullableTime** | Date and time when the profile was created, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 
**Updated** | Pointer to **NullableTime** | Date and time when the profile was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 
**LastEventDate** | Pointer to **NullableTime** | Date and time of the most recent event the triggered an update to the profile, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 
**Location** | Pointer to [**ProfileLocation**](ProfileLocation.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | An object containing key/value pairs for any custom properties assigned to this profile | [optional] 
**JoinedGroupAt** | **time.Time** | The datetime when this profile most recently joined the list. | 

## Methods

### NewListMemberResponseObjectResourceAttributes

`func NewListMemberResponseObjectResourceAttributes(joinedGroupAt time.Time, ) *ListMemberResponseObjectResourceAttributes`

NewListMemberResponseObjectResourceAttributes instantiates a new ListMemberResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMemberResponseObjectResourceAttributesWithDefaults

`func NewListMemberResponseObjectResourceAttributesWithDefaults() *ListMemberResponseObjectResourceAttributes`

NewListMemberResponseObjectResourceAttributesWithDefaults instantiates a new ListMemberResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ListMemberResponseObjectResourceAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ListMemberResponseObjectResourceAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ListMemberResponseObjectResourceAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ListMemberResponseObjectResourceAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ListMemberResponseObjectResourceAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ListMemberResponseObjectResourceAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *ListMemberResponseObjectResourceAttributes) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ListMemberResponseObjectResourceAttributes) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ListMemberResponseObjectResourceAttributes) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ListMemberResponseObjectResourceAttributes) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ListMemberResponseObjectResourceAttributes) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ListMemberResponseObjectResourceAttributes) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetExternalId

`func (o *ListMemberResponseObjectResourceAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListMemberResponseObjectResourceAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListMemberResponseObjectResourceAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListMemberResponseObjectResourceAttributes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListMemberResponseObjectResourceAttributes) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListMemberResponseObjectResourceAttributes) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetFirstName

`func (o *ListMemberResponseObjectResourceAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ListMemberResponseObjectResourceAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ListMemberResponseObjectResourceAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ListMemberResponseObjectResourceAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ListMemberResponseObjectResourceAttributes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ListMemberResponseObjectResourceAttributes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ListMemberResponseObjectResourceAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ListMemberResponseObjectResourceAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ListMemberResponseObjectResourceAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ListMemberResponseObjectResourceAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ListMemberResponseObjectResourceAttributes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ListMemberResponseObjectResourceAttributes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetOrganization

`func (o *ListMemberResponseObjectResourceAttributes) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ListMemberResponseObjectResourceAttributes) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ListMemberResponseObjectResourceAttributes) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ListMemberResponseObjectResourceAttributes) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ListMemberResponseObjectResourceAttributes) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ListMemberResponseObjectResourceAttributes) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetTitle

`func (o *ListMemberResponseObjectResourceAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListMemberResponseObjectResourceAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListMemberResponseObjectResourceAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListMemberResponseObjectResourceAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ListMemberResponseObjectResourceAttributes) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ListMemberResponseObjectResourceAttributes) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetImage

`func (o *ListMemberResponseObjectResourceAttributes) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ListMemberResponseObjectResourceAttributes) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ListMemberResponseObjectResourceAttributes) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ListMemberResponseObjectResourceAttributes) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *ListMemberResponseObjectResourceAttributes) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *ListMemberResponseObjectResourceAttributes) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetCreated

`func (o *ListMemberResponseObjectResourceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListMemberResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListMemberResponseObjectResourceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListMemberResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *ListMemberResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ListMemberResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *ListMemberResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ListMemberResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ListMemberResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ListMemberResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *ListMemberResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *ListMemberResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetLastEventDate

`func (o *ListMemberResponseObjectResourceAttributes) GetLastEventDate() time.Time`

GetLastEventDate returns the LastEventDate field if non-nil, zero value otherwise.

### GetLastEventDateOk

`func (o *ListMemberResponseObjectResourceAttributes) GetLastEventDateOk() (*time.Time, bool)`

GetLastEventDateOk returns a tuple with the LastEventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventDate

`func (o *ListMemberResponseObjectResourceAttributes) SetLastEventDate(v time.Time)`

SetLastEventDate sets LastEventDate field to given value.

### HasLastEventDate

`func (o *ListMemberResponseObjectResourceAttributes) HasLastEventDate() bool`

HasLastEventDate returns a boolean if a field has been set.

### SetLastEventDateNil

`func (o *ListMemberResponseObjectResourceAttributes) SetLastEventDateNil(b bool)`

 SetLastEventDateNil sets the value for LastEventDate to be an explicit nil

### UnsetLastEventDate
`func (o *ListMemberResponseObjectResourceAttributes) UnsetLastEventDate()`

UnsetLastEventDate ensures that no value is present for LastEventDate, not even an explicit nil
### GetLocation

`func (o *ListMemberResponseObjectResourceAttributes) GetLocation() ProfileLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ListMemberResponseObjectResourceAttributes) GetLocationOk() (*ProfileLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ListMemberResponseObjectResourceAttributes) SetLocation(v ProfileLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ListMemberResponseObjectResourceAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProperties

`func (o *ListMemberResponseObjectResourceAttributes) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ListMemberResponseObjectResourceAttributes) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ListMemberResponseObjectResourceAttributes) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ListMemberResponseObjectResourceAttributes) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ListMemberResponseObjectResourceAttributes) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ListMemberResponseObjectResourceAttributes) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetJoinedGroupAt

`func (o *ListMemberResponseObjectResourceAttributes) GetJoinedGroupAt() time.Time`

GetJoinedGroupAt returns the JoinedGroupAt field if non-nil, zero value otherwise.

### GetJoinedGroupAtOk

`func (o *ListMemberResponseObjectResourceAttributes) GetJoinedGroupAtOk() (*time.Time, bool)`

GetJoinedGroupAtOk returns a tuple with the JoinedGroupAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedGroupAt

`func (o *ListMemberResponseObjectResourceAttributes) SetJoinedGroupAt(v time.Time)`

SetJoinedGroupAt sets JoinedGroupAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


