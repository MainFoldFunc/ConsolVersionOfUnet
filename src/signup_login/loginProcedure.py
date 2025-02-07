from checkIfPasswordCorr import checkIfPasswordCorr
def loginProcedure():
    run = True
    while run:
        login = input("What is your login: ")
        password = input("What is your password: ")

        passCorrect = checkIfPasswordCorr(login, password)
        if passCorrect:
            print("Correct login and password")
            run = False
        elif not passCorrect:
            print("Login not correct")
