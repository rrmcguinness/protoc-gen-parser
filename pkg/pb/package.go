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

import (
	"log"

	"github.com/GoogleCloudPlatform/proto-gen-parser/pkg/api"
)

// option an option implementation
type option struct {
	NameValue     string
	ValueString   string
	CommentString string
}

func (o *option) Name() string {
	return o.NameValue
}

func (o *option) Value() string {
	return o.ValueString
}

func (o *option) Comment() string {
	return o.CommentString
}

func (o *option) SetComment(in string) {
	o.CommentString = in
}

// _import an import implementation
type _import struct {
	path    string
	comment string
}

func (i *_import) Path() string {
	return i.path
}

func (i *_import) Comment() string {
	return i.comment
}

func (i *_import) SetComment(in string) {
	i.comment = in
}

// _package - the package implementation
type _package struct {
	api.Qualified
	OptionValues  []api.Option
	ImportValues  []api.Import
	MessageValues []api.Message
	EnumValues    []api.Enum
	ServiceValues []api.Service
	Graph         api.Graph
}

func addVertex(pkg api.Package, vertexType api.VertexType, qualified api.Qualified) {
	err := pkg.GetGraph().AddVertex(api.NewVertex(qualified.Name(), vertexType))
	if err != nil {
		log.Default().Printf("failed to add vertex to package: %v", err)
	}
}

func (p *_package) Options() []api.Option {
	return p.OptionValues
}

func (p *_package) AddOption(name string, value string, comment string) api.Package {
	p.OptionValues = append(p.OptionValues, &option{NameValue: name, ValueString: value, CommentString: comment})
	return p
}

func (p *_package) Imports() []api.Import {
	return p.ImportValues
}

func (p *_package) AddImport(path string, comment string) api.Package {
	p.ImportValues = append(p.ImportValues, &_import{path: path, comment: comment})
	return p
}

func (p *_package) Messages() []api.Message {
	return p.MessageValues
}

func (p *_package) AddMessage(message api.Message) api.Package {
	addVertex(p, api.MESSAGE, message)
	p.MessageValues = append(p.MessageValues, message)
	return p
}

func (p *_package) Enums() []api.Enum {
	return p.EnumValues
}

func (p *_package) AddEnum(enum api.Enum) api.Package {
	addVertex(p, api.ENUM, enum)
	p.EnumValues = append(p.EnumValues, enum)
	return p
}

func (p *_package) Services() []api.Service {
	return p.ServiceValues
}

func (p *_package) AddService(service api.Service) api.Package {
	addVertex(p, api.SERVICE, service)
	p.ServiceValues = append(p.ServiceValues, service)
	return p
}

func (p _package) GetGraph() api.Graph {
	return p.Graph
}
