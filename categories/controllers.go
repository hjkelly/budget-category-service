package categories

import (
	"cloud.google.com/go/datastore"
	"context"
)

type Cat struct {
	Value string
}

const KIND = "categories"

func getDatastoreClient(ctx context.Context) (*datastore.Client, error) {
	return datastore.NewClient(ctx, "zero-balance-budget")
}

func Create(cat Category) (Category, error) {
	ctx := context.Background()

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	dsClient, err := getDatastoreClient(ctx)
	if err != nil {
		panic(err)
	}

	k := datastore.NameKey(KIND, cat.ID, nil)
	//c := new(Cat)
	//if err := dsClient.Get(ctx, k, c); err != nil {

	if _, err := dsClient.Put(ctx, k, &cat); err != nil {
		// Handle error.
		panic(err)
	}

	return cat, nil
}

func List() ([]Category, error) {
	ctx := context.Background()

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	dsClient, err := getDatastoreClient(ctx)
	if err != nil {
		panic(err)
	}

	results := make([]Category, 0, 0)
	query := datastore.NewQuery(KIND)
	//c := new(Cat)
	//if err := dsClient.Get(ctx, k, c); err != nil {

	if _, err := dsClient.GetAll(ctx, query, &results); err != nil {
		// Handle error.
		panic(err)
	}

	return results, nil
}
