# GetTemplateResponseCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TemplateResponseObjectResource**](TemplateResponseObjectResource.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewGetTemplateResponseCollection

`func NewGetTemplateResponseCollection(data []TemplateResponseObjectResource, links CollectionLinks, ) *GetTemplateResponseCollection`

NewGetTemplateResponseCollection instantiates a new GetTemplateResponseCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTemplateResponseCollectionWithDefaults

`func NewGetTemplateResponseCollectionWithDefaults() *GetTemplateResponseCollection`

NewGetTemplateResponseCollectionWithDefaults instantiates a new GetTemplateResponseCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTemplateResponseCollection) GetData() []TemplateResponseObjectResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTemplateResponseCollection) GetDataOk() (*[]TemplateResponseObjectResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTemplateResponseCollection) SetData(v []TemplateResponseObjectResource)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetTemplateResponseCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTemplateResponseCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTemplateResponseCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


