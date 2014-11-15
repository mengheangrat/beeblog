package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

var rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/comments","description":"oprations for Comments\n"},{"path":"/posts","description":"oprations for Posts\n"},{"path":"/users","description":"oprations for Users\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
var subapi string = `{"/comments":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/comments","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"Post","type":"","summary":"create Comments","parameters":[{"paramType":"body","name":"body","description":"\"body for Comments content\"","dataType":"Comments","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Comments.Id","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get Comments by id","parameters":[{"paramType":"path","name":"id","description":"\"The key for staticblock\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Comments","responseModel":"Comments"},{"code":403,"message":":id is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"Get All","type":"","summary":"get Comments","parameters":[{"paramType":"query","name":"query","description":"\"Filter. e.g. col1:v1,col2:v2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"fields","description":"\"Fields returned. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"sortby","description":"\"Sorted-by fields. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"order","description":"\"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"limit","description":"\"Limit the size of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"offset","description":"\"Start position of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Comments","responseModel":"Comments"},{"code":403,"message":"","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"PUT","nickname":"Update","type":"","summary":"update the Comments","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"body for Comments content\"","dataType":"Comments","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Comments","responseModel":"Comments"},{"code":403,"message":":id is not int","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"DELETE","nickname":"Delete","type":"","summary":"delete the Comments","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success!","responseModel":""},{"code":403,"message":"id is empty","responseModel":""}]}]}],"models":{"Comments":{"id":"Comments","properties":{"Comment":{"type":"string","description":"","format":""},"CommenterId":{"type":"Users","description":"","format":""},"CreatedAt":{"type":"\u0026{time Time}","description":"","format":""},"Id":{"type":"int","description":"","format":""},"PostId":{"type":"Posts","description":"","format":""}}},"Posts":{"id":"Posts","properties":{"CreatedAt":{"type":"\u0026{time Time}","description":"","format":""},"Id":{"type":"int","description":"","format":""},"OwnerId":{"type":"Users","description":"","format":""},"Text":{"type":"string","description":"","format":""},"Title":{"type":"string","description":"","format":""}}},"Users":{"id":"Users","properties":{"CreatedAt":{"type":"\u0026{time Time}","description":"","format":""},"Id":{"type":"int","description":"","format":""},"Username":{"type":"string","description":"","format":""}}}}},"/posts":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/posts","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"Post","type":"","summary":"create Posts","parameters":[{"paramType":"body","name":"body","description":"\"body for Posts content\"","dataType":"Posts","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Posts.Id","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get Posts by id","parameters":[{"paramType":"path","name":"id","description":"\"The key for staticblock\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Posts","responseModel":"Posts"},{"code":403,"message":":id is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"Get All","type":"","summary":"get Posts","parameters":[{"paramType":"query","name":"query","description":"\"Filter. e.g. col1:v1,col2:v2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"fields","description":"\"Fields returned. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"sortby","description":"\"Sorted-by fields. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"order","description":"\"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"limit","description":"\"Limit the size of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"offset","description":"\"Start position of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Posts","responseModel":"Posts"},{"code":403,"message":"","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"PUT","nickname":"Update","type":"","summary":"update the Posts","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"body for Posts content\"","dataType":"Posts","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Posts","responseModel":"Posts"},{"code":403,"message":":id is not int","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"DELETE","nickname":"Delete","type":"","summary":"delete the Posts","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success!","responseModel":""},{"code":403,"message":"id is empty","responseModel":""}]}]}],"models":{"Posts":{"id":"Posts","properties":{"CreatedAt":{"type":"\u0026{time Time}","description":"","format":""},"Id":{"type":"int","description":"","format":""},"OwnerId":{"type":"Users","description":"","format":""},"Text":{"type":"string","description":"","format":""},"Title":{"type":"string","description":"","format":""}}},"Users":{"id":"Users","properties":{"CreatedAt":{"type":"\u0026{time Time}","description":"","format":""},"Id":{"type":"int","description":"","format":""},"Username":{"type":"string","description":"","format":""}}}}},"/users":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/users","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"Post","type":"","summary":"create Users","parameters":[{"paramType":"body","name":"body","description":"\"body for Users content\"","dataType":"Users","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Users.Id","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get Users by id","parameters":[{"paramType":"path","name":"id","description":"\"The key for staticblock\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Users","responseModel":"Users"},{"code":403,"message":":id is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"Get All","type":"","summary":"get Users","parameters":[{"paramType":"query","name":"query","description":"\"Filter. e.g. col1:v1,col2:v2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"fields","description":"\"Fields returned. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"sortby","description":"\"Sorted-by fields. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"order","description":"\"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"limit","description":"\"Limit the size of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"offset","description":"\"Start position of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Users","responseModel":"Users"},{"code":403,"message":"","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"PUT","nickname":"Update","type":"","summary":"update the Users","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"body for Users content\"","dataType":"Users","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Users","responseModel":"Users"},{"code":403,"message":":id is not int","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"DELETE","nickname":"Delete","type":"","summary":"delete the Users","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success!","responseModel":""},{"code":403,"message":"id is empty","responseModel":""}]}]}],"models":{"Users":{"id":"Users","properties":{"CreatedAt":{"type":"\u0026{time Time}","description":"","format":""},"Id":{"type":"int","description":"","format":""},"Username":{"type":"string","description":"","format":""}}}}}}`
var rootapi swagger.ResourceListing

var apilist map[string]*swagger.ApiDeclaration

func init() {
	basepath := "/v1"
	err := json.Unmarshal([]byte(rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = basepath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
