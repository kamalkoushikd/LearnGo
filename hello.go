package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"os"
	"strings"
)

// import "encoding/json"
type User struct {
	Username string `json:"username"`
	Time int64 `json:"time"`
	Authorizer string `json:"authorizer"`
}

func main(){
	// x := 2;
	print("Initialized Client\n");
	users := []User{};
	for {
		user := User{}
		reader := bufio.NewReader(os.Stdin);
		fmt.Print("Enter username: ");
		username, _ := reader.ReadString('\n');
		user.Username = strings.TrimSpace(username);
		// fmt.Print("Enter time: ");
		user.Time = time.Now().Unix();
		fmt.Print("Enter Authorizer: ");
		authorizer, _ := reader.ReadString('\n');
		user.Authorizer = strings.TrimSpace(authorizer);
		users = append(users, user);
		if user.Username == "exit" {
			break;
		}

		// data, := json.Marshal(user);
		jsonData, jerr := json.Marshal(user);
		if jerr != nil {
			fmt.Println("Error: ", jerr);
			return;
		}
		buf := bytes.NewBuffer(jsonData);
		req, _ := http.NewRequest("POST", "http://gotest.requestcatcher.com/test", buf);
		fmt.Println(buf);
		req.Header.Set("Content-Type", "application/json");
		req.Header.Set("Authorization", "Bearer token");
		client := &http.Client{};
		resp, err := client.Do(req);
		// fmt.Println("jsonData: ", bytes.NewBuffer(jsonData));
		if err != nil {
			fmt.Println("Error: ", err);
			return;
		}
		defer resp.Body.Close();
		body, err := io.ReadAll(resp.Body);
		// fmt.Print("Body: ", resp.Body);
		if err != nil {
			fmt.Println("Error: ", err);
			return;
		}
		fmt.Println("hello, World!");
		fmt.Println("Status", resp.Status);
		fmt.Println("Body: ", string(body));
		// fmt.Println(users);
		data, _ := json.Marshal(users);
		fmt.Println("Data: ", string(data));
	}
	
}
