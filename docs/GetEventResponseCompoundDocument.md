# GetEventResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetEventResponseCollectionCompoundDocumentDataInner**](GetEventResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]GetEventResponseCollectionCompoundDocumentIncludedInner**](GetEventResponseCollectionCompoundDocumentIncludedInner.md) |  | [optional] 

## Methods

### NewGetEventResponseCompoundDocument

`func NewGetEventResponseCompoundDocument(data GetEventResponseCollectionCompoundDocumentDataInner, ) *GetEventResponseCompoundDocument`

NewGetEventResponseCompoundDocument instantiates a new GetEventResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventResponseCompoundDocumentWithDefaults

`func NewGetEventResponseCompoundDocumentWithDefaults() *GetEventResponseCompoundDocument`

NewGetEventResponseCompoundDocumentWithDefaults instantiates a new GetEventResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetEventResponseCompoundDocument) GetData() GetEventResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEventResponseCompoundDocument) GetDataOk() (*GetEventResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEventResponseCompoundDocument) SetData(v GetEventResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetEventResponseCompoundDocument) GetIncluded() []GetEventResponseCollectionCompoundDocumentIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetEventResponseCompoundDocument) GetIncludedOk() (*[]GetEventResponseCollectionCompoundDocumentIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetEventResponseCompoundDocument) SetIncluded(v []GetEventResponseCollectionCompoundDocumentIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetEventResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


