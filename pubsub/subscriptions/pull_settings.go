// Copyright 2019 Google LLC
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

package subscriptions

// [START pubsub_subscriber_flow_settings]
import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/pubsub"
)

func pullMsgsSettings(w io.Writer, projectID, subName string) error {
	// projectID := "my-project-id"
	// subName := projectID + "-example-sub"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	sub := client.Subscription(subName)
	sub.ReceiveSettings.Synchronous = true
	// MaxOutstandingMessages is the maximum number of unprocessed messages (unacknowledged but not yet expired).
	// This is only guaranteed when ReceiveSettings.Synchronous is set to true.
	sub.ReceiveSettings.MaxOutstandingMessages = 10
	// MaxOutstandingBytes is the maximum size of unprocessed messages (unacknowledged but not yet expired).
	sub.ReceiveSettings.MaxOutstandingBytes = 1e10
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("Receive: %v", err)
	}
	return nil
}

// [END pubsub_subscriber_flow_settings]
