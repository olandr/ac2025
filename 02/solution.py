## Day xx

sample_data = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

data = "9595822750-9596086139,1957-2424,88663-137581,48152-65638,12354817-12385558,435647-489419,518494-609540,2459-3699,646671-688518,195-245,295420-352048,346-514,8686839668-8686892985,51798991-51835611,8766267-8977105,2-17,967351-995831,6184891-6331321,6161577722-6161678622,912862710-913019953,6550936-6625232,4767634976-4767662856,2122995-2257010,1194-1754,779-1160,22-38,4961-6948,39-53,102-120,169741-245433,92902394-92956787,531-721,64-101,15596-20965,774184-943987,8395-11781,30178-47948,94338815-94398813"

def solution_01(data):
  ranges = data.split(",")
  count = 0
  for rang in ranges:
    ids = rang.split("-")
    le = len(ids[0])
    num=0
    if le % 2 == 0:
      if int(ids[0][:int(le/2)]) >= int(ids[0][int(le/2):]):
        num = int(ids[0][:int(le/2)])
      else:
        num = int(ids[0][:int(le/2)])+1
    else:
      num = 10**int((le-1)/2)
    val = int(str(num)*2)
    while val <= int(ids[1]):
      count += val
      num += 1
      val = int(str(num)*2)
  return(count)

def solution_02(data):
  ranges = data.split(",")
  invalid_ids = set()
  for rang in ranges:
    ids = rang.split("-")
    le_min = len(ids[0])
    le_max = len(ids[1])
    le_cur = le_min
    num=0
    for le_piece in range(1,int(le_max/2)+1):
      le_cur = le_min
      if le_cur % le_piece == 0:
        num = int(ids[0][:int(le_piece)])
      else:
        num = 10**(le_piece-1)
        le_cur += 1
      val = int(str(num)*int(round(le_cur/le_piece)))
      while val <= int(ids[1]):
        if val >= int(ids[0]) and val > 10:
          invalid_ids.add(val)
        if num == (10**le_piece)-1:
          le_cur += le_piece
          num = 10**(le_piece-1)
        else:
          num += 1
        val = int(str(num)*int(le_cur/le_piece))
  return(sum(invalid_ids))


#print("Solution 1 test: ",solution_01(sample_data))
print("Solution 1: ",solution_01(data))  ## Solution 1:  40398804950

#print("Solution 2 test: ",solution_02(sample_data))
print("Solution 2: ",solution_02(data)) ## Solution 2: 65794984339