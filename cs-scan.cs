using System;
using System.Collections.Generic;
using System.Text;

namespace cs_scan
{
    class Program
    {
        static int[] ports;

        static void Main(string[] args)
        {
            ports = new int[45];
            ports[0] = 7;
            ports[1] = 18;
            ports[2] = 20;
            ports[3] = 21;
            ports[4] = 22;
            ports[5] = 23;
            ports[6] = 25;
            ports[7] = 29;
            ports[8] = 37;
            ports[9] = 42;
            ports[10] = 43;
            ports[11] = 49;
            ports[12] = 53;
            ports[13] = 69;
            ports[14] = 70;
            ports[15] = 79;
            ports[16] = 80;
            ports[17] = 103;
            ports[18] = 108;
            ports[19] = 109;
            ports[20] = 110;
            ports[21] = 115;
            ports[22] = 118;
            ports[23] = 119;
            ports[24] = 137;
            ports[25] = 139;
            ports[26] = 143;
            ports[27] = 150;
            ports[28] = 156;
            ports[29] = 161;
            ports[30] = 179;
            ports[31] = 190;
            ports[32] = 194;
            ports[33] = 197;
            ports[34] = 389;
            ports[35] = 396;
            ports[36] = 443;
            ports[37] = 444;
            ports[38] = 445;
            ports[39] = 458;
            ports[40] = 546;
            ports[41] = 547;
            ports[42] = 563;
            ports[43] = 569;
            ports[44] = 1080;


            string ip = args[0];
            int timeout = Convert.ToInt32(args[1]);
            int threads = Convert.ToInt32(args[2]);
            bool openonly = (args.Length > 3);

            if (ip[ip.Length - 1] == '.')
            {
                char[] spl = new char[1];
                spl[0] = '.';

                if (ip.Split(spl).Length == 2)
                {
                    for (int k = 1; k < 255; k++)
                    {
                        for (int j = 1; j < 255; j++)
                        {
                            string ips = ip + k + "." + j;
                            for (int i = 0; i < threads; i++)
                            {
                                //threadSleep(threads, timeout);
                                //(new Thread(new ScanSeg(ips, ports, i, threads, timeout, openonly))).start();
                            }
                        }
                    }
                }
                else
                {
                    for (int j = 1; j < 255; j++)
                    {
                        string ips = ip + j;
                        for (int i = 0; i < threads; i++)
                        {
                            //threadSleep(threads, timeout);
                            //(new Thread(new ScanSeg(ips, ports, i, threads, timeout, openonly))).start();
                        }
                    }
                }
            }
            else
                for (int i = 0; i < threads; i++)
                {
                    //(new Thread(new ScanSeg(ip, ports, i, threads, timeout, openonly))).start();
                }
		
            Console.In.ReadLine();
        }
    }
}
