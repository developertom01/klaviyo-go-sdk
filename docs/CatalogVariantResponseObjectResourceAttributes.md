# CatalogVariantResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **NullableString** | The ID of the catalog item variant in an external system. | [optional] 
**Title** | Pointer to **NullableString** | The title of the catalog item variant. | [optional] 
**Description** | Pointer to **NullableString** | A description of the catalog item variant. | [optional] 
**Sku** | Pointer to **NullableString** | The SKU of the catalog item variant. | [optional] 
**InventoryPolicy** | Pointer to **NullableInt32** | This field controls the visibility of this catalog item variant in product feeds/blocks. This field supports the following values: &#x60;1&#x60;: a product will not appear in dynamic product recommendation feeds and blocks if it is out of stock. &#x60;0&#x60; or &#x60;2&#x60;: a product can appear in dynamic product recommendation feeds and blocks regardless of inventory quantity. | [optional] 
**InventoryQuantity** | Pointer to **NullableFloat32** | The quantity of the catalog item variant currently in stock. | [optional] 
**Price** | Pointer to **NullableFloat32** | This field can be used to set the price on the catalog item variant, which is what gets displayed for the item variant when included in emails. For most price-update use cases, you will also want to update the &#x60;price&#x60; on any parent items using the [Update Catalog Item Endpoint](https://developers.klaviyo.com/en/reference/update_catalog_item). | [optional] 
**Url** | Pointer to **NullableString** | URL pointing to the location of the catalog item variant on your website. | [optional] 
**ImageFullUrl** | Pointer to **NullableString** | URL pointing to the location of a full image of the catalog item variant. | [optional] 
**ImageThumbnailUrl** | Pointer to **NullableString** | URL pointing to the location of an image thumbnail of the catalog item variant. | [optional] 
**Images** | Pointer to **[]string** | List of URLs pointing to the locations of images of the catalog item variant. | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | Flat JSON blob to provide custom metadata about the catalog item variant. May not exceed 100kb. | [optional] 
**Published** | Pointer to **NullableBool** | Boolean value indicating whether the catalog item variant is published. | [optional] 
**Created** | Pointer to **NullableTime** | Date and time when the catalog item  variant was created, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 
**Updated** | Pointer to **NullableTime** | Date and time when the catalog item variant was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 

## Methods

### NewCatalogVariantResponseObjectResourceAttributes

`func NewCatalogVariantResponseObjectResourceAttributes() *CatalogVariantResponseObjectResourceAttributes`

NewCatalogVariantResponseObjectResourceAttributes instantiates a new CatalogVariantResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogVariantResponseObjectResourceAttributesWithDefaults

`func NewCatalogVariantResponseObjectResourceAttributesWithDefaults() *CatalogVariantResponseObjectResourceAttributes`

NewCatalogVariantResponseObjectResourceAttributesWithDefaults instantiates a new CatalogVariantResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CatalogVariantResponseObjectResourceAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CatalogVariantResponseObjectResourceAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CatalogVariantResponseObjectResourceAttributes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetTitle

`func (o *CatalogVariantResponseObjectResourceAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogVariantResponseObjectResourceAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CatalogVariantResponseObjectResourceAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CatalogVariantResponseObjectResourceAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogVariantResponseObjectResourceAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogVariantResponseObjectResourceAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSku

`func (o *CatalogVariantResponseObjectResourceAttributes) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CatalogVariantResponseObjectResourceAttributes) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CatalogVariantResponseObjectResourceAttributes) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetInventoryPolicy

`func (o *CatalogVariantResponseObjectResourceAttributes) GetInventoryPolicy() int32`

GetInventoryPolicy returns the InventoryPolicy field if non-nil, zero value otherwise.

### GetInventoryPolicyOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetInventoryPolicyOk() (*int32, bool)`

GetInventoryPolicyOk returns a tuple with the InventoryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPolicy

`func (o *CatalogVariantResponseObjectResourceAttributes) SetInventoryPolicy(v int32)`

SetInventoryPolicy sets InventoryPolicy field to given value.

### HasInventoryPolicy

`func (o *CatalogVariantResponseObjectResourceAttributes) HasInventoryPolicy() bool`

HasInventoryPolicy returns a boolean if a field has been set.

### SetInventoryPolicyNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetInventoryPolicyNil(b bool)`

 SetInventoryPolicyNil sets the value for InventoryPolicy to be an explicit nil

### UnsetInventoryPolicy
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetInventoryPolicy()`

UnsetInventoryPolicy ensures that no value is present for InventoryPolicy, not even an explicit nil
### GetInventoryQuantity

`func (o *CatalogVariantResponseObjectResourceAttributes) GetInventoryQuantity() float32`

GetInventoryQuantity returns the InventoryQuantity field if non-nil, zero value otherwise.

### GetInventoryQuantityOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetInventoryQuantityOk() (*float32, bool)`

