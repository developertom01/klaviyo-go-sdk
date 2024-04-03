# ClientBISSubscriptionCreateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**BackInStockSubscriptionEnum**](BackInStockSubscriptionEnum.md) |  | 
**Attributes** | [**ClientBISSubscriptionCreateQueryResourceObjectAttributes**](ClientBISSubscriptionCreateQueryResourceObjectAttributes.md) |  | 
**Relationships** | [**ServerBISSubscriptionCreateQueryResourceObjectRelationships**](ServerBISSubscriptionCreateQueryResourceObjectRelationships.md) |  | 

## Methods

### NewClientBISSubscriptionCreateQueryResourceObject

`func NewClientBISSubscriptionCreateQueryResourceObject(type_ BackInStockSubscriptionEnum, attributes ClientBISSubscriptionCreateQueryResourceObjectAttributes, relationships ServerBISSubscriptionCreateQueryResourceObjectRelationships, ) *ClientBISSubscriptionCreateQueryResourceObject`

NewClientBISSubscriptionCreateQueryResourceObject instantiates a new ClientBISSubscriptionCreateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientBISSubscriptionCreateQueryResourceObjectWithDefaults

`func NewClientBISSubscriptionCreateQueryResourceObjectWithDefaults() *ClientBISSubscriptionCreateQueryResourceObject`

NewClientBISSubscriptionCreateQueryResourceObjectWithDefaults instantiates a new ClientBISSubscriptionCreateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClientBISSubscriptionCreateQueryResourceObject) GetType() BackInStockSubscriptionEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientBISSubscriptionCreateQueryResourceObject) GetTypeOk() (*BackInStockSubscriptionEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientBISSubscriptionCreateQueryResourceObject) SetType(v BackInStockSubscriptionEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ClientBISSubscriptionCreateQueryResourceObject) GetAttributes() ClientBISSubscriptionCreateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ClientBISSubscriptionCreateQueryResourceObject) GetAttributesOk() (*ClientBISSubscriptionCreateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ClientBISSubscriptionCreateQueryResourceObject) SetAttributes(v ClientBISSubscriptionCreateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ClientBISSubscriptionCreateQueryResourceObject) GetRelationships() ServerBISSubscriptionCreateQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ClientBISSubscriptionCreateQueryResourceObject) GetRelationshipsOk() (*ServerBISSubscriptionCreateQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ClientBISSubscriptionCreateQueryResourceObject) SetRelationships(v ServerBISSubscriptionCreateQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


