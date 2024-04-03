# CampaignRecipientEstimationJobResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignRecipientEstimationJobEnum**](CampaignRecipientEstimationJobEnum.md) |  | 
**Id** | **string** | The ID of the campaign used for estimating recipients | 
**Attributes** | [**CampaignRecipientEstimationJobResponseObjectResourceAttributes**](CampaignRecipientEstimationJobResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewCampaignRecipientEstimationJobResponseObjectResource

`func NewCampaignRecipientEstimationJobResponseObjectResource(type_ CampaignRecipientEstimationJobEnum, id string, attributes CampaignRecipientEstimationJobResponseObjectResourceAttributes, links ObjectLinks, ) *CampaignRecipientEstimationJobResponseObjectResource`

NewCampaignRecipientEstimationJobResponseObjectResource instantiates a new CampaignRecipientEstimationJobResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignRecipientEstimationJobResponseObjectResourceWithDefaults

`func NewCampaignRecipientEstimationJobResponseObjectResourceWithDefaults() *CampaignRecipientEstimationJobResponseObjectResource`

NewCampaignRecipientEstimationJobResponseObjectResourceWithDefaults instantiates a new CampaignRecipientEstimationJobResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignRecipientEstimationJobResponseObjectResource) GetType() CampaignRecipientEstimationJobEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignRecipientEstimationJobResponseObjectResource) GetTypeOk() (*CampaignRecipientEstimationJobEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignRecipientEstimationJobResponseObjectResource) SetType(v CampaignRecipientEstimationJobEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CampaignRecipientEstimationJobResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignRecipientEstimationJobResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignRecipientEstimationJobResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CampaignRecipientEstimationJobResponseObjectResource) GetAttributes() CampaignRecipientEstimationJobResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CampaignRecipientEstimationJobResponseObjectResource) GetAttributesOk() (*CampaignRecipientEstimationJobResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CampaignRecipientEstimationJobResponseObjectResource) SetAttributes(v CampaignRecipientEstimationJobResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *CampaignRecipientEstimationJobResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CampaignRecipientEstimationJobResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CampaignRecipientEstimationJobResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


