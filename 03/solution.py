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


def solution(data,len_batt=2):
  count = 0
  for bank in data:
    out = ""
    bank_new=bank[:len(bank)-len_batt+1]
    bank_end=bank[len(bank)-len_batt+1:]

    for x in range(len_batt):
      val = max(bank_new)
      out+=val
      idx = bank_new.find(val)
      if x < len_batt-1:
        bank_new = bank_new[idx+1:]
        bank_new += bank_end[x]

    count += int(out)
  return(count)

#print("Solution 1 test: ",solution(sample_data,2))
print("Solution 1: ",solution(data,2))  ## Solution 1:  16927

#print("Solution 2 test: ",solution(sample_data,12))
print("Solution 2: ",solution(data,12)) ## Solution 2:  167384358365132
