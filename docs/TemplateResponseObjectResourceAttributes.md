# TemplateResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the template | 
**EditorType** | **string** | &#x60;editor_type&#x60; has a fixed set of values: * SYSTEM_DRAGGABLE: indicates a drag-and-drop editor template * SIMPLE: A rich text editor template * CODE: A custom HTML template * USER_DRAGGABLE: A hybrid template, using custom HTML in the drag-and-drop editor | 
**Html** | **string** | The rendered HTML of the template | 
**Text** | Pointer to **NullableString** | The template plain_text | [optional] 
**Created** | Pointer to **NullableTime** | The date the template was created in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 
**Updated** | Pointer to **NullableTime** | The date the template was updated in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 

## Methods

### NewTemplateResponseObjectResourceAttributes

`func NewTemplateResponseObjectResourceAttributes(name string, editorType string, html string, ) *TemplateResponseObjectResourceAttributes`

NewTemplateResponseObjectResourceAttributes instantiates a new TemplateResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateResponseObjectResourceAttributesWithDefaults

`func NewTemplateResponseObjectResourceAttributesWithDefaults() *TemplateResponseObjectResourceAttributes`

NewTemplateResponseObjectResourceAttributesWithDefaults instantiates a new TemplateResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateResponseObjectResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateResponseObjectResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateResponseObjectResourceAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetEditorType

`func (o *TemplateResponseObjectResourceAttributes) GetEditorType() string`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *TemplateResponseObjectResourceAttributes) GetEditorTypeOk() (*string, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *TemplateResponseObjectResourceAttributes) SetEditorType(v string)`

SetEditorType sets EditorType field to given value.


### GetHtml

`func (o *TemplateResponseObjectResourceAttributes) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *TemplateResponseObjectResourceAttributes) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *TemplateResponseObjectResourceAttributes) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetText

`func (o *TemplateResponseObjectResourceAttributes) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TemplateResponseObjectResourceAttributes) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TemplateResponseObjectResourceAttributes) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TemplateResponseObjectResourceAttributes) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *TemplateResponseObjectResourceAttributes) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *TemplateResponseObjectResourceAttributes) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetCreated

`func (o *TemplateResponseObjectResourceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TemplateResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TemplateResponseObjectResourceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TemplateResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *TemplateResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *TemplateResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *TemplateResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TemplateResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TemplateResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *TemplateResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *TemplateResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *TemplateResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


