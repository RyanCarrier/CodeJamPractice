#D
##Option 1
 - Take all combinations of K down to C complexity, excluding K L's.
 - Find the lowest amount of columns (indexes of the final complexity), which can have a G in every column.
 - Return those columns, or INSOMNIA if len(columns) > S


##Option 2
Find the pattern.


A reference to each initial item is needed;

As any initial G will leave a G in the end (Even on an L branch).

So from the root of the tree, down to level C;

  - Take branch 1, then 2, 3, 4.... min(K,C)

Then we need to keep going for any remaining 'untouched' initial items;
  - Take branch C, C+1, C+2...., 2C
Then;
  - Take branch 2C, 2C+1, 2C+2...., 3C
  - ... BC, BC+1...., (B+1)C; Where B is branch

For first index to check;

C\\K|1|2|3|4|5|6
   1|1|1|1|1|1|1
   2|1|2|2|2|2|2
   3|1|2|6|7|8|9
   4|1|2|6|20|39
   5|1|2|6|20|180
   6|1|2|6|20|
