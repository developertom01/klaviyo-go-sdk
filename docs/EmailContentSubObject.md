# EmailContentSubObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **NullableString** | The subject of the message | [optional] 
**PreviewText** | Pointer to **NullableString** | Preview text associated with the message | [optional] 
**FromEmail** | Pointer to **NullableString** | The email the message should be sent from | [optional] 
**FromLabel** | Pointer to **NullableString** | The label associated with the from_email | [optional] 
**ReplyToEmail** | Pointer to **NullableString** | Optional Reply-To email address | [optional] 
**CcEmail** | Pointer to **NullableString** | Optional CC email address | [optional] 
**BccEmail** | Pointer to **NullableString** | Optional BCC email address | [optional] 

## Methods

### NewEmailContentSubObject

`func NewEmailContentSubObject() *EmailContentSubObject`

NewEmailContentSubObject instantiates a new EmailContentSubObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailContentSubObjectWithDefaults

`func NewEmailContentSubObjectWithDefaults() *EmailContentSubObject`

NewEmailContentSubObjectWithDefaults instantiates a new EmailContentSubObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *EmailContentSubObject) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailContentSubObject) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailContentSubObject) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailContentSubObject) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *EmailContentSubObject) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *EmailContentSubObject) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetPreviewText

`func (o *EmailContentSubObject) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *EmailContentSubObject) GetPreviewTextOk() (*string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *EmailContentSubObject) SetPreviewText(v string)`

SetPreviewText sets PreviewText field to given value.

### HasPreviewText

`func (o *EmailContentSubObject) HasPreviewText() bool`

HasPreviewText returns a boolean if a field has been set.

### SetPreviewTextNil

`func (o *EmailContentSubObject) SetPreviewTextNil(b bool)`

 SetPreviewTextNil sets the value for PreviewText to be an explicit nil

### UnsetPreviewText
`func (o *EmailContentSubObject) UnsetPreviewText()`

UnsetPreviewText ensures that no value is present for PreviewText, not even an explicit nil
### GetFromEmail

`func (o *EmailContentSubObject) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *EmailContentSubObject) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *EmailContentSubObject) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *EmailContentSubObject) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### SetFromEmailNil

`func (o *EmailContentSubObject) SetFromEmailNil(b bool)`

 SetFromEmailNil sets the value for FromEmail to be an explicit nil

### UnsetFromEmail
`func (o *EmailContentSubObject) UnsetFromEmail()`

UnsetFromEmail ensures that no value is present for FromEmail, not even an explicit nil
### GetFromLabel

`func (o *EmailContentSubObject) GetFromLabel() string`

GetFromLabel returns the FromLabel field if non-nil, zero value otherwise.

### GetFromLabelOk

`func (o *EmailContentSubObject) GetFromLabelOk() (*string, bool)`

GetFromLabelOk returns a tuple with the FromLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromLabel

`func (o *EmailContentSubObject) SetFromLabel(v string)`

SetFromLabel sets FromLabel field to given value.

### HasFromLabel

`func (o *EmailContentSubObject) HasFromLabel() bool`

HasFromLabel returns a boolean if a field has been set.

### SetFromLabelNil

`func (o *EmailContentSubObject) SetFromLabelNil(b bool)`

 SetFromLabelNil sets the value for FromLabel to be an explicit nil

### UnsetFromLabel
`func (o *EmailContentSubObject) UnsetFromLabel()`

UnsetFromLabel ensures that no value is present for FromLabel, not even an explicit nil
### GetReplyToEmail

`func (o *EmailContentSubObject) GetReplyToEmail() string`

GetReplyToEmail returns the ReplyToEmail field if non-nil, zero value otherwise.

### GetReplyToEmailOk

`func (o *EmailContentSubObject) GetReplyToEmailOk() (*string, bool)`

GetReplyToEmailOk returns a tuple with the ReplyToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmail

`func (o *EmailContentSubObject) SetReplyToEmail(v string)`

SetReplyToEmail sets ReplyToEmail field to given value.

### HasReplyToEmail

`func (o *EmailContentSubObject) HasReplyToEmail() bool`

HasReplyToEmail returns a boolean if a field has been set.

### SetReplyToEmailNil

`func (o *EmailContentSubObject) SetReplyToEmailNil(b bool)`

 SetReplyToEmailNil sets the value for ReplyToEmail to be an explicit nil

### UnsetReplyToEmail
`func (o *EmailContentSubObject) UnsetReplyToEmail()`

UnsetReplyToEmail ensures that no value is present for ReplyToEmail, not even an explicit nil
### GetCcEmail

`func (o *EmailContentSubObject) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *EmailContentSubObject) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *EmailContentSubObject) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *EmailContentSubObject) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### SetCcEmailNil

`func (o *EmailContentSubObject) SetCcEmailNil(b bool)`

 SetCcEmailNil sets the value for CcEmail to be an explicit nil

### UnsetCcEmail
`func (o *EmailContentSubObject) UnsetCcEmail()`

UnsetCcEmail ensures that no value is present for CcEmail, not even an explicit nil
### GetBccEmail

`func (o *EmailContentSubObject) GetBccEmail() string`

GetBccEmail returns the BccEmail field if non-nil, zero value otherwise.

### GetBccEmailOk

`func (o *EmailContentSubObject) GetBccEmailOk() (*string, bool)`

GetBccEmailOk returns a tuple with the BccEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccEmail

`func (o *EmailContentSubObject) SetBccEmail(v string)`

SetBccEmail sets BccEmail field to given value.

### HasBccEmail

`func (o *EmailContentSubObject) HasBccEmail() bool`

HasBccEmail returns a boolean if a field has been set.

### SetBccEmailNil

`func (o *EmailContentSubObject) SetBccEmailNil(b bool)`

 SetBccEmailNil sets the value for BccEmail to be an explicit nil

### UnsetBccEmail
`func (o *EmailContentSubObject) UnsetBccEmail()`

UnsetBccEmail ensures that no value is present for BccEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


