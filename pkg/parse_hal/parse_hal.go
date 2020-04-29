/*
 * Copyright The Spinnaker Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package parse_hal

import (
	"io/ioutil"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/protoyaml"
)

func ParseHalConfig(fn string) (*config.Hal, error) {
	dat, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	h := &config.Hal{}
	if err = protoyaml.Unmarshal(dat, h); err != nil {
		return nil, err
	}
	return h, nil
}