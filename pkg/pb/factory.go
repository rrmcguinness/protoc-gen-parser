/*
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pb

import "github.com/GoogleCloudPlatform/proto-gen-parser/pkg/api"

type ProtobufFactory struct{}

func (p ProtobufFactory) NewPackage(path string, name string, comment string) api.Package {
	return &_package{
		Qualified:     newQualified(path, name, comment),
		OptionValues:  make([]api.Option, 0),
		ImportValues:  make([]api.Import, 0),
		MessageValues: make([]api.Message, 0),
		EnumValues:    make([]api.Enum, 0),
		ServiceValues: make([]api.Service, 0),
		Graph:         api.NewGraph(),
	}
}

func (p ProtobufFactory) NewMessage(qualifier string, name string, comment string) api.Message {
	return &message{
		Qualified:       newQualified(qualifier, name, comment),
		AttributeValues: make([]api.Attribute, 0),
		MessageValues:   make([]api.Message, 0),
		EnumValues:      make([]api.Enum, 0),
		ReservedValues:  make([]api.Reserved, 0),
		Graph:           api.NewGraph(),
	}
}

func (p ProtobufFactory) NewReserved(start int32, end int32) api.Reserved {
	return &reserved{
		start: start,
		end:   end,
	}
}

func (p ProtobufFactory) NewService(qualifier string, name string, comment string) api.Service {
	return &service{
		Qualified:            newQualified(qualifier, name, comment),
		remoteProcedureCalls: make([]api.RPC, 0),
		graph:                api.NewGraph(),
	}
}

func (p ProtobufFactory) NewAnnotation(name string, value string) api.Annotation {
	return &annotation{NameValue: name, ValueString: value}
}

func (p ProtobufFactory) NewAttribute(qualifier string, name string, comment string, isRepeated bool, isMap bool, ordinal int, kinds ...string) api.Attribute {
	out := &attribute{
		Qualified:     newQualified(qualifier, name, comment),
		RepeatedValue: isRepeated,
		IsMapValue:    isMap,
		KindValues:    kinds,
		OrdinalValue:  ordinal,
	}
	return out
}

func (p ProtobufFactory) NewEnum(qualifier string, name string, comment string, values ...api.EnumValue) api.Enum {
	return &enum{
		Qualified:  newQualified(qualifier, name, comment),
		EnumValues: values,
	}
}

func (p ProtobufFactory) NewEnumValue(qualifier string, name string, comment string, ordinal int) api.EnumValue {
	return &enumValue{
		Qualified: newQualified(qualifier, name, comment),
		ordinal:   ordinal,
	}
}

func (p ProtobufFactory) NewImport(path string, comment string) api.Import {
	return &_import{
		path:    path,
		comment: comment,
	}
}

func (p ProtobufFactory) NewOption(name string, value string, comment string) api.Option {
	return &option{
		NameValue:     name,
		ValueString:   value,
		CommentString: comment,
	}
}

func (p ProtobufFactory) NewMessageOption(name string, value string, comment string) api.MessageOption {
	return &messageOption{
		OptionValue: &option{
			NameValue:     name,
			ValueString:   value,
			CommentString: comment,
		},
	}
}

func (p ProtobufFactory) NewRPC(qualifier string, name string, comment string) api.RPC {
	return &rPC{
		Qualified:        newQualified(qualifier, name, comment),
		inputParameters:  make([]api.RPCParameter, 0),
		returnParameters: make([]api.RPCParameter, 0),
		rpcOptions:       make([]api.RPCOption, 0),
	}
}

func (p ProtobufFactory) NewRPCParameter(stream bool, kind string) api.RPCParameter {
	return &rPCParameter{
		stream: stream,
		kind:   kind,
	}
}
