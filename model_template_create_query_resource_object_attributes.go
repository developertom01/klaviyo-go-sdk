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

// checks if the TemplateCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateCreateQueryResourceObjectAttributes{}

// TemplateCreateQueryResourceObjectAttributes struct for TemplateCreateQueryResourceObjectAttributes
type TemplateCreateQueryResourceObjectAttributes struct {
	// The name of the template
	Name string `json:"name"`
	// Restricted to CODE
	EditorType string `json:"editor_type"`
	// The HTML contents of the template
	Html NullableString `json:"html,omitempty"`
	// The plaintext version of the template
	Text NullableString `json:"text,omitempty"`
}

type _TemplateCreateQueryResourceObjectAttributes TemplateCreateQueryResourceObjectAttributes

// NewTemplateCreateQueryResourceObjectAttributes instantiates a new TemplateCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateCreateQueryResourceObjectAttributes(name string, editorType string) *TemplateCreateQueryResourceObjectAttributes {
	this := TemplateCreateQueryResourceObjectAttributes{}
	this.Name = name
	this.EditorType = editorType
	return &this
}

// NewTemplateCreateQueryResourceObjectAttributesWithDefaults instantiates a new TemplateCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateCreateQueryResourceObjectAttributesWithDefaults() *TemplateCreateQueryResourceObjectAttributes {
	this := TemplateCreateQueryResourceObjectAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *TemplateCreateQueryResourceObjectAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateCreateQueryResourceObjectAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateCreateQueryResourceObjectAttributes) SetName(v string) {
	o.Name = v
}

// GetEditorType returns the EditorType field value
func (o *TemplateCreateQueryResourceObjectAttributes) GetEditorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EditorType
}

// GetEditorTypeOk returns a tuple with the EditorType field value
// and a boolean to check if the value has been set.
func (o *TemplateCreateQueryResourceObjectAttributes) GetEditorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditorType, true
}

// SetEditorType sets field value
func (o *TemplateCreateQueryResourceObjectAttributes) SetEditorType(v string) {
	o.EditorType = v
}

// GetHtml returns the Html field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateCreateQueryResourceObjectAttributes) GetHtml() string {
	if o == nil || IsNil(o.Html.Get()) {
		var ret string
		return ret
	}
	return *o.Html.Get()
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateCreateQueryResourceObjectAttributes) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Html.Get(), o.Html.IsSet()
}

// HasHtml returns a boolean if a field has been set.
func (o *TemplateCreateQueryResourceObjectAttributes) HasHtml() bool {
	if o != nil && o.Html.IsSet() {
		return true
	}

	return false
}

// SetHtml gets a reference to the given NullableString and assigns it to the Html field.
func (o *TemplateCreateQueryResourceObjectAttributes) SetHtml(v string) {
	o.Html.Set(&v)
}
// SetHtmlNil sets the value for Html to be an explicit nil
func (o *TemplateCreateQueryResourceObjectAttributes) SetHtmlNil() {
	o.Html.Set(nil)
}

// UnsetHtml ensures that no value is present for Html, not even an explicit nil
func (o *TemplateCreateQueryResourceObjectAttributes) UnsetHtml() {
	o.Html.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateCreateQueryResourceObjectAttributes) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateCreateQueryResourceObjectAttributes) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *TemplateCreateQueryResourceObjectAttributes) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *TemplateCreateQueryResourceObjectAttributes) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *TemplateCreateQueryResourceObjectAttributes) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *TemplateCreateQueryResourceObjectAttributes) UnsetText() {
	o.Text.Unset()
}

func (o TemplateCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["editor_type"] = o.EditorType
	if o.Html.IsSet() {
		toSerialize["html"] = o.Html.Get()
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	return toSerialize, nil
}

func (o *TemplateCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"editor_type",
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

	varTemplateCreateQueryResourceObjectAttributes := _TemplateCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplateCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = TemplateCreateQueryResourceObjectAttributes(varTemplateCreateQueryResourceObjectAttributes)

	return err
}

type NullableTemplateCreateQueryResourceObjectAttributes struct {
	value *TemplateCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableTemplateCreateQueryResourceObjectAttributes) Get() *TemplateCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableTemplateCreateQueryResourceObjectAttributes) Set(val *TemplateCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateCreateQueryResourceObjectAttributes(val *TemplateCreateQueryResourceObjectAttributes) *NullableTemplateCreateQueryResourceObjectAttributes {
	return &NullableTemplateCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableTemplateCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


