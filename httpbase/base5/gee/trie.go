package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 是否是一个完整的 URL ，不是则为 ""
	part     string  // URL 块值，用 / 分割的部分，eg：/abe/123， abd 和 123 就是两个 part
	children []*node // 子节点，例如[doc, tutorial,intro]
	isWild   bool    // 是否精确匹配，比如：filename 或 *filename 这样的 node 就为 true
}

// 找到匹配的子节点，用在插入场景，找到 1 个匹配的就立即返回
func (n *node) matchChild(part string) *node {
	println("匹配的子节点 " + n.String())
	// 遍历 n 节点的所有子节点，看是否能找到匹配的子节点，返回
	for _, child := range n.children {
		// 如果有模糊匹配也会成功匹配上
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
			fmt.Printf("[matchChildren] node1: %s\n", nodes)
		}
	}
	fmt.Printf("[matchChildren] node2: %s\n", nodes)
	return nodes
}

// 一边匹配一边插入
func (n *node) insert(pattern string, parts []string, height int) {
	fmt.Printf("[insert] node: %s\n", n.String())
	fmt.Printf("[insert 1] pattern: %s\n", pattern)
	fmt.Printf("[insert 2] parts: %s\n", parts)
	fmt.Printf("[insert 3] height: %d\n", height)
	if len(parts) == height {
		// 如果已经匹配完，那么将 pattern 赋值给该 node ，表示它是一个完整的 url
		// 这是递归的终止条件
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// 如果没有匹配上，那么进行生成，放到 n 节点的子列表中
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		fmt.Printf("[search ] %s\n", n)
		return n
	}

	part := parts[height]
	fmt.Printf("part %s\n", part)
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}

	for _, child := range n.children {
		child.travel(list)
	}
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t, children%s}", n.pattern, n.part, n.isWild, n.children)
}
