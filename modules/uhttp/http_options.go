// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package uhttp

import (
	"go.uber.org/fx/modules"
	"go.uber.org/fx/service"
)

const _filterKey = "uhttpFilter"

// WithFilters adds Filters to uhttp Module that will be applied to all incoming http requests.
func WithFilters(fs ...Filter) modules.Option {
	return func(mci *service.ModuleCreateInfo) error {
		filters := filtersFromCreateInfo(*mci)
		filters = append(filters, fs...)
		if mci.Items == nil {
			mci.Items = make(map[string]interface{})
		}
		mci.Items[_filterKey] = filters

		return nil
	}
}

<<<<<<< HEAD:modules/uhttp/http_options.go
func filtersFromCreateInfo(mci service.ModuleCreateInfo) []Filter {
	items, ok := mci.Items[_filterKey]
	if !ok {
		return nil
	}
=======
// ServeHTTP calls Handler.ServeHTTP( w, r) and injects a new service context for use.
func (h *handlerWithHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := ulog.NewLogContext(r.Context(), ulog.Logger(r.Context()))
	stopwatch := stats.HTTPMethodTimer.Timer(r.Method).Start()
	defer stopwatch.Stop()
>>>>>>> b69958d... rebase from master:modules/uhttp/handler.go

	// Intentionally panic if programmer adds non-filter slice to the data
	return items.([]Filter)
}
