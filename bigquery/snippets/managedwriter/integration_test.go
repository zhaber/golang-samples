// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package managedwriter

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/golang-samples/bigquery/snippets/bqtestutil"
)

func BenchmarkBQWriteAPI(b *testing.B) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, getProjectId())
	if err != nil {
		b.Fatal(err)
	}
	defer client.Close()
	meta := &bigquery.DatasetMetadata{
		Location: "us-central1", // See https://cloud.google.com/bigquery/docs/locations
	}
	testDatasetID, err := bqtestutil.UniqueBQName("snippet_managedwriter_tests")
	if err != nil {
		b.Fatalf("couldn't generate unique resource name: %v", err)
	}
	if err := client.Dataset(testDatasetID).Create(ctx, meta); err != nil {
		b.Fatalf("failed to create test dataset: %v", err)
	}
	defer client.Dataset(testDatasetID).DeleteWithContents(ctx) // Cleanup table at end of test.

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pendingSchemaSmall := bigquery.Schema{
				{Name: "bool_col", Type: bigquery.BooleanFieldType},
				{Name: "float64_col", Type: bigquery.FloatFieldType},
				{Name: "int64_col", Type: bigquery.IntegerFieldType},
				{Name: "string_col", Type: bigquery.StringFieldType},
			}

			testTableIDSmall, err := bqtestutil.UniqueBQName("small")
			if err != nil {
				b.Fatalf("couldn't generate unique table id: %v", err)
			}

			if err := client.Dataset(testDatasetID).Table(testTableIDSmall).Create(ctx, &bigquery.TableMetadata{
				Schema: pendingSchemaSmall,
			}); err != nil {
				b.Fatalf("failed to create destination table(%q %q): %v", testDatasetID, testTableIDSmall, err)
			}
			appendToStreamJson(ioutil.Discard, getProjectId(), testDatasetID, testTableIDSmall)
		}
	})
}

func getProjectId() string {
	return os.Getenv("GOLANG_SAMPLES_PROJECT_ID")
}
