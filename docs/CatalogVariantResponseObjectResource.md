# CatalogVariantResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogVariantEnum**](CatalogVariantEnum.md) |  | 
**Id** | **string** | The catalog variant ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 
**Attributes** | [**CatalogVariantResponseObjectResourceAttributes**](CatalogVariantResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewCatalogVariantResponseObjectResource

`func NewCatalogVariantResponseObjectResource(type_ CatalogVariantEnum, id string, attributes CatalogVariantResponseObjectResourceAttributes, links ObjectLinks, ) *CatalogVariantResponseObjectResource`

NewCatalogVariantResponseObjectResource instantiates a new CatalogVariantResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogVariantResponseObjectResourceWithDefaults

`func NewCatalogVariantResponseObjectResourceWithDefaults() *CatalogVariantResponseObjectResource`

NewCatalogVariantResponseObjectResourceWithDefaults instantiates a new CatalogVariantResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogVariantResponseObjectResource) GetType() CatalogVariantEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogVariantResponseObjectResource) GetTypeOk() (*CatalogVariantEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogVariantResponseObjectResource) SetType(v CatalogVariantEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CatalogVariantResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogVariantResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogVariantResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CatalogVariantResponseObjectResource) GetAttributes() CatalogVariantResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogVariantResponseObjectResource) GetAttributesOk() (*CatalogVariantResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogVariantResponseObjectResource) SetAttributes(v CatalogVariantResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *CatalogVariantResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CatalogVariantResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CatalogVariantResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


