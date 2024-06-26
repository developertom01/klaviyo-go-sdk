# Go API client for klaviyo

The Klaviyo REST API. Please visit <https://developers.klaviyo.com> for more details.

## Overview

- API version: 2024-02-15
- Package version: 0.0.1
For more information, please visit [https://developers.klaviyo.com](https://developers.klaviyo.com)

## Installation

Install the following dependencies:

```sh
 go get github.com/developertom01/klaviyo-go-sdk
```

Put the package under your project folder and add the following in import:

```go
import klaviyo "github.com/developertom01/klaviyo-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `klaviyo.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), klaviyo.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `klaviyo.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), klaviyo.ContextServerVariables, map[string]string{
 "basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `klaviyo.ContextOperationServerIndices` and `klaviyo.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), klaviyo.ContextOperationServerIndices, map[string]int{
 "{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), klaviyo.ContextOperationServerVariables, map[string]map[string]string{
 "{classname}Service.{nickname}": {
  "port": "8443",
 },
})
```

## Documentation for API Endpoints

All URIs are relative to *<https://a.klaviyo.com>*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsAPI* | [**GetAccount**](docs/AccountsAPI.md#getaccount) | **Get** /api/accounts/{id}/ | Get Account
*AccountsAPI* | [**GetAccounts**](docs/AccountsAPI.md#getaccounts) | **Get** /api/accounts/ | Get Accounts
*CampaignsAPI* | [**CreateCampaign**](docs/CampaignsAPI.md#createcampaign) | **Post** /api/campaigns/ | Create Campaign
*CampaignsAPI* | [**CreateCampaignClone**](docs/CampaignsAPI.md#createcampaignclone) | **Post** /api/campaign-clone/ | Create Campaign Clone
*CampaignsAPI* | [**CreateCampaignMessageAssignTemplate**](docs/CampaignsAPI.md#createcampaignmessageassigntemplate) | **Post** /api/campaign-message-assign-template/ | Assign Campaign Message Template
*CampaignsAPI* | [**CreateCampaignRecipientEstimationJob**](docs/CampaignsAPI.md#createcampaignrecipientestimationjob) | **Post** /api/campaign-recipient-estimation-jobs/ | Create Campaign Recipient Estimation Job
*CampaignsAPI* | [**CreateCampaignSendJob**](docs/CampaignsAPI.md#createcampaignsendjob) | **Post** /api/campaign-send-jobs/ | Create Campaign Send Job
*CampaignsAPI* | [**DeleteCampaign**](docs/CampaignsAPI.md#deletecampaign) | **Delete** /api/campaigns/{id}/ | Delete Campaign
*CampaignsAPI* | [**GetCampaign**](docs/CampaignsAPI.md#getcampaign) | **Get** /api/campaigns/{id}/ | Get Campaign
*CampaignsAPI* | [**GetCampaignCampaignMessages**](docs/CampaignsAPI.md#getcampaigncampaignmessages) | **Get** /api/campaigns/{id}/campaign-messages/ | Get Campaign Campaign Messages
*CampaignsAPI* | [**GetCampaignMessage**](docs/CampaignsAPI.md#getcampaignmessage) | **Get** /api/campaign-messages/{id}/ | Get Campaign Message
*CampaignsAPI* | [**GetCampaignMessageCampaign**](docs/CampaignsAPI.md#getcampaignmessagecampaign) | **Get** /api/campaign-messages/{id}/campaign/ | Get Campaign Message Campaign
*CampaignsAPI* | [**GetCampaignMessageRelationshipsCampaign**](docs/CampaignsAPI.md#getcampaignmessagerelationshipscampaign) | **Get** /api/campaign-messages/{id}/relationships/campaign/ | Get Campaign Message Relationships Campaign
*CampaignsAPI* | [**GetCampaignMessageRelationshipsTemplate**](docs/CampaignsAPI.md#getcampaignmessagerelationshipstemplate) | **Get** /api/campaign-messages/{id}/relationships/template/ | Get Campaign Message Relationships Template
*CampaignsAPI* | [**GetCampaignMessageTemplate**](docs/CampaignsAPI.md#getcampaignmessagetemplate) | **Get** /api/campaign-messages/{id}/template/ | Get Campaign Message Template
*CampaignsAPI* | [**GetCampaignRecipientEstimation**](docs/CampaignsAPI.md#getcampaignrecipientestimation) | **Get** /api/campaign-recipient-estimations/{id}/ | Get Campaign Recipient Estimation
*CampaignsAPI* | [**GetCampaignRecipientEstimationJob**](docs/CampaignsAPI.md#getcampaignrecipientestimationjob) | **Get** /api/campaign-recipient-estimation-jobs/{id}/ | Get Campaign Recipient Estimation Job
*CampaignsAPI* | [**GetCampaignRelationshipsCampaignMessages**](docs/CampaignsAPI.md#getcampaignrelationshipscampaignmessages) | **Get** /api/campaigns/{id}/relationships/campaign-messages/ | Get Campaign Relationships Campaign Messages
*CampaignsAPI* | [**GetCampaignRelationshipsTags**](docs/CampaignsAPI.md#getcampaignrelationshipstags) | **Get** /api/campaigns/{id}/relationships/tags/ | Get Campaign Relationships Tags
*CampaignsAPI* | [**GetCampaignSendJob**](docs/CampaignsAPI.md#getcampaignsendjob) | **Get** /api/campaign-send-jobs/{id}/ | Get Campaign Send Job
*CampaignsAPI* | [**GetCampaignTags**](docs/CampaignsAPI.md#getcampaigntags) | **Get** /api/campaigns/{id}/tags/ | Get Campaign Tags
*CampaignsAPI* | [**GetCampaigns**](docs/CampaignsAPI.md#getcampaigns) | **Get** /api/campaigns/ | Get Campaigns
*CampaignsAPI* | [**UpdateCampaign**](docs/CampaignsAPI.md#updatecampaign) | **Patch** /api/campaigns/{id}/ | Update Campaign
*CampaignsAPI* | [**UpdateCampaignMessage**](docs/CampaignsAPI.md#updatecampaignmessage) | **Patch** /api/campaign-messages/{id}/ | Update Campaign Message
*CampaignsAPI* | [**UpdateCampaignSendJob**](docs/CampaignsAPI.md#updatecampaignsendjob) | **Patch** /api/campaign-send-jobs/{id}/ | Update Campaign Send Job
*CatalogsAPI* | [**CreateBackInStockSubscription**](docs/CatalogsAPI.md#createbackinstocksubscription) | **Post** /api/back-in-stock-subscriptions/ | Create Back In Stock Subscription
*CatalogsAPI* | [**CreateCatalogCategory**](docs/CatalogsAPI.md#createcatalogcategory) | **Post** /api/catalog-categories/ | Create Catalog Category
*CatalogsAPI* | [**CreateCatalogCategoryRelationshipsItems**](docs/CatalogsAPI.md#createcatalogcategoryrelationshipsitems) | **Post** /api/catalog-categories/{id}/relationships/items/ | Create Catalog Category Relationships Items
*CatalogsAPI* | [**CreateCatalogItem**](docs/CatalogsAPI.md#createcatalogitem) | **Post** /api/catalog-items/ | Create Catalog Item
*CatalogsAPI* | [**CreateCatalogItemRelationshipsCategories**](docs/CatalogsAPI.md#createcatalogitemrelationshipscategories) | **Post** /api/catalog-items/{id}/relationships/categories/ | Create Catalog Item Relationships Categories
*CatalogsAPI* | [**CreateCatalogVariant**](docs/CatalogsAPI.md#createcatalogvariant) | **Post** /api/catalog-variants/ | Create Catalog Variant
*CatalogsAPI* | [**DeleteCatalogCategory**](docs/CatalogsAPI.md#deletecatalogcategory) | **Delete** /api/catalog-categories/{id}/ | Delete Catalog Category
*CatalogsAPI* | [**DeleteCatalogCategoryRelationshipsItems**](docs/CatalogsAPI.md#deletecatalogcategoryrelationshipsitems) | **Delete** /api/catalog-categories/{id}/relationships/items/ | Delete Catalog Category Relationships Items
*CatalogsAPI* | [**DeleteCatalogItem**](docs/CatalogsAPI.md#deletecatalogitem) | **Delete** /api/catalog-items/{id}/ | Delete Catalog Item
*CatalogsAPI* | [**DeleteCatalogItemRelationshipsCategories**](docs/CatalogsAPI.md#deletecatalogitemrelationshipscategories) | **Delete** /api/catalog-items/{id}/relationships/categories/ | Delete Catalog Item Relationships Categories
*CatalogsAPI* | [**DeleteCatalogVariant**](docs/CatalogsAPI.md#deletecatalogvariant) | **Delete** /api/catalog-variants/{id}/ | Delete Catalog Variant
*CatalogsAPI* | [**GetCatalogCategories**](docs/CatalogsAPI.md#getcatalogcategories) | **Get** /api/catalog-categories/ | Get Catalog Categories
*CatalogsAPI* | [**GetCatalogCategory**](docs/CatalogsAPI.md#getcatalogcategory) | **Get** /api/catalog-categories/{id}/ | Get Catalog Category
*CatalogsAPI* | [**GetCatalogCategoryItems**](docs/CatalogsAPI.md#getcatalogcategoryitems) | **Get** /api/catalog-categories/{id}/items/ | Get Catalog Category Items
*CatalogsAPI* | [**GetCatalogCategoryRelationshipsItems**](docs/CatalogsAPI.md#getcatalogcategoryrelationshipsitems) | **Get** /api/catalog-categories/{id}/relationships/items/ | Get Catalog Category Relationships Items
*CatalogsAPI* | [**GetCatalogItem**](docs/CatalogsAPI.md#getcatalogitem) | **Get** /api/catalog-items/{id}/ | Get Catalog Item
*CatalogsAPI* | [**GetCatalogItemCategories**](docs/CatalogsAPI.md#getcatalogitemcategories) | **Get** /api/catalog-items/{id}/categories/ | Get Catalog Item Categories
*CatalogsAPI* | [**GetCatalogItemRelationshipsCategories**](docs/CatalogsAPI.md#getcatalogitemrelationshipscategories) | **Get** /api/catalog-items/{id}/relationships/categories/ | Get Catalog Item Relationships Categories
*CatalogsAPI* | [**GetCatalogItemVariants**](docs/CatalogsAPI.md#getcatalogitemvariants) | **Get** /api/catalog-items/{id}/variants/ | Get Catalog Item Variants
*CatalogsAPI* | [**GetCatalogItems**](docs/CatalogsAPI.md#getcatalogitems) | **Get** /api/catalog-items/ | Get Catalog Items
*CatalogsAPI* | [**GetCatalogVariant**](docs/CatalogsAPI.md#getcatalogvariant) | **Get** /api/catalog-variants/{id}/ | Get Catalog Variant
*CatalogsAPI* | [**GetCatalogVariants**](docs/CatalogsAPI.md#getcatalogvariants) | **Get** /api/catalog-variants/ | Get Catalog Variants
*CatalogsAPI* | [**GetCreateCategoriesJob**](docs/CatalogsAPI.md#getcreatecategoriesjob) | **Get** /api/catalog-category-bulk-create-jobs/{job_id}/ | Get Create Categories Job
*CatalogsAPI* | [**GetCreateCategoriesJobs**](docs/CatalogsAPI.md#getcreatecategoriesjobs) | **Get** /api/catalog-category-bulk-create-jobs/ | Get Create Categories Jobs
*CatalogsAPI* | [**GetCreateItemsJob**](docs/CatalogsAPI.md#getcreateitemsjob) | **Get** /api/catalog-item-bulk-create-jobs/{job_id}/ | Get Create Items Job
*CatalogsAPI* | [**GetCreateItemsJobs**](docs/CatalogsAPI.md#getcreateitemsjobs) | **Get** /api/catalog-item-bulk-create-jobs/ | Get Create Items Jobs
*CatalogsAPI* | [**GetCreateVariantsJob**](docs/CatalogsAPI.md#getcreatevariantsjob) | **Get** /api/catalog-variant-bulk-create-jobs/{job_id}/ | Get Create Variants Job
*CatalogsAPI* | [**GetCreateVariantsJobs**](docs/CatalogsAPI.md#getcreatevariantsjobs) | **Get** /api/catalog-variant-bulk-create-jobs/ | Get Create Variants Jobs
*CatalogsAPI* | [**GetDeleteCategoriesJob**](docs/CatalogsAPI.md#getdeletecategoriesjob) | **Get** /api/catalog-category-bulk-delete-jobs/{job_id}/ | Get Delete Categories Job
*CatalogsAPI* | [**GetDeleteCategoriesJobs**](docs/CatalogsAPI.md#getdeletecategoriesjobs) | **Get** /api/catalog-category-bulk-delete-jobs/ | Get Delete Categories Jobs
*CatalogsAPI* | [**GetDeleteItemsJob**](docs/CatalogsAPI.md#getdeleteitemsjob) | **Get** /api/catalog-item-bulk-delete-jobs/{job_id}/ | Get Delete Items Job
*CatalogsAPI* | [**GetDeleteItemsJobs**](docs/CatalogsAPI.md#getdeleteitemsjobs) | **Get** /api/catalog-item-bulk-delete-jobs/ | Get Delete Items Jobs
*CatalogsAPI* | [**GetDeleteVariantsJob**](docs/CatalogsAPI.md#getdeletevariantsjob) | **Get** /api/catalog-variant-bulk-delete-jobs/{job_id}/ | Get Delete Variants Job
*CatalogsAPI* | [**GetDeleteVariantsJobs**](docs/CatalogsAPI.md#getdeletevariantsjobs) | **Get** /api/catalog-variant-bulk-delete-jobs/ | Get Delete Variants Jobs
*CatalogsAPI* | [**GetUpdateCategoriesJob**](docs/CatalogsAPI.md#getupdatecategoriesjob) | **Get** /api/catalog-category-bulk-update-jobs/{job_id}/ | Get Update Categories Job
*CatalogsAPI* | [**GetUpdateCategoriesJobs**](docs/CatalogsAPI.md#getupdatecategoriesjobs) | **Get** /api/catalog-category-bulk-update-jobs/ | Get Update Categories Jobs
*CatalogsAPI* | [**GetUpdateItemsJob**](docs/CatalogsAPI.md#getupdateitemsjob) | **Get** /api/catalog-item-bulk-update-jobs/{job_id}/ | Get Update Items Job
*CatalogsAPI* | [**GetUpdateItemsJobs**](docs/CatalogsAPI.md#getupdateitemsjobs) | **Get** /api/catalog-item-bulk-update-jobs/ | Get Update Items Jobs
*CatalogsAPI* | [**GetUpdateVariantsJob**](docs/CatalogsAPI.md#getupdatevariantsjob) | **Get** /api/catalog-variant-bulk-update-jobs/{job_id}/ | Get Update Variants Job
*CatalogsAPI* | [**GetUpdateVariantsJobs**](docs/CatalogsAPI.md#getupdatevariantsjobs) | **Get** /api/catalog-variant-bulk-update-jobs/ | Get Update Variants Jobs
*CatalogsAPI* | [**SpawnCreateCategoriesJob**](docs/CatalogsAPI.md#spawncreatecategoriesjob) | **Post** /api/catalog-category-bulk-create-jobs/ | Spawn Create Categories Job
*CatalogsAPI* | [**SpawnCreateItemsJob**](docs/CatalogsAPI.md#spawncreateitemsjob) | **Post** /api/catalog-item-bulk-create-jobs/ | Spawn Create Items Job
*CatalogsAPI* | [**SpawnCreateVariantsJob**](docs/CatalogsAPI.md#spawncreatevariantsjob) | **Post** /api/catalog-variant-bulk-create-jobs/ | Spawn Create Variants Job
*CatalogsAPI* | [**SpawnDeleteCategoriesJob**](docs/CatalogsAPI.md#spawndeletecategoriesjob) | **Post** /api/catalog-category-bulk-delete-jobs/ | Spawn Delete Categories Job
*CatalogsAPI* | [**SpawnDeleteItemsJob**](docs/CatalogsAPI.md#spawndeleteitemsjob) | **Post** /api/catalog-item-bulk-delete-jobs/ | Spawn Delete Items Job
*CatalogsAPI* | [**SpawnDeleteVariantsJob**](docs/CatalogsAPI.md#spawndeletevariantsjob) | **Post** /api/catalog-variant-bulk-delete-jobs/ | Spawn Delete Variants Job
*CatalogsAPI* | [**SpawnUpdateCategoriesJob**](docs/CatalogsAPI.md#spawnupdatecategoriesjob) | **Post** /api/catalog-category-bulk-update-jobs/ | Spawn Update Categories Job
*CatalogsAPI* | [**SpawnUpdateItemsJob**](docs/CatalogsAPI.md#spawnupdateitemsjob) | **Post** /api/catalog-item-bulk-update-jobs/ | Spawn Update Items Job
*CatalogsAPI* | [**SpawnUpdateVariantsJob**](docs/CatalogsAPI.md#spawnupdatevariantsjob) | **Post** /api/catalog-variant-bulk-update-jobs/ | Spawn Update Variants Job
*CatalogsAPI* | [**UpdateCatalogCategory**](docs/CatalogsAPI.md#updatecatalogcategory) | **Patch** /api/catalog-categories/{id}/ | Update Catalog Category
*CatalogsAPI* | [**UpdateCatalogCategoryRelationshipsItems**](docs/CatalogsAPI.md#updatecatalogcategoryrelationshipsitems) | **Patch** /api/catalog-categories/{id}/relationships/items/ | Update Catalog Category Relationships Items
*CatalogsAPI* | [**UpdateCatalogItem**](docs/CatalogsAPI.md#updatecatalogitem) | **Patch** /api/catalog-items/{id}/ | Update Catalog Item
*CatalogsAPI* | [**UpdateCatalogItemRelationshipsCategories**](docs/CatalogsAPI.md#updatecatalogitemrelationshipscategories) | **Patch** /api/catalog-items/{id}/relationships/categories/ | Update Catalog Item Relationships Categories
*CatalogsAPI* | [**UpdateCatalogVariant**](docs/CatalogsAPI.md#updatecatalogvariant) | **Patch** /api/catalog-variants/{id}/ | Update Catalog Variant
*ClientAPI* | [**BulkCreateClientEvents**](docs/ClientAPI.md#bulkcreateclientevents) | **Post** /client/event-bulk-create/ | Bulk Create Client Events
*ClientAPI* | [**CreateClientBackInStockSubscription**](docs/ClientAPI.md#createclientbackinstocksubscription) | **Post** /client/back-in-stock-subscriptions/ | Create Client Back In Stock Subscription
*ClientAPI* | [**CreateClientEvent**](docs/ClientAPI.md#createclientevent) | **Post** /client/events/ | Create Client Event
*ClientAPI* | [**CreateClientProfile**](docs/ClientAPI.md#createclientprofile) | **Post** /client/profiles/ | Create or Update Client Profile
*ClientAPI* | [**CreateClientPushToken**](docs/ClientAPI.md#createclientpushtoken) | **Post** /client/push-tokens/ | Create or Update Client Push Token
*ClientAPI* | [**CreateClientSubscription**](docs/ClientAPI.md#createclientsubscription) | **Post** /client/subscriptions/ | Create Client Subscription
*ClientAPI* | [**UnregisterClientPushToken**](docs/ClientAPI.md#unregisterclientpushtoken) | **Post** /client/push-token-unregister/ | Unregister Client Push Token
*CouponsAPI* | [**CreateCoupon**](docs/CouponsAPI.md#createcoupon) | **Post** /api/coupons/ | Create Coupon
*CouponsAPI* | [**CreateCouponCode**](docs/CouponsAPI.md#createcouponcode) | **Post** /api/coupon-codes/ | Create Coupon Code
*CouponsAPI* | [**DeleteCoupon**](docs/CouponsAPI.md#deletecoupon) | **Delete** /api/coupons/{id}/ | Delete Coupon
*CouponsAPI* | [**DeleteCouponCode**](docs/CouponsAPI.md#deletecouponcode) | **Delete** /api/coupon-codes/{id}/ | Delete Coupon Code
*CouponsAPI* | [**GetCoupon**](docs/CouponsAPI.md#getcoupon) | **Get** /api/coupons/{id}/ | Get Coupon
*CouponsAPI* | [**GetCouponCode**](docs/CouponsAPI.md#getcouponcode) | **Get** /api/coupon-codes/{id}/ | Get Coupon Code
*CouponsAPI* | [**GetCouponCodeBulkCreateJob**](docs/CouponsAPI.md#getcouponcodebulkcreatejob) | **Get** /api/coupon-code-bulk-create-jobs/{job_id}/ | Get Coupon Code Bulk Create Job
*CouponsAPI* | [**GetCouponCodeBulkCreateJobs**](docs/CouponsAPI.md#getcouponcodebulkcreatejobs) | **Get** /api/coupon-code-bulk-create-jobs/ | Get Coupon Code Bulk Create Jobs
*CouponsAPI* | [**GetCouponCodeRelationshipsCoupon**](docs/CouponsAPI.md#getcouponcoderelationshipscoupon) | **Get** /api/coupons/{id}/relationships/coupon-codes/ | Get Coupon Code Relationships Coupon
*CouponsAPI* | [**GetCouponCodes**](docs/CouponsAPI.md#getcouponcodes) | **Get** /api/coupon-codes/ | Get Coupon Codes
*CouponsAPI* | [**GetCouponCodesForCoupon**](docs/CouponsAPI.md#getcouponcodesforcoupon) | **Get** /api/coupons/{id}/coupon-codes/ | Get Coupon Codes For Coupon
*CouponsAPI* | [**GetCouponForCouponCode**](docs/CouponsAPI.md#getcouponforcouponcode) | **Get** /api/coupon-codes/{id}/coupon/ | Get Coupon For Coupon Code
*CouponsAPI* | [**GetCouponRelationshipsCouponCodes**](docs/CouponsAPI.md#getcouponrelationshipscouponcodes) | **Get** /api/coupon-codes/{id}/relationships/coupon/ | Get Coupon Relationships Coupon Codes
*CouponsAPI* | [**GetCoupons**](docs/CouponsAPI.md#getcoupons) | **Get** /api/coupons/ | Get Coupons
*CouponsAPI* | [**SpawnCouponCodeBulkCreateJob**](docs/CouponsAPI.md#spawncouponcodebulkcreatejob) | **Post** /api/coupon-code-bulk-create-jobs/ | Spawn Coupon Code Bulk Create Job
*CouponsAPI* | [**UpdateCoupon**](docs/CouponsAPI.md#updatecoupon) | **Patch** /api/coupons/{id}/ | Update Coupon
*CouponsAPI* | [**UpdateCouponCode**](docs/CouponsAPI.md#updatecouponcode) | **Patch** /api/coupon-codes/{id}/ | Update Coupon Code
*DataPrivacyAPI* | [**RequestProfileDeletion**](docs/DataPrivacyAPI.md#requestprofiledeletion) | **Post** /api/data-privacy-deletion-jobs/ | Request Profile Deletion
*EventsAPI* | [**CreateEvent**](docs/EventsAPI.md#createevent) | **Post** /api/events/ | Create Event
*EventsAPI* | [**GetEvent**](docs/EventsAPI.md#getevent) | **Get** /api/events/{id}/ | Get Event
*EventsAPI* | [**GetEventMetric**](docs/EventsAPI.md#geteventmetric) | **Get** /api/events/{id}/metric/ | Get Event Metric
*EventsAPI* | [**GetEventProfile**](docs/EventsAPI.md#geteventprofile) | **Get** /api/events/{id}/profile/ | Get Event Profile
*EventsAPI* | [**GetEventRelationshipsMetric**](docs/EventsAPI.md#geteventrelationshipsmetric) | **Get** /api/events/{id}/relationships/metric/ | Get Event Relationships Metric
*EventsAPI* | [**GetEventRelationshipsProfile**](docs/EventsAPI.md#geteventrelationshipsprofile) | **Get** /api/events/{id}/relationships/profile/ | Get Event Relationships Profile
*EventsAPI* | [**GetEvents**](docs/EventsAPI.md#getevents) | **Get** /api/events/ | Get Events
*FlowsAPI* | [**GetFlow**](docs/FlowsAPI.md#getflow) | **Get** /api/flows/{id}/ | Get Flow
*FlowsAPI* | [**GetFlowAction**](docs/FlowsAPI.md#getflowaction) | **Get** /api/flow-actions/{id}/ | Get Flow Action
*FlowsAPI* | [**GetFlowActionFlow**](docs/FlowsAPI.md#getflowactionflow) | **Get** /api/flow-actions/{id}/flow/ | Get Flow For Flow Action
*FlowsAPI* | [**GetFlowActionMessages**](docs/FlowsAPI.md#getflowactionmessages) | **Get** /api/flow-actions/{id}/flow-messages/ | Get Flow Action Messages
*FlowsAPI* | [**GetFlowActionRelationshipsFlow**](docs/FlowsAPI.md#getflowactionrelationshipsflow) | **Get** /api/flow-actions/{id}/relationships/flow/ | Get Flow Action Relationships Flow
*FlowsAPI* | [**GetFlowActionRelationshipsMessages**](docs/FlowsAPI.md#getflowactionrelationshipsmessages) | **Get** /api/flow-actions/{id}/relationships/flow-messages/ | Get Flow Action Relationships Messages
*FlowsAPI* | [**GetFlowFlowActions**](docs/FlowsAPI.md#getflowflowactions) | **Get** /api/flows/{id}/flow-actions/ | Get Flow Flow Actions
*FlowsAPI* | [**GetFlowMessage**](docs/FlowsAPI.md#getflowmessage) | **Get** /api/flow-messages/{id}/ | Get Flow Message
*FlowsAPI* | [**GetFlowMessageAction**](docs/FlowsAPI.md#getflowmessageaction) | **Get** /api/flow-messages/{id}/flow-action/ | Get Flow Action For Message
*FlowsAPI* | [**GetFlowMessageRelationshipsAction**](docs/FlowsAPI.md#getflowmessagerelationshipsaction) | **Get** /api/flow-messages/{id}/relationships/flow-action/ | Get Flow Message Relationships Action
*FlowsAPI* | [**GetFlowMessageRelationshipsTemplate**](docs/FlowsAPI.md#getflowmessagerelationshipstemplate) | **Get** /api/flow-messages/{id}/relationships/template/ | Get Flow Message Relationships Template
*FlowsAPI* | [**GetFlowMessageTemplate**](docs/FlowsAPI.md#getflowmessagetemplate) | **Get** /api/flow-messages/{id}/template/ | Get Flow Message Template
*FlowsAPI* | [**GetFlowRelationshipsFlowActions**](docs/FlowsAPI.md#getflowrelationshipsflowactions) | **Get** /api/flows/{id}/relationships/flow-actions/ | Get Flow Relationships Flow Actions
*FlowsAPI* | [**GetFlowRelationshipsTags**](docs/FlowsAPI.md#getflowrelationshipstags) | **Get** /api/flows/{id}/relationships/tags/ | Get Flow Relationships Tags
*FlowsAPI* | [**GetFlowTags**](docs/FlowsAPI.md#getflowtags) | **Get** /api/flows/{id}/tags/ | Get Flow Tags
*FlowsAPI* | [**GetFlows**](docs/FlowsAPI.md#getflows) | **Get** /api/flows/ | Get Flows
*FlowsAPI* | [**UpdateFlow**](docs/FlowsAPI.md#updateflow) | **Patch** /api/flows/{id}/ | Update Flow Status
*ImagesAPI* | [**GetImage**](docs/ImagesAPI.md#getimage) | **Get** /api/images/{id}/ | Get Image
*ImagesAPI* | [**GetImages**](docs/ImagesAPI.md#getimages) | **Get** /api/images/ | Get Images
*ImagesAPI* | [**UpdateImage**](docs/ImagesAPI.md#updateimage) | **Patch** /api/images/{id}/ | Update Image
*ImagesAPI* | [**UploadImageFromFile**](docs/ImagesAPI.md#uploadimagefromfile) | **Post** /api/image-upload/ | Upload Image From File
*ImagesAPI* | [**UploadImageFromUrl**](docs/ImagesAPI.md#uploadimagefromurl) | **Post** /api/images/ | Upload Image From URL
*ListsAPI* | [**CreateList**](docs/ListsAPI.md#createlist) | **Post** /api/lists/ | Create List
*ListsAPI* | [**CreateListRelationships**](docs/ListsAPI.md#createlistrelationships) | **Post** /api/lists/{id}/relationships/profiles/ | Add Profile To List
*ListsAPI* | [**DeleteList**](docs/ListsAPI.md#deletelist) | **Delete** /api/lists/{id}/ | Delete List
*ListsAPI* | [**DeleteListRelationships**](docs/ListsAPI.md#deletelistrelationships) | **Delete** /api/lists/{id}/relationships/profiles/ | Remove Profile From List
*ListsAPI* | [**GetList**](docs/ListsAPI.md#getlist) | **Get** /api/lists/{id}/ | Get List
*ListsAPI* | [**GetListProfiles**](docs/ListsAPI.md#getlistprofiles) | **Get** /api/lists/{id}/profiles/ | Get List Profiles
*ListsAPI* | [**GetListRelationshipsProfiles**](docs/ListsAPI.md#getlistrelationshipsprofiles) | **Get** /api/lists/{id}/relationships/profiles/ | Get List Relationships Profiles
*ListsAPI* | [**GetListRelationshipsTags**](docs/ListsAPI.md#getlistrelationshipstags) | **Get** /api/lists/{id}/relationships/tags/ | Get List Relationships Tags
*ListsAPI* | [**GetListTags**](docs/ListsAPI.md#getlisttags) | **Get** /api/lists/{id}/tags/ | Get List Tags
*ListsAPI* | [**GetLists**](docs/ListsAPI.md#getlists) | **Get** /api/lists/ | Get Lists
*ListsAPI* | [**UpdateList**](docs/ListsAPI.md#updatelist) | **Patch** /api/lists/{id}/ | Update List
*MetricsAPI* | [**GetMetric**](docs/MetricsAPI.md#getmetric) | **Get** /api/metrics/{id}/ | Get Metric
*MetricsAPI* | [**GetMetrics**](docs/MetricsAPI.md#getmetrics) | **Get** /api/metrics/ | Get Metrics
*MetricsAPI* | [**QueryMetricAggregates**](docs/MetricsAPI.md#querymetricaggregates) | **Post** /api/metric-aggregates/ | Query Metric Aggregates
*ProfilesAPI* | [**CreateOrUpdateProfile**](docs/ProfilesAPI.md#createorupdateprofile) | **Post** /api/profile-import/ | Create or Update Profile
*ProfilesAPI* | [**CreateProfile**](docs/ProfilesAPI.md#createprofile) | **Post** /api/profiles/ | Create Profile
*ProfilesAPI* | [**CreatePushToken**](docs/ProfilesAPI.md#createpushtoken) | **Post** /api/push-tokens/ | Create or Update Push Token
*ProfilesAPI* | [**GetBulkProfileImportJob**](docs/ProfilesAPI.md#getbulkprofileimportjob) | **Get** /api/profile-bulk-import-jobs/{job_id}/ | Get Bulk Profile Import Job
*ProfilesAPI* | [**GetBulkProfileImportJobImportErrors**](docs/ProfilesAPI.md#getbulkprofileimportjobimporterrors) | **Get** /api/profile-bulk-import-jobs/{id}/import-errors/ | Get Bulk Profile Import Job Errors
*ProfilesAPI* | [**GetBulkProfileImportJobLists**](docs/ProfilesAPI.md#getbulkprofileimportjoblists) | **Get** /api/profile-bulk-import-jobs/{id}/lists/ | Get Bulk Profile Import Job Lists
*ProfilesAPI* | [**GetBulkProfileImportJobProfiles**](docs/ProfilesAPI.md#getbulkprofileimportjobprofiles) | **Get** /api/profile-bulk-import-jobs/{id}/profiles/ | Get Bulk Profile Import Job Profiles
*ProfilesAPI* | [**GetBulkProfileImportJobRelationshipsLists**](docs/ProfilesAPI.md#getbulkprofileimportjobrelationshipslists) | **Get** /api/profile-bulk-import-jobs/{id}/relationships/lists/ | Get Bulk Profile Import Job Relationships Lists
*ProfilesAPI* | [**GetBulkProfileImportJobRelationshipsProfiles**](docs/ProfilesAPI.md#getbulkprofileimportjobrelationshipsprofiles) | **Get** /api/profile-bulk-import-jobs/{id}/relationships/profiles/ | Get Bulk Profile Import Job Relationships Profiles
*ProfilesAPI* | [**GetBulkProfileImportJobs**](docs/ProfilesAPI.md#getbulkprofileimportjobs) | **Get** /api/profile-bulk-import-jobs/ | Get Bulk Profile Import Jobs
*ProfilesAPI* | [**GetProfile**](docs/ProfilesAPI.md#getprofile) | **Get** /api/profiles/{id}/ | Get Profile
*ProfilesAPI* | [**GetProfileLists**](docs/ProfilesAPI.md#getprofilelists) | **Get** /api/profiles/{id}/lists/ | Get Profile Lists
*ProfilesAPI* | [**GetProfileRelationshipsLists**](docs/ProfilesAPI.md#getprofilerelationshipslists) | **Get** /api/profiles/{id}/relationships/lists/ | Get Profile Relationships Lists
*ProfilesAPI* | [**GetProfileRelationshipsSegments**](docs/ProfilesAPI.md#getprofilerelationshipssegments) | **Get** /api/profiles/{id}/relationships/segments/ | Get Profile Relationships Segments
*ProfilesAPI* | [**GetProfileSegments**](docs/ProfilesAPI.md#getprofilesegments) | **Get** /api/profiles/{id}/segments/ | Get Profile Segments
*ProfilesAPI* | [**GetProfiles**](docs/ProfilesAPI.md#getprofiles) | **Get** /api/profiles/ | Get Profiles
*ProfilesAPI* | [**MergeProfiles**](docs/ProfilesAPI.md#mergeprofiles) | **Post** /api/profile-merge/ | Merge Profiles
*ProfilesAPI* | [**SpawnBulkProfileImportJob**](docs/ProfilesAPI.md#spawnbulkprofileimportjob) | **Post** /api/profile-bulk-import-jobs/ | Spawn Bulk Profile Import Job
*ProfilesAPI* | [**SubscribeProfiles**](docs/ProfilesAPI.md#subscribeprofiles) | **Post** /api/profile-subscription-bulk-create-jobs/ | Subscribe Profiles
*ProfilesAPI* | [**SuppressProfiles**](docs/ProfilesAPI.md#suppressprofiles) | **Post** /api/profile-suppression-bulk-create-jobs/ | Suppress Profiles
*ProfilesAPI* | [**UnsubscribeProfiles**](docs/ProfilesAPI.md#unsubscribeprofiles) | **Post** /api/profile-subscription-bulk-delete-jobs/ | Unsubscribe Profiles
*ProfilesAPI* | [**UnsuppressProfiles**](docs/ProfilesAPI.md#unsuppressprofiles) | **Post** /api/profile-suppression-bulk-delete-jobs/ | Unsuppress Profiles
*ProfilesAPI* | [**UpdateProfile**](docs/ProfilesAPI.md#updateprofile) | **Patch** /api/profiles/{id}/ | Update Profile
*ReportingAPI* | [**QueryCampaignValues**](docs/ReportingAPI.md#querycampaignvalues) | **Post** /api/campaign-values-reports/ | Query Campaign Values
*ReportingAPI* | [**QueryFlowSeries**](docs/ReportingAPI.md#queryflowseries) | **Post** /api/flow-series-reports/ | Query Flow Series
*ReportingAPI* | [**QueryFlowValues**](docs/ReportingAPI.md#queryflowvalues) | **Post** /api/flow-values-reports/ | Query Flow Values
*SegmentsAPI* | [**GetSegment**](docs/SegmentsAPI.md#getsegment) | **Get** /api/segments/{id}/ | Get Segment
*SegmentsAPI* | [**GetSegmentProfiles**](docs/SegmentsAPI.md#getsegmentprofiles) | **Get** /api/segments/{id}/profiles/ | Get Segment Profiles
*SegmentsAPI* | [**GetSegmentRelationshipsProfiles**](docs/SegmentsAPI.md#getsegmentrelationshipsprofiles) | **Get** /api/segments/{id}/relationships/profiles/ | Get Segment Relationships Profiles
*SegmentsAPI* | [**GetSegmentRelationshipsTags**](docs/SegmentsAPI.md#getsegmentrelationshipstags) | **Get** /api/segments/{id}/relationships/tags/ | Get Segment Relationships Tags
*SegmentsAPI* | [**GetSegmentTags**](docs/SegmentsAPI.md#getsegmenttags) | **Get** /api/segments/{id}/tags/ | Get Segment Tags
*SegmentsAPI* | [**GetSegments**](docs/SegmentsAPI.md#getsegments) | **Get** /api/segments/ | Get Segments
*SegmentsAPI* | [**UpdateSegment**](docs/SegmentsAPI.md#updatesegment) | **Patch** /api/segments/{id}/ | Update Segment
*TagsAPI* | [**CreateTag**](docs/TagsAPI.md#createtag) | **Post** /api/tags/ | Create Tag
*TagsAPI* | [**CreateTagGroup**](docs/TagsAPI.md#createtaggroup) | **Post** /api/tag-groups/ | Create Tag Group
*TagsAPI* | [**CreateTagRelationshipsCampaigns**](docs/TagsAPI.md#createtagrelationshipscampaigns) | **Post** /api/tags/{id}/relationships/campaigns/ | Create Tag Relationships Campaigns
*TagsAPI* | [**CreateTagRelationshipsFlows**](docs/TagsAPI.md#createtagrelationshipsflows) | **Post** /api/tags/{id}/relationships/flows/ | Create Tag Relationships Flows
*TagsAPI* | [**CreateTagRelationshipsLists**](docs/TagsAPI.md#createtagrelationshipslists) | **Post** /api/tags/{id}/relationships/lists/ | Create Tag Relationships Lists
*TagsAPI* | [**CreateTagRelationshipsSegments**](docs/TagsAPI.md#createtagrelationshipssegments) | **Post** /api/tags/{id}/relationships/segments/ | Create Tag Relationships Segments
*TagsAPI* | [**DeleteTag**](docs/TagsAPI.md#deletetag) | **Delete** /api/tags/{id}/ | Delete Tag
*TagsAPI* | [**DeleteTagGroup**](docs/TagsAPI.md#deletetaggroup) | **Delete** /api/tag-groups/{id}/ | Delete Tag Group
*TagsAPI* | [**DeleteTagRelationshipsCampaigns**](docs/TagsAPI.md#deletetagrelationshipscampaigns) | **Delete** /api/tags/{id}/relationships/campaigns/ | Delete Tag Relationships Campaigns
*TagsAPI* | [**DeleteTagRelationshipsFlows**](docs/TagsAPI.md#deletetagrelationshipsflows) | **Delete** /api/tags/{id}/relationships/flows/ | Delete Tag Relationships Flows
*TagsAPI* | [**DeleteTagRelationshipsLists**](docs/TagsAPI.md#deletetagrelationshipslists) | **Delete** /api/tags/{id}/relationships/lists/ | Delete Tag Relationships Lists
*TagsAPI* | [**DeleteTagRelationshipsSegments**](docs/TagsAPI.md#deletetagrelationshipssegments) | **Delete** /api/tags/{id}/relationships/segments/ | Delete Tag Relationships Segments
*TagsAPI* | [**GetTag**](docs/TagsAPI.md#gettag) | **Get** /api/tags/{id}/ | Get Tag
*TagsAPI* | [**GetTagGroup**](docs/TagsAPI.md#gettaggroup) | **Get** /api/tag-groups/{id}/ | Get Tag Group
*TagsAPI* | [**GetTagGroupRelationshipsTags**](docs/TagsAPI.md#gettaggrouprelationshipstags) | **Get** /api/tag-groups/{id}/relationships/tags/ | Get Tag Group Relationships Tags
*TagsAPI* | [**GetTagGroupTags**](docs/TagsAPI.md#gettaggrouptags) | **Get** /api/tag-groups/{id}/tags/ | Get Tag Group Tags
*TagsAPI* | [**GetTagGroups**](docs/TagsAPI.md#gettaggroups) | **Get** /api/tag-groups/ | Get Tag Groups
*TagsAPI* | [**GetTagRelationshipsCampaigns**](docs/TagsAPI.md#gettagrelationshipscampaigns) | **Get** /api/tags/{id}/relationships/campaigns/ | Get Tag Relationships Campaigns
*TagsAPI* | [**GetTagRelationshipsFlows**](docs/TagsAPI.md#gettagrelationshipsflows) | **Get** /api/tags/{id}/relationships/flows/ | Get Tag Relationships Flows
*TagsAPI* | [**GetTagRelationshipsLists**](docs/TagsAPI.md#gettagrelationshipslists) | **Get** /api/tags/{id}/relationships/lists/ | Get Tag Relationships Lists
*TagsAPI* | [**GetTagRelationshipsSegments**](docs/TagsAPI.md#gettagrelationshipssegments) | **Get** /api/tags/{id}/relationships/segments/ | Get Tag Relationships Segments
*TagsAPI* | [**GetTagRelationshipsTagGroup**](docs/TagsAPI.md#gettagrelationshipstaggroup) | **Get** /api/tags/{id}/relationships/tag-group/ | Get Tag Relationships Tag Group
*TagsAPI* | [**GetTagTagGroup**](docs/TagsAPI.md#gettagtaggroup) | **Get** /api/tags/{id}/tag-group/ | Get Tag Tag Group
*TagsAPI* | [**GetTags**](docs/TagsAPI.md#gettags) | **Get** /api/tags/ | Get Tags
*TagsAPI* | [**UpdateTag**](docs/TagsAPI.md#updatetag) | **Patch** /api/tags/{id}/ | Update Tag
*TagsAPI* | [**UpdateTagGroup**](docs/TagsAPI.md#updatetaggroup) | **Patch** /api/tag-groups/{id}/ | Update Tag Group
*TemplatesAPI* | [**CreateTemplate**](docs/TemplatesAPI.md#createtemplate) | **Post** /api/templates/ | Create Template
*TemplatesAPI* | [**CreateTemplateClone**](docs/TemplatesAPI.md#createtemplateclone) | **Post** /api/template-clone/ | Create Template Clone
*TemplatesAPI* | [**CreateTemplateRender**](docs/TemplatesAPI.md#createtemplaterender) | **Post** /api/template-render/ | Create Template Render
*TemplatesAPI* | [**DeleteTemplate**](docs/TemplatesAPI.md#deletetemplate) | **Delete** /api/templates/{id}/ | Delete Template
*TemplatesAPI* | [**GetTemplate**](docs/TemplatesAPI.md#gettemplate) | **Get** /api/templates/{id}/ | Get Template
*TemplatesAPI* | [**GetTemplates**](docs/TemplatesAPI.md#gettemplates) | **Get** /api/templates/ | Get Templates
*TemplatesAPI* | [**UpdateTemplate**](docs/TemplatesAPI.md#updatetemplate) | **Patch** /api/templates/{id}/ | Update Template

## Documentation For Models

- [APIJobErrorPayload](docs/APIJobErrorPayload.md)
- [AccountEnum](docs/AccountEnum.md)
- [AccountResponseObjectResource](docs/AccountResponseObjectResource.md)
- [AccountResponseObjectResourceAttributes](docs/AccountResponseObjectResourceAttributes.md)
- [AttributionEnum](docs/AttributionEnum.md)
- [AttributionResponseObjectResource](docs/AttributionResponseObjectResource.md)
- [AttributionResponseObjectResourceRelationships](docs/AttributionResponseObjectResourceRelationships.md)
- [AttributionResponseObjectResourceRelationshipsAttributedEvent](docs/AttributionResponseObjectResourceRelationshipsAttributedEvent.md)
- [AttributionResponseObjectResourceRelationshipsAttributedEventData](docs/AttributionResponseObjectResourceRelationshipsAttributedEventData.md)
- [AttributionResponseObjectResourceRelationshipsCampaign](docs/AttributionResponseObjectResourceRelationshipsCampaign.md)
- [AttributionResponseObjectResourceRelationshipsCampaignData](docs/AttributionResponseObjectResourceRelationshipsCampaignData.md)
- [AttributionResponseObjectResourceRelationshipsCampaignMessage](docs/AttributionResponseObjectResourceRelationshipsCampaignMessage.md)
- [AttributionResponseObjectResourceRelationshipsCampaignMessageData](docs/AttributionResponseObjectResourceRelationshipsCampaignMessageData.md)
- [AttributionResponseObjectResourceRelationshipsEvent](docs/AttributionResponseObjectResourceRelationshipsEvent.md)
- [AttributionResponseObjectResourceRelationshipsEventData](docs/AttributionResponseObjectResourceRelationshipsEventData.md)
- [AttributionResponseObjectResourceRelationshipsFlow](docs/AttributionResponseObjectResourceRelationshipsFlow.md)
- [AttributionResponseObjectResourceRelationshipsFlowData](docs/AttributionResponseObjectResourceRelationshipsFlowData.md)
- [AttributionResponseObjectResourceRelationshipsFlowMessage](docs/AttributionResponseObjectResourceRelationshipsFlowMessage.md)
- [AttributionResponseObjectResourceRelationshipsFlowMessageData](docs/AttributionResponseObjectResourceRelationshipsFlowMessageData.md)
- [AttributionResponseObjectResourceRelationshipsFlowMessageVariation](docs/AttributionResponseObjectResourceRelationshipsFlowMessageVariation.md)
- [AttributionResponseObjectResourceRelationshipsFlowMessageVariationData](docs/AttributionResponseObjectResourceRelationshipsFlowMessageVariationData.md)
- [AudiencesSubObject](docs/AudiencesSubObject.md)
- [BackInStockSubscriptionEnum](docs/BackInStockSubscriptionEnum.md)
- [BaseEventCreateQueryResourceObject](docs/BaseEventCreateQueryResourceObject.md)
- [BaseEventCreateQueryResourceObjectAttributes](docs/BaseEventCreateQueryResourceObjectAttributes.md)
- [CampaignCloneQuery](docs/CampaignCloneQuery.md)
- [CampaignCloneQueryResourceObject](docs/CampaignCloneQueryResourceObject.md)
- [CampaignCloneQueryResourceObjectAttributes](docs/CampaignCloneQueryResourceObjectAttributes.md)
- [CampaignCreateQuery](docs/CampaignCreateQuery.md)
- [CampaignCreateQueryResourceObject](docs/CampaignCreateQueryResourceObject.md)
- [CampaignCreateQueryResourceObjectAttributes](docs/CampaignCreateQueryResourceObjectAttributes.md)
- [CampaignCreateQueryResourceObjectAttributesCampaignMessages](docs/CampaignCreateQueryResourceObjectAttributesCampaignMessages.md)
- [CampaignCreateQueryResourceObjectAttributesSendOptions](docs/CampaignCreateQueryResourceObjectAttributesSendOptions.md)
- [CampaignCreateQueryResourceObjectAttributesTrackingOptions](docs/CampaignCreateQueryResourceObjectAttributesTrackingOptions.md)
- [CampaignEnum](docs/CampaignEnum.md)
- [CampaignMessageAssignTemplateQuery](docs/CampaignMessageAssignTemplateQuery.md)
- [CampaignMessageAssignTemplateQueryResourceObject](docs/CampaignMessageAssignTemplateQueryResourceObject.md)
- [CampaignMessageAssignTemplateQueryResourceObjectRelationships](docs/CampaignMessageAssignTemplateQueryResourceObjectRelationships.md)
- [CampaignMessageAssignTemplateQueryResourceObjectRelationshipsTemplate](docs/CampaignMessageAssignTemplateQueryResourceObjectRelationshipsTemplate.md)
- [CampaignMessageAssignTemplateQueryResourceObjectRelationshipsTemplateData](docs/CampaignMessageAssignTemplateQueryResourceObjectRelationshipsTemplateData.md)
- [CampaignMessageCreateQueryResourceObject](docs/CampaignMessageCreateQueryResourceObject.md)
- [CampaignMessageCreateQueryResourceObjectAttributes](docs/CampaignMessageCreateQueryResourceObjectAttributes.md)
- [CampaignMessageCreateQueryResourceObjectAttributesContent](docs/CampaignMessageCreateQueryResourceObjectAttributesContent.md)
- [CampaignMessageEnum](docs/CampaignMessageEnum.md)
- [CampaignMessagePartialUpdateQuery](docs/CampaignMessagePartialUpdateQuery.md)
- [CampaignMessagePartialUpdateQueryResourceObject](docs/CampaignMessagePartialUpdateQueryResourceObject.md)
- [CampaignMessagePartialUpdateQueryResourceObjectAttributes](docs/CampaignMessagePartialUpdateQueryResourceObjectAttributes.md)
- [CampaignMessageResponseObjectResource](docs/CampaignMessageResponseObjectResource.md)
- [CampaignMessageResponseObjectResourceAttributes](docs/CampaignMessageResponseObjectResourceAttributes.md)
- [CampaignMessageResponseObjectResourceAttributesContent](docs/CampaignMessageResponseObjectResourceAttributesContent.md)
- [CampaignPartialUpdateQuery](docs/CampaignPartialUpdateQuery.md)
- [CampaignPartialUpdateQueryResourceObject](docs/CampaignPartialUpdateQueryResourceObject.md)
- [CampaignPartialUpdateQueryResourceObjectAttributes](docs/CampaignPartialUpdateQueryResourceObjectAttributes.md)
- [CampaignRecipientEstimationEnum](docs/CampaignRecipientEstimationEnum.md)
- [CampaignRecipientEstimationJobCreateQuery](docs/CampaignRecipientEstimationJobCreateQuery.md)
- [CampaignRecipientEstimationJobCreateQueryResourceObject](docs/CampaignRecipientEstimationJobCreateQueryResourceObject.md)
- [CampaignRecipientEstimationJobEnum](docs/CampaignRecipientEstimationJobEnum.md)
- [CampaignRecipientEstimationJobResponseObjectResource](docs/CampaignRecipientEstimationJobResponseObjectResource.md)
- [CampaignRecipientEstimationJobResponseObjectResourceAttributes](docs/CampaignRecipientEstimationJobResponseObjectResourceAttributes.md)
- [CampaignRecipientEstimationResponseObjectResource](docs/CampaignRecipientEstimationResponseObjectResource.md)
- [CampaignRecipientEstimationResponseObjectResourceAttributes](docs/CampaignRecipientEstimationResponseObjectResourceAttributes.md)
- [CampaignResponseObjectResource](docs/CampaignResponseObjectResource.md)
- [CampaignResponseObjectResourceAttributes](docs/CampaignResponseObjectResourceAttributes.md)
- [CampaignResponseObjectResourceAttributesSendOptions](docs/CampaignResponseObjectResourceAttributesSendOptions.md)
- [CampaignResponseObjectResourceAttributesTrackingOptions](docs/CampaignResponseObjectResourceAttributesTrackingOptions.md)
- [CampaignSendJobCreateQuery](docs/CampaignSendJobCreateQuery.md)
- [CampaignSendJobCreateQueryResourceObject](docs/CampaignSendJobCreateQueryResourceObject.md)
- [CampaignSendJobEnum](docs/CampaignSendJobEnum.md)
- [CampaignSendJobPartialUpdateQuery](docs/CampaignSendJobPartialUpdateQuery.md)
- [CampaignSendJobPartialUpdateQueryResourceObject](docs/CampaignSendJobPartialUpdateQueryResourceObject.md)
- [CampaignSendJobPartialUpdateQueryResourceObjectAttributes](docs/CampaignSendJobPartialUpdateQueryResourceObjectAttributes.md)
- [CampaignSendJobResponseObjectResource](docs/CampaignSendJobResponseObjectResource.md)
- [CampaignSendJobResponseObjectResourceAttributes](docs/CampaignSendJobResponseObjectResourceAttributes.md)
- [CampaignValuesReportEnum](docs/CampaignValuesReportEnum.md)
- [CampaignValuesRequestDTO](docs/CampaignValuesRequestDTO.md)
- [CampaignValuesRequestDTOResourceObject](docs/CampaignValuesRequestDTOResourceObject.md)
- [CampaignValuesRequestDTOResourceObjectAttributes](docs/CampaignValuesRequestDTOResourceObjectAttributes.md)
- [CampaignValuesRequestDTOResourceObjectAttributesTimeframe](docs/CampaignValuesRequestDTOResourceObjectAttributesTimeframe.md)
- [CatalogCategoryBulkCreateJobEnum](docs/CatalogCategoryBulkCreateJobEnum.md)
- [CatalogCategoryBulkDeleteJobEnum](docs/CatalogCategoryBulkDeleteJobEnum.md)
- [CatalogCategoryBulkUpdateJobEnum](docs/CatalogCategoryBulkUpdateJobEnum.md)
- [CatalogCategoryCreateJobCreateQuery](docs/CatalogCategoryCreateJobCreateQuery.md)
- [CatalogCategoryCreateJobCreateQueryResourceObject](docs/CatalogCategoryCreateJobCreateQueryResourceObject.md)
- [CatalogCategoryCreateJobCreateQueryResourceObjectAttributes](docs/CatalogCategoryCreateJobCreateQueryResourceObjectAttributes.md)
- [CatalogCategoryCreateJobCreateQueryResourceObjectAttributesCategories](docs/CatalogCategoryCreateJobCreateQueryResourceObjectAttributesCategories.md)
- [CatalogCategoryCreateJobResponseObjectResource](docs/CatalogCategoryCreateJobResponseObjectResource.md)
- [CatalogCategoryCreateQuery](docs/CatalogCategoryCreateQuery.md)
- [CatalogCategoryCreateQueryResourceObject](docs/CatalogCategoryCreateQueryResourceObject.md)
- [CatalogCategoryCreateQueryResourceObjectAttributes](docs/CatalogCategoryCreateQueryResourceObjectAttributes.md)
- [CatalogCategoryCreateQueryResourceObjectRelationships](docs/CatalogCategoryCreateQueryResourceObjectRelationships.md)
- [CatalogCategoryCreateQueryResourceObjectRelationshipsItems](docs/CatalogCategoryCreateQueryResourceObjectRelationshipsItems.md)
- [CatalogCategoryDeleteJobCreateQuery](docs/CatalogCategoryDeleteJobCreateQuery.md)
- [CatalogCategoryDeleteJobCreateQueryResourceObject](docs/CatalogCategoryDeleteJobCreateQueryResourceObject.md)
- [CatalogCategoryDeleteJobCreateQueryResourceObjectAttributes](docs/CatalogCategoryDeleteJobCreateQueryResourceObjectAttributes.md)
- [CatalogCategoryDeleteJobCreateQueryResourceObjectAttributesCategories](docs/CatalogCategoryDeleteJobCreateQueryResourceObjectAttributesCategories.md)
- [CatalogCategoryDeleteJobResponseObjectResource](docs/CatalogCategoryDeleteJobResponseObjectResource.md)
- [CatalogCategoryDeleteQueryResourceObject](docs/CatalogCategoryDeleteQueryResourceObject.md)
- [CatalogCategoryEnum](docs/CatalogCategoryEnum.md)
- [CatalogCategoryItemOp](docs/CatalogCategoryItemOp.md)
- [CatalogCategoryResponseObjectResource](docs/CatalogCategoryResponseObjectResource.md)
- [CatalogCategoryResponseObjectResourceAttributes](docs/CatalogCategoryResponseObjectResourceAttributes.md)
- [CatalogCategoryUpdateJobCreateQuery](docs/CatalogCategoryUpdateJobCreateQuery.md)
- [CatalogCategoryUpdateJobCreateQueryResourceObject](docs/CatalogCategoryUpdateJobCreateQueryResourceObject.md)
- [CatalogCategoryUpdateJobCreateQueryResourceObjectAttributes](docs/CatalogCategoryUpdateJobCreateQueryResourceObjectAttributes.md)
- [CatalogCategoryUpdateJobCreateQueryResourceObjectAttributesCategories](docs/CatalogCategoryUpdateJobCreateQueryResourceObjectAttributesCategories.md)
- [CatalogCategoryUpdateJobResponseObjectResource](docs/CatalogCategoryUpdateJobResponseObjectResource.md)
- [CatalogCategoryUpdateQuery](docs/CatalogCategoryUpdateQuery.md)
- [CatalogCategoryUpdateQueryResourceObject](docs/CatalogCategoryUpdateQueryResourceObject.md)
- [CatalogCategoryUpdateQueryResourceObjectAttributes](docs/CatalogCategoryUpdateQueryResourceObjectAttributes.md)
- [CatalogItemBulkCreateJobEnum](docs/CatalogItemBulkCreateJobEnum.md)
- [CatalogItemBulkDeleteJobEnum](docs/CatalogItemBulkDeleteJobEnum.md)
- [CatalogItemBulkUpdateJobEnum](docs/CatalogItemBulkUpdateJobEnum.md)
- [CatalogItemCategoryOp](docs/CatalogItemCategoryOp.md)
- [CatalogItemCreateJobCreateQuery](docs/CatalogItemCreateJobCreateQuery.md)
- [CatalogItemCreateJobCreateQueryResourceObject](docs/CatalogItemCreateJobCreateQueryResourceObject.md)
- [CatalogItemCreateJobCreateQueryResourceObjectAttributes](docs/CatalogItemCreateJobCreateQueryResourceObjectAttributes.md)
- [CatalogItemCreateJobCreateQueryResourceObjectAttributesItems](docs/CatalogItemCreateJobCreateQueryResourceObjectAttributesItems.md)
- [CatalogItemCreateJobResponseObjectResource](docs/CatalogItemCreateJobResponseObjectResource.md)
- [CatalogItemCreateQuery](docs/CatalogItemCreateQuery.md)
- [CatalogItemCreateQueryResourceObject](docs/CatalogItemCreateQueryResourceObject.md)
- [CatalogItemCreateQueryResourceObjectAttributes](docs/CatalogItemCreateQueryResourceObjectAttributes.md)
- [CatalogItemCreateQueryResourceObjectRelationships](docs/CatalogItemCreateQueryResourceObjectRelationships.md)
- [CatalogItemCreateQueryResourceObjectRelationshipsCategories](docs/CatalogItemCreateQueryResourceObjectRelationshipsCategories.md)
- [CatalogItemDeleteJobCreateQuery](docs/CatalogItemDeleteJobCreateQuery.md)
- [CatalogItemDeleteJobCreateQueryResourceObject](docs/CatalogItemDeleteJobCreateQueryResourceObject.md)
- [CatalogItemDeleteJobCreateQueryResourceObjectAttributes](docs/CatalogItemDeleteJobCreateQueryResourceObjectAttributes.md)
- [CatalogItemDeleteJobCreateQueryResourceObjectAttributesItems](docs/CatalogItemDeleteJobCreateQueryResourceObjectAttributesItems.md)
- [CatalogItemDeleteJobResponseObjectResource](docs/CatalogItemDeleteJobResponseObjectResource.md)
- [CatalogItemDeleteQueryResourceObject](docs/CatalogItemDeleteQueryResourceObject.md)
- [CatalogItemEnum](docs/CatalogItemEnum.md)
- [CatalogItemResponseObjectResource](docs/CatalogItemResponseObjectResource.md)
- [CatalogItemResponseObjectResourceAttributes](docs/CatalogItemResponseObjectResourceAttributes.md)
- [CatalogItemUpdateJobCreateQuery](docs/CatalogItemUpdateJobCreateQuery.md)
- [CatalogItemUpdateJobCreateQueryResourceObject](docs/CatalogItemUpdateJobCreateQueryResourceObject.md)
- [CatalogItemUpdateJobCreateQueryResourceObjectAttributes](docs/CatalogItemUpdateJobCreateQueryResourceObjectAttributes.md)
- [CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems](docs/CatalogItemUpdateJobCreateQueryResourceObjectAttributesItems.md)
- [CatalogItemUpdateJobResponseObjectResource](docs/CatalogItemUpdateJobResponseObjectResource.md)
- [CatalogItemUpdateQuery](docs/CatalogItemUpdateQuery.md)
- [CatalogItemUpdateQueryResourceObject](docs/CatalogItemUpdateQueryResourceObject.md)
- [CatalogItemUpdateQueryResourceObjectAttributes](docs/CatalogItemUpdateQueryResourceObjectAttributes.md)
- [CatalogVariantBulkCreateJobEnum](docs/CatalogVariantBulkCreateJobEnum.md)
- [CatalogVariantBulkDeleteJobEnum](docs/CatalogVariantBulkDeleteJobEnum.md)
- [CatalogVariantBulkUpdateJobEnum](docs/CatalogVariantBulkUpdateJobEnum.md)
- [CatalogVariantCreateJobCreateQuery](docs/CatalogVariantCreateJobCreateQuery.md)
- [CatalogVariantCreateJobCreateQueryResourceObject](docs/CatalogVariantCreateJobCreateQueryResourceObject.md)
- [CatalogVariantCreateJobCreateQueryResourceObjectAttributes](docs/CatalogVariantCreateJobCreateQueryResourceObjectAttributes.md)
- [CatalogVariantCreateJobCreateQueryResourceObjectAttributesVariants](docs/CatalogVariantCreateJobCreateQueryResourceObjectAttributesVariants.md)
- [CatalogVariantCreateJobResponseObjectResource](docs/CatalogVariantCreateJobResponseObjectResource.md)
- [CatalogVariantCreateQuery](docs/CatalogVariantCreateQuery.md)
- [CatalogVariantCreateQueryResourceObject](docs/CatalogVariantCreateQueryResourceObject.md)
- [CatalogVariantCreateQueryResourceObjectAttributes](docs/CatalogVariantCreateQueryResourceObjectAttributes.md)
- [CatalogVariantCreateQueryResourceObjectRelationships](docs/CatalogVariantCreateQueryResourceObjectRelationships.md)
- [CatalogVariantCreateQueryResourceObjectRelationshipsItem](docs/CatalogVariantCreateQueryResourceObjectRelationshipsItem.md)
- [CatalogVariantCreateQueryResourceObjectRelationshipsItemData](docs/CatalogVariantCreateQueryResourceObjectRelationshipsItemData.md)
- [CatalogVariantDeleteJobCreateQuery](docs/CatalogVariantDeleteJobCreateQuery.md)
- [CatalogVariantDeleteJobCreateQueryResourceObject](docs/CatalogVariantDeleteJobCreateQueryResourceObject.md)
- [CatalogVariantDeleteJobCreateQueryResourceObjectAttributes](docs/CatalogVariantDeleteJobCreateQueryResourceObjectAttributes.md)
- [CatalogVariantDeleteJobCreateQueryResourceObjectAttributesVariants](docs/CatalogVariantDeleteJobCreateQueryResourceObjectAttributesVariants.md)
- [CatalogVariantDeleteJobResponseObjectResource](docs/CatalogVariantDeleteJobResponseObjectResource.md)
- [CatalogVariantDeleteQueryResourceObject](docs/CatalogVariantDeleteQueryResourceObject.md)
- [CatalogVariantEnum](docs/CatalogVariantEnum.md)
- [CatalogVariantResponseObjectResource](docs/CatalogVariantResponseObjectResource.md)
- [CatalogVariantResponseObjectResourceAttributes](docs/CatalogVariantResponseObjectResourceAttributes.md)
- [CatalogVariantUpdateJobCreateQuery](docs/CatalogVariantUpdateJobCreateQuery.md)
- [CatalogVariantUpdateJobCreateQueryResourceObject](docs/CatalogVariantUpdateJobCreateQueryResourceObject.md)
- [CatalogVariantUpdateJobCreateQueryResourceObjectAttributes](docs/CatalogVariantUpdateJobCreateQueryResourceObjectAttributes.md)
- [CatalogVariantUpdateJobCreateQueryResourceObjectAttributesVariants](docs/CatalogVariantUpdateJobCreateQueryResourceObjectAttributesVariants.md)
- [CatalogVariantUpdateJobResponseObjectResource](docs/CatalogVariantUpdateJobResponseObjectResource.md)
- [CatalogVariantUpdateQuery](docs/CatalogVariantUpdateQuery.md)
- [CatalogVariantUpdateQueryResourceObject](docs/CatalogVariantUpdateQueryResourceObject.md)
- [CatalogVariantUpdateQueryResourceObjectAttributes](docs/CatalogVariantUpdateQueryResourceObjectAttributes.md)
- [ClientBISSubscriptionCreateQuery](docs/ClientBISSubscriptionCreateQuery.md)
- [ClientBISSubscriptionCreateQueryResourceObject](docs/ClientBISSubscriptionCreateQueryResourceObject.md)
- [ClientBISSubscriptionCreateQueryResourceObjectAttributes](docs/ClientBISSubscriptionCreateQueryResourceObjectAttributes.md)
- [ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile](docs/ClientBISSubscriptionCreateQueryResourceObjectAttributesProfile.md)
- [CollectionLinks](docs/CollectionLinks.md)
- [ContactInformation](docs/ContactInformation.md)
- [CouponCodeBulkCreateJobEnum](docs/CouponCodeBulkCreateJobEnum.md)
- [CouponCodeCreateJobCreateQuery](docs/CouponCodeCreateJobCreateQuery.md)
- [CouponCodeCreateJobCreateQueryResourceObject](docs/CouponCodeCreateJobCreateQueryResourceObject.md)
- [CouponCodeCreateJobCreateQueryResourceObjectAttributes](docs/CouponCodeCreateJobCreateQueryResourceObjectAttributes.md)
- [CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes](docs/CouponCodeCreateJobCreateQueryResourceObjectAttributesCouponCodes.md)
- [CouponCodeCreateJobResponseObjectResource](docs/CouponCodeCreateJobResponseObjectResource.md)
- [CouponCodeCreateJobResponseObjectResourceAttributes](docs/CouponCodeCreateJobResponseObjectResourceAttributes.md)
- [CouponCodeCreateQuery](docs/CouponCodeCreateQuery.md)
- [CouponCodeCreateQueryResourceObject](docs/CouponCodeCreateQueryResourceObject.md)
- [CouponCodeCreateQueryResourceObjectAttributes](docs/CouponCodeCreateQueryResourceObjectAttributes.md)
- [CouponCodeCreateQueryResourceObjectRelationships](docs/CouponCodeCreateQueryResourceObjectRelationships.md)
- [CouponCodeCreateQueryResourceObjectRelationshipsCoupon](docs/CouponCodeCreateQueryResourceObjectRelationshipsCoupon.md)
- [CouponCodeEnum](docs/CouponCodeEnum.md)
- [CouponCodeResponseObjectResource](docs/CouponCodeResponseObjectResource.md)
- [CouponCodeResponseObjectResourceAttributes](docs/CouponCodeResponseObjectResourceAttributes.md)
- [CouponCodeUpdateQuery](docs/CouponCodeUpdateQuery.md)
- [CouponCodeUpdateQueryResourceObject](docs/CouponCodeUpdateQueryResourceObject.md)
- [CouponCodeUpdateQueryResourceObjectAttributes](docs/CouponCodeUpdateQueryResourceObjectAttributes.md)
- [CouponCreateQuery](docs/CouponCreateQuery.md)
- [CouponCreateQueryResourceObject](docs/CouponCreateQueryResourceObject.md)
- [CouponEnum](docs/CouponEnum.md)
- [CouponResponseObjectResource](docs/CouponResponseObjectResource.md)
- [CouponResponseObjectResourceAttributes](docs/CouponResponseObjectResourceAttributes.md)
- [CouponUpdateQuery](docs/CouponUpdateQuery.md)
- [CouponUpdateQueryResourceObject](docs/CouponUpdateQueryResourceObject.md)
- [CouponUpdateQueryResourceObjectAttributes](docs/CouponUpdateQueryResourceObjectAttributes.md)
- [CustomTimeframe](docs/CustomTimeframe.md)
- [DataPrivacyCreateDeletionJobQuery](docs/DataPrivacyCreateDeletionJobQuery.md)
- [DataPrivacyCreateDeletionJobQueryResourceObject](docs/DataPrivacyCreateDeletionJobQueryResourceObject.md)
- [DataPrivacyCreateDeletionJobQueryResourceObjectAttributes](docs/DataPrivacyCreateDeletionJobQueryResourceObjectAttributes.md)
- [DataPrivacyCreateDeletionJobQueryResourceObjectAttributesProfile](docs/DataPrivacyCreateDeletionJobQueryResourceObjectAttributesProfile.md)
- [DataPrivacyDeletionJobEnum](docs/DataPrivacyDeletionJobEnum.md)
- [DataPrivacyProfileQueryResourceObject](docs/DataPrivacyProfileQueryResourceObject.md)
- [DataPrivacyProfileQueryResourceObjectAttributes](docs/DataPrivacyProfileQueryResourceObjectAttributes.md)
- [DeleteTagGroupResponse](docs/DeleteTagGroupResponse.md)
- [DeviceMetadata](docs/DeviceMetadata.md)
- [EmailChannel](docs/EmailChannel.md)
- [EmailContentSubObject](docs/EmailContentSubObject.md)
- [EmailMarketing](docs/EmailMarketing.md)
- [EmailMarketingListSuppression](docs/EmailMarketingListSuppression.md)
- [EmailMarketingSuppression](docs/EmailMarketingSuppression.md)
- [EmailMessageContent](docs/EmailMessageContent.md)
- [EmailSendOptionsSubObject](docs/EmailSendOptionsSubObject.md)
- [EmailSubscriptionParameters](docs/EmailSubscriptionParameters.md)
- [EmailTrackingOptions](docs/EmailTrackingOptions.md)
- [EmailTrackingOptionsSubObject](docs/EmailTrackingOptionsSubObject.md)
- [ErrorSource](docs/ErrorSource.md)
- [EventBulkCreateEnum](docs/EventBulkCreateEnum.md)
- [EventCreateQueryV2](docs/EventCreateQueryV2.md)
- [EventCreateQueryV2ResourceObject](docs/EventCreateQueryV2ResourceObject.md)
- [EventCreateQueryV2ResourceObjectAttributes](docs/EventCreateQueryV2ResourceObjectAttributes.md)
- [EventCreateQueryV2ResourceObjectAttributesMetric](docs/EventCreateQueryV2ResourceObjectAttributesMetric.md)
- [EventCreateQueryV2ResourceObjectAttributesProfile](docs/EventCreateQueryV2ResourceObjectAttributesProfile.md)
- [EventEnum](docs/EventEnum.md)
- [EventResponseObjectResource](docs/EventResponseObjectResource.md)
- [EventResponseObjectResourceAttributes](docs/EventResponseObjectResourceAttributes.md)
- [EventsBulkCreateQuery](docs/EventsBulkCreateQuery.md)
- [EventsBulkCreateQueryResourceObject](docs/EventsBulkCreateQueryResourceObject.md)
- [EventsBulkCreateQueryResourceObjectAttributes](docs/EventsBulkCreateQueryResourceObjectAttributes.md)
- [EventsBulkCreateQueryResourceObjectAttributesEvents](docs/EventsBulkCreateQueryResourceObjectAttributesEvents.md)
- [FlowActionEnum](docs/FlowActionEnum.md)
- [FlowActionResponseObjectResource](docs/FlowActionResponseObjectResource.md)
- [FlowActionResponseObjectResourceAttributes](docs/FlowActionResponseObjectResourceAttributes.md)
- [FlowActionResponseObjectResourceAttributesTrackingOptions](docs/FlowActionResponseObjectResourceAttributesTrackingOptions.md)
- [FlowEnum](docs/FlowEnum.md)
- [FlowMessageEnum](docs/FlowMessageEnum.md)
- [FlowMessageResponseObjectResource](docs/FlowMessageResponseObjectResource.md)
- [FlowMessageResponseObjectResourceAttributes](docs/FlowMessageResponseObjectResourceAttributes.md)
- [FlowMessageResponseObjectResourceAttributesContent](docs/FlowMessageResponseObjectResourceAttributesContent.md)
- [FlowResponseObjectResource](docs/FlowResponseObjectResource.md)
- [FlowResponseObjectResourceAttributes](docs/FlowResponseObjectResourceAttributes.md)
- [FlowSeriesReportEnum](docs/FlowSeriesReportEnum.md)
- [FlowSeriesRequestDTO](docs/FlowSeriesRequestDTO.md)
- [FlowSeriesRequestDTOResourceObject](docs/FlowSeriesRequestDTOResourceObject.md)
- [FlowSeriesRequestDTOResourceObjectAttributes](docs/FlowSeriesRequestDTOResourceObjectAttributes.md)
- [FlowUpdateQuery](docs/FlowUpdateQuery.md)
- [FlowUpdateQueryResourceObject](docs/FlowUpdateQueryResourceObject.md)
- [FlowUpdateQueryResourceObjectAttributes](docs/FlowUpdateQueryResourceObjectAttributes.md)
- [FlowValuesReportEnum](docs/FlowValuesReportEnum.md)
- [FlowValuesRequestDTO](docs/FlowValuesRequestDTO.md)
- [FlowValuesRequestDTOResourceObject](docs/FlowValuesRequestDTOResourceObject.md)
- [FlowValuesRequestDTOResourceObjectAttributes](docs/FlowValuesRequestDTOResourceObjectAttributes.md)
- [GetAccountResponse](docs/GetAccountResponse.md)
- [GetAccountResponseCollection](docs/GetAccountResponseCollection.md)
- [GetAccounts4XXResponse](docs/GetAccounts4XXResponse.md)
- [GetAccounts4XXResponseErrorsInner](docs/GetAccounts4XXResponseErrorsInner.md)
- [GetAccounts4XXResponseErrorsInnerSource](docs/GetAccounts4XXResponseErrorsInnerSource.md)
- [GetCampaignMessageCampaignRelationshipListResponse](docs/GetCampaignMessageCampaignRelationshipListResponse.md)
- [GetCampaignMessageCampaignRelationshipListResponseData](docs/GetCampaignMessageCampaignRelationshipListResponseData.md)
- [GetCampaignMessageResponseCollectionCompoundDocument](docs/GetCampaignMessageResponseCollectionCompoundDocument.md)
- [GetCampaignMessageResponseCompoundDocument](docs/GetCampaignMessageResponseCompoundDocument.md)
- [GetCampaignMessageResponseCompoundDocumentData](docs/GetCampaignMessageResponseCompoundDocumentData.md)
- [GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships](docs/GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships.md)
- [GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaign](docs/GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaign.md)
- [GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData](docs/GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaignData.md)
- [GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate](docs/GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate.md)
- [GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData](docs/GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData.md)
- [GetCampaignMessageResponseCompoundDocumentIncludedInner](docs/GetCampaignMessageResponseCompoundDocumentIncludedInner.md)
- [GetCampaignMessageTemplateRelationshipListResponse](docs/GetCampaignMessageTemplateRelationshipListResponse.md)
- [GetCampaignMessagesRelationshipListResponseCollection](docs/GetCampaignMessagesRelationshipListResponseCollection.md)
- [GetCampaignMessagesRelationshipListResponseCollectionDataInner](docs/GetCampaignMessagesRelationshipListResponseCollectionDataInner.md)
- [GetCampaignRecipientEstimationJobResponse](docs/GetCampaignRecipientEstimationJobResponse.md)
- [GetCampaignRecipientEstimationResponse](docs/GetCampaignRecipientEstimationResponse.md)
- [GetCampaignResponse](docs/GetCampaignResponse.md)
- [GetCampaignResponseCollectionCompoundDocument](docs/GetCampaignResponseCollectionCompoundDocument.md)
- [GetCampaignResponseCollectionCompoundDocumentDataInner](docs/GetCampaignResponseCollectionCompoundDocumentDataInner.md)
- [GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCampaignMessages](docs/GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCampaignMessages.md)
- [GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCampaignMessagesDataInner](docs/GetCampaignResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCampaignMessagesDataInner.md)
- [GetCampaignResponseCollectionCompoundDocumentIncludedInner](docs/GetCampaignResponseCollectionCompoundDocumentIncludedInner.md)
- [GetCampaignResponseCompoundDocument](docs/GetCampaignResponseCompoundDocument.md)
- [GetCampaignResponseData](docs/GetCampaignResponseData.md)
- [GetCampaignResponseDataAllOfRelationships](docs/GetCampaignResponseDataAllOfRelationships.md)
- [GetCampaignSendJobResponse](docs/GetCampaignSendJobResponse.md)
- [GetCampaignTagRelationshipListResponseCollection](docs/GetCampaignTagRelationshipListResponseCollection.md)
- [GetCatalogCategoryCreateJobResponseCollectionCompoundDocument](docs/GetCatalogCategoryCreateJobResponseCollectionCompoundDocument.md)
- [GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner](docs/GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner.md)
- [GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategories](docs/GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategories.md)
- [GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategoriesDataInner](docs/GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategoriesDataInner.md)
- [GetCatalogCategoryCreateJobResponseCompoundDocument](docs/GetCatalogCategoryCreateJobResponseCompoundDocument.md)
- [GetCatalogCategoryDeleteJobResponse](docs/GetCatalogCategoryDeleteJobResponse.md)
- [GetCatalogCategoryDeleteJobResponseCollection](docs/GetCatalogCategoryDeleteJobResponseCollection.md)
- [GetCatalogCategoryDeleteJobResponseCollectionDataInner](docs/GetCatalogCategoryDeleteJobResponseCollectionDataInner.md)
- [GetCatalogCategoryDeleteJobResponseCollectionDataInnerAllOfRelationships](docs/GetCatalogCategoryDeleteJobResponseCollectionDataInnerAllOfRelationships.md)
- [GetCatalogCategoryItemListResponseCollection](docs/GetCatalogCategoryItemListResponseCollection.md)
- [GetCatalogCategoryItemListResponseCollectionDataInner](docs/GetCatalogCategoryItemListResponseCollectionDataInner.md)
- [GetCatalogCategoryResponse](docs/GetCatalogCategoryResponse.md)
- [GetCatalogCategoryResponseCollection](docs/GetCatalogCategoryResponseCollection.md)
- [GetCatalogCategoryResponseCollectionDataInner](docs/GetCatalogCategoryResponseCollectionDataInner.md)
- [GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships](docs/GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships.md)
- [GetCatalogCategoryUpdateJobResponseCollectionCompoundDocument](docs/GetCatalogCategoryUpdateJobResponseCollectionCompoundDocument.md)
- [GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInner](docs/GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInner.md)
- [GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategories](docs/GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategories.md)
- [GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategoriesDataInner](docs/GetCatalogCategoryUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCategoriesDataInner.md)
- [GetCatalogCategoryUpdateJobResponseCompoundDocument](docs/GetCatalogCategoryUpdateJobResponseCompoundDocument.md)
- [GetCatalogItemCategoryListResponseCollection](docs/GetCatalogItemCategoryListResponseCollection.md)
- [GetCatalogItemCategoryListResponseCollectionDataInner](docs/GetCatalogItemCategoryListResponseCollectionDataInner.md)
- [GetCatalogItemCreateJobResponseCollectionCompoundDocument](docs/GetCatalogItemCreateJobResponseCollectionCompoundDocument.md)
- [GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner](docs/GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner.md)
- [GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsItems](docs/GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsItems.md)
- [GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsItemsDataInner](docs/GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsItemsDataInner.md)
- [GetCatalogItemCreateJobResponseCompoundDocument](docs/GetCatalogItemCreateJobResponseCompoundDocument.md)
- [GetCatalogItemDeleteJobResponse](docs/GetCatalogItemDeleteJobResponse.md)
- [GetCatalogItemDeleteJobResponseCollection](docs/GetCatalogItemDeleteJobResponseCollection.md)
- [GetCatalogItemDeleteJobResponseCollectionDataInner](docs/GetCatalogItemDeleteJobResponseCollectionDataInner.md)
- [GetCatalogItemResponseCollectionCompoundDocument](docs/GetCatalogItemResponseCollectionCompoundDocument.md)
- [GetCatalogItemResponseCollectionCompoundDocumentDataInner](docs/GetCatalogItemResponseCollectionCompoundDocumentDataInner.md)
- [GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants](docs/GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants.md)
- [GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner](docs/GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner.md)
- [GetCatalogItemResponseCompoundDocument](docs/GetCatalogItemResponseCompoundDocument.md)
- [GetCatalogItemUpdateJobResponseCollectionCompoundDocument](docs/GetCatalogItemUpdateJobResponseCollectionCompoundDocument.md)
- [GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner](docs/GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner.md)
- [GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsItems](docs/GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsItems.md)
- [GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsItemsDataInner](docs/GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsItemsDataInner.md)
- [GetCatalogItemUpdateJobResponseCompoundDocument](docs/GetCatalogItemUpdateJobResponseCompoundDocument.md)
- [GetCatalogVariantCreateJobResponseCollectionCompoundDocument](docs/GetCatalogVariantCreateJobResponseCollectionCompoundDocument.md)
- [GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner](docs/GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner.md)
- [GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants](docs/GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants.md)
- [GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner](docs/GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner.md)
- [GetCatalogVariantCreateJobResponseCompoundDocument](docs/GetCatalogVariantCreateJobResponseCompoundDocument.md)
- [GetCatalogVariantDeleteJobResponse](docs/GetCatalogVariantDeleteJobResponse.md)
- [GetCatalogVariantDeleteJobResponseCollection](docs/GetCatalogVariantDeleteJobResponseCollection.md)
- [GetCatalogVariantDeleteJobResponseCollectionDataInner](docs/GetCatalogVariantDeleteJobResponseCollectionDataInner.md)
- [GetCatalogVariantDeleteJobResponseCollectionDataInnerAllOfRelationships](docs/GetCatalogVariantDeleteJobResponseCollectionDataInnerAllOfRelationships.md)
- [GetCatalogVariantResponse](docs/GetCatalogVariantResponse.md)
- [GetCatalogVariantResponseCollection](docs/GetCatalogVariantResponseCollection.md)
- [GetCatalogVariantResponseCollectionDataInner](docs/GetCatalogVariantResponseCollectionDataInner.md)
- [GetCatalogVariantResponseCollectionDataInnerAllOfRelationships](docs/GetCatalogVariantResponseCollectionDataInnerAllOfRelationships.md)
- [GetCatalogVariantUpdateJobResponseCollectionCompoundDocument](docs/GetCatalogVariantUpdateJobResponseCollectionCompoundDocument.md)
- [GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInner](docs/GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInner.md)
- [GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants](docs/GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants.md)
- [GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner](docs/GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner.md)
- [GetCatalogVariantUpdateJobResponseCompoundDocument](docs/GetCatalogVariantUpdateJobResponseCompoundDocument.md)
- [GetCouponCodeCreateJobResponseCollectionCompoundDocument](docs/GetCouponCodeCreateJobResponseCollectionCompoundDocument.md)
- [GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInner](docs/GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInner.md)
- [GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponCodes](docs/GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponCodes.md)
- [GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponCodesDataInner](docs/GetCouponCodeCreateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponCodesDataInner.md)
- [GetCouponCodeCreateJobResponseCompoundDocument](docs/GetCouponCodeCreateJobResponseCompoundDocument.md)
- [GetCouponCodeRelationshipCouponResponse](docs/GetCouponCodeRelationshipCouponResponse.md)
- [GetCouponCodeRelationshipCouponResponseData](docs/GetCouponCodeRelationshipCouponResponseData.md)
- [GetCouponCodeResponseCollection](docs/GetCouponCodeResponseCollection.md)
- [GetCouponCodeResponseCollectionCompoundDocument](docs/GetCouponCodeResponseCollectionCompoundDocument.md)
- [GetCouponCodeResponseCollectionCompoundDocumentDataInner](docs/GetCouponCodeResponseCollectionCompoundDocumentDataInner.md)
- [GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCoupon](docs/GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCoupon.md)
- [GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponData](docs/GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsCouponData.md)
- [GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile](docs/GetCouponCodeResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile.md)
- [GetCouponCodeResponseCollectionDataInner](docs/GetCouponCodeResponseCollectionDataInner.md)
- [GetCouponCodeResponseCollectionDataInnerAllOfRelationships](docs/GetCouponCodeResponseCollectionDataInnerAllOfRelationships.md)
- [GetCouponCodeResponseCompoundDocument](docs/GetCouponCodeResponseCompoundDocument.md)
- [GetCouponRelationshipCouponCodesListResponseCollection](docs/GetCouponRelationshipCouponCodesListResponseCollection.md)
- [GetCouponRelationshipCouponCodesListResponseCollectionDataInner](docs/GetCouponRelationshipCouponCodesListResponseCollectionDataInner.md)
- [GetCouponResponse](docs/GetCouponResponse.md)
- [GetCouponResponseCollection](docs/GetCouponResponseCollection.md)
- [GetEventMetricsRelationshipListResponse](docs/GetEventMetricsRelationshipListResponse.md)
- [GetEventMetricsRelationshipListResponseData](docs/GetEventMetricsRelationshipListResponseData.md)
- [GetEventProfilesRelationshipListResponse](docs/GetEventProfilesRelationshipListResponse.md)
- [GetEventProfilesRelationshipListResponseData](docs/GetEventProfilesRelationshipListResponseData.md)
- [GetEventResponseCollectionCompoundDocument](docs/GetEventResponseCollectionCompoundDocument.md)
- [GetEventResponseCollectionCompoundDocumentDataInner](docs/GetEventResponseCollectionCompoundDocumentDataInner.md)
- [GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsAttributions](docs/GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsAttributions.md)
- [GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsAttributionsDataInner](docs/GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsAttributionsDataInner.md)
- [GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsMetric](docs/GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsMetric.md)
- [GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsMetricData](docs/GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsMetricData.md)
- [GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile](docs/GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile.md)
- [GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfileData](docs/GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfileData.md)
- [GetEventResponseCollectionCompoundDocumentIncludedInner](docs/GetEventResponseCollectionCompoundDocumentIncludedInner.md)
- [GetEventResponseCompoundDocument](docs/GetEventResponseCompoundDocument.md)
- [GetFlowActionFlowMessageRelationshipResponseCollection](docs/GetFlowActionFlowMessageRelationshipResponseCollection.md)
- [GetFlowActionFlowRelationshipResponse](docs/GetFlowActionFlowRelationshipResponse.md)
- [GetFlowActionResponse](docs/GetFlowActionResponse.md)
- [GetFlowActionResponseCollection](docs/GetFlowActionResponseCollection.md)
- [GetFlowActionResponseCollectionDataInner](docs/GetFlowActionResponseCollectionDataInner.md)
- [GetFlowActionResponseCollectionDataInnerAllOfRelationships](docs/GetFlowActionResponseCollectionDataInnerAllOfRelationships.md)
- [GetFlowActionResponseCompoundDocument](docs/GetFlowActionResponseCompoundDocument.md)
- [GetFlowActionResponseCompoundDocumentData](docs/GetFlowActionResponseCompoundDocumentData.md)
- [GetFlowActionResponseCompoundDocumentDataAllOfRelationships](docs/GetFlowActionResponseCompoundDocumentDataAllOfRelationships.md)
- [GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlow](docs/GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlow.md)
- [GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowData](docs/GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowData.md)
- [GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowMessages](docs/GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowMessages.md)
- [GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowMessagesDataInner](docs/GetFlowActionResponseCompoundDocumentDataAllOfRelationshipsFlowMessagesDataInner.md)
- [GetFlowActionResponseCompoundDocumentIncludedInner](docs/GetFlowActionResponseCompoundDocumentIncludedInner.md)
- [GetFlowFlowActionRelationshipListResponseCollection](docs/GetFlowFlowActionRelationshipListResponseCollection.md)
- [GetFlowMessageFlowActionRelationshipResponse](docs/GetFlowMessageFlowActionRelationshipResponse.md)
- [GetFlowMessageResponseCollection](docs/GetFlowMessageResponseCollection.md)
- [GetFlowMessageResponseCollectionDataInner](docs/GetFlowMessageResponseCollectionDataInner.md)
- [GetFlowMessageResponseCollectionDataInnerAllOfRelationships](docs/GetFlowMessageResponseCollectionDataInnerAllOfRelationships.md)
- [GetFlowMessageResponseCompoundDocument](docs/GetFlowMessageResponseCompoundDocument.md)
- [GetFlowMessageResponseCompoundDocumentData](docs/GetFlowMessageResponseCompoundDocumentData.md)
- [GetFlowMessageResponseCompoundDocumentDataAllOfRelationships](docs/GetFlowMessageResponseCompoundDocumentDataAllOfRelationships.md)
- [GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction](docs/GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsFlowAction.md)
- [GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate](docs/GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate.md)
- [GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData](docs/GetFlowMessageResponseCompoundDocumentDataAllOfRelationshipsTemplateData.md)
- [GetFlowMessageResponseCompoundDocumentIncludedInner](docs/GetFlowMessageResponseCompoundDocumentIncludedInner.md)
- [GetFlowMessageTemplateRelationshipResponse](docs/GetFlowMessageTemplateRelationshipResponse.md)
- [GetFlowResponse](docs/GetFlowResponse.md)
- [GetFlowResponseCollectionCompoundDocument](docs/GetFlowResponseCollectionCompoundDocument.md)
- [GetFlowResponseCollectionCompoundDocumentDataInner](docs/GetFlowResponseCollectionCompoundDocumentDataInner.md)
- [GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions](docs/GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActions.md)
- [GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner](docs/GetFlowResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsFlowActionsDataInner.md)
- [GetFlowResponseCollectionCompoundDocumentIncludedInner](docs/GetFlowResponseCollectionCompoundDocumentIncludedInner.md)
- [GetFlowResponseCompoundDocument](docs/GetFlowResponseCompoundDocument.md)
- [GetFlowResponseData](docs/GetFlowResponseData.md)
- [GetFlowResponseDataAllOfRelationships](docs/GetFlowResponseDataAllOfRelationships.md)
- [GetFlowTagRelationshipListResponseCollection](docs/GetFlowTagRelationshipListResponseCollection.md)
- [GetImageResponse](docs/GetImageResponse.md)
- [GetImageResponseCollection](docs/GetImageResponseCollection.md)
- [GetImportErrorResponseCollection](docs/GetImportErrorResponseCollection.md)
- [GetListListResponseCollectionCompoundDocument](docs/GetListListResponseCollectionCompoundDocument.md)
- [GetListListResponseCollectionCompoundDocumentDataInner](docs/GetListListResponseCollectionCompoundDocumentDataInner.md)
- [GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTags](docs/GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTags.md)
- [GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagsDataInner](docs/GetListListResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagsDataInner.md)
- [GetListMemberResponseCollection](docs/GetListMemberResponseCollection.md)
- [GetListMemberResponseCollectionDataInner](docs/GetListMemberResponseCollectionDataInner.md)
- [GetListRelationshipsResponseCollection](docs/GetListRelationshipsResponseCollection.md)
- [GetListRelationshipsResponseCollectionDataInner](docs/GetListRelationshipsResponseCollectionDataInner.md)
- [GetListResponseCollection](docs/GetListResponseCollection.md)
- [GetListResponseCollectionDataInner](docs/GetListResponseCollectionDataInner.md)
- [GetListResponseCollectionDataInnerAllOfRelationships](docs/GetListResponseCollectionDataInnerAllOfRelationships.md)
- [GetListRetrieveResponseCompoundDocument](docs/GetListRetrieveResponseCompoundDocument.md)
- [GetListRetrieveResponseCompoundDocumentData](docs/GetListRetrieveResponseCompoundDocumentData.md)
- [GetListRetrieveResponseCompoundDocumentDataAllOfAttributes](docs/GetListRetrieveResponseCompoundDocumentDataAllOfAttributes.md)
- [GetListTagRelationshipListResponseCollection](docs/GetListTagRelationshipListResponseCollection.md)
- [GetMetricResponse](docs/GetMetricResponse.md)
- [GetMetricResponseCollection](docs/GetMetricResponseCollection.md)
- [GetProfileImportJobListRelationshipsResponseCollection](docs/GetProfileImportJobListRelationshipsResponseCollection.md)
- [GetProfileImportJobProfileRelationshipsResponseCollection](docs/GetProfileImportJobProfileRelationshipsResponseCollection.md)
- [GetProfileImportJobResponseCollectionCompoundDocument](docs/GetProfileImportJobResponseCollectionCompoundDocument.md)
- [GetProfileImportJobResponseCollectionCompoundDocumentDataInner](docs/GetProfileImportJobResponseCollectionCompoundDocumentDataInner.md)
- [GetProfileImportJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetProfileImportJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetProfileImportJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsLists](docs/GetProfileImportJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsLists.md)
- [GetProfileImportJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsListsDataInner](docs/GetProfileImportJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsListsDataInner.md)
- [GetProfileImportJobResponseCompoundDocument](docs/GetProfileImportJobResponseCompoundDocument.md)
- [GetProfileListRelationshipsResponseCollection](docs/GetProfileListRelationshipsResponseCollection.md)
- [GetProfileResponse](docs/GetProfileResponse.md)
- [GetProfileResponseCollection](docs/GetProfileResponseCollection.md)
- [GetProfileResponseCollectionCompoundDocument](docs/GetProfileResponseCollectionCompoundDocument.md)
- [GetProfileResponseCompoundDocument](docs/GetProfileResponseCompoundDocument.md)
- [GetProfileResponseCompoundDocumentData](docs/GetProfileResponseCompoundDocumentData.md)
- [GetProfileResponseCompoundDocumentDataAllOfRelationships](docs/GetProfileResponseCompoundDocumentDataAllOfRelationships.md)
- [GetProfileResponseCompoundDocumentDataAllOfRelationshipsLists](docs/GetProfileResponseCompoundDocumentDataAllOfRelationshipsLists.md)
- [GetProfileResponseCompoundDocumentDataAllOfRelationshipsListsDataInner](docs/GetProfileResponseCompoundDocumentDataAllOfRelationshipsListsDataInner.md)
- [GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments](docs/GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments.md)
- [GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegmentsDataInner](docs/GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegmentsDataInner.md)
- [GetProfileResponseCompoundDocumentIncludedInner](docs/GetProfileResponseCompoundDocumentIncludedInner.md)
- [GetProfileResponseData](docs/GetProfileResponseData.md)
- [GetProfileResponseDataAllOfAttributes](docs/GetProfileResponseDataAllOfAttributes.md)
- [GetProfileResponseDataAllOfRelationships](docs/GetProfileResponseDataAllOfRelationships.md)
- [GetProfileSegmentRelationshipsResponseCollection](docs/GetProfileSegmentRelationshipsResponseCollection.md)
- [GetSegmentListResponseCollectionCompoundDocument](docs/GetSegmentListResponseCollectionCompoundDocument.md)
- [GetSegmentListResponseCollectionCompoundDocumentDataInner](docs/GetSegmentListResponseCollectionCompoundDocumentDataInner.md)
- [GetSegmentMemberResponseCollection](docs/GetSegmentMemberResponseCollection.md)
- [GetSegmentMemberResponseCollectionDataInner](docs/GetSegmentMemberResponseCollectionDataInner.md)
- [GetSegmentRelationshipsResponseCollection](docs/GetSegmentRelationshipsResponseCollection.md)
- [GetSegmentResponseCollection](docs/GetSegmentResponseCollection.md)
- [GetSegmentResponseCollectionDataInner](docs/GetSegmentResponseCollectionDataInner.md)
- [GetSegmentRetrieveResponseCompoundDocument](docs/GetSegmentRetrieveResponseCompoundDocument.md)
- [GetSegmentRetrieveResponseCompoundDocumentData](docs/GetSegmentRetrieveResponseCompoundDocumentData.md)
- [GetSegmentTagRelationshipListResponseCollection](docs/GetSegmentTagRelationshipListResponseCollection.md)
- [GetTagCampaignRelationshipsResponseCollection](docs/GetTagCampaignRelationshipsResponseCollection.md)
- [GetTagCampaignRelationshipsResponseCollectionDataInner](docs/GetTagCampaignRelationshipsResponseCollectionDataInner.md)
- [GetTagFlowRelationshipsResponseCollection](docs/GetTagFlowRelationshipsResponseCollection.md)
- [GetTagFlowRelationshipsResponseCollectionDataInner](docs/GetTagFlowRelationshipsResponseCollectionDataInner.md)
- [GetTagGroupResponse](docs/GetTagGroupResponse.md)
- [GetTagGroupResponseCollection](docs/GetTagGroupResponseCollection.md)
- [GetTagGroupResponseCollectionDataInner](docs/GetTagGroupResponseCollectionDataInner.md)
- [GetTagGroupResponseCollectionDataInnerAllOfRelationships](docs/GetTagGroupResponseCollectionDataInnerAllOfRelationships.md)
- [GetTagGroupTagRelationshipsResponseCollection](docs/GetTagGroupTagRelationshipsResponseCollection.md)
- [GetTagGroupTagRelationshipsResponseCollectionDataInner](docs/GetTagGroupTagRelationshipsResponseCollectionDataInner.md)
- [GetTagListRelationshipsResponseCollection](docs/GetTagListRelationshipsResponseCollection.md)
- [GetTagListRelationshipsResponseCollectionDataInner](docs/GetTagListRelationshipsResponseCollectionDataInner.md)
- [GetTagResponseCollection](docs/GetTagResponseCollection.md)
- [GetTagResponseCollectionCompoundDocument](docs/GetTagResponseCollectionCompoundDocument.md)
- [GetTagResponseCollectionCompoundDocumentDataInner](docs/GetTagResponseCollectionCompoundDocumentDataInner.md)
- [GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationships](docs/GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md)
- [GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroup](docs/GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroup.md)
- [GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData](docs/GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData.md)
- [GetTagResponseCollectionDataInner](docs/GetTagResponseCollectionDataInner.md)
- [GetTagResponseCollectionDataInnerAllOfRelationships](docs/GetTagResponseCollectionDataInnerAllOfRelationships.md)
- [GetTagResponseCompoundDocument](docs/GetTagResponseCompoundDocument.md)
- [GetTagSegmentRelationshipsResponseCollection](docs/GetTagSegmentRelationshipsResponseCollection.md)
- [GetTagSegmentRelationshipsResponseCollectionDataInner](docs/GetTagSegmentRelationshipsResponseCollectionDataInner.md)
- [GetTagTagGroupRelationshipsResponse](docs/GetTagTagGroupRelationshipsResponse.md)
- [GetTagTagGroupRelationshipsResponseData](docs/GetTagTagGroupRelationshipsResponseData.md)
- [GetTemplateResponse](docs/GetTemplateResponse.md)
- [GetTemplateResponseCollection](docs/GetTemplateResponseCollection.md)
- [ImageCreateQuery](docs/ImageCreateQuery.md)
- [ImageCreateQueryResourceObject](docs/ImageCreateQueryResourceObject.md)
- [ImageCreateQueryResourceObjectAttributes](docs/ImageCreateQueryResourceObjectAttributes.md)
- [ImageEnum](docs/ImageEnum.md)
- [ImagePartialUpdateQuery](docs/ImagePartialUpdateQuery.md)
- [ImagePartialUpdateQueryResourceObject](docs/ImagePartialUpdateQueryResourceObject.md)
- [ImagePartialUpdateQueryResourceObjectAttributes](docs/ImagePartialUpdateQueryResourceObjectAttributes.md)
- [ImageResponseObjectResource](docs/ImageResponseObjectResource.md)
- [ImageResponseObjectResourceAttributes](docs/ImageResponseObjectResourceAttributes.md)
- [ImportErrorEnum](docs/ImportErrorEnum.md)
- [ImportErrorResponseObjectResource](docs/ImportErrorResponseObjectResource.md)
- [ImportErrorResponseObjectResourceAttributes](docs/ImportErrorResponseObjectResourceAttributes.md)
- [ListCreateQuery](docs/ListCreateQuery.md)
- [ListCreateQueryResourceObject](docs/ListCreateQueryResourceObject.md)
- [ListCreateQueryResourceObjectAttributes](docs/ListCreateQueryResourceObjectAttributes.md)
- [ListEnum](docs/ListEnum.md)
- [ListListResponseObjectResource](docs/ListListResponseObjectResource.md)
- [ListListResponseObjectResourceAttributes](docs/ListListResponseObjectResourceAttributes.md)
- [ListMemberResponseObjectResource](docs/ListMemberResponseObjectResource.md)
- [ListMemberResponseObjectResourceAttributes](docs/ListMemberResponseObjectResourceAttributes.md)
- [ListMembersAddQuery](docs/ListMembersAddQuery.md)
- [ListMembersDeleteQuery](docs/ListMembersDeleteQuery.md)
- [ListPartialUpdateQuery](docs/ListPartialUpdateQuery.md)
- [ListPartialUpdateQueryResourceObject](docs/ListPartialUpdateQueryResourceObject.md)
- [ListResponseObjectResource](docs/ListResponseObjectResource.md)
- [ListRetrieveResponseObjectResource](docs/ListRetrieveResponseObjectResource.md)
- [MarketingSubscriptionParameters](docs/MarketingSubscriptionParameters.md)
- [MetricAggregateEnum](docs/MetricAggregateEnum.md)
- [MetricAggregateQuery](docs/MetricAggregateQuery.md)
- [MetricAggregateQueryResourceObject](docs/MetricAggregateQueryResourceObject.md)
- [MetricAggregateQueryResourceObjectAttributes](docs/MetricAggregateQueryResourceObjectAttributes.md)
- [MetricAggregateRowDTO](docs/MetricAggregateRowDTO.md)
- [MetricCreateQueryResourceObject](docs/MetricCreateQueryResourceObject.md)
- [MetricCreateQueryResourceObjectAttributes](docs/MetricCreateQueryResourceObjectAttributes.md)
- [MetricEnum](docs/MetricEnum.md)
- [MetricResponseObjectResource](docs/MetricResponseObjectResource.md)
- [MetricResponseObjectResourceAttributes](docs/MetricResponseObjectResourceAttributes.md)
- [ObjectLinks](docs/ObjectLinks.md)
- [OnsiteProfileCreateQuery](docs/OnsiteProfileCreateQuery.md)
- [OnsiteProfileCreateQueryResourceObject](docs/OnsiteProfileCreateQueryResourceObject.md)
- [OnsiteProfileCreateQueryResourceObjectAttributes](docs/OnsiteProfileCreateQueryResourceObjectAttributes.md)
- [OnsiteProfileMeta](docs/OnsiteProfileMeta.md)
- [OnsiteSubscriptionCreateQuery](docs/OnsiteSubscriptionCreateQuery.md)
- [OnsiteSubscriptionCreateQueryResourceObject](docs/OnsiteSubscriptionCreateQueryResourceObject.md)
- [OnsiteSubscriptionCreateQueryResourceObjectAttributes](docs/OnsiteSubscriptionCreateQueryResourceObjectAttributes.md)
- [OnsiteSubscriptionCreateQueryResourceObjectAttributesProfile](docs/OnsiteSubscriptionCreateQueryResourceObjectAttributesProfile.md)
- [OnsiteSubscriptionCreateQueryResourceObjectRelationships](docs/OnsiteSubscriptionCreateQueryResourceObjectRelationships.md)
- [OnsiteSubscriptionCreateQueryResourceObjectRelationshipsList](docs/OnsiteSubscriptionCreateQueryResourceObjectRelationshipsList.md)
- [OnsiteSubscriptionCreateQueryResourceObjectRelationshipsListData](docs/OnsiteSubscriptionCreateQueryResourceObjectRelationshipsListData.md)
- [PatchCampaignMessageResponse](docs/PatchCampaignMessageResponse.md)
- [PatchCampaignResponse](docs/PatchCampaignResponse.md)
- [PatchCatalogCategoryResponse](docs/PatchCatalogCategoryResponse.md)
- [PatchCatalogItemResponse](docs/PatchCatalogItemResponse.md)
- [PatchCatalogVariantResponse](docs/PatchCatalogVariantResponse.md)
- [PatchCouponCodeResponse](docs/PatchCouponCodeResponse.md)
- [PatchCouponResponse](docs/PatchCouponResponse.md)
- [PatchFlowResponse](docs/PatchFlowResponse.md)
- [PatchFlowResponseData](docs/PatchFlowResponseData.md)
- [PatchImageResponse](docs/PatchImageResponse.md)
- [PatchListPartialUpdateResponse](docs/PatchListPartialUpdateResponse.md)
- [PatchProfileResponse](docs/PatchProfileResponse.md)
- [PatchSegmentPartialUpdateResponse](docs/PatchSegmentPartialUpdateResponse.md)
- [PatchSegmentPartialUpdateResponseData](docs/PatchSegmentPartialUpdateResponseData.md)
- [PatchTagGroupResponse](docs/PatchTagGroupResponse.md)
- [PatchTemplateResponse](docs/PatchTemplateResponse.md)
- [PostCampaignMessageResponse](docs/PostCampaignMessageResponse.md)
- [PostCampaignMessageResponseData](docs/PostCampaignMessageResponseData.md)
- [PostCampaignRecipientEstimationJobResponse](docs/PostCampaignRecipientEstimationJobResponse.md)
- [PostCampaignRecipientEstimationJobResponseData](docs/PostCampaignRecipientEstimationJobResponseData.md)
- [PostCampaignResponse](docs/PostCampaignResponse.md)
- [PostCampaignResponseData](docs/PostCampaignResponseData.md)
- [PostCampaignSendJobResponse](docs/PostCampaignSendJobResponse.md)
- [PostCampaignSendJobResponseData](docs/PostCampaignSendJobResponseData.md)
- [PostCampaignValuesResponseDTO](docs/PostCampaignValuesResponseDTO.md)
- [PostCampaignValuesResponseDTOData](docs/PostCampaignValuesResponseDTOData.md)
- [PostCampaignValuesResponseDTODataAttributes](docs/PostCampaignValuesResponseDTODataAttributes.md)
- [PostCampaignValuesResponseDTODataRelationships](docs/PostCampaignValuesResponseDTODataRelationships.md)
- [PostCatalogCategoryCreateJobResponse](docs/PostCatalogCategoryCreateJobResponse.md)
- [PostCatalogCategoryCreateJobResponseData](docs/PostCatalogCategoryCreateJobResponseData.md)
- [PostCatalogCategoryDeleteJobResponse](docs/PostCatalogCategoryDeleteJobResponse.md)
- [PostCatalogCategoryDeleteJobResponseData](docs/PostCatalogCategoryDeleteJobResponseData.md)
- [PostCatalogCategoryDeleteJobResponseDataRelationships](docs/PostCatalogCategoryDeleteJobResponseDataRelationships.md)
- [PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories](docs/PostCatalogCategoryDeleteJobResponseDataRelationshipsCategories.md)
- [PostCatalogCategoryDeleteJobResponseDataRelationshipsCategoriesDataInner](docs/PostCatalogCategoryDeleteJobResponseDataRelationshipsCategoriesDataInner.md)
- [PostCatalogCategoryResponse](docs/PostCatalogCategoryResponse.md)
- [PostCatalogCategoryResponseData](docs/PostCatalogCategoryResponseData.md)
- [PostCatalogCategoryResponseDataRelationships](docs/PostCatalogCategoryResponseDataRelationships.md)
- [PostCatalogCategoryResponseDataRelationshipsItems](docs/PostCatalogCategoryResponseDataRelationshipsItems.md)
- [PostCatalogCategoryUpdateJobResponse](docs/PostCatalogCategoryUpdateJobResponse.md)
- [PostCatalogCategoryUpdateJobResponseData](docs/PostCatalogCategoryUpdateJobResponseData.md)
- [PostCatalogItemCreateJobResponse](docs/PostCatalogItemCreateJobResponse.md)
- [PostCatalogItemCreateJobResponseData](docs/PostCatalogItemCreateJobResponseData.md)
- [PostCatalogItemDeleteJobResponse](docs/PostCatalogItemDeleteJobResponse.md)
- [PostCatalogItemDeleteJobResponseData](docs/PostCatalogItemDeleteJobResponseData.md)
- [PostCatalogItemDeleteJobResponseDataRelationships](docs/PostCatalogItemDeleteJobResponseDataRelationships.md)
- [PostCatalogItemDeleteJobResponseDataRelationshipsItems](docs/PostCatalogItemDeleteJobResponseDataRelationshipsItems.md)
- [PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner](docs/PostCatalogItemDeleteJobResponseDataRelationshipsItemsDataInner.md)
- [PostCatalogItemResponse](docs/PostCatalogItemResponse.md)
- [PostCatalogItemResponseData](docs/PostCatalogItemResponseData.md)
- [PostCatalogItemUpdateJobResponse](docs/PostCatalogItemUpdateJobResponse.md)
- [PostCatalogItemUpdateJobResponseData](docs/PostCatalogItemUpdateJobResponseData.md)
- [PostCatalogVariantCreateJobResponse](docs/PostCatalogVariantCreateJobResponse.md)
- [PostCatalogVariantCreateJobResponseData](docs/PostCatalogVariantCreateJobResponseData.md)
- [PostCatalogVariantDeleteJobResponse](docs/PostCatalogVariantDeleteJobResponse.md)
- [PostCatalogVariantDeleteJobResponseData](docs/PostCatalogVariantDeleteJobResponseData.md)
- [PostCatalogVariantDeleteJobResponseDataRelationships](docs/PostCatalogVariantDeleteJobResponseDataRelationships.md)
- [PostCatalogVariantDeleteJobResponseDataRelationshipsVariants](docs/PostCatalogVariantDeleteJobResponseDataRelationshipsVariants.md)
- [PostCatalogVariantDeleteJobResponseDataRelationshipsVariantsDataInner](docs/PostCatalogVariantDeleteJobResponseDataRelationshipsVariantsDataInner.md)
- [PostCatalogVariantResponse](docs/PostCatalogVariantResponse.md)
- [PostCatalogVariantResponseData](docs/PostCatalogVariantResponseData.md)
- [PostCatalogVariantResponseDataRelationships](docs/PostCatalogVariantResponseDataRelationships.md)
- [PostCatalogVariantResponseDataRelationshipsItem](docs/PostCatalogVariantResponseDataRelationshipsItem.md)
- [PostCatalogVariantResponseDataRelationshipsItemData](docs/PostCatalogVariantResponseDataRelationshipsItemData.md)
- [PostCatalogVariantUpdateJobResponse](docs/PostCatalogVariantUpdateJobResponse.md)
- [PostCatalogVariantUpdateJobResponseData](docs/PostCatalogVariantUpdateJobResponseData.md)
- [PostCouponCodeCreateJobResponse](docs/PostCouponCodeCreateJobResponse.md)
- [PostCouponCodeCreateJobResponseData](docs/PostCouponCodeCreateJobResponseData.md)
- [PostCouponCodeResponse](docs/PostCouponCodeResponse.md)
- [PostCouponCodeResponseData](docs/PostCouponCodeResponseData.md)
- [PostCouponCodeResponseDataRelationships](docs/PostCouponCodeResponseDataRelationships.md)
- [PostCouponCodeResponseDataRelationshipsProfile](docs/PostCouponCodeResponseDataRelationshipsProfile.md)
- [PostCouponCodeResponseDataRelationshipsProfileData](docs/PostCouponCodeResponseDataRelationshipsProfileData.md)
- [PostCouponResponse](docs/PostCouponResponse.md)
- [PostCouponResponseData](docs/PostCouponResponseData.md)
- [PostFlowSeriesResponseDTO](docs/PostFlowSeriesResponseDTO.md)
- [PostFlowSeriesResponseDTOData](docs/PostFlowSeriesResponseDTOData.md)
- [PostFlowSeriesResponseDTODataAttributes](docs/PostFlowSeriesResponseDTODataAttributes.md)
- [PostFlowValuesResponseDTO](docs/PostFlowValuesResponseDTO.md)
- [PostFlowValuesResponseDTOData](docs/PostFlowValuesResponseDTOData.md)
- [PostFlowValuesResponseDTODataAttributes](docs/PostFlowValuesResponseDTODataAttributes.md)
- [PostFlowValuesResponseDTODataRelationships](docs/PostFlowValuesResponseDTODataRelationships.md)
- [PostImageResponse](docs/PostImageResponse.md)
- [PostImageResponseData](docs/PostImageResponseData.md)
- [PostListCreateResponse](docs/PostListCreateResponse.md)
- [PostListCreateResponseData](docs/PostListCreateResponseData.md)
- [PostListCreateResponseDataRelationships](docs/PostListCreateResponseDataRelationships.md)
- [PostListCreateResponseDataRelationshipsProfiles](docs/PostListCreateResponseDataRelationshipsProfiles.md)
- [PostMetricAggregateResponse](docs/PostMetricAggregateResponse.md)
- [PostMetricAggregateResponseData](docs/PostMetricAggregateResponseData.md)
- [PostMetricAggregateResponseDataAttributes](docs/PostMetricAggregateResponseDataAttributes.md)
- [PostProfileImportJobResponse](docs/PostProfileImportJobResponse.md)
- [PostProfileImportJobResponseData](docs/PostProfileImportJobResponseData.md)
- [PostProfileImportJobResponseDataRelationships](docs/PostProfileImportJobResponseDataRelationships.md)
- [PostProfileImportJobResponseDataRelationshipsImportErrors](docs/PostProfileImportJobResponseDataRelationshipsImportErrors.md)
- [PostProfileImportJobResponseDataRelationshipsImportErrorsDataInner](docs/PostProfileImportJobResponseDataRelationshipsImportErrorsDataInner.md)
- [PostProfileImportJobResponseDataRelationshipsProfiles](docs/PostProfileImportJobResponseDataRelationshipsProfiles.md)
- [PostProfileImportJobResponseDataRelationshipsProfilesDataInner](docs/PostProfileImportJobResponseDataRelationshipsProfilesDataInner.md)
- [PostProfileMergeResponse](docs/PostProfileMergeResponse.md)
- [PostProfileMergeResponseData](docs/PostProfileMergeResponseData.md)
- [PostProfileResponse](docs/PostProfileResponse.md)
- [PostProfileResponseData](docs/PostProfileResponseData.md)
- [PostProfileResponseDataAttributes](docs/PostProfileResponseDataAttributes.md)
- [PostTagGroupResponse](docs/PostTagGroupResponse.md)
- [PostTagGroupResponseData](docs/PostTagGroupResponseData.md)
- [PostTagGroupResponseDataRelationships](docs/PostTagGroupResponseDataRelationships.md)
- [PostTagResponse](docs/PostTagResponse.md)
- [PostTagResponseData](docs/PostTagResponseData.md)
- [PostTagResponseDataRelationships](docs/PostTagResponseDataRelationships.md)
- [PostTagResponseDataRelationshipsCampaigns](docs/PostTagResponseDataRelationshipsCampaigns.md)
- [PostTagResponseDataRelationshipsFlows](docs/PostTagResponseDataRelationshipsFlows.md)
- [PostTemplateResponse](docs/PostTemplateResponse.md)
- [PostTemplateResponseData](docs/PostTemplateResponseData.md)
- [PredictiveAnalytics](docs/PredictiveAnalytics.md)
- [ProfileBulkImportJobEnum](docs/ProfileBulkImportJobEnum.md)
- [ProfileCreateQuery](docs/ProfileCreateQuery.md)
- [ProfileCreateQueryResourceObject](docs/ProfileCreateQueryResourceObject.md)
- [ProfileCreateQueryResourceObjectAttributes](docs/ProfileCreateQueryResourceObjectAttributes.md)
- [ProfileEnum](docs/ProfileEnum.md)
- [ProfileIdentifierDTOResourceObject](docs/ProfileIdentifierDTOResourceObject.md)
- [ProfileIdentifierDTOResourceObjectAttributes](docs/ProfileIdentifierDTOResourceObjectAttributes.md)
- [ProfileImportJobCreateQuery](docs/ProfileImportJobCreateQuery.md)
- [ProfileImportJobCreateQueryResourceObject](docs/ProfileImportJobCreateQueryResourceObject.md)
- [ProfileImportJobCreateQueryResourceObjectAttributes](docs/ProfileImportJobCreateQueryResourceObjectAttributes.md)
- [ProfileImportJobCreateQueryResourceObjectAttributesProfiles](docs/ProfileImportJobCreateQueryResourceObjectAttributesProfiles.md)
- [ProfileImportJobCreateQueryResourceObjectRelationships](docs/ProfileImportJobCreateQueryResourceObjectRelationships.md)
- [ProfileImportJobCreateQueryResourceObjectRelationshipsLists](docs/ProfileImportJobCreateQueryResourceObjectRelationshipsLists.md)
- [ProfileImportJobCreateQueryResourceObjectRelationshipsListsDataInner](docs/ProfileImportJobCreateQueryResourceObjectRelationshipsListsDataInner.md)
- [ProfileImportJobResponseObjectResource](docs/ProfileImportJobResponseObjectResource.md)
- [ProfileImportJobResponseObjectResourceAttributes](docs/ProfileImportJobResponseObjectResourceAttributes.md)
- [ProfileLocation](docs/ProfileLocation.md)
- [ProfileLocationLatitude](docs/ProfileLocationLatitude.md)
- [ProfileLocationLongitude](docs/ProfileLocationLongitude.md)
- [ProfileMergeEnum](docs/ProfileMergeEnum.md)
- [ProfileMergeQuery](docs/ProfileMergeQuery.md)
- [ProfileMergeQueryResourceObject](docs/ProfileMergeQueryResourceObject.md)
- [ProfileMergeQueryResourceObjectRelationships](docs/ProfileMergeQueryResourceObjectRelationships.md)
- [ProfileMergeQueryResourceObjectRelationshipsProfiles](docs/ProfileMergeQueryResourceObjectRelationshipsProfiles.md)
- [ProfileMergeQueryResourceObjectRelationshipsProfilesDataInner](docs/ProfileMergeQueryResourceObjectRelationshipsProfilesDataInner.md)
- [ProfileMeta](docs/ProfileMeta.md)
- [ProfileMetaPatchProperties](docs/ProfileMetaPatchProperties.md)
- [ProfileMetaPatchPropertiesUnset](docs/ProfileMetaPatchPropertiesUnset.md)
- [ProfilePartialUpdateQuery](docs/ProfilePartialUpdateQuery.md)
- [ProfilePartialUpdateQueryResourceObject](docs/ProfilePartialUpdateQueryResourceObject.md)
- [ProfilePartialUpdateQueryResourceObjectAttributes](docs/ProfilePartialUpdateQueryResourceObjectAttributes.md)
- [ProfileResponseObjectResource](docs/ProfileResponseObjectResource.md)
- [ProfileResponseObjectResourceAttributes](docs/ProfileResponseObjectResourceAttributes.md)
- [ProfileSubscriptionBulkCreateJobEnum](docs/ProfileSubscriptionBulkCreateJobEnum.md)
- [ProfileSubscriptionBulkDeleteJobEnum](docs/ProfileSubscriptionBulkDeleteJobEnum.md)
- [ProfileSubscriptionCreateQueryResourceObject](docs/ProfileSubscriptionCreateQueryResourceObject.md)
- [ProfileSubscriptionCreateQueryResourceObjectAttributes](docs/ProfileSubscriptionCreateQueryResourceObjectAttributes.md)
- [ProfileSubscriptionDeleteQueryResourceObject](docs/ProfileSubscriptionDeleteQueryResourceObject.md)
- [ProfileSubscriptionDeleteQueryResourceObjectAttributes](docs/ProfileSubscriptionDeleteQueryResourceObjectAttributes.md)
- [ProfileSuppressionBulkCreateJobEnum](docs/ProfileSuppressionBulkCreateJobEnum.md)
- [ProfileSuppressionBulkDeleteJobEnum](docs/ProfileSuppressionBulkDeleteJobEnum.md)
- [ProfileSuppressionCreateQueryResourceObject](docs/ProfileSuppressionCreateQueryResourceObject.md)
- [ProfileSuppressionCreateQueryResourceObjectAttributes](docs/ProfileSuppressionCreateQueryResourceObjectAttributes.md)
- [ProfileSuppressionDeleteQueryResourceObject](docs/ProfileSuppressionDeleteQueryResourceObject.md)
- [ProfileSuppressionDeleteQueryResourceObjectAttributes](docs/ProfileSuppressionDeleteQueryResourceObjectAttributes.md)
- [ProfileUpsertQuery](docs/ProfileUpsertQuery.md)
- [ProfileUpsertQueryResourceObject](docs/ProfileUpsertQueryResourceObject.md)
- [ProfileUpsertQueryResourceObjectAttributes](docs/ProfileUpsertQueryResourceObjectAttributes.md)
- [PushTokenCreateQuery](docs/PushTokenCreateQuery.md)
- [PushTokenCreateQueryResourceObject](docs/PushTokenCreateQueryResourceObject.md)
- [PushTokenCreateQueryResourceObjectAttributes](docs/PushTokenCreateQueryResourceObjectAttributes.md)
- [PushTokenCreateQueryResourceObjectAttributesProfile](docs/PushTokenCreateQueryResourceObjectAttributesProfile.md)
- [PushTokenEnum](docs/PushTokenEnum.md)
- [PushTokenUnregisterEnum](docs/PushTokenUnregisterEnum.md)
- [PushTokenUnregisterQuery](docs/PushTokenUnregisterQuery.md)
- [PushTokenUnregisterQueryResourceObject](docs/PushTokenUnregisterQueryResourceObject.md)
- [PushTokenUnregisterQueryResourceObjectAttributes](docs/PushTokenUnregisterQueryResourceObjectAttributes.md)
- [RelationshipLinks](docs/RelationshipLinks.md)
- [RenderOptionsSubObject](docs/RenderOptionsSubObject.md)
- [SMSChannel](docs/SMSChannel.md)
- [SMSContentSubObject](docs/SMSContentSubObject.md)
- [SMSContentSubObjectCreate](docs/SMSContentSubObjectCreate.md)
- [SMSMarketing](docs/SMSMarketing.md)
- [SMSMessageContent](docs/SMSMessageContent.md)
- [SMSRenderOptions](docs/SMSRenderOptions.md)
- [SMSSendOptionsSubObject](docs/SMSSendOptionsSubObject.md)
- [SMSSubscriptionParameters](docs/SMSSubscriptionParameters.md)
- [SMSTrackingOptions](docs/SMSTrackingOptions.md)
- [SMSTrackingOptionsSubObject](docs/SMSTrackingOptionsSubObject.md)
- [STOScheduleOptions](docs/STOScheduleOptions.md)
- [SegmentEnum](docs/SegmentEnum.md)
- [SegmentListResponseObjectResource](docs/SegmentListResponseObjectResource.md)
- [SegmentListResponseObjectResourceAttributes](docs/SegmentListResponseObjectResourceAttributes.md)
- [SegmentMemberResponseObjectResource](docs/SegmentMemberResponseObjectResource.md)
- [SegmentMemberResponseObjectResourceAttributes](docs/SegmentMemberResponseObjectResourceAttributes.md)
- [SegmentPartialUpdateQuery](docs/SegmentPartialUpdateQuery.md)
- [SegmentPartialUpdateQueryResourceObject](docs/SegmentPartialUpdateQueryResourceObject.md)
- [SegmentPartialUpdateQueryResourceObjectAttributes](docs/SegmentPartialUpdateQueryResourceObjectAttributes.md)
- [SegmentResponseObjectResource](docs/SegmentResponseObjectResource.md)
- [SegmentRetrieveResponseObjectResource](docs/SegmentRetrieveResponseObjectResource.md)
- [SendOptions](docs/SendOptions.md)
- [SendStrategySubObject](docs/SendStrategySubObject.md)
- [SendTimeSubObject](docs/SendTimeSubObject.md)
- [SeriesData](docs/SeriesData.md)
- [ServerBISSubscriptionCreateQuery](docs/ServerBISSubscriptionCreateQuery.md)
- [ServerBISSubscriptionCreateQueryResourceObject](docs/ServerBISSubscriptionCreateQueryResourceObject.md)
- [ServerBISSubscriptionCreateQueryResourceObjectAttributes](docs/ServerBISSubscriptionCreateQueryResourceObjectAttributes.md)
- [ServerBISSubscriptionCreateQueryResourceObjectAttributesProfile](docs/ServerBISSubscriptionCreateQueryResourceObjectAttributesProfile.md)
- [ServerBISSubscriptionCreateQueryResourceObjectRelationships](docs/ServerBISSubscriptionCreateQueryResourceObjectRelationships.md)
- [ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariant](docs/ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariant.md)
- [ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData](docs/ServerBISSubscriptionCreateQueryResourceObjectRelationshipsVariantData.md)
- [StaticScheduleOptions](docs/StaticScheduleOptions.md)
- [StreetAddress](docs/StreetAddress.md)
- [SubscriptionChannels](docs/SubscriptionChannels.md)
- [SubscriptionCreateJobCreateQuery](docs/SubscriptionCreateJobCreateQuery.md)
- [SubscriptionCreateJobCreateQueryResourceObject](docs/SubscriptionCreateJobCreateQueryResourceObject.md)
- [SubscriptionCreateJobCreateQueryResourceObjectAttributes](docs/SubscriptionCreateJobCreateQueryResourceObjectAttributes.md)
- [SubscriptionCreateJobCreateQueryResourceObjectAttributesProfiles](docs/SubscriptionCreateJobCreateQueryResourceObjectAttributesProfiles.md)
- [SubscriptionCreateJobCreateQueryResourceObjectRelationships](docs/SubscriptionCreateJobCreateQueryResourceObjectRelationships.md)
- [SubscriptionCreateJobCreateQueryResourceObjectRelationshipsList](docs/SubscriptionCreateJobCreateQueryResourceObjectRelationshipsList.md)
- [SubscriptionCreateJobCreateQueryResourceObjectRelationshipsListData](docs/SubscriptionCreateJobCreateQueryResourceObjectRelationshipsListData.md)
- [SubscriptionDeleteJobCreateQuery](docs/SubscriptionDeleteJobCreateQuery.md)
- [SubscriptionDeleteJobCreateQueryResourceObject](docs/SubscriptionDeleteJobCreateQueryResourceObject.md)
- [SubscriptionDeleteJobCreateQueryResourceObjectAttributes](docs/SubscriptionDeleteJobCreateQueryResourceObjectAttributes.md)
- [SubscriptionDeleteJobCreateQueryResourceObjectAttributesProfiles](docs/SubscriptionDeleteJobCreateQueryResourceObjectAttributesProfiles.md)
- [SubscriptionDeleteJobCreateQueryResourceObjectRelationships](docs/SubscriptionDeleteJobCreateQueryResourceObjectRelationships.md)
- [SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList](docs/SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsList.md)
- [SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsListData](docs/SubscriptionDeleteJobCreateQueryResourceObjectRelationshipsListData.md)
- [SubscriptionEnum](docs/SubscriptionEnum.md)
- [Subscriptions](docs/Subscriptions.md)
- [SuppressionCreateJobCreateQuery](docs/SuppressionCreateJobCreateQuery.md)
- [SuppressionCreateJobCreateQueryResourceObject](docs/SuppressionCreateJobCreateQueryResourceObject.md)
- [SuppressionCreateJobCreateQueryResourceObjectAttributes](docs/SuppressionCreateJobCreateQueryResourceObjectAttributes.md)
- [SuppressionCreateJobCreateQueryResourceObjectAttributesProfiles](docs/SuppressionCreateJobCreateQueryResourceObjectAttributesProfiles.md)
- [SuppressionDeleteJobCreateQuery](docs/SuppressionDeleteJobCreateQuery.md)
- [SuppressionDeleteJobCreateQueryResourceObject](docs/SuppressionDeleteJobCreateQueryResourceObject.md)
- [SuppressionDeleteJobCreateQueryResourceObjectAttributes](docs/SuppressionDeleteJobCreateQueryResourceObjectAttributes.md)
- [SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles](docs/SuppressionDeleteJobCreateQueryResourceObjectAttributesProfiles.md)
- [TagCampaignOp](docs/TagCampaignOp.md)
- [TagCampaignOpDataInner](docs/TagCampaignOpDataInner.md)
- [TagCreateQuery](docs/TagCreateQuery.md)
- [TagCreateQueryResourceObject](docs/TagCreateQueryResourceObject.md)
- [TagCreateQueryResourceObjectRelationships](docs/TagCreateQueryResourceObjectRelationships.md)
- [TagCreateQueryResourceObjectRelationshipsTagGroup](docs/TagCreateQueryResourceObjectRelationshipsTagGroup.md)
- [TagCreateQueryResourceObjectRelationshipsTagGroupData](docs/TagCreateQueryResourceObjectRelationshipsTagGroupData.md)
- [TagEnum](docs/TagEnum.md)
- [TagFlowOp](docs/TagFlowOp.md)
- [TagFlowOpDataInner](docs/TagFlowOpDataInner.md)
- [TagGroupCreateQuery](docs/TagGroupCreateQuery.md)
- [TagGroupCreateQueryResourceObject](docs/TagGroupCreateQueryResourceObject.md)
- [TagGroupCreateQueryResourceObjectAttributes](docs/TagGroupCreateQueryResourceObjectAttributes.md)
- [TagGroupEnum](docs/TagGroupEnum.md)
- [TagGroupResponseObjectResource](docs/TagGroupResponseObjectResource.md)
- [TagGroupResponseObjectResourceAttributes](docs/TagGroupResponseObjectResourceAttributes.md)
- [TagGroupUpdateQuery](docs/TagGroupUpdateQuery.md)
- [TagGroupUpdateQueryResourceObject](docs/TagGroupUpdateQueryResourceObject.md)
- [TagGroupUpdateQueryResourceObjectAttributes](docs/TagGroupUpdateQueryResourceObjectAttributes.md)
- [TagListOp](docs/TagListOp.md)
- [TagListOpDataInner](docs/TagListOpDataInner.md)
- [TagResponseObjectResource](docs/TagResponseObjectResource.md)
- [TagResponseObjectResourceAttributes](docs/TagResponseObjectResourceAttributes.md)
- [TagSegmentOp](docs/TagSegmentOp.md)
- [TagSegmentOpDataInner](docs/TagSegmentOpDataInner.md)
- [TagUpdateQuery](docs/TagUpdateQuery.md)
- [TagUpdateQueryResourceObject](docs/TagUpdateQueryResourceObject.md)
- [TemplateCloneQuery](docs/TemplateCloneQuery.md)
- [TemplateCloneQueryResourceObject](docs/TemplateCloneQueryResourceObject.md)
- [TemplateCloneQueryResourceObjectAttributes](docs/TemplateCloneQueryResourceObjectAttributes.md)
- [TemplateCreateQuery](docs/TemplateCreateQuery.md)
- [TemplateCreateQueryResourceObject](docs/TemplateCreateQueryResourceObject.md)
- [TemplateCreateQueryResourceObjectAttributes](docs/TemplateCreateQueryResourceObjectAttributes.md)
- [TemplateEnum](docs/TemplateEnum.md)
- [TemplateRenderQuery](docs/TemplateRenderQuery.md)
- [TemplateRenderQueryResourceObject](docs/TemplateRenderQueryResourceObject.md)
- [TemplateRenderQueryResourceObjectAttributes](docs/TemplateRenderQueryResourceObjectAttributes.md)
- [TemplateResponseObjectResource](docs/TemplateResponseObjectResource.md)
- [TemplateResponseObjectResourceAttributes](docs/TemplateResponseObjectResourceAttributes.md)
- [TemplateUpdateQuery](docs/TemplateUpdateQuery.md)
- [TemplateUpdateQueryResourceObject](docs/TemplateUpdateQueryResourceObject.md)
- [TemplateUpdateQueryResourceObjectAttributes](docs/TemplateUpdateQueryResourceObjectAttributes.md)
- [ThrottledScheduleOptions](docs/ThrottledScheduleOptions.md)
- [Timeframe](docs/Timeframe.md)
- [UTMParamsSubObject](docs/UTMParamsSubObject.md)
- [UtmParamInfo](docs/UtmParamInfo.md)
- [ValuesData](docs/ValuesData.md)

## Documentation For Authorization

Authentication schemes defined for the API:

### Klaviyo-API-Key

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
  context.Background(),
  klaviyo.ContextAPIKeys,
  map[string]klaviyo.APIKey{
   "Authorization": {Key: "API_KEY_STRING"},
  },
 )
r, err := client.Service.Operation(auth, args)
```

## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

- `PtrBool`
- `PtrInt`
- `PtrInt32`
- `PtrInt64`
- `PtrFloat`
- `PtrFloat32`
- `PtrFloat64`
- `PtrString`
- `PtrTime`

## Author

developers@klaviyo.com
