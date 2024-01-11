package goo

import (
	"fmt"
	"strings"
)

type node struct {
	pattern string // 待匹配路由，例如 /p/:lang
	part string // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild bool // part 含有 : 或 * 时为true
}

func (n *node) matchChild(part string, isWild bool) *node {
	for _, child := range n.children {
		// isWild 为 true 时，仅匹配模糊节点
		if child.part == part && !isWild || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part {
			nodes = append(nodes, child)
		}
	}
	// isWild 节点匹配优先度低
	for _, child := range n.children {
		if child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) String() string  {
	return n.pattern
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]

	// 不允许同一层出现两个模糊匹配节点
	if part[0] == ':' || part[0] == '*' {
		child := n.matchChild(part, true)
		fmt.Println(pattern)
		if child != nil {
			fmt.Println(child.part, part)
		}
		if child != nil && child.part != part {
			panic(fmt.Sprintf("now %s(in %s) is conflict with %s", part, pattern, child.part))
		}
	}
	child := n.matchChild(part, false)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node  {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children{
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}