package cache

// NewCache creates a new cache instance
func NewCache(size int) *Cache {
	return &Cache{
		size:      size,
		Responses: make(map[string][]byte),
	}
}

// AddResponse adds a response to the cache for the given request
func (c *Cache) AddResponse(request string, response []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.Responses[request]; exists {
		return
	}

	c.Requests = append(c.Requests, request)
	if len(c.Requests) > c.size {
		oldestRequest := c.Requests[0]
		delete(c.Responses, oldestRequest)
		c.Requests = c.Requests[1:]
	}

	c.Responses[request] = response
}

// GetResponse retrieves the response for the given request from the cache
func (c *Cache) GetResponse(request string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	response, exists := c.Responses[request]
	return response, exists
}

// ResetRequest removes a request from the cache
func (c *Cache) ResetRequest(request string) {
	delete(c.Responses, request)
	var requests []string
	for _, req := range c.Requests {
		if req != request {
			requests = append(requests, req)
		}
	}
	c.Requests = requests
}
