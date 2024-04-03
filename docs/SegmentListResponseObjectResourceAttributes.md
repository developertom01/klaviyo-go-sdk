# SegmentListResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | A helpful name to label the segment | [optional] 
**Created** | Pointer to **NullableTime** | Date and time when the segment was created, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 
**Updated** | Pointer to **NullableTime** | Date and time when the segment was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 
**IsActive** | **bool** |  | 
**IsProcessing** | **bool** |  | 
**IsStarred** | **bool** |  | 

## Methods

### NewSegmentListResponseObjectResourceAttributes

`func NewSegmentListResponseObjectResourceAttributes(isActive bool, isProcessing bool, isStarred bool, ) *SegmentListResponseObjectResourceAttributes`

NewSegmentListResponseObjectResourceAttributes instantiates a new SegmentListResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentListResponseObjectResourceAttributesWithDefaults

`func NewSegmentListResponseObjectResourceAttributesWithDefaults() *SegmentListResponseObjectResourceAttributes`

NewSegmentListResponseObjectResourceAttributesWithDefaults instantiates a new SegmentListResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SegmentListResponseObjectResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentListResponseObjectResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentListResponseObjectResourceAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SegmentListResponseObjectResourceAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SegmentListResponseObjectResourceAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SegmentListResponseObjectResourceAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreated

`func (o *SegmentListResponseObjectResourceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SegmentListResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SegmentListResponseObjectResourceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SegmentListResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *SegmentListResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SegmentListResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *SegmentListResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SegmentListResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SegmentListResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SegmentListResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *SegmentListResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *SegmentListResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetIsActive

`func (o *SegmentListResponseObjectResourceAttributes) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SegmentListResponseObjectResourceAttributes) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SegmentListResponseObjectResourceAttributes) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsProcessing

`func (o *SegmentListResponseObjectResourceAttributes) GetIsProcessing() bool`

GetIsProcessing returns the IsProcessing field if non-nil, zero value otherwise.

### GetIsProcessingOk

`func (o *SegmentListResponseObjectResourceAttributes) GetIsProcessingOk() (*bool, bool)`

GetIsProcessingOk returns a tuple with the IsProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProcessing

`func (o *SegmentListResponseObjectResourceAttributes) SetIsProcessing(v bool)`

SetIsProcessing sets IsProcessing field to given value.


### GetIsStarred

`func (o *SegmentListResponseObjectResourceAttributes) GetIsStarred() bool`

GetIsStarred returns the IsStarred field if non-nil, zero value otherwise.

### GetIsStarredOk

`func (o *SegmentListResponseObjectResourceAttributes) GetIsStarredOk() (*bool, bool)`

GetIsStarredOk returns a tuple with the IsStarred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStarred

`func (o *SegmentListResponseObjectResourceAttributes) SetIsStarred(v bool)`

SetIsStarred sets IsStarred field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


