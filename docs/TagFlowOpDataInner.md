# TagFlowOpDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowEnum**](FlowEnum.md) |  | 
**Id** | **string** | The IDs of the flows to link or unlink with the given Tag ID | 

## Methods

### NewTagFlowOpDataInner

`func NewTagFlowOpDataInner(type_ FlowEnum, id string, ) *TagFlowOpDataInner`

NewTagFlowOpDataInner instantiates a new TagFlowOpDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagFlowOpDataInnerWithDefaults

`func NewTagFlowOpDataInnerWithDefaults() *TagFlowOpDataInner`

NewTagFlowOpDataInnerWithDefaults instantiates a new TagFlowOpDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagFlowOpDataInner) GetType() FlowEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagFlowOpDataInner) GetTypeOk() (*FlowEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagFlowOpDataInner) SetType(v FlowEnum)`

SetType sets Type field to given value.


### GetId

`func (o *TagFlowOpDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagFlowOpDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagFlowOpDataInner) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


