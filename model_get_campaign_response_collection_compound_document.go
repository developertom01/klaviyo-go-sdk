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

// checks if the GetCampaignResponseCollectionCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignResponseCollectionCompoundDocument{}

// GetCampaignResponseCollectionCompoundDocument struct for GetCampaignResponseCollectionCompoundDocument
type GetCampaignResponseCollectionCompoundDocument struct {
	Data []GetCampaignResponseCollectionCompoundDocumentDataInner `json:"data"`
	Links CollectionLinks `json:"links"`
	Included []GetCampaignResponseCollectionCompoundDocumentIncludedInner `json:"included,omitempty"`
}

type _GetCampaignResponseCollectionCompoundDocument GetCampaignResponseCollectionCompoundDocument

// NewGetCampaignResponseCollectionCompoundDocument instantiates a new GetCampaignResponseCollectionCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignResponseCollectionCompoundDocument(data []GetCampaignResponseCollectionCompoundDocumentDataInner, links CollectionLinks) *GetCampaignResponseCollectionCompoundDocument {
	this := GetCampaignResponseCollectionCompoundDocument{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetCampaignResponseCollectionCompoundDocumentWithDefaults instantiates a new GetCampaignResponseCollectionCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignResponseCollectionCompoundDocumentWithDefaults() *GetCampaignResponseCollectionCompoundDocument {
	this := GetCampaignResponseCollectionCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetCampaignResponseCollectionCompoundDocument) GetData() []GetCampaignResponseCollectionCompoundDocumentDataInner {
	if o == nil {
		var ret []GetCampaignResponseCollectionCompoundDocumentDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCampaignResponseCollectionCompoundDocument) GetDataOk() ([]GetCampaignResponseCollectionCompoundDocumentDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCampaignResponseCollectionCompoundDocument) SetData(v []GetCampaignResponseCollectionCompoundDocumentDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetCampaignResponseCollectionCompoundDocument) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCampaignResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCampaignResponseCollectionCompoundDocument) SetLinks(v CollectionLinks) {
	o.Links = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GetCampaignResponseCollectionCompoundDocument) GetIncluded() []GetCampaignResponseCollectionCompoundDocumentIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GetCampaignResponseCollectionCompoundDocumentIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaignResponseCollectionCompoundDocument) GetIncludedOk() ([]GetCampaignResponseCollectionCompoundDocumentIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GetCampaignResponseCollectionCompoundDocument) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GetCampaignResponseCollectionCompoundDocumentIncludedInner and assigns it to the Included field.
func (o *GetCampaignResponseCollectionCompoundDocument) SetIncluded(v []GetCampaignResponseCollectionCompoundDocumentIncludedInner) {
	o.Included = v
}

func (o GetCampaignResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignResponseCollectionCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *GetCampaignResponseCollectionCompoundDocument) UnmarshalJSON(data []byte) (err error) {
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

	varGetCampaignResponseCollectionCompoundDocument := _GetCampaignResponseCollectionCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCampaignResponseCollectionCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetCampaignResponseCollectionCompoundDocument(varGetCampaignResponseCollectionCompoundDocument)

	return err
}

type NullableGetCampaignResponseCollectionCompoundDocument struct {
	value *GetCampaignResponseCollectionCompoundDocument
	isSet bool
}

func (v NullableGetCampaignResponseCollectionCompoundDocument) Get() *GetCampaignResponseCollectionCompoundDocument {
	return v.value
}

func (v *NullableGetCampaignResponseCollectionCompoundDocument) Set(val *GetCampaignResponseCollectionCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignResponseCollectionCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignResponseCollectionCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignResponseCollectionCompoundDocument(val *GetCampaignResponseCollectionCompoundDocument) *NullableGetCampaignResponseCollectionCompoundDocument {
	return &NullableGetCampaignResponseCollectionCompoundDocument{value: val, isSet: true}
}

func (v NullableGetCampaignResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignResponseCollectionCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


