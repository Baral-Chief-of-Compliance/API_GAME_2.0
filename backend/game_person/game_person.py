from app import Blueprint, jsonify, request
from app.use_db import game_person_db


game_person = Blueprint('game_person', __name__)


@game_person.route('/game_persons', methods=['GET', 'POST'])
def all_game_persons():
    if request.method == 'POST':
        name_gp = request.json['name_gp']
        type_gp = request.json['type_gp']

        game_person_db.create_game_person([name_gp, type_gp])

        return jsonify(f"{name_gp} is add")

    elif request.method == 'GET':
        game_persons_json = []
        game_persons = game_person_db.get_all_game_person()

        for gp in game_persons:
            game_persons_json.append({
                "id": gp[0],
                "name": gp[1],
                "type_gp": gp[2]
            })

        return jsonify({
            "game_persons": game_persons_json
        })


@game_person.route('/game_persons/<int:id_gp>', methods=['GET', 'DELETE', 'PUT'])
def inf_game_person(id_gp):
    if request.method == 'GET':
        inf = game_person_db.inf_game_person(id_gp)

        return jsonify({
            "name_gp": inf[1],
            "type_gp": inf[2]
        })

    elif request.method == 'DELETE':
        game_person_db.del_game_person(id_gp)
        return jsonify(f"game person with {id_gp} is deleted")

    elif request.method == 'PUT':
        name_gp = request.json['name_gp']
        type_gp = request.json['type_gp']

        game_person_db.update_gp([name_gp, type_gp, id_gp])

        return jsonify(f"game person width {id_gp} is updated")


@game_person.route('/game_persons/<int:id_gp>/information', methods=['GET', 'POST', 'PUT'])
def game_person_information(id_gp):
    if request.method == 'GET':
        stories_json = []
        stories = game_person_db.get_stories(id_gp)

        for story in stories:
            stories_json.append(
                {
                    "id_stgp": story[0],
                    "confidence_gp": story[1],
                    "story": story[2]
                }
            )

        return jsonify({
            "stories": stories_json
        })