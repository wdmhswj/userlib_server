package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func postHandler_datastoreGet(w http.ResponseWriter, r *http.Request) {
	log.Println("postHandler_datastoreGet:")
	// 仅处理 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体中的数据
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// 处理 POST 数据
	log.Printf("Received POST data: %s", string(body))
	// 反序列化
	// 创建结构体变量，用于存储解码后的数据
	var key UUID
	// 使用解码器将 JSON 数据反序列化为多个结构体对象
	err = json.Unmarshal(body, &key)
	if err != nil {
		log.Println("Error decoding key:", err)
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	value, ok := datastoreGet(key)

	if !ok {
		log.Println("The value does not exist")
		http.Error(w, "The value does not exist", http.StatusBadRequest)
		return
	}
	// 返回响应
	w.WriteHeader(http.StatusOK)
	w.Write(value)
}

func postHandler_keystoreSet(w http.ResponseWriter, r *http.Request) {
	log.Println("postHandler_keystoreSet:")
	// 仅处理 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体中的数据
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	// 处理 POST 数据
	// 反序列化
	buffer := bytes.NewBuffer(body)
	decoder := json.NewDecoder(buffer)

	var key string
	var value PublicKeyType

	err = decoder.Decode(&key)
	if err != nil {
		log.Println("Error decoding key string:", err)
		http.Error(w, "Error decoding key string", http.StatusBadRequest)
		return
	}

	err = decoder.Decode(&value)
	if err != nil {
		log.Println("Error decoding value PublicKeyType:", err)
		http.Error(w, "Error decoding value PublicKeyType", http.StatusBadRequest)
		return
	}

	log.Printf("Received POST data: \n\tkey: %v\n\tvalue: %v", key, value)
	err = keystoreSet(key, value)
	w.WriteHeader(http.StatusOK)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("nil"))

}

func postHandler_datastoreSet(w http.ResponseWriter, r *http.Request) {
	log.Println("postHandler_datastoreSet:")
	// 仅处理 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体中的数据
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	// 处理 POST 数据
	// 反序列化
	buffer := bytes.NewBuffer(body)
	decoder := json.NewDecoder(buffer)

	var key UUID
	var value []byte
	err = decoder.Decode(&key)
	if err != nil {
		log.Println("Error decoding key UUID:", err)
		http.Error(w, "Error decoding key UUID", http.StatusBadRequest)
		return
	}

	err = decoder.Decode(&value)
	if err != nil {
		log.Println("Error decoding value []byte:", err)
		http.Error(w, "Error decoding value []byte", http.StatusBadRequest)
		return
	}

	log.Printf("Received POST data: \n\tkey: %v\n\tvalue: %v", key, "value: ...")
	datastoreSet(key, value)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("nil"))

}

func postHandler_keystoreGet(w http.ResponseWriter, r *http.Request) {
	log.Println("postHandler_keystoreGet:")
	// 仅处理 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体中的数据
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// 处理 POST 数据

	// 反序列化
	var key string
	err = json.Unmarshal(body, &key)
	if err != nil {
		log.Println("Error decoding key string:", err)
		http.Error(w, "Error decoding key string", http.StatusBadRequest)
		return
	}
	log.Printf("Received POST data: %v", key)
	value, ok := keystoreGet(key)

	if !ok {
		log.Println("The value does not exist")
		http.Error(w, "The value does not exist", http.StatusBadRequest)
		return
	}
	// 返回响应
	// 序列化
	data, err := json.Marshal(value)
	if err != nil {
		log.Println("Error json.Marsharl: " + err.Error())
		http.Error(w, "Error json.Marsharl: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func postHandler_datastoreDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("postHandler_datastoreDelete:")
	// 仅处理 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体中的数据
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// 处理 POST 数据

	// 反序列化
	var key UUID
	err = json.Unmarshal(body, &key)
	if err != nil {
		log.Println("Error decoding key UUID:", err)
		http.Error(w, "Error decoding key UUID", http.StatusBadRequest)
		return
	}
	log.Printf("Received POST data: %v", key)
	datastoreDelete(key)

	// 返回响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("nil"))
}

func postHandler_datastoreDeleteAll(w http.ResponseWriter, r *http.Request) {
	log.Println("postHandler_datastoreDeleteAll:")
	// 仅处理 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 清除 datastore 中的所有数据
	datastoreClear()

	// 返回响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("nil"))
}

func postHandler_keystoreDeleteAll(w http.ResponseWriter, r *http.Request) {
	log.Println("postHandler_keystoreDeleteAll:")
	// 仅处理 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 清除 keystore 中的所有数据
	keystoreClear()

	// 返回响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("nil"))
}
