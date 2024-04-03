# ClientBISSubscriptionCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **[]string** | The channel(s) through which the profile would like to receive the back in stock notification. This can be leveraged within a back in stock flow to notify the subscriber through their preferred channel(s). | 
**Profile** | [**ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile**](ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile.md) |  | 

## Methods

### NewClientBISSubscriptionCreateQueryResourceObjectAttributes

`func NewClientBISSubscriptionCreateQueryResourceObjectAttributes(channels []string, profile ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile, ) *ClientBISSubscriptionCreateQueryResourceObjectAttributes`

NewClientBISSubscriptionCreateQueryResourceObjectAttributes instantiates a new ClientBISSubscriptionCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientBISSubscriptionCreateQueryResourceObjectAttributesWithDefaults

`func NewClientBISSubscriptionCreateQueryResourceObjectAttributesWithDefaults() *ClientBISSubscriptionCreateQueryResourceObjectAttributes`

NewClientBISSubscriptionCreateQueryResourceObjectAttributesWithDefaults instantiates a new ClientBISSubscriptionCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ClientBISSubscriptionCreateQueryResourceObjectAttributes) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ClientBISSubscriptionCreateQueryResourceObjectAttributes) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ClientBISSubscriptionCreateQueryResourceObjectAttributes) SetChannels(v []string)`

SetChannels sets Channels field to given value.


### GetProfile

`func (o *ClientBISSubscriptionCreateQueryResourceObjectAttributes) GetProfile() ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ClientBISSubscriptionCreateQueryResourceObjectAttributes) GetProfileOk() (*ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ClientBISSubscriptionCreateQueryResourceObjectAttributes) SetProfile(v ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


