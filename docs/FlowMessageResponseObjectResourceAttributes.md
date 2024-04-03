# FlowMessageResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Channel** | **string** |  | 
**Content** | [**FlowMessageResponseObjectResourceAttributesContent**](FlowMessageResponseObjectResourceAttributesContent.md) |  | 
**Created** | Pointer to **NullableTime** |  | [optional] 
**Updated** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewFlowMessageResponseObjectResourceAttributes

`func NewFlowMessageResponseObjectResourceAttributes(name string, channel string, content FlowMessageResponseObjectResourceAttributesContent, ) *FlowMessageResponseObjectResourceAttributes`

NewFlowMessageResponseObjectResourceAttributes instantiates a new FlowMessageResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowMessageResponseObjectResourceAttributesWithDefaults

`func NewFlowMessageResponseObjectResourceAttributesWithDefaults() *FlowMessageResponseObjectResourceAttributes`

NewFlowMessageResponseObjectResourceAttributesWithDefaults instantiates a new FlowMessageResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlowMessageResponseObjectResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowMessageResponseObjectResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowMessageResponseObjectResourceAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetChannel

`func (o *FlowMessageResponseObjectResourceAttributes) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *FlowMessageResponseObjectResourceAttributes) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *FlowMessageResponseObjectResourceAttributes) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetContent

`func (o *FlowMessageResponseObjectResourceAttributes) GetContent() FlowMessageResponseObjectResourceAttributesContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FlowMessageResponseObjectResourceAttributes) GetContentOk() (*FlowMessageResponseObjectResourceAttributesContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FlowMessageResponseObjectResourceAttributes) SetContent(v FlowMessageResponseObjectResourceAttributesContent)`

SetContent sets Content field to given value.


### GetCreated

`func (o *FlowMessageResponseObjectResourceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowMessageResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowMessageResponseObjectResourceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowMessageResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *FlowMessageResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *FlowMessageResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *FlowMessageResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FlowMessageResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FlowMessageResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FlowMessageResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *FlowMessageResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *FlowMessageResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


