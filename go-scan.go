package main

import (
	"net"
	"fmt"
	"flag"
	"strconv"
	"time"
	"strings"
)

var ports [45]int


func main() {
	c := make(chan int); 
		
	var jobs []string
	
	scanPort := func (ip string, timeout int64, openonly bool) {
		conn, err := net.DialTimeout("tcp", ip, time.Duration(timeout*1000000))
		parts := strings.Split(ip, ":")
		if err != nil {
			fmt.Println(parts[0]+"\t"+parts[1]+"\tF\t"+err.Error())
		} else {
			fmt.Println(parts[0]+"\t"+parts[1]+"\tS")
			conn.Close()
		}
	}
	
	scanSlice := func (seg []string, timeout int64, openonly bool) {
		for i:=0; i<len(seg); i++ {
			scanPort(seg[i], timeout, openonly)
		}
		c <- 0 //tell main we are done with this slice
	}
	
	ports[	0	] = 7;
	ports[	1	] =	18	;
	ports[	2	] =	20	;
	ports[	3	] =	21	;
	ports[	4	] =	22	;
	ports[	5	] =	23	;
	ports[	6	] =	25	;
	ports[	7	] =	29	;
	ports[	8	] =	37	;
	ports[	9	] =	42	;
	ports[	10	] =	43	;
	ports[	11	] =	49	;
	ports[	12	] =	53	;
	ports[	13	] =	69	;
	ports[	14	] =	70	;
	ports[	15	] =	79	;
	ports[	16	] =	80	;
	ports[	17	] =	103	;
	ports[	18	] =	108	;
	ports[	19	] =	109	;
	ports[	20	] =	110	;
	ports[	21	] =	115	;
	ports[	22	] =	118	;
	ports[	23	] =	119	;
	ports[	24	] =	137	;
	ports[	25	] =	139	;
	ports[	26	] =	143	;
	ports[	27	] =	150	;
	ports[	28	] =	156	;
	ports[	29	] =	161	;
	ports[	30	] =	179	;
	ports[	31	] =	190	;
	ports[	32	] =	194	;
	ports[	33	] =	197	;
	ports[	34	] =	389	;
	ports[	35	] =	396	;
	ports[	36	] =	443	;
	ports[	37	] =	444	;
	ports[	38	] =	445	;
	ports[	39	] =	458	;
	ports[	40	] =	546	;
	ports[	41	] =	547	;
	ports[	42	] =	563	;
	ports[	43	] =	569	;
	ports[	44	] =	1080;
	
	
	flag.Parse()
	args := flag.Args()
	
	ip := args[0]
	timeout, _ := strconv.ParseInt(args[1], 0, 64)
	threads, _ := strconv.ParseInt(args[2], 0, 64)
	openonly := (len(args) > 3)
	var _ = ip 
	var _ = timeout 
	var _ = threads
	var _ = openonly
	
	
	if ip[len(ip)-1]=='.' {
		if len(strings.Split(ip, "."))==3 {
			for k:=1; k<255; k++ {
				for j:=1; j<255; j++ {
					for i := 0; i < len(ports); i++ {
						ips := ip+strconv.Itoa(k)+"."+strconv.Itoa(j)+":"+strconv.Itoa(ports[i])
						jobs = append(jobs, ips)
					}
				}
			}
		} else {			
			for j:=1; j<255; j++ {
				for i := 0; i < len(ports); i++ {
					ips := ip+strconv.Itoa(j)+":"+strconv.Itoa(ports[i])
					jobs = append(jobs, ips)
				}
			}
		}
	} else {
		//(new Thread(new ScanSeg(ip, ports, i, threads, timeout, openonly))).start();
		for i := 0; i < len(ports); i++ {
			ips := ip+":"+strconv.Itoa(ports[i])
			jobs = append(jobs, ips)
		}
	}
	
	//create routines with their job slices
	segsize := len(jobs) / int(threads)
	for i:=0; int64(i)<threads; i++ {
		var jobpart []string
		
		if int64(i)==threads-1 {
			jobpart = jobs[segsize*i:]
		} else {
			jobpart = jobs[segsize*i:segsize*(i+1)]
		}
		
		go scanSlice(jobpart, timeout, openonly)
	}
	
	//wait for all routines to finish
	for i:=0; int64(i)<threads; i++ {
		<- c
	}
}



