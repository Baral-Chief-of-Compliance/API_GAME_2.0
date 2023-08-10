from app.use_db.tools import quarry


def add_faction(*args):
    quarry.call("insert into factions(name_f, description_f) values "
                "(%s, %s)", *args, commit=True, fetchall=False)


def all_factions():
    factions = quarry.call("select * from factions", commit=False, fetchall=True)

    return factions


def inf_faction(id_f):
    inf = quarry.call("select * from factions where id_f = %s", commit=False, fetchall=False)
    return inf


def del_faction(id_f):
    quarry.call("delete from factions where id_f = %s", [id_f], commit=True, fetchall=False)


def update_faction(*args):
    quarry.call("update factions set name_f = %s, description_f = %s where id_f = %s",
                *args, commit=True, fetchall=False)