# PostProfileResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileEnum**](ProfileEnum.md) |  | 
**Id** | Pointer to **NullableString** | Primary key that uniquely identifies this profile. Generated by Klaviyo. | [optional] 
**Attributes** | [**PostProfileResponseDataAttributes**](PostProfileResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**GetProfileResponseCompoundDocumentDataAllOfRelationships**](GetProfileResponseCompoundDocumentDataAllOfRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostProfileResponseData

`func NewPostProfileResponseData(type_ ProfileEnum, attributes PostProfileResponseDataAttributes, links ObjectLinks, ) *PostProfileResponseData`

NewPostProfileResponseData instantiates a new PostProfileResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostProfileResponseDataWithDefaults

`func NewPostProfileResponseDataWithDefaults() *PostProfileResponseData`

NewPostProfileResponseDataWithDefaults instantiates a new PostProfileResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostProfileResponseData) GetType() ProfileEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostProfileResponseData) GetTypeOk() (*ProfileEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostProfileResponseData) SetType(v ProfileEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostProfileResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostProfileResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostProfileResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostProfileResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PostProfileResponseData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PostProfileResponseData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PostProfileResponseData) GetAttributes() PostProfileResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostProfileResponseData) GetAttributesOk() (*PostProfileResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostProfileResponseData) SetAttributes(v PostProfileResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PostProfileResponseData) GetRelationships() GetProfileResponseCompoundDocumentDataAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PostProfileResponseData) GetRelationshipsOk() (*GetProfileResponseCompoundDocumentDataAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PostProfileResponseData) SetRelationships(v GetProfileResponseCompoundDocumentDataAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PostProfileResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PostProfileResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostProfileResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostProfileResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


