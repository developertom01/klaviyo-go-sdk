# GetListMemberResponseCollectionDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileEnum**](ProfileEnum.md) |  | 
**Id** | Pointer to **NullableString** | Primary key that uniquely identifies this profile. Generated by Klaviyo. | [optional] 
**Attributes** | [**GetProfileResponseDataAllOfAttributes**](GetProfileResponseDataAllOfAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetProfileResponseDataAllOfRelationships**](GetProfileResponseDataAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetListMemberResponseCollectionDataInner

`func NewGetListMemberResponseCollectionDataInner(type_ ProfileEnum, attributes GetProfileResponseDataAllOfAttributes, links ObjectLinks, ) *GetListMemberResponseCollectionDataInner`

NewGetListMemberResponseCollectionDataInner instantiates a new GetListMemberResponseCollectionDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListMemberResponseCollectionDataInnerWithDefaults

`func NewGetListMemberResponseCollectionDataInnerWithDefaults() *GetListMemberResponseCollectionDataInner`

NewGetListMemberResponseCollectionDataInnerWithDefaults instantiates a new GetListMemberResponseCollectionDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetListMemberResponseCollectionDataInner) GetType() ProfileEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetListMemberResponseCollectionDataInner) GetTypeOk() (*ProfileEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetListMemberResponseCollectionDataInner) SetType(v ProfileEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetListMemberResponseCollectionDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetListMemberResponseCollectionDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetListMemberResponseCollectionDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetListMemberResponseCollectionDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GetListMemberResponseCollectionDataInner) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GetListMemberResponseCollectionDataInner) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *GetListMemberResponseCollectionDataInner) GetAttributes() GetProfileResponseDataAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetListMemberResponseCollectionDataInner) GetAttributesOk() (*GetProfileResponseDataAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetListMemberResponseCollectionDataInner) SetAttributes(v GetProfileResponseDataAllOfAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetListMemberResponseCollectionDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetListMemberResponseCollectionDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetListMemberResponseCollectionDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetListMemberResponseCollectionDataInner) GetRelationships() GetProfileResponseDataAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetListMemberResponseCollectionDataInner) GetRelationshipsOk() (*GetProfileResponseDataAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetListMemberResponseCollectionDataInner) SetRelationships(v GetProfileResponseDataAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetListMemberResponseCollectionDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


