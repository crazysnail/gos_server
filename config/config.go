package config

import (
	"bufio"
	"log"
	"os"
	"io"
	"strings"
	"sync"
)

var _DEF_CONFIG = os.Getenv("GOS_ROOT") + "/config/config.ini"

var (
	_map  map[string]string
	_lock sync.Mutex
)

func init() {
	Reload()
}

func Get() map[string]string {
	_lock.Lock()
	defer _lock.Unlock()
	return _map
}

func Reload() {
	var path = _DEF_CONFIG
	_lock.Lock()
	log.Println("Loading Config from:", path)
	defer log.Println("Config Loaded.")
	_map = LoadConfig(path)
	_lock.Unlock()
}

func LoadConfig(path string) map[string]string {
    //初始化
    myMap := make(map[string]string)

    //打开文件指定目录，返回一个文件f和错误信息
    f, err := os.Open(path)

    //异常处理 以及确保函数结尾关闭文件流
    if err != nil {
        panic(err)
    }
    defer f.Close()

    //创建一个输出流向该文件的缓冲流*Reader
    r := bufio.NewReader(f)
    for {
        //读取，返回[]byte 单行切片给b
        b, _, err := r.ReadLine()
        if err != nil {
            if err == io.EOF {
                break
            }
            panic(err)
        }

        //去除单行属性两端的空格
        s := strings.TrimSpace(string(b))
        //fmt.Println(s)

        //判断等号=在该行的位置
        index := strings.Index(s, "=")
        if index < 0 {
            continue
        }
        //取得等号左边的key值，判断是否为空
        key := strings.TrimSpace(s[:index])
        if len(key) == 0 {
            continue
        }

        //取得等号右边的value值，判断是否为空
        value := strings.TrimSpace(s[index+1:])
        if len(value) == 0 {
            continue
        }
        //这样就成功吧配置文件里的属性key=value对，成功载入到内存中c对象里
        myMap[key] = value
    }
    return myMap
}
//
//func _load_config(path string) (ret map[string]string) {
//	ret = make(map[string]string)
//	f, err := os.Open(path)
//	if err != nil {
//		log.Println(path, err)
//		return
//	}
//
//	re := regexp.MustCompile(`[\t ]*([0-9A-Za-z_]+)[\t ]*=[\t ]*([^\t\n\f\r# ]+)[\t #]*`)
//
//	// using scanner to read config file
//	scanner := bufio.NewScanner(f)
//	scanner.Split(bufio.ScanLines)
//
//	for scanner.Scan() {
//		line := strings.TrimSpace(scanner.Text())
//		// expression match
//		slice := re.FindStringSubmatch(line)
//
//		if slice != nil {
//			ret[slice[1]] = slice[2]
//			log.Println(slice[1], "=", slice[2])
//		}
//	}
//
//	return
//}
