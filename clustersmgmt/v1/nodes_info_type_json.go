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

// MarshalNodesInfo writes a value of the 'nodes_info' type to the given writer.
func MarshalNodesInfo(object *NodesInfo, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeNodesInfo(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeNodesInfo writes a value of the 'nodes_info' type to the given stream.
func writeNodesInfo(object *NodesInfo, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = object.bitmap_&1 != 0 && object.nodes != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("nodes")
		writeNodeInfoList(object.nodes, stream)
	}
	stream.WriteObjectEnd()
}

// UnmarshalNodesInfo reads a value of the 'nodes_info' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalNodesInfo(source interface{}) (object *NodesInfo, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readNodesInfo(iterator)
	err = iterator.Error
	return
}

// readNodesInfo reads a value of the 'nodes_info' type from the given iterator.
func readNodesInfo(iterator *jsoniter.Iterator) *NodesInfo {
	object := &NodesInfo{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "nodes":
			value := readNodeInfoList(iterator)
			object.nodes = value
			object.bitmap_ |= 1
		default:
			iterator.ReadAny()
		}
	}
	return object
}
