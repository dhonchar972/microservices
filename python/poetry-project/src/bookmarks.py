from flask import Blueprint


bookmarks = Blueprint("bookmarks", __name__, url_prefix="/api/v1/bookmarks")

@bookmarks.post("/")
def get_all():
    return {"bookmarks": []}
