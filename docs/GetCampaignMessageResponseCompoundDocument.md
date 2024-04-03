# GetCampaignMessageResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetCampaignMessageResponseCompoundDocumentData**](GetCampaignMessageResponseCompoundDocumentData.md) |  | 
**Included** | Pointer to [**[]GetCampaignMessageResponseCompoundDocumentIncludedInner**](GetCampaignMessageResponseCompoundDocumentIncludedInner.md) |  | [optional] 

## Methods

### NewGetCampaignMessageResponseCompoundDocument

`func NewGetCampaignMessageResponseCompoundDocument(data GetCampaignMessageResponseCompoundDocumentData, ) *GetCampaignMessageResponseCompoundDocument`

NewGetCampaignMessageResponseCompoundDocument instantiates a new GetCampaignMessageResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignMessageResponseCompoundDocumentWithDefaults

`func NewGetCampaignMessageResponseCompoundDocumentWithDefaults() *GetCampaignMessageResponseCompoundDocument`

NewGetCampaignMessageResponseCompoundDocumentWithDefaults instantiates a new GetCampaignMessageResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCampaignMessageResponseCompoundDocument) GetData() GetCampaignMessageResponseCompoundDocumentData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCampaignMessageResponseCompoundDocument) GetDataOk() (*GetCampaignMessageResponseCompoundDocumentData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCampaignMessageResponseCompoundDocument) SetData(v GetCampaignMessageResponseCompoundDocumentData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetCampaignMessageResponseCompoundDocument) GetIncluded() []GetCampaignMessageResponseCompoundDocumentIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCampaignMessageResponseCompoundDocument) GetIncludedOk() (*[]GetCampaignMessageResponseCompoundDocumentIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCampaignMessageResponseCompoundDocument) SetIncluded(v []GetCampaignMessageResponseCompoundDocumentIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCampaignMessageResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


