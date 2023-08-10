from app import app, jsonify, request, Blueprint
from game_person.game_person import game_person
from faction.faction import faction
from stories.stories import stories

url = '/the_stone_sword/api/v1.0'

app.register_blueprint(game_person, url_prefix=url)
app.register_blueprint(faction, url_prefix=url)
app.register_blueprint(stories, url_prefix=url)