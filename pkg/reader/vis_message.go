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

package reader

import (
	"strings"

	"github.com/GoogleCloudPlatform/proto-gen-parser/pkg/api"
	"github.com/GoogleCloudPlatform/proto-gen-parser/pkg/logging"
)

// MessageVisitor is used for interpreting message text
type MessageVisitor struct {
	Log *logging.Logger
}

func NewMessageVisitor(debug bool) *MessageVisitor {
	return &MessageVisitor{Log: logging.NewLogger(debug, "message visitor")}
}

// CanVisit visits if the line starts with 'message' and ends with an open brace '{'
func (mv *MessageVisitor) CanVisit(in *Line) bool {
	return strings.HasPrefix(in.Syntax, "message ") && in.Token == OpenBrace
}

// Visit evaluates the current line and parses the message until the closed brace
// is evaluated.
func (mv *MessageVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	mv.Log.Debugf("Visiting Message: %s :: %s\n", in.Syntax, in.Token)

	values := SplitSyntax(in.Syntax)
	out := ProtobufFactory.NewMessage(
		Join(Period, namespace, values[1]),
		values[1],
		in.Comment)

	comment := ""

	for scanner.Scan() {
		line := scanner.ReadLine()

		mv.Log.Debugf("Current Line: `%s` :: `%s`\n", line.Syntax, line.Token)

		if strings.HasSuffix(line.Token, ClosedBrace) {
			break
		}
		for _, visitor := range RegisteredVisitors {
			if visitor.CanVisit(line) {
				rt := visitor.Visit(
					scanner,
					line,
					Join(Period, namespace, out.Name()))
				switch t := rt.(type) {
				case api.Message:
					t.SetComment(Join(Space, comment, t.Comment()))
					out.AddMessage(t)
					comment = ""
				case api.Enum:
					t.SetComment(Join(Space, comment, t.Comment()))
					out.AddEnum(t)
					comment = ""
				case api.Attribute:
					if t.Validate() {
						t.SetComment(Join(Space, comment, t.Comment()))
						out.AddAttribute(t)
						comment = ""
					}
				case []api.Annotation:
					attrs := out.Attributes()
					if len(attrs) > 0 {
						cAttr := attrs[len(attrs)-1]
						for _, a := range t {
							cAttr.AddAnnotation(a.Name(), a.Value())
						}
					}
				case api.MessageOption:
					t.SetComment(Join(Space, comment, t.Comment()))
					out.AddOption(t)
					comment = ""
				case api.Reserved:
					out.AddReserved(t.Start(), t.End())
				case string:
					comment = Join(Space, comment, t)
				}
			}
		}
	}
	return out
}
