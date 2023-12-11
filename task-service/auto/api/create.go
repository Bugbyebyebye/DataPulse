package api

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func CreateAPi(portStr string, pathStr string, funcName string) {

	port := flag.String(portStr, portStr, "Port number to run the server on")
	path := flag.String(pathStr, pathStr, "Path to serve")

	flag.Parse()
	// 创建JavaScript文件的内容
	jsCode := fmt.Sprintf(`
    const http = require('http');
	const mysql = require('mysql');
	const conn = mysql.createConnection({
		host: '222.186.50.126',
		user: 'root',
		password: 'maojiukeai1412',
		port: '20134',
		database: 'df_system'
	});
	conn.connect();
	
	const server = http.createServer((req, res) => {
		if (req.method === 'GET' && req.url === '%s/get') {
			res.writeHead(200, { 'Content-Type': 'text/plain;charset=utf-8' });
			new Promise((resolve, reject) => {
				conn.query('select * from t_user', function (error, results, fields) {
					if (error) {
						console.log(error);
						throw(error);
					} else {
						console.log(results);
						resolve(JSON.stringify(results));
					}
				});
			}).then(data => {
				res.write(data);
				res.end();
			}).catch(err => {
				console.error(err);
				res.writeHead(500, { 'Content-Type': 'text/plain;charset=utf-8' });
				res.end('Internal Server Error');
			});
		} else if (req.method === 'GET' && req.url === '%s') {
			res.writeHead(200, { 'Content-Type': 'text/plain;charset=utf-8' });
			res.end('Hello from Go-generated Node.js server!');
		} else {
			res.writeHead(404);
			res.end();
		}
	});
	
	server.listen(%s, 'localhost', () => {
		console.log('Server running at localhost:3000');
	});
    `, *path, *path, *port)

	// 将JS代码写入指定的文件夹（例如：./node/server.js）
	jsFilePath := "/home/sora/项目/期末作业项目/软件系统管理/DataPulse/task-service/auto/api/src" + funcName + ".js"
	err := ioutil.WriteFile(jsFilePath, []byte(jsCode), 0644)
	if err != nil {
		panic(err)
	}

	// 使用Node.js运行生成的JS文件
	cmd := exec.Command("node", jsFilePath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("err => %s", err)
	}

	fmt.Println(string(out))
}

func RunDocker(portStr string, funcName string) {
	// 设置环境变量
	setEnv("Host", "222.186.50.126")
	setEnv("Port", "20134")
	setEnv("Username", "root")
	setEnv("Password", "maojiukeai1412")
	setEnv("Database", "df_system")
	setEnv("ServerPort", portStr)
	portMapping := portStr + ":" + portStr

	// 构建 Docker 命令
	cmd := exec.Command("podman", "run", "--env", "Port", "--env", "Host", "--env", "Username", "--env", "Password", "--env", "Database", "--env", "ServerPort", "-p", portMapping, "autodocker")

	// 执行 Docker 命令
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running Docker command:", err)
		return
	}

	// 打印命令输出
	fmt.Println(string(output))
}

func setEnv(key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		fmt.Printf("Error setting environment variable %s: %s\n", key, err)
	}
}
