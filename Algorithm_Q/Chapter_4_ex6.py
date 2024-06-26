"""
Generated by Google bard (https://bard.google.com/)
"""

def divide_and_conquer_count_positives(array):

  if len(array) == 1:
    if array[0] > 0:
      return 1
    else:
      return 0
  else:
    mid = len(array) // 2
    left = divide_and_conquer_count_positives(array[:mid])
    right = divide_and_conquer_count_positives(array[mid:])
    
    return left + right 


def main():
    array = [4,2,3,2,1,-1,0]
    result = divide_and_conquer_count_positives(array)
    print(result)

if __name__ == "__main__":
    main()
