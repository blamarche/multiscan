
var net = require('net');


var ip = process.argv[2];
var timeout = parseInt(process.argv[3]);
var threads = parseInt(process.argv[4]);
var openonly = (process.argv.length > 5);


function scan()
{	
	if (ip.charAt(ip.length-1)=='.') {
		if (ip.split(".").length==2)
		{
			for (var k=1; k<255; k++) {
				for (var j=1; j<255; j++) {
					var ips = ip+k+"."+j;
					scanIp(ips, timeout);
				}
			}
		}
		else
		{			
			for (var j=1; j<255; j++) {
				var ips = ip+j;
				scanIp(ips, timeout);
			}
		}
	}
	else
		scanIp(ip, timeout);
}

function scanIp(ip, timeout) {
	for (var i=0; i<ports.length; i++) {
		var p = ports[i];
		
		var c = net.connect({port: p, host: ip}, function() {
			console.log(this.ip+"\t"+this.port+"\tS");
			this.end();
		});
		
		c.setTimeout(timeout);
		c.ip = ip;
		c.port = p;
		
		
		c.on('timeout', function() {
			if (!openonly)
				console.log(this.ip+"\t"+this.port+"\tF\tTimeout");
			this.end();
		});
		
		c.on('error', function(err) {
			if (!openonly)
				console.log(this.ip+"\t"+this.port+"\tF\t"+err.toString());
		});
	}
}


var ports = [];
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

scan();