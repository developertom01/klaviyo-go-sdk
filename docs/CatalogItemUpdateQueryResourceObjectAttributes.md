# CatalogItemUpdateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | The title of the catalog item. | [optional] 
**Price** | Pointer to **NullableFloat32** | This field can be used to set the price on the catalog item, which is what gets displayed for the item when included in emails. For most price-update use cases, you will also want to update the &#x60;price&#x60; on any child variants, using the [Update Catalog Variant Endpoint](https://developers.klaviyo.com/en/reference/update_catalog_variant). | [optional] 
**Description** | Pointer to **NullableString** | A description of the catalog item. | [optional] 
**Url** | Pointer to **NullableString** | URL pointing to the location of the catalog item on your website. | [optional] 
**ImageFullUrl** | Pointer to **NullableString** | URL pointing to the location of a full image of the catalog item. | [optional] 
**ImageThumbnailUrl** | Pointer to **NullableString** | URL pointing to the location of an image thumbnail of the catalog item | [optional] 
**Images** | Pointer to **[]string** | List of URLs pointing to the locations of images of the catalog item. | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | Flat JSON blob to provide custom metadata about the catalog item. May not exceed 100kb. | [optional] 
**Published** | Pointer to **NullableBool** | Boolean value indicating whether the catalog item is published. | [optional] 

## Methods

### NewCatalogItemUpdateQueryResourceObjectAttributes

`func NewCatalogItemUpdateQueryResourceObjectAttributes() *CatalogItemUpdateQueryResourceObjectAttributes`

NewCatalogItemUpdateQueryResourceObjectAttributes instantiates a new CatalogItemUpdateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemUpdateQueryResourceObjectAttributesWithDefaults

`func NewCatalogItemUpdateQueryResourceObjectAttributesWithDefaults() *CatalogItemUpdateQueryResourceObjectAttributes`

NewCatalogItemUpdateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogItemUpdateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CatalogItemUpdateQueryResourceObjectAttributes) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPrice

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CatalogItemUpdateQueryResourceObjectAttributes) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetDescription

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogItemUpdateQueryResourceObjectAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CatalogItemUpdateQueryResourceObjectAttributes) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetImageFullUrl

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetImageFullUrl() string`

GetImageFullUrl returns the ImageFullUrl field if non-nil, zero value otherwise.

### GetImageFullUrlOk

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetImageFullUrlOk() (*string, bool)`

GetImageFullUrlOk returns a tuple with the ImageFullUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFullUrl

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetImageFullUrl(v string)`

SetImageFullUrl sets ImageFullUrl field to given value.

### HasImageFullUrl

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) HasImageFullUrl() bool`

HasImageFullUrl returns a boolean if a field has been set.

### SetImageFullUrlNil

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetImageFullUrlNil(b bool)`

 SetImageFullUrlNil sets the value for ImageFullUrl to be an explicit nil

### UnsetImageFullUrl
`func (o *CatalogItemUpdateQueryResourceObjectAttributes) UnsetImageFullUrl()`

UnsetImageFullUrl ensures that no value is present for ImageFullUrl, not even an explicit nil
### GetImageThumbnailUrl

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetImageThumbnailUrl() string`

GetImageThumbnailUrl returns the ImageThumbnailUrl field if non-nil, zero value otherwise.

### GetImageThumbnailUrlOk

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetImageThumbnailUrlOk() (*string, bool)`

GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageThumbnailUrl

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetImageThumbnailUrl(v string)`

SetImageThumbnailUrl sets ImageThumbnailUrl field to given value.

### HasImageThumbnailUrl

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) HasImageThumbnailUrl() bool`

HasImageThumbnailUrl returns a boolean if a field has been set.

### SetImageThumbnailUrlNil

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetImageThumbnailUrlNil(b bool)`

 SetImageThumbnailUrlNil sets the value for ImageThumbnailUrl to be an explicit nil

### UnsetImageThumbnailUrl
`func (o *CatalogItemUpdateQueryResourceObjectAttributes) UnsetImageThumbnailUrl()`

UnsetImageThumbnailUrl ensures that no value is present for ImageThumbnailUrl, not even an explicit nil
### GetImages

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CatalogItemUpdateQueryResourceObjectAttributes) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetCustomMetadata

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### SetCustomMetadataNil

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetCustomMetadataNil(b bool)`

 SetCustomMetadataNil sets the value for CustomMetadata to be an explicit nil

### UnsetCustomMetadata
`func (o *CatalogItemUpdateQueryResourceObjectAttributes) UnsetCustomMetadata()`

UnsetCustomMetadata ensures that no value is present for CustomMetadata, not even an explicit nil
### GetPublished

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *CatalogItemUpdateQueryResourceObjectAttributes) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *CatalogItemUpdateQueryResourceObjectAttributes) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


