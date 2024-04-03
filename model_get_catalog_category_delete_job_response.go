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

// checks if the GetCatalogCategoryDeleteJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogCategoryDeleteJobResponse{}

// GetCatalogCategoryDeleteJobResponse struct for GetCatalogCategoryDeleteJobResponse
type GetCatalogCategoryDeleteJobResponse struct {
	Data GetCatalogCategoryDeleteJobResponseCollectionDataInner `json:"data"`
}

type _GetCatalogCategoryDeleteJobResponse GetCatalogCategoryDeleteJobResponse

// NewGetCatalogCategoryDeleteJobResponse instantiates a new GetCatalogCategoryDeleteJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogCategoryDeleteJobResponse(data GetCatalogCategoryDeleteJobResponseCollectionDataInner) *GetCatalogCategoryDeleteJobResponse {
	this := GetCatalogCategoryDeleteJobResponse{}
	this.Data = data
	return &this
}

// NewGetCatalogCategoryDeleteJobResponseWithDefaults instantiates a new GetCatalogCategoryDeleteJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogCategoryDeleteJobResponseWithDefaults() *GetCatalogCategoryDeleteJobResponse {
	this := GetCatalogCategoryDeleteJobResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCatalogCategoryDeleteJobResponse) GetData() GetCatalogCategoryDeleteJobResponseCollectionDataInner {
	if o == nil {
		var ret GetCatalogCategoryDeleteJobResponseCollectionDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryDeleteJobResponse) GetDataOk() (*GetCatalogCategoryDeleteJobResponseCollectionDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCatalogCategoryDeleteJobResponse) SetData(v GetCatalogCategoryDeleteJobResponseCollectionDataInner) {
	o.Data = v
}

func (o GetCatalogCategoryDeleteJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogCategoryDeleteJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCatalogCategoryDeleteJobResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetCatalogCategoryDeleteJobResponse := _GetCatalogCategoryDeleteJobResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogCategoryDeleteJobResponse)

	if err != nil {
		return err
	}

	*o = GetCatalogCategoryDeleteJobResponse(varGetCatalogCategoryDeleteJobResponse)

	return err
}

type NullableGetCatalogCategoryDeleteJobResponse struct {
	value *GetCatalogCategoryDeleteJobResponse
	isSet bool
}

func (v NullableGetCatalogCategoryDeleteJobResponse) Get() *GetCatalogCategoryDeleteJobResponse {
	return v.value
}

func (v *NullableGetCatalogCategoryDeleteJobResponse) Set(val *GetCatalogCategoryDeleteJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogCategoryDeleteJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogCategoryDeleteJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogCategoryDeleteJobResponse(val *GetCatalogCategoryDeleteJobResponse) *NullableGetCatalogCategoryDeleteJobResponse {
	return &NullableGetCatalogCategoryDeleteJobResponse{value: val, isSet: true}
}

func (v NullableGetCatalogCategoryDeleteJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogCategoryDeleteJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


