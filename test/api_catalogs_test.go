/*
Klaviyo API

Testing CatalogsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package klaviyo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/developertom01/klaviyo-go-sdk"
)

func Test_klaviyo_CatalogsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CatalogsAPIService CreateBackInStockSubscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CatalogsAPI.CreateBackInStockSubscription(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService CreateCatalogCategory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.CreateCatalogCategory(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService CreateCatalogCategoryRelationshipsItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CatalogsAPI.CreateCatalogCategoryRelationshipsItems(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService CreateCatalogItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.CreateCatalogItem(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService CreateCatalogItemRelationshipsCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CatalogsAPI.CreateCatalogItemRelationshipsCategories(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService CreateCatalogVariant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.CreateCatalogVariant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService DeleteCatalogCategory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CatalogsAPI.DeleteCatalogCategory(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService DeleteCatalogCategoryRelationshipsItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CatalogsAPI.DeleteCatalogCategoryRelationshipsItems(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService DeleteCatalogItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CatalogsAPI.DeleteCatalogItem(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService DeleteCatalogItemRelationshipsCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CatalogsAPI.DeleteCatalogItemRelationshipsCategories(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService DeleteCatalogVariant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CatalogsAPI.DeleteCatalogVariant(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogCategories(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogCategory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogCategory(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogCategoryItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogCategoryItems(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogCategoryRelationshipsItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogCategoryRelationshipsItems(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogItem(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogItemCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogItemCategories(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogItemRelationshipsCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogItemRelationshipsCategories(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogItemVariants", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogItemVariants(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogVariant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogVariant(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCatalogVariants", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetCatalogVariants(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCreateCategoriesJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCreateCategoriesJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCreateCategoriesJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetCreateCategoriesJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCreateItemsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCreateItemsJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCreateItemsJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetCreateItemsJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCreateVariantsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.CatalogsAPI.GetCreateVariantsJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetCreateVariantsJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetCreateVariantsJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetDeleteCategoriesJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.CatalogsAPI.GetDeleteCategoriesJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetDeleteCategoriesJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetDeleteCategoriesJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetDeleteItemsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.CatalogsAPI.GetDeleteItemsJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetDeleteItemsJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetDeleteItemsJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetDeleteVariantsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.CatalogsAPI.GetDeleteVariantsJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetDeleteVariantsJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetDeleteVariantsJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetUpdateCategoriesJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.CatalogsAPI.GetUpdateCategoriesJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetUpdateCategoriesJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetUpdateCategoriesJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetUpdateItemsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.CatalogsAPI.GetUpdateItemsJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetUpdateItemsJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetUpdateItemsJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetUpdateVariantsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.CatalogsAPI.GetUpdateVariantsJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService GetUpdateVariantsJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.GetUpdateVariantsJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService SpawnCreateCategoriesJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.SpawnCreateCategoriesJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService SpawnCreateItemsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.SpawnCreateItemsJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService SpawnCreateVariantsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.SpawnCreateVariantsJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService SpawnDeleteCategoriesJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.SpawnDeleteCategoriesJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService SpawnDeleteItemsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.SpawnDeleteItemsJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService SpawnDeleteVariantsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.SpawnDeleteVariantsJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService SpawnUpdateCategoriesJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.SpawnUpdateCategoriesJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService SpawnUpdateItemsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.SpawnUpdateItemsJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService SpawnUpdateVariantsJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogsAPI.SpawnUpdateVariantsJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService UpdateCatalogCategory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.UpdateCatalogCategory(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService UpdateCatalogCategoryRelationshipsItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CatalogsAPI.UpdateCatalogCategoryRelationshipsItems(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService UpdateCatalogItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.UpdateCatalogItem(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService UpdateCatalogItemRelationshipsCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.CatalogsAPI.UpdateCatalogItemRelationshipsCategories(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogsAPIService UpdateCatalogVariant", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CatalogsAPI.UpdateCatalogVariant(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
