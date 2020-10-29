import time
import multiprocessing


def basic_fun(x):
    if x == 0:
        return 'zero'
    elif x%2 == 0:
        return 'even'
    esle:
        return 'odd'


def multiprocessing_fun(x):
    y = x*x
    time.sleep(2)
    print