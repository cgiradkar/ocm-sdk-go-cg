/*
Copyright (c) 2022 Red Hat, Inc.

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

package testing

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"sync"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// This stores the zap.Sink objects that write to io.Writer objects, indexed by identifier.
var writerSinks *sync.Map = &sync.Map{}

func init() {
	// Register the zap sink that writes to a io.Writer:
	err := zap.RegisterSink("writer", func(url *url.URL) (sink zap.Sink, err error) {
		key := url.Host
		value, ok := writerSinks.Load(key)
		if !ok {
			err = fmt.Errorf("can't find writer sink for identifier '%s'", key)
			return
		}
		sink = value.(zap.Sink)
		return
	})
	if err != nil {
		panic(err)
	}
}

// MakeLogger returns a logger that writes to the given writer.
func MakeLogger(writer io.Writer) logr.Logger {
	// Create the object:
	sink := &writerSink{
		id:     uuid.NewString(),
		writer: writer,
	}

	// Store the sik in the map:
	writerSinks.Store(sink.id, sink)

	// Create the zap logger:
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.OutputPaths = []string{
		fmt.Sprintf("writer://%s", sink.id),
	}
	zapper, err := config.Build()
	if err != nil {
		panic(err)
	}

	// Redirect standard logging to the Ginkgo writer so that error messages generated by the
	// HTTP clients (for example when a protocol error happens) will not interfere with the
	// Ginkgo output:
	log.SetOutput(writer)

	// Create the logr logger:
	return zapr.NewLogger(zapper)
}

// writerSink is an implementation of the zap.Sink interface that writes to GinkgoWriter.
// This is intended for tests, so that the logs of the tests are written to the GinkgoWriter and
// displayed only when the test fails or when the verbose mode is enabled.
type writerSink struct {
	id     string
	writer io.Writer
}

func (s *writerSink) Write(b []byte) (n int, err error) {
	n, err = s.writer.Write(b)
	return
}

func (s *writerSink) Sync() error {
	return nil
}

func (s *writerSink) Close() error {
	writerSinks.Delete(s.id)
	return nil
}