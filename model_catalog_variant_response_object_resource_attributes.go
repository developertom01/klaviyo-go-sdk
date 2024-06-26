/*
Klaviyo API

The Klaviyo REST API. Please visit https://developers.klaviyo.com for more details.

API version: 2024-02-15
Contact: developers@klaviyo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package klaviyo

import (
	"encoding/json"
	"time"
)

// checks if the CatalogVariantResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogVariantResponseObjectResourceAttributes{}

// CatalogVariantResponseObjectResourceAttributes struct for CatalogVariantResponseObjectResourceAttributes
type CatalogVariantResponseObjectResourceAttributes struct {
	// The ID of the catalog item variant in an external system.
	ExternalId NullableString `json:"external_id,omitempty"`
	// The title of the catalog item variant.
	Title NullableString `json:"title,omitempty"`
	// A description of the catalog item variant.
	Description NullableString `json:"description,omitempty"`
	// The SKU of the catalog item variant.
	Sku NullableString `json:"sku,omitempty"`
	// This field controls the visibility of this catalog item variant in product feeds/blocks. This field supports the following values: `1`: a product will not appear in dynamic product recommendation feeds and blocks if it is out of stock. `0` or `2`: a product can appear in dynamic product recommendation feeds and blocks regardless of inventory quantity.
	InventoryPolicy NullableInt32 `json:"inventory_policy,omitempty"`
	// The quantity of the catalog item variant currently in stock.
	InventoryQuantity NullableFloat32 `json:"inventory_quantity,omitempty"`
	// This field can be used to set the price on the catalog item variant, which is what gets displayed for the item variant when included in emails. For most price-update use cases, you will also want to update the `price` on any parent items using the [Update Catalog Item Endpoint](https://developers.klaviyo.com/en/reference/update_catalog_item).
	Price NullableFloat32 `json:"price,omitempty"`
	// URL pointing to the location of the catalog item variant on your website.
	Url NullableString `json:"url,omitempty"`
	// URL pointing to the location of a full image of the catalog item variant.
	ImageFullUrl NullableString `json:"image_full_url,omitempty"`
	// URL pointing to the location of an image thumbnail of the catalog item variant.
	ImageThumbnailUrl NullableString `json:"image_thumbnail_url,omitempty"`
	// List of URLs pointing to the locations of images of the catalog item variant.
	Images []string `json:"images,omitempty"`
	// Flat JSON blob to provide custom metadata about the catalog item variant. May not exceed 100kb.
	CustomMetadata map[string]interface{} `json:"custom_metadata,omitempty"`
	// Boolean value indicating whether the catalog item variant is published.
	Published NullableBool `json:"published,omitempty"`
	// Date and time when the catalog item  variant was created, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
	Created NullableTime `json:"created,omitempty"`
	// Date and time when the catalog item variant was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
	Updated NullableTime `json:"updated,omitempty"`
}

// NewCatalogVariantResponseObjectResourceAttributes instantiates a new CatalogVariantResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogVariantResponseObjectResourceAttributes() *CatalogVariantResponseObjectResourceAttributes {
	this := CatalogVariantResponseObjectResourceAttributes{}
	return &this
}

// NewCatalogVariantResponseObjectResourceAttributesWithDefaults instantiates a new CatalogVariantResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogVariantResponseObjectResourceAttributesWithDefaults() *CatalogVariantResponseObjectResourceAttributes {
	this := CatalogVariantResponseObjectResourceAttributes{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetSku returns the Sku field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetSku() string {
	if o == nil || IsNil(o.Sku.Get()) {
		var ret string
		return ret
	}
	return *o.Sku.Get()
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sku.Get(), o.Sku.IsSet()
}

// HasSku returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasSku() bool {
	if o != nil && o.Sku.IsSet() {
		return true
	}

	return false
}

// SetSku gets a reference to the given NullableString and assigns it to the Sku field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetSku(v string) {
	o.Sku.Set(&v)
}
// SetSkuNil sets the value for Sku to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetSkuNil() {
	o.Sku.Set(nil)
}

// UnsetSku ensures that no value is present for Sku, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetSku() {
	o.Sku.Unset()
}

// GetInventoryPolicy returns the InventoryPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetInventoryPolicy() int32 {
	if o == nil || IsNil(o.InventoryPolicy.Get()) {
		var ret int32
		return ret
	}
	return *o.InventoryPolicy.Get()
}

// GetInventoryPolicyOk returns a tuple with the InventoryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetInventoryPolicyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryPolicy.Get(), o.InventoryPolicy.IsSet()
}

// HasInventoryPolicy returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasInventoryPolicy() bool {
	if o != nil && o.InventoryPolicy.IsSet() {
		return true
	}

	return false
}

// SetInventoryPolicy gets a reference to the given NullableInt32 and assigns it to the InventoryPolicy field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetInventoryPolicy(v int32) {
	o.InventoryPolicy.Set(&v)
}
// SetInventoryPolicyNil sets the value for InventoryPolicy to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetInventoryPolicyNil() {
	o.InventoryPolicy.Set(nil)
}

// UnsetInventoryPolicy ensures that no value is present for InventoryPolicy, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetInventoryPolicy() {
	o.InventoryPolicy.Unset()
}

// GetInventoryQuantity returns the InventoryQuantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetInventoryQuantity() float32 {
	if o == nil || IsNil(o.InventoryQuantity.Get()) {
		var ret float32
		return ret
	}
	return *o.InventoryQuantity.Get()
}

// GetInventoryQuantityOk returns a tuple with the InventoryQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetInventoryQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryQuantity.Get(), o.InventoryQuantity.IsSet()
}

// HasInventoryQuantity returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasInventoryQuantity() bool {
	if o != nil && o.InventoryQuantity.IsSet() {
		return true
	}

	return false
}

// SetInventoryQuantity gets a reference to the given NullableFloat32 and assigns it to the InventoryQuantity field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetInventoryQuantity(v float32) {
	o.InventoryQuantity.Set(&v)
}
// SetInventoryQuantityNil sets the value for InventoryQuantity to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetInventoryQuantityNil() {
	o.InventoryQuantity.Set(nil)
}

// UnsetInventoryQuantity ensures that no value is present for InventoryQuantity, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetInventoryQuantity() {
	o.InventoryQuantity.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetPrice() float32 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat32 and assigns it to the Price field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetPrice(v float32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetPrice() {
	o.Price.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetUrl() {
	o.Url.Unset()
}

// GetImageFullUrl returns the ImageFullUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetImageFullUrl() string {
	if o == nil || IsNil(o.ImageFullUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageFullUrl.Get()
}

// GetImageFullUrlOk returns a tuple with the ImageFullUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetImageFullUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageFullUrl.Get(), o.ImageFullUrl.IsSet()
}

// HasImageFullUrl returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasImageFullUrl() bool {
	if o != nil && o.ImageFullUrl.IsSet() {
		return true
	}

	return false
}

// SetImageFullUrl gets a reference to the given NullableString and assigns it to the ImageFullUrl field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetImageFullUrl(v string) {
	o.ImageFullUrl.Set(&v)
}
// SetImageFullUrlNil sets the value for ImageFullUrl to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetImageFullUrlNil() {
	o.ImageFullUrl.Set(nil)
}

// UnsetImageFullUrl ensures that no value is present for ImageFullUrl, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetImageFullUrl() {
	o.ImageFullUrl.Unset()
}

// GetImageThumbnailUrl returns the ImageThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetImageThumbnailUrl() string {
	if o == nil || IsNil(o.ImageThumbnailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageThumbnailUrl.Get()
}

// GetImageThumbnailUrlOk returns a tuple with the ImageThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetImageThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageThumbnailUrl.Get(), o.ImageThumbnailUrl.IsSet()
}

// HasImageThumbnailUrl returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasImageThumbnailUrl() bool {
	if o != nil && o.ImageThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetImageThumbnailUrl gets a reference to the given NullableString and assigns it to the ImageThumbnailUrl field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetImageThumbnailUrl(v string) {
	o.ImageThumbnailUrl.Set(&v)
}
// SetImageThumbnailUrlNil sets the value for ImageThumbnailUrl to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetImageThumbnailUrlNil() {
	o.ImageThumbnailUrl.Set(nil)
}

// UnsetImageThumbnailUrl ensures that no value is present for ImageThumbnailUrl, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetImageThumbnailUrl() {
	o.ImageThumbnailUrl.Unset()
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetImages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetImagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []string and assigns it to the Images field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetImages(v []string) {
	o.Images = v
}

// GetCustomMetadata returns the CustomMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetCustomMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomMetadata
}

// GetCustomMetadataOk returns a tuple with the CustomMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetCustomMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomMetadata) {
		return map[string]interface{}{}, false
	}
	return o.CustomMetadata, true
}

// HasCustomMetadata returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasCustomMetadata() bool {
	if o != nil && !IsNil(o.CustomMetadata) {
		return true
	}

	return false
}

// SetCustomMetadata gets a reference to the given map[string]interface{} and assigns it to the CustomMetadata field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetCustomMetadata(v map[string]interface{}) {
	o.CustomMetadata = v
}

// GetPublished returns the Published field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetPublished() bool {
	if o == nil || IsNil(o.Published.Get()) {
		var ret bool
		return ret
	}
	return *o.Published.Get()
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Published.Get(), o.Published.IsSet()
}

// HasPublished returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasPublished() bool {
	if o != nil && o.Published.IsSet() {
		return true
	}

	return false
}

// SetPublished gets a reference to the given NullableBool and assigns it to the Published field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetPublished(v bool) {
	o.Published.Set(&v)
}
// SetPublishedNil sets the value for Published to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetPublishedNil() {
	o.Published.Set(nil)
}

// UnsetPublished ensures that no value is present for Published, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetPublished() {
	o.Published.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetCreated() {
	o.Created.Unset()
}

// GetUpdated returns the Updated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogVariantResponseObjectResourceAttributes) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Updated.Get()
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogVariantResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updated.Get(), o.Updated.IsSet()
}

// HasUpdated returns a boolean if a field has been set.
func (o *CatalogVariantResponseObjectResourceAttributes) HasUpdated() bool {
	if o != nil && o.Updated.IsSet() {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given NullableTime and assigns it to the Updated field.
func (o *CatalogVariantResponseObjectResourceAttributes) SetUpdated(v time.Time) {
	o.Updated.Set(&v)
}
// SetUpdatedNil sets the value for Updated to be an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) SetUpdatedNil() {
	o.Updated.Set(nil)
}

// UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
func (o *CatalogVariantResponseObjectResourceAttributes) UnsetUpdated() {
	o.Updated.Unset()
}

func (o CatalogVariantResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogVariantResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId.IsSet() {
		toSerialize["external_id"] = o.ExternalId.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Sku.IsSet() {
		toSerialize["sku"] = o.Sku.Get()
	}
	if o.InventoryPolicy.IsSet() {
		toSerialize["inventory_policy"] = o.InventoryPolicy.Get()
	}
	if o.InventoryQuantity.IsSet() {
		toSerialize["inventory_quantity"] = o.InventoryQuantity.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.ImageFullUrl.IsSet() {
		toSerialize["image_full_url"] = o.ImageFullUrl.Get()
	}
	if o.ImageThumbnailUrl.IsSet() {
		toSerialize["image_thumbnail_url"] = o.ImageThumbnailUrl.Get()
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if o.CustomMetadata != nil {
		toSerialize["custom_metadata"] = o.CustomMetadata
	}
	if o.Published.IsSet() {
		toSerialize["published"] = o.Published.Get()
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Updated.IsSet() {
		toSerialize["updated"] = o.Updated.Get()
	}
	return toSerialize, nil
}

type NullableCatalogVariantResponseObjectResourceAttributes struct {
	value *CatalogVariantResponseObjectResourceAttributes
	isSet bool
}

func (v NullableCatalogVariantResponseObjectResourceAttributes) Get() *CatalogVariantResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableCatalogVariantResponseObjectResourceAttributes) Set(val *CatalogVariantResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogVariantResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogVariantResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogVariantResponseObjectResourceAttributes(val *CatalogVariantResponseObjectResourceAttributes) *NullableCatalogVariantResponseObjectResourceAttributes {
	return &NullableCatalogVariantResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableCatalogVariantResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogVariantResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


