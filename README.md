# HarryPotterWordlist
A word list themed Harry Potter for dictionary attacks

- ```harry1.txt``` contains Harry Potter related words.
- ```harry1_nums.txt``` contains Harry Potter related words and all numbers from 1 to 4 digits
- ```harry2.txt``` contains all the permutations from ```harry1.txt``` in format ```[word]<space>[word]``` i.e. two words separated by a single space. <br>
Made using: ```awk 'NR==FNR{a[$0];next}{for(i in a)print i" "$0}' harry1.txt harry1.txt > harry2.txt```
- ```harry2_nums.txt``` is same as above but is made using ```harry1_nums.txt``` instead of ```harry1.txt``` i.e. contains all numbers of 1 to 4 digits.
- ```harry3.go``` is a go program to generate ```[word]<space>[word]<space>[word]```. It is made in go language because the performance was so low in bash and awk when creating such a large wordlist.

Note: The dictionary which contains two words does not contain one words, only: ```[word]<space>[word]``` no ```[word]```
