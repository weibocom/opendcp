// Copyright 2016 Weibo Inc. All Rights Reserved.
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

package models

// Contains the output of Regions.
type RegionsResp struct {

	// Information about one or more regions.
	Regions []Region `locationName:"regionInfo" locationNameList:"item" type:"list"`
}

// Describes a region.
type Region struct {

	// The region service endpoint.
	Endpoint string `locationName:"regionEndpoint" type:"string"`

	// The name of the region.
	RegionName string `locationName:"regionName" type:"string"`
}
