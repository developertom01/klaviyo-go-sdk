# GetCatalogCategoryResponseCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetCatalogCategoryResponseCollectionDataInner**](GetCatalogCategoryResponseCollectionDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewGetCatalogCategoryResponseCollection

`func NewGetCatalogCategoryResponseCollection(data []GetCatalogCategoryResponseCollectionDataInner, links CollectionLinks, ) *GetCatalogCategoryResponseCollection`

NewGetCatalogCategoryResponseCollection instantiates a new GetCatalogCategoryResponseCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogCategoryResponseCollectionWithDefaults

`func NewGetCatalogCategoryResponseCollectionWithDefaults() *GetCatalogCategoryResponseCollection`

NewGetCatalogCategoryResponseCollectionWithDefaults instantiates a new GetCatalogCategoryResponseCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCatalogCategoryResponseCollection) GetData() []GetCatalogCategoryResponseCollectionDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCatalogCategoryResponseCollection) GetDataOk() (*[]GetCatalogCategoryResponseCollectionDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCatalogCategoryResponseCollection) SetData(v []GetCatalogCategoryResponseCollectionDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetCatalogCategoryResponseCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCatalogCategoryResponseCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCatalogCategoryResponseCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


