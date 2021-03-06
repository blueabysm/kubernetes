/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package dialer

import (
	"net"
	"os"
	"syscall"
	"time"

	winio "github.com/Microsoft/go-winio"
)

func isNoent(err error) bool {
	if err != nil {
		if oerr, ok := err.(*os.PathError); ok {
			if oerr.Err == syscall.ENOENT {
				return true
			}
		}
	}
	return false
}

func dialer(address string, timeout time.Duration) (net.Conn, error) {
	return winio.DialPipe(address, &timeout)
}

// DialAddress returns the dial address
func DialAddress(address string) string {
	return address
}
