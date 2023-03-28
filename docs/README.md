###Bugs

##Bug01
1. There is a deadlock at the send op
2. Added a go routine for sending
3. Because they are in different go routines the main go routine can recieve

##Bug02
1. main ends before print prints 11
2. Added a waitgroup
3. Because now the main waits for all goroutines to be done

#many senders and many recievers
1. What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?

Ans: Some producers try to send to the already closed channel.

2. What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?

Ans: The channel gets closed when the first producer is done which makes the other producers not able to send anymore values.

3. What happens if you remove the statement close(ch) completely?

Ans: Nothing. The program will die after it anyways. But better todo it for the consumers

4. What happens if you increase the number of consumers from 2 to 4?

Ans: Will be faster because more go routines.

5. Can you be sure that all strings are printed before the program stops?

Ans: No. There is nothing that prevents the program to stop when all consumers all finished.