GetInventoryQuantityOk returns a tuple with the InventoryQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryQuantity

`func (o *CatalogVariantResponseObjectResourceAttributes) SetInventoryQuantity(v float32)`

SetInventoryQuantity sets InventoryQuantity field to given value.

### HasInventoryQuantity

`func (o *CatalogVariantResponseObjectResourceAttributes) HasInventoryQuantity() bool`

HasInventoryQuantity returns a boolean if a field has been set.

### SetInventoryQuantityNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetInventoryQuantityNil(b bool)`

 SetInventoryQuantityNil sets the value for InventoryQuantity to be an explicit nil

### UnsetInventoryQuantity
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetInventoryQuantity()`

UnsetInventoryQuantity ensures that no value is present for InventoryQuantity, not even an explicit nil
### GetPrice

`func (o *CatalogVariantResponseObjectResourceAttributes) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogVariantResponseObjectResourceAttributes) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogVariantResponseObjectResourceAttributes) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetUrl

`func (o *CatalogVariantResponseObjectResourceAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CatalogVariantResponseObjectResourceAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CatalogVariantResponseObjectResourceAttributes) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetImageFullUrl

`func (o *CatalogVariantResponseObjectResourceAttributes) GetImageFullUrl() string`

GetImageFullUrl returns the ImageFullUrl field if non-nil, zero value otherwise.

### GetImageFullUrlOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetImageFullUrlOk() (*string, bool)`

GetImageFullUrlOk returns a tuple with the ImageFullUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFullUrl

`func (o *CatalogVariantResponseObjectResourceAttributes) SetImageFullUrl(v string)`

SetImageFullUrl sets ImageFullUrl field to given value.

### HasImageFullUrl

`func (o *CatalogVariantResponseObjectResourceAttributes) HasImageFullUrl() bool`

HasImageFullUrl returns a boolean if a field has been set.

### SetImageFullUrlNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetImageFullUrlNil(b bool)`

 SetImageFullUrlNil sets the value for ImageFullUrl to be an explicit nil

### UnsetImageFullUrl
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetImageFullUrl()`

UnsetImageFullUrl ensures that no value is present for ImageFullUrl, not even an explicit nil
### GetImageThumbnailUrl

`func (o *CatalogVariantResponseObjectResourceAttributes) GetImageThumbnailUrl() string`

GetImageThumbnailUrl returns the ImageThumbnailUrl field if non-nil, zero value otherwise.

### GetImageThumbnailUrlOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetImageThumbnailUrlOk() (*string, bool)`

GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageThumbnailUrl

`func (o *CatalogVariantResponseObjectResourceAttributes) SetImageThumbnailUrl(v string)`

SetImageThumbnailUrl sets ImageThumbnailUrl field to given value.

### HasImageThumbnailUrl

`func (o *CatalogVariantResponseObjectResourceAttributes) HasImageThumbnailUrl() bool`

HasImageThumbnailUrl returns a boolean if a field has been set.

### SetImageThumbnailUrlNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetImageThumbnailUrlNil(b bool)`

 SetImageThumbnailUrlNil sets the value for ImageThumbnailUrl to be an explicit nil

### UnsetImageThumbnailUrl
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetImageThumbnailUrl()`

UnsetImageThumbnailUrl ensures that no value is present for ImageThumbnailUrl, not even an explicit nil
### GetImages

`func (o *CatalogVariantResponseObjectResourceAttributes) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CatalogVariantResponseObjectResourceAttributes) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *CatalogVariantResponseObjectResourceAttributes) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetCustomMetadata

`func (o *CatalogVariantResponseObjectResourceAttributes) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *CatalogVariantResponseObjectResourceAttributes) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *CatalogVariantResponseObjectResourceAttributes) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### SetCustomMetadataNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetCustomMetadataNil(b bool)`

 SetCustomMetadataNil sets the value for CustomMetadata to be an explicit nil

### UnsetCustomMetadata
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetCustomMetadata()`

UnsetCustomMetadata ensures that no value is present for CustomMetadata, not even an explicit nil
### GetPublished

`func (o *CatalogVariantResponseObjectResourceAttributes) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CatalogVariantResponseObjectResourceAttributes) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CatalogVariantResponseObjectResourceAttributes) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil
### GetCreated

`func (o *CatalogVariantResponseObjectResourceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CatalogVariantResponseObjectResourceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CatalogVariantResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *CatalogVariantResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CatalogVariantResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CatalogVariantResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CatalogVariantResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *CatalogVariantResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *CatalogVariantResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


