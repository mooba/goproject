// Copyright 2021 Shopee, Inc.
// author pengchengbai
// date 2021/1/25

package main

import (
	"errors"
	"fmt"
	"strings"
)

type Region string

const (
	ID Region = "ID"
	TH Region = "TH"
	SG Region = "SG"
	VN Region = "VN"
	PH Region = "PH"
	TW Region = "TW"
	MY Region = "MY"
	BR Region = "BR"
)

func (r *Region) getRegion(regionStr string) error {
	regionStr = strings.ToUpper(regionStr)
	tmpRegion := Region(regionStr)
	switch tmpRegion {
	case ID,TH,SG,VN,PH,TW,MY,BR:
		*r = tmpRegion
		return nil
	}

	return errors.New("Invalid leave type.")
}

func main() {
	var region Region
	err := region.getRegion("ss")
	if err == nil {
		fmt.Println(region)
	} else {
		fmt.Println(err.Error())
	}

}
