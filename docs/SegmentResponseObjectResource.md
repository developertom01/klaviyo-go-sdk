# SegmentResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SegmentEnum**](SegmentEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**SegmentListResponseObjectResourceAttributes**](SegmentListResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewSegmentResponseObjectResource

`func NewSegmentResponseObjectResource(type_ SegmentEnum, id string, attributes SegmentListResponseObjectResourceAttributes, links ObjectLinks, ) *SegmentResponseObjectResource`

NewSegmentResponseObjectResource instantiates a new SegmentResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentResponseObjectResourceWithDefaults

`func NewSegmentResponseObjectResourceWithDefaults() *SegmentResponseObjectResource`

NewSegmentResponseObjectResourceWithDefaults instantiates a new SegmentResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SegmentResponseObjectResource) GetType() SegmentEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SegmentResponseObjectResource) GetTypeOk() (*SegmentEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SegmentResponseObjectResource) SetType(v SegmentEnum)`

SetType sets Type field to given value.


### GetId

`func (o *SegmentResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *SegmentResponseObjectResource) GetAttributes() SegmentListResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SegmentResponseObjectResource) GetAttributesOk() (*SegmentListResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SegmentResponseObjectResource) SetAttributes(v SegmentListResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *SegmentResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SegmentResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SegmentResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


