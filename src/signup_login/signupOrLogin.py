def signupOrLogin():
    loginOrSignup = input("Hello welcome to Unet, Do you have an account: ")
    if(loginOrSignup.lower() == "yes"):
        return True
    else:
        return False

