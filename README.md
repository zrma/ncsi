# ncsi
simple ncsi response server

----

윈도우에서는 NCSI(Network Connection Status Indicator)를 통해서 네트워크가 인터넷에 연결 되어 있는지 판단합니다.

윈도우 10 version 1607 이전
- http://www.msftncsi.com/ncsi.txt
- http://ipv6.msftncsi.com/ncsi.txt

`Microsoft NCSI`

문자열 회신 성공

윈도우 10 version 1607 이후
- http://www.msftconnecttest.com/connecttest.txt
- http://ipv6.msftconnecttest.com/connecttest.txt

`Microsoft Connect Test`

문자열 회신 성공

dns.msftncsi.com의 FQDN 확인 성공까지 마친 경우에 인터넷 연결 성공으로 판단합니다.

----

ref1: [Windows でインターネット接続しているのに「インターネットなし」と表示される](https://hebikuzure.wordpress.com/2018/09/26/network-connection-status-icon/)\
ref2: [An Internet Explorer or Edge window opens when your computer connects to a corporate network or a public network](https://docs.microsoft.com/en-US/troubleshoot/windows-client/networking/internet-explorer-edge-open-connect-corporate-public-network) - [NCSI active probes and the network status alert](https://docs.microsoft.com/en-US/troubleshoot/windows-client/networking/internet-explorer-edge-open-connect-corporate-public-network#ncsi-active-probes-and-the-network-status-alert)
