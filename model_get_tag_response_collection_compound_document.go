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

// checks if the GetTagResponseCollectionCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTagResponseCollectionCompoundDocument{}

// GetTagResponseCollectionCompoundDocument struct for GetTagResponseCollectionCompoundDocument
type GetTagResponseCollectionCompoundDocument struct {
	Data []GetTagResponseCollectionCompoundDocumentDataInner `json:"data"`
	Links CollectionLinks `json:"links"`
	Included []TagGroupResponseObjectResource `json:"included,omitempty"`
}

type _GetTagResponseCollectionCompoundDocument GetTagResponseCollectionCompoundDocument

// NewGetTagResponseCollectionCompoundDocument instantiates a new GetTagResponseCollectionCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTagResponseCollectionCompoundDocument(data []GetTagResponseCollectionCompoundDocumentDataInner, links CollectionLinks) *GetTagResponseCollectionCompoundDocument {
	this := GetTagResponseCollectionCompoundDocument{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetTagResponseCollectionCompoundDocumentWithDefaults instantiates a new GetTagResponseCollectionCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTagResponseCollectionCompoundDocumentWithDefaults() *GetTagResponseCollectionCompoundDocument {
	this := GetTagResponseCollectionCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetTagResponseCollectionCompoundDocument) GetData() []GetTagResponseCollectionCompoundDocumentDataInner {
	if o == nil {
		var ret []GetTagResponseCollectionCompoundDocumentDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetTagResponseCollectionCompoundDocument) GetDataOk() ([]GetTagResponseCollectionCompoundDocumentDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetTagResponseCollectionCompoundDocument) SetData(v []GetTagResponseCollectionCompoundDocumentDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetTagResponseCollectionCompoundDocument) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetTagResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetTagResponseCollectionCompoundDocument) SetLinks(v CollectionLinks) {
	o.Links = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GetTagResponseCollectionCompoundDocument) GetIncluded() []TagGroupResponseObjectResource {
	if o == nil || IsNil(o.Included) {
		var ret []TagGroupResponseObjectResource
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTagResponseCollectionCompoundDocument) GetIncludedOk() ([]TagGroupResponseObjectResource, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GetTagResponseCollectionCompoundDocument) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []TagGroupResponseObjectResource and assigns it to the Included field.
func (o *GetTagResponseCollectionCompoundDocument) SetIncluded(v []TagGroupResponseObjectResource) {
	o.Included = v
}

func (o GetTagResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTagResponseCollectionCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *GetTagResponseCollectionCompoundDocument) UnmarshalJSON(data []byte) (err error) {
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

	varGetTagResponseCollectionCompoundDocument := _GetTagResponseCollectionCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTagResponseCollectionCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetTagResponseCollectionCompoundDocument(varGetTagResponseCollectionCompoundDocument)

	return err
}

type NullableGetTagResponseCollectionCompoundDocument struct {
	value *GetTagResponseCollectionCompoundDocument
	isSet bool
}

func (v NullableGetTagResponseCollectionCompoundDocument) Get() *GetTagResponseCollectionCompoundDocument {
	return v.value
}

func (v *NullableGetTagResponseCollectionCompoundDocument) Set(val *GetTagResponseCollectionCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTagResponseCollectionCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTagResponseCollectionCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTagResponseCollectionCompoundDocument(val *GetTagResponseCollectionCompoundDocument) *NullableGetTagResponseCollectionCompoundDocument {
	return &NullableGetTagResponseCollectionCompoundDocument{value: val, isSet: true}
}

func (v NullableGetTagResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTagResponseCollectionCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


