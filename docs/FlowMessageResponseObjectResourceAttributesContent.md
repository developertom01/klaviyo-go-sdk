# FlowMessageResponseObjectResourceAttributesContent

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
**Body** | **string** |  | 
**MediaUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFlowMessageResponseObjectResourceAttributesContent

`func NewFlowMessageResponseObjectResourceAttributesContent(subject string, previewText string, fromEmail string, fromLabel string, body string, ) *FlowMessageResponseObjectResourceAttributesContent`

NewFlowMessageResponseObjectResourceAttributesContent instantiates a new FlowMessageResponseObjectResourceAttributesContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowMessageResponseObjectResourceAttributesContentWithDefaults

`func NewFlowMessageResponseObjectResourceAttributesContentWithDefaults() *FlowMessageResponseObjectResourceAttributesContent`

NewFlowMessageResponseObjectResourceAttributesContentWithDefaults instantiates a new FlowMessageResponseObjectResourceAttributesContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetPreviewText

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetPreviewTextOk() (*string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetPreviewText(v string)`

SetPreviewText sets PreviewText field to given value.


### GetFromEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.


### GetFromLabel

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetFromLabel() string`

GetFromLabel returns the FromLabel field if non-nil, zero value otherwise.

### GetFromLabelOk

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetFromLabelOk() (*string, bool)`

GetFromLabelOk returns a tuple with the FromLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromLabel

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetFromLabel(v string)`

SetFromLabel sets FromLabel field to given value.


### GetReplyToEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetReplyToEmail() string`

GetReplyToEmail returns the ReplyToEmail field if non-nil, zero value otherwise.

### GetReplyToEmailOk

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetReplyToEmailOk() (*string, bool)`

GetReplyToEmailOk returns a tuple with the ReplyToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetReplyToEmail(v string)`

SetReplyToEmail sets ReplyToEmail field to given value.

### HasReplyToEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) HasReplyToEmail() bool`

HasReplyToEmail returns a boolean if a field has been set.

### SetReplyToEmailNil

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetReplyToEmailNil(b bool)`

 SetReplyToEmailNil sets the value for ReplyToEmail to be an explicit nil

### UnsetReplyToEmail
`func (o *FlowMessageResponseObjectResourceAttributesContent) UnsetReplyToEmail()`

UnsetReplyToEmail ensures that no value is present for ReplyToEmail, not even an explicit nil
### GetCcEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### SetCcEmailNil

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetCcEmailNil(b bool)`

 SetCcEmailNil sets the value for CcEmail to be an explicit nil

### UnsetCcEmail
`func (o *FlowMessageResponseObjectResourceAttributesContent) UnsetCcEmail()`

UnsetCcEmail ensures that no value is present for CcEmail, not even an explicit nil
### GetBccEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetBccEmail() string`

GetBccEmail returns the BccEmail field if non-nil, zero value otherwise.

### GetBccEmailOk

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetBccEmailOk() (*string, bool)`

GetBccEmailOk returns a tuple with the BccEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetBccEmail(v string)`

SetBccEmail sets BccEmail field to given value.

### HasBccEmail

`func (o *FlowMessageResponseObjectResourceAttributesContent) HasBccEmail() bool`

HasBccEmail returns a boolean if a field has been set.

### SetBccEmailNil

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetBccEmailNil(b bool)`

 SetBccEmailNil sets the value for BccEmail to be an explicit nil

### UnsetBccEmail
`func (o *FlowMessageResponseObjectResourceAttributesContent) UnsetBccEmail()`

UnsetBccEmail ensures that no value is present for BccEmail, not even an explicit nil
### GetBody

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetBody(v string)`

SetBody sets Body field to given value.


### GetMediaUrl

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *FlowMessageResponseObjectResourceAttributesContent) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *FlowMessageResponseObjectResourceAttributesContent) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### SetMediaUrlNil

`func (o *FlowMessageResponseObjectResourceAttributesContent) SetMediaUrlNil(b bool)`

 SetMediaUrlNil sets the value for MediaUrl to be an explicit nil

### UnsetMediaUrl
`func (o *FlowMessageResponseObjectResourceAttributesContent) UnsetMediaUrl()`

UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


