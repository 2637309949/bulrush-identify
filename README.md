# bulrush-identify
Provide basic user authorization and authentication.
- EXAMPLE:   
```go
app.Use(&identify.Identify{
    Auth: func(c *gin.Context) (interface{}, error) {
        var login binds.Login
        if err := c.ShouldBindJSON(&login); err != nil {
            return nil, err
        }
        if login.Password == "xx" && login.UserName == "xx" {
            return map[string]interface{}{
                "id":       "3e4r56u80a55",
                "username": login.UserName,
            }, nil
        }
        return nil, errors.New("user authentication failed")
    },
    Tokens: identify.TokensGroup{
        Save:   utils.Rds.Hooks.SaveToken,
        Revoke: utils.Rds.Hooks.RevokeToken,
        Find:   utils.Rds.Hooks.FindToken,
    },
    FakeURLs: []interface{}{`^/api/v1/ignore$`, `^/api/v1/docs/*`, `^/public/*`, `^/api/v1/ptest$`},
})
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