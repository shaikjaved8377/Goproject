1.Fibonacci series

Here im checking fibonacci sereies for 10, so i took number as 10 and fibonnaci starts with 0,1 so 
i have took e,d variable to initial, used for loop to print all the given 10 numer series
fmt.Printf("%d ", e)-> which neans every time its prints new value e, d = d, e+d -> which means now new e value is old d and d value is e+d (adding perviois two numbers)

2.Factorial program

Here im doing factorial of 10, so the logic here im using like factorial starts with 1, using for loop to run the till 10 and Factorial *= i -> which means each time when it runs it multiplies the current Factorial value by i after that it prints both the number and its factorial value.

3. Plaindrome program

Here im doing like created a function called palindrome and string s as input and i have used mention called
s = strings.ToLower(s) -> which means it converts the string to lowercase and decleared reversed as empyty string after that i used foor loop which it checks the lengh of string and its starting from last charcters len(s)-1, so it
adds each character to reversed and its building the reversed string, so then if given string matches the reversed string then its a palindrome else its not. now i declared word variable with value DAD and i called the function palindrome in main function to print.

4. Printing count of the characters:

Here as per i learned im printing the count of the characters which we gave as input so here what i have done is 
created a string variable str and gave as (go run run yes go) and creates an empty map like key is a string (string) which means each unique word and the value is an integer (int) → how many times that word appears and count is empty 
used for loop and used mentioned called strings.Fields(str) -> it will splits the string by spaces and takes as single word, after loop increments the count for the word in the map if the word appear it will increments and prints the word and its count.

5. using goroutines print diff languages:

Here as per instructor i have created 3 different functions and those functions i have called in main function using go key word and after that in main i have also used method call time.sleep() which means its goroutines have enough time to finish before the program exits without this menthod the program will end before printing gorounties and beacuse go routines will stop when main exits.