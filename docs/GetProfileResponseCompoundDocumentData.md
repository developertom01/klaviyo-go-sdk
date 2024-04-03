# GetProfileResponseCompoundDocumentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileEnum**](ProfileEnum.md) |  | 
**Id** | Pointer to **NullableString** | Primary key that uniquely identifies this profile. Generated by Klaviyo. | [optional] 
**Attributes** | [**GetProfileResponseDataAllOfAttributes**](GetProfileResponseDataAllOfAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetProfileResponseCompoundDocumentDataAllOfRelationships**](GetProfileResponseCompoundDocumentDataAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetProfileResponseCompoundDocumentData

`func NewGetProfileResponseCompoundDocumentData(type_ ProfileEnum, attributes GetProfileResponseDataAllOfAttributes, links ObjectLinks, ) *GetProfileResponseCompoundDocumentData`

NewGetProfileResponseCompoundDocumentData instantiates a new GetProfileResponseCompoundDocumentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfileResponseCompoundDocumentDataWithDefaults

`func NewGetProfileResponseCompoundDocumentDataWithDefaults() *GetProfileResponseCompoundDocumentData`

NewGetProfileResponseCompoundDocumentDataWithDefaults instantiates a new GetProfileResponseCompoundDocumentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetProfileResponseCompoundDocumentData) GetType() ProfileEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetProfileResponseCompoundDocumentData) GetTypeOk() (*ProfileEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetProfileResponseCompoundDocumentData) SetType(v ProfileEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetProfileResponseCompoundDocumentData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetProfileResponseCompoundDocumentData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetProfileResponseCompoundDocumentData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetProfileResponseCompoundDocumentData) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GetProfileResponseCompoundDocumentData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GetProfileResponseCompoundDocumentData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *GetProfileResponseCompoundDocumentData) GetAttributes() GetProfileResponseDataAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetProfileResponseCompoundDocumentData) GetAttributesOk() (*GetProfileResponseDataAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetProfileResponseCompoundDocumentData) SetAttributes(v GetProfileResponseDataAllOfAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetProfileResponseCompoundDocumentData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetProfileResponseCompoundDocumentData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetProfileResponseCompoundDocumentData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetProfileResponseCompoundDocumentData) GetRelationships() GetProfileResponseCompoundDocumentDataAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetProfileResponseCompoundDocumentData) GetRelationshipsOk() (*GetProfileResponseCompoundDocumentDataAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetProfileResponseCompoundDocumentData) SetRelationships(v GetProfileResponseCompoundDocumentDataAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetProfileResponseCompoundDocumentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

