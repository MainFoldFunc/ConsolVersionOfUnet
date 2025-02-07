import sys
import os

# Get the parent directory (the folder containing both 'checkIfPasswordCorr.py' and 'app' folder)
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))

# Now you can import from sibling folders
from checkIfPasswordCorr import checkIfPasswordCorr
from app.application import app  # No need for '../'

def loginProcedure():
    run = True
    while run:
        login = input("What is your login: ")
        password = input("What is your password: ")

        passCorrect = checkIfPasswordCorr(login, password)
        if passCorrect:
            print("Correct login and password")
            run = False
        else:
            print("Login not correct")
    
    app()  # Call app function after successful login

