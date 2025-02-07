import requests
def checkIfPasswordCorr(login, password):
    url = "http://localhost:8080/login"
    dataJson = {
        "login": login,
        "password": password
    }

    response =requests.post(url, json=dataJson)
    print(f"Status code: {response.status_code}")
    print(f"Response: {response.text}")
    if response.text == True:
        return True
    else:
        return False
