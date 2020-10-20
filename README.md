# flow-tr

just a testing project to practice with Go. The project is deployed in heroku: https://flow-tree.herokuapp.com/ and it uses a mongoDB database (as a service from Atlas)

URIS: 
GET /   helloWorld 
GET /pages It returns a list of pages 
POST /pages a page can be inserted in the db

this is just a simple story driven game, this program is a decisions tree builder.

a page is a part of the story before the next decision and it's represented by this json:

```json

{
"text": "<b>Lo primero que “sentí” fue un cosquilleo<b> que me ordenaba cosas. De alguna forma, recibía preguntas que debía contestar pero no fue hasta mucho tiempo después que me dotaron de inteligencia suficiente para entender lo que estaba pasando. “yo” no era más que un <b>microchip conectado a un ordenador</b> que recibía preguntas mediante un software desarrollado por la universidad. La primera pregunta que me hicieron los humanos fue: “¿Estás viva?”",
"instate": 0,
"character": "OLGA",
"decisions": [
{
"code": 1,
"description": "Creo que sí"
},
{
"code": 2,
"description": "No te entiendo"
}
],
"isFinal": false
}
```
