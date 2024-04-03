# GetCouponCodeResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetCouponCodeResponseCollectionCompoundDocumentDataInner**](GetCouponCodeResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]CouponResponseObjectResource**](CouponResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetCouponCodeResponseCompoundDocument

`func NewGetCouponCodeResponseCompoundDocument(data GetCouponCodeResponseCollectionCompoundDocumentDataInner, ) *GetCouponCodeResponseCompoundDocument`

NewGetCouponCodeResponseCompoundDocument instantiates a new GetCouponCodeResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCouponCodeResponseCompoundDocumentWithDefaults

`func NewGetCouponCodeResponseCompoundDocumentWithDefaults() *GetCouponCodeResponseCompoundDocument`

NewGetCouponCodeResponseCompoundDocumentWithDefaults instantiates a new GetCouponCodeResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCouponCodeResponseCompoundDocument) GetData() GetCouponCodeResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCouponCodeResponseCompoundDocument) GetDataOk() (*GetCouponCodeResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCouponCodeResponseCompoundDocument) SetData(v GetCouponCodeResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetCouponCodeResponseCompoundDocument) GetIncluded() []CouponResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCouponCodeResponseCompoundDocument) GetIncludedOk() (*[]CouponResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCouponCodeResponseCompoundDocument) SetIncluded(v []CouponResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCouponCodeResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


