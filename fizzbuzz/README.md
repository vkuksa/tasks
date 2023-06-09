# Description:
Given an integer n, return a string array answer (1-indexed) where:

    answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
    answer[i] == "Fizz" if i is divisible by 3.
    answer[i] == "Buzz" if i is divisible by 5.
    answer[i] == i (as a string) if none of the above conditions are true.

Example 1:

Input: n = 3
Output: ["1","2","Fizz"]

Example 2:

Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]

Example 3:

Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]

Constraints:
    1 <= n <= 104

# Overview
In brutforce version approach achieves decent execution speed and simplicity

In parametrised version as golang for now does not have any constraints over values that can have operations of substraction
Only constraints.Integer supported 

Parallel version execution speed is very low, and it's in inverse dependency with the number of workers.
The approach uses worker pool approach

# Important! 
In brutforce and parametrised version the sequence of result is linear and expected.
In parallel, as a result of concurrency, the sequence is not defined and presumed. It satisfies only the presence of element as constraint