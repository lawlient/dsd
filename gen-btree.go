package main

import (
	"bufio"
	"fmt"
	"strings"
)

type bTreeGen struct {
	w   *strings.Builder
	l2n map[int][]*bnode
}

func init() {
	G["btree"] = func(w *strings.Builder) Generator {
		return &bTreeGen{
			w:   w,
			l2n: make(map[int][]*bnode, 0),
		}
	}
}

func (g *bTreeGen) Generate(req string) error {
	w := g.w
	w.WriteString("digraph BTree {\n")
	w.WriteString("    splines=false\n")

	/* bnode */
	scanner := bufio.NewScanner(strings.NewReader(req))
	for scanner.Scan() {
		line := scanner.Text()
		n := newBNode(g, line)
		if n == nil {
			continue
		}

		n.dotnode()
	}

	/* edge */
	for d, li := range g.l2n {
		if d+1 == len(g.l2n) {
			/* leaf */
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

type bnode struct {
	g     *bTreeGen
	items []string
	level int
	order int
}

func (n *bnode) name() string {
	if n.level == 0 {
		return "\"root\""
	}
	return fmt.Sprintf("\"node-%d-%d\"", n.level, n.order)
}

func (n *bnode) childName(i, first int) string {
	if i > len(n.items) {
		return "\"null\""
	}
	return fmt.Sprintf("\"node-%d-%d\"", n.level+1, i+first)
}

func newBNode(g *bTreeGen, l string) *bnode {
	indent := 0
	for _, c := range l {
		if c != ' ' {
			break
		}
		indent++
	}
	level := indent / 4
	n := &bnode{
		g:     g,
		items: strings.Split(l[indent:], " "),
		level: level,
		order: len(g.l2n[level]),
	}
	g.l2n[level] = append(g.l2n[level], n)
	return n
}

func (n *bnode) dotnode() {
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

func (n *bnode) dotedge() {
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
