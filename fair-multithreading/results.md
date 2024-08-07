# Task
- To find the count of prime numbers till 100 million

## Sequential

```
Time taken:  3m53.31377251s
The number of primes:  5761455
```

## Threaded, Non Fair

```
Time taken by thread  0 :  15.248318599s
Time taken by thread  1 :  22.110497828s
Time taken by thread  2 :  26.67205913s
Time taken by thread  3 :  29.729364939s
Time taken by thread  4 :  32.530676501s
Time taken by thread  5 :  34.995951909s
Time taken by thread  6 :  36.146459206s
Time taken by thread  7 :  40.569081647s
Time taken by thread  8 :  40.82184847s
Time taken by thread  9 :  41.062439545s
Total time taken:  41.062563862s
Num primes found:  5761455
```
## Threaded, Fair

```
Time taken by thread  8 :  35.777029534s
Time taken by thread  5 :  35.777409976s
Time taken by thread  9 :  35.77817987s
Time taken by thread  0 :  35.778212907s
Time taken by thread  2 :  35.778024675s
Time taken by thread  3 :  35.777738873s
Time taken by thread  6 :  35.777560001s
Time taken by thread  7 :  35.777114464s
Time taken by thread  4 :  35.778190835s
Time taken by thread  1 :  35.777923891s
Total time taken:  35.778364539s
Num primes found:  5761455
```