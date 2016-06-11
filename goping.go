package main
import (
    "bytes"
 //   "container/list"
    "encoding/binary"
    "fmt"
    "net"
  //  "os"
    "time"
)
//检验和算法
func CheckSum(data []byte) uint16{
	var(
		sum uint32
		length int=len(data)
		index int
	)
	for length >1{
	sum+=uint32(data[index])<<8+uint32(data[index+1])
	index +=2
	length -=2
	}
	if length>0{
		sum+=uint32(data[index])
	}
	sum +=(sum>>16)
	return uint16(^sum)
}


type ICMP struct {
	Type    uint8
	Code    uint8
	Checksum  uint16
	Identifier uint16                
	SequenceNum uint16
}          /*定义ICMP报文体*/

func goping(deviceip string) (errip int){
var(
	icmp ICMP                   
	laddr = net.IPAddr{IP:net.ParseIP("0.0.0.0")} //源地址laddr可以是0.0.0.0也可以是自己的ip，这个并不影响ICMP的工作。
	raddr,_ = net.ResolveIPAddr("ip",deviceip)
	//目的地址raddr是一个URL，这里使用Resolve进行DNS解析，注意返回值是一个指针，所以下面的DialIP方法中参数表示没有取地址符。
)			
conn,err := net.DialIP("ip4:icmp",&laddr,raddr)  //DialIP在网络协议netProto上连接本地地址laddr和远端地址raddr，netProto必须是"ip"、"ip4"或"ip6"后跟冒号和协议名或协议号。
if err!=nil{
	fmt.Println(err.Error())
	return 1                           //diaip失败
}
defer conn.Close()

//构造ICMP报文
icmp.Type = 8
	icmp.Code =0
	icmp.Checksum=0
	icmp.Identifier=0
	icmp.SequenceNum=0

	var buffer bytes.Buffer
	binary.Write(&buffer,binary.BigEndian,icmp)
	icmp.Checksum=CheckSum(buffer.Bytes())
	buffer.Reset()
	binary.Write(&buffer,binary.BigEndian,icmp)

//Ping的Request过程
//fmt.Printf("\n正在 Ping %s 具有 0 字节的数据:\n",raddr.String())
recv :=make([]byte,1024)
//statistic :=list.New()
sended_packets:=0

for i:=1;i>0;i--{
	if  _,err:=conn.Write(buffer.Bytes());err!=nil{  //conn.Write方法执行之后也就发送了一条ICMP请求，同时进行计时和计次。
	fmt.Println(err.Error())
	return 1
	}
	sended_packets++
//	t_start :=time.Now() 
	conn.SetReadDeadline((time.Now().Add(time.Second *1)))
//SetReadDeadline可以在未收到数据的指定时间内停止Read等待，并返回错误err，然后判定请求超时。否则，收到回应后，计算来回所用时间，并放入一个list方便后续统计。
	_,err:=conn.Read(recv)
	if err!=nil{
//	fmt.Println("请求超时")
	return 1}
/*	t_end :=time.Now()
	dur :=t_end.Sub(t_start).Nanoseconds()/1e6
	fmt.Printf("来自 %s 的回复: 时间 = %dms\n", raddr.String(), dur)
	statistic.PushBack(dur) */

				}
					/*
				
				defer func(){
	fmt.Println("")
	   //信息统计
	var min,max,sum int64
	if statistic.Len()==0{
		min,max,sum=0,0,0
	}else{
		min,max,sum=statistic.Front().Value.(int64),statistic.Front().Value.(int64),int64(0)
	}
	for v:=statistic.Front();v!=nil;v=v.Next(){
		val :=v.Value.(int64)
		switch{
		case val <min:
			min=val
		case val>max:
			max=val
		}
		sum=sum+val
	}
	recved,losted :=statistic.Len(),sended_packets-statistic.Len()
	fmt.Printf("%s 的 Ping 统计信息：\n  数据包：已发送 = %d，已接收 = %d，丢失 = %d (%.1f%% 丢失)，\n往返行程的估计时间(以毫秒为单位)：\n  最短 = %dms，最长 = %dms，平均 = %.0fms\n",
	raddr.String(),
	sended_packets,recved,losted,float32(losted)/float32(sended_packets)*100,
	min,max,float32(sum)/float32(recved),
	)
	}()*/
return 0
	}			
				
				

