package pkg

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"os/exec"
	"runtime"
	"time"
)

const (
	UnixStaTimestamp = 2208988800
)

/**
NTP协议 http://www.ntp.org/documentation.html
@author mengdj@outlook.com
*/
type Ntp struct {
	//1:32bits
	Li        uint8 //2 bits
	Vn        uint8 //3 bits
	Mode      uint8 //3 bits
	Stratum   uint8
	Poll      uint8
	Precision uint8
	//2:
	RootDelay           int32
	RootDispersion      int32
	ReferenceIdentifier int32
	//64位时间戳
	ReferenceTimestamp uint64 //指示系统时钟最后一次校准的时间
	OriginateTimestamp uint64 //指示客户向服务器发起请求的时间
	ReceiveTimestamp   uint64 //指服务器收到客户请求的时间
	TransmitTimestamp  uint64 //指示服务器向客户发时间戳的时间
}

func NewNtp() (p *Ntp) {
	//其他参数通常都是服务器返回的
	p = &Ntp{Li: 0, Vn: 3, Mode: 3, Stratum: 0}
	return p
}

/**
构建NTP协议信息
*/
func (e *Ntp) GetBytes() []byte {
	//注意网络上使用的是大端字节排序
	buf := &bytes.Buffer{}
	head := (e.Li << 6) | (e.Vn << 3) | ((e.Mode << 5) >> 5)
	_ = binary.Write(buf, binary.BigEndian, uint8(head))
	_ = binary.Write(buf, binary.BigEndian, e.Stratum)
	_ = binary.Write(buf, binary.BigEndian, e.Poll)
	_ = binary.Write(buf, binary.BigEndian, e.Precision)
	//写入其他字节数据
	_ = binary.Write(buf, binary.BigEndian, e.RootDelay)
	_ = binary.Write(buf, binary.BigEndian, e.RootDispersion)
	_ = binary.Write(buf, binary.BigEndian, e.ReferenceIdentifier)
	_ = binary.Write(buf, binary.BigEndian, e.ReferenceTimestamp)
	_ = binary.Write(buf, binary.BigEndian, e.OriginateTimestamp)
	_ = binary.Write(buf, binary.BigEndian, e.ReceiveTimestamp)
	_ = binary.Write(buf, binary.BigEndian, e.TransmitTimestamp)
	//[27 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	return buf.Bytes()
}

func (e *Ntp) Parse(bf []byte, useUnixSec bool) {
	var (
		bit8  uint8
		bit32 int32
		bit64 uint64
		rb    *bytes.Reader
	)
	//貌似这binary.Read只能顺序读，不能跳着读，想要跳着读只能使用切片bf
	rb = bytes.NewReader(bf)
	_ = binary.Read(rb, binary.BigEndian, &bit8)
	//向右偏移6位得到前两位LI即可
	e.Li = bit8 >> 6
	//向右偏移2位,向右偏移5位,得到前中间3位
	e.Vn = (bit8 << 2) >> 5
	//向左偏移5位，然后右偏移5位得到最后3位
	e.Mode = (bit8 << 5) >> 5
	_ = binary.Read(rb, binary.BigEndian, &bit8)
	e.Stratum = bit8
	_ = binary.Read(rb, binary.BigEndian, &bit8)
	e.Poll = bit8
	_ = binary.Read(rb, binary.BigEndian, &bit8)
	e.Precision = bit8

	//32bits
	_ = binary.Read(rb, binary.BigEndian, &bit32)
	e.RootDelay = bit32
	_ = binary.Read(rb, binary.BigEndian, &bit32)
	e.RootDispersion = bit32
	_ = binary.Read(rb, binary.BigEndian, &bit32)
	e.ReferenceIdentifier = bit32

	//以下几个字段都是64位时间戳(NTP都是64位的时间戳)
	_ = binary.Read(rb, binary.BigEndian, &bit64)
	e.ReferenceTimestamp = bit64
	_ = binary.Read(rb, binary.BigEndian, &bit64)
	e.OriginateTimestamp = bit64
	_ = binary.Read(rb, binary.BigEndian, &bit64)
	e.ReceiveTimestamp = bit64
	_ = binary.Read(rb, binary.BigEndian, &bit64)
	e.TransmitTimestamp = bit64
	//转换为unix时间戳,先左偏移32位拿到64位时间戳的整数部分，然后ntp的起始时间戳 1900年1月1日 0时0分0秒 2208988800
	if useUnixSec {
		e.ReferenceTimestamp = (e.ReceiveTimestamp >> 32) - UnixStaTimestamp
		if e.OriginateTimestamp > 0 {
			e.OriginateTimestamp = (e.OriginateTimestamp >> 32) - UnixStaTimestamp
		}
		e.ReceiveTimestamp = (e.ReceiveTimestamp >> 32) - UnixStaTimestamp
		e.TransmitTimestamp = (e.TransmitTimestamp >> 32) - UnixStaTimestamp
	}
}

/**
 * @SynchronizationTime   同步时间
 */

// @Tags NTP
// @Summary 同步时间
// @Security SynchronizationTime
// @Success serverUrl string 服务端地址
func SynchronizationTime(serverUrl string) {
	// 初始化
	ntp := NewNtp()
	buffer := make([]byte, 2048)
	if serverUrl == "" {
		serverUrl = "ntp1.aliyun.com:123"
	}
	if conn, err := net.Dial("udp", serverUrl); err == nil {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
			conn.Close()
		}()
		// 写初始化
		_, _ = conn.Write(ntp.GetBytes())
		if ret, err := conn.Read(buffer); err == nil && ret > 0 {
			ntp.Parse(buffer, true)
			// 转换时间戳
			t := time.Unix(int64(ntp.ReferenceTimestamp), 0)
			switch runtime.GOOS {
			case "windows":
				hms := t.Format("15:04:05")
				_ = exec.Command("cmd", "/C", "time", hms).Run()
				date := t.Format("2006-01-02")
				_ = exec.Command("cmd", "/C", "date", date).Run()
			default:
				date := t.Format("2006-01-02 01:01:01")
				_ = exec.Command("bash", "-c", "date -s "+date).Run()
			}
		}
	}
}
