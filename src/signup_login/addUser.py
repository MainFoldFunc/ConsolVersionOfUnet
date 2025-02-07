import requests
def addUser(login, password):
    url = "http://localhost:8080/signup" 
    dataJson = {
        "login": login,
        "password": password
    }

    response = requests.post(url, json=dataJson)
    print(f"Status code: {response.status_code}") 
    print(f"Response: {response.text}")
    if response.status_code == 200:
        return True
    else:
        return False
