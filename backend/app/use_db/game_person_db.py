from app.use_db.tools import quarry


def create_game_person(*args):
    quarry.call("insert into game_person(name_gp, type_gp) values "
                "(%s, %s)", *args, commit=True, fetchall=False)


def get_all_game_person():
    game_persons = quarry.call("select * from game_person order by game_person.id_gp DESC", commit=False, fetchall=True)
    return game_persons


def del_game_person(id_gp):
    quarry.call("delete from game_person where id_gp = %s", [id_gp], commit=True, fetchall=False)


def inf_game_person(id_gp):
    inf = quarry.call("select * from game_person where id_gp = %s", [id_gp], commit=False, fetchall=False)
    return inf


def update_gp(*args):
    quarry.call("update game_person set name_gp = %s, type_gp = %s where id_gp = %s",
                *args, commit=True, fetchall=False)