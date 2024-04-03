# GetCampaignResponseCollectionCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetCampaignResponseCollectionCompoundDocumentDataInner**](GetCampaignResponseCollectionCompoundDocumentDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 
**Included** | Pointer to [**[]GetCampaignResponseCollectionCompoundDocumentIncludedInner**](GetCampaignResponseCollectionCompoundDocumentIncludedInner.md) |  | [optional] 

## Methods

### NewGetCampaignResponseCollectionCompoundDocument

`func NewGetCampaignResponseCollectionCompoundDocument(data []GetCampaignResponseCollectionCompoundDocumentDataInner, links CollectionLinks, ) *GetCampaignResponseCollectionCompoundDocument`

NewGetCampaignResponseCollectionCompoundDocument instantiates a new GetCampaignResponseCollectionCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignResponseCollectionCompoundDocumentWithDefaults

`func NewGetCampaignResponseCollectionCompoundDocumentWithDefaults() *GetCampaignResponseCollectionCompoundDocument`

NewGetCampaignResponseCollectionCompoundDocumentWithDefaults instantiates a new GetCampaignResponseCollectionCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCampaignResponseCollectionCompoundDocument) GetData() []GetCampaignResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCampaignResponseCollectionCompoundDocument) GetDataOk() (*[]GetCampaignResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCampaignResponseCollectionCompoundDocument) SetData(v []GetCampaignResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetCampaignResponseCollectionCompoundDocument) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCampaignResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCampaignResponseCollectionCompoundDocument) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.


### GetIncluded

`func (o *GetCampaignResponseCollectionCompoundDocument) GetIncluded() []GetCampaignResponseCollectionCompoundDocumentIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCampaignResponseCollectionCompoundDocument) GetIncludedOk() (*[]GetCampaignResponseCollectionCompoundDocumentIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCampaignResponseCollectionCompoundDocument) SetIncluded(v []GetCampaignResponseCollectionCompoundDocumentIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCampaignResponseCollectionCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


