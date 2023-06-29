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
for ((i=1; i<=$n; i++)); do
  echo $i
done | sort -R | tr '\n' ' '
echo
