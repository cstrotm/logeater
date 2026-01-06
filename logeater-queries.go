package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
	"strings"
)

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

func main() {
	listdomains := flag.Bool("d", false, "list domain names")
	listtld := flag.Bool("1", false, "list top level domains")
	listqueryip := flag.Bool("i", false, "list query IP addresses")
	listqueryclass := flag.Bool("c", false, "list query network classes")
	listquerytype := flag.Bool("t", false, "list query type")
	liststats := flag.Bool("s", false, "list statistics")
	nolookup := flag.Bool("n", false, "no reverse IP lookup")
	flag.Parse()

	querynames := make(map[string]int)
	querytld := make(map[string]int)
	queryips := make(map[string]int)
	queryclasses := make(map[string]int)
	querytypes := make(map[string]int)

	qnum := 0
	qrecursive := 0
	qiterative := 0
	qedns := 0
	qdnssecok := 0
	qcd := 0
	qsigned := 0
	qtcp := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		qnum++

		// collect query names
		i := strings.Index(txt, "query: ")
		qname := strings.SplitN(txt[i+7:], " ", 5)
		querynames[strings.ToLower(qname[0])]++
		tld := qname[0][strings.LastIndex(qname[0], ".")+1:]
		querytld[strings.ToLower(tld)]++

		// collect DNS query classes
		queryclasses[qname[1]]++

		// collect DNS query types
		querytypes[qname[2]]++

		// collect DNS query flags
		qflags := qname[3]
		if qflags[:1] == "-" {
			qiterative++
		} else {
			qrecursive++
		}

		if strings.ContainsRune(qflags,'E') {
			qedns++
		}
		if strings.ContainsRune(qflags,'D') {
			qdnssecok++
		}
		if strings.ContainsRune(qflags,'C') {
			qcd++
		}
		if strings.ContainsRune(qflags,'S') {
			qsigned++
		}
		if strings.ContainsRune(qflags,'T') {
			qtcp++
		}

		// collect IP address of clients
		i = strings.Index(txt, "client @")
		ip := strings.SplitN(txt[i+19:], "#", 2)
		queryips[strings.ToLower(ip[0])]++

	}

	if *listdomains {
		fmt.Println("Query-Domain-Names\n")

		for _, res := range sortedKeys(querynames) {
			fmt.Println(querynames[res], ":", res)
		}
	}

	if *listtld {
		fmt.Println("Query-TLD-Names\n")

		for _, res := range sortedKeys(querytld) {
			fmt.Println(querytld[res], ":", res)
		}
	}

	if *listqueryclass {
		fmt.Println("Query-Network-Classes\n")

		for _, class := range sortedKeys(queryclasses) {
			fmt.Println(queryclasses[class], ":", class)
		}
	}

	if *listquerytype {
		fmt.Println("Query-Network-Types\n")

		for _, qtype := range sortedKeys(querytypes) {
			fmt.Println(querytypes[qtype], ":", qtype)
		}
	}

	if *listqueryip {
		fmt.Println("Query-IP-Addresses\n")

		for _, ip := range sortedKeys(queryips) {
			num := queryips[ip]
			if *nolookup == false {
				name, _ := net.LookupAddr(ip)
				fmt.Println(num, ":", ip, name)
			} else {
				fmt.Println(num, ":", ip)
			}
		}
	}

	if *liststats {
		fmt.Println("Query-Statistics\n")

		fmt.Println(qnum,": total queries ( 100 % )")
		fmt.Println(qiterative,": iterative queries (", qiterative*100/qnum,"% )")
		fmt.Println(qrecursive,": recursive queries (", qrecursive*100/qnum," % )")
		fmt.Println(qtcp,": queries over TCP (", qtcp*100/qnum," % )")
		fmt.Println(qedns,": queries with EDNS support (", qedns*100/qnum," % )")
		fmt.Println(qdnssecok,": queries indicate DNSSEC support (", qdnssecok*100/qnum," % )")
		fmt.Println(qcd,": queries with DNSSEC validation disabled (CD-flag) (", qcd*100/qnum," % )")
		fmt.Println(qsigned,": queries TSIG signed (", qsigned*100/qnum," % )")
	}

}
