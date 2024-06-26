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
	"bytes"
	"fmt"
)

// checks if the GetCampaignMessageResponseCollectionCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignMessageResponseCollectionCompoundDocument{}

// GetCampaignMessageResponseCollectionCompoundDocument struct for GetCampaignMessageResponseCollectionCompoundDocument
type GetCampaignMessageResponseCollectionCompoundDocument struct {
	Data []GetCampaignMessageResponseCompoundDocumentData `json:"data"`
	Links CollectionLinks `json:"links"`
	Included []GetCampaignMessageResponseCompoundDocumentIncludedInner `json:"included,omitempty"`
}

type _GetCampaignMessageResponseCollectionCompoundDocument GetCampaignMessageResponseCollectionCompoundDocument

// NewGetCampaignMessageResponseCollectionCompoundDocument instantiates a new GetCampaignMessageResponseCollectionCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignMessageResponseCollectionCompoundDocument(data []GetCampaignMessageResponseCompoundDocumentData, links CollectionLinks) *GetCampaignMessageResponseCollectionCompoundDocument {
	this := GetCampaignMessageResponseCollectionCompoundDocument{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetCampaignMessageResponseCollectionCompoundDocumentWithDefaults instantiates a new GetCampaignMessageResponseCollectionCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignMessageResponseCollectionCompoundDocumentWithDefaults() *GetCampaignMessageResponseCollectionCompoundDocument {
	this := GetCampaignMessageResponseCollectionCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetData() []GetCampaignMessageResponseCompoundDocumentData {
	if o == nil {
		var ret []GetCampaignMessageResponseCompoundDocumentData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetDataOk() ([]GetCampaignMessageResponseCompoundDocumentData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCampaignMessageResponseCollectionCompoundDocument) SetData(v []GetCampaignMessageResponseCompoundDocumentData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCampaignMessageResponseCollectionCompoundDocument) SetLinks(v CollectionLinks) {
	o.Links = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetIncluded() []GetCampaignMessageResponseCompoundDocumentIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GetCampaignMessageResponseCompoundDocumentIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetIncludedOk() ([]GetCampaignMessageResponseCompoundDocumentIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GetCampaignMessageResponseCollectionCompoundDocument) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GetCampaignMessageResponseCompoundDocumentIncludedInner and assigns it to the Included field.
func (o *GetCampaignMessageResponseCollectionCompoundDocument) SetIncluded(v []GetCampaignMessageResponseCompoundDocumentIncludedInner) {
	o.Included = v
}

func (o GetCampaignMessageResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignMessageResponseCollectionCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *GetCampaignMessageResponseCollectionCompoundDocument) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetCampaignMessageResponseCollectionCompoundDocument := _GetCampaignMessageResponseCollectionCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCampaignMessageResponseCollectionCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetCampaignMessageResponseCollectionCompoundDocument(varGetCampaignMessageResponseCollectionCompoundDocument)

	return err
}

type NullableGetCampaignMessageResponseCollectionCompoundDocument struct {
	value *GetCampaignMessageResponseCollectionCompoundDocument
	isSet bool
}

func (v NullableGetCampaignMessageResponseCollectionCompoundDocument) Get() *GetCampaignMessageResponseCollectionCompoundDocument {
	return v.value
}

func (v *NullableGetCampaignMessageResponseCollectionCompoundDocument) Set(val *GetCampaignMessageResponseCollectionCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignMessageResponseCollectionCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignMessageResponseCollectionCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignMessageResponseCollectionCompoundDocument(val *GetCampaignMessageResponseCollectionCompoundDocument) *NullableGetCampaignMessageResponseCollectionCompoundDocument {
	return &NullableGetCampaignMessageResponseCollectionCompoundDocument{value: val, isSet: true}
}

func (v NullableGetCampaignMessageResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignMessageResponseCollectionCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


