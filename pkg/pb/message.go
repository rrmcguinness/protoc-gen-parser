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
	"github.com/GoogleCloudPlatform/proto-gen-parser/pkg/api"
	"log"
)

type message struct {
	api.Qualified
	OptionValues    []api.MessageOption
	AttributeValues []api.Attribute
	MessageValues   []api.Message
	EnumValues      []api.Enum
	ReservedValues  []api.Reserved
	Graph           api.Graph
}

func (m *message) Attributes() []api.Attribute {
	return m.AttributeValues
}

func (m *message) Options() []api.MessageOption {
	return m.OptionValues
}

func (m *message) AddOption(option api.MessageOption) api.Message {
	if !m.ContainsOptionName(option.Name()) {
		m.OptionValues = append(m.OptionValues, option)
	}
	return m
}

func (m *message) ContainsOptionName(name string) bool {
	contains := false
	for _, o := range m.OptionValues {
		if o.Name() == name {
			contains = true
			break
		}
	}
	return contains
}

func (m *message) AddAttribute(attribute api.Attribute) api.Message {
	m.AttributeValues = append(m.AttributeValues, attribute)
	return m
}

func (m *message) Messages() []api.Message {
	return m.MessageValues
}

func (m *message) AddMessage(message api.Message) api.Message {
	err := m.Graph.AddVertex(api.NewVertex(message.Name(), api.MESSAGE))
	if err != nil {
		log.Default().Printf("error adding message to Graph: %v", err)
	}
	m.MessageValues = append(m.MessageValues, message)
	return m
}

func (m *message) Enums() []api.Enum {
	return m.EnumValues
}

func (m *message) AddEnum(enum api.Enum) api.Message {
	err := m.Graph.AddVertex(api.NewVertex(enum.Name(), api.ENUM))
	if err != nil {
		log.Default().Printf("error adding enum to Graph: %v", err)
	}
	m.EnumValues = append(m.EnumValues, enum)
	return m
}

func (m *message) Reserved() []api.Reserved {
	return m.ReservedValues
}

func (m *message) AddReserved(start int32, end int32) api.Message {
	m.ReservedValues = append(m.ReservedValues, &reserved{start: start, end: end})
	return m
}

func (m *message) GetGraph() api.Graph {
	return m.Graph
}

// reserved the default implementation for api.Reserved
type reserved struct {
	start int32
	end   int32
}

func (r *reserved) Start() int32 {
	return r.start
}

func (r *reserved) End() int32 {
	return r.end
}

// annotation implements api.Annotation
type annotation struct {
	NameValue   string
	ValueString string
}

func (a *annotation) Name() string {
	return a.NameValue
}

func (a *annotation) Value() string {
	return a.ValueString
}

// attribute is the implementation for api.Attribute
type attribute struct {
	api.Qualified
	RepeatedValue    bool
	IsMapValue       bool
	KindValues       []string
	OrdinalValue     int
	AnnotationValues []api.Annotation
}

func (a *attribute) Validate() bool {
	return len(a.Name()) > 0 && a.Kinds() != nil && len(a.Kinds()) >= 1 && a.OrdinalValue >= 1
}

func (a *attribute) Repeated() bool {
	return a.RepeatedValue
}

func (a *attribute) Map() bool {
	return a.IsMapValue
}

func (a *attribute) Kinds() []string {
	return a.KindValues
}

func (a *attribute) Annotations() []api.Annotation {
	return a.AnnotationValues
}

func (a *attribute) Ordinal() int {
	return a.OrdinalValue
}

func (a *attribute) AddAnnotation(name string, value string) api.Attribute {
	a.AnnotationValues = append(a.AnnotationValues, &annotation{NameValue: name, ValueString: value})
	return a
}

type messageOption struct {
	OptionValue api.Option
}

func (m *messageOption) SetComment(comment string) {
	m.OptionValue.SetComment(comment)
}

func (m *messageOption) Name() string {
	return m.OptionValue.Name()
}

func (m *messageOption) Comment() string {
	return m.OptionValue.Comment()
}

func (m *messageOption) Value() string {
	return m.Value()
}
