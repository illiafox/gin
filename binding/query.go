// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import "net/http"

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (q queryBinding) Bind(req *http.Request, obj any) error {
	values := req.URL.Query()
	if err := mapFormByTag(obj, values, q.Name()); err != nil {
		return err
	}
	return validate(obj)
}
