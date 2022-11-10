# Merge Sort

1. [+50%] Given the head of a singly linked list, sort it using an **in-place**
   the merge sort, and return the list after sorting it in ascending order.

   * The implementation must have time complexity $O(n \log n)$ and require
     auxiliary space up to $O(\log n)$.
   * You are not allowed to create new `node` structs in `merge`, `split` or
     `mergeSort` functions.
   * You are not allowed to use any packages apart from
     `fmt` for printing an output in `main` function.
   
   Implement the program in `merge_sort.go` and add tests in
   `merge_sort_test.go`.

2. [+50%] Write a recurrence equation for the merge sort. Prove that the merge
   sort complexity is $O(n \log n)$ by solving the equation.

3. [optional, +25%] We can express insertion sort as a recursive procedure as follows. In
   order to sort $A[1..n]$, we recursively sort $A[1..n-1]$ and then insert
   $A[n]$ into the sorted array $A[1..n-1]$.

   Write a recurrence equation for the insertion sort. Prove that the insertion
   sort complexity is $O(n^2)$ by solving the equation.
