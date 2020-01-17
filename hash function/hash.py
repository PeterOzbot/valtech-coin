def hash(n):
    return n % (2 ** 126)

x = hash(100000000000000000)
print(x)



