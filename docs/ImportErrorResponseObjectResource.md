# ImportErrorResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ImportErrorEnum**](ImportErrorEnum.md) |  | 
**Id** | **string** | Unique identifier for the error. | 
**Attributes** | [**ImportErrorResponseObjectResourceAttributes**](ImportErrorResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewImportErrorResponseObjectResource

`func NewImportErrorResponseObjectResource(type_ ImportErrorEnum, id string, attributes ImportErrorResponseObjectResourceAttributes, links ObjectLinks, ) *ImportErrorResponseObjectResource`

NewImportErrorResponseObjectResource instantiates a new ImportErrorResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportErrorResponseObjectResourceWithDefaults

`func NewImportErrorResponseObjectResourceWithDefaults() *ImportErrorResponseObjectResource`

NewImportErrorResponseObjectResourceWithDefaults instantiates a new ImportErrorResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ImportErrorResponseObjectResource) GetType() ImportErrorEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportErrorResponseObjectResource) GetTypeOk() (*ImportErrorEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportErrorResponseObjectResource) SetType(v ImportErrorEnum)`

SetType sets Type field to given value.


### GetId

`func (o *ImportErrorResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportErrorResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportErrorResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *ImportErrorResponseObjectResource) GetAttributes() ImportErrorResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ImportErrorResponseObjectResource) GetAttributesOk() (*ImportErrorResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ImportErrorResponseObjectResource) SetAttributes(v ImportErrorResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *ImportErrorResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ImportErrorResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ImportErrorResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


