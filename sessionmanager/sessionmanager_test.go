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

package sessionmanager

import (
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	sm := &SessionManager{}
	sm.Connect("localhost:8021", "ClueCon")
	for {
		ev := sm.ReadNextEvent()
		log.Print(ev.Fields["Event-Name"])
		log.Print(ev)
	}
}