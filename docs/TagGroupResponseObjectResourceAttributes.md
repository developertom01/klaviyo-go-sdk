# TagGroupResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The Tag Group name | 
**Exclusive** | **bool** | If a tag group is non-exclusive, any given related resource (campaign, flow, etc.) can be linked to multiple tags from that tag group. If a tag group is exclusive, any given related resource can only be linked to one tag from that tag group. | 
**Default** | **bool** | Every company automatically has one Default Tag Group. The Default Tag Group cannot be deleted, and no other Default Tag Groups can be created. This value is true for the Default Tag Group and false for all other Tag Groups. | 

## Methods

### NewTagGroupResponseObjectResourceAttributes

`func NewTagGroupResponseObjectResourceAttributes(name string, exclusive bool, default_ bool, ) *TagGroupResponseObjectResourceAttributes`

NewTagGroupResponseObjectResourceAttributes instantiates a new TagGroupResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagGroupResponseObjectResourceAttributesWithDefaults

`func NewTagGroupResponseObjectResourceAttributesWithDefaults() *TagGroupResponseObjectResourceAttributes`

NewTagGroupResponseObjectResourceAttributesWithDefaults instantiates a new TagGroupResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagGroupResponseObjectResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagGroupResponseObjectResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagGroupResponseObjectResourceAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetExclusive

`func (o *TagGroupResponseObjectResourceAttributes) GetExclusive() bool`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *TagGroupResponseObjectResourceAttributes) GetExclusiveOk() (*bool, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *TagGroupResponseObjectResourceAttributes) SetExclusive(v bool)`

SetExclusive sets Exclusive field to given value.


### GetDefault

`func (o *TagGroupResponseObjectResourceAttributes) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TagGroupResponseObjectResourceAttributes) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TagGroupResponseObjectResourceAttributes) SetDefault(v bool)`

SetDefault sets Default field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


