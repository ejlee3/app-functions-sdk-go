//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package transforms

import (
	"errors"

	"github.com/edgexfoundry/app-functions-sdk-go/v2/appcontext"
	"github.com/edgexfoundry/app-functions-sdk-go/v2/pkg/util"
)

type CoreData struct {
	DeviceName  string
	ReadingName string
}

// NewCoreData Is provided to interact with CoreData
func NewCoreData() *CoreData {
	coredata := &CoreData{}
	return coredata
}

// PushToCoreData pushes the provided value as an event to CoreData using the device name and reading name that have been set. If validation is turned on in
// CoreServices then your deviceName and readingName must exist in the CoreMetadata and be properly registered in EdgeX.
func (cdc *CoreData) PushToCoreData(edgexcontext *appcontext.Context, params ...interface{}) (bool, interface{}) {
	if len(params) < 1 {
		// We didn't receive a result
		return false, errors.New("No Data Received")
	}
	val, err := util.CoerceType(params[0])
	if err != nil {
		return false, err
	}
	result, err := edgexcontext.PushToCoreData(cdc.DeviceName, cdc.ReadingName, val)
	if err != nil {
		return false, err
	}
	return true, result
}
