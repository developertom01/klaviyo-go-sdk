# GetEventResponseCollectionCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetEventResponseCollectionCompoundDocumentDataInner**](GetEventResponseCollectionCompoundDocumentDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 
**Included** | Pointer to [**[]GetEventResponseCollectionCompoundDocumentIncludedInner**](GetEventResponseCollectionCompoundDocumentIncludedInner.md) |  | [optional] 

## Methods

### NewGetEventResponseCollectionCompoundDocument

`func NewGetEventResponseCollectionCompoundDocument(data []GetEventResponseCollectionCompoundDocumentDataInner, links CollectionLinks, ) *GetEventResponseCollectionCompoundDocument`

NewGetEventResponseCollectionCompoundDocument instantiates a new GetEventResponseCollectionCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventResponseCollectionCompoundDocumentWithDefaults

`func NewGetEventResponseCollectionCompoundDocumentWithDefaults() *GetEventResponseCollectionCompoundDocument`

NewGetEventResponseCollectionCompoundDocumentWithDefaults instantiates a new GetEventResponseCollectionCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetEventResponseCollectionCompoundDocument) GetData() []GetEventResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEventResponseCollectionCompoundDocument) GetDataOk() (*[]GetEventResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEventResponseCollectionCompoundDocument) SetData(v []GetEventResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetEventResponseCollectionCompoundDocument) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetEventResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetEventResponseCollectionCompoundDocument) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.


### GetIncluded

`func (o *GetEventResponseCollectionCompoundDocument) GetIncluded() []GetEventResponseCollectionCompoundDocumentIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetEventResponseCollectionCompoundDocument) GetIncludedOk() (*[]GetEventResponseCollectionCompoundDocumentIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetEventResponseCollectionCompoundDocument) SetIncluded(v []GetEventResponseCollectionCompoundDocumentIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetEventResponseCollectionCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


