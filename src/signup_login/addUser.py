import requests
import logging
def addUser(login, password):
    url = "http://localhost:8080/signup" 
    dataJson = {
        "login": login,
        "password": password
    }

    response = requests.post(url, json=dataJson)
    logging.info(f"Status code: {response.status_code}") 
    logging.info(f"Response: {response.text}")
    if response.status_code == 200:
        return True
    else:
        return False
