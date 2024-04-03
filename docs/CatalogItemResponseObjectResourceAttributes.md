# CatalogItemResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **NullableString** | The ID of the catalog item in an external system. | [optional] 
**Title** | Pointer to **NullableString** | The title of the catalog item. | [optional] 
**Description** | Pointer to **NullableString** | A description of the catalog item. | [optional] 
**Price** | Pointer to **NullableFloat32** | This field can be used to set the price on the catalog item, which is what gets displayed for the item when included in emails. For most price-update use cases, you will also want to update the &#x60;price&#x60; on any child variants, using the [Update Catalog Variant Endpoint](https://developers.klaviyo.com/en/reference/update_catalog_variant). | [optional] 
**Url** | Pointer to **NullableString** | URL pointing to the location of the catalog item on your website. | [optional] 
**ImageFullUrl** | Pointer to **NullableString** | URL pointing to the location of a full image of the catalog item. | [optional] 
**ImageThumbnailUrl** | Pointer to **NullableString** | URL pointing to the location of an image thumbnail of the catalog item | [optional] 
**Images** | Pointer to **[]string** | List of URLs pointing to the locations of images of the catalog item. | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | Flat JSON blob to provide custom metadata about the catalog item. May not exceed 100kb. | [optional] 
**Published** | Pointer to **NullableBool** | Boolean value indicating whether the catalog item is published. | [optional] 
**Created** | Pointer to **NullableTime** | Date and time when the catalog item was created, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 
**Updated** | Pointer to **NullableTime** | Date and time when the catalog item was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm). | [optional] 

## Methods

### NewCatalogItemResponseObjectResourceAttributes

`func NewCatalogItemResponseObjectResourceAttributes() *CatalogItemResponseObjectResourceAttributes`

NewCatalogItemResponseObjectResourceAttributes instantiates a new CatalogItemResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemResponseObjectResourceAttributesWithDefaults

`func NewCatalogItemResponseObjectResourceAttributesWithDefaults() *CatalogItemResponseObjectResourceAttributes`

NewCatalogItemResponseObjectResourceAttributesWithDefaults instantiates a new CatalogItemResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CatalogItemResponseObjectResourceAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CatalogItemResponseObjectResourceAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CatalogItemResponseObjectResourceAttributes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetTitle

`func (o *CatalogItemResponseObjectResourceAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogItemResponseObjectResourceAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CatalogItemResponseObjectResourceAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CatalogItemResponseObjectResourceAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogItemResponseObjectResourceAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogItemResponseObjectResourceAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrice

`func (o *CatalogItemResponseObjectResourceAttributes) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogItemResponseObjectResourceAttributes) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogItemResponseObjectResourceAttributes) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetUrl

`func (o *CatalogItemResponseObjectResourceAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CatalogItemResponseObjectResourceAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CatalogItemResponseObjectResourceAttributes) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetImageFullUrl

`func (o *CatalogItemResponseObjectResourceAttributes) GetImageFullUrl() string`

GetImageFullUrl returns the ImageFullUrl field if non-nil, zero value otherwise.

### GetImageFullUrlOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetImageFullUrlOk() (*string, bool)`

GetImageFullUrlOk returns a tuple with the ImageFullUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFullUrl

`func (o *CatalogItemResponseObjectResourceAttributes) SetImageFullUrl(v string)`

SetImageFullUrl sets ImageFullUrl field to given value.

### HasImageFullUrl

`func (o *CatalogItemResponseObjectResourceAttributes) HasImageFullUrl() bool`

HasImageFullUrl returns a boolean if a field has been set.

### SetImageFullUrlNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetImageFullUrlNil(b bool)`

 SetImageFullUrlNil sets the value for ImageFullUrl to be an explicit nil

### UnsetImageFullUrl
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetImageFullUrl()`

UnsetImageFullUrl ensures that no value is present for ImageFullUrl, not even an explicit nil
### GetImageThumbnailUrl

`func (o *CatalogItemResponseObjectResourceAttributes) GetImageThumbnailUrl() string`

GetImageThumbnailUrl returns the ImageThumbnailUrl field if non-nil, zero value otherwise.

### GetImageThumbnailUrlOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetImageThumbnailUrlOk() (*string, bool)`

GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageThumbnailUrl

`func (o *CatalogItemResponseObjectResourceAttributes) SetImageThumbnailUrl(v string)`

SetImageThumbnailUrl sets ImageThumbnailUrl field to given value.

### HasImageThumbnailUrl

`func (o *CatalogItemResponseObjectResourceAttributes) HasImageThumbnailUrl() bool`

HasImageThumbnailUrl returns a boolean if a field has been set.

### SetImageThumbnailUrlNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetImageThumbnailUrlNil(b bool)`

 SetImageThumbnailUrlNil sets the value for ImageThumbnailUrl to be an explicit nil

### UnsetImageThumbnailUrl
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetImageThumbnailUrl()`

UnsetImageThumbnailUrl ensures that no value is present for ImageThumbnailUrl, not even an explicit nil
### GetImages

`func (o *CatalogItemResponseObjectResourceAttributes) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CatalogItemResponseObjectResourceAttributes) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *CatalogItemResponseObjectResourceAttributes) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetCustomMetadata

`func (o *CatalogItemResponseObjectResourceAttributes) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *CatalogItemResponseObjectResourceAttributes) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *CatalogItemResponseObjectResourceAttributes) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### SetCustomMetadataNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetCustomMetadataNil(b bool)`

 SetCustomMetadataNil sets the value for CustomMetadata to be an explicit nil

### UnsetCustomMetadata
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetCustomMetadata()`

UnsetCustomMetadata ensures that no value is present for CustomMetadata, not even an explicit nil
### GetPublished

`func (o *CatalogItemResponseObjectResourceAttributes) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CatalogItemResponseObjectResourceAttributes) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CatalogItemResponseObjectResourceAttributes) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil
### GetCreated

`func (o *CatalogItemResponseObjectResourceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CatalogItemResponseObjectResourceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CatalogItemResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *CatalogItemResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CatalogItemResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CatalogItemResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CatalogItemResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *CatalogItemResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *CatalogItemResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


