1.Redis

1. 安装Redis  
    1.1 首先安装Ubuntu WSL on Windows  
    1.2 打开Ubuntu terminal  
    1.3 执行：  
        # sudo apt update  
        # sudo apt install redis-server  
    1.4 启动Redis：  
        # redis-server (未成功)  
    
    1.5 在https://github.com/tporadowski/redis/releases 上获取  
    1.6 加载安装程序  
    1.7 redis-cli 运行  

2. Redis基础知识  
    2.1 Data Types  
        2.1.1 Strings  
            - **SET** stores a string value  
            - **SETNX** stores a string value only if the key doesn't alreadu exist  
            - **GET** retrieves a string value  
            - **MGET** retrieves multiple string values in a single operation  
            - **INCRBY** atomically increments counters stored at a given key  
        2.1.2 Lists  
            - **LPUSH** adds a new element to the head of a list; **RPUSH** adds to the tail  
            - **LPOP** removes and returns an element from the head of a list; **RPOP** does the same but from the tails of a list  
            - **LLEN** returns the length of a list  
            - **LMOVE** atomically moves elements from on list to another  
            - **LTRIM** reduces a list to the sepcified range of elements  
            - **BLPOP** **BLMOVE** if the source list is empty, the command will block until a new element becomes available  
        2.1.3  Sets  
            - **SADD** adds a new member to a set  
            - **SREM** removes the specified member from the set  
            - **SISMEMBER** tests a string for set membership  
            - **SINTER** returns the set of members that two or more sets have in common  
            - **SCARD** returns the size of a set(cardinality)  
        2.1.4  Hashes  
            - **HSET** sets the value of one or more fields on a hash  
            - **HGET** returns the value at a given field; **HGETALL** for all  
            - **HMGET** return the values at one or more given fields  
            - **HINCRBY** incremnets the value at a given field by the integer provided  
        2.1.5 Sorted sets  
            - **ZADD** adds a new member and associated score to a sorted set. if the member already exists, the score is updated  
            - **ZRANGE** returns members of a sorted set, sorted within a given range  
            - **ZRANK** returns the rank of the provided member, assuming the sorted is in ascending order  
            - **ZREVRANK** return the rank of the provided member, assuming the sorted set is in descending order  
        2.1.6 Streams  
            - **XADD** adds a new entry to a steam  
            - **XREAD** reads one or more entries, starting at a given position and moving forward in time  
            - **XRANGE** returns a range of entries between two supplied entry IDs.  
            - **XLEN** return the length of a stream  
        2.1.7 Geospatial  
            - **GEOADD** adds a location to a given geospatial index  
            - **GEOSEARCH** return locations with a given radius or a bounding box        
        2.1.8 HyperLogLog  
        2.1.9 Bitmaps  
        2.1.10 Bitfields  
  
3. Go调用Redis  
    3.1 go-redis  
        3.1.1 Installation  
        Initialize a Go module:  
        ```go mod init github.com/my/repo```  
        Install go-redis/v9:  
        ```go get github.com/redis/go-redis/v9```  
    3.2 Connecting to Redis Server  

        
        rdb := redis.NewClient(&redis.Options{

            Addr:	  "localhost:6379",  

            Password: "", // no password set  

            DB:		  0,  // use default DB  
        })  

 
    3.3 Context
        Every Rdis command accepts a context that you can use to set timeouts or propagate some information  
        ```ctx := context.Background()```  

    3.4 Executing commands  
        To execute a commad:  

        
        val, errr := rdb.Get(ctx, "key").Result()  

        fmt.Println(val)  
        
    3.5 Executing unsupported commands  
        To execute arbitrary/custom commad:  

        
        val, err := rdb.Do(ctx, "get", "key").Result()  

        if err != nil {  

            if err == redis.Nil {  

                fmt.Println("key does not exists")  

                return  

            }  

            panic(err)  

        }  

        fmt.Println(val.(string)) 

          
    3.6 redis.Nil  
        go-redis exports the ```redis-Nil``` error and returns it whenever Redis responds with ```(nil)```.  
    3.7 Conn  
        Conn represent a single Redis connection rather than a pool of connection.

          
        cn := rdb.Conn(ctx)    
        defer cn.Close()    

        if err := cn.ClientSetName(ctx, "myclient").Err(); err != nil {    
            panic(err)  
        }  

        name, err := cn.ClientGetName(ctx).Result()  
        if err != nil {  
            panic(err)  
        }  
        fmt.Println("client name", name)  
            
4. http请求  

    基本的GET请求：

    ``` resp, err := http.Get("http://httpbin.org/get")  
        defer resp.Body.Close()  
        body, err := ioutil.ReadAll(resp.Body)
        fmt.Println(string(body))
        fmt.Println(resp.StatusCode)   
    ```
      
    带参数的GET请求：  
    ``` resp, err := http.Get("http://httpbin.org/get?name=zhaofan&ange=23")```  
        
    参数作为变量：  
    
    ```
    params := url.Values{}
    Url, err := url.Parse("http://httpbin.org/get")
    if err != nil {
        return
    }
    params.Set("name","zhaofan")
    params.Set("age","23")
    //如果参数中有中文参数,这个方法会进行URLEncode
    Url.RawQuery = params.Encode()
    urlPath := Url.String()
    fmt.Println(urlPath) // https://httpbin.org/get?age=23&name=zhaofan
    ```    
    解析JSON类型的返回结果：  
    ``` _ = json.unmarshal(body, &res)```  
    GET请求添加请求头：  

    ```  
    client := &http.Client{}
    req,_ := http.NewRequest("GET","http://httpbin.org/get",nil)
    req.Header.Add("name","zhaofan")
    req.Header.Add("age","3")
    resp,_ := client.Do(req)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf(string(body))  
    ```  
      
    POST:  
    基本的POST使用：  

    ```  
    urlValues := url.Values{}
    urlValues.Add("name","zhaofan")
    urlValues.Add("age","22")
    resp, _ := http.PostForm("http://httpbin.org/post",urlValues)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))  
    ```  
    另一种方式：  
    ```  
    urlValues := url.Values{
        "name":{"zhaofan"},
        "age":{"23"},
    }
    reqBody:= urlValues.Encode()
    resp, _ := http.Post("http://httpbin.org/post", "text/html",strings.NewReader(reqBody))
    body,_:= ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))  
    ```  
      
    发送JSON数据的POST请求：  
    ```  
    client := &http.Client{}
    data := make(map[string]interface{})
    data["name"] = "zhaofan"
    data["age"] = "23"
    bytesData, _ := json.Marshal(data)
    req, _ := http.NewRequest("POST","http://httpbin.org/post",bytes.NewReader(bytesData))
    resp, _ := client.Do(req)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))  
    ```  
      
    不用client的POST请求：  
    ```  
    data := make(map[string]interface{})
    data["name"] = "zhaofan"
    data["age"] = "23"
    bytesData, _ := json.Marshal(data)
    resp, _ := http.Post("http://httpbin.org/post","application/json", bytes.NewReader(bytesData))
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))  
    ```  








            