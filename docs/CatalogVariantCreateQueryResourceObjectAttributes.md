# CatalogVariantCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The ID of the catalog item variant in an external system. | 
**CatalogType** | Pointer to **NullableString** | The type of catalog. Currently only \&quot;$default\&quot; is supported. | [optional] [default to "$default"]
**IntegrationType** | Pointer to **NullableString** | The integration type. Currently only \&quot;$custom\&quot; is supported. | [optional] [default to "$custom"]
**Title** | **string** | The title of the catalog item variant. | 
**Description** | **string** | A description of the catalog item variant. | 
**Sku** | **string** | The SKU of the catalog item variant. | 
**InventoryPolicy** | Pointer to **NullableInt32** | This field controls the visibility of this catalog item variant in product feeds/blocks. This field supports the following values: &#x60;1&#x60;: a product will not appear in dynamic product recommendation feeds and blocks if it is out of stock. &#x60;0&#x60; or &#x60;2&#x60;: a product can appear in dynamic product recommendation feeds and blocks regardless of inventory quantity. | [optional] [default to 0]
**InventoryQuantity** | **float32** | The quantity of the catalog item variant currently in stock. | 
**Price** | **float32** | This field can be used to set the price on the catalog item variant, which is what gets displayed for the item variant when included in emails. For most price-update use cases, you will also want to update the &#x60;price&#x60; on any parent items using the [Update Catalog Item Endpoint](https://developers.klaviyo.com/en/reference/update_catalog_item). | 
**Url** | **string** | URL pointing to the location of the catalog item variant on your website. | 
**ImageFullUrl** | Pointer to **NullableString** | URL pointing to the location of a full image of the catalog item variant. | [optional] 
**ImageThumbnailUrl** | Pointer to **NullableString** | URL pointing to the location of an image thumbnail of the catalog item variant. | [optional] 
**Images** | Pointer to **[]string** | List of URLs pointing to the locations of images of the catalog item variant. | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | Flat JSON blob to provide custom metadata about the catalog item variant. May not exceed 100kb. | [optional] 
**Published** | Pointer to **NullableBool** | Boolean value indicating whether the catalog item variant is published. | [optional] [default to true]

## Methods

### NewCatalogVariantCreateQueryResourceObjectAttributes

`func NewCatalogVariantCreateQueryResourceObjectAttributes(externalId string, title string, description string, sku string, inventoryQuantity float32, price float32, url string, ) *CatalogVariantCreateQueryResourceObjectAttributes`

NewCatalogVariantCreateQueryResourceObjectAttributes instantiates a new CatalogVariantCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogVariantCreateQueryResourceObjectAttributesWithDefaults

`func NewCatalogVariantCreateQueryResourceObjectAttributesWithDefaults() *CatalogVariantCreateQueryResourceObjectAttributes`

NewCatalogVariantCreateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogVariantCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCatalogType

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetCatalogType() string`

GetCatalogType returns the CatalogType field if non-nil, zero value otherwise.

### GetCatalogTypeOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetCatalogTypeOk() (*string, bool)`

GetCatalogTypeOk returns a tuple with the CatalogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogType

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetCatalogType(v string)`

SetCatalogType sets CatalogType field to given value.

### HasCatalogType

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) HasCatalogType() bool`

HasCatalogType returns a boolean if a field has been set.

### SetCatalogTypeNil

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetCatalogTypeNil(b bool)`

 SetCatalogTypeNil sets the value for CatalogType to be an explicit nil

### UnsetCatalogType
`func (o *CatalogVariantCreateQueryResourceObjectAttributes) UnsetCatalogType()`

UnsetCatalogType ensures that no value is present for CatalogType, not even an explicit nil
### GetIntegrationType

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### SetIntegrationTypeNil

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetIntegrationTypeNil(b bool)`

 SetIntegrationTypeNil sets the value for IntegrationType to be an explicit nil

### UnsetIntegrationType
`func (o *CatalogVariantCreateQueryResourceObjectAttributes) UnsetIntegrationType()`

UnsetIntegrationType ensures that no value is present for IntegrationType, not even an explicit nil
### GetTitle

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSku

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetSku(v string)`

SetSku sets Sku field to given value.


### GetInventoryPolicy

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetInventoryPolicy() int32`

