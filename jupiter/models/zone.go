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

// Contains the output of DescribeAvailabiltyZones.
type AvailabilityZonesResp struct {

	// Information about one or more Availability Zones.
	AvailabilityZones []AvailabilityZone `locationName:"availabilityZoneInfo" locationNameList:"item" type:"list"`
}

// Describes an Availability Zone.
type AvailabilityZone struct {

	// Any messages about the Availability Zone.
	Messages []AvailabilityZoneMessage `locationName:"messageSet" locationNameList:"item" type:"list"`

	// The name of the region.
	RegionName string

	// The state of the Availability Zone.
	State string `locationName:"zoneState" type:"string" enum:"AvailabilityZoneState"`

	// The name of the Availability Zone.
	ZoneName string
}

// Describes a message about an Availability Zone.
type AvailabilityZoneMessage struct {

	// The message about the Availability Zone.
	Message string `locationName:"message" type:"string"`
}
