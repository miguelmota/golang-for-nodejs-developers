const sqlite3 = require('sqlite3').verbose()
const db = new sqlite3.Database('./sqlite3.db')

db.serialize(() => {
  db.run('CREATE TABLE persons (name TEXT)')

  const stmt = db.prepare('INSERT INTO persons VALUES (?)')
  const names = ['alice', 'bob', 'charlie']
  for (let i = 0; i < names.length; i++) {
    stmt.run(names[i])
  }

  stmt.finalize()

  db.each('SELECT rowid AS id, name FROM persons', (err, row) => {
    if (err) {
      console.error(err)
      return
    }

    console.log(row.id, row.name)
  })
})

db.close()
