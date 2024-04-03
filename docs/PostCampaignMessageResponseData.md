# PostCampaignMessageResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignMessageEnum**](CampaignMessageEnum.md) |  | 
**Id** | **string** | The message ID | 
**Attributes** | [**CampaignMessageResponseObjectResourceAttributes**](CampaignMessageResponseObjectResourceAttributes.md) |  | 
**Relationships** | Pointer to [**GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships**](GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostCampaignMessageResponseData

`func NewPostCampaignMessageResponseData(type_ CampaignMessageEnum, id string, attributes CampaignMessageResponseObjectResourceAttributes, links ObjectLinks, ) *PostCampaignMessageResponseData`

NewPostCampaignMessageResponseData instantiates a new PostCampaignMessageResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCampaignMessageResponseDataWithDefaults

`func NewPostCampaignMessageResponseDataWithDefaults() *PostCampaignMessageResponseData`

NewPostCampaignMessageResponseDataWithDefaults instantiates a new PostCampaignMessageResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostCampaignMessageResponseData) GetType() CampaignMessageEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCampaignMessageResponseData) GetTypeOk() (*CampaignMessageEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCampaignMessageResponseData) SetType(v CampaignMessageEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostCampaignMessageResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCampaignMessageResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCampaignMessageResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PostCampaignMessageResponseData) GetAttributes() CampaignMessageResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostCampaignMessageResponseData) GetAttributesOk() (*CampaignMessageResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostCampaignMessageResponseData) SetAttributes(v CampaignMessageResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PostCampaignMessageResponseData) GetRelationships() GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PostCampaignMessageResponseData) GetRelationshipsOk() (*GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PostCampaignMessageResponseData) SetRelationships(v GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PostCampaignMessageResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PostCampaignMessageResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostCampaignMessageResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostCampaignMessageResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


