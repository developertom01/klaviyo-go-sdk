# GetCampaignMessageResponseCollectionCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetCampaignMessageResponseCompoundDocumentData**](GetCampaignMessageResponseCompoundDocumentData.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 
**Included** | Pointer to [**[]GetCampaignMessageResponseCompoundDocumentIncludedInner**](GetCampaignMessageResponseCompoundDocumentIncludedInner.md) |  | [optional] 

## Methods

### NewGetCampaignMessageResponseCollectionCompoundDocument

`func NewGetCampaignMessageResponseCollectionCompoundDocument(data []GetCampaignMessageResponseCompoundDocumentData, links CollectionLinks, ) *GetCampaignMessageResponseCollectionCompoundDocument`

NewGetCampaignMessageResponseCollectionCompoundDocument instantiates a new GetCampaignMessageResponseCollectionCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignMessageResponseCollectionCompoundDocumentWithDefaults

`func NewGetCampaignMessageResponseCollectionCompoundDocumentWithDefaults() *GetCampaignMessageResponseCollectionCompoundDocument`

NewGetCampaignMessageResponseCollectionCompoundDocumentWithDefaults instantiates a new GetCampaignMessageResponseCollectionCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetData() []GetCampaignMessageResponseCompoundDocumentData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetDataOk() (*[]GetCampaignMessageResponseCompoundDocumentData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCampaignMessageResponseCollectionCompoundDocument) SetData(v []GetCampaignMessageResponseCompoundDocumentData)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCampaignMessageResponseCollectionCompoundDocument) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.


### GetIncluded

`func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetIncluded() []GetCampaignMessageResponseCompoundDocumentIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCampaignMessageResponseCollectionCompoundDocument) GetIncludedOk() (*[]GetCampaignMessageResponseCompoundDocumentIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCampaignMessageResponseCollectionCompoundDocument) SetIncluded(v []GetCampaignMessageResponseCompoundDocumentIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCampaignMessageResponseCollectionCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


