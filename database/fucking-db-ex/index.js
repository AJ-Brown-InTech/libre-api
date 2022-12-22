const pg  = require('pg')
let host =  "localhost"
let port  =  5432
let user =  "postgres"
let password  = 3596
let dbname = "postgres"

const connectDb = async () => {
        try {
            const pool = new pg.Pool({
                user: user,
                host: host,
                database: dbname,
                password: password,
                port: port,
            })
    await pool.connect()
           
        } catch (error) {
            console.log(error)
        }
}

connectDb().then((res)=> {console.log("success")})