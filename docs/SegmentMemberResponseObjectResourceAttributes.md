# SegmentMemberResponseObjectResourceAttributes

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
**JoinedGroupAt** | **time.Time** | The datetime when this profile most recently joined the segment. | 

## Methods

### NewSegmentMemberResponseObjectResourceAttributes

`func NewSegmentMemberResponseObjectResourceAttributes(joinedGroupAt time.Time, ) *SegmentMemberResponseObjectResourceAttributes`

NewSegmentMemberResponseObjectResourceAttributes instantiates a new SegmentMemberResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentMemberResponseObjectResourceAttributesWithDefaults

`func NewSegmentMemberResponseObjectResourceAttributesWithDefaults() *SegmentMemberResponseObjectResourceAttributes`

NewSegmentMemberResponseObjectResourceAttributesWithDefaults instantiates a new SegmentMemberResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SegmentMemberResponseObjectResourceAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SegmentMemberResponseObjectResourceAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SegmentMemberResponseObjectResourceAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *SegmentMemberResponseObjectResourceAttributes) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *SegmentMemberResponseObjectResourceAttributes) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *SegmentMemberResponseObjectResourceAttributes) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetExternalId

`func (o *SegmentMemberResponseObjectResourceAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SegmentMemberResponseObjectResourceAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SegmentMemberResponseObjectResourceAttributes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetFirstName

`func (o *SegmentMemberResponseObjectResourceAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SegmentMemberResponseObjectResourceAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SegmentMemberResponseObjectResourceAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *SegmentMemberResponseObjectResourceAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SegmentMemberResponseObjectResourceAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SegmentMemberResponseObjectResourceAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetOrganization

`func (o *SegmentMemberResponseObjectResourceAttributes) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SegmentMemberResponseObjectResourceAttributes) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SegmentMemberResponseObjectResourceAttributes) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetTitle

`func (o *SegmentMemberResponseObjectResourceAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SegmentMemberResponseObjectResourceAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SegmentMemberResponseObjectResourceAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetImage

`func (o *SegmentMemberResponseObjectResourceAttributes) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SegmentMemberResponseObjectResourceAttributes) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *SegmentMemberResponseObjectResourceAttributes) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetCreated

`func (o *SegmentMemberResponseObjectResourceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SegmentMemberResponseObjectResourceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SegmentMemberResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *SegmentMemberResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SegmentMemberResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SegmentMemberResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetLastEventDate

`func (o *SegmentMemberResponseObjectResourceAttributes) GetLastEventDate() time.Time`

GetLastEventDate returns the LastEventDate field if non-nil, zero value otherwise.

### GetLastEventDateOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetLastEventDateOk() (*time.Time, bool)`

GetLastEventDateOk returns a tuple with the LastEventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventDate

`func (o *SegmentMemberResponseObjectResourceAttributes) SetLastEventDate(v time.Time)`

SetLastEventDate sets LastEventDate field to given value.

### HasLastEventDate

`func (o *SegmentMemberResponseObjectResourceAttributes) HasLastEventDate() bool`

HasLastEventDate returns a boolean if a field has been set.

### SetLastEventDateNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetLastEventDateNil(b bool)`

 SetLastEventDateNil sets the value for LastEventDate to be an explicit nil

### UnsetLastEventDate
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetLastEventDate()`

UnsetLastEventDate ensures that no value is present for LastEventDate, not even an explicit nil
### GetLocation

`func (o *SegmentMemberResponseObjectResourceAttributes) GetLocation() ProfileLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetLocationOk() (*ProfileLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SegmentMemberResponseObjectResourceAttributes) SetLocation(v ProfileLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SegmentMemberResponseObjectResourceAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProperties

`func (o *SegmentMemberResponseObjectResourceAttributes) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SegmentMemberResponseObjectResourceAttributes) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SegmentMemberResponseObjectResourceAttributes) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *SegmentMemberResponseObjectResourceAttributes) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *SegmentMemberResponseObjectResourceAttributes) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetJoinedGroupAt

`func (o *SegmentMemberResponseObjectResourceAttributes) GetJoinedGroupAt() time.Time`

GetJoinedGroupAt returns the JoinedGroupAt field if non-nil, zero value otherwise.

### GetJoinedGroupAtOk

`func (o *SegmentMemberResponseObjectResourceAttributes) GetJoinedGroupAtOk() (*time.Time, bool)`

GetJoinedGroupAtOk returns a tuple with the JoinedGroupAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedGroupAt

`func (o *SegmentMemberResponseObjectResourceAttributes) SetJoinedGroupAt(v time.Time)`

SetJoinedGroupAt sets JoinedGroupAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


