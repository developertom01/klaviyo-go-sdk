# ListMemberResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileEnum**](ProfileEnum.md) |  | 
**Id** | Pointer to **NullableString** | Primary key that uniquely identifies this profile. Generated by Klaviyo. | [optional] 
**Attributes** | [**ListMemberResponseObjectResourceAttributes**](ListMemberResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewListMemberResponseObjectResource

`func NewListMemberResponseObjectResource(type_ ProfileEnum, attributes ListMemberResponseObjectResourceAttributes, links ObjectLinks, ) *ListMemberResponseObjectResource`

NewListMemberResponseObjectResource instantiates a new ListMemberResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMemberResponseObjectResourceWithDefaults

`func NewListMemberResponseObjectResourceWithDefaults() *ListMemberResponseObjectResource`

NewListMemberResponseObjectResourceWithDefaults instantiates a new ListMemberResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListMemberResponseObjectResource) GetType() ProfileEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListMemberResponseObjectResource) GetTypeOk() (*ProfileEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListMemberResponseObjectResource) SetType(v ProfileEnum)`

SetType sets Type field to given value.


### GetId

`func (o *ListMemberResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMemberResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMemberResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListMemberResponseObjectResource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ListMemberResponseObjectResource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ListMemberResponseObjectResource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *ListMemberResponseObjectResource) GetAttributes() ListMemberResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ListMemberResponseObjectResource) GetAttributesOk() (*ListMemberResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ListMemberResponseObjectResource) SetAttributes(v ListMemberResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *ListMemberResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListMemberResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListMemberResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


