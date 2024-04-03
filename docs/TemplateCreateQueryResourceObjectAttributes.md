# TemplateCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the template | 
**EditorType** | **string** | Restricted to CODE | 
**Html** | Pointer to **NullableString** | The HTML contents of the template | [optional] 
**Text** | Pointer to **NullableString** | The plaintext version of the template | [optional] 

## Methods

### NewTemplateCreateQueryResourceObjectAttributes

`func NewTemplateCreateQueryResourceObjectAttributes(name string, editorType string, ) *TemplateCreateQueryResourceObjectAttributes`

NewTemplateCreateQueryResourceObjectAttributes instantiates a new TemplateCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateCreateQueryResourceObjectAttributesWithDefaults

`func NewTemplateCreateQueryResourceObjectAttributesWithDefaults() *TemplateCreateQueryResourceObjectAttributes`

NewTemplateCreateQueryResourceObjectAttributesWithDefaults instantiates a new TemplateCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateCreateQueryResourceObjectAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateCreateQueryResourceObjectAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateCreateQueryResourceObjectAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetEditorType

`func (o *TemplateCreateQueryResourceObjectAttributes) GetEditorType() string`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *TemplateCreateQueryResourceObjectAttributes) GetEditorTypeOk() (*string, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *TemplateCreateQueryResourceObjectAttributes) SetEditorType(v string)`

SetEditorType sets EditorType field to given value.


### GetHtml

`func (o *TemplateCreateQueryResourceObjectAttributes) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *TemplateCreateQueryResourceObjectAttributes) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *TemplateCreateQueryResourceObjectAttributes) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *TemplateCreateQueryResourceObjectAttributes) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### SetHtmlNil

`func (o *TemplateCreateQueryResourceObjectAttributes) SetHtmlNil(b bool)`

 SetHtmlNil sets the value for Html to be an explicit nil

### UnsetHtml
`func (o *TemplateCreateQueryResourceObjectAttributes) UnsetHtml()`

UnsetHtml ensures that no value is present for Html, not even an explicit nil
### GetText

`func (o *TemplateCreateQueryResourceObjectAttributes) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TemplateCreateQueryResourceObjectAttributes) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TemplateCreateQueryResourceObjectAttributes) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TemplateCreateQueryResourceObjectAttributes) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *TemplateCreateQueryResourceObjectAttributes) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *TemplateCreateQueryResourceObjectAttributes) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


