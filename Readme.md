vmMath is an attempt at creating a cross-platform (all the ones Go supports) high-performance vector and matrix math library in *pure* Go (see below).

'high-performance' having been said, there will be a fairly specific progression of development as far as optimizations are concerned:

1. A pure Go implimentation; in this context 'pure Go' means not using the Go assembly language. v1.0 will represent the ending of this phase. In this first iteration the types and library interface will be laid out, experimented with, and finalized. Yes, I said *finalized*; mplimentation will be laid out with knowledge of how some processor optimizations work in order to prevent hackery in the future, and other vector and matrix math libaries will be referenced when deciding on the interface, but once we reach version 1.0 no breaking changes will be considred until all three phases have completed a full round. This is to help encourage people to start using this library to experiment, while it'll be made performant later.

2. Convert the lower level operations to use processor specific optimisations. eg: SIMD instructions, and/or prefetch commands.

3. Begin considering optimizations that sacrifice code clarity. ie: iterative processes instead of recursive ones to avoid function overhead, or other things I can't think of to remove runtime overhead.
