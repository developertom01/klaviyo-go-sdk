# CampaignMessagePartialUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignMessageEnum**](CampaignMessageEnum.md) |  | 
**Id** | **string** | The message ID to be retrieved | 
**Attributes** | [**CampaignMessagePartialUpdateQueryResourceObjectAttributes**](CampaignMessagePartialUpdateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewCampaignMessagePartialUpdateQueryResourceObject

`func NewCampaignMessagePartialUpdateQueryResourceObject(type_ CampaignMessageEnum, id string, attributes CampaignMessagePartialUpdateQueryResourceObjectAttributes, ) *CampaignMessagePartialUpdateQueryResourceObject`

NewCampaignMessagePartialUpdateQueryResourceObject instantiates a new CampaignMessagePartialUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignMessagePartialUpdateQueryResourceObjectWithDefaults

`func NewCampaignMessagePartialUpdateQueryResourceObjectWithDefaults() *CampaignMessagePartialUpdateQueryResourceObject`

NewCampaignMessagePartialUpdateQueryResourceObjectWithDefaults instantiates a new CampaignMessagePartialUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignMessagePartialUpdateQueryResourceObject) GetType() CampaignMessageEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignMessagePartialUpdateQueryResourceObject) GetTypeOk() (*CampaignMessageEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignMessagePartialUpdateQueryResourceObject) SetType(v CampaignMessageEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CampaignMessagePartialUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignMessagePartialUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignMessagePartialUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CampaignMessagePartialUpdateQueryResourceObject) GetAttributes() CampaignMessagePartialUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CampaignMessagePartialUpdateQueryResourceObject) GetAttributesOk() (*CampaignMessagePartialUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CampaignMessagePartialUpdateQueryResourceObject) SetAttributes(v CampaignMessagePartialUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


