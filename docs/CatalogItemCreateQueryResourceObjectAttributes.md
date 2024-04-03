# CatalogItemCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The ID of the catalog item in an external system. | 
**IntegrationType** | Pointer to **NullableString** | The integration type. Currently only \&quot;$custom\&quot; is supported. | [optional] [default to "$custom"]
**Title** | **string** | The title of the catalog item. | 
**Price** | Pointer to **NullableFloat32** | This field can be used to set the price on the catalog item, which is what gets displayed for the item when included in emails. For most price-update use cases, you will also want to update the &#x60;price&#x60; on any child variants, using the [Update Catalog Variant Endpoint](https://developers.klaviyo.com/en/reference/update_catalog_variant). | [optional] 
**CatalogType** | Pointer to **NullableString** | The type of catalog. Currently only \&quot;$default\&quot; is supported. | [optional] [default to "$default"]
**Description** | **string** | A description of the catalog item. | 
**Url** | **string** | URL pointing to the location of the catalog item on your website. | 
**ImageFullUrl** | Pointer to **NullableString** | URL pointing to the location of a full image of the catalog item. | [optional] 
**ImageThumbnailUrl** | Pointer to **NullableString** | URL pointing to the location of an image thumbnail of the catalog item | [optional] 
**Images** | Pointer to **[]string** | List of URLs pointing to the locations of images of the catalog item. | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | Flat JSON blob to provide custom metadata about the catalog item. May not exceed 100kb. | [optional] 
**Published** | Pointer to **NullableBool** | Boolean value indicating whether the catalog item is published. | [optional] [default to true]

## Methods

### NewCatalogItemCreateQueryResourceObjectAttributes

`func NewCatalogItemCreateQueryResourceObjectAttributes(externalId string, title string, description string, url string, ) *CatalogItemCreateQueryResourceObjectAttributes`

NewCatalogItemCreateQueryResourceObjectAttributes instantiates a new CatalogItemCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemCreateQueryResourceObjectAttributesWithDefaults

`func NewCatalogItemCreateQueryResourceObjectAttributesWithDefaults() *CatalogItemCreateQueryResourceObjectAttributes`

NewCatalogItemCreateQueryResourceObjectAttributesWithDefaults instantiates a new CatalogItemCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetIntegrationType

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *CatalogItemCreateQueryResourceObjectAttributes) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### SetIntegrationTypeNil

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetIntegrationTypeNil(b bool)`

 SetIntegrationTypeNil sets the value for IntegrationType to be an explicit nil

### UnsetIntegrationType
`func (o *CatalogItemCreateQueryResourceObjectAttributes) UnsetIntegrationType()`

UnsetIntegrationType ensures that no value is present for IntegrationType, not even an explicit nil
### GetTitle

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPrice

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogItemCreateQueryResourceObjectAttributes) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CatalogItemCreateQueryResourceObjectAttributes) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCatalogType

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetCatalogType() string`

GetCatalogType returns the CatalogType field if non-nil, zero value otherwise.

### GetCatalogTypeOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetCatalogTypeOk() (*string, bool)`

GetCatalogTypeOk returns a tuple with the CatalogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogType

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetCatalogType(v string)`

SetCatalogType sets CatalogType field to given value.

### HasCatalogType

`func (o *CatalogItemCreateQueryResourceObjectAttributes) HasCatalogType() bool`

HasCatalogType returns a boolean if a field has been set.

### SetCatalogTypeNil

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetCatalogTypeNil(b bool)`

 SetCatalogTypeNil sets the value for CatalogType to be an explicit nil

### UnsetCatalogType
`func (o *CatalogItemCreateQueryResourceObjectAttributes) UnsetCatalogType()`

UnsetCatalogType ensures that no value is present for CatalogType, not even an explicit nil
### GetDescription

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetImageFullUrl

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetImageFullUrl() string`

GetImageFullUrl returns the ImageFullUrl field if non-nil, zero value otherwise.

### GetImageFullUrlOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetImageFullUrlOk() (*string, bool)`

GetImageFullUrlOk returns a tuple with the ImageFullUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFullUrl

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetImageFullUrl(v string)`

SetImageFullUrl sets ImageFullUrl field to given value.

### HasImageFullUrl

`func (o *CatalogItemCreateQueryResourceObjectAttributes) HasImageFullUrl() bool`

HasImageFullUrl returns a boolean if a field has been set.

### SetImageFullUrlNil

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetImageFullUrlNil(b bool)`

 SetImageFullUrlNil sets the value for ImageFullUrl to be an explicit nil

### UnsetImageFullUrl
`func (o *CatalogItemCreateQueryResourceObjectAttributes) UnsetImageFullUrl()`

UnsetImageFullUrl ensures that no value is present for ImageFullUrl, not even an explicit nil
### GetImageThumbnailUrl

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetImageThumbnailUrl() string`

GetImageThumbnailUrl returns the ImageThumbnailUrl field if non-nil, zero value otherwise.

### GetImageThumbnailUrlOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetImageThumbnailUrlOk() (*string, bool)`

GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageThumbnailUrl

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetImageThumbnailUrl(v string)`

SetImageThumbnailUrl sets ImageThumbnailUrl field to given value.

### HasImageThumbnailUrl

`func (o *CatalogItemCreateQueryResourceObjectAttributes) HasImageThumbnailUrl() bool`

HasImageThumbnailUrl returns a boolean if a field has been set.

### SetImageThumbnailUrlNil

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetImageThumbnailUrlNil(b bool)`

 SetImageThumbnailUrlNil sets the value for ImageThumbnailUrl to be an explicit nil

### UnsetImageThumbnailUrl
`func (o *CatalogItemCreateQueryResourceObjectAttributes) UnsetImageThumbnailUrl()`

UnsetImageThumbnailUrl ensures that no value is present for ImageThumbnailUrl, not even an explicit nil
### GetImages

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *CatalogItemCreateQueryResourceObjectAttributes) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CatalogItemCreateQueryResourceObjectAttributes) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetCustomMetadata

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *CatalogItemCreateQueryResourceObjectAttributes) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### SetCustomMetadataNil

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetCustomMetadataNil(b bool)`

 SetCustomMetadataNil sets the value for CustomMetadata to be an explicit nil

### UnsetCustomMetadata
`func (o *CatalogItemCreateQueryResourceObjectAttributes) UnsetCustomMetadata()`

UnsetCustomMetadata ensures that no value is present for CustomMetadata, not even an explicit nil
### GetPublished

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CatalogItemCreateQueryResourceObjectAttributes) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CatalogItemCreateQueryResourceObjectAttributes) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *CatalogItemCreateQueryResourceObjectAttributes) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *CatalogItemCreateQueryResourceObjectAttributes) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


