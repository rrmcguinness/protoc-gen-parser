// Copyright 2023 Ryan McGuinness
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reader

import (
	"strings"
)

type MessageOptionVisitor struct{}

func (o *MessageOptionVisitor) CanVisit(line *Line) bool {
	return strings.HasPrefix(line.Syntax, "option ") && strings.Index(line.Syntax, "(") > 0 && strings.Index(line.Syntax, ")") > 0 && strings.Index(line.Syntax, "=") > 0
}

func (o *MessageOptionVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
	// option (gen_bq_schema.bigquery_opts).table_name = "tbl_audit_record";
	optionName := strings.Trim(in.Syntax[strings.Index(in.Syntax, "("):strings.Index(in.Syntax, "=")], " ")
	optionValue := in.Syntax[strings.Index(in.Syntax, "\"")+1 : strings.LastIndex(in.Syntax, "\"")]

	return ProtobufFactory.NewOption(optionName, optionValue, "")
}
