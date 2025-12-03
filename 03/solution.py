## Day 03

sample_data = [
"987654321111111",
"811111111111119",
"234234234234278",
"818181911112111"
]

data = []
with open("./input.txt", "r+") as f:
  #data = [list(map(int, line.rstrip('\n'))) for line in f]
  data = [line.rstrip('\n') for line in f]


def solution_01(data):
  count = 0
  for bank in data:
    val = max(bank)
    idx = bank.find(val)
    if idx == len(bank)-1:
      bank_new = bank[:idx] + bank[idx+1:]
      val_2 = max(bank_new)
      count += int(val_2+val)
      #print(bank,bank_new,val,idx,val_2,int(val_2+val))
    else:
      bank_new = bank[idx+1:]
      val_2 = max(bank_new)
      count += int(val+val_2)
      #print(bank,bank_new,val,idx,val_2,int(val+val_2))
  return(count)


def solution_02(data):
  return(0)


print("Solution 1 test: ",solution_01(sample_data))
print("Solution 1: ",solution_01(data))  ## Solution 1:  xxx

#print("Solution 2 test: ",solution_02(sample_data))
#print("Solution 2: ",solution_02(data)) ## Solution 2:  xxx
