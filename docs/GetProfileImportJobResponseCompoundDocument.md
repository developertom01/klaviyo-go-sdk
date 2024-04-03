# GetProfileImportJobResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetProfileImportJobResponseCollectionCompoundDocumentDataInner**](GetProfileImportJobResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]ListResponseObjectResource**](ListResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetProfileImportJobResponseCompoundDocument

`func NewGetProfileImportJobResponseCompoundDocument(data GetProfileImportJobResponseCollectionCompoundDocumentDataInner, ) *GetProfileImportJobResponseCompoundDocument`

NewGetProfileImportJobResponseCompoundDocument instantiates a new GetProfileImportJobResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfileImportJobResponseCompoundDocumentWithDefaults

`func NewGetProfileImportJobResponseCompoundDocumentWithDefaults() *GetProfileImportJobResponseCompoundDocument`

NewGetProfileImportJobResponseCompoundDocumentWithDefaults instantiates a new GetProfileImportJobResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetProfileImportJobResponseCompoundDocument) GetData() GetProfileImportJobResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetProfileImportJobResponseCompoundDocument) GetDataOk() (*GetProfileImportJobResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetProfileImportJobResponseCompoundDocument) SetData(v GetProfileImportJobResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetProfileImportJobResponseCompoundDocument) GetIncluded() []ListResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetProfileImportJobResponseCompoundDocument) GetIncludedOk() (*[]ListResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetProfileImportJobResponseCompoundDocument) SetIncluded(v []ListResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetProfileImportJobResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


