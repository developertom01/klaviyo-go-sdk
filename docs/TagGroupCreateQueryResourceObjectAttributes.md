# TagGroupCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The Tag Group name | 
**Exclusive** | Pointer to **NullableBool** |  | [optional] [default to false]

## Methods

### NewTagGroupCreateQueryResourceObjectAttributes

`func NewTagGroupCreateQueryResourceObjectAttributes(name string, ) *TagGroupCreateQueryResourceObjectAttributes`

NewTagGroupCreateQueryResourceObjectAttributes instantiates a new TagGroupCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagGroupCreateQueryResourceObjectAttributesWithDefaults

`func NewTagGroupCreateQueryResourceObjectAttributesWithDefaults() *TagGroupCreateQueryResourceObjectAttributes`

NewTagGroupCreateQueryResourceObjectAttributesWithDefaults instantiates a new TagGroupCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagGroupCreateQueryResourceObjectAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagGroupCreateQueryResourceObjectAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagGroupCreateQueryResourceObjectAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetExclusive

`func (o *TagGroupCreateQueryResourceObjectAttributes) GetExclusive() bool`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *TagGroupCreateQueryResourceObjectAttributes) GetExclusiveOk() (*bool, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *TagGroupCreateQueryResourceObjectAttributes) SetExclusive(v bool)`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *TagGroupCreateQueryResourceObjectAttributes) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### SetExclusiveNil

`func (o *TagGroupCreateQueryResourceObjectAttributes) SetExclusiveNil(b bool)`

 SetExclusiveNil sets the value for Exclusive to be an explicit nil

### UnsetExclusive
`func (o *TagGroupCreateQueryResourceObjectAttributes) UnsetExclusive()`

UnsetExclusive ensures that no value is present for Exclusive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


