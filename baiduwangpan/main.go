package main

import(
    "io/ioutil"
    "net/http"
    "fmt"
    "strings"
    "os/exec"
    "runtime"
    "encoding/json"
    "strconv"
    "container/list"
) 

var commands = map[string]string{
    "windows": "start",
    "darwin":  "open",
    "linux":   "xdg-open",
}
// Token access_token object
type Token struct{
    ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	AccessToken string `json:"access_token"`
	SessionSecret string `json:"session_secret"`
	SessionKey string `json:"session_key"`
	Scope string `json:"scope"`
}
// ShareLinkDir 外链中的文件夹
type ShareLinkDir struct {
	Errno int `json:"errno"`
	RequestID int64 `json:"request_id"`
	ServerTime int `json:"server_time"`
	Title string `json:"title"`
	List []ShareLinkDirList `json:"list"`
}
// ShareLinkFile 外链中的文件
type ShareLinkFile struct {
	Errno int `json:"errno"`
	RequestID int64 `json:"request_id"`
	ServerTime int `json:"server_time"`
	List []ShareLinkFileList `json:"list"`
}
//ShareLinkDirList 文件夹列表
type ShareLinkDirList struct {
	FsID string `json:"fs_id"`
	ServerFilename string `json:"server_filename"`
	Size string `json:"size"`
	ServerMtime string `json:"server_mtime"`
	ServerCtime string `json:"server_ctime"`
	LocalMtime string `json:"local_mtime"`
	LocalCtime string `json:"local_ctime"`
	Isdir string `json:"isdir"`
	Category string `json:"category"`
	Path string `json:"path"`
	Md5 string `json:"md5"`
}
//ShareLinkFileThumbs 文件缩略图地址
type ShareLinkFileThumbs struct {
	URL1 string `json:"url1"`
	URL2 string `json:"url2"`
	URL3 string `json:"url3"`
	Icon string `json:"icon"`
}
//ShareLinkFileList 文件列表
type ShareLinkFileList struct {
	FsID int64 `json:"fs_id"`
	Path string `json:"path"`
	ServerFilename string `json:"server_filename"`
	Size int `json:"size"`
	ServerMtime int `json:"server_mtime"`
	ServerCtime int `json:"server_ctime"`
	LocalMtime int `json:"local_mtime"`
	LocalCtime int `json:"local_ctime"`
	Isdir int `json:"isdir"`
	Category int `json:"category"`
	Md5 string `json:"md5"`
	Thumbs ShareLinkFileThumbs `json:"thumbs"`
}

