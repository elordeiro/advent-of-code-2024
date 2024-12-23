package main

import (
	"adventofcode/utils"
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Set map[string]bool

func newGraph(fileName string) map[string][]string {
	conns := utils.ReadStringSlice(fileName)
	graph := map[string][]string{}
	for _, conn := range conns {
		parts := strings.Split(conn, "-")
		src, dst := parts[0], parts[1]
		sns, dns := graph[src], graph[dst]
		sns = append(sns, dst)
		dns = append(dns, src)
		graph[src], graph[dst] = sns, dns
	}
	return graph
}

func solve(fileName string) string {
	graph := newGraph(fileName)

	R := Set{}
	P := Set{}
	X := Set{}

	for v := range graph {
		P[v] = true
	}

	var Res Set

	var BronKerbosch func(Set, Set, Set)
	BronKerbosch = func(R, P, X Set) {
		if len(P) == 0 && len(X) == 0 {
			if len(R) > len(Res) {
				Res = Set{}
				maps.Copy(Res, R)
			}
			return
		}
		for v := range P {
			PN := Set{}
			XN := Set{}
			for _, n := range graph[v] {
				if P[n] {
					PN[n] = true
				}
				if X[n] {
					XN[n] = true
				}
			}
			R[v] = true
			BronKerbosch(R, PN, XN)
			delete(R, v)
			P[v] = false
			X[v] = true
		}
	}

	BronKerbosch(R, P, X)

	res := slices.Collect(maps.Keys(Res))
	slices.Sort(res)
	return strings.Join(res, ",")
}

func main() {
	tests := []struct {
		fileName string
		want     string
	}{
		{"../test1.txt", "co,de,ka,ta"},
		{"../input.txt", "bg,bu,ce,ga,hw,jw,nf,nt,ox,tj,uu,vk,wp"},
	}

	for _, test := range tests {
		got := solve(test.fileName)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %s, Want %s\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %s\n", test.fileName, got)
	}
}
