# GetListRetrieveResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetListRetrieveResponseCompoundDocumentData**](GetListRetrieveResponseCompoundDocumentData.md) |  | 
**Included** | Pointer to [**[]TagResponseObjectResource**](TagResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetListRetrieveResponseCompoundDocument

`func NewGetListRetrieveResponseCompoundDocument(data GetListRetrieveResponseCompoundDocumentData, ) *GetListRetrieveResponseCompoundDocument`

NewGetListRetrieveResponseCompoundDocument instantiates a new GetListRetrieveResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListRetrieveResponseCompoundDocumentWithDefaults

`func NewGetListRetrieveResponseCompoundDocumentWithDefaults() *GetListRetrieveResponseCompoundDocument`

NewGetListRetrieveResponseCompoundDocumentWithDefaults instantiates a new GetListRetrieveResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetListRetrieveResponseCompoundDocument) GetData() GetListRetrieveResponseCompoundDocumentData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetListRetrieveResponseCompoundDocument) GetDataOk() (*GetListRetrieveResponseCompoundDocumentData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetListRetrieveResponseCompoundDocument) SetData(v GetListRetrieveResponseCompoundDocumentData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetListRetrieveResponseCompoundDocument) GetIncluded() []TagResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetListRetrieveResponseCompoundDocument) GetIncludedOk() (*[]TagResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetListRetrieveResponseCompoundDocument) SetIncluded(v []TagResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetListRetrieveResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


