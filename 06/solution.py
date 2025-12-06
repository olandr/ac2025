## Day 06

sample_data = [
"123 328  51 64 ",
" 45 64  387 23 ",
"  6 98  215 314",
"*   +   *   +  " 
]

data = []
with open("./input.txt", "r+") as f:
  #data = [list(map(int, line.rstrip('\n'))) for line in f]
  data = [line.rstrip('\n') for line in f]


def solution_01(data):
  count=0
  operations=data[len(data)-1]
  
  idx_old=len(operations)
  for idx,value in reversed(list(enumerate(operations))):
    if value != " ":
      c = 0 if value == "+" else 1
      for x in data[:len(data)-1]:
        if value == "*":
          c *= int(x[idx:idx_old])
        else:
          c += int(x[idx:idx_old])
      idx_old=idx
      count += c
  return(count)

def solution_02(data):
  return(0)


print("Solution 1 test: ",solution_01(sample_data))
print("Solution 1: ",solution_01(data))  ## Solution 1:  4805473544166

#print("Solution 2 test: ",solution_02(sample_data))
#print("Solution 2: ",solution_02(data)) ## Solution 2:  xxx
