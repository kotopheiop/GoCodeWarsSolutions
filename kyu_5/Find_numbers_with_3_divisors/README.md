# Find numbers with 3 divisors

Your task is to write a function that takes two integers `n` and `m`, and returns a sorted array of all integers from `n` to `m` inclusive, which have exactly 3 divisors (excluding 1 and the number itself).

Example:

```solution(2, 20) -> [16]```

`16` has 3 divisors: `2, 4, 8` (`1` and `16` aren't included)

Input:
* `n` - integer `(2 ≤ n ≤ 10^14)`
* `m` - integer `(2 ≤ m ≤ 10^18)`

Output:
* `result` - array of integers