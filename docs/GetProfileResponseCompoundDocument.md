# GetProfileResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetProfileResponseCompoundDocumentData**](GetProfileResponseCompoundDocumentData.md) |  | 
**Included** | Pointer to [**[]GetProfileResponseCompoundDocumentIncludedInner**](GetProfileResponseCompoundDocumentIncludedInner.md) |  | [optional] 

## Methods

### NewGetProfileResponseCompoundDocument

`func NewGetProfileResponseCompoundDocument(data GetProfileResponseCompoundDocumentData, ) *GetProfileResponseCompoundDocument`

NewGetProfileResponseCompoundDocument instantiates a new GetProfileResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfileResponseCompoundDocumentWithDefaults

`func NewGetProfileResponseCompoundDocumentWithDefaults() *GetProfileResponseCompoundDocument`

NewGetProfileResponseCompoundDocumentWithDefaults instantiates a new GetProfileResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetProfileResponseCompoundDocument) GetData() GetProfileResponseCompoundDocumentData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetProfileResponseCompoundDocument) GetDataOk() (*GetProfileResponseCompoundDocumentData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetProfileResponseCompoundDocument) SetData(v GetProfileResponseCompoundDocumentData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetProfileResponseCompoundDocument) GetIncluded() []GetProfileResponseCompoundDocumentIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetProfileResponseCompoundDocument) GetIncludedOk() (*[]GetProfileResponseCompoundDocumentIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetProfileResponseCompoundDocument) SetIncluded(v []GetProfileResponseCompoundDocumentIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetProfileResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


