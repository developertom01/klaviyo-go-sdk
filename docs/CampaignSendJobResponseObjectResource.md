# CampaignSendJobResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignSendJobEnum**](CampaignSendJobEnum.md) |  | 
**Id** | **string** | The ID of the campaign to send | 
**Attributes** | [**CampaignSendJobResponseObjectResourceAttributes**](CampaignSendJobResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewCampaignSendJobResponseObjectResource

`func NewCampaignSendJobResponseObjectResource(type_ CampaignSendJobEnum, id string, attributes CampaignSendJobResponseObjectResourceAttributes, links ObjectLinks, ) *CampaignSendJobResponseObjectResource`

NewCampaignSendJobResponseObjectResource instantiates a new CampaignSendJobResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSendJobResponseObjectResourceWithDefaults

`func NewCampaignSendJobResponseObjectResourceWithDefaults() *CampaignSendJobResponseObjectResource`

NewCampaignSendJobResponseObjectResourceWithDefaults instantiates a new CampaignSendJobResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignSendJobResponseObjectResource) GetType() CampaignSendJobEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignSendJobResponseObjectResource) GetTypeOk() (*CampaignSendJobEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignSendJobResponseObjectResource) SetType(v CampaignSendJobEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CampaignSendJobResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignSendJobResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignSendJobResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CampaignSendJobResponseObjectResource) GetAttributes() CampaignSendJobResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CampaignSendJobResponseObjectResource) GetAttributesOk() (*CampaignSendJobResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CampaignSendJobResponseObjectResource) SetAttributes(v CampaignSendJobResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *CampaignSendJobResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CampaignSendJobResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CampaignSendJobResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


