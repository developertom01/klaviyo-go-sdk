# FlowResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowEnum**](FlowEnum.md) |  | 
**Id** | **string** |  | 
**Attributes** | [**FlowResponseObjectResourceAttributes**](FlowResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewFlowResponseObjectResource

`func NewFlowResponseObjectResource(type_ FlowEnum, id string, attributes FlowResponseObjectResourceAttributes, links ObjectLinks, ) *FlowResponseObjectResource`

NewFlowResponseObjectResource instantiates a new FlowResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowResponseObjectResourceWithDefaults

`func NewFlowResponseObjectResourceWithDefaults() *FlowResponseObjectResource`

NewFlowResponseObjectResourceWithDefaults instantiates a new FlowResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlowResponseObjectResource) GetType() FlowEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowResponseObjectResource) GetTypeOk() (*FlowEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowResponseObjectResource) SetType(v FlowEnum)`

SetType sets Type field to given value.


### GetId

`func (o *FlowResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *FlowResponseObjectResource) GetAttributes() FlowResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FlowResponseObjectResource) GetAttributesOk() (*FlowResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FlowResponseObjectResource) SetAttributes(v FlowResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *FlowResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlowResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlowResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


