# OnsiteSubscriptionCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomSource** | Pointer to **NullableString** | A custom method detail or source to store on the consent records for this subscription. | [optional] 
**Profile** | [**OnsiteSubscriptionCreateQueryResourceObjectAttributesProfile**](OnsiteSubscriptionCreateQueryResourceObjectAttributesProfile.md) |  | 

## Methods

### NewOnsiteSubscriptionCreateQueryResourceObjectAttributes

`func NewOnsiteSubscriptionCreateQueryResourceObjectAttributes(profile OnsiteSubscriptionCreateQueryResourceObjectAttributesProfile, ) *OnsiteSubscriptionCreateQueryResourceObjectAttributes`

NewOnsiteSubscriptionCreateQueryResourceObjectAttributes instantiates a new OnsiteSubscriptionCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnsiteSubscriptionCreateQueryResourceObjectAttributesWithDefaults

`func NewOnsiteSubscriptionCreateQueryResourceObjectAttributesWithDefaults() *OnsiteSubscriptionCreateQueryResourceObjectAttributes`

NewOnsiteSubscriptionCreateQueryResourceObjectAttributesWithDefaults instantiates a new OnsiteSubscriptionCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomSource

`func (o *OnsiteSubscriptionCreateQueryResourceObjectAttributes) GetCustomSource() string`

GetCustomSource returns the CustomSource field if non-nil, zero value otherwise.

### GetCustomSourceOk

`func (o *OnsiteSubscriptionCreateQueryResourceObjectAttributes) GetCustomSourceOk() (*string, bool)`

GetCustomSourceOk returns a tuple with the CustomSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSource

`func (o *OnsiteSubscriptionCreateQueryResourceObjectAttributes) SetCustomSource(v string)`

SetCustomSource sets CustomSource field to given value.

### HasCustomSource

`func (o *OnsiteSubscriptionCreateQueryResourceObjectAttributes) HasCustomSource() bool`

HasCustomSource returns a boolean if a field has been set.

### SetCustomSourceNil

`func (o *OnsiteSubscriptionCreateQueryResourceObjectAttributes) SetCustomSourceNil(b bool)`

 SetCustomSourceNil sets the value for CustomSource to be an explicit nil

### UnsetCustomSource
`func (o *OnsiteSubscriptionCreateQueryResourceObjectAttributes) UnsetCustomSource()`

UnsetCustomSource ensures that no value is present for CustomSource, not even an explicit nil
### GetProfile

`func (o *OnsiteSubscriptionCreateQueryResourceObjectAttributes) GetProfile() OnsiteSubscriptionCreateQueryResourceObjectAttributesProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OnsiteSubscriptionCreateQueryResourceObjectAttributes) GetProfileOk() (*OnsiteSubscriptionCreateQueryResourceObjectAttributesProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OnsiteSubscriptionCreateQueryResourceObjectAttributes) SetProfile(v OnsiteSubscriptionCreateQueryResourceObjectAttributesProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


