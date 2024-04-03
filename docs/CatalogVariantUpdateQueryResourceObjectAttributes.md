# CatalogVariantUpdateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewCatalogVariantUpdateQueryResourceObjectAttributes

`func NewCatalogVariantUpdateQueryResourceObjectAttributes() *CatalogVariantUpdateQueryResourceObjectAttributes`

NewCatalogVariantUpdateQueryResourceObjectAttributes instantiates a new CatalogVariantUpdateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogVariantUpdateQueryResourceObjectAttributesWithDefaults

`func NewCatalogVariantUpdateQueryResourceObjectAttributesWithDefaults() *CatalogVariantUpdateQueryResourceObjectAttributes`

NewCatalogVariantUpdateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogVariantUpdateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSku

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetInventoryPolicy

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetInventoryPolicy() int32`

GetInventoryPolicy returns the InventoryPolicy field if non-nil, zero value otherwise.

### GetInventoryPolicyOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetInventoryPolicyOk() (*int32, bool)`

GetInventoryPolicyOk returns a tuple with the InventoryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPolicy

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetInventoryPolicy(v int32)`

SetInventoryPolicy sets InventoryPolicy field to given value.

### HasInventoryPolicy

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasInventoryPolicy() bool`

HasInventoryPolicy returns a boolean if a field has been set.

### SetInventoryPolicyNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetInventoryPolicyNil(b bool)`

 SetInventoryPolicyNil sets the value for InventoryPolicy to be an explicit nil

### UnsetInventoryPolicy
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetInventoryPolicy()`

UnsetInventoryPolicy ensures that no value is present for InventoryPolicy, not even an explicit nil
### GetInventoryQuantity

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetInventoryQuantity() float32`

GetInventoryQuantity returns the InventoryQuantity field if non-nil, zero value otherwise.

### GetInventoryQuantityOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetInventoryQuantityOk() (*float32, bool)`

GetInventoryQuantityOk returns a tuple with the InventoryQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryQuantity

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetInventoryQuantity(v float32)`

SetInventoryQuantity sets InventoryQuantity field to given value.

### HasInventoryQuantity

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasInventoryQuantity() bool`

HasInventoryQuantity returns a boolean if a field has been set.

### SetInventoryQuantityNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetInventoryQuantityNil(b bool)`

 SetInventoryQuantityNil sets the value for InventoryQuantity to be an explicit nil

### UnsetInventoryQuantity
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetInventoryQuantity()`

UnsetInventoryQuantity ensures that no value is present for InventoryQuantity, not even an explicit nil
### GetPrice

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetUrl

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetImageFullUrl

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetImageFullUrl() string`

GetImageFullUrl returns the ImageFullUrl field if non-nil, zero value otherwise.

### GetImageFullUrlOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetImageFullUrlOk() (*string, bool)`

GetImageFullUrlOk returns a tuple with the ImageFullUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFullUrl

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetImageFullUrl(v string)`

SetImageFullUrl sets ImageFullUrl field to given value.

### HasImageFullUrl

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasImageFullUrl() bool`

HasImageFullUrl returns a boolean if a field has been set.

### SetImageFullUrlNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetImageFullUrlNil(b bool)`

 SetImageFullUrlNil sets the value for ImageFullUrl to be an explicit nil

### UnsetImageFullUrl
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetImageFullUrl()`

UnsetImageFullUrl ensures that no value is present for ImageFullUrl, not even an explicit nil
### GetImageThumbnailUrl

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetImageThumbnailUrl() string`

GetImageThumbnailUrl returns the ImageThumbnailUrl field if non-nil, zero value otherwise.

### GetImageThumbnailUrlOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetImageThumbnailUrlOk() (*string, bool)`

GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageThumbnailUrl

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetImageThumbnailUrl(v string)`

SetImageThumbnailUrl sets ImageThumbnailUrl field to given value.

### HasImageThumbnailUrl

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasImageThumbnailUrl() bool`

HasImageThumbnailUrl returns a boolean if a field has been set.

### SetImageThumbnailUrlNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetImageThumbnailUrlNil(b bool)`

 SetImageThumbnailUrlNil sets the value for ImageThumbnailUrl to be an explicit nil

### UnsetImageThumbnailUrl
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetImageThumbnailUrl()`

UnsetImageThumbnailUrl ensures that no value is present for ImageThumbnailUrl, not even an explicit nil
### GetImages

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetCustomMetadata

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### SetCustomMetadataNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetCustomMetadataNil(b bool)`

 SetCustomMetadataNil sets the value for CustomMetadata to be an explicit nil

### UnsetCustomMetadata
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetCustomMetadata()`

UnsetCustomMetadata ensures that no value is present for CustomMetadata, not even an explicit nil
### GetPublished

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *CatalogVariantUpdateQueryResourceObjectAttributes) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


