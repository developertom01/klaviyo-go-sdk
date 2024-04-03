# TagCampaignOpDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignEnum**](CampaignEnum.md) |  | 
**Id** | **string** | The IDs of the campaigns to link or unlink with the given Tag ID | 

## Methods

### NewTagCampaignOpDataInner

`func NewTagCampaignOpDataInner(type_ CampaignEnum, id string, ) *TagCampaignOpDataInner`

NewTagCampaignOpDataInner instantiates a new TagCampaignOpDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCampaignOpDataInnerWithDefaults

`func NewTagCampaignOpDataInnerWithDefaults() *TagCampaignOpDataInner`

NewTagCampaignOpDataInnerWithDefaults instantiates a new TagCampaignOpDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagCampaignOpDataInner) GetType() CampaignEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagCampaignOpDataInner) GetTypeOk() (*CampaignEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagCampaignOpDataInner) SetType(v CampaignEnum)`

SetType sets Type field to given value.


### GetId

`func (o *TagCampaignOpDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagCampaignOpDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagCampaignOpDataInner) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


