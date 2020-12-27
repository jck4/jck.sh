package services
 
import (
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "log"
    "os"
    "net/http"
    "github.com/gin-gonic/contrib/sessions"
    "github.com/gin-gonic/gin"
    "golang.org/x/oauth2"
    "encoding/gob"
)
 
type Credentials struct {
    Cid     string `json:"cid"`
    Csecret string `json:"csecret"`
}
 
// User is a retrieved and authenticated user.
type User struct {
    Email string `json:"email"`
}
 
var cred Credentials
var conf *oauth2.Config
var state string
 
func randToken() string {
    b := make([]byte, 32)
    rand.Read(b)
    return base64.StdEncoding.EncodeToString(b)
}
 
func init() {
	file, err := ioutil.ReadFile("./creds.json")
	
    if err != nil {
        log.Printf("File error: %v\n", err)
        os.Exit(1)
	}
	
	json.Unmarshal(file, &cred)
	
	var Endpoint = oauth2.Endpoint{
		AuthURL:  "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token",
	}
    
    conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  githubAuthorizeUrl,
			TokenURL: githubTokenUrl,
		},
        RedirectURL:  "https://jck.sh/auth",
        Scopes: []string{"user:email"}
	}
}
 
 
func getLoginURL(state string) string {
    return conf.AuthCodeURL(state)
}
 
func AuthHandler(contex *gin.Context) {
    // Handle the exchange code to initiate a transport.
    session := sessions.Default(contex)
    retrievedState := session.Get("state")
    if retrievedState != contex.Query("state") {
        contex.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
        return
    }
 
    tok, err := conf.Exchange(oauth2.NoContext, contex.Query("code"))
    if err != nil {
        contex.AbortWithError(http.StatusBadRequest, err)
        return
    }
 
    client := conf.Client(oauth2.NoContext, tok)
    email, err := client.Get("https://api.github.com/user/emails")
    if err != nil {
        contex.AbortWithError(http.StatusBadRequest, err)
        return
    }
    defer email.Body.Close()
    data, _ := ioutil.ReadAll(email.Body)
    log.Println("Email body: ", string(data))
    contex.Status(http.StatusOK)
}
 
func LoginHandler(contex *gin.Context) {
    state = randToken()
    session := sessions.Default(contex)
    session.Set("state", state)
    session.Save()
    contex.Writer.Write([]byte("<html><title>Golang Github</title> <body> <a href='" + getLoginURL(state) + "'><button>Login with Github!</button> </a> </body></html>"))
}
