from chekIfPasswordCorr import checkIfPasswordCorr
def loginProcedure():
    login = input("What is your login: ")
    password = input("What is your password: ")

    passCorrect = checkIfPasswordCorr(login, password)
    if passCorrect:
        # TODO: What hapens if credentials are ok
        pass
    elif not passCorrect:
        #TODO: What hapens if credentials are not ok
        pass
