# ServerBISSubscriptionCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **[]string** | The channel(s) through which the profile would like to receive the back in stock notification. This can be leveraged within a back in stock flow to notify the subscriber through their preferred channel(s). | 
**Profile** | Pointer to [**NullableServerBISSubscriptionCreateQueryResourceObjectAttributesProfile**](ServerBISSubscriptionCreateQueryResourceObjectAttributesProfile.md) |  | [optional] 

## Methods

### NewServerBISSubscriptionCreateQueryResourceObjectAttributes

`func NewServerBISSubscriptionCreateQueryResourceObjectAttributes(channels []string, ) *ServerBISSubscriptionCreateQueryResourceObjectAttributes`

NewServerBISSubscriptionCreateQueryResourceObjectAttributes instantiates a new ServerBISSubscriptionCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerBISSubscriptionCreateQueryResourceObjectAttributesWithDefaults

`func NewServerBISSubscriptionCreateQueryResourceObjectAttributesWithDefaults() *ServerBISSubscriptionCreateQueryResourceObjectAttributes`

NewServerBISSubscriptionCreateQueryResourceObjectAttributesWithDefaults instantiates a new ServerBISSubscriptionCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ServerBISSubscriptionCreateQueryResourceObjectAttributes) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ServerBISSubscriptionCreateQueryResourceObjectAttributes) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ServerBISSubscriptionCreateQueryResourceObjectAttributes) SetChannels(v []string)`

SetChannels sets Channels field to given value.


### GetProfile

`func (o *ServerBISSubscriptionCreateQueryResourceObjectAttributes) GetProfile() ServerBISSubscriptionCreateQueryResourceObjectAttributesProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ServerBISSubscriptionCreateQueryResourceObjectAttributes) GetProfileOk() (*ServerBISSubscriptionCreateQueryResourceObjectAttributesProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ServerBISSubscriptionCreateQueryResourceObjectAttributes) SetProfile(v ServerBISSubscriptionCreateQueryResourceObjectAttributesProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ServerBISSubscriptionCreateQueryResourceObjectAttributes) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *ServerBISSubscriptionCreateQueryResourceObjectAttributes) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *ServerBISSubscriptionCreateQueryResourceObjectAttributes) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


