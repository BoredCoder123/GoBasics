package main

import  (
    "github.com/go-sql-driver/mysql"
    "fmt"
    "log"
    "database/sql"
)

var db *sql.DB

type Album struct {
    ID int64
    Title string
    Artist string
    Price float32
}

func main() {
    cfg := mysql.Config{
        User: "root",
        Passwd: "password",
        Net: "tcp",
        Addr: "127.0.0.1:3306",
        DBName: "recordings",
    }

    var err error
    db, err = sql.Open("mysql" , cfg.FormatDSN())
    if err!= nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil{
        log.Fatal(pingErr)
    }
    fmt.Println("Connected")

    albums, err := albumsByArtist("John Coltrane")
    if err != nil{
        log.Fatal(err)
    }
    for i, ele := range albums{
        fmt.Printf("%v: %v %v %v %v\n", i, ele.ID, ele.Title, ele.Artist, ele.Price)
    }

    album, err := albumById(1);
    if err != nil{
        log.Fatal(err)
    }
    fmt.Printf("%d %v %v %v", album.ID, album.Title, album.Artist, album.Price)

    var insAlb Album
    insAlb = Album{
        Title: "Test",
        Artist: "Test",
        Price: 123.0,
    }
    lastInsId, err := addAlbum(insAlb)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Last inserted element id = %v", lastInsId)
}

func albumsByArtist(name string) ([]Album, error){
    var albums []Album
    rows, err := db.Query("select * from album where artist = ?", name)
    if err != nil{
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }

    defer rows.Close()

    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err!= nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err!= nil{
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}

func albumById(id int64) (Album, error){
    var album Album

    row := db.QueryRow("select * from album where id=?", id)

    if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err!= nil {
        return album, fmt.Errorf("albumById %d: %v", id, err)
    }

    return album, nil;
}

func addAlbum(album Album) (int64, error) {
    result, err := db.Exec("insert into album(title, artist, price) values (?, ?, ?)", album.Title, album.Artist, album.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum error: %v", err)
    }

    id, err := result.LastInsertId()

    if err != nil {
        return 0, fmt.Errorf("addAlbum error: %v", err)
    }

    return id, nil
}
