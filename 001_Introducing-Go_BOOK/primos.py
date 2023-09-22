"""
Created on Sun Mar  7 17:11:12 2021

@author: User

SUMMATION OF PRIMES

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

21min19s to find. 

"""
from datetime import datetime as date
import os

def nextPrime(n, primes):
    isPrime = False
    
    if n % 2 == 0: 
        n +=1
    else: 
        if n % primes[-1] == 0: 
            n += 2
    
    while not isPrime:
        for prime in primes:
            # FIRST: Cheking if if is prime or should go next. SECOND: it is
            # a waste to try division by a prime which returns less than 3
            if n % prime == 0:
                n +=2
                # skipping 5
                if str(n)[-1] == "5":
                    n += 2
                break
            if prime == primes[-1]:
                isPrime = not isPrime
            if n / prime < 3 : 
                isPrime = not isPrime
                break
    return n

# INPUTS
target = 2000000
primes = [2, 3, 5, 7, 11, 13, 17, 19]
control = target / 10

#INITIALIZING
print("INICIANDO BUSCA DOS NÚMEROS PRIMOS MENORES QUE", target)
start = date.now()

# PROCESSING
while primes[-1] < target :
    candidate = nextPrime(primes[-1], primes)
    if candidate > target : 
        break
    primes.append(candidate)
# CONTROLLING
    if candidate >= control:
        print("O", len(primes), "º primo é", candidate, "em", date.now() - start)
        control += target / 10

# OUTPUT
print("\n")
print("RESULTADOS:")
print("ENCONTRAR OS NÚMEROS PRIMOS MENORES QUE", target)
print("FORAM ENCONTRADOS", len(primes), "NÚMEROS PRIMOS")
print("ÚLTIMO PRIMO DA LISTA:", primes[-1])
print("SOMA DOS PRIMOS ENCONTRADOS:", sum(primes))
print("TEMPO TOTAL DA BUSCA:", date.now() - start)

# TO FILE
try:
    path = os.getcwd()  
    print(path)
    file = "primos_ate_" + str(target) + ".csv"  
    with open(path + "/" + file, "w+") as f:
        for i in range(len(primes)):
            f.write(str(i+1))
            f.write("\t")               # tab
            f.write(str(primes[i]))
            f.write("\r")               # carriage return
        f.close()
except IOError as e:
    print("Erro ao abrir ou manipular o arquivo:", str(e))