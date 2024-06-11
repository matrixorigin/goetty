package buf

// Allocator memory allocation for ByteBuf
type Allocator interface {
	// Alloc allocate a []byte with len(data) >= size, and the returned []byte cannot
	// be expanded in use.
	Alloc(capacity int) ([]byte, error)
	// Free free the allocated memory
	Free([]byte)
}

type nonReusableAllocator struct {
}

func newNonReusableAllocator() Allocator {
	return &nonReusableAllocator{}
}

func (ma *nonReusableAllocator) Alloc(size int) ([]byte, error) {
	return make([]byte, size), nil
}

func (ma *nonReusableAllocator) Free([]byte) {

}
