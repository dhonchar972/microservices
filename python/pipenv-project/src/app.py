from flask import Flask
from flask_restful import Api, Resource, reqparse, marshal_with, fields, abort
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
api = Api(app)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///database.db'
db = SQLAlchemy(app)
db.create_all()

persons_parser = reqparse.RequestParser()  # validation
persons_parser.add_argument("name", type=str, help="Person name", required=True)
persons_parser.add_argument("age", type=int, help="Person age", required=True)

persons = {
    "persons": {
        1: {
            "name": "Tim",
            "age": 17
        },
        2: {
            "name": "Bob",
            "age": 27
        }
    }
}

resource_fields = {
    'name': fields.String,
    'age': fields.Integer,
}

class Persons(Resource):
    def get(self):
        return persons


class HelloWorld(Resource):
    def get(self):
        return {"date": "Hello!"}

    def post(self):
        return {"date": "Posted"}


class HelloWorldWithParameters(Resource):
    @marshal_with(resource_fields)
    def get(self, name, age):
        return {"name": name, "age": age}


api.add_resource(Persons, "/")
api.add_resource(HelloWorld, "/hello")
api.add_resource(HelloWorldWithParameters, "/hello/<string:name>/<int:age>")

if __name__ == "__main__":
    app.run(debug=True)
