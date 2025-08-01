// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.

package modbus

import (
	"fmt"
	"testing"
	"time"
)

func TestReconnect(t *testing.T) {
	handler := NewTCPClientHandler("127.0.0.1:502")
	handler.Timeout = 3 * time.Second
	handler.SlaveId = byte(1)
	handler.IdleTimeout = 3 * time.Second
	err := handler.Connect()
	if err != nil {
		fmt.Printf("Connect err: %v\n", err)
	}

	NewClient(handler)

}
