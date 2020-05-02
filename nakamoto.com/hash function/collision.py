from hashlib import md5
import random
import string

def generateWord() -> str:
    availableChars = string.ascii_lowercase + string.ascii_uppercase + string.digits
    return ''.join(random.choice(availableChars) for _ in range(10))

def md125(s: str) -> str: # use this hash function to generate a collision
  return md5(s.encode()).hexdigest()[:8]

def generate_md125_collisions() -> (str, str):
    isCollision = False
    firstPreimage = "nakamotoPeter"
    secondPreimage = "nakamoto"
    #secondPreimage = "nakamoto492T1eA8NX"
    #secondPreimage = "nakamotoPL2Ja6uben"
    counter = 0

    while not isCollision:

        #firstPreimage = "nakamoto" + generateWord()
        secondPreimage = "nakamoto" + generateWord()

        firstHash = md125(firstPreimage)
        secondHash = md125(secondPreimage)

        counter = counter + 1
        if counter % 100000 == 0:
            print(counter)

        isCollision = firstHash == secondHash

    return (firstPreimage, secondPreimage)


# run collision detection
collision = generate_md125_collisions()

print(collision[0])
print(md125(collision[0]))
print(collision[1])
print(md125(collision[1]))