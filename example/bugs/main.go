//
//  Copyright 2024, easysoft
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

package main

import (
	"log"

	zentao "github.com/xinnss/go-zentao/v21/zentao"
)

func main() {
	zt, err := zentao.NewBasicAuthClient(
		"admin",
		"zzxxcc1476@A",
		zentao.WithBaseURL("http://127.0.0.1"),
		zentao.WithDevMode(),
		zentao.WithDumpAll(),
		zentao.WithoutProxy(),
	)
	if err != nil {
		panic(err)
	}

	p5, _, err := zt.Bugs.Create(3, zentao.BugCreateMeta{
		BugMeta: zentao.BugMeta{
			Title:      "go sdk bug",
			Type:       zentao.CodeErrorBugType,
			Severity:   1,
			Pri:        1,
			Project:    3,
			Steps:      "<p>[步骤]</p><p></p><p>[结果]</p><p></p><p>[期望]</p>",
			AssignedTo: "productManager",
			Mailto:     []string{"admin", "yangwenkai"},
		},
		Openedbuild: []string{"trunk", "v1"}, // 主干
	})
	if err != nil {
		panic(err)
	}
	log.Printf("product 1 create bug: %v", p5.Title)

}
