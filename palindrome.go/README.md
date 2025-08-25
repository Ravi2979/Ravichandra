 # Code explanation of each

 <pre> 1. Palindrome with Fibonacci length check</pre>

 **palindrome - It is nothing but a word that reads the same forward and backward.**

 Firstly, I declare the var name string 
 I have used keyword called scanln (it is used to read the input from name)

Secondly, if find the distance of the word called name

Third, I have started with a = 1, b = 2 and generate numbers like 1, 2, 3….

If the name’s length matches one of these numbers, I mark has isFibonacci = true.
I reverse the word using a for loop (from last character to first). Then I compare the reversed word with the original.

If the length is not Fibonacci 
I simply print that the name is not in Fibonacci sequence and don’t reverse it.

<pre> 2. By Go Routine I'm trying to print differnt lang </pre>

Go Routines - It allow multiple tasks to be executed simultaneously to improve performance and speed of program.

First, I have taken 5 input String values in the list
<p>I have used **len(languages)** tells how many items are in the list.</p>
<p>i starts from 0 and increases by 1 each time.</p>
<p>The loop runs while i < 5, so i = 0, 1, 2, 3, 4.</p>
<p>If main program ends, the whole program ends and sometimes goroutines don’t get a chance to print.
That’s why I add **time.Sleep**, it pauses the main program for a little while. so the goroutines can finish.</p>

<pre> 3. Factorial program. </pre>
<p>Initially I have declare the var f to hold an integer, user had a chance to type a number and press enter<p>
<p>Fact := 1<p>
<p>fact will hold the result of the factorial. I can start with 1 because factorial is multiplication and multiplying by 0 would always give 0. i will start for loop from i = 1 until i<=f, each time it will multiply by i </p>

<pre> 4. Count the character's </pre>

<p>First, I have 4 string name letters "go run fast go"<p>
<p>second, when we start the program the count will start from 0. I'm trying to print how many times the words "go", "run", and "fast" appear in the list.<p>
<p>Third, I used **len(letters)** gives the number of items and loop start from i=0,1,2,3.<p>
For each index i, letters[i] is the current word. I compare that each word to the exact strings "go", "run", and "fast".If it matches "go", I will do goCount++ (add 1). otherwise, Else if it matches "run", add 1 to runCount. Else if it matches "fast", add 1 to fastCount. If the word is something else (example - "ravi"), nothing happens there is no final else, so it’s ignored.