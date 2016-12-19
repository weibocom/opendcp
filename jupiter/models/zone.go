/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */



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
