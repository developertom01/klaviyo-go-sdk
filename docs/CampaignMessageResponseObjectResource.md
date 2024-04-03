# CampaignMessageResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignMessageEnum**](CampaignMessageEnum.md) |  | 
**Id** | **string** | The message ID | 
**Attributes** | [**CampaignMessageResponseObjectResourceAttributes**](CampaignMessageResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewCampaignMessageResponseObjectResource

`func NewCampaignMessageResponseObjectResource(type_ CampaignMessageEnum, id string, attributes CampaignMessageResponseObjectResourceAttributes, links ObjectLinks, ) *CampaignMessageResponseObjectResource`

NewCampaignMessageResponseObjectResource instantiates a new CampaignMessageResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignMessageResponseObjectResourceWithDefaults

`func NewCampaignMessageResponseObjectResourceWithDefaults() *CampaignMessageResponseObjectResource`

NewCampaignMessageResponseObjectResourceWithDefaults instantiates a new CampaignMessageResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignMessageResponseObjectResource) GetType() CampaignMessageEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignMessageResponseObjectResource) GetTypeOk() (*CampaignMessageEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignMessageResponseObjectResource) SetType(v CampaignMessageEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CampaignMessageResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignMessageResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignMessageResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CampaignMessageResponseObjectResource) GetAttributes() CampaignMessageResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CampaignMessageResponseObjectResource) GetAttributesOk() (*CampaignMessageResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CampaignMessageResponseObjectResource) SetAttributes(v CampaignMessageResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *CampaignMessageResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CampaignMessageResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CampaignMessageResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


