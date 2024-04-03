# GetCampaignResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetCampaignResponseCollectionCompoundDocumentDataInner**](GetCampaignResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]GetCampaignResponseCollectionCompoundDocumentIncludedInner**](GetCampaignResponseCollectionCompoundDocumentIncludedInner.md) |  | [optional] 

## Methods

### NewGetCampaignResponseCompoundDocument

`func NewGetCampaignResponseCompoundDocument(data GetCampaignResponseCollectionCompoundDocumentDataInner, ) *GetCampaignResponseCompoundDocument`

NewGetCampaignResponseCompoundDocument instantiates a new GetCampaignResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignResponseCompoundDocumentWithDefaults

`func NewGetCampaignResponseCompoundDocumentWithDefaults() *GetCampaignResponseCompoundDocument`

NewGetCampaignResponseCompoundDocumentWithDefaults instantiates a new GetCampaignResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCampaignResponseCompoundDocument) GetData() GetCampaignResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCampaignResponseCompoundDocument) GetDataOk() (*GetCampaignResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCampaignResponseCompoundDocument) SetData(v GetCampaignResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetCampaignResponseCompoundDocument) GetIncluded() []GetCampaignResponseCollectionCompoundDocumentIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCampaignResponseCompoundDocument) GetIncludedOk() (*[]GetCampaignResponseCollectionCompoundDocumentIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCampaignResponseCompoundDocument) SetIncluded(v []GetCampaignResponseCollectionCompoundDocumentIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCampaignResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


