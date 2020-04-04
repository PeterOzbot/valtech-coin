from enum import Enum
from hashlib import sha256


class Side(Enum):
    LEFT = 0
    RIGHT = 1


def hash(word: str) -> str:
    return sha256(word.encode()).hexdigest()


def validate_proof(root: str, data: str, proof: [(str, Side)]) -> bool:
    nodeHash = hash(data)
    for itemHash, itemSide in proof:
        if itemSide == Side.LEFT:
            nodeHash = hash(itemHash + nodeHash)
        if itemSide == Side.RIGHT:
            nodeHash = hash(nodeHash + itemHash)

    return root == nodeHash


# print(hash("3ef3529a089c3b33c0a55a1d888e11dc58f3e36b4be29db74b06520f3b48a27c"+"835cf5feb8c46994f6b1d679fad81fed4c9783d6ee7df77936b3808cf73d163d"))

root = "9a1706f7ce1acd8f536cfa70113812e7828ffda02f6da8ab3e724a0ad3e8b1f8"
data = "silly"
proof = [("655b71a31e76c6fa4ca9e227cba95e50fac84f22283503a1d82fb6bc4c37f2d3", Side.RIGHT),
         ("3ef3529a089c3b33c0a55a1d888e11dc58f3e36b4be29db74b06520f3b48a27c", Side.LEFT)]
print(validate_proof(root, data, proof))
