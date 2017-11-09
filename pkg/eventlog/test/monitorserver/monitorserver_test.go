
// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.
 
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.
 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
 
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package monitorserver

import (
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/pebbe/zmq4"
)

var urlData = `
2017-05-19 11:33:34 APPS SumTimeByUrl [{"tenant":"o2o","service":"zzcplus","url":"/active/js/wx_share.js","avgtime":"1.453","sumtime":"1.453","counts":"1"}]
`

func BenchmarkMonitorServer(t *testing.B) {
	client, _ := zmq4.NewSocket(zmq4.PUB)
	// client.Monitor("inproc://monitor.rep", zmq4.EVENT_ALL)
	// go monitor()
	client.Bind("tcp://0.0.0.0:9442")
	defer client.Close()
	var size int64
	for i := 0; i < t.N; i++ {
		client.Send("ceptop", zmq4.SNDMORE)
		_, err := client.Send(urlData, zmq4.DONTWAIT)
		if err != nil {
			logrus.Error("Send Error:", err)
		}
		size++
	}
	logrus.Info(size)
}

func monitor() {
	mo, _ := zmq4.NewSocket(zmq4.PAIR)
	mo.Connect("inproc://monitor.rep")
	for {
		a, b, c, err := mo.RecvEvent(0)
		if err != nil {
			logrus.Error(err)
			return
		}
		logrus.Infof("A:%s B:%s C:%d", a, b, c)
	}

}