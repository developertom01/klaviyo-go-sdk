# GetCampaignResponseCollectionCompoundDocumentDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignEnum**](CampaignEnum.md) |  | 
**Id** | **string** | The campaign ID | 
**Attributes** | [**CampaignResponseObjectResourceAttributes**](CampaignResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationships**](GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetCampaignResponseCollectionCompoundDocumentDataInner

`func NewGetCampaignResponseCollectionCompoundDocumentDataInner(type_ CampaignEnum, id string, attributes CampaignResponseObjectResourceAttributes, links ObjectLinks, ) *GetCampaignResponseCollectionCompoundDocumentDataInner`

NewGetCampaignResponseCollectionCompoundDocumentDataInner instantiates a new GetCampaignResponseCollectionCompoundDocumentDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignResponseCollectionCompoundDocumentDataInnerWithDefaults

`func NewGetCampaignResponseCollectionCompoundDocumentDataInnerWithDefaults() *GetCampaignResponseCollectionCompoundDocumentDataInner`

NewGetCampaignResponseCollectionCompoundDocumentDataInnerWithDefaults instantiates a new GetCampaignResponseCollectionCompoundDocumentDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) GetType() CampaignEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) GetTypeOk() (*CampaignEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) SetType(v CampaignEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) GetAttributes() CampaignResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) GetAttributesOk() (*CampaignResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) SetAttributes(v CampaignResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) GetRelationships() GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) GetRelationshipsOk() (*GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) SetRelationships(v GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetCampaignResponseCollectionCompoundDocumentDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


