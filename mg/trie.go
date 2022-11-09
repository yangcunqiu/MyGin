package mg

import "strings"

// node 前缀树节点
type node struct {
	pattern  string  // 路由规则, 例 /p/:lang
	part     string  // 路由中的一部分, 例 :lang
	children []*node // 子节点
	isWild   bool    // 是否是通配符
}

// insert 插入节点
func (n *node) insert(pattern string, parts []string, height int) {
	// 最后一个节点
	if len(parts) == height {
		// 只有最后的叶子节点的pattern才赋值
		n.pattern = pattern
		return
	}

	part := parts[height]
	// 查找当前节点下的子节点
	child := n.matchChild(part)
	if child == nil {
		// 新建子节点
		child = &node{
			part:   part,
			isWild: strings.HasPrefix(part, ":") || strings.HasPrefix(part, "*"),
		}
		n.children = append(n.children, child)
	}
	// 继续给子节点继续往下添加子节点
	child.insert(pattern, parts, height+1)
}

// search 查询节点
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

// 第一个成功匹配的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 匹配所有节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
