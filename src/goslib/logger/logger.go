/*
The MIT License (MIT)

Copyright (c) 2018 SavinMax. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package logger

import (
	"fmt"
	"log"
	"strings"
)

var (
	_debug_open bool
)

func ERR(v ...interface{}) {
	msg := fmt.Sprintf("\033[1;4;31m[ERROR] %v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
	//raven.CaptureMessage(msg, map[string]string{"category": "error"})
	log.Printf(msg)
}

func ERRDirect(v ...interface{}) {
	msg := fmt.Sprintf("\033[1;4;31m[ERROR] %v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
	log.Printf(msg)
}

func WARN(v ...interface{}) {
	log.Printf("\033[1;33m[WARN] %v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

func INFO(v ...interface{}) {
	log.Printf("\033[32m[INFO] %v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

func NOTICE(v ...interface{}) {
	log.Printf("[NOTICE] %v\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

func DEBUG(v ...interface{}) {
	if _debug_open {
		log.Printf("\033[1;35m[DEBUG] %v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
	}
}

func format(v ...interface{}) string {
	return fmt.Sprintf("%v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
}
