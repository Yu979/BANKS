
package main

import (
	"encoding/csv"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"io"
	"os"

	"net/http"

	"log"

	"fmt"

	"gopkg.in/mgo.v2"

	_ "gopkg.in/mgo.v2/bson"

	"encoding/json"

	"strings"

	"time"

	"github.com/codegangsta/negroni"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/dgrijalva/jwt-go/request"
)

const (
	SecretKey = "welcome to wangshubo's blog"
)

func fatal(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

type UserCredentials struct {

	Username string `json:"username"`
	Password string `json:"password"`
}

type goods struct{

	Name string
	Money int
	Seat int
}

type User struct {

	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}



type Response struct {
	Data string `json:"data"`
}



type Token struct {
	Token string `json:"token"`
}

//JWT 验证
func StartServer() {

	http.HandleFunc("/login", LoginHandler)

	http.Handle("/resource", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),

	))

	log.Println("Now listening...")


}


const MAXVEX int = 10
const MAXWEIGHT int = 999

var shortTablePath = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT,MAXWEIGHT}

func main() {
	StartServer()
	router := mux.NewRouter().StrictSlash(true)

	//redis缓存机制
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	//写入值并可可设置过期时间//
	_, err = c.Do("SET", "mykey", "superYu", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	time.Sleep(8 * time.Second)

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	//restful暴露API
	router.HandleFunc("/", Index)


	//创建CSV文件
	f, err := os.Create("BANK.csv") //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f) //创建一个新的写入文件流
	data := [][]string{
		{"1", "transport machine", "200"},
		{"2", "fork", "500"},
		{"3", "desk", "700"},
		{"4", "pens", "400"},
		{"5", "box", "20"},
		{"6", "toy", "90"},
		{"7", "ATM", "100000"},
		{"8", "house", "5000000"},
		{"9", "computer", "5000"},
		{"10", "rubber", "3"},
	}
	w.WriteAll(data) //写入数据
	w.Flush()
	//打开座位图csv文件
	file, err:= os.Open("BANK.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	//输出座位图
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(record) // record has the type []string
	}


   //设置座位优先级

	a := make(chan int, 1)
	b := make(chan int, 1)
	y := make(chan int, 1)
    d := make(chan int, 1)
    e := make(chan int, 1)
    q := make(chan int, 1)
    g := make(chan int, 1)
    h := make(chan int, 1)
    i := make(chan int, 1)
    j := make(chan int, 1)
//给channel赋值
var p int
   fmt.Scanf( "%d",&p)

	switch p {
	case 8:
	a<-8
	case 7:
		b<-7
	case 9:
		y<-9
	case 3:
		d<-3
	case 2:
		e<-2
	case 4:
		q<-4
	case 1:
		g<-1
	case 6:
		h<-6
	case 5:
		i<-5
	case 10:
		j<-10
	}




	select {
	case <-a:
		fmt.Println("从house开始找")
	default:
		select {
		case <-b:
			fmt.Println("从ATM开始找")
		default:
			select {
			case <-y:
				fmt.Println("从computer开始找")
			default:
				select {
				case <-d:
					fmt.Println("从desk开始找")
				default:
					select {
					case <-e:
						fmt.Println("从fork开始找")
					default:
						select {
						case <-q:
							fmt.Println("从pens开始找")
						default:
							select {
							case <-g:
								fmt.Println("从transport machine开始找")
							default:
								select {
								case <-h:
									fmt.Println("从toy开始找")
								default:
									select {
									case <-i:
										fmt.Println("从box开始找")
									default:
										select {
										case <-j:
											fmt.Println("rubber")
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}


    //最优路径算法,以V0为起点寻找最短路径
	    graph := NewGraph()
		var TablePathMin int       //存放shortTablePath中,未遍历的最小结点的值
		var Vx int                 //存放shortTablePath中,未遍历的最小结点的下标
		var isgetPath [MAXVEX]bool //记录结点是否已经找到v0到vx的最小路径

		// 获取v0这一行的权值数组
		for v := 0; v < len(graph); v++ {
			shortTablePath[v] = graph[0][v]
		}
		shortTablePath[0] = 0
		isgetPath[0] = true
		//遍历v1 ~ v9
		for v := 1; v < len(graph); v++ {
			TablePathMin = MAXWEIGHT

			//找出shortTablePath中,未遍历的最小结点的值
			for w := 0; w < len(graph); w++ {
				if !isgetPath[w] && shortTablePath[w] < TablePathMin {
					Vx = w
					TablePathMin = shortTablePath[w]
				}
			}
			isgetPath[Vx] = true
			for j := 0; j < len(graph); j++ {
				if !isgetPath[j] && TablePathMin+graph[Vx][j] < shortTablePath[j] {
					shortTablePath[j] = TablePathMin + graph[Vx][j]
				}
			}
  //
			fmt.Println("遍历完V", v, "后:", shortTablePath)

		}

		//输出最短路径
		for i := 0; i < len(shortTablePath); i++ {
			fmt.Println("V0到V", i, "最小路径:", shortTablePath[i])
		}

	}


//登入
    log.Fatal(http.ListenAndServe(":8080", router))

}




func Index(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintln(writer, "Welcome!")

	//连接本地mongodb服务
	url:="mongodb://localhost"
	session,err:=mgo.Dial(url)

	if err!=nil{
		panic(err)
	}
	defer session.Close()

	//打开数据库BANK
	session.SetMode(mgo.Monotonic,true)

	c:=session.DB("BANK").C("goods")

	//插入goods
	c.Insert(&goods{"transport machine",200,1},
		&goods{"fork",500,2},&goods{"desk",700,3},&goods{"pens",400,4},&goods{"box",20,5},
		&goods{"toy",90,6},&goods{"ATM",100000,7},&goods{"house",5000000,8},
		&goods{"computer",5000,9},&goods{"rubber",3,10})
	//查找全部
	goodss := make([]goods, 20)
	err = c.Find(nil).All(&goodss)
	//输出
	fmt.Fprintln(writer,goodss)
}
//设置JWT 认证
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

	response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)
}

//登入
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	if strings.ToLower(user.Username) != "someone" {
//设置密码
		if user.Password != "ppaassssword" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
		}

	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims



	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error extracting the key")
		fatal(err)

	}

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		fatal(err)
	}
	response := Token{tokenString}
	JsonResponse(response, w)
}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
//设置座位图
func NewGraph() [MAXVEX][MAXVEX]int {
	var graph [MAXVEX][MAXVEX]int
	var v0 = [MAXVEX]int{0,MAXWEIGHT, 1, 5, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
	var v1 = [MAXVEX]int{1, 0, 3, 7, 5, MAXWEIGHT,MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
	var v2 = [MAXVEX]int{5, 3, 0, MAXWEIGHT, 1,MAXWEIGHT, 7, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
	var v3 = [MAXVEX]int{MAXWEIGHT, 7, MAXWEIGHT, 0, 2, MAXWEIGHT,MAXWEIGHT, 3, MAXWEIGHT, MAXWEIGHT}
	var v4 = [MAXVEX]int{MAXWEIGHT, 5, 1, 2, 0, 3, 6, 9, MAXWEIGHT}
	var v5 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, 7, MAXWEIGHT, 3, 0, MAXWEIGHT, 5, MAXWEIGHT,MAXWEIGHT}
	var v6 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 3, 6, MAXWEIGHT, 0, 2, 7,MAXWEIGHT}
	var v7 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 9, 5, 2, 0, 4,MAXWEIGHT}
	var v8 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 7, 4, 0,MAXWEIGHT}
	var v9=  [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 2,MAXWEIGHT , 0,MAXWEIGHT}

	graph[0] = v0
	graph[1] = v1
	graph[2] = v2
	graph[3] = v3
	graph[4] = v4
	graph[5] = v5
	graph[6] = v6
	graph[7] = v7
	graph[8] = v8
	graph[9] = v9
	return graph
}