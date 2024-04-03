# AttributionResponseObjectResourceRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**AttributionResponseObjectResourceRelationshipsEvent**](AttributionResponseObjectResourceRelationshipsEvent.md) |  | [optional] 
**AttributedEvent** | Pointer to [**AttributionResponseObjectResourceRelationshipsAttributedEvent**](AttributionResponseObjectResourceRelationshipsAttributedEvent.md) |  | [optional] 
**Campaign** | Pointer to [**AttributionResponseObjectResourceRelationshipsCampaign**](AttributionResponseObjectResourceRelationshipsCampaign.md) |  | [optional] 
**CampaignMessage** | Pointer to [**AttributionResponseObjectResourceRelationshipsCampaignMessage**](AttributionResponseObjectResourceRelationshipsCampaignMessage.md) |  | [optional] 
**Flow** | Pointer to [**AttributionResponseObjectResourceRelationshipsFlow**](AttributionResponseObjectResourceRelationshipsFlow.md) |  | [optional] 
**FlowMessage** | Pointer to [**AttributionResponseObjectResourceRelationshipsFlowMessage**](AttributionResponseObjectResourceRelationshipsFlowMessage.md) |  | [optional] 
**FlowMessageVariation** | Pointer to [**AttributionResponseObjectResourceRelationshipsFlowMessageVariation**](AttributionResponseObjectResourceRelationshipsFlowMessageVariation.md) |  | [optional] 

## Methods

### NewAttributionResponseObjectResourceRelationships

`func NewAttributionResponseObjectResourceRelationships() *AttributionResponseObjectResourceRelationships`

NewAttributionResponseObjectResourceRelationships instantiates a new AttributionResponseObjectResourceRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributionResponseObjectResourceRelationshipsWithDefaults

`func NewAttributionResponseObjectResourceRelationshipsWithDefaults() *AttributionResponseObjectResourceRelationships`

NewAttributionResponseObjectResourceRelationshipsWithDefaults instantiates a new AttributionResponseObjectResourceRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AttributionResponseObjectResourceRelationships) GetEvent() AttributionResponseObjectResourceRelationshipsEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AttributionResponseObjectResourceRelationships) GetEventOk() (*AttributionResponseObjectResourceRelationshipsEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AttributionResponseObjectResourceRelationships) SetEvent(v AttributionResponseObjectResourceRelationshipsEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *AttributionResponseObjectResourceRelationships) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetAttributedEvent

`func (o *AttributionResponseObjectResourceRelationships) GetAttributedEvent() AttributionResponseObjectResourceRelationshipsAttributedEvent`

GetAttributedEvent returns the AttributedEvent field if non-nil, zero value otherwise.

### GetAttributedEventOk

`func (o *AttributionResponseObjectResourceRelationships) GetAttributedEventOk() (*AttributionResponseObjectResourceRelationshipsAttributedEvent, bool)`

GetAttributedEventOk returns a tuple with the AttributedEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributedEvent

`func (o *AttributionResponseObjectResourceRelationships) SetAttributedEvent(v AttributionResponseObjectResourceRelationshipsAttributedEvent)`

SetAttributedEvent sets AttributedEvent field to given value.

### HasAttributedEvent

`func (o *AttributionResponseObjectResourceRelationships) HasAttributedEvent() bool`

HasAttributedEvent returns a boolean if a field has been set.

### GetCampaign

`func (o *AttributionResponseObjectResourceRelationships) GetCampaign() AttributionResponseObjectResourceRelationshipsCampaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *AttributionResponseObjectResourceRelationships) GetCampaignOk() (*AttributionResponseObjectResourceRelationshipsCampaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *AttributionResponseObjectResourceRelationships) SetCampaign(v AttributionResponseObjectResourceRelationshipsCampaign)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *AttributionResponseObjectResourceRelationships) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetCampaignMessage

`func (o *AttributionResponseObjectResourceRelationships) GetCampaignMessage() AttributionResponseObjectResourceRelationshipsCampaignMessage`

GetCampaignMessage returns the CampaignMessage field if non-nil, zero value otherwise.

### GetCampaignMessageOk

`func (o *AttributionResponseObjectResourceRelationships) GetCampaignMessageOk() (*AttributionResponseObjectResourceRelationshipsCampaignMessage, bool)`

GetCampaignMessageOk returns a tuple with the CampaignMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignMessage

`func (o *AttributionResponseObjectResourceRelationships) SetCampaignMessage(v AttributionResponseObjectResourceRelationshipsCampaignMessage)`

SetCampaignMessage sets CampaignMessage field to given value.

### HasCampaignMessage

`func (o *AttributionResponseObjectResourceRelationships) HasCampaignMessage() bool`

HasCampaignMessage returns a boolean if a field has been set.

### GetFlow

`func (o *AttributionResponseObjectResourceRelationships) GetFlow() AttributionResponseObjectResourceRelationshipsFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *AttributionResponseObjectResourceRelationships) GetFlowOk() (*AttributionResponseObjectResourceRelationshipsFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *AttributionResponseObjectResourceRelationships) SetFlow(v AttributionResponseObjectResourceRelationshipsFlow)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *AttributionResponseObjectResourceRelationships) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetFlowMessage

`func (o *AttributionResponseObjectResourceRelationships) GetFlowMessage() AttributionResponseObjectResourceRelationshipsFlowMessage`

GetFlowMessage returns the FlowMessage field if non-nil, zero value otherwise.

### GetFlowMessageOk

`func (o *AttributionResponseObjectResourceRelationships) GetFlowMessageOk() (*AttributionResponseObjectResourceRelationshipsFlowMessage, bool)`

GetFlowMessageOk returns a tuple with the FlowMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowMessage

`func (o *AttributionResponseObjectResourceRelationships) SetFlowMessage(v AttributionResponseObjectResourceRelationshipsFlowMessage)`

SetFlowMessage sets FlowMessage field to given value.

### HasFlowMessage

`func (o *AttributionResponseObjectResourceRelationships) HasFlowMessage() bool`

HasFlowMessage returns a boolean if a field has been set.

### GetFlowMessageVariation

`func (o *AttributionResponseObjectResourceRelationships) GetFlowMessageVariation() AttributionResponseObjectResourceRelationshipsFlowMessageVariation`

GetFlowMessageVariation returns the FlowMessageVariation field if non-nil, zero value otherwise.

### GetFlowMessageVariationOk

`func (o *AttributionResponseObjectResourceRelationships) GetFlowMessageVariationOk() (*AttributionResponseObjectResourceRelationshipsFlowMessageVariation, bool)`

GetFlowMessageVariationOk returns a tuple with the FlowMessageVariation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowMessageVariation

`func (o *AttributionResponseObjectResourceRelationships) SetFlowMessageVariation(v AttributionResponseObjectResourceRelationshipsFlowMessageVariation)`

SetFlowMessageVariation sets FlowMessageVariation field to given value.

### HasFlowMessageVariation

`func (o *AttributionResponseObjectResourceRelationships) HasFlowMessageVariation() bool`

HasFlowMessageVariation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


