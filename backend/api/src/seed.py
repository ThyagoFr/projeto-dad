import requests


def load_data():
    subjects = [
        "horror",
        "drama",
        "romance",
        "adventure",
        "fiction"
    ]
    print("Carregando dados para o banco...")
    for subject in subjects:
        response = requests.get("https://www.googleapis.com/books/v1/volumes?q=subject{0}".format(subject))
        for r in response.json()["items"]:
            print(r["volumeInfo"]["title"])
    print("Finalizado...")


load_data()