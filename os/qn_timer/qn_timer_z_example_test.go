// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_timer_test

import (
	"fmt"
	"time"

	"github.com/qnsoft/common/os/qn_timer"
)

func Example_add() {
	now := time.Now()
	interval := 1400 * time.Millisecond
	qn_timer.Add(interval, func() {
		fmt.Println(time.Now(), time.Duration(time.Now().UnixNano()-now.UnixNano()))
		now = time.Now()
	})

	select {}
}
