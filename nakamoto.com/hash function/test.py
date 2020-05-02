from hashlib import sha1
x = sha1(b"peter").hexdigest()
print(x)


# import binascii # to convert binary to ASCII bytes
# from hashlib import sha1

# x1 = sha1(binascii.b2a_uu(b"11111111")).hexdigest()
# print(x1)
# # 'a31518892830eec9ea21762e8bb101ce13890aee'

# # let's flip a single bit and see what happens

# x2 = sha1(binascii.b2a_uu(b"11011111")).hexdigest()
# print(x2)
# # 'c68b4d1bb5154c76370f33895d5d9350a4c73ba9'

