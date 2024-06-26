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

// checks if the GetSegmentListResponseCollectionCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSegmentListResponseCollectionCompoundDocument{}

// GetSegmentListResponseCollectionCompoundDocument struct for GetSegmentListResponseCollectionCompoundDocument
type GetSegmentListResponseCollectionCompoundDocument struct {
	Data []GetSegmentListResponseCollectionCompoundDocumentDataInner `json:"data"`
	Links CollectionLinks `json:"links"`
	Included []TagResponseObjectResource `json:"included,omitempty"`
}

type _GetSegmentListResponseCollectionCompoundDocument GetSegmentListResponseCollectionCompoundDocument

// NewGetSegmentListResponseCollectionCompoundDocument instantiates a new GetSegmentListResponseCollectionCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSegmentListResponseCollectionCompoundDocument(data []GetSegmentListResponseCollectionCompoundDocumentDataInner, links CollectionLinks) *GetSegmentListResponseCollectionCompoundDocument {
	this := GetSegmentListResponseCollectionCompoundDocument{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetSegmentListResponseCollectionCompoundDocumentWithDefaults instantiates a new GetSegmentListResponseCollectionCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSegmentListResponseCollectionCompoundDocumentWithDefaults() *GetSegmentListResponseCollectionCompoundDocument {
	this := GetSegmentListResponseCollectionCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetSegmentListResponseCollectionCompoundDocument) GetData() []GetSegmentListResponseCollectionCompoundDocumentDataInner {
	if o == nil {
		var ret []GetSegmentListResponseCollectionCompoundDocumentDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetSegmentListResponseCollectionCompoundDocument) GetDataOk() ([]GetSegmentListResponseCollectionCompoundDocumentDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetSegmentListResponseCollectionCompoundDocument) SetData(v []GetSegmentListResponseCollectionCompoundDocumentDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetSegmentListResponseCollectionCompoundDocument) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetSegmentListResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetSegmentListResponseCollectionCompoundDocument) SetLinks(v CollectionLinks) {
	o.Links = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GetSegmentListResponseCollectionCompoundDocument) GetIncluded() []TagResponseObjectResource {
	if o == nil || IsNil(o.Included) {
		var ret []TagResponseObjectResource
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSegmentListResponseCollectionCompoundDocument) GetIncludedOk() ([]TagResponseObjectResource, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GetSegmentListResponseCollectionCompoundDocument) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []TagResponseObjectResource and assigns it to the Included field.
func (o *GetSegmentListResponseCollectionCompoundDocument) SetIncluded(v []TagResponseObjectResource) {
	o.Included = v
}

func (o GetSegmentListResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSegmentListResponseCollectionCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *GetSegmentListResponseCollectionCompoundDocument) UnmarshalJSON(data []byte) (err error) {
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

	varGetSegmentListResponseCollectionCompoundDocument := _GetSegmentListResponseCollectionCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSegmentListResponseCollectionCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetSegmentListResponseCollectionCompoundDocument(varGetSegmentListResponseCollectionCompoundDocument)

	return err
}

type NullableGetSegmentListResponseCollectionCompoundDocument struct {
	value *GetSegmentListResponseCollectionCompoundDocument
	isSet bool
}

func (v NullableGetSegmentListResponseCollectionCompoundDocument) Get() *GetSegmentListResponseCollectionCompoundDocument {
	return v.value
}

func (v *NullableGetSegmentListResponseCollectionCompoundDocument) Set(val *GetSegmentListResponseCollectionCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSegmentListResponseCollectionCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSegmentListResponseCollectionCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSegmentListResponseCollectionCompoundDocument(val *GetSegmentListResponseCollectionCompoundDocument) *NullableGetSegmentListResponseCollectionCompoundDocument {
	return &NullableGetSegmentListResponseCollectionCompoundDocument{value: val, isSet: true}
}

func (v NullableGetSegmentListResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSegmentListResponseCollectionCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


