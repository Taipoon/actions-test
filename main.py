import sys


def main():
  print("Hello, world!")

def average(numbers):
  sum = 0
  for number in numbers:
    sum += number
  return sum / len(numbers)

def maximum(numbers):
  tmp = (-1) * sys.maxsize
  for number in numbers:
    if tmp < number:
      tmp = number
  return tmp

def minimum(numbers):
  tmp = sys.maxsize
  for number in numbers:
    if number < tmp:
      tmp = number
  return tmp

if __name__ == "__main__":
  nums = [10, 54, 22, 43, 86, 45, 12]
  main()
  print("Numbers:", nums)
  print("Max:", maximum(nums))
  print("Avg:", average(nums))
  print("Min:", minimum(nums))
