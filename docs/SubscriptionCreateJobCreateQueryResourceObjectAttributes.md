# SubscriptionCreateJobCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomSource** | Pointer to **NullableString** | A custom method detail or source to store on the consent records. | [optional] 
**Profiles** | [**SubscriptionCreateJobCreateQueryResourceObjectAttributesProfiles**](SubscriptionCreateJobCreateQueryResourceObjectAttributesProfiles.md) |  | 

## Methods

### NewSubscriptionCreateJobCreateQueryResourceObjectAttributes

`func NewSubscriptionCreateJobCreateQueryResourceObjectAttributes(profiles SubscriptionCreateJobCreateQueryResourceObjectAttributesProfiles, ) *SubscriptionCreateJobCreateQueryResourceObjectAttributes`

NewSubscriptionCreateJobCreateQueryResourceObjectAttributes instantiates a new SubscriptionCreateJobCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreateJobCreateQueryResourceObjectAttributesWithDefaults

`func NewSubscriptionCreateJobCreateQueryResourceObjectAttributesWithDefaults() *SubscriptionCreateJobCreateQueryResourceObjectAttributes`

NewSubscriptionCreateJobCreateQueryResourceObjectAttributesWithDefaults instantiates a new SubscriptionCreateJobCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomSource

`func (o *SubscriptionCreateJobCreateQueryResourceObjectAttributes) GetCustomSource() string`

GetCustomSource returns the CustomSource field if non-nil, zero value otherwise.

### GetCustomSourceOk

`func (o *SubscriptionCreateJobCreateQueryResourceObjectAttributes) GetCustomSourceOk() (*string, bool)`

GetCustomSourceOk returns a tuple with the CustomSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSource

`func (o *SubscriptionCreateJobCreateQueryResourceObjectAttributes) SetCustomSource(v string)`

SetCustomSource sets CustomSource field to given value.

### HasCustomSource

`func (o *SubscriptionCreateJobCreateQueryResourceObjectAttributes) HasCustomSource() bool`

HasCustomSource returns a boolean if a field has been set.

### SetCustomSourceNil

`func (o *SubscriptionCreateJobCreateQueryResourceObjectAttributes) SetCustomSourceNil(b bool)`

 SetCustomSourceNil sets the value for CustomSource to be an explicit nil

### UnsetCustomSource
`func (o *SubscriptionCreateJobCreateQueryResourceObjectAttributes) UnsetCustomSource()`

UnsetCustomSource ensures that no value is present for CustomSource, not even an explicit nil
### GetProfiles

`func (o *SubscriptionCreateJobCreateQueryResourceObjectAttributes) GetProfiles() SubscriptionCreateJobCreateQueryResourceObjectAttributesProfiles`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SubscriptionCreateJobCreateQueryResourceObjectAttributes) GetProfilesOk() (*SubscriptionCreateJobCreateQueryResourceObjectAttributesProfiles, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SubscriptionCreateJobCreateQueryResourceObjectAttributes) SetProfiles(v SubscriptionCreateJobCreateQueryResourceObjectAttributesProfiles)`

SetProfiles sets Profiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


