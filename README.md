目录结构  
app //应用文件目录  
├── module //应用模块 示例  
    ├── ── api  //控制器 中间件   
    ├── ── dao  //DAO层  
    ├── ── model //模型层  
    ├── ── service // 服务层  
conf // 配置文件目录  
├── conf.go  
db // 数据库连接文件目录  
├── mysql.go  
├── redis.go  
route //路由文件目录  
├── route.go  
static //静态资源文件 示例  
├── css  
├── js  
├── img  
├── uploads   
test //测试文件目录  
├── test_test.go  
web //静态页面目录  
├── index.html  
main.go  
由于现在基本是前后端分离模式开发，如果使用前端框架打包的文件。例如，vue框架直接把index.html放在web里，然后static里的文件也放static里即可  
//避免与前端一些框架冲突，例如vue框架 我把原来的渲染分隔符改成了这样 "{[{", "}]}"  

数据库框架使用的是gorm（github.com/jinzhu/gorm）  
redis 使用的包是go-redis（github.com/go-redis/redis）  
