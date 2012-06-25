/*
Rating system designed to be used in VoIP Carriers World
Copyright (C) 2012  Radu Ioan Fericean

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package main

import (
	"github.com/cgrates/cgrates/timespans"
	"log"
	"strconv"
	"time"
)

type Rate struct {
	DestinationsTag                        string
	ConnectFee, Price, BillingUnit, Weight float64
}

func NewRate(destinationsTag, connectFee, price, billingUnit, weight string) (r *Rate, err error) {
	cf, err := strconv.ParseFloat(connectFee, 64)
	if err != nil {
		log.Printf("Error parsing connect fee from: %v", connectFee)
		return
	}
	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Printf("Error parsing price from: %v", price)
		return
	}
	bu, err := strconv.ParseFloat(billingUnit, 64)
	if err != nil {
		log.Printf("Error parsing billing unit from: %v", billingUnit)
		return
	}
	w, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		log.Printf("Error parsing weight from: %v", weight)
		return
	}
	r = &Rate{
		DestinationsTag: destinationsTag,
		ConnectFee:      cf,
		Price:           p,
		BillingUnit:     bu,
		Weight:          w,
	}
	return
}

type Timing struct {
	Months    timespans.Months
	MonthDays timespans.MonthDays
	WeekDays  timespans.WeekDays
	StartTime string
}

func NewTiming(timeingInfo ...string) (rt *Timing) {
	rt = &Timing{}
	rt.Months.Parse(timeingInfo[0], ";")
	rt.MonthDays.Parse(timeingInfo[1], ";")
	rt.WeekDays.Parse(timeingInfo[2], ";")
	if timeingInfo[3] == "*now" {
		rt.StartTime = time.Now().Format("00:00:00")
	} else {
		rt.StartTime = timeingInfo[3]
	}
	return
}

type RateTiming struct {
	RatesTag string
	timing   *Timing
}

func NewRateTiming(ratesTag string, timing *Timing) (rt *RateTiming) {
	rt = &RateTiming{
		RatesTag: ratesTag,
		timing:   timing,
	}
	return
}

func (rt *RateTiming) GetInterval(r *Rate) (i *timespans.Interval) {
	i = &timespans.Interval{
		Months:      rt.timing.Months,
		MonthDays:   rt.timing.MonthDays,
		WeekDays:    rt.timing.WeekDays,
		StartTime:   rt.timing.StartTime,
		ConnectFee:  r.ConnectFee,
		Price:       r.Price,
		BillingUnit: r.BillingUnit,
		Weight:      r.Weight,
	}
	return
}

type CallDescriptors []*timespans.CallDescriptor

func (cds CallDescriptors) getKey(key string) *timespans.CallDescriptor {
	for _, cd := range cds {
		if cd.GetKey() == key {
			return cd
		}
	}
	return nil
}