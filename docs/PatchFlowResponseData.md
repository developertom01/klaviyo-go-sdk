# PatchFlowResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowEnum**](FlowEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**FlowResponseObjectResourceAttributes**](FlowResponseObjectResourceAttributes.md) |  | 
**Relationships** | Pointer to [**GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships**](GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPatchFlowResponseData

`func NewPatchFlowResponseData(type_ FlowEnum, id string, attributes FlowResponseObjectResourceAttributes, links ObjectLinks, ) *PatchFlowResponseData`

NewPatchFlowResponseData instantiates a new PatchFlowResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchFlowResponseDataWithDefaults

`func NewPatchFlowResponseDataWithDefaults() *PatchFlowResponseData`

NewPatchFlowResponseDataWithDefaults instantiates a new PatchFlowResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PatchFlowResponseData) GetType() FlowEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchFlowResponseData) GetTypeOk() (*FlowEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchFlowResponseData) SetType(v FlowEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PatchFlowResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchFlowResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchFlowResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PatchFlowResponseData) GetAttributes() FlowResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchFlowResponseData) GetAttributesOk() (*FlowResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchFlowResponseData) SetAttributes(v FlowResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PatchFlowResponseData) GetRelationships() GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchFlowResponseData) GetRelationshipsOk() (*GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchFlowResponseData) SetRelationships(v GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchFlowResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PatchFlowResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PatchFlowResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PatchFlowResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


