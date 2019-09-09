#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the reverseArray function below.
def reverseArray(a):
    print(a.reverse())
    return a

if __name__ == '__main__':

    a = input("please enter the number")
    arr = list(map(int,a.strip().split(' ')))


    res = reverseArray(arr)
    print(res)