func httpGet(url string) []byte {
    resp, err := http.Get(url)
    defer resp.Body.Close()
    if err != nil {
        fmt.Printf("Error with GET url %s", url)
        panic(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error with read body from url %s",url)
        panic(err)
    }
    return body
}

// Open calls the OS default program for uri
//通过调用各个平台打开文件的命令来实现,如果系统没有安装对应文件的打开工具，cmd.Start应该会返回err
func openURI(uri string) error {
    run, ok := commands[runtime.GOOS]
    if !ok {
        return fmt.Errorf("don't know how to open uri on %s platform", runtime.GOOS)
    }
    cmd := exec.Command(run, uri)
    return cmd.Start()
}

// 首先获取code,通过code换取token
func authToken(ak,sk string) Token{
    //百度帐号接入参考https://developer.baidu.com/newwiki/dev-wiki/?t=1557733846879
    //调用浏览器打开授权页面获取code
    var code string
    var tk Token
    err :=openURI("https://openapi.baidu.com/oauth/2.0/authorize?response_type=code&client_id="+ak+
    "&redirect_uri=oob&scope=basic,netdisk&display=popup&force_login=1&confirm_login=1")
    if err != nil {
        fmt.Printf("open auth page fail")
        panic(err)
    }
    fmt.Printf("Please enter your auth code: ")
    fmt.Scanln(&code)
    /*通过code换取网页授权access_token
    jsonStr :=`
    {
        "expires_in":2592000,
        "refresh_token":"22.eee5220f2f283a00f5d976b9c33c2940.315360000.1891315406.590339695-17979131",
        "access_token":"21.6207ea06baa2847b62854e80d2346b78.2592000.1578547406.590339695-17979131",
        "session_secret":"",
        "session_key":"",
        "scope":"basic netdisk"
    }
    `*/
    token :=httpGet("https://openapi.baidu.com/oauth/2.0/token?grant_type=authorization_code&code="+code+
    "&client_id="+ak+
    "&client_secret="+sk+"&redirect_uri=oob")
    if err := json.Unmarshal(token, &tk); err != nil {
        fmt.Printf("Umarshal token fail")
        panic(err)
    }
    return tk
}
//通过外链验证码获取外链密码
func sharelinkPwd(shareid,surl,vcode string) string {
    client := &http.Client{}
    var req *http.Request
    req,_ =http.NewRequest("POST", 
    "https://pan.baidu.com/rest/2.0/xpan/share?method=verify?shareid="+shareid+"&surl="+surl,
    strings.NewReader("pwd="+vcode))

	//添加cookie，可以添加多个cookie
	// cookie1 := &http.Cookie{Name: "X-Xsrftoken",Value: "df41ba54db5011e89861002324e63af81", HttpOnly: true}
    // req.AddCookie(cookie1)
    
    //添加header
    req.Header.Add("User-Agent","pan.baidu.com")
    req.Header.Add("Referer","pan.baidu.com")
 
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get share link password fail")
        panic(err)
	}
	defer resp.Body.Close()  
    body, _ := ioutil.ReadAll(resp.Body)
    /*{"errno":0,"err_msg":"","request_id":9123308402920936109,
    "randsk":"MP%2FJFKWzBzE91bRmkRsukpzD0WYNBbQ3gB%2B5Bnpwdas%3D"}*/
    var str =string(body)
    return str[strings.Index(str,"randsk")+9 : len(str)-2]
}
//获取外链创建者UK
func shareLinkUK(shareid,surl,randsk string) string{
    body :=httpGet("https://pan.baidu.com/api/shorturlinfo?shareid="+shareid+"&shorturl=1"+surl+"&spd="+randsk)
    /*{"shareid":1231251527,"uk":4094053687,"dir":"","type":0,"prod_type":"share","page":1,"root":0,"third":0,
    "longurl":"shareid=1231251527&uk=4094053687","fid":0,"errno":-3,"expire_days":7,"ctime":1575960893,
    "expiredtype":596566,"fcount":1,"uk_str":"4094053687"}*/
    var str =string(body)
    uk := strings.Index(str, "uk")
    pos := strings.Index(str[uk:], ",")
    return str[uk+4:uk+pos]
}
func transferImpl(access_token,shareid,randsk,uk,flist,targetPath string){
    client := &http.Client{}
    var req *http.Request
    req,_ =http.NewRequest("POST", 
    "https://pan.baidu.com/share/transfer?access_token="+access_token+"&shareid="+shareid+"&from="+uk+
    "&sekey="+randsk,
    strings.NewReader("fsidlist=["+flist+"]&path="+targetPath))
    //添加header
    req.Header.Add("User-Agent","pan.baidu.com")
    req.Header.Add("Referer","pan.baidu.com")
    resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("transfer flist[ %s ] to %s fail",flist,targetPath)
        panic(err)
	}
	defer resp.Body.Close()  
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
//遍历文件夹下文件并转存
func transferSharelinkFiles(access_token string) {
    var shareid,surl,vcode string
    fmt.Printf("Please enter your shareid: ")
    fmt.Scanln(&shareid)
    fmt.Printf("Please enter your short url: ")
    fmt.Scanln(&surl)
    fmt.Printf("Please enter your verification code: ")
    fmt.Scanln(&vcode)
    randsk :=sharelinkPwd(shareid,surl,vcode)
    var shareDir ShareLinkDir
    directory :=httpGet("https://pan.baidu.com/share/list?shareid="+shareid+"&shorturl="+surl+"&sekey="+
    randsk+"&root=1")
    if err := json.Unmarshal(directory, &shareDir); err != nil {
        fmt.Printf("Umarshal share link file fail")
        panic(err)
    }
    uk :=shareLinkUK(shareid,surl,randsk)
    for _,sf:= range shareDir.List {
        if strings.Compare(sf.Isdir,"1")==0 {
            page :=1
            counter :=0
            for {
                body :=httpGet("https://pan.baidu.com/share/list?shareid="+shareid+"&shorturl="+surl+"&sekey="+
                randsk+"&dir="+sf.Path+"&page="+strconv.Itoa(page)+"&num=10000")
                var flist ShareLinkFile
                if err := json.Unmarshal(body, &flist); err != nil {
                    fmt.Printf("Umarshal share link file fail")
                    panic(err)
                }
                count :=len(flist.List)
                if count==0{
                    break;
                }

                if()
                page ++
                counter +=count
            }
            fmt.Println("page %d counter %d",page,counter)
        }else{
            fmt.Println(sf.Path+ "is file and ignored")
        } 
    }
}

func httpPost(url,params string) []byte {
    resp, err := http.Post(url,"application/x-www-form-urlencoded;charset=utf-8",strings.NewReader(params))
    defer resp.Body.Close()
    if err != nil {
        fmt.Printf("Error with POST url %s", url)
        panic(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error with read body from url %s",url)
        panic(err)
    }
   return body
}

func main(){
    //申请百度开发者中心应用
    // var ak="bNskj1G251ayiYf9yI2DRI4Q"
    // var sk="9VyfB19hMdeRFBLpmEGpTLPuc5VgvMAE"
    //获取token
    // tk :=authToken(ak,sk)
    /* 获取外链密码
    shareid可以通过浏览器f12在分享的时候监视到,如下：
    {"errno":0,
    "request_id":7969551761846580210,
    "shareid":1231251527,
    "link":"https:\/\/pan.baidu.com\/s\/1yl0Ak3paNXKClQvYPBJxNg",
    "shorturl":"https:\/\/pan.baidu.com\/s\/1yl0Ak3paNXKClQvYPBJxNg",
    "ctime":1575960894,"expiredType":7,"createsharetips_ldlj":"","premis":false}
    验证码直接拷贝页面提示框里即可
    */
    transferSharelinkFiles()
    // shareLinkUK("1231251527","yl0Ak3paNXKClQvYPBJxNg","MP%2FJFKWzBzE91bRmkRsukpzD0WYNBbQ3gB%2B5Bnpwdas%3D")
    // var flURI =`
    // https://pan.baidu.com/share/list?
    // shareid=1231251527&shorturl=yl0Ak3paNXKClQvYPBJxNg&sekey=S3l2qkR0bI1Iuc%2f2Jxv2g2a4vLQkzx9KQlztg%2b04YNQ=&root=1" -H "User-Agent: pan.baidu.com
    // `
    // flist :=httpPost("https://pan.baidu.com/rest/2.0/xpan/share?method=list")
}