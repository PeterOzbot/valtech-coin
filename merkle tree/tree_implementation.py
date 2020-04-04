from hashlib import sha256
import math

def hash(word: str)->str:
    return sha256(word.encode()).hexdigest()

def getAppropriateNumberOfWords(numberOfWords: int)->int:
    # issue with this if there is only one word
    # the result for one word is 1 which seems wrong
    # course for which this was written required it like this....
    closestToThePowerOfTwo = math.ceil(math.log(numberOfWords, 2))
    appropriateNumberOfWords = math.pow(2, closestToThePowerOfTwo)

    # round to cast float to int
    return math.floor(appropriateNumberOfWords) 

def merkleize(sentence: str) -> str:
    # null checking
    if not sentence:
        return None

    # split words
    splitSentence = sentence.strip().split()
    numberOfWords = len(splitSentence)

    # checking if there are any "words" in sentence
    if numberOfWords == 0:
        return None

    # hash the base layer
    hashedWords = []
    for word in splitSentence:
        hashedWords.append(hash(word))

    # padding if the blocks are not a power of 2
    appropriateNumberOfWords = getAppropriateNumberOfWords(numberOfWords)
    if numberOfWords != appropriateNumberOfWords:
        for _ in range(appropriateNumberOfWords - numberOfWords):
            hashedWords.append(chr(0))
            numberOfWords = numberOfWords + 1
    
    # get total lvls for a tree depending on words    
    maxLvl = math.floor(math.log(numberOfWords) / math.log(2))
    
    # initialize empty tree
    tree = [[]]
    for lvl in range(maxLvl):
        tree.append([])

    # set leaf
    tree[0] = hashedWords
        
    #build tree
    for lvl in range(maxLvl):
        for index in range(len(tree[lvl])):
            if index % 2 == 0:
                tree[lvl + 1].append(hash(tree[lvl][index] + tree[lvl][index + 1]))

    # return root
    return tree[maxLvl][0]

#print(merkleize("I love chicken! But I dont like pees."))
#d9e642064279e23d14443f446d071a5aa0f40257641d38d46b48a3111601c46b

#print(merkleize("Unbelievable..."))
#49527622a02f579bd6eb9b4bfadc2883cc22bfc5d9d05be936deb519eaa90b87

#print(merkleize("In our village, folks say God crumbles up the old moon into stars."))
#dddf9c7317f31d40714814749dac3a0c5ebab164262d49b576ed35f95850797a

print(merkleize("Hi there silly boy."))
#9a1706f7ce1acd8f536cfa70113812e7828ffda02f6da8ab3e724a0ad3e8b1f8