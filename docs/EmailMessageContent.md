# EmailMessageContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** |  | 
**PreviewText** | **string** |  | 
**FromEmail** | **string** |  | 
**FromLabel** | **string** |  | 
**ReplyToEmail** | Pointer to **NullableString** |  | [optional] 
**CcEmail** | Pointer to **NullableString** |  | [optional] 
**BccEmail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailMessageContent

`func NewEmailMessageContent(subject string, previewText string, fromEmail string, fromLabel string, ) *EmailMessageContent`

NewEmailMessageContent instantiates a new EmailMessageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMessageContentWithDefaults

`func NewEmailMessageContentWithDefaults() *EmailMessageContent`

NewEmailMessageContentWithDefaults instantiates a new EmailMessageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *EmailMessageContent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailMessageContent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailMessageContent) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetPreviewText

`func (o *EmailMessageContent) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *EmailMessageContent) GetPreviewTextOk() (*string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *EmailMessageContent) SetPreviewText(v string)`

SetPreviewText sets PreviewText field to given value.


### GetFromEmail

`func (o *EmailMessageContent) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *EmailMessageContent) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *EmailMessageContent) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.


### GetFromLabel

`func (o *EmailMessageContent) GetFromLabel() string`

GetFromLabel returns the FromLabel field if non-nil, zero value otherwise.

### GetFromLabelOk

`func (o *EmailMessageContent) GetFromLabelOk() (*string, bool)`

GetFromLabelOk returns a tuple with the FromLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromLabel

`func (o *EmailMessageContent) SetFromLabel(v string)`

SetFromLabel sets FromLabel field to given value.


### GetReplyToEmail

`func (o *EmailMessageContent) GetReplyToEmail() string`

GetReplyToEmail returns the ReplyToEmail field if non-nil, zero value otherwise.

### GetReplyToEmailOk

`func (o *EmailMessageContent) GetReplyToEmailOk() (*string, bool)`

GetReplyToEmailOk returns a tuple with the ReplyToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmail

`func (o *EmailMessageContent) SetReplyToEmail(v string)`

SetReplyToEmail sets ReplyToEmail field to given value.

### HasReplyToEmail

`func (o *EmailMessageContent) HasReplyToEmail() bool`

HasReplyToEmail returns a boolean if a field has been set.

### SetReplyToEmailNil

`func (o *EmailMessageContent) SetReplyToEmailNil(b bool)`

 SetReplyToEmailNil sets the value for ReplyToEmail to be an explicit nil

### UnsetReplyToEmail
`func (o *EmailMessageContent) UnsetReplyToEmail()`

UnsetReplyToEmail ensures that no value is present for ReplyToEmail, not even an explicit nil
### GetCcEmail

`func (o *EmailMessageContent) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *EmailMessageContent) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *EmailMessageContent) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *EmailMessageContent) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### SetCcEmailNil

`func (o *EmailMessageContent) SetCcEmailNil(b bool)`

 SetCcEmailNil sets the value for CcEmail to be an explicit nil

### UnsetCcEmail
`func (o *EmailMessageContent) UnsetCcEmail()`

UnsetCcEmail ensures that no value is present for CcEmail, not even an explicit nil
### GetBccEmail

`func (o *EmailMessageContent) GetBccEmail() string`

GetBccEmail returns the BccEmail field if non-nil, zero value otherwise.

### GetBccEmailOk

`func (o *EmailMessageContent) GetBccEmailOk() (*string, bool)`

GetBccEmailOk returns a tuple with the BccEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccEmail

`func (o *EmailMessageContent) SetBccEmail(v string)`

SetBccEmail sets BccEmail field to given value.

### HasBccEmail

`func (o *EmailMessageContent) HasBccEmail() bool`

HasBccEmail returns a boolean if a field has been set.

### SetBccEmailNil

`func (o *EmailMessageContent) SetBccEmailNil(b bool)`

 SetBccEmailNil sets the value for BccEmail to be an explicit nil

### UnsetBccEmail
`func (o *EmailMessageContent) UnsetBccEmail()`

UnsetBccEmail ensures that no value is present for BccEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


