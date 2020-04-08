from hashlib import sha256

VERSION = "1"


def hash(word: str) -> str:
    return sha256(word.encode()).hexdigest()


def binary_leading_0s(hex_str: str):
    binary_representation = bin(int(hex_str, 16))[2:].zfill(256)
    return len(binary_representation) - len(binary_representation.lstrip('0'))


def is_valid(token: str, date: str, email: str, difficulty: int) -> bool:
    # validate date length
    if date is None or len(date) > 6:
        return False

    # split token to validate its parts
    tokenParts = token.split(':')
    if(len(tokenParts) != 4):
        return False
    tokenVersion, tokenDate, tokenEmail, tokenNonce = tokenParts

    # valdiate that date and email match function input
    if email != tokenEmail or date != tokenDate:
        return False

    # validate version
    if tokenVersion != VERSION:
        return False

    # validate nonce length
    if len(tokenNonce) > 16:
        return False

    # has the token and check leading zeros
    hex_str = hash(token)
    return binary_leading_0s(hex_str) == difficulty
