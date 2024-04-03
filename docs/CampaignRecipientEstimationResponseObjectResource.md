# CampaignRecipientEstimationResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignRecipientEstimationEnum**](CampaignRecipientEstimationEnum.md) |  | 
**Id** | **string** | The ID of the campaign for which to get the estimated number of recipients | 
**Attributes** | [**CampaignRecipientEstimationResponseObjectResourceAttributes**](CampaignRecipientEstimationResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewCampaignRecipientEstimationResponseObjectResource

`func NewCampaignRecipientEstimationResponseObjectResource(type_ CampaignRecipientEstimationEnum, id string, attributes CampaignRecipientEstimationResponseObjectResourceAttributes, links ObjectLinks, ) *CampaignRecipientEstimationResponseObjectResource`

NewCampaignRecipientEstimationResponseObjectResource instantiates a new CampaignRecipientEstimationResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignRecipientEstimationResponseObjectResourceWithDefaults

`func NewCampaignRecipientEstimationResponseObjectResourceWithDefaults() *CampaignRecipientEstimationResponseObjectResource`

NewCampaignRecipientEstimationResponseObjectResourceWithDefaults instantiates a new CampaignRecipientEstimationResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignRecipientEstimationResponseObjectResource) GetType() CampaignRecipientEstimationEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignRecipientEstimationResponseObjectResource) GetTypeOk() (*CampaignRecipientEstimationEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignRecipientEstimationResponseObjectResource) SetType(v CampaignRecipientEstimationEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CampaignRecipientEstimationResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignRecipientEstimationResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignRecipientEstimationResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CampaignRecipientEstimationResponseObjectResource) GetAttributes() CampaignRecipientEstimationResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CampaignRecipientEstimationResponseObjectResource) GetAttributesOk() (*CampaignRecipientEstimationResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CampaignRecipientEstimationResponseObjectResource) SetAttributes(v CampaignRecipientEstimationResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *CampaignRecipientEstimationResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CampaignRecipientEstimationResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CampaignRecipientEstimationResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


