from ..db import db
from ..model.models import Book


class BookService:
    @staticmethod
    def store(data):
        bk = Book()
        bk.author = data["authors"][0]
        bk.genre = data["categories"][0]
        bk.title = data["volumeInfo"]["title"]
        bk.summary = data["volumeInfo"]["description"]
        db.session.add(bk)
        db.session.commit()
