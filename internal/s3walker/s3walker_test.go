// Copyright 2017 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package s3walker

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/grailbio/testutil/s3test"
)

const bucket = "test"

func TestS3Walker(t *testing.T) {
	client := s3test.NewClient(t, bucket)
	want := []string{"test/x", "test/y", "test/z/foobar"}
	keys := append([]string{"unrelated"}, want...)
	for _, key := range keys {
		client.SetFile(key, []byte(key), "unused")
	}
	ctx := context.Background()
	w := &S3Walker{S3: client, Bucket: bucket, Prefix: "test/"}
	var got []string
	for w.Scan(ctx) {
		got = append(got, aws.StringValue(w.Object().Key))
	}
	if err := w.Err(); err != nil {
		t.Error(err)
	}
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestS3WalkerFile(t *testing.T) {
	client := s3test.NewClient(t, bucket)
	const key = "path/to/a/file"
	client.SetFile(key, []byte("contents"), "unused")
	ctx := context.Background()
	w := &S3Walker{S3: client, Bucket: bucket, Prefix: key}
	var got []string
	for w.Scan(ctx) {
		got = append(got, aws.StringValue(w.Object().Key))
	}
	if err := w.Err(); err != nil {
		t.Error(err)
	}
	if want := []string{key}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
