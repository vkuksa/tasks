Task: Solve the "FizzBuzz" problem in Golang or Node.js. Make your code as extensible and flexible as possible.

Description:
Given an integer n, return a string array answer (1-indexed) where:
    answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
    answer[i] == "Fizz" if i is divisible by 3.
    answer[i] == "Buzz" if i is divisible by 5.
    answer[i] == i (as a string) if none of the above conditions are true.

In this implementation i assumed we do not need ordered results rather we expect set of results.
If we do need ordered results, concurrent version is not a case.

In parametrised version as golang for now does not have any constraints over values that can have operations of substraction
Only constraints.Integer supported 
