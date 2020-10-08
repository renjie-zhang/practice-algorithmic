/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package lru_cache

type LRUCache struct {
	size int
	capacity int
	cache map[int]*DoubleLinkedNode
	head,tail *DoubleLinkedNode
}

type DoubleLinkedNode struct {
	key,value int
	pre,next *DoubleLinkedNode
}

func initDoubleLinkedList(key,value int) *DoubleLinkedNode{
	return &DoubleLinkedNode{
		key:   key,
		value: value,
		pre:   nil,
		next:  nil,
	}
}

func Constructor(capacity int) LRUCache{
	lru := LRUCache{
		capacity: capacity,
		cache: map[int]*DoubleLinkedNode{},
		head:     initDoubleLinkedList(0,0),
		tail:     initDoubleLinkedList(0,0),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (c *LRUCache) Get(key int) int{
	if _,ok := c.cache[key];!ok{
		return -1
	}
	n := c.cache[key]
	c.moveToHead(n)
	return n.value
}

func (c *LRUCache) Put(key int,value int){
	if _,ok := c.cache[key];!ok{
		n := initDoubleLinkedList(key,value)
		c.cache[key] = n
		c.addToHead(n)
		c.size++
		if c.size > c.capacity{
			re := c.removeTail()
			delete(c.cache,re.key)
			c.size--
		}
	}else {
		n := c.cache[key]
		n.value = value
		c.moveToHead(n)
	}
}


func (c *LRUCache) addToHead(node *DoubleLinkedNode){
	node.pre = c.head
	node.next = c.head.next
	c.head.next.pre = node
	c.head.next = node
}

func (c *LRUCache) removeNode(node *DoubleLinkedNode){
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (c *LRUCache) moveToHead(node *DoubleLinkedNode){
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeTail() *DoubleLinkedNode{
	n := c.tail.pre
	c.removeNode(n)
	return n
}
