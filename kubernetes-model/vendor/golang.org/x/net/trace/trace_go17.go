/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package trace

import "context"

// NewContext returns a copy of the parent context
// and associates it with a Trace.
func NewContext(ctx context.Context, tr Trace) context.Context {
	return context.WithValue(ctx, contextKey, tr)
}

// FromContext returns the Trace bound to the context, if any.
func FromContext(ctx context.Context) (tr Trace, ok bool) {
	tr, ok = ctx.Value(contextKey).(Trace)
	return
}
