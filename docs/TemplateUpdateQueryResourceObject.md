# TemplateUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TemplateEnum**](TemplateEnum.md) |  | 
**Id** | **string** | The ID of template | 
**Attributes** | [**TemplateUpdateQueryResourceObjectAttributes**](TemplateUpdateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewTemplateUpdateQueryResourceObject

`func NewTemplateUpdateQueryResourceObject(type_ TemplateEnum, id string, attributes TemplateUpdateQueryResourceObjectAttributes, ) *TemplateUpdateQueryResourceObject`

NewTemplateUpdateQueryResourceObject instantiates a new TemplateUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateUpdateQueryResourceObjectWithDefaults

`func NewTemplateUpdateQueryResourceObjectWithDefaults() *TemplateUpdateQueryResourceObject`

NewTemplateUpdateQueryResourceObjectWithDefaults instantiates a new TemplateUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateUpdateQueryResourceObject) GetType() TemplateEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateUpdateQueryResourceObject) GetTypeOk() (*TemplateEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateUpdateQueryResourceObject) SetType(v TemplateEnum)`

SetType sets Type field to given value.


### GetId

`func (o *TemplateUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *TemplateUpdateQueryResourceObject) GetAttributes() TemplateUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TemplateUpdateQueryResourceObject) GetAttributesOk() (*TemplateUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TemplateUpdateQueryResourceObject) SetAttributes(v TemplateUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


