// Copyright 2017 XUEQIU.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dump

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xueqiu/rdr/decoder"
)

func TestGetLargestKeyPrefixes(t *testing.T) {
	e := &decoder.Entry{
		Key:                "RELATIONSFOLLOWERIDS6420000664",
		Bytes:              1,
		Type:               "sortedset",
		NumOfElem:          1,
		LenOfLargestElem:   1,
		FieldOfLargestElem: "test",
	}
	c := NewCounter()
	c.countByKeyPrefix(e)
	c.calcuLargestKeyPrefix(1)
	for _, p := range c.GetLargestKeyPrefixes() {
		assert.Equal(t, "RELATIONSFOLLOWERIDS0000000000", p.Key)
	}
}
