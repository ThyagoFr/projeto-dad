from flask import Flask
from flask_cors import CORS
from flask_jwt_extended import JWTManager
from .db import initialize_db
from .seed import load_data


def create_app():
    app = Flask(__name__)
    CORS(app)
    JWTManager(app)
    initialize_db(app)
    load_data()
    return app


if __name__ == "__main__":
    application = create_app()
    application.run(debug=True)
