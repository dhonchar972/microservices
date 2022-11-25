from src.constants.http_status_codes import HTTP_400_BAD_REQUEST, HTTP_201_CREATED
from src.db import Bookmark, db
from flask import Blueprint, request
import validators
from flask.json import jsonify
from flask_jwt_extended import get_jwt_identity, jwt_required

bookmarks = Blueprint("bookmarks", __name__, url_prefix="/api/v1/bookmarks")


@bookmarks.route('/', methods=['POST', 'GET'])
@jwt_required
def bookmarks_route():
    current_user = get_jwt_identity()

    if request.method == "POST":
        body = request.get_json().get("body", "")
        url = request.get_json().get("url", "")

        if not validators.url(url):
            return jsonify({"error": "Enter a valid url"}), HTTP_400_BAD_REQUEST

        if Bookmark.query.filter_by(url=url).first():
            return jsonify({"error": "Enter a valid url"}), HTTP_400_BAD_REQUEST

        bookmark = Bookmark(url=url, body=body, user_id=current_user)
        db.session.add(bookmark)
        db.session.commit()

        return jsonify({
            'id': bookmark.id,
            'url': bookmark.url,
            'short_url': bookmark.short_url,
            'visit': bookmark.visits,
            'body': bookmark.body,
            'created_at': bookmark.created_at,
            'updated_at': bookmark.updated_at,
        }), HTTP_201_CREATED

    else:
        bms = Bookmark.query.filter_by(user_id=current_user)

        data = []

        for bm in bms:
            data.append({
                'id': bm.id,
                'url': bm.url,
                'short_url': bm.short_url,
                'visit': bm.visits,
                'body': bm.body,
                'created_at': bm.created_at,
                'updated_at': bm.updated_at,
            })
