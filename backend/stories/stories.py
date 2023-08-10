from app import Blueprint, jsonify, request
from app.use_db import stories_db


stories = Blueprint('stories', __name__)


@stories.route('/stories', methods=['GET', 'POST'])
def all_stories():
    if request.method == 'POST':
        name_st = request.json["name_st"]
        content_st = request.json["content_st"]

        stories_db.add_stories(name_st, content_st)

        return jsonify(f'story {name_st} is added')

    elif request.method == 'GET':
        stories_json = []
        stories = stories_db.get_all_stories()

        for s in stories:
            stories_json.append(
                {
                    "id": s[0],
                    "name": s[1]
                }
            )

        return jsonify({
            "stories": stories_json
        })

