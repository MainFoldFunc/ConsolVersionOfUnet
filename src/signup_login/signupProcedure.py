from addUser import addUser
def signupProcedure():
    run = True
    while run:
        setLogin = input("What would you like your username to be: ") 
        setPassword = input("Set your password: ")

        setOk = addUser(setLogin, setPassword)
        if setOk:
            print("Sign up succesfull now login again by re starting application")
            run = False
        elif not setOk:
            print("Signing up not correct")
