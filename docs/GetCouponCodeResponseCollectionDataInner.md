# GetCouponCodeResponseCollectionDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CouponCodeEnum**](CouponCodeEnum.md) |  | 
**Id** | **string** | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with. | 
**Attributes** | [**CouponCodeResponseObjectResourceAttributes**](CouponCodeResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetCouponCodeResponseCollectionDataInnerAllOfRelationships**](GetCouponCodeResponseCollectionDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetCouponCodeResponseCollectionDataInner

`func NewGetCouponCodeResponseCollectionDataInner(type_ CouponCodeEnum, id string, attributes CouponCodeResponseObjectResourceAttributes, links ObjectLinks, ) *GetCouponCodeResponseCollectionDataInner`

NewGetCouponCodeResponseCollectionDataInner instantiates a new GetCouponCodeResponseCollectionDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCouponCodeResponseCollectionDataInnerWithDefaults

`func NewGetCouponCodeResponseCollectionDataInnerWithDefaults() *GetCouponCodeResponseCollectionDataInner`

NewGetCouponCodeResponseCollectionDataInnerWithDefaults instantiates a new GetCouponCodeResponseCollectionDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCouponCodeResponseCollectionDataInner) GetType() CouponCodeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCouponCodeResponseCollectionDataInner) GetTypeOk() (*CouponCodeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCouponCodeResponseCollectionDataInner) SetType(v CouponCodeEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetCouponCodeResponseCollectionDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCouponCodeResponseCollectionDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCouponCodeResponseCollectionDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetCouponCodeResponseCollectionDataInner) GetAttributes() CouponCodeResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetCouponCodeResponseCollectionDataInner) GetAttributesOk() (*CouponCodeResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetCouponCodeResponseCollectionDataInner) SetAttributes(v CouponCodeResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetCouponCodeResponseCollectionDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCouponCodeResponseCollectionDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCouponCodeResponseCollectionDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetCouponCodeResponseCollectionDataInner) GetRelationships() GetCouponCodeResponseCollectionDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetCouponCodeResponseCollectionDataInner) GetRelationshipsOk() (*GetCouponCodeResponseCollectionDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetCouponCodeResponseCollectionDataInner) SetRelationships(v GetCouponCodeResponseCollectionDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetCouponCodeResponseCollectionDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


