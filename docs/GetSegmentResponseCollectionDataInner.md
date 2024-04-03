# GetSegmentResponseCollectionDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SegmentEnum**](SegmentEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**SegmentListResponseObjectResourceAttributes**](SegmentListResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetListResponseCollectionDataInnerAllOfRelationships**](GetListResponseCollectionDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetSegmentResponseCollectionDataInner

`func NewGetSegmentResponseCollectionDataInner(type_ SegmentEnum, id string, attributes SegmentListResponseObjectResourceAttributes, links ObjectLinks, ) *GetSegmentResponseCollectionDataInner`

NewGetSegmentResponseCollectionDataInner instantiates a new GetSegmentResponseCollectionDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSegmentResponseCollectionDataInnerWithDefaults

`func NewGetSegmentResponseCollectionDataInnerWithDefaults() *GetSegmentResponseCollectionDataInner`

NewGetSegmentResponseCollectionDataInnerWithDefaults instantiates a new GetSegmentResponseCollectionDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetSegmentResponseCollectionDataInner) GetType() SegmentEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSegmentResponseCollectionDataInner) GetTypeOk() (*SegmentEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSegmentResponseCollectionDataInner) SetType(v SegmentEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetSegmentResponseCollectionDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSegmentResponseCollectionDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSegmentResponseCollectionDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetSegmentResponseCollectionDataInner) GetAttributes() SegmentListResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetSegmentResponseCollectionDataInner) GetAttributesOk() (*SegmentListResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetSegmentResponseCollectionDataInner) SetAttributes(v SegmentListResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetSegmentResponseCollectionDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSegmentResponseCollectionDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSegmentResponseCollectionDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetSegmentResponseCollectionDataInner) GetRelationships() GetListResponseCollectionDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetSegmentResponseCollectionDataInner) GetRelationshipsOk() (*GetListResponseCollectionDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetSegmentResponseCollectionDataInner) SetRelationships(v GetListResponseCollectionDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetSegmentResponseCollectionDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


