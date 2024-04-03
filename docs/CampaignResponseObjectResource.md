# CampaignResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignEnum**](CampaignEnum.md) |  | 
**Id** | **string** | The campaign ID | 
**Attributes** | [**CampaignResponseObjectResourceAttributes**](CampaignResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewCampaignResponseObjectResource

`func NewCampaignResponseObjectResource(type_ CampaignEnum, id string, attributes CampaignResponseObjectResourceAttributes, links ObjectLinks, ) *CampaignResponseObjectResource`

NewCampaignResponseObjectResource instantiates a new CampaignResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignResponseObjectResourceWithDefaults

`func NewCampaignResponseObjectResourceWithDefaults() *CampaignResponseObjectResource`

NewCampaignResponseObjectResourceWithDefaults instantiates a new CampaignResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignResponseObjectResource) GetType() CampaignEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignResponseObjectResource) GetTypeOk() (*CampaignEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignResponseObjectResource) SetType(v CampaignEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CampaignResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CampaignResponseObjectResource) GetAttributes() CampaignResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CampaignResponseObjectResource) GetAttributesOk() (*CampaignResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CampaignResponseObjectResource) SetAttributes(v CampaignResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *CampaignResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CampaignResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CampaignResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


