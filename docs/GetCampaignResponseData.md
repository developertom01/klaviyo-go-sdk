# GetCampaignResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignEnum**](CampaignEnum.md) |  | 
**Id** | **string** | The campaign ID | 
**Attributes** | [**CampaignResponseObjectResourceAttributes**](CampaignResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetCampaignResponseDataAllOfRelationships**](GetCampaignResponseDataAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetCampaignResponseData

`func NewGetCampaignResponseData(type_ CampaignEnum, id string, attributes CampaignResponseObjectResourceAttributes, links ObjectLinks, ) *GetCampaignResponseData`

NewGetCampaignResponseData instantiates a new GetCampaignResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignResponseDataWithDefaults

`func NewGetCampaignResponseDataWithDefaults() *GetCampaignResponseData`

NewGetCampaignResponseDataWithDefaults instantiates a new GetCampaignResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCampaignResponseData) GetType() CampaignEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCampaignResponseData) GetTypeOk() (*CampaignEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCampaignResponseData) SetType(v CampaignEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetCampaignResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCampaignResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCampaignResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetCampaignResponseData) GetAttributes() CampaignResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetCampaignResponseData) GetAttributesOk() (*CampaignResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetCampaignResponseData) SetAttributes(v CampaignResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetCampaignResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCampaignResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCampaignResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetCampaignResponseData) GetRelationships() GetCampaignResponseDataAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetCampaignResponseData) GetRelationshipsOk() (*GetCampaignResponseDataAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetCampaignResponseData) SetRelationships(v GetCampaignResponseDataAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetCampaignResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


