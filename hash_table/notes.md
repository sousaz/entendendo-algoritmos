## When to Use:

When you need fast lookups, insertions, and deletions (on average O(1)).

When you have a large dataset and need a data structure optimized for quick access.

When data can be mapped to unique keys and hash functions can be efficiently designed.

Ideal for applications like caches, symbol tables, or sets.

## Properties:

Fast Access: Provides constant average-time complexity for search, insert, and delete operations.

Uses Hashing: Data is stored in an array format, and its index is determined by a hash function.

Flexible Key Types: Keys can be of various types as long as they are hashable.

Collision Handling: Uses strategies like separate chaining or open addressing to manage hash collisions.

Load Factor: Performance is influenced by the load factor (number of elements / size of the hash table), which is typically kept below a threshold to minimize collisions.

## Big O:
Best Case: O(1) (when there are no collisions, and the hash function distributes keys evenly).

Average Case: O(1) (with a good hash function and low load factor).

Worst Case: O(n) (occurs when all elements hash to the same index, degrading to a linear search).