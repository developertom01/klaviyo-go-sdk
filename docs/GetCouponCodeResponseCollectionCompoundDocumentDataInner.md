# GetCouponCodeResponseCollectionCompoundDocumentDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CouponCodeEnum**](CouponCodeEnum.md) |  | 
**Id** | **string** | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with. | 
**Attributes** | [**CouponCodeResponseObjectResourceAttributes**](CouponCodeResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships**](GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetCouponCodeResponseCollectionCompoundDocumentDataInner

`func NewGetCouponCodeResponseCollectionCompoundDocumentDataInner(type_ CouponCodeEnum, id string, attributes CouponCodeResponseObjectResourceAttributes, links ObjectLinks, ) *GetCouponCodeResponseCollectionCompoundDocumentDataInner`

NewGetCouponCodeResponseCollectionCompoundDocumentDataInner instantiates a new GetCouponCodeResponseCollectionCompoundDocumentDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCouponCodeResponseCollectionCompoundDocumentDataInnerWithDefaults

`func NewGetCouponCodeResponseCollectionCompoundDocumentDataInnerWithDefaults() *GetCouponCodeResponseCollectionCompoundDocumentDataInner`

NewGetCouponCodeResponseCollectionCompoundDocumentDataInnerWithDefaults instantiates a new GetCouponCodeResponseCollectionCompoundDocumentDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) GetType() CouponCodeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) GetTypeOk() (*CouponCodeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) SetType(v CouponCodeEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) GetAttributes() CouponCodeResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) GetAttributesOk() (*CouponCodeResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) SetAttributes(v CouponCodeResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) GetRelationships() GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) GetRelationshipsOk() (*GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) SetRelationships(v GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetCouponCodeResponseCollectionCompoundDocumentDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


