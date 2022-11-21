import requests

BASE = "http://127.0.0.1:5000/"

response = requests.get(BASE + "hello")
print(response.json())

response = requests.get(f"{BASE}hello/Bob/27")
print(response.json())