GetInventoryPolicy returns the InventoryPolicy field if non-nil, zero value otherwise.

### GetInventoryPolicyOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetInventoryPolicyOk() (*int32, bool)`

GetInventoryPolicyOk returns a tuple with the InventoryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPolicy

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetInventoryPolicy(v int32)`

SetInventoryPolicy sets InventoryPolicy field to given value.

### HasInventoryPolicy

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) HasInventoryPolicy() bool`

HasInventoryPolicy returns a boolean if a field has been set.

### SetInventoryPolicyNil

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetInventoryPolicyNil(b bool)`

 SetInventoryPolicyNil sets the value for InventoryPolicy to be an explicit nil

### UnsetInventoryPolicy
`func (o *CatalogVariantCreateQueryResourceObjectAttributes) UnsetInventoryPolicy()`

UnsetInventoryPolicy ensures that no value is present for InventoryPolicy, not even an explicit nil
### GetInventoryQuantity

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetInventoryQuantity() float32`

GetInventoryQuantity returns the InventoryQuantity field if non-nil, zero value otherwise.

### GetInventoryQuantityOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetInventoryQuantityOk() (*float32, bool)`

GetInventoryQuantityOk returns a tuple with the InventoryQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryQuantity

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetInventoryQuantity(v float32)`

SetInventoryQuantity sets InventoryQuantity field to given value.


### GetPrice

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetUrl

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetImageFullUrl

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetImageFullUrl() string`

GetImageFullUrl returns the ImageFullUrl field if non-nil, zero value otherwise.

### GetImageFullUrlOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetImageFullUrlOk() (*string, bool)`

GetImageFullUrlOk returns a tuple with the ImageFullUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFullUrl

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetImageFullUrl(v string)`

SetImageFullUrl sets ImageFullUrl field to given value.

### HasImageFullUrl

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) HasImageFullUrl() bool`

HasImageFullUrl returns a boolean if a field has been set.

### SetImageFullUrlNil

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetImageFullUrlNil(b bool)`

 SetImageFullUrlNil sets the value for ImageFullUrl to be an explicit nil

### UnsetImageFullUrl
`func (o *CatalogVariantCreateQueryResourceObjectAttributes) UnsetImageFullUrl()`

UnsetImageFullUrl ensures that no value is present for ImageFullUrl, not even an explicit nil
### GetImageThumbnailUrl

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetImageThumbnailUrl() string`

GetImageThumbnailUrl returns the ImageThumbnailUrl field if non-nil, zero value otherwise.

### GetImageThumbnailUrlOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetImageThumbnailUrlOk() (*string, bool)`

GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageThumbnailUrl

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetImageThumbnailUrl(v string)`

SetImageThumbnailUrl sets ImageThumbnailUrl field to given value.

### HasImageThumbnailUrl

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) HasImageThumbnailUrl() bool`

HasImageThumbnailUrl returns a boolean if a field has been set.

### SetImageThumbnailUrlNil

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetImageThumbnailUrlNil(b bool)`

 SetImageThumbnailUrlNil sets the value for ImageThumbnailUrl to be an explicit nil

### UnsetImageThumbnailUrl
`func (o *CatalogVariantCreateQueryResourceObjectAttributes) UnsetImageThumbnailUrl()`

UnsetImageThumbnailUrl ensures that no value is present for ImageThumbnailUrl, not even an explicit nil
### GetImages

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CatalogVariantCreateQueryResourceObjectAttributes) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetCustomMetadata

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### SetCustomMetadataNil

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetCustomMetadataNil(b bool)`

 SetCustomMetadataNil sets the value for CustomMetadata to be an explicit nil

### UnsetCustomMetadata
`func (o *CatalogVariantCreateQueryResourceObjectAttributes) UnsetCustomMetadata()`

UnsetCustomMetadata ensures that no value is present for CustomMetadata, not even an explicit nil
### GetPublished

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *CatalogVariantCreateQueryResourceObjectAttributes) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *CatalogVariantCreateQueryResourceObjectAttributes) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


