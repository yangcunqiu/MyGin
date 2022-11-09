package mg

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		log.Printf("%v-%v start", c.Method, c.Path)
		start := time.Now()
		c.Next()
		log.Printf("%v-%v end, cost %v ns", c.Method, c.Path, time.Since(start).Nanoseconds())
	}
}
