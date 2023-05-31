package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "github.com/gorilla/mux"
)

type Article struct {
    Id      string `json:"id"`
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// déclarons un tableau global d'articles
// que nous pourrons ensuite remplir dans notre fonction principale
// pour simuler une base de données
var Articles []Article

// fonction homePage gère toutes les requêtes à notre URL racine

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

//fonction returnAllArticles renvoie notre variable Articles nouvellement rense>

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

//Obtenir la valeur de l'id à partir de notre URL
//et nous pouvons renvoyer l'article qui correspond à ce critère
func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    fmt.Fprintf(w, "Key: " + key)

   // Boucle sur tous nos articles
    // si l'article.Id est égal à la clé que nous avons passée
    // retourne l'article encodé en JSON
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

//fonction pour ajouter un nouvel article
func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // récupérer le corps de notre requête POST
    // retourner la réponse sous forme de chaîne contenant le corps de la 
    //requête
    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Fprintf(w, "%+v", string(reqBody))

    var article Article 
    json.Unmarshal(reqBody, &article)
    
    // mettre à jour notre tableau global d'articles pour inclure
    // notre nouvel article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

//fonction deleteArticle  supprime les articles s'ils correspondent au paramètr>
//Id path donné

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    // une fois de plus, nous devrons analyser les paramètres du chemin d'accès.
    vars := mux.Vars(r)

    // nous devrons extraire le `id` de l'article que nous souhaitons supprimer
    id := vars["id"]

    // nous devons ensuite passer en boucle tous nos articles
    for index, article := range Articles {

    // si notre paramètre id path correspond à l'un de nos articles
    if article.Id == id {
        // met à jour notre tableau d'articles pour supprimer 
        //l'article
         Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}

//fonction updateArticle mettra à jour un article à l'aide de son id
func updateArticle(w http.ResponseWriter, r *http.Request) {
     // une fois de plus, nous devrons analyser les paramètres du chemin d'accè>
    vars := mux.Vars(r)

    // nous devrons extraire le `id` de l'article que nous souhaitons update
    id := vars["id"]

    var article Article

    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Fprintf(w, "%+v", string(reqBody))


    //reqBody, err := ioutil.ReadAll(r.Body)
    //if err != nil {
    //        fmt.Fprintf(w, "Kindly enter data with the event title and descri>
    //}
    json.Unmarshal(reqBody, &article) 

    // nous devons ensuite passer en boucle tous nos articles
    for index, singleArticle := range Articles {

        // si notre paramètre id path correspond à l'un de nos articles
        if singleArticle.Id == id {

        // mettre à jour notre tableau global d'articles pour mettre à jour
        // notre nouvel article
        singleArticle.Title = article.Title
        singleArticle.Desc = article.Desc
        singleArticle.Content = article.Content
        Articles = append(Articles[:index],singleArticle)
        json.NewEncoder(w).Encode(singleArticle)
      }
  }
}

//fonction handleRequests fera correspondre le chemin 
//de l'URL avec une fonction définie

func handleRequests() {
    // crée une nouvelle instance d'un routeur mux
    myRouter := mux.NewRouter().StrictSlash(true)

    // remplacer http.HandleFunc par myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles)

    // NOTE : L'ordre est important ici ! Ceci doit être défini avant
    // l'autre point de terminaison `/article`.
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    // ajouter notre nouveau point de terminaison DELETE ici
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    // ajouter notre nouveau point de terminaison PUT ici
    myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PATCH")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)

    // enfin, au lieu de passer dans nil,
    // nous voulons passer notre routeur nouvellement créé comme deuxième
    // argument
    log.Fatal(http.ListenAndServe(":80", myRouter))
}

//fonction main qui lancera notre API

func main() {
//variable Articles est remplie avec des données que nous pourrons récupérer et>
        Articles = []Article{
        Article{Id: "1",Title: "Hello", Desc: "Article Description", Content: "12"},
        Article{Id: "2",Title: "Hello 2", Desc: "Article Description", Content: "13"},
 }
    handleRequests()

}


