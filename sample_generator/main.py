import requests
import json
from datetime import datetime
import random
import time


def send_post_request(url, data):
    headers = {
        'Content-Type': 'application/json'
    }
    response = requests.post(url, headers=headers, data=json.dumps(data))
    if response.status_code == 200:
        return response.json()
    else:
        print(f"Error: Status code {response.status_code}\n{response.text}")
        return None

def main():
    url = 'http://10.17.42.2:8080/handle'
    tmp = 20
    moi = 60
    air = 1024
    while True:
        now = datetime.utcnow()
        formatted_date = now.strftime('%Y-%m-%dT%H:%M:%SZ')
        tmp = tmp + 3 - 6* random.random()
        moi = moi + 5 - 10* random.random()
        air = air + 5 - 10* random.random()
        data = {
            "surroundings": [
                {
                    "number": 1,
                    "timestamp": formatted_date,
                    "tempreture": tmp,
                    "moisuture": moi,
                    "airPressure": air,
                },
            ]
        }
        response = send_post_request(url, data)
        if response:
            print(response)
        else:
            print("POST request failed")
            break
        time.sleep(3)

    return

if __name__ == "__main__":
    main()