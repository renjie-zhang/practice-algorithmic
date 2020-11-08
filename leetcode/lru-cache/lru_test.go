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
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)

	assert.Equal(t, true, reflect.DeepEqual(1, obj.Get(1)))

	obj.Put(3, 3)
	assert.Equal(t, true, reflect.DeepEqual(-1, obj.Get(2)))
	obj.Put(4, 4)
	assert.Equal(t, true, reflect.DeepEqual(-1, obj.Get(1)))
	assert.Equal(t, true, reflect.DeepEqual(3, obj.Get(3)))
	assert.Equal(t, true, reflect.DeepEqual(4, obj.Get(4)))
}
