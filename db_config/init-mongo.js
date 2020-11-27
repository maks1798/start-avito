db.createUser(
    {
        user : "qst1798",
        pwd : "1234",
        roles : [
            {
                role : "readWrite",
                db : "start-avito"
            }
        ]
    }
)