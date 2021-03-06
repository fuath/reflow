// Copyright 2018 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package blob

import (
	"context"
	"io"
	"net/url"
	"strings"

	"github.com/grailbio/reflow"
	"github.com/grailbio/reflow/errors"
)

// Mux multiplexes a number of blob store implementations. Mux
// implements bucket operations based on blob store URLs. URLs
// that are passed into Mux are intepreted as:
//
//	store://bucket/key
type Mux map[string]Store

// Bucket parses the provided URL, looks up its implementation, and
// returns the store's Bucket and the prefix implied by the URL. A
// errors.NotSupported is returned if there is no implementation for
// the requested scheme.
func (m Mux) Bucket(ctx context.Context, rawurl string) (Bucket, string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, "", err
	}
	store, ok := m[u.Scheme]
	if !ok {
		return nil, "", errors.E(errors.NotSupported, "blob.Bucket", rawurl,
			errors.Errorf("no implementation for scheme %s", u.Scheme))
	}
	bucket, err := store.Bucket(ctx, u.Host)
	return bucket, strings.TrimPrefix(u.Path, "/"), err
}

// File returns file metadata for the provided URL.
func (m Mux) File(ctx context.Context, url string) (reflow.File, error) {
	bucket, key, err := m.Bucket(ctx, url)
	if err != nil {
		return reflow.File{}, err
	}
	return bucket.File(ctx, key)
}

// Scan returns a scanner for the provided URL (which represents a
// prefix).
func (m Mux) Scan(ctx context.Context, url string) (Scanner, error) {
	bucket, prefix, err := m.Bucket(ctx, url)
	if err != nil {
		return nil, err
	}
	return bucket.Scan(prefix), nil
}

// Download downloads the object named by the provided URL to the
// provided io.WriterAt. If the provided etag is nonempty, then it is
// checked as a precondition on downloading the object. Download may
// download multiple chunks concurrently.
func (m Mux) Download(ctx context.Context, url, etag string, w io.WriterAt) (int64, error) {
	bucket, key, err := m.Bucket(ctx, url)
	if err != nil {
		return -1, err
	}
	return bucket.Download(ctx, key, etag, w)
}

// Get returns a (streaming) reader of the object named by the
// provided URL. If the provided etag is nonempty, then it is checked
// as a precondition on streaming the object.
func (m Mux) Get(ctx context.Context, url, etag string) (io.ReadCloser, reflow.File, error) {
	bucket, key, err := m.Bucket(ctx, url)
	if err != nil {
		return nil, reflow.File{}, err
	}
	return bucket.Get(ctx, key, etag)
}

// Put stores the contents of the provided io.Reader at the provided URL.
func (m Mux) Put(ctx context.Context, url string, body io.Reader) error {
	bucket, key, err := m.Bucket(ctx, url)
	if err != nil {
		return err
	}
	return bucket.Put(ctx, key, body)
}

// Snapshot returns an un-loaded Reflow fileset representing the contents
// of the provided URL.
func (m Mux) Snapshot(ctx context.Context, url string) (reflow.Fileset, error) {
	bucket, prefix, err := m.Bucket(ctx, url)
	if err != nil {
		return reflow.Fileset{}, err
	}
	return bucket.Snapshot(ctx, prefix)
}
