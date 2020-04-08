from validator import is_valid, VERSION

# this is the maximum number that gets hexadecimal number with length 16
MAX_VALUE = 18446744073709551615


def mint(date: str, email: str, difficulty: int) -> str:

    # loop until the correct token is found
    for count in range(0, MAX_VALUE + 1):

        # construct token
        token = "%s:%s:%s:%s" % (VERSION, date, email, format(count, 'x'))

        # if token is valud then stop checking new ones
        if is_valid(token, date, email, difficulty):
            break

    # return result
    return token