# PostProfileResponseDataAttributes

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
**Subscriptions** | Pointer to [**Subscriptions**](Subscriptions.md) |  | [optional] 
**PredictiveAnalytics** | Pointer to [**PredictiveAnalytics**](PredictiveAnalytics.md) |  | [optional] 

## Methods

### NewPostProfileResponseDataAttributes

`func NewPostProfileResponseDataAttributes() *PostProfileResponseDataAttributes`

NewPostProfileResponseDataAttributes instantiates a new PostProfileResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostProfileResponseDataAttributesWithDefaults

`func NewPostProfileResponseDataAttributesWithDefaults() *PostProfileResponseDataAttributes`

NewPostProfileResponseDataAttributesWithDefaults instantiates a new PostProfileResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PostProfileResponseDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostProfileResponseDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostProfileResponseDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PostProfileResponseDataAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PostProfileResponseDataAttributes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PostProfileResponseDataAttributes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *PostProfileResponseDataAttributes) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PostProfileResponseDataAttributes) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PostProfileResponseDataAttributes) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PostProfileResponseDataAttributes) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *PostProfileResponseDataAttributes) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *PostProfileResponseDataAttributes) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetExternalId

`func (o *PostProfileResponseDataAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PostProfileResponseDataAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PostProfileResponseDataAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PostProfileResponseDataAttributes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *PostProfileResponseDataAttributes) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *PostProfileResponseDataAttributes) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetFirstName

`func (o *PostProfileResponseDataAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PostProfileResponseDataAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PostProfileResponseDataAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PostProfileResponseDataAttributes) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *PostProfileResponseDataAttributes) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *PostProfileResponseDataAttributes) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *PostProfileResponseDataAttributes) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PostProfileResponseDataAttributes) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PostProfileResponseDataAttributes) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PostProfileResponseDataAttributes) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *PostProfileResponseDataAttributes) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *PostProfileResponseDataAttributes) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetOrganization

`func (o *PostProfileResponseDataAttributes) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PostProfileResponseDataAttributes) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PostProfileResponseDataAttributes) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PostProfileResponseDataAttributes) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *PostProfileResponseDataAttributes) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *PostProfileResponseDataAttributes) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetTitle

`func (o *PostProfileResponseDataAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostProfileResponseDataAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostProfileResponseDataAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PostProfileResponseDataAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PostProfileResponseDataAttributes) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PostProfileResponseDataAttributes) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetImage

`func (o *PostProfileResponseDataAttributes) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PostProfileResponseDataAttributes) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PostProfileResponseDataAttributes) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PostProfileResponseDataAttributes) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *PostProfileResponseDataAttributes) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *PostProfileResponseDataAttributes) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetCreated

`func (o *PostProfileResponseDataAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PostProfileResponseDataAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PostProfileResponseDataAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PostProfileResponseDataAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *PostProfileResponseDataAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *PostProfileResponseDataAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *PostProfileResponseDataAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PostProfileResponseDataAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PostProfileResponseDataAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PostProfileResponseDataAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *PostProfileResponseDataAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *PostProfileResponseDataAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetLastEventDate

`func (o *PostProfileResponseDataAttributes) GetLastEventDate() time.Time`

GetLastEventDate returns the LastEventDate field if non-nil, zero value otherwise.

### GetLastEventDateOk

`func (o *PostProfileResponseDataAttributes) GetLastEventDateOk() (*time.Time, bool)`

GetLastEventDateOk returns a tuple with the LastEventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventDate

`func (o *PostProfileResponseDataAttributes) SetLastEventDate(v time.Time)`

SetLastEventDate sets LastEventDate field to given value.

### HasLastEventDate

`func (o *PostProfileResponseDataAttributes) HasLastEventDate() bool`

HasLastEventDate returns a boolean if a field has been set.

### SetLastEventDateNil

`func (o *PostProfileResponseDataAttributes) SetLastEventDateNil(b bool)`

 SetLastEventDateNil sets the value for LastEventDate to be an explicit nil

### UnsetLastEventDate
`func (o *PostProfileResponseDataAttributes) UnsetLastEventDate()`

UnsetLastEventDate ensures that no value is present for LastEventDate, not even an explicit nil
### GetLocation

`func (o *PostProfileResponseDataAttributes) GetLocation() ProfileLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PostProfileResponseDataAttributes) GetLocationOk() (*ProfileLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PostProfileResponseDataAttributes) SetLocation(v ProfileLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PostProfileResponseDataAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProperties

`func (o *PostProfileResponseDataAttributes) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PostProfileResponseDataAttributes) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PostProfileResponseDataAttributes) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PostProfileResponseDataAttributes) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *PostProfileResponseDataAttributes) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *PostProfileResponseDataAttributes) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetSubscriptions

`func (o *PostProfileResponseDataAttributes) GetSubscriptions() Subscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *PostProfileResponseDataAttributes) GetSubscriptionsOk() (*Subscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *PostProfileResponseDataAttributes) SetSubscriptions(v Subscriptions)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *PostProfileResponseDataAttributes) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetPredictiveAnalytics

`func (o *PostProfileResponseDataAttributes) GetPredictiveAnalytics() PredictiveAnalytics`

GetPredictiveAnalytics returns the PredictiveAnalytics field if non-nil, zero value otherwise.

### GetPredictiveAnalyticsOk

`func (o *PostProfileResponseDataAttributes) GetPredictiveAnalyticsOk() (*PredictiveAnalytics, bool)`

GetPredictiveAnalyticsOk returns a tuple with the PredictiveAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictiveAnalytics

`func (o *PostProfileResponseDataAttributes) SetPredictiveAnalytics(v PredictiveAnalytics)`

SetPredictiveAnalytics sets PredictiveAnalytics field to given value.

### HasPredictiveAnalytics

`func (o *PostProfileResponseDataAttributes) HasPredictiveAnalytics() bool`

HasPredictiveAnalytics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


