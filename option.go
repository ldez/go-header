// Copyright (c) 2020-2025 Denis Tingaikin
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package goheader

type Option interface {
	apply(*Analyzer)
}

type applyAnalyzerOptionFunc func(*Analyzer)

func (f applyAnalyzerOptionFunc) apply(a *Analyzer) {
	f(a)
}

func WithValues(values map[string]Value) Option {
	return applyAnalyzerOptionFunc(func(a *Analyzer) {
		a.values = make(map[string]Value)
		for k, v := range values {
			a.values[k] = v
		}
	})
}

// WithDelims replaces default delims for parsing.
func WithDelims(delims string) Option {
	return applyAnalyzerOptionFunc(func(a *Analyzer) {
		var left = delims[:len(delims)/2]
		var right = delims[len(delims)/2:]

		a.delimsLeft = left
		a.delimsRight = right
	})
}

func WithTemplate(template string) Option {
	return applyAnalyzerOptionFunc(func(a *Analyzer) {
		a.template = template
	})
}
