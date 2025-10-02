package parser

import (
	"fmt"
	"sort"
	"strconv"
)

type jsonType int

const (
	jsonTypeNil jsonType = iota
	jsonTypeObject
	jsonTypeArray
	jsonTypeStr
	jsonTypeNum
	jsonTypeBool
)

type JsonNode struct {
	objectVal map[string]JsonNode
	arrayVal  []JsonNode
	strVal    string
	numVal    float64
	boolVal   bool
	nodeType  jsonType
}

func NewObjectNode() JsonNode {
	return JsonNode{objectVal: make(map[string]JsonNode), nodeType: jsonTypeObject}
}

func NewArrayNode() JsonNode {
	return JsonNode{arrayVal: []JsonNode{}, nodeType: jsonTypeArray}
}

func NewStrNode(v string) JsonNode {
	return JsonNode{strVal: v, nodeType: jsonTypeStr}
}

func NewNumNode(v float64) JsonNode {
	return JsonNode{numVal: v, nodeType: jsonTypeNum}
}

func NewBoolNode(v bool) JsonNode {
	return JsonNode{boolVal: v, nodeType: jsonTypeBool}
}

func NewNode() JsonNode {
	return JsonNode{nodeType: jsonTypeNil}
}

func (n *JsonNode) String(sorted bool) string {
	switch n.nodeType {
	case jsonTypeNil:
		return "null"
	case jsonTypeObject:
		objStr := ""
		keys := getKeys(n.objectVal, sorted)
		for _, key := range keys {
			objStr = objStr + fmt.Sprintf("%q", key) + ":"
			value := n.objectVal[key]
			objStr = objStr + value.String(sorted) + ","
		}
		if len(objStr) > 0 {
			objStr = objStr[:len(objStr)-1]
		}
		return fmt.Sprintf("{%s}", objStr)
	case jsonTypeArray:
		arrayStr := ""
		for _, val := range n.arrayVal {
			next := val.String(sorted)
			arrayStr = arrayStr + next + ","
		}
		if len(arrayStr) > 0 {
			arrayStr = arrayStr[:len(arrayStr)-1]
		}
		return fmt.Sprintf("[%s]", arrayStr)
	case jsonTypeStr:
		return fmt.Sprintf("\"%s\"", n.strVal)
	case jsonTypeNum:
		return strconv.FormatFloat(n.numVal, 'f', -1, 64)
	case jsonTypeBool:
		if n.boolVal {
			return "true"
		} else {
			return "false"
		}
	default:
		panic(fmt.Sprintf("not supported jnode type %v", n.nodeType))
	}
}

func getKeys(kvps map[string]JsonNode, sorted bool) []string {
	keys := make([]string, len(kvps))

	i := 0
	for key := range kvps {
		keys[i] = key
		i++
	}

	if sorted {
		sort.Strings(keys)
	}

	return keys
}
