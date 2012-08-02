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

package timespans

import (
	"testing"
)

var (
	dest = `
Tag,Prefix
GERMANY,49
GERMANY_O2,41
GERMANY_PREMIUM,43
ALL,49
ALL,41
ALL,43
NAT,0256
NAT,0257
NAT,0723
RET,0723
RET,0724
`
	rts = `
RT_STANDARD,GERMANY,0,0.2,60,1
RT_STANDARD,GERMANY_O2,0,0.1,60,1
RT_STANDARD,GERMANY_PREMIUM,0,0.1,60,1
RT_DEFAULT,ALL,0,0.1,60,1
RT_STD_WEEKEND,GERMANY,0,0.1,60,1
RT_STD_WEEKEND,GERMANY_O2,0,0.05,60,1
P1,NAT,0,1,1,1
P2,NAT,0,0.5,1,1
`
	ts = `
WORKDAYS_00,*all,*all,*all,1;2;3;4;5,00:00:00
WORKDAYS_18,*all,*all,*all,1;2;3;4;5,18:00:00
WEEKENDS,*all,*all,*all,6;7,00:00:00
ONE_TIME_RUN,2012,*none,*none,*none,*asap
`
	rtts = `
STANDARD,RT_STANDARD,WORKDAYS_00,10
STANDARD,RT_STD_WEEKEND,WORKDAYS_18,10
STANDARD,RT_STD_WEEKEND,WEEKENDS,10
PREMIUM,RT_STD_WEEKEND,WEEKENDS,10
DEFAULT,RT_DEFAULT,WORKDAYS_00,10
EVENING,P1,WORKDAYS_00,10
EVENING,P2,WORKDAYS_18,10
EVENING,P2,WEEKENDS,10
`
	rp = `
CUSTOMER_1,0,OUT,rif:from:tm,danb,PREMIUM,2012-01-01T00:00:00Z
CUSTOMER_1,0,OUT,rif:from:tm,danb,STANDARD,2012-02-28T00:00:00Z
CUSTOMER_2,0,OUT,danb:87.139.12.167,danb,STANDARD,2012-01-01T00:00:00Z
CUSTOMER_1,0,OUT,danb,,PREMIUM,2012-01-01T00:00:00Z
vdf,0,OUT,rif,,EVENING,2012-01-01T00:00:00Z
vdf,0,OUT,rif,,EVENING,2012-02-28T00:00:00Z
vdf,0,OUT,minu,,EVENING,2012-01-01T00:00:00Z
vdf,0,OUT,minu,,EVENING,2012-02-28T00:00:00Z
`
	a = `
MINI,TOPUP,MINUTES,OUT,100,NAT,ABSOLUTE,0,10,10
`
	atms = `
MORE_MINUTES,MINI,ONE_TIME_RUN,10
`
	atrs = `
STANDARD_TRIGGER,MINUTES,OUT,10,GERMANY_O2,SOME_1,10
STANDARD_TRIGGER,MINUTES,OUT,200,GERMANY,SOME_2,10
`
	accs = `
vdf,minitsboy,OUT,MORE_MINUTES,STANDARD_TRIGGER
`
)

func init() {
	csvr := &CSVReader{OpenStringCSVReader}
	csvr.LoadDestinations(dest, ',')
	csvr.LoadRates(rts, ',')
	csvr.LoadTimings(ts, ',')
	csvr.LoadRateTimings(rtts, ',')
	csvr.LoadRatingProfiles(rp, ',')
	csvr.LoadActions(a, ',')
	csvr.LoadActionTimings(atms, ',')
	csvr.LoadActionTriggers(atrs, ',')
	csvr.LoadAccountActions(accs, ',')
	WriteToDatabase(storageGetter, false, false)

}

func TestLoadDestinations(t *testing.T) {
	if len(destinations) != 6 {
		t.Error("Failed to load destinations: ", destinations)
	}
}

func TestLoadRates(t *testing.T) {
	if len(rates) != 5 {
		t.Error("Failed to load rates: ", rates)
	}
}

func TestLoadTimimgs(t *testing.T) {
	if len(timings) != 4 {
		t.Error("Failed to load timings: ", timings)
	}
}

func TestLoadRateTimings(t *testing.T) {
	if len(activationPeriods) != 4 {
		t.Error("Failed to load rate timings: ", activationPeriods)
	}
}

func TestLoadRatingProfiles(t *testing.T) {
	if len(ratingProfiles) != 7 {
		t.Error("Failed to load rating profiles: ", len(ratingProfiles), ratingProfiles)
	}
}

func TestLoadActions(t *testing.T) {
	if len(actions) != 1 {
		t.Error("Failed to load actions: ", actions)
	}
}

func TestLoadActionTimings(t *testing.T) {
	if len(actionsTimings) != 1 {
		t.Error("Failed to load action timings: ", actionsTimings)
	}
}

func TestLoadActionTriggers(t *testing.T) {
	if len(actionsTriggers) != 1 {
		t.Error("Failed to load action triggers: ", actionsTriggers)
	}
}

func TestLoadAccountActions(t *testing.T) {
	if len(accountActions) != 1 {
		t.Error("Failed to load account actions: ", accountActions)
	}
}