import requests

def checkIfPasswordCorr(login, password):
    url = "http://localhost:8080/login"
    dataJson = {
        "login": login,
        "password": password
    }

    response = requests.post(url, json=dataJson)
    print(f"Status code: {response.status_code}")
    print(f"Response: {response.text}")

    if response.status_code == 200:
        # Parse the JSON response
        response_json = response.json()
        if response_json.get("status") == "success":
            return True
        else:
            return False
    else:
        # Invalid login or other error
        return False

