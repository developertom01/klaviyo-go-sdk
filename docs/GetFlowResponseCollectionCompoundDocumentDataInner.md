# GetFlowResponseCollectionCompoundDocumentDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowEnum**](FlowEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**FlowResponseObjectResourceAttributes**](FlowResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships**](GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetFlowResponseCollectionCompoundDocumentDataInner

`func NewGetFlowResponseCollectionCompoundDocumentDataInner(type_ FlowEnum, id string, attributes FlowResponseObjectResourceAttributes, links ObjectLinks, ) *GetFlowResponseCollectionCompoundDocumentDataInner`

NewGetFlowResponseCollectionCompoundDocumentDataInner instantiates a new GetFlowResponseCollectionCompoundDocumentDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFlowResponseCollectionCompoundDocumentDataInnerWithDefaults

`func NewGetFlowResponseCollectionCompoundDocumentDataInnerWithDefaults() *GetFlowResponseCollectionCompoundDocumentDataInner`

NewGetFlowResponseCollectionCompoundDocumentDataInnerWithDefaults instantiates a new GetFlowResponseCollectionCompoundDocumentDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) GetType() FlowEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) GetTypeOk() (*FlowEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) SetType(v FlowEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) GetAttributes() FlowResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) GetAttributesOk() (*FlowResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) SetAttributes(v FlowResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) GetRelationships() GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) GetRelationshipsOk() (*GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) SetRelationships(v GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetFlowResponseCollectionCompoundDocumentDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


