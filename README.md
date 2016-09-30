# Data Center Recovery

Simulates a distributed data center recovery procedure. 

Input:
------

Reads from standard input.

The first line of input will contain an integer between 0 and 999999
inclusive, representing the number of data centers.

Following that will be one line of input for each data center with a
space-separated list of data set ids currently present at that data
center. Data set ids are each an integer between 1 and 999999, inclusive.
Each line will be at most 1000 characters long.

A data center with no data sets is represented with a blank line. Data
set ids are not necessarily consecutive. The list of data sets will not
be in any particular order.

There will be exactly one line for each data center, meaning the total
number of lines in the file will be `N+1`, where `N` is the number of data
centers.


Output:
-------

The program outputs an optimal data set copy strategy to ensure that
every data center has a copy of every data set. Output one line for every
copy instruction.

No data set is copied to a data center that already contains
said data set, and at the end of operation every data center has a single
copy of every data set.

A copy instruction is of the form

`ds_id` `from_dc` `to_dc`

where `ds_id` is the data set id, `from_dc` is the index of the data center
the data set will be copied to (1 = the first data center) and `to_dc` is the
index of the data center to copy the data set to.

The last line of your ouptut will be `done`.


Reads from standard input:
------------

```
$ ./datasync < input.txt
```

Example:
----------
Input:

```
3
1 3 4 5 7
1 3
2
```

Possible Correct Output:

```
2 3 2
2 3 1
1 1 3
4 1 2
5 1 3
5 3 2
4 2 3
3 1 3
7 1 2
7 1 3
done
```
