# ImageCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | A name for the image.  Defaults to the filename if not provided.  If the name matches an existing image, a suffix will be added. | [optional] 
**ImportFromUrl** | **string** | An existing image url to import the image from. Alternatively, you may specify a base-64 encoded data-uri (&#x60;data:image/...&#x60;). Supported image formats: jpeg,png,gif. Maximum image size: 5MB. | 
**Hidden** | Pointer to **NullableBool** | If true, this image is not shown in the asset library. | [optional] [default to false]

## Methods

### NewImageCreateQueryResourceObjectAttributes

`func NewImageCreateQueryResourceObjectAttributes(importFromUrl string, ) *ImageCreateQueryResourceObjectAttributes`

NewImageCreateQueryResourceObjectAttributes instantiates a new ImageCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageCreateQueryResourceObjectAttributesWithDefaults

`func NewImageCreateQueryResourceObjectAttributesWithDefaults() *ImageCreateQueryResourceObjectAttributes`

NewImageCreateQueryResourceObjectAttributesWithDefaults instantiates a new ImageCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageCreateQueryResourceObjectAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageCreateQueryResourceObjectAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageCreateQueryResourceObjectAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageCreateQueryResourceObjectAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ImageCreateQueryResourceObjectAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImageCreateQueryResourceObjectAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImportFromUrl

`func (o *ImageCreateQueryResourceObjectAttributes) GetImportFromUrl() string`

GetImportFromUrl returns the ImportFromUrl field if non-nil, zero value otherwise.

### GetImportFromUrlOk

`func (o *ImageCreateQueryResourceObjectAttributes) GetImportFromUrlOk() (*string, bool)`

GetImportFromUrlOk returns a tuple with the ImportFromUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFromUrl

`func (o *ImageCreateQueryResourceObjectAttributes) SetImportFromUrl(v string)`

SetImportFromUrl sets ImportFromUrl field to given value.


### GetHidden

`func (o *ImageCreateQueryResourceObjectAttributes) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ImageCreateQueryResourceObjectAttributes) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ImageCreateQueryResourceObjectAttributes) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ImageCreateQueryResourceObjectAttributes) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *ImageCreateQueryResourceObjectAttributes) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *ImageCreateQueryResourceObjectAttributes) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


