# Program to check if a number is prime or not

import time
import math 
num = 27064032706411   #9999999992999999999
start_time = time.time()
# To take input from the user
#num = int(input("Enter a number: "))

# prime numbers are greater than 1
if num > 1:
   # check for factors
   for i in range(2,int(math.sqrt(num))):
       if (num % i) == 0:
           print(num,"is not a prime number")
           print(i,"times",num//i,"is",num)
           break
   else:
       print(num,"is a prime number")
       
# if input number is less than
# or equal to 1, it is not prime
else:
   print(num,"is not a prime number")

print("---Time taken by python is  %s seconds ---" % (time.time() - start_time))