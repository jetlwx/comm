package comm

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type PingResult struct {
	Host               string
	PacketsTransmitted int //发出个数和
	PacketsReceived    int //收到个数
	PacketLoss         int //丢包个数
	PacketLossPercent  int //丢包率
	RoundTripMin       int //最短时间
	RoundTripAvg       int //平均时间
	RoundTripMax       int //最大时间
}

//host:ing 指定的主机
//count:要发送的回显请求数。
//size ;要发送缓冲区大小。建议size=32
//timeout :"等待每次回复的超时时间(毫秒)。"
//neverstop""Ping 指定的主机，直到停止。" ,默认false
//showdetail:是否显示详细 信息
func Ping(host string, count int, size int, showdetail bool) (res PingResult) {
	timeout := 1000
	neverstop := false
	cname, _ := net.LookupCNAME(host)
	starttime := time.Now()
	conn, err := net.DialTimeout("ip4:icmp", host, time.Duration(timeout*1000*1000))
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
		}
	}()
	ip := conn.RemoteAddr()
	// if err != nil {
	// 	log.Println("err--->", err)
	// }

	if showdetail {
		fmt.Println("正在 Ping " + cname + " [" + ip.String() + "] 具有 32 字节的数据:")
	}
	var seq int16 = 1
	id0, id1 := genidentifier(host)
	const ECHO_REQUEST_HEAD_LEN = 8

	sendN := 0
	recvN := 0
	lostN := 0
	shortT := -1
	longT := -1
	sumT := 0

	for count > 0 || neverstop {
		sendN++
		var msg []byte = make([]byte, size+ECHO_REQUEST_HEAD_LEN)
		msg[0] = 8                        // echo
		msg[1] = 0                        // code 0
		msg[2] = 0                        // checksum
		msg[3] = 0                        // checksum
		msg[4], msg[5] = id0, id1         //identifier[0] identifier[1]
		msg[6], msg[7] = gensequence(seq) //sequence[0], sequence[1]

		length := size + ECHO_REQUEST_HEAD_LEN

		check := checkSum(msg[0:length])
		msg[2] = byte(check >> 8)
		msg[3] = byte(check & 255)

		conn, err = net.DialTimeout("ip:icmp", host, time.Duration(timeout*1000*1000))
		if err != nil {
			if showdetail {
				stat(ip.String(), 0, 0, 0, 0, 0, 0, showdetail)
			}
			continue
		}
		//checkError(err)

		starttime = time.Now()
		conn.SetDeadline(starttime.Add(time.Duration(timeout * 1000 * 1000)))
		_, err = conn.Write(msg[0:length])

		const ECHO_REPLY_HEAD_LEN = 20

		var receive []byte = make([]byte, ECHO_REPLY_HEAD_LEN+length)
		n, err := conn.Read(receive)
		_ = n

		var endduration int = int(int64(time.Since(starttime)) / (1000 * 1000))

		sumT += endduration

		time.Sleep(1000 * 1000 * 1000)

		if err != nil || receive[ECHO_REPLY_HEAD_LEN+4] != msg[4] || receive[ECHO_REPLY_HEAD_LEN+5] != msg[5] || receive[ECHO_REPLY_HEAD_LEN+6] != msg[6] || receive[ECHO_REPLY_HEAD_LEN+7] != msg[7] || endduration >= int(timeout) || receive[ECHO_REPLY_HEAD_LEN] == 11 {
			lostN++
			if showdetail {
				fmt.Println("对 " + cname + "[" + ip.String() + "]" + " 的请求超时。")
			}
		} else {
			if shortT == -1 {
				shortT = endduration
			} else if shortT > endduration {
				shortT = endduration
			}
			if longT == -1 {
				longT = endduration
			} else if longT < endduration {
				longT = endduration
			}
			recvN++
			ttl := int(receive[8])
			if showdetail {
				//			fmt.Println(ttl)
				fmt.Println("来自 " + cname + "[" + ip.String() + "]" + " 的回复: 字节=32 时间=" + strconv.Itoa(endduration) + "ms TTL=" + strconv.Itoa(ttl))
			}
		}

		seq++
		count--
	}

	stat(ip.String(), sendN, lostN, recvN, shortT, longT, sumT, showdetail)
	//	c <- 1
	res.PacketsTransmitted = sendN
	res.PacketLoss = lostN
	res.PacketLossPercent = int(lostN * 100 / sendN)
	res.PacketsReceived = recvN
	res.RoundTripMin = shortT
	res.RoundTripMax = longT
	res.RoundTripAvg = sumT / sendN
	res.Host = ip.String()
	return res
}

func checkSum(msg []byte) uint16 {
	sum := 0

	length := len(msg)
	for i := 0; i < length-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	if length%2 == 1 {
		sum += int(msg[length-1]) * 256 // notice here, why *256?
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

// func checkError(err error) {
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
// 		os.Exit(1)
// 	}
// }

func gensequence(v int16) (byte, byte) {
	ret1 := byte(v >> 8)
	ret2 := byte(v & 255)
	return ret1, ret2
}

func genidentifier(host string) (byte, byte) {
	return host[0], host[1]
}

func stat(ip string, sendN int, lostN int, recvN int, shortT int, longT int, sumT int, showdetail bool) {
	if !showdetail {
		return
	}
	fmt.Println()
	fmt.Println(ip, " 的 Ping 统计信息:")
	fmt.Printf("    数据包: 已发送 = %d，已接收 = %d，丢失 = %d (%d%% 丢失)，\n", sendN, recvN, lostN, int(lostN*100/sendN))
	fmt.Println("往返行程的估计时间(以毫秒为单位):")
	if recvN != 0 {
		fmt.Printf("    最短 = %dms，最长 = %dms，平均 = %dms\n", shortT, longT, sumT/sendN)
	}
	return
}
