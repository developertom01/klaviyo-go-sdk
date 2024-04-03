# CampaignMessageResponseObjectResourceAttributesContent

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
**Body** | Pointer to **NullableString** | The message body | [optional] 
**MediaUrl** | Pointer to **NullableString** | URL for included media | [optional] 

## Methods

### NewCampaignMessageResponseObjectResourceAttributesContent

`func NewCampaignMessageResponseObjectResourceAttributesContent() *CampaignMessageResponseObjectResourceAttributesContent`

NewCampaignMessageResponseObjectResourceAttributesContent instantiates a new CampaignMessageResponseObjectResourceAttributesContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignMessageResponseObjectResourceAttributesContentWithDefaults

`func NewCampaignMessageResponseObjectResourceAttributesContentWithDefaults() *CampaignMessageResponseObjectResourceAttributesContent`

NewCampaignMessageResponseObjectResourceAttributesContentWithDefaults instantiates a new CampaignMessageResponseObjectResourceAttributesContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CampaignMessageResponseObjectResourceAttributesContent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *CampaignMessageResponseObjectResourceAttributesContent) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetPreviewText

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetPreviewTextOk() (*string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetPreviewText(v string)`

SetPreviewText sets PreviewText field to given value.

### HasPreviewText

`func (o *CampaignMessageResponseObjectResourceAttributesContent) HasPreviewText() bool`

HasPreviewText returns a boolean if a field has been set.

### SetPreviewTextNil

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetPreviewTextNil(b bool)`

 SetPreviewTextNil sets the value for PreviewText to be an explicit nil

### UnsetPreviewText
`func (o *CampaignMessageResponseObjectResourceAttributesContent) UnsetPreviewText()`

UnsetPreviewText ensures that no value is present for PreviewText, not even an explicit nil
### GetFromEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### SetFromEmailNil

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetFromEmailNil(b bool)`

 SetFromEmailNil sets the value for FromEmail to be an explicit nil

### UnsetFromEmail
`func (o *CampaignMessageResponseObjectResourceAttributesContent) UnsetFromEmail()`

UnsetFromEmail ensures that no value is present for FromEmail, not even an explicit nil
### GetFromLabel

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetFromLabel() string`

GetFromLabel returns the FromLabel field if non-nil, zero value otherwise.

### GetFromLabelOk

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetFromLabelOk() (*string, bool)`

GetFromLabelOk returns a tuple with the FromLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromLabel

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetFromLabel(v string)`

SetFromLabel sets FromLabel field to given value.

### HasFromLabel

`func (o *CampaignMessageResponseObjectResourceAttributesContent) HasFromLabel() bool`

HasFromLabel returns a boolean if a field has been set.

### SetFromLabelNil

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetFromLabelNil(b bool)`

 SetFromLabelNil sets the value for FromLabel to be an explicit nil

### UnsetFromLabel
`func (o *CampaignMessageResponseObjectResourceAttributesContent) UnsetFromLabel()`

UnsetFromLabel ensures that no value is present for FromLabel, not even an explicit nil
### GetReplyToEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetReplyToEmail() string`

GetReplyToEmail returns the ReplyToEmail field if non-nil, zero value otherwise.

### GetReplyToEmailOk

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetReplyToEmailOk() (*string, bool)`

GetReplyToEmailOk returns a tuple with the ReplyToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetReplyToEmail(v string)`

SetReplyToEmail sets ReplyToEmail field to given value.

### HasReplyToEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) HasReplyToEmail() bool`

HasReplyToEmail returns a boolean if a field has been set.

### SetReplyToEmailNil

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetReplyToEmailNil(b bool)`

 SetReplyToEmailNil sets the value for ReplyToEmail to be an explicit nil

### UnsetReplyToEmail
`func (o *CampaignMessageResponseObjectResourceAttributesContent) UnsetReplyToEmail()`

UnsetReplyToEmail ensures that no value is present for ReplyToEmail, not even an explicit nil
### GetCcEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### SetCcEmailNil

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetCcEmailNil(b bool)`

 SetCcEmailNil sets the value for CcEmail to be an explicit nil

### UnsetCcEmail
`func (o *CampaignMessageResponseObjectResourceAttributesContent) UnsetCcEmail()`

UnsetCcEmail ensures that no value is present for CcEmail, not even an explicit nil
### GetBccEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetBccEmail() string`

GetBccEmail returns the BccEmail field if non-nil, zero value otherwise.

### GetBccEmailOk

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetBccEmailOk() (*string, bool)`

GetBccEmailOk returns a tuple with the BccEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetBccEmail(v string)`

SetBccEmail sets BccEmail field to given value.

### HasBccEmail

`func (o *CampaignMessageResponseObjectResourceAttributesContent) HasBccEmail() bool`

HasBccEmail returns a boolean if a field has been set.

### SetBccEmailNil

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetBccEmailNil(b bool)`

 SetBccEmailNil sets the value for BccEmail to be an explicit nil

### UnsetBccEmail
`func (o *CampaignMessageResponseObjectResourceAttributesContent) UnsetBccEmail()`

UnsetBccEmail ensures that no value is present for BccEmail, not even an explicit nil
### GetBody

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CampaignMessageResponseObjectResourceAttributesContent) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *CampaignMessageResponseObjectResourceAttributesContent) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetMediaUrl

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *CampaignMessageResponseObjectResourceAttributesContent) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *CampaignMessageResponseObjectResourceAttributesContent) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### SetMediaUrlNil

`func (o *CampaignMessageResponseObjectResourceAttributesContent) SetMediaUrlNil(b bool)`

 SetMediaUrlNil sets the value for MediaUrl to be an explicit nil

### UnsetMediaUrl
`func (o *CampaignMessageResponseObjectResourceAttributesContent) UnsetMediaUrl()`

UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


