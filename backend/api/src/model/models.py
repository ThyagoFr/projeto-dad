from ..db import db


class Reader(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(255))
    email = db.Column(db.String(255))
    age = db.Column(db.Integer())
    photo = db.Column(db.String(255))
    password = db.Column(db.String(255))
    comments = db.relationship("Comment", backref="reader", lazy=True)
    interests = db.relationship("Interest", backref="reader", lazy=True)


class Book(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(255))
    subtitle = db.Column(db.String(255))
    cover = db.Column(db.String(100))
    genre = db.Column(db.String(255))
    author = db.Column(db.String(255))
    summary = db.Column(db.Text)
    comments = db.relationship("Comment", backref="book", lazy=True)


class Comment(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    comment = db.Column(db.Text)
    rate = db.Column(db.Integer())
    book_id = db.Column(db.Integer, db.ForeignKey("book.id"), nullable=False)
    reader_id = db.Column(db.Integer, db.ForeignKey("reader.id"), nullable=False)


class Interest(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    book_id = db.Column(db.Integer, db.ForeignKey("book.id"), nullable=False)
    reader_id = db.Column(db.Integer, db.ForeignKey("reader.id"), nullable=False)


class RecoverPassword(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    email = db.Column(db.String(255))
    token = db.Column(db.String(10))
    retrieved = db.Column(db.DateTime)
