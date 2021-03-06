// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_proc

import (
	"errors"
	"io"

	"github.com/qnsoft/common/internal/json"
	"github.com/qnsoft/common/net/qn_tcp"
)

// Send sends data to specified process of given pid.
func Send(pid int, data []byte, group ...string) error {
	msg := MsgRequest{
		SendPid: Pid(),
		RecvPid: pid,
		Group:   gPROC_COMM_DEFAULT_GRUOP_NAME,
		Data:    data,
	}
	if len(group) > 0 {
		msg.Group = group[0]
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	var conn *qn_tcp.PoolConn
	conn, err = getConnByPid(pid)
	if err != nil {
		return err
	}
	defer conn.Close()
	// Do the sending.
	var result []byte
	result, err = conn.SendRecvPkg(msgBytes, qn_tcp.PkgOption{
		Retry: qn_tcp.Retry{
			Count: 3,
		},
	})
	if len(result) > 0 {
		response := new(Msqn_response)
		err = json.Unmarshal(result, response)
		if err == nil {
			if response.Code != 1 {
				err = errors.New(response.Message)
			}
		}
	}
	// EOF is not really an error.
	if err == io.EOF {
		err = nil
	}
	return err
}
