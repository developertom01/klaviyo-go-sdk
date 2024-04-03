# GetSegmentRetrieveResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetSegmentRetrieveResponseCompoundDocumentData**](GetSegmentRetrieveResponseCompoundDocumentData.md) |  | 
**Included** | Pointer to [**[]TagResponseObjectResource**](TagResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetSegmentRetrieveResponseCompoundDocument

`func NewGetSegmentRetrieveResponseCompoundDocument(data GetSegmentRetrieveResponseCompoundDocumentData, ) *GetSegmentRetrieveResponseCompoundDocument`

NewGetSegmentRetrieveResponseCompoundDocument instantiates a new GetSegmentRetrieveResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSegmentRetrieveResponseCompoundDocumentWithDefaults

`func NewGetSegmentRetrieveResponseCompoundDocumentWithDefaults() *GetSegmentRetrieveResponseCompoundDocument`

NewGetSegmentRetrieveResponseCompoundDocumentWithDefaults instantiates a new GetSegmentRetrieveResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSegmentRetrieveResponseCompoundDocument) GetData() GetSegmentRetrieveResponseCompoundDocumentData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSegmentRetrieveResponseCompoundDocument) GetDataOk() (*GetSegmentRetrieveResponseCompoundDocumentData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSegmentRetrieveResponseCompoundDocument) SetData(v GetSegmentRetrieveResponseCompoundDocumentData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetSegmentRetrieveResponseCompoundDocument) GetIncluded() []TagResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetSegmentRetrieveResponseCompoundDocument) GetIncludedOk() (*[]TagResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetSegmentRetrieveResponseCompoundDocument) SetIncluded(v []TagResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetSegmentRetrieveResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


