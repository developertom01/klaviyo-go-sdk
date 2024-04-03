# PostCouponCodeResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CouponCodeEnum**](CouponCodeEnum.md) |  | 
**Id** | **string** | The id of a coupon code is a combination of its unique code and the id of the coupon it is associated with. | 
**Attributes** | [**CouponCodeResponseObjectResourceAttributes**](CouponCodeResponseObjectResourceAttributes.md) |  | 
**Relationships** | Pointer to [**PostCouponCodeResponseDataRelationships**](PostCouponCodeResponseDataRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostCouponCodeResponseData

`func NewPostCouponCodeResponseData(type_ CouponCodeEnum, id string, attributes CouponCodeResponseObjectResourceAttributes, links ObjectLinks, ) *PostCouponCodeResponseData`

NewPostCouponCodeResponseData instantiates a new PostCouponCodeResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCouponCodeResponseDataWithDefaults

`func NewPostCouponCodeResponseDataWithDefaults() *PostCouponCodeResponseData`

NewPostCouponCodeResponseDataWithDefaults instantiates a new PostCouponCodeResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostCouponCodeResponseData) GetType() CouponCodeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCouponCodeResponseData) GetTypeOk() (*CouponCodeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCouponCodeResponseData) SetType(v CouponCodeEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostCouponCodeResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCouponCodeResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCouponCodeResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PostCouponCodeResponseData) GetAttributes() CouponCodeResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostCouponCodeResponseData) GetAttributesOk() (*CouponCodeResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostCouponCodeResponseData) SetAttributes(v CouponCodeResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PostCouponCodeResponseData) GetRelationships() PostCouponCodeResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PostCouponCodeResponseData) GetRelationshipsOk() (*PostCouponCodeResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PostCouponCodeResponseData) SetRelationships(v PostCouponCodeResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PostCouponCodeResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PostCouponCodeResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostCouponCodeResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostCouponCodeResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


