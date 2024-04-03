# FlowActionResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowActionEnum**](FlowActionEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**FlowActionResponseObjectResourceAttributes**](FlowActionResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewFlowActionResponseObjectResource

`func NewFlowActionResponseObjectResource(type_ FlowActionEnum, id string, attributes FlowActionResponseObjectResourceAttributes, links ObjectLinks, ) *FlowActionResponseObjectResource`

NewFlowActionResponseObjectResource instantiates a new FlowActionResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowActionResponseObjectResourceWithDefaults

`func NewFlowActionResponseObjectResourceWithDefaults() *FlowActionResponseObjectResource`

NewFlowActionResponseObjectResourceWithDefaults instantiates a new FlowActionResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlowActionResponseObjectResource) GetType() FlowActionEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowActionResponseObjectResource) GetTypeOk() (*FlowActionEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowActionResponseObjectResource) SetType(v FlowActionEnum)`

SetType sets Type field to given value.


### GetId

`func (o *FlowActionResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowActionResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowActionResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *FlowActionResponseObjectResource) GetAttributes() FlowActionResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FlowActionResponseObjectResource) GetAttributesOk() (*FlowActionResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FlowActionResponseObjectResource) SetAttributes(v FlowActionResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *FlowActionResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlowActionResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlowActionResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


