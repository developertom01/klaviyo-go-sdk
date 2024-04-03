# ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogVariantEnum**](CatalogVariantEnum.md) |  | 
**Id** | **string** | The catalog variant ID for which the profile is subscribing to back in stock notifications. This ID is made up of the integration type, catalog ID, and and the external ID of the variant like so: &#x60;integrationType:::catalogId:::externalId&#x60;. If the integration you are using is not set up for multi-catalog storage, the &#39;catalogId&#39; will be &#x60;$default&#x60;. For Shopify &#x60;$shopify:::$default:::33001893429341&#x60; | 

## Methods

### NewServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData

`func NewServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData(type_ CatalogVariantEnum, id string, ) *ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData`

NewServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData instantiates a new ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantDataWithDefaults

`func NewServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantDataWithDefaults() *ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData`

NewServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantDataWithDefaults instantiates a new ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData) GetType() CatalogVariantEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData) GetTypeOk() (*CatalogVariantEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData) SetType(v CatalogVariantEnum)`

SetType sets Type field to given value.


### GetId

`func (o *ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


