# GetFlowActionResponseCompoundDocumentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowActionEnum**](FlowActionEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**FlowActionResponseObjectResourceAttributes**](FlowActionResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetFlowActionResponseCompoundDocumentDataAllOfRelationships**](GetFlowActionResponseCompoundDocumentDataAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetFlowActionResponseCompoundDocumentData

`func NewGetFlowActionResponseCompoundDocumentData(type_ FlowActionEnum, id string, attributes FlowActionResponseObjectResourceAttributes, links ObjectLinks, ) *GetFlowActionResponseCompoundDocumentData`

NewGetFlowActionResponseCompoundDocumentData instantiates a new GetFlowActionResponseCompoundDocumentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFlowActionResponseCompoundDocumentDataWithDefaults

`func NewGetFlowActionResponseCompoundDocumentDataWithDefaults() *GetFlowActionResponseCompoundDocumentData`

NewGetFlowActionResponseCompoundDocumentDataWithDefaults instantiates a new GetFlowActionResponseCompoundDocumentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetFlowActionResponseCompoundDocumentData) GetType() FlowActionEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetFlowActionResponseCompoundDocumentData) GetTypeOk() (*FlowActionEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetFlowActionResponseCompoundDocumentData) SetType(v FlowActionEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetFlowActionResponseCompoundDocumentData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetFlowActionResponseCompoundDocumentData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetFlowActionResponseCompoundDocumentData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetFlowActionResponseCompoundDocumentData) GetAttributes() FlowActionResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetFlowActionResponseCompoundDocumentData) GetAttributesOk() (*FlowActionResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetFlowActionResponseCompoundDocumentData) SetAttributes(v FlowActionResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetFlowActionResponseCompoundDocumentData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetFlowActionResponseCompoundDocumentData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetFlowActionResponseCompoundDocumentData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetFlowActionResponseCompoundDocumentData) GetRelationships() GetFlowActionResponseCompoundDocumentDataAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetFlowActionResponseCompoundDocumentData) GetRelationshipsOk() (*GetFlowActionResponseCompoundDocumentDataAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetFlowActionResponseCompoundDocumentData) SetRelationships(v GetFlowActionResponseCompoundDocumentDataAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetFlowActionResponseCompoundDocumentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


