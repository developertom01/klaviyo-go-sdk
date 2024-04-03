# EventResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EventEnum**](EventEnum.md) |  | 
**Id** | **string** | The Event ID | 
**Attributes** | [**EventResponseObjectResourceAttributes**](EventResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewEventResponseObjectResource

`func NewEventResponseObjectResource(type_ EventEnum, id string, attributes EventResponseObjectResourceAttributes, links ObjectLinks, ) *EventResponseObjectResource`

NewEventResponseObjectResource instantiates a new EventResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseObjectResourceWithDefaults

`func NewEventResponseObjectResourceWithDefaults() *EventResponseObjectResource`

NewEventResponseObjectResourceWithDefaults instantiates a new EventResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventResponseObjectResource) GetType() EventEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventResponseObjectResource) GetTypeOk() (*EventEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventResponseObjectResource) SetType(v EventEnum)`

SetType sets Type field to given value.


### GetId

`func (o *EventResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *EventResponseObjectResource) GetAttributes() EventResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EventResponseObjectResource) GetAttributesOk() (*EventResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EventResponseObjectResource) SetAttributes(v EventResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *EventResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


