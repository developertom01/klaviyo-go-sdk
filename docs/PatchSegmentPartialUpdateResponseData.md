# PatchSegmentPartialUpdateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SegmentEnum**](SegmentEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**SegmentListResponseObjectResourceAttributes**](SegmentListResponseObjectResourceAttributes.md) |  | 
**Relationships** | Pointer to [**PostListCreateResponseDataRelationships**](PostListCreateResponseDataRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPatchSegmentPartialUpdateResponseData

`func NewPatchSegmentPartialUpdateResponseData(type_ SegmentEnum, id string, attributes SegmentListResponseObjectResourceAttributes, links ObjectLinks, ) *PatchSegmentPartialUpdateResponseData`

NewPatchSegmentPartialUpdateResponseData instantiates a new PatchSegmentPartialUpdateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSegmentPartialUpdateResponseDataWithDefaults

`func NewPatchSegmentPartialUpdateResponseDataWithDefaults() *PatchSegmentPartialUpdateResponseData`

NewPatchSegmentPartialUpdateResponseDataWithDefaults instantiates a new PatchSegmentPartialUpdateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PatchSegmentPartialUpdateResponseData) GetType() SegmentEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchSegmentPartialUpdateResponseData) GetTypeOk() (*SegmentEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchSegmentPartialUpdateResponseData) SetType(v SegmentEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PatchSegmentPartialUpdateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchSegmentPartialUpdateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchSegmentPartialUpdateResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PatchSegmentPartialUpdateResponseData) GetAttributes() SegmentListResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchSegmentPartialUpdateResponseData) GetAttributesOk() (*SegmentListResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchSegmentPartialUpdateResponseData) SetAttributes(v SegmentListResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PatchSegmentPartialUpdateResponseData) GetRelationships() PostListCreateResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchSegmentPartialUpdateResponseData) GetRelationshipsOk() (*PostListCreateResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchSegmentPartialUpdateResponseData) SetRelationships(v PostListCreateResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchSegmentPartialUpdateResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PatchSegmentPartialUpdateResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PatchSegmentPartialUpdateResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PatchSegmentPartialUpdateResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


