# CampaignCloneQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignEnum**](CampaignEnum.md) |  | 
**Id** | **string** | The campaign ID to be cloned | 
**Attributes** | [**CampaignCloneQueryResourceObjectAttributes**](CampaignCloneQueryResourceObjectAttributes.md) |  | 

## Methods

### NewCampaignCloneQueryResourceObject

`func NewCampaignCloneQueryResourceObject(type_ CampaignEnum, id string, attributes CampaignCloneQueryResourceObjectAttributes, ) *CampaignCloneQueryResourceObject`

NewCampaignCloneQueryResourceObject instantiates a new CampaignCloneQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCloneQueryResourceObjectWithDefaults

`func NewCampaignCloneQueryResourceObjectWithDefaults() *CampaignCloneQueryResourceObject`

NewCampaignCloneQueryResourceObjectWithDefaults instantiates a new CampaignCloneQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignCloneQueryResourceObject) GetType() CampaignEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignCloneQueryResourceObject) GetTypeOk() (*CampaignEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignCloneQueryResourceObject) SetType(v CampaignEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CampaignCloneQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignCloneQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignCloneQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CampaignCloneQueryResourceObject) GetAttributes() CampaignCloneQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CampaignCloneQueryResourceObject) GetAttributesOk() (*CampaignCloneQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CampaignCloneQueryResourceObject) SetAttributes(v CampaignCloneQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


