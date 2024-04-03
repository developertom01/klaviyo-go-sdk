# CampaignMessageAssignTemplateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignMessageEnum**](CampaignMessageEnum.md) |  | 
**Id** | **string** | The message ID to be assigned to | 
**Relationships** | [**CampaignMessageAssignTemplateQueryResourceObjectRelationships**](CampaignMessageAssignTemplateQueryResourceObjectRelationships.md) |  | 

## Methods

### NewCampaignMessageAssignTemplateQueryResourceObject

`func NewCampaignMessageAssignTemplateQueryResourceObject(type_ CampaignMessageEnum, id string, relationships CampaignMessageAssignTemplateQueryResourceObjectRelationships, ) *CampaignMessageAssignTemplateQueryResourceObject`

NewCampaignMessageAssignTemplateQueryResourceObject instantiates a new CampaignMessageAssignTemplateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignMessageAssignTemplateQueryResourceObjectWithDefaults

`func NewCampaignMessageAssignTemplateQueryResourceObjectWithDefaults() *CampaignMessageAssignTemplateQueryResourceObject`

NewCampaignMessageAssignTemplateQueryResourceObjectWithDefaults instantiates a new CampaignMessageAssignTemplateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignMessageAssignTemplateQueryResourceObject) GetType() CampaignMessageEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignMessageAssignTemplateQueryResourceObject) GetTypeOk() (*CampaignMessageEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignMessageAssignTemplateQueryResourceObject) SetType(v CampaignMessageEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CampaignMessageAssignTemplateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignMessageAssignTemplateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignMessageAssignTemplateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *CampaignMessageAssignTemplateQueryResourceObject) GetRelationships() CampaignMessageAssignTemplateQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CampaignMessageAssignTemplateQueryResourceObject) GetRelationshipsOk() (*CampaignMessageAssignTemplateQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CampaignMessageAssignTemplateQueryResourceObject) SetRelationships(v CampaignMessageAssignTemplateQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


