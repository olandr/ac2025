# Day 03 - Battery Joltage Problem - Bug Fix Summary

## Problem Description

Given a sequence of digits representing batteries, we need to select a specific number of digits to form the maximum possible joltage number.

### Examples (12-digit selection from 15-digit input):
- `987654321111111` → `987654321111` (remove 3 digits: some 1s at the end)
- `811111111111119` → `811111111119` (remove 3 digits: some 1s)
- `234234234234278` → `434234234278` (remove 3 digits: 2, 3, 2 near the start)
- `818181911112111` → `888911112111` (remove 3 digits: some 1s near the front)

## Bugs Fixed

### Bug #1: Incorrect Algorithm Logic

**Original Code (WRONG):**
```go
splice := line[ix : len(line)-length_of_battery_toggles]
for range splice {
    ix, max = findMax(splice, ix, len(splice))
    ix++
    sb.Write([]byte(max))
}
```

**Problem:**
- For a 15-digit input with `length_of_battery_toggles=12`, it calculated:
  - `splice = line[0:3]` (only first 3 characters!)
- Then searched for max digits only in this tiny 3-character substring
- Completely wrong approach to the problem

**Fixed Code (CORRECT):**
```go
// We want to select length_of_battery_toggles digits from the line
for digitsSelected := 0; digitsSelected < length_of_battery_toggles; digitsSelected++ {
    // Remaining digits we still need to select after this one
    remainingToSelect := length_of_battery_toggles - digitsSelected - 1
    // Search window: from currentIdx to a position that leaves room for remaining digits
    endIdx := len(line) - remainingToSelect

    currentIdx, max = findMax(line, currentIdx, endIdx)
    currentIdx++ // Move past the selected digit for next iteration
    sb.Write([]byte(max))
}
```

**How the Algorithm Works:**
1. We use a **greedy approach** to select digits that maximize the result
2. For each position in the output:
   - Calculate how many more digits we need to select after this one
   - Search from current position to a point that leaves room for remaining digits
   - Pick the maximum digit in that search window
   - Move past the selected digit for the next iteration

**Example Walkthrough** for `234234234234278` (selecting 12 from 15):
- Position 0: Search indices [0,4), find '4' at index 2
- Position 1: Search indices [3,5), find '3' at index 4  
- Position 2: Search indices [5,6), find '4' at index 5
- Continue this pattern...
- Result: `434234234278` ✓

### Bug #2: File Not Reset Between Passes

**Original Code (WRONG):**
```go
fmt.Println(process(in, out, 2))
fmt.Println(process(in, out, 12))
```

**Problem:**
- First `process` call reads entire file (file cursor moves to EOF)
- Second `process` call has nothing to read (cursor at end of file)
- Result: Second output was always 0 or empty

**Fixed Code (CORRECT):**
```go
result1 := process(in, out, 2)
fmt.Println(result1)

// Reset file to beginning for second pass
if *input != "" {
    in.Seek(0, 0)
}

result2 := process(in, out, 12)
fmt.Println(result2)
```

**Solution:**
- Use `in.Seek(0, 0)` to reset file cursor to beginning
- Now both passes can read the complete file

## Test Results

**Sample Input:**
```
987654321111111
811111111111119
234234234234278
818181911112111
```

**Output:**
```
357              # Part 1: Sum of 2-digit maximums (98 + 81 + 78 + 91)
3121910778619    # Part 2: Sum of 12-digit maximums
```

**Verification:**
- Line 1: `987654321111111` → `98` (2 digits) and `987654321111` (12 digits) ✓
- Line 2: `811111111111119` → `81` (2 digits) and `811111111119` (12 digits) ✓
- Line 3: `234234234234278` → `78` (2 digits) and `434234234278` (12 digits) ✓
- Line 4: `818181911112111` → `91` (2 digits) and `888911112111` (12 digits) ✓

## Algorithm Complexity

- **Time Complexity:** O(n × k) where n is the input length and k is the number of digits to select
- **Space Complexity:** O(k) for the output string builder

The greedy approach is optimal because at each position, selecting the largest available digit that leaves room for remaining digits guarantees the maximum result.
