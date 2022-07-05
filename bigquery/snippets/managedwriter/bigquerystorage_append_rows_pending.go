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
	"fmt"
	"io"
	"log"
	"math/rand"

	"cloud.google.com/go/bigquery/storage/managedwriter"
	"cloud.google.com/go/bigquery/storage/managedwriter/adapt"
	"github.com/GoogleCloudPlatform/golang-samples/bigquery/snippets/managedwriter/exampleproto"
	"github.com/google/uuid"
	storagepb "google.golang.org/genproto/googleapis/cloud/bigquery/storage/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func jsonToProtobuf(json string, pb proto.Message) {
	// Unmarshal accept []byte, so we'll convert json from string to []byte
	err := protojson.Unmarshal([]byte(json), pb)
	if err != nil {
		log.Println("Can't convert from JSON to Protobuf", err)
	}
}

func generateJsonMessagesFromString(numMessages int) ([][]byte, error) {
	msgs := make([][]byte, numMessages)
	for i := 0; i < numMessages; i++ {
		json := `{"bool_col": true,` +
			`"float64_col": ` + fmt.Sprintf("%f", rand.Float64()) + `,` +
			`"int64_col": ` + fmt.Sprint(rand.Int63()) + `,` +
			`"string_col": "` + uuid.New().String() + `"}`
		m := &exampleproto.JsonData{}
		jsonToProtobuf(json, m)
		b, err := proto.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("error generating message %d: %v", i, err)
		}
		msgs[i] = b
	}

	return msgs, nil
}

func generateJsonMessages(numMessages int) ([][]byte, error) {
	msgs := make([][]byte, numMessages)
	for i := 0; i < numMessages; i++ {
		m := &exampleproto.JsonData{
			BoolCol:    proto.Bool(true),
			Float64Col: proto.Float64(rand.Float64()),
			Int64Col:   proto.Int64(rand.Int63()),
			StringCol:  proto.String(uuid.New().String()),
		}
		b, err := proto.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("error generating message %d: %v", i, err)
		}
		msgs[i] = b
	}

	return msgs, nil
}

func appendToStreamJson(w io.Writer, projectID, datasetID, tableID string) error {
	ctx := context.Background()
	client, err := managedwriter.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("managedwriter.NewClient: %v", err)
	}
	defer client.Close()

	pendingStream, err := client.CreateWriteStream(ctx, &storagepb.CreateWriteStreamRequest{
		Parent: fmt.Sprintf("projects/%s/datasets/%s/tables/%s", projectID, datasetID, tableID),
		WriteStream: &storagepb.WriteStream{
			Type: storagepb.WriteStream_COMMITTED,
		},
	})

	m := &exampleproto.JsonData{}
	descriptorProto, err := adapt.NormalizeDescriptor(m.ProtoReflect().Descriptor())
	if err != nil {
		return fmt.Errorf("NormalizeDescriptor: %v", err)
	}

	// Instantiate a ManagedStream, which manages low level details like connection state and provides
	// additional features like a future-like callback for appends, etc.  NewManagedStream can also create
	// the stream on your behalf, but in this example we're being explicit about stream creation.
	managedStream, err := client.NewManagedStream(ctx, managedwriter.WithStreamName(pendingStream.GetName()),
		managedwriter.WithSchemaDescriptor(descriptorProto))
	if err != nil {
		return fmt.Errorf("NewManagedStream: %v", err)
	}

	n := 100000 // maximum batch size
	rows, err := generateJsonMessages(n)
	if err != nil {
		return fmt.Errorf("generateExampleMessages: %v", err)
	}

	var curOffset int64

	for i := 1; i <= 50000000/n; i++ {
		result, err := managedStream.AppendRows(ctx, rows, managedwriter.WithOffset(curOffset))
		if err != nil {
			return fmt.Errorf("AppendRows call error: %v", err)
		}
		_, err = result.GetResult(ctx)
		if err != nil {
			return fmt.Errorf("AppendRows call error: %v", err)
		}
		curOffset = curOffset + int64(n)
	}
	return nil
}
