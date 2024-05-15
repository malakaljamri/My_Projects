"""
Generated by Google bard (https://bard.google.com/)
"""

def divide_and_conquer_search_for_key(array,key):

  if len(array) == 1:
    # if array[0] == key:
    #   return True
    # else:
    #   return False

    # or in one line
    return array[0] == key

  else:
    mid = len(array) // 2
    left = divide_and_conquer_search_for_key(array[:mid],key)
    right = divide_and_conquer_search_for_key(array[mid:],key)
    
    return left or right 


def main():
    array = [4,2,3,2,1]
    key = 1
    result = divide_and_conquer_search_for_key(array,key)
    print(result)

if __name__ == "__main__":
    main()
