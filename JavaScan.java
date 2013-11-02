
import java.net.*;
import java.io.*;


public class JavaScan {

	public static int[] ports;

    public static void main(String[] args) {
	
		ports = new int[45];
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
	
		
		String ip = args[0];
		int timeout = Integer.parseInt(args[1]);
        int threads = Integer.parseInt(args[2]);
		boolean openonly = (args.length > 3);

		if (ip.charAt(ip.length()-1)=='.') {
			for (int j=1; j<255; j++) {
				String ips = ip+j;
				for (int i=0; i<threads; i++)
					(new Thread(new ScanSeg(ips, ports, i, threads, timeout, openonly))).start();
			}
		}
		else
			for (int i=0; i<threads; i++)
				(new Thread(new ScanSeg(ip, ports, i, threads, timeout, openonly))).start();
		
    }

}


class ScanSeg implements Runnable {

	public int starti;
	public int increment;
	public int timeout;
	public int[] ports;
	String ip;
	boolean openonly;

	public ScanSeg(String ipaddr, int[] portA, int start, int inc, int timeo, boolean openon) 
	{
		this.openonly=openon;
		this.ip = ipaddr;
		this.ports = portA;
		this.starti=start;
		this.timeout=timeo;
		this.increment=inc;
	}
	
    public void run() {
        for (int i=starti; i<ports.length; i+=increment)
			trySocket(ip, ports[i], timeout);
    }

	public void trySocket(String ip, int port, int timeout)
	{
		//Create socket connection
		try {
			Socket socket=new Socket();   
			socket.connect(new InetSocketAddress(ip,port),timeout);
			socket.close();
			System.out.println(ip+"\t"+port+"\tS");
		} catch (Exception e) {
			if (!openonly)
				System.out.println(ip+"\t"+port+"\tF\t"+e.toString());
		}
	}
	
}
	
