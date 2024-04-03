# SegmentPartialUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SegmentEnum**](SegmentEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**SegmentPartialUpdateQueryResourceObjectAttributes**](SegmentPartialUpdateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewSegmentPartialUpdateQueryResourceObject

`func NewSegmentPartialUpdateQueryResourceObject(type_ SegmentEnum, id string, attributes SegmentPartialUpdateQueryResourceObjectAttributes, ) *SegmentPartialUpdateQueryResourceObject`

NewSegmentPartialUpdateQueryResourceObject instantiates a new SegmentPartialUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentPartialUpdateQueryResourceObjectWithDefaults

`func NewSegmentPartialUpdateQueryResourceObjectWithDefaults() *SegmentPartialUpdateQueryResourceObject`

NewSegmentPartialUpdateQueryResourceObjectWithDefaults instantiates a new SegmentPartialUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SegmentPartialUpdateQueryResourceObject) GetType() SegmentEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SegmentPartialUpdateQueryResourceObject) GetTypeOk() (*SegmentEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SegmentPartialUpdateQueryResourceObject) SetType(v SegmentEnum)`

SetType sets Type field to given value.


### GetId

`func (o *SegmentPartialUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentPartialUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentPartialUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *SegmentPartialUpdateQueryResourceObject) GetAttributes() SegmentPartialUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SegmentPartialUpdateQueryResourceObject) GetAttributesOk() (*SegmentPartialUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SegmentPartialUpdateQueryResourceObject) SetAttributes(v SegmentPartialUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


