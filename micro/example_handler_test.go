// Copyright 2022-2023 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package micro_test

import (
	"log"
	"strconv"

	"github.com/matsuev/nats.go"
	"github.com/matsuev/nats.go/micro"
)

type rectangle struct {
	height int
	width  int
}

// Handle is an implementation of micro.Handler used to
// calculate the area of a rectangle
func (r rectangle) Handle(req micro.Request) {
	area := r.height * r.width
	req.Respond([]byte(strconv.Itoa(area)))
}

func ExampleHandler() {
	nc, err := nats.Connect("127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	rec := rectangle{10, 5}

	config := micro.Config{
		Name:    "RectangleAreaService",
		Version: "0.1.0",
		Endpoint: &micro.EndpointConfig{
			Subject: "area.rectangle",
			Handler: rec,
		},
	}
	svc, err := micro.AddService(nc, config)
	if err != nil {
		log.Fatal(err)
	}
	defer svc.Stop()
}
