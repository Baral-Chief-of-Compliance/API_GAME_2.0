from app.use_db.tools import quarry


def get_all_stories():
    stories = quarry.call("select * from story", fetchall=True, commit=False)
    return stories


def add_stories(name_st, content_st):
    quarry.call("insert into story (name_st, content_st) value(%s, %s)",
                [name_st, content_st], commit=True, fetchall=False)
