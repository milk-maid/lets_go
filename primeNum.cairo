def main(upper_limit: felt) -> ():
    for i in range(2, upper_limit + 1):
        if is_prime(i):
            debug_print(i)

def is_prime(n: felt) -> (bit):
    if n < 2:
        return 0
    for i in range(2, sqrt(n) + 1):
        if n % i == 0:
            return 0
    return 1
