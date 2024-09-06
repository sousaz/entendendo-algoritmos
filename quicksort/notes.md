## When to use:

When you need an efficient sorting algorithm, especially for large arrays

When speed is a priority, and you can handle occasional worst-case performance

When minimizing additional memory is important, as Quicksort is an in-place algorithm

## Property

It is one of the fastest sorting algorithms for most cases, with an average time complexity of O(n log n)

Works well on large arrays and is optimized for fast data access (e.g., in memory)

Can be improved with techniques like optimized pivot selection to reduce the risk of the worst case (O(n²))

## Big O:  

Worst case: O(n²) (occurs when the pivot consistently selects the largest or smallest element)

Average case: O(n log n)