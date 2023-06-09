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

type enum struct {
	api.Qualified
	EnumValues []api.EnumValue
}

func (e *enum) Values() []api.EnumValue {
	return e.EnumValues
}

func (e *enum) AddValue(value api.EnumValue) api.Enum {
	e.EnumValues = append(e.EnumValues, value)
	return e
}

type enumValue struct {
	api.Qualified
	ordinal int
}

func (e *enumValue) Ordinal() int {
	return e.ordinal
}
