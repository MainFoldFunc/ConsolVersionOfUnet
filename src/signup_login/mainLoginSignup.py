from loginProcedure import loginProcedure
from signupProcedure import signupProcedure
from signupOrLogin import signupOrLogin
def main():
    login = signupOrLogin()
    if login == True:
        loginProcedure()
    else:
        signupProcedure()

main()
