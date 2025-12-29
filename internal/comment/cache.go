package comment

import (
	"go/ast"
	"go/token"
	"sync"
	"sync/atomic"
)

// Cache provides thread-safe caching of comment maps.
type Cache struct {
	comments map[*ast.File]ast.CommentMap
	mu       sync.RWMutex

	hits   atomic.Uint64
	misses atomic.Uint64
}

const cachePreallocSize = 64

// Get returns a comment map for a given file, creating and caching it if needed.
func (c *Cache) Get(fset *token.FileSet, f *ast.File) ast.CommentMap {
	c.mu.RLock()

	cm, ok := c.comments[f]

	c.mu.RUnlock()

	if ok {
		c.hits.Add(1)

		return cm
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// Double-check after acquiring write lock
	if cm, ok = c.comments[f]; ok {
		c.hits.Add(1)

		return cm
	}

	if c.comments == nil {
		c.comments = make(map[*ast.File]ast.CommentMap, cachePreallocSize)
	}

	c.misses.Add(1)

	cm = ast.NewCommentMap(fset, f, f.Comments)
	c.comments[f] = cm

	return cm
}

// Stats returns cache hit and miss counts.
func (c *Cache) Stats() (hits, misses uint64) {
	return c.hits.Load(), c.misses.Load()
}
