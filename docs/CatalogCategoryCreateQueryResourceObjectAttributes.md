# CatalogCategoryCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The ID of the catalog category in an external system. | 
**Name** | **string** | The name of the catalog category. | 
**IntegrationType** | Pointer to **NullableString** | The integration type. Currently only \&quot;$custom\&quot; is supported. | [optional] [default to "$custom"]
**CatalogType** | Pointer to **NullableString** | The type of catalog. Currently only \&quot;$default\&quot; is supported. | [optional] [default to "$default"]

## Methods

### NewCatalogCategoryCreateQueryResourceObjectAttributes

`func NewCatalogCategoryCreateQueryResourceObjectAttributes(externalId string, name string, ) *CatalogCategoryCreateQueryResourceObjectAttributes`

NewCatalogCategoryCreateQueryResourceObjectAttributes instantiates a new CatalogCategoryCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogCategoryCreateQueryResourceObjectAttributesWithDefaults

`func NewCatalogCategoryCreateQueryResourceObjectAttributesWithDefaults() *CatalogCategoryCreateQueryResourceObjectAttributes`

NewCatalogCategoryCreateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogCategoryCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetName

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetIntegrationType

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### SetIntegrationTypeNil

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) SetIntegrationTypeNil(b bool)`

 SetIntegrationTypeNil sets the value for IntegrationType to be an explicit nil

### UnsetIntegrationType
`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) UnsetIntegrationType()`

UnsetIntegrationType ensures that no value is present for IntegrationType, not even an explicit nil
### GetCatalogType

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) GetCatalogType() string`

GetCatalogType returns the CatalogType field if non-nil, zero value otherwise.

### GetCatalogTypeOk

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) GetCatalogTypeOk() (*string, bool)`

GetCatalogTypeOk returns a tuple with the CatalogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogType

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) SetCatalogType(v string)`

SetCatalogType sets CatalogType field to given value.

### HasCatalogType

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) HasCatalogType() bool`

HasCatalogType returns a boolean if a field has been set.

### SetCatalogTypeNil

`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) SetCatalogTypeNil(b bool)`

 SetCatalogTypeNil sets the value for CatalogType to be an explicit nil

### UnsetCatalogType
`func (o *CatalogCategoryCreateQueryResourceObjectAttributes) UnsetCatalogType()`

UnsetCatalogType ensures that no value is present for CatalogType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


