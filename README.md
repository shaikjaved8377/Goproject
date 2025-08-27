
Goroutines: go routines are lightweight threads and its started with go keyword. we can run millions of 
            go routines. The main function will end the program when it completes even the other go routines 
            still running, So for that we have use time.sleep() methods to wait the main program until the go routines completes and the exit the main program.

syn.wait groups in go: its a same as time.sleep but in real word we can decide the program execution
                       time right so in that case we have use wait groups for example if the program execution is completes in 2 seconds and we are giving time.sleep for 5 seconds in that case we are wasting 3 seconds right so there the wait groups comes into picture, So here we have 3 methods like 
                       var wg sync.WaitGroup we have declare the variable wg for waitgroup and then we have different methods, wg.Add(n) , wg.Done(), wg.Wait().

                       so wg.Add(n) means we have to add some value for example 1 atleast if you give 0 there is no 
                       use of adding the value because we are running go routines here so we have tell like these many go routines that we are running and we have to wait.

                       wg.Done(): which means onces the go routines finishes or executed the go will call this function to tell that im done.
                        
                       wg.Wait(): which means it will wait like the Add value to become zero then the program exits 
                                  it will wait untill the all go routines are done.



channels: Channels in go are nothing but the communication between two go routines, and we will use variable chan 
          to declear the channel alago with data type and will pass the value into the channel from one go to another go routines like ch <- i. there are two types of channles in go 

          1.Buffered channel: Buffered channel is like sender will send the data but reciver will pick the 
                              data later. and also buffered channel is like has a fixed capacity. based on 
                              the capacity its completed the program. example: ch:= make(chan int, 5)


          2. Unbuffered channel: unbuffered channel is like sending the data and reciver will recive the data from 
                                 channel directly. for unbuffered is no capacity. ex: ch:= make(chan int)
       

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