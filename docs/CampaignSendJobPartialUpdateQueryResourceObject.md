# CampaignSendJobPartialUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignSendJobEnum**](CampaignSendJobEnum.md) |  | 
**Id** | **string** | The ID of the currently sending campaign to cancel or revert | 
**Attributes** | [**CampaignSendJobPartialUpdateQueryResourceObjectAttributes**](CampaignSendJobPartialUpdateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewCampaignSendJobPartialUpdateQueryResourceObject

`func NewCampaignSendJobPartialUpdateQueryResourceObject(type_ CampaignSendJobEnum, id string, attributes CampaignSendJobPartialUpdateQueryResourceObjectAttributes, ) *CampaignSendJobPartialUpdateQueryResourceObject`

NewCampaignSendJobPartialUpdateQueryResourceObject instantiates a new CampaignSendJobPartialUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSendJobPartialUpdateQueryResourceObjectWithDefaults

`func NewCampaignSendJobPartialUpdateQueryResourceObjectWithDefaults() *CampaignSendJobPartialUpdateQueryResourceObject`

NewCampaignSendJobPartialUpdateQueryResourceObjectWithDefaults instantiates a new CampaignSendJobPartialUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignSendJobPartialUpdateQueryResourceObject) GetType() CampaignSendJobEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignSendJobPartialUpdateQueryResourceObject) GetTypeOk() (*CampaignSendJobEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignSendJobPartialUpdateQueryResourceObject) SetType(v CampaignSendJobEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CampaignSendJobPartialUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignSendJobPartialUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignSendJobPartialUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CampaignSendJobPartialUpdateQueryResourceObject) GetAttributes() CampaignSendJobPartialUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CampaignSendJobPartialUpdateQueryResourceObject) GetAttributesOk() (*CampaignSendJobPartialUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CampaignSendJobPartialUpdateQueryResourceObject) SetAttributes(v CampaignSendJobPartialUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


