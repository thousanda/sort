#!/bin/bash

# $ ./create_input.sh 5 > input10.txt
#
# こんな感じのデータができる
# ```
# 5
# 5 1 3 4 2
# ```
# 

n=$1
echo $n
seq $n | sort -R | tr '\n' ' '
echo
