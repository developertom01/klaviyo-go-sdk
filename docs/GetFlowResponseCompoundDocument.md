# GetFlowResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetFlowResponseCollectionCompoundDocumentDataInner**](GetFlowResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]GetFlowResponseCollectionCompoundDocumentIncludedInner**](GetFlowResponseCollectionCompoundDocumentIncludedInner.md) |  | [optional] 

## Methods

### NewGetFlowResponseCompoundDocument

`func NewGetFlowResponseCompoundDocument(data GetFlowResponseCollectionCompoundDocumentDataInner, ) *GetFlowResponseCompoundDocument`

NewGetFlowResponseCompoundDocument instantiates a new GetFlowResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFlowResponseCompoundDocumentWithDefaults

`func NewGetFlowResponseCompoundDocumentWithDefaults() *GetFlowResponseCompoundDocument`

NewGetFlowResponseCompoundDocumentWithDefaults instantiates a new GetFlowResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetFlowResponseCompoundDocument) GetData() GetFlowResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFlowResponseCompoundDocument) GetDataOk() (*GetFlowResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFlowResponseCompoundDocument) SetData(v GetFlowResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetFlowResponseCompoundDocument) GetIncluded() []GetFlowResponseCollectionCompoundDocumentIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetFlowResponseCompoundDocument) GetIncludedOk() (*[]GetFlowResponseCollectionCompoundDocumentIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetFlowResponseCompoundDocument) SetIncluded(v []GetFlowResponseCollectionCompoundDocumentIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetFlowResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


