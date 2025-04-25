package main

import (
	"encoding/json"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
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
	for {
		user := User{}
		users := []User{};
		fmt.Print("Enter username: ");
		fmt.Scanf("%s", &user.Username);
		// fmt.Print("Enter time: ");
		user.Time = time.Now().Unix();
		fmt.Print("Enter Authorizer: ");
		fmt.Scanf("%s", &user.Authorizer);
		users = append(users, user);
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
	}
}
