// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package servermanager

import (
	"context"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type LogWriter struct {
	ctx context.Context
}

func (l *LogWriter) Write(p []byte) (n int, err error) {
	runtime.LogDebug(l.ctx, strings.TrimSuffix(string(p), "\n"))
	return len(p), nil
}
