# GetCouponCodeCreateJobResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInner**](GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]CouponCodeResponseObjectResource**](CouponCodeResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetCouponCodeCreateJobResponseCompoundDocument

`func NewGetCouponCodeCreateJobResponseCompoundDocument(data GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInner, ) *GetCouponCodeCreateJobResponseCompoundDocument`

NewGetCouponCodeCreateJobResponseCompoundDocument instantiates a new GetCouponCodeCreateJobResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCouponCodeCreateJobResponseCompoundDocumentWithDefaults

`func NewGetCouponCodeCreateJobResponseCompoundDocumentWithDefaults() *GetCouponCodeCreateJobResponseCompoundDocument`

NewGetCouponCodeCreateJobResponseCompoundDocumentWithDefaults instantiates a new GetCouponCodeCreateJobResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCouponCodeCreateJobResponseCompoundDocument) GetData() GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCouponCodeCreateJobResponseCompoundDocument) GetDataOk() (*GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCouponCodeCreateJobResponseCompoundDocument) SetData(v GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetCouponCodeCreateJobResponseCompoundDocument) GetIncluded() []CouponCodeResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCouponCodeCreateJobResponseCompoundDocument) GetIncludedOk() (*[]CouponCodeResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCouponCodeCreateJobResponseCompoundDocument) SetIncluded(v []CouponCodeResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCouponCodeCreateJobResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


