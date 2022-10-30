package lib

import (
	"github.com/dghubble/trie"
	"github.com/gogf/gf/util/gconv"
	"sort"
	"strings"
	"sync"
)

type Trie struct {
	tree *trie.PathTrie
}

type Item struct {
	Path      string
	Timestamp int
}

var instance *Trie

var once sync.Once

func GetTrie() *Trie {
	once.Do(func() {
		instance = &Trie{tree: trie.NewPathTrie()}
	})
	return instance
}

func Add(path string, timestamp int) {
	path = strings.Trim(path, "/")
	GetTrie().tree.Put(path, timestamp)
}

func Walk(path string) int {
	path = strings.Trim(path, "/")
	tmp := make([]*Item, 0)
	GetTrie().tree.WalkPath(path, func(key string, value interface{}) error {
		tmp = append(tmp, &Item{
			Path:      path,
			Timestamp: gconv.Int(value),
		})
		return nil
	})
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].Timestamp > tmp[j].Timestamp
	})
	if len(tmp) == 0 {
		return 0
	}
	return tmp[0].Timestamp
}
