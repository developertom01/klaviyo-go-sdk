# GetFlowMessageResponseCompoundDocumentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowMessageEnum**](FlowMessageEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**FlowMessageResponseObjectResourceAttributes**](FlowMessageResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetFlowMessageResponseCompoundDocumentDataAllOfRelationships**](GetFlowMessageResponseCompoundDocumentDataAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetFlowMessageResponseCompoundDocumentData

`func NewGetFlowMessageResponseCompoundDocumentData(type_ FlowMessageEnum, id string, attributes FlowMessageResponseObjectResourceAttributes, links ObjectLinks, ) *GetFlowMessageResponseCompoundDocumentData`

NewGetFlowMessageResponseCompoundDocumentData instantiates a new GetFlowMessageResponseCompoundDocumentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFlowMessageResponseCompoundDocumentDataWithDefaults

`func NewGetFlowMessageResponseCompoundDocumentDataWithDefaults() *GetFlowMessageResponseCompoundDocumentData`

NewGetFlowMessageResponseCompoundDocumentDataWithDefaults instantiates a new GetFlowMessageResponseCompoundDocumentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetFlowMessageResponseCompoundDocumentData) GetType() FlowMessageEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetFlowMessageResponseCompoundDocumentData) GetTypeOk() (*FlowMessageEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetFlowMessageResponseCompoundDocumentData) SetType(v FlowMessageEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetFlowMessageResponseCompoundDocumentData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetFlowMessageResponseCompoundDocumentData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetFlowMessageResponseCompoundDocumentData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetFlowMessageResponseCompoundDocumentData) GetAttributes() FlowMessageResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetFlowMessageResponseCompoundDocumentData) GetAttributesOk() (*FlowMessageResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetFlowMessageResponseCompoundDocumentData) SetAttributes(v FlowMessageResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetFlowMessageResponseCompoundDocumentData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetFlowMessageResponseCompoundDocumentData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetFlowMessageResponseCompoundDocumentData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetFlowMessageResponseCompoundDocumentData) GetRelationships() GetFlowMessageResponseCompoundDocumentDataAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetFlowMessageResponseCompoundDocumentData) GetRelationshipsOk() (*GetFlowMessageResponseCompoundDocumentDataAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetFlowMessageResponseCompoundDocumentData) SetRelationships(v GetFlowMessageResponseCompoundDocumentDataAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetFlowMessageResponseCompoundDocumentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

