package tree

import (
	"fmt"
	"leetcode/utils"
	"strconv"
	"strings"
)

// leetcode 297 困难

var seqArr []string // 写算法题多用全局变量，可解决很多问题，如避免用闭包
var desArr []string

// Serializes a tree to a single string.
func Serialize(root *utils.TreeNode) string {
	seqArr = make([]string, 0) // 每一次测试用例都需把全局变量置为0
	serialize0(root)
	res := ""
	for i := 0; i < len(seqArr); i++ {
		res += seqArr[i]
	}
	return res
}
func serialize0(root *utils.TreeNode) {
	// 前序
	if root == nil {
		seqArr = append(seqArr, "#")
		seqArr = append(seqArr, ",")
		return
	}
	seqArr = append(seqArr, strconv.Itoa(root.Value))
	seqArr = append(seqArr, ",")

	serialize0(root.LeftNode)
	serialize0(root.RightNode)
}

// Deserializes your encoded data to tree.
func Deserialize(data string) *utils.TreeNode {
	desArr = make([]string, 0)
	desArr = strings.Split(data, ",")
	desArr = desArr[:len(desArr)-1] // 字符串结尾多一个","，因此切片结尾多一个空元素，去除
	return deserialize0()
}
func deserialize0() *utils.TreeNode {
	// 用递归最简单
	if len(desArr) == 0 {
		return nil
	}
	first := removeFirst()
	if first == "#" {
		return nil
	}
	firstInt, err := strconv.Atoi(first)
	if err != nil {
		fmt.Println("err: ", err)
	}

	root := &utils.TreeNode{Value: firstInt}
	root.LeftNode = deserialize0()
	root.RightNode = deserialize0()
	return root
}
func removeFirst() string {
	if len(desArr) == 0 {
		return ""
	}
	res := desArr[0]
	desArr = desArr[1:]
	return res
}
