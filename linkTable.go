// linktable
package main

import (
    "fmt"
    "sync"
    "unsafe"
)

const SUCCESS = 0
const FAILURE = -1

//定义节点
type LinkTableNode struct {
    pNext *LinkTableNode
}

//定义LinkTable
type LinkTable struct {
    pHead     *LinkTableNode
    pTail     *LinkTableNode
    SumOfNode int
    mutex     sync.Mutex
}

//创建LinkTable
func CreateLinkTable() *LinkTable {
    var pLinkTable *LinkTable = new(LinkTable)
    if pLinkTable == nil {
        return nil
    }
    pLinkTable.pHead = nil
    pLinkTable.pTail = nil
    pLinkTable.SumOfNode = 0
    return pLinkTable
}

//删除LinkTable
func DeleteLinkTable(pLinkTable *LinkTable) int {
    if pLinkTable == nil {
        return FAILURE
    }
    for pLinkTable.pHead != nil {
        var p *LinkTableNode = pLinkTable.pHead
        pLinkTable.mutex.Lock()
        pLinkTable.pHead = p.pNext
        pLinkTable.SumOfNode--
        pLinkTable.mutex.Unlock()

    }
    pLinkTable.pHead = nil
    pLinkTable.pTail = nil
    pLinkTable.SumOfNode = 0
    return SUCCESS
}

//添加节点
func AddLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int {
    if pLinkTable == nil || pNode == nil {
        return FAILURE
    }
    pLinkTable.mutex.Lock()
    if pLinkTable.pHead == nil && pLinkTable.pTail == nil {
        pLinkTable.pHead = pNode
        pLinkTable.pTail = pNode
        pLinkTable.pTail.pNext = nil
        pLinkTable.SumOfNode = 1
    } else {
        pLinkTable.pTail.pNext = pNode
        pLinkTable.pTail = pNode
        pLinkTable.pTail.pNext = nil
        pLinkTable.SumOfNode++
    }
    pLinkTable.mutex.Unlock()
    return SUCCESS
}

//删除节点
func DelLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int {
    if pLinkTable == nil || pNode == nil {
        return FAILURE
    }
    pLinkTable.mutex.Lock()
    var pWork *LinkTableNode = pLinkTable.pHead
    var pre *LinkTableNode = pWork
    if pLinkTable.pHead == pNode {
        pLinkTable.pHead = pWork.pNext
        pLinkTable.SumOfNode--
        return SUCCESS
    }
    for pWork != nil {
        if pWork == pNode {
            pre.pNext = pWork.pNext
            pLinkTable.SumOfNode--
            return SUCCESS
        }
        pre = pWork
        pWork = pWork.pNext
    }
    return FAILURE
}

//查询节点
func SearchLinkTableNode(pLinkTable *LinkTable, Condition func(pNode *LinkTableNode, args unsafe.Pointer) int, args unsafe.Pointer) *LinkTableNode {
    if pLinkTable == nil || Condition == nil {
        return nil
    }
    var pNode *LinkTableNode = pLinkTable.pHead
    for pNode != nil {
        if Condition(pNode, args) == SUCCESS {
            return pNode
        }
        pNode = pNode.pNext
    }
    return nil
}

//获取LinkTable头节点
func getLinkTableHead(pLinkTable *LinkTable) *LinkTableNode {
    if pLinkTable == nil {
        fmt.Println("LinkTable is empty")
        return nil
    }
    return pLinkTable.pHead
}

//获取下一个节点
func getNextLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) *LinkTableNode {
    if pLinkTable == nil || pNode == nil {
        fmt.Println("Linktable is empty")
        return nil
    }
    var pWork *LinkTableNode = pLinkTable.pHead
    for pWork != nil {
        if pWork == pNode {
            return pWork.pNext
        }
        pWork = pWork.pNext
    }
    return nil
}
