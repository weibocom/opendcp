package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type RegionService interface {
	DescribeRegions(params map[string]interface{}) (DescribeRegionsResponse, error)
	DescribeZones(params map[string]interface{}) (DescribeZonesResponse, error)
}

type RegionOperator struct {
	Common *util.CommonParam
}

// Response struct for DescribeRegions
type DescribeRegionsResponse struct {
	util.ErrorResponse
	AllRegions Regions `json:"Regions"`
}

type Regions struct {
	AllRegion []RegionType `json:"Region"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&regiontype
type RegionType struct {
	RegionId   string `json:"RegionId"`
	RegionName string `json:"RegionName"`
}

// Response struct for DescribeZones
type DescribeZonesResponse struct {
	util.ErrorResponse
	AllZones Zones `json:"Zones"`
}

type Zones struct {
	AllZone []Zone `json:"Zone"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&zonetype
type Zone struct {
	ZoneId                    string                        `json:"ZoneId"`
	LocalName                 string                        `json:"LocalName"`
	AvailableResourceCreation AvailableResourceCreationType `json:"AvailableResourceCreation"`
	AvailableDiskCategories   AvailableDiskCategoriesType   `json:"AvailableDiskCategories"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&availableresourcecreationtype
type AvailableResourceCreationType struct {
	ResourceTypes []string `json:"ResourceTypes"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&availablediskcategoriestype
type AvailableDiskCategoriesType struct {
	DiskCategories []string `json:"DiskCategories"`
}

func (op *RegionOperator) DescribeRegions(params map[string]interface{}) (DescribeRegionsResponse, error) {
	var resp DescribeRegionsResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *RegionOperator) DescribeZones(params map[string]interface{}) (DescribeZonesResponse, error) {
	var resp DescribeZonesResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
