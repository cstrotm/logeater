# logeater
Log-Aggregation Tools for BIND 9 logs


Example use of ```logeater-queries```:
=====================================

Statistics:

```
cat query.n | ./logeater-queries -s | column -t -s ":"
Query-Statistics
21385              total queries ( 100 % )
20471              iterative queries ( 95 % )
914                recursive queries ( 4  % )
863                queries over TCP ( 4  % )
16987              queries with EDNS support ( 79  % )
15197              queries indicate DNSSEC support ( 71  % )
8804               queries with DNSSEC validation disabled (CD-flag) ( 41  % )
1571               queries TSIG signed ( 7  % )
```

Example - prining top 10 of TLDs used in queries

```
# ./logeater-queries -1 < query.log | head
Query-TLD-Names

324632 : net
214054 : bind
16751 : de
10147 : com
4764 : org
1188 : host
271 : social
197 : arpa
```

Example - printing network classes and query types of queries:

```
$ cat query.log | ./logeater-queries -c -t | column -t
Query-Network-Classes
21379                  :  IN
6                      :  CH

Query-Network-Types
8514                   :  A
4060                   :  AAAA
3079                   :  SOA
2372                   :  DNSKEY
927                    :  PTR
658                    :  MX
543                    :  NS
312                    :  DS
286                    :  TXT
186                    :  NSEC
129                    :  ANY
115                    :  CNAME
```

Example - printing the top ten query IP addresses with reverse name
resolution (can be slow):

```
$ cat query.log | ./logeater-queries -i | head | column -t

Query-IP-Addresses
1571                :  212.114.206.217                   [muc.example.de.]
821                 :  2620:74:13:4400::41               [dnsviz-db.verisignlabs.com.]
794                 :  72.13.58.112                      [dnsviz-db.verisignlabs.com.]
704                 :  54.234.42.241                     [241.compute-1.amazonaws.com.]
682                 :  2001:19f0:5001:df:76d7:5703:ba0a:e220  []
565                 :  185.92.221.212                    [185.92.221.212.vultr.com.]
467                 :  185.22.143.29                     [b9168f1d.cgn.dg-w.de.]
314                 :  91.51.184.46                      [3b82e.dip0.t-ipconnect.de.]
```

```logeater-dnssec``` analyses the a log file with messages from the
"DNSSEC" category and groups the error messages :

```
$ cat dnssec.log | ./logeater-dnssec | head
8727 : 0C9F6LGOE6NADAS8KG1CLIK9UO9G7EIG.ad/NSEC3: no valid signature found
6953 : ad/SOA: no valid signature found
3976 : sat-tv.com/A: got insecure response; parent indicates it should be secure
1730 : mozilla.com/SOA: no valid signature found
1586 : stream.bestvideostreaming.is/A: no valid signature found
1577 : 8FC1DQ3C2Q3ERFD4UO40ENDBTSFME5JO5.ad/NSEC3: no valid signature found
1576 : sat-tv.com/SOA: got insecure response; parent indicates it should be secure
1576 : cdws.eu-west-1.amazonaws.com.Cisco/AAAA: bad cache hit (amazonaws.com.Cisco/DS)
1483 : 0c9f6lgoe6n13ad9iu1clik9uo9g7eig.ad/NSEC3: no valid signature found
968 : cbr.de/NSEC: no valid signature found
```

```logeater-resolver``` analyses the a log file with messages from the
"resolver" category and groups the error messages:

```
$ cat resolvers.log | ./logeater-resolvers | head
42908 : s-cnc1.qq.com/AAAA: Name qq.com (SOA) not subdomain of zone ns-cnc1.qq.com -- invalid response
42713 : s-tel1.qq.com/AAAA: Name qq.com (SOA) not subdomain of zone ns-tel1.qq.com -- invalid response
42484 : s-os1.qq.com/AAAA: Name qq.com (SOA) not subdomain of zone ns-os1.qq.com -- invalid response
42297 : s-cmn1.qq.com/AAAA: Name qq.com (SOA) not subdomain of zone ns-cmn1.qq.com -- invalid response
20346 : mails.sonymusicfans.com/DS: invalid response
10920 : tp1.glb.nist.gov/DS: invalid response
9693 : media.netd.com.tr/AAAA for client 192.0.2.165#3347: Name netd.com.tr (SOA) not subdomain of zone media.netd.com.tr -- invalid response
7932 : service.superc.net/AAAA for client 192.0.2.11#3073: Name superc.net (SOA) not subdomain of zone service.superc.net — invalid response
04597 : brickleonavon.com/NS for client 192.0.2.46#3073: Name . (SOA) not subdomain of zone brickleonavon.com -- invalid response
4474 : promo.mobile.de/AAAA for client 2001:db8:1800:88:78f9:ba4:45fe:d438#48296: Name mobile.de (SOA) not subdomain of zone promo.mobile.de -- invalid response
```

Compiling from source
=====================

Install ```git``` (https://git-scm.com/) and the ```go```
(https://golang.org) programming language. Both are available in the
repositories of most Linux and BSD distributions.

Download the source

```
git clone https://github.com/menandmice-services/logeater.git
```

Compile the source

```
cd logeater
go build logeater-queries.go
go build logeater-resolvers.go
go build logeater-dnssec.go
```

BIND 9 versions
===============

Tested with BIND 9.10 - BIND 9.20.
