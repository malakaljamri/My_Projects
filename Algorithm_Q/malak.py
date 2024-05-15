
def malak(array):

  if len(array) == 1:
    return True
  else:
    mid = len(array) // 2
    left = malak(array[:mid])
    right = malak(array[mid:])
    
    return left and right and array[mid - 1] % 2 != array[mid] % 2    



# def main():
#     array = [1,1,4]
#     sum = malak(array)
#     print(sum)

def main():
    array = [1,2,3,4,5]
    result = malak(array)
    print(result)

if __name__ == "__main__":
    main()
