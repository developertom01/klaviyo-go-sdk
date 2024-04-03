# GetSegmentRetrieveResponseCompoundDocumentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SegmentEnum**](SegmentEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**GetListRetrieveResponseCompoundDocumentDataAllOfAttributes**](GetListRetrieveResponseCompoundDocumentDataAllOfAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships**](GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetSegmentRetrieveResponseCompoundDocumentData

`func NewGetSegmentRetrieveResponseCompoundDocumentData(type_ SegmentEnum, id string, attributes GetListRetrieveResponseCompoundDocumentDataAllOfAttributes, links ObjectLinks, ) *GetSegmentRetrieveResponseCompoundDocumentData`

NewGetSegmentRetrieveResponseCompoundDocumentData instantiates a new GetSegmentRetrieveResponseCompoundDocumentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSegmentRetrieveResponseCompoundDocumentDataWithDefaults

`func NewGetSegmentRetrieveResponseCompoundDocumentDataWithDefaults() *GetSegmentRetrieveResponseCompoundDocumentData`

NewGetSegmentRetrieveResponseCompoundDocumentDataWithDefaults instantiates a new GetSegmentRetrieveResponseCompoundDocumentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) GetType() SegmentEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) GetTypeOk() (*SegmentEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) SetType(v SegmentEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) GetAttributes() GetListRetrieveResponseCompoundDocumentDataAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) GetAttributesOk() (*GetListRetrieveResponseCompoundDocumentDataAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) SetAttributes(v GetListRetrieveResponseCompoundDocumentDataAllOfAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) GetRelationships() GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) GetRelationshipsOk() (*GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) SetRelationships(v GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetSegmentRetrieveResponseCompoundDocumentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


