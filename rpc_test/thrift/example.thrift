namespace go echo

struct EchoReq {
	1: string msg;
	2: i32  tag;
}

struct EchoRes {
	1: string msg;
	2: i32 tag
}

service Echo {
	EchoRes echo(1: EchoReq req);
}