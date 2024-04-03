# TemplateCloneQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TemplateEnum**](TemplateEnum.md) |  | 
**Id** | **string** | The ID of template to be cloned | 
**Attributes** | [**TemplateCloneQueryResourceObjectAttributes**](TemplateCloneQueryResourceObjectAttributes.md) |  | 

## Methods

### NewTemplateCloneQueryResourceObject

`func NewTemplateCloneQueryResourceObject(type_ TemplateEnum, id string, attributes TemplateCloneQueryResourceObjectAttributes, ) *TemplateCloneQueryResourceObject`

NewTemplateCloneQueryResourceObject instantiates a new TemplateCloneQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateCloneQueryResourceObjectWithDefaults

`func NewTemplateCloneQueryResourceObjectWithDefaults() *TemplateCloneQueryResourceObject`

NewTemplateCloneQueryResourceObjectWithDefaults instantiates a new TemplateCloneQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateCloneQueryResourceObject) GetType() TemplateEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateCloneQueryResourceObject) GetTypeOk() (*TemplateEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateCloneQueryResourceObject) SetType(v TemplateEnum)`

SetType sets Type field to given value.


### GetId

`func (o *TemplateCloneQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateCloneQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateCloneQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *TemplateCloneQueryResourceObject) GetAttributes() TemplateCloneQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TemplateCloneQueryResourceObject) GetAttributesOk() (*TemplateCloneQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TemplateCloneQueryResourceObject) SetAttributes(v TemplateCloneQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


