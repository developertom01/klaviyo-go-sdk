# GetCouponCodeResponseCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetCouponCodeResponseCollectionDataInner**](GetCouponCodeResponseCollectionDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewGetCouponCodeResponseCollection

`func NewGetCouponCodeResponseCollection(data []GetCouponCodeResponseCollectionDataInner, links CollectionLinks, ) *GetCouponCodeResponseCollection`

NewGetCouponCodeResponseCollection instantiates a new GetCouponCodeResponseCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCouponCodeResponseCollectionWithDefaults

`func NewGetCouponCodeResponseCollectionWithDefaults() *GetCouponCodeResponseCollection`

NewGetCouponCodeResponseCollectionWithDefaults instantiates a new GetCouponCodeResponseCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCouponCodeResponseCollection) GetData() []GetCouponCodeResponseCollectionDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCouponCodeResponseCollection) GetDataOk() (*[]GetCouponCodeResponseCollectionDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCouponCodeResponseCollection) SetData(v []GetCouponCodeResponseCollectionDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetCouponCodeResponseCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCouponCodeResponseCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCouponCodeResponseCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


