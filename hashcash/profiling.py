import cProfile
from validator import is_valid
from minting import mint
 
pr = cProfile.Profile()
pr.enable()

# for i in range(0, 100000):
#     is_valid("1:193040:val.idus@gmx.com:147d75", "193040", "val.idus@gmx.com", 20)
# print("---------------------------------------------------------------------------------------------")

print(mint("193040", "val.idus@gmx.com", 20))
print("---------------------------------------------------------------------------------------------")

pr.disable()
 
pr.print_stats(sort='time')