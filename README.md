# bulrush-identify
Provide basic user authorization and authentication

```go
var Identify = identify.
	New().
	AddOptions(identify.FakeURLsOption([]string{`^/api/ignore$`, `^/api/gorm/mock`})).
	AddOptions(identify.FakeTokensOption([]string{})).
	AddOptions(identify.ModelOption(&identify.RedisModel{
		Redis: addition.Redis,
	})).
	Init(func(iden *identify.Identify) {
		iden.AddOptions(
			identify.AuthOption(func(ctx *gin.Context) (interface{}, error) {
				login := struct {
					UserName string `form:"username" json:"username" xml:"username" binding:"required"`
					Password string `form:"password" json:"password" xml:"password" binding:"required"`
					Type     string `form:"type" json:"type" xml:"type"`
				}{}
				// captcha := ctx.GetString("captcha")
				if err := ctx.ShouldBind(&login); err != nil {
					return nil, err
				}
				user := addition.GORMExt.Var("User")
				if err := addition.GORMExt.DB.Find(user, map[string]interface{}{"name": login.UserName, "password": login.Password}).Error; err != nil {
					return nil, errors.New("user authentication failed")
				}
				return user, nil
			}),
		)
	})
app.Use(Identify)
```

#### ObtainToken
```curl
double@double:~/bulrush-template$ curl http://127.0.0.1:3001/api/v1/obtainToken -X POST -d "username=123&password=xxx"
{
  "accessToken":"E8V3TZX3XQEWN6E8B0VJTXK26Z8OF6F5",
  "refreshToken":"U1MPQIYQZKD86CYGWKDI04O4R5SK4FEW",
  "created":1559287003,
  "updated":1559287003,
  "expired":1559373403
}
```

#### RefleshToken
```curl
double@double:~/bulrush-template$ curl http://127.0.0.1:3001/api/v1/refleshToken -X POST -d "accessToken=U1MPQIYQZKD86CYGWKDI04O4R5SK4FEW"
{
  "accessToken":"0ILMFDM0JILCLJ5987HYMT2J18BYJDXI",
  "refreshToken":"U1MPQIYQZKD86CYGWKDI04O4R5SK4FEW",
  "created":1559287003,
  "updated":1559287592,
  "expired":1559373992,
  "extra":{
    "id":"5ce8ba5e51603f528e568b47","username":"123","roles":[]
  }
}
```

## MIT License

Copyright (c) 2018-2020 Double

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.