// Copyright 2021 The ZKits Project Authors.
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

package main

import (
	"fmt"

	"github.com/edoger/zkits/internal/parser"
)

var src = `
# 中文 😁 xxx
@api(/api/v1/get) {
    @method(GET)
    @handler(GetHandler(Request) return (Response))
}

@api(/api/v1/post) {
    @handlerS(404)
}

`

func main() {
	p := parser.NewLexer([]byte(src))
	text, err := p.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("|" + text + "|")
}
