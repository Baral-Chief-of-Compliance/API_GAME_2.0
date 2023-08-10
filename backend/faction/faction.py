from app import Blueprint, jsonify, request
from app.use_db import faction_db


faction = Blueprint('faction', __name__)


@faction.route('/faction', methods=['GET', 'POST'])
def all_faction():
    if request.method == 'POST':
        name_f = request.json['name_f']
        description_f = request.json['description_f']

        faction_db.add_faction([name_f, description_f])

        return jsonify(f"faction {name_f} is added")

    elif request.method == 'GET':
        factions_json = []
        factions = faction_db.all_factions()

        for f in factions:
            factions_json.append({
                "id": f[0],
                "name": f[1]
            })

        return jsonify({
            "factions": factions_json
        })


@faction.route('/faction/<int:id_f>', methods=['GET', 'DELETE', 'PUT'])
def inf_faction(id_f):
    if request.method == 'GET':
        inf = faction_db.inf_faction(id_f)

        return jsonify({
            "id_f": inf[0],
            "name": inf[1],
            "description_f": inf[2]
        })

    elif request.method == "DELETE":
        faction_db.del_faction(id_f)
        return jsonify(f"faction with id_f {id_f} is deleted")

    elif request.method == 'PUT':
        name_f = request.json['name_f']
        description_f = request.json['description_f']

        faction_db.update_faction([name_f, description_f, id_f])

        return jsonify(f"faction {id_f} is updated")
