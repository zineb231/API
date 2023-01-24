package main
//les packages
import(
		"os"
		"database/sql"
		"fmt"
		"log"
		_ "github.com/mattn/go-sqlite3"
)


var DB *sql.DB

//l'objet Document
type Document struct {
    ID  int
    nom  string
    description string
  
}


//initialisation de la base de données et création de la table
func init_database(){
	db, err := sql.Open("sqlite3", "Documents.db")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    query:=`CREATE TABLE IF NOT EXISTS documents(ID INTEGER PRIMARY KEY, nom VARCHAR(64) NULL, description text)`
	_, err = db.Exec(query)
 	if err != nil {
  		log.Printf("%q: %s\n", err, query)
  	return
 	}
}



//insérer des lignes dans la base de données(dans la table Documents)
func insertion(d Document){
	db, err := sql.Open("sqlite3", "Documents.db")
	q := "INSERT INTO `Documents` (ID,nom,description) VALUES (?, ?, ?);"
	insert, err := db.Prepare(q)
	if err != nil {
		fmt.Println(err)
		}

	insert.Exec(d.ID,d.nom,d.description)
	insert.Close()

}

//affichage du contenu de la table dans la base de données
func return_data(){
	db, err := sql.Open("sqlite3", "Documents.db")
	rows, err := db.Query("select ID,nom,description from Documents")
 	if err != nil {
  		log.Fatal(err)
 					}
 	defer rows.Close()
 	for rows.Next() {
  	var id int
  	var nom string
  	var description string
  	err = rows.Scan(&id, &nom,&description)
  	if err != nil {
   		log.Fatal(err)
  	}
  	fmt.Println(id, nom,description)
 		}
	 err = rows.Err()
 	if err != nil {
  	log.Fatal(err)
 	}
					}

//suppression par ID
func delete_by_id(id int){
	db, err := sql.Open("sqlite3", "Documents.db")
	_, err = db.Exec("delete from Documents where ID=?",id)
 	if err != nil {
  	log.Fatal(err)
 					}
						}






//fonction main
func main(){
	//initialisation de la base de données
	init_database()
	var ID int
	var name string
	var description string
	var num int
	fmt.Println("Entrez le nombre de documents à ajouter:")
	fmt.Scanf("%d",&num)
	for  i:=0; i<num ;i++ {
		fmt.Println("Entrez un document:")
		fmt.Println("ID")
		fmt.Scanf("%d",&ID)
		fmt.Println("nom")
		fmt.Scanf("%s", &name)
		fmt.Println("description")
    	fmt.Scanf("%s", &description)
		//insertion
		insertion(Document{ID,name,description})
						}
	//Affichage du contenu de la table
	return_data()
	//suppression par ID
	delete_by_id(14)
	fmt.Println("Après suppression")
	//Affichage après suppression
	return_data()


	

}