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

import (
	"container/list"
)

type LRUCache struct {
	v        map[int]int
	capacity int
	list *list.List
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		v:        make(map[int]int, capacity),
		capacity: capacity,
		list: list.New(),
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.v[key]
	var v *list.Element
	if ok {
		for e := this.list.Front();e != nil;e = e.Next(){
			if e.Value.(int) == key{
				v = e
				break
			}
		}
		this.list.MoveToFront(v)
		return value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if len(this.v) == this.capacity{
		_, ok := this.v[key]
		var v *list.Element
		if ok {
			for e := this.list.Front();e != nil;e = e.Next(){
				if e.Value == key{
					v = e
					break
				}

			}
			this.list.MoveToFront(v)
			this.v[key] = value
		}else {
			//delete last use element
			temp := this.list.Back()
			if temp != nil{
				this.list.Remove(temp)
				delete(this.v,temp.Value.(int))
			}
			this.list.PushFront(key)
			this.v[key] = value
		}
	}else {
		this.list.PushFront(key)
		this.v[key] = value
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
