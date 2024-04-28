package main

import (
	"bufio"
	"fmt"
	"strings"
)

type bPlusTreeGen struct {
	w   *strings.Builder
	l2n map[int][]*bpnode
}

func init() {
	G["bplustree"] = func(w *strings.Builder) Generator {
		return &bPlusTreeGen{
			w:   w,
			l2n: make(map[int][]*bpnode, 0),
		}
	}
}

func (g *bPlusTreeGen) Generate(req string) error {
	g.w.WriteString("digraph BPlusTree {\n")
	g.w.WriteString("    splines=false\n")

	/* bpnode */
	scanner := bufio.NewScanner(strings.NewReader(req))
	for scanner.Scan() {
		line := scanner.Text()
		n := newBPNode(g, line)
		if n == nil {
			continue
		}

		n.dotnode()
	}

	/* edge */
	for d, li := range g.l2n {
		if d+1 == len(g.l2n) {
			/* leaf */
			for i, n := range li {
				if i+1 == len(li) {
					continue
				}
				g.w.WriteString(fmt.Sprintf("    %s:d%d:e -> %s:d0:w;\n", n.name(), len(n.items)-1, li[i+1].name()))
			}
			continue
		}
		for _, n := range li {
			n.dotedge()
		}
	}

	g.w.WriteString("\n")

	/* rank */
	for _, li := range g.l2n {
		if len(li) < 2 {
			continue
		}
		g.w.WriteString("    {rank=same; ")
		for _, n := range li {
			g.w.WriteString(n.name() + " ")
		}
		g.w.WriteString("}\n")

	}

	g.w.WriteString("}")
	return nil
}

type bpnode struct {
	g     *bPlusTreeGen
	items []string
	level int
	order int
}

func newBPNode(g *bPlusTreeGen, l string) *bpnode {
	indent := 0
	for _, c := range l {
		if c != ' ' {
			break
		}
		indent++
	}
	//  if indent == len(l) {
	//  	/* space line */
	//  	return nil
	//  }
	level := indent / 4
	n := &bpnode{
		g:     g,
		items: strings.Split(l[indent:], " "),
		level: level,
		order: len(g.l2n[level]),
	}
	g.l2n[level] = append(g.l2n[level], n)
	return n
}

func (n *bpnode) name() string {
	if n.level == 0 {
		return "\"root\""
	}
	return fmt.Sprintf("\"node-%d-%d\"", n.level, n.order)
}

func (n *bpnode) childName(i, first int) string {
	if i > len(n.items) {
		return "\"null\""
	}
	return fmt.Sprintf("\"node-%d-%d\"", n.level+1, i+first)
}

func (n *bpnode) dotnode() {
	w := n.g.w
	w.WriteString("\n")
	w.WriteString("    " + n.name() + " [\n")
	w.WriteString("        shape=none\n")
	w.WriteString("        label=<<table border=\"0\" cellborder=\"1\" cellspacing=\"0\" ><tr>\n")
	for i, e := range n.items {
		w.WriteString(fmt.Sprintf("            <td port=\"d%d\">%s</td>\n", i, e))
	}
	w.WriteString("        </tr></table>>\n")
	w.WriteString("    ]\n\n")
	w.WriteString("\n")
}

func (n *bpnode) dotedge() {
	w := n.g.w
	first_child_order := 0
	for _, b := range n.g.l2n[n.level] {
		if b == n {
			break
		}
		first_child_order += len(b.items) + 1
	}
	for i, _ := range n.items {
		w.WriteString(fmt.Sprintf("    %s:d%d:sw -> %s:n\n", n.name(), i, n.childName(i, first_child_order)))
	}
	w.WriteString(fmt.Sprintf("    %s:d%d:se -> %s:n\n", n.name(), len(n.items)-1, n.childName(len(n.items), first_child_order)))
}
