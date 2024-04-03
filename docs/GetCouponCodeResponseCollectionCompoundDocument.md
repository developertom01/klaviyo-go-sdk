# GetCouponCodeResponseCollectionCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetCouponCodeResponseCollectionCompoundDocumentDataInner**](GetCouponCodeResponseCollectionCompoundDocumentDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 
**Included** | Pointer to [**[]CouponResponseObjectResource**](CouponResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetCouponCodeResponseCollectionCompoundDocument

`func NewGetCouponCodeResponseCollectionCompoundDocument(data []GetCouponCodeResponseCollectionCompoundDocumentDataInner, links CollectionLinks, ) *GetCouponCodeResponseCollectionCompoundDocument`

NewGetCouponCodeResponseCollectionCompoundDocument instantiates a new GetCouponCodeResponseCollectionCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCouponCodeResponseCollectionCompoundDocumentWithDefaults

`func NewGetCouponCodeResponseCollectionCompoundDocumentWithDefaults() *GetCouponCodeResponseCollectionCompoundDocument`

NewGetCouponCodeResponseCollectionCompoundDocumentWithDefaults instantiates a new GetCouponCodeResponseCollectionCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCouponCodeResponseCollectionCompoundDocument) GetData() []GetCouponCodeResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCouponCodeResponseCollectionCompoundDocument) GetDataOk() (*[]GetCouponCodeResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCouponCodeResponseCollectionCompoundDocument) SetData(v []GetCouponCodeResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetCouponCodeResponseCollectionCompoundDocument) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCouponCodeResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCouponCodeResponseCollectionCompoundDocument) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.


### GetIncluded

`func (o *GetCouponCodeResponseCollectionCompoundDocument) GetIncluded() []CouponResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCouponCodeResponseCollectionCompoundDocument) GetIncludedOk() (*[]CouponResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCouponCodeResponseCollectionCompoundDocument) SetIncluded(v []CouponResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCouponCodeResponseCollectionCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


