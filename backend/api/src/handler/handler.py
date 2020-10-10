from flask import Blueprint
from json import dumps
from datetime import datetime

errors_bp = Blueprint("errors", __name__)


class HandlerException(Exception):
    def __init__(self, message, status=None, payload=None):
        super().__init__()
        self.message = message
        if status is not None:
            self.status = status
        self.timestamp = str(datetime.now())
        self.payload = payload


@errors_bp.app_errorhandler(HandlerException)
def handler(exc):
    return dumps(exc.__dict__), exc.status
