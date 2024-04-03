# CatalogCategoryResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **NullableString** | The ID of the catalog category in an external system. | [optional] 
**Name** | Pointer to **NullableString** | The name of the catalog category. | [optional] 
**Updated** | Pointer to **NullableTime** | Date and time when the catalog category was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 

## Methods

### NewCatalogCategoryResponseObjectResourceAttributes

`func NewCatalogCategoryResponseObjectResourceAttributes() *CatalogCategoryResponseObjectResourceAttributes`

NewCatalogCategoryResponseObjectResourceAttributes instantiates a new CatalogCategoryResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogCategoryResponseObjectResourceAttributesWithDefaults

`func NewCatalogCategoryResponseObjectResourceAttributesWithDefaults() *CatalogCategoryResponseObjectResourceAttributes`

NewCatalogCategoryResponseObjectResourceAttributesWithDefaults instantiates a new CatalogCategoryResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CatalogCategoryResponseObjectResourceAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CatalogCategoryResponseObjectResourceAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CatalogCategoryResponseObjectResourceAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CatalogCategoryResponseObjectResourceAttributes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CatalogCategoryResponseObjectResourceAttributes) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CatalogCategoryResponseObjectResourceAttributes) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetName

`func (o *CatalogCategoryResponseObjectResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogCategoryResponseObjectResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogCategoryResponseObjectResourceAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogCategoryResponseObjectResourceAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CatalogCategoryResponseObjectResourceAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CatalogCategoryResponseObjectResourceAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUpdated

`func (o *CatalogCategoryResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CatalogCategoryResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CatalogCategoryResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CatalogCategoryResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *CatalogCategoryResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *CatalogCategoryResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


