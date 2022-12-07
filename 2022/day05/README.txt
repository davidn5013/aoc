I never solves day 5, trapped in parsing hell.

But I learn it better to use bufio.Scanner och fmt.Sscanf for string comprehension

I only thought of scanf as  taking input and storing in variable using the buffer string 
fmt.Scanf("%s %s %d %s", &name, &temp, &amount, &unit) 

But using to read a bufio.Scanner like this:

fmt.Sscanf(sc.Text(), "move %d from %d to %d", &toMove, &from, &to)

to interpret file line of "move 1 from 2 to 3" or strings with file 
contemns It's just really easier then doing strings.Split
