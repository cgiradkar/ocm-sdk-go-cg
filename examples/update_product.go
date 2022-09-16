/*
Copyright (c) 2021 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This example shows how to update the display name of a product.

package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/renan-campos/ocm-sdk-go"
	"github.com/renan-campos/ocm-sdk-go/logging"
	sb "github.com/renan-campos/ocm-sdk-go/statusboard/v1"
)

func main() {
	// Create a context:
	ctx := context.Background()

	// Create a logger that has the debug level enabled:
	logger, err := logging.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}

	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		URL("http://localhost:8000").
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of products:
	collection := connection.StatusBoard().V1().Products()

	// Get the client for the resource that manages the product that we want to update. Note
	// that this will not yet send any request to the server, so it will succeed even if the
	// product doesn't exist.
	resource := collection.Product("ea7ee64f-978d-4705-a271-85b072bc5241")

	// Prepare the patch to send:
	patch, err := sb.NewProduct().
		Name("SomeProduct").
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create product patch: %v\n", err)
		os.Exit(1)
	}

	// Send the request to update the product:
	_, err = resource.Update().
		Body(patch).
		SendContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't update product: %v\n", err)
		os.Exit(1)
	}
}
