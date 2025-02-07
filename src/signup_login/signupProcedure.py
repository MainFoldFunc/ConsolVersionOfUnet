from addUser import addUser
def signupProcedure():
    setLogin = input("What would you like your username to be: ") 
    setPassword = input("Set your password: ")

    setOk = addUser(setLogin, setPassword)
    if setOk:
        #TODO: What to do if signup is ok
        pass
    elif not setOk:
        #TODO: what to do if function returned false
        pass
