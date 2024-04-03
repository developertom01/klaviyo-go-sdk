# GetCampaignMessageResponseCompoundDocumentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignMessageEnum**](CampaignMessageEnum.md) |  | 
**Id** | **string** | The message ID | 
**Attributes** | [**CampaignMessageResponseObjectResourceAttributes**](CampaignMessageResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships**](GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetCampaignMessageResponseCompoundDocumentData

`func NewGetCampaignMessageResponseCompoundDocumentData(type_ CampaignMessageEnum, id string, attributes CampaignMessageResponseObjectResourceAttributes, links ObjectLinks, ) *GetCampaignMessageResponseCompoundDocumentData`

NewGetCampaignMessageResponseCompoundDocumentData instantiates a new GetCampaignMessageResponseCompoundDocumentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignMessageResponseCompoundDocumentDataWithDefaults

`func NewGetCampaignMessageResponseCompoundDocumentDataWithDefaults() *GetCampaignMessageResponseCompoundDocumentData`

NewGetCampaignMessageResponseCompoundDocumentDataWithDefaults instantiates a new GetCampaignMessageResponseCompoundDocumentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCampaignMessageResponseCompoundDocumentData) GetType() CampaignMessageEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCampaignMessageResponseCompoundDocumentData) GetTypeOk() (*CampaignMessageEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCampaignMessageResponseCompoundDocumentData) SetType(v CampaignMessageEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetCampaignMessageResponseCompoundDocumentData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCampaignMessageResponseCompoundDocumentData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCampaignMessageResponseCompoundDocumentData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetCampaignMessageResponseCompoundDocumentData) GetAttributes() CampaignMessageResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetCampaignMessageResponseCompoundDocumentData) GetAttributesOk() (*CampaignMessageResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetCampaignMessageResponseCompoundDocumentData) SetAttributes(v CampaignMessageResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetCampaignMessageResponseCompoundDocumentData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCampaignMessageResponseCompoundDocumentData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCampaignMessageResponseCompoundDocumentData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetCampaignMessageResponseCompoundDocumentData) GetRelationships() GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetCampaignMessageResponseCompoundDocumentData) GetRelationshipsOk() (*GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetCampaignMessageResponseCompoundDocumentData) SetRelationships(v GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetCampaignMessageResponseCompoundDocumentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


