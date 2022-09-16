/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/renan-campos/ocm-sdk-go/clustersmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/renan-campos/ocm-sdk-go/helpers"
)

// MarshalAWSSpotMarketOptionsList writes a list of values of the 'AWS_spot_market_options' type to
// the given writer.
func MarshalAWSSpotMarketOptionsList(list []*AWSSpotMarketOptions, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeAWSSpotMarketOptionsList(list, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeAWSSpotMarketOptionsList writes a list of value of the 'AWS_spot_market_options' type to
// the given stream.
func writeAWSSpotMarketOptionsList(list []*AWSSpotMarketOptions, stream *jsoniter.Stream) {
	stream.WriteArrayStart()
	for i, value := range list {
		if i > 0 {
			stream.WriteMore()
		}
		writeAWSSpotMarketOptions(value, stream)
	}
	stream.WriteArrayEnd()
}

// UnmarshalAWSSpotMarketOptionsList reads a list of values of the 'AWS_spot_market_options' type
// from the given source, which can be a slice of bytes, a string or a reader.
func UnmarshalAWSSpotMarketOptionsList(source interface{}) (items []*AWSSpotMarketOptions, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	items = readAWSSpotMarketOptionsList(iterator)
	err = iterator.Error
	return
}

// readAWSSpotMarketOptionsList reads list of values of the ''AWS_spot_market_options' type from
// the given iterator.
func readAWSSpotMarketOptionsList(iterator *jsoniter.Iterator) []*AWSSpotMarketOptions {
	list := []*AWSSpotMarketOptions{}
	for iterator.ReadArray() {
		item := readAWSSpotMarketOptions(iterator)
		list = append(list, item)
	}
	return list
}
