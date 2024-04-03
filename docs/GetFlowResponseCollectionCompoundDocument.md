# GetFlowResponseCollectionCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetFlowResponseCollectionCompoundDocumentDataInner**](GetFlowResponseCollectionCompoundDocumentDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 
**Included** | Pointer to [**[]GetFlowResponseCollectionCompoundDocumentIncludedInner**](GetFlowResponseCollectionCompoundDocumentIncludedInner.md) |  | [optional] 

## Methods

### NewGetFlowResponseCollectionCompoundDocument

`func NewGetFlowResponseCollectionCompoundDocument(data []GetFlowResponseCollectionCompoundDocumentDataInner, links CollectionLinks, ) *GetFlowResponseCollectionCompoundDocument`

NewGetFlowResponseCollectionCompoundDocument instantiates a new GetFlowResponseCollectionCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFlowResponseCollectionCompoundDocumentWithDefaults

`func NewGetFlowResponseCollectionCompoundDocumentWithDefaults() *GetFlowResponseCollectionCompoundDocument`

NewGetFlowResponseCollectionCompoundDocumentWithDefaults instantiates a new GetFlowResponseCollectionCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetFlowResponseCollectionCompoundDocument) GetData() []GetFlowResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFlowResponseCollectionCompoundDocument) GetDataOk() (*[]GetFlowResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFlowResponseCollectionCompoundDocument) SetData(v []GetFlowResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetFlowResponseCollectionCompoundDocument) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetFlowResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetFlowResponseCollectionCompoundDocument) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.


### GetIncluded

`func (o *GetFlowResponseCollectionCompoundDocument) GetIncluded() []GetFlowResponseCollectionCompoundDocumentIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetFlowResponseCollectionCompoundDocument) GetIncludedOk() (*[]GetFlowResponseCollectionCompoundDocumentIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetFlowResponseCollectionCompoundDocument) SetIncluded(v []GetFlowResponseCollectionCompoundDocumentIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetFlowResponseCollectionCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


