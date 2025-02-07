import requests
import logging
def checkIfPasswordCorr(login, password):
    url = "http://localhost:8080/login"
    dataJson = {
        "login": login,
        "password": password
    }

    response =requests.post(url, json=dataJson)
    logging.info(f"Status code: {response.status_code}")
    logging.info(f"Response: {response.text}")
    if response.text == True:
        return True
    else:
        return False
