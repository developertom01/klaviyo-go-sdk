# FlowMessageResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowMessageEnum**](FlowMessageEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**FlowMessageResponseObjectResourceAttributes**](FlowMessageResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewFlowMessageResponseObjectResource

`func NewFlowMessageResponseObjectResource(type_ FlowMessageEnum, id string, attributes FlowMessageResponseObjectResourceAttributes, links ObjectLinks, ) *FlowMessageResponseObjectResource`

NewFlowMessageResponseObjectResource instantiates a new FlowMessageResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowMessageResponseObjectResourceWithDefaults

`func NewFlowMessageResponseObjectResourceWithDefaults() *FlowMessageResponseObjectResource`

NewFlowMessageResponseObjectResourceWithDefaults instantiates a new FlowMessageResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlowMessageResponseObjectResource) GetType() FlowMessageEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowMessageResponseObjectResource) GetTypeOk() (*FlowMessageEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowMessageResponseObjectResource) SetType(v FlowMessageEnum)`

SetType sets Type field to given value.


### GetId

`func (o *FlowMessageResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowMessageResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowMessageResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *FlowMessageResponseObjectResource) GetAttributes() FlowMessageResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FlowMessageResponseObjectResource) GetAttributesOk() (*FlowMessageResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FlowMessageResponseObjectResource) SetAttributes(v FlowMessageResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *FlowMessageResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlowMessageResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlowMessageResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


