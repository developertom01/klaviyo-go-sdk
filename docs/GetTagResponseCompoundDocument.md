# GetTagResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetTagResponseCollectionCompoundDocumentDataInner**](GetTagResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]TagGroupResponseObjectResource**](TagGroupResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetTagResponseCompoundDocument

`func NewGetTagResponseCompoundDocument(data GetTagResponseCollectionCompoundDocumentDataInner, ) *GetTagResponseCompoundDocument`

NewGetTagResponseCompoundDocument instantiates a new GetTagResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTagResponseCompoundDocumentWithDefaults

`func NewGetTagResponseCompoundDocumentWithDefaults() *GetTagResponseCompoundDocument`

NewGetTagResponseCompoundDocumentWithDefaults instantiates a new GetTagResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTagResponseCompoundDocument) GetData() GetTagResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTagResponseCompoundDocument) GetDataOk() (*GetTagResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTagResponseCompoundDocument) SetData(v GetTagResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetTagResponseCompoundDocument) GetIncluded() []TagGroupResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetTagResponseCompoundDocument) GetIncludedOk() (*[]TagGroupResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetTagResponseCompoundDocument) SetIncluded(v []TagGroupResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetTagResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


