import requests


def scanservice_task(ip: str):
    data = {
        'image-name': 'scanservice',
        'tag': '1.0.3',
        'task-name': 'test-' + ip,
        'priority': '5',
        'params[]': [ip, '80,443,3306,53']
    }
    requests.post('http://10.0.21.229:8080/task/', data=data)
