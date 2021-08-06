
# Auth

## ResendAppTicket

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Auth.ResendAppTicket(ctx, &lark.ResendAppTicketReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

飞书每隔 1 小时会给应用推送一次最新的 app_ticket，应用也可以主动调用此接口，触发飞书进行及时的重新推送。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQjNz4CN2MjL0YzM](https://open.feishu.cn/document/ukTMukTMukTM/uQjNz4CN2MjL0YzM)

### URL

`https://open.feishu.cn/open-apis/auth/v3/app_ticket/resend/`

### Method

`POST`

## GetAccessToken

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Auth.GetAccessToken(ctx, &lark.GetAccessTokenReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

通过此接口获取登录预授权码 code 对应的登录用户身份。

该接口仅适用于通过网页应用登录方式获取的预授权码，小程序登录中用户身份的获取，请使用[小程序 code2session 接口](/ssl:ttdoc/uYjL24iN/ukjM04SOyQjL5IDN)


#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uEDO4UjLxgDO14SM4gTN](https://open.feishu.cn/document/ukTMukTMukTM/uEDO4UjLxgDO14SM4gTN)

### URL

`https://open.feishu.cn/open-apis/authen/v1/access_token`

### Method

`POST`

## RefreshAccessToken

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Auth.RefreshAccessToken(ctx, &lark.RefreshAccessTokenReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

user_access_token 具有一定的时效性，默认最长有效期为7200秒。该接口用于在 user_access_token 过期时用 refresh_token 重新获取 access_token。此时会返回新的 refresh_token，再次刷新 access_token 时需要使用新的 refresh_token。

调用该接口之后，之前的 user_access_token 即使没有到过期时间也会马上失效

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQDO4UjL0gDO14CN4gTN](https://open.feishu.cn/document/ukTMukTMukTM/uQDO4UjL0gDO14CN4gTN)

### URL

`https://open.feishu.cn/open-apis/authen/v1/refresh_access_token`

### Method

`POST`

## GetUserInfo

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Auth.GetUserInfo(ctx, &lark.GetUserInfoReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

此接口仅用于获取登录用户的信息。调用此接口需要在 Header 中带上 user_access_token。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uIDO4UjLygDO14iM4gTN](https://open.feishu.cn/document/ukTMukTMukTM/uIDO4UjLygDO14iM4gTN)

### URL

`https://open.feishu.cn/open-apis/authen/v1/user_info`

### Method

`GET`


# Contact

## CreateUser

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.CreateUser(ctx, &lark.CreateUserReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

使用该接口向通讯录创建一个用户，可以理解为员工入职。创建部门后只返回有数据权限的数据。具体的数据权限的与字段的对应关系请参照[应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)

新增用户的所有部门必须都在当前应用的通讯录授权范围内才允许新增用户，如果想要在根部门下新增用户，必须要有全员权限。 应用商店应用无权限调用此接口。







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users`

### Method

`POST`

## DeleteUser

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.DeleteUser(ctx, &lark.DeleteUserReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口向通讯录删除一个用户信息，可以理解为员工离职。

应用需要待删除用户的所有部门的通讯录权限才能删除该用户。应用商店应用无权限调用接口。用户可以在删除员工时设置删除员工数据的接收者，如果不设置则由其leader接收，如果该员工没有leader，则会将该员工的数据删除。







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/delete)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users/:user_id`

### Method

`DELETE`

## GetUser

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetUser(ctx, &lark.GetUserReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取通讯录中单个用户的信息。

<b>以应用身份访问通讯录</b> 权限为历史版本，不推荐申请。应用访问通讯录相关接口请申请 <b>以应用身份读取通讯录</b>


该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users/:user_id`

### Method

`GET`

## BatchGetUser

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.BatchGetUser(ctx, &lark.BatchGetUserReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于批量获取用户详细信息。

调用该接口需要申请“以应用身份读取通讯录”以及[用户数据权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)。请求的用户如果在当前应用的通讯录授权范围内，会返回该用户的详细信息；否则不会返回。



#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uIzNz4iM3MjLyczM](https://open.feishu.cn/document/ukTMukTMukTM/uIzNz4iM3MjLyczM)

### URL

`https://open.feishu.cn/open-apis/contact/v1/user/batch_get`

### Method

`GET`

## BatchGetUserByID

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.BatchGetUserByID(ctx, &lark.BatchGetUserByIDReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



根据用户邮箱或手机号查询用户 open_id 和 user_id，支持批量查询。<br>
调用该接口需要申请 `通过手机号或者邮箱获取用户ID` 权限。<br>只能查询到应用可用性范围内的用户 ID，不在范围内的用户会表现为不存在。



#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUzMyUjL1MjM14SNzITN](https://open.feishu.cn/document/ukTMukTMukTM/uUzMyUjL1MjM14SNzITN)

### URL

`https://open.feishu.cn/open-apis/user/v1/batch_get_id`

### Method

`GET`

## GetUserList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetUserList(ctx, &lark.GetUserListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



当应用的通讯录权限范围为全员时，传入 department_id，以获取该部门下所有的直属成员。根部门的 department_id 是 0。

当应用的通讯录权限范围为非全员时，传入权限范围内的 department_id，以获取该部门下所有的直属成员。不传 department_id 时，会获取到权限范围内的独立成员。（当权限范围包含了某成员，但不包含成员所在部门，则该成员视为权限范围内的独立成员）{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=user&method=list)




<b>以应用身份访问通讯录</b> 权限为历史版本，不推荐申请。应用访问通讯录相关接口请申请 <b>以应用身份读取通讯录</b>


该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/list)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users`

### Method

`GET`

## UpdateUserPatch

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.UpdateUserPatch(ctx, &lark.UpdateUserPatchReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新通讯录中用户的字段，未传递的参数不会更新。

该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/patch)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users/:user_id`

### Method

`PATCH`

## UpdateUser

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.UpdateUser(ctx, &lark.UpdateUserReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新通讯录中用户的字段。

应用需要拥有待更新用户的通讯录授权，如果涉及到用户部门变更，还需要同时拥有所有新部门的通讯录授权。应用商店应用无权限调用此接口。


该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/update)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users/:user_id`

### Method

`PUT`

## CreateDepartment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.CreateDepartment(ctx, &lark.CreateDepartmentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于向通讯录中创建部门。

只可在应用的通讯录权限范围内的部门下创建部门。若需要在根部门下创建子部门，则应用通讯录权限范围需要设置为“全部成员”。应用商店应用无权限调用此接口。


该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/create)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments`

### Method

`POST`

## GetDepartment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetDepartment(ctx, &lark.GetDepartmentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于向通讯录获取单个部门信息。

使用tenant_access_token时，应用需要拥有待查询部门的通讯录授权。如果需要获取根部门信息，则需要拥有全员权限。
使用user_access_token时，用户需要有待查询部门的可见性，如果需要获取根部门信息，则要求员工可见所有人。

<b>以应用身份访问通讯录</b> 权限为历史版本，不推荐申请。应用访问通讯录相关接口请申请 <b>以应用身份读取通讯录</b>


该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/get)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments/:department_id`

### Method

`GET`

## GetDepartmentList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetDepartmentList(ctx, &lark.GetDepartmentListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取当前部门子部门列表。

- 使用tenant_access_token时,只返回权限范围内的部门。

 - 使用user_access_token时，返回可见性范围内的所有可见部门。当进行递归查询时，只筛查最多1000个部门的可见性。

 - fetch_child字段填写为false：如果填写具体的部门ID，则返回该部门下的一级子部门；如果没有填写部门ID， 若有全员权限，返回根部门信息，若没有全员权限则返回通讯录范围中配置的部门及其一级子部门。

 - fetch_child字段填写为true：如果填写具体的部门ID，则返回该部门下所有子部门；如果没有填写部门ID， 若有全员权限，返回根部门信息，可以根据根部门ID获取其下的一级子部门，若没有全员权限则返回通讯录范围中配置的部门及其子部门。

<b>以应用身份访问通讯录</b> 权限为历史版本，不推荐申请。应用访问通讯录相关接口请申请 <b>以应用身份读取通讯录</b>


该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/list)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments`

### Method

`GET`

## GetParentDepartment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetParentDepartment(ctx, &lark.GetParentDepartmentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用来递归获取部门父部门的信息，并按照由子到父的顺序返回有权限的父部门信息列表。

使用tenant_access_token时,该接口只返回可见性范围内的父部门信息

例如：A >>B>>C>>D四级部门，通讯录权限只到B，那么查询D部门的parent，会返回B和C两级部门。
使用user_access_token时,该接口只返回对于用户可见的父部门信息

<b>以应用身份访问通讯录</b> 权限为历史版本，不推荐申请。应用访问通讯录相关接口请申请 <b>以应用身份读取通讯录</b>


该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/parent](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/parent)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments/parent`

### Method

`GET`

## SearchDepartment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.SearchDepartment(ctx, &lark.SearchDepartmentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

搜索部门，用户通过关键词查询可见的部门数据，部门可见性需要管理员在后台配置

部门存在，但用户搜索不到并不一定是搜索有问题，可能是管理员在后台配置了权限控制，导致用户无法搜索到该部门







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/search)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments/search`

### Method

`POST`

## UpdateDepartmentPatch

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.UpdateDepartmentPatch(ctx, &lark.UpdateDepartmentPatchReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新通讯录中部门的信息中的任一个字段。

调用该接口需要具有该部门以及更新操作涉及的部门的通讯录权限。应用商店应用无权限调用此接口。


该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/patch)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments/:department_id`

### Method

`PATCH`

## UpdateDepartment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.UpdateDepartment(ctx, &lark.UpdateDepartmentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新当前部门所有信息。

- 调用该接口需要具有该部门以及更新操作涉及的部门的通讯录权限。应用商店应用无权限调用此接口。

 - 没有填写的字段会被置为空值（order字段除外）。


该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/update)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments/:department_id`

### Method

`PUT`

## DeleteDepartment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.DeleteDepartment(ctx, &lark.DeleteDepartmentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于向通讯录中删除部门。

应用需要同时拥有待删除部门及其父部门的通讯录授权。应用商店应用无权限调用该接口。







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/delete)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments/:department_id`

### Method

`DELETE`

## GetEmployeeTypeEnumList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetEmployeeTypeEnumList(ctx, &lark.GetEmployeeTypeEnumListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取员工的人员类型

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list)

### URL

`https://open.feishu.cn/open-apis/contact/v3/employee_type_enums`

### Method

`GET`

## UpdateEmployeeTypeEnumPatch

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.UpdateEmployeeTypeEnumPatch(ctx, &lark.UpdateEmployeeTypeEnumPatchReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新自定义人员类型

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/update)

### URL

`https://open.feishu.cn/open-apis/contact/v3/employee_type_enums/:enum_id`

### Method

`PUT`

## DeleteEmployeeTypeEnum

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.DeleteEmployeeTypeEnum(ctx, &lark.DeleteEmployeeTypeEnumReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除自定义人员枚举

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/delete)

### URL

`https://open.feishu.cn/open-apis/contact/v3/employee_type_enums/:enum_id`

### Method

`DELETE`

## CreateEmployeeTypeEnum

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.CreateEmployeeTypeEnum(ctx, &lark.CreateEmployeeTypeEnumReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

新增自定义人员类型

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/create)

### URL

`https://open.feishu.cn/open-apis/contact/v3/employee_type_enums`

### Method

`POST`

## GetContactCustomAttrList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetContactCustomAttrList(ctx, &lark.GetContactCustomAttrListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取企业自定义的用户字段配置信息

调用该接口前，需要先确认企业管理员已经在 企业管理后台 > 组织架构 > 成员字段管理 自定义字段管理栏开启了“允许开放平台API调用“。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/custom_attr/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/custom_attr/list)

### URL

`https://open.feishu.cn/open-apis/contact/v3/custom_attrs`

### Method

`GET`


# Message

## SendRawMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.SendRawMessage(ctx, &lark.SendRawMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

给指定用户或者会话发送消息，支持文本、富文本、卡片、群名片、个人名片、图片、视频、音频、文件、表情包。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 给用户发送消息，需要机器人对用户有可见性
- 给群组发送消息，需要机器人在群中
- 消息请求体最大不能超过30k




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages`

### Method

`POST`

## SendRawMessageOld

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.SendRawMessageOld(ctx, &lark.SendRawMessageOldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```





### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUjNz4SN2MjL1YzM](https://open.feishu.cn/document/ukTMukTMukTM/uUjNz4SN2MjL1YzM)

### URL

`https://open.feishu.cn/open-apis/message/v4/send/`

### Method

`POST`

## ReplyRawMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.ReplyRawMessage(ctx, &lark.ReplyRawMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

回复指定消息，支持文本、富文本、卡片、群名片、个人名片、图片、视频、文件等多种消息类型。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 回复私聊消息，需要机器人对用户有可见性
- 回复群组消息，需要机器人在群中


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/reply](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/reply)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id/reply`

### Method

`POST`

## DeleteMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.DeleteMessage(ctx, &lark.DeleteMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

机器人撤回机器人自己发送的消息或群主撤回群内消息。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  ，撤回消息时机器人仍需要在会话内
- 机器人可以撤回单聊和群组内，自己发送 且 发送时间不超过1天(24小时)的消息
- 若机器人要撤回群内他人发送的消息，则机器人必须是该群的群主 或者 得到群主的授权，且消息发送时间不超过1天（24小时）
- 无法撤回通过「批量发送消息接口」发送的消息


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/delete)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id`

### Method

`DELETE`

## UpdateMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.UpdateMessage(ctx, &lark.UpdateMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新应用已发送的消息卡片内容。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 当前仅更新支持卡片消息
- 只支持对所有人都更新的「共享卡片」。如果你只想更新特定人的消息卡片，必须要用户在卡片操作交互后触发，开发文档参考[「独享卡片」](https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN#49904b71)


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/patch)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id`

### Method

`PATCH`

## GetMessageReadUserList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.GetMessageReadUserList(ctx, &lark.GetMessageReadUserListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

查询消息的已读用户信息。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 只能查询机器人自己发送，且发送时间不超过7天的消息
- 查询消息已读信息时机器人仍需要在会话内


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/read_users](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/read_users)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id/read_users`

### Method

`GET`

## GetMessageList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.GetMessageList(ctx, &lark.GetMessageListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取会话（包括单聊、群组）的历史消息。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 获取消息时，机器人必须在群组中
- 获取群组消息时，应用必须拥有 获取群组中所有的消息 权限


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/list)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages`

### Method

`GET`

## GetMessageFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.GetMessageFile(ctx, &lark.GetMessageFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取消息中的资源文件，包括音频，视频，图片和文件。当前仅支持 100M 以内的资源文件的下载。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 机器人和消息需要在同一会话中
- 请求的 file_key 和 message_id 需要匹配
- 获取群组消息时，应用必须拥有 获取群组中所有的消息 权限


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-resource/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-resource/get)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id/resources/:file_key`

### Method

`GET`

## GetMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.GetMessage(ctx, &lark.GetMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

通过 message_id 查询消息内容

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 机器人必须在群组中


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/get)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id`

### Method

`GET`

## DeleteEphemeralMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.DeleteEphemeralMessage(ctx, &lark.DeleteEphemeralMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



在群会话中删除指定用户的临时消息卡片<br>
临时卡片消息可以通过该接口进行显式删除，临时卡片消息删除后将不会在该设备上留下任何痕迹。
**权限说明** ：需要启用机器人能力；需要机器人在会话群里

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITOyYjLykjM24iM5IjN](https://open.feishu.cn/document/ukTMukTMukTM/uITOyYjLykjM24iM5IjN)

### URL

`https://open.feishu.cn/open-apis/ephemeral/v1/delete`

### Method

`POST`


# Chat

## CreateChat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.CreateChat(ctx, &lark.CreateChatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

创建群并设置群头像、群名、群描述等。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/create)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats`

### Method

`POST`

## GetChat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.GetChat(ctx, &lark.GetChatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取群名称、群描述、群头像、群主 ID 等群基本信息。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
 - 机器人或授权用户必须在群里


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id`

### Method

`GET`

## UpdateChat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.UpdateChat(ctx, &lark.UpdateChatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新群头像、群名称、群描述、群配置、转让群主等。

注意事项：
- 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 若群未开启==仅群主可编辑群信息== 配置：
 	- 群主 或 创建群组且具备==更新应用所创建群的群信息==权限的机器人，可更新所有信息
 	- 不满足上述条件的群成员或机器人，仅可更新群头像、群名称、群描述、群国际化名称信息 
- 若群开启了==仅群主可编辑群信息==配置：
 	- 群主 或 创建群组且具备==更新应用所创建群的群信息==权限的机器人，可更新所有信息
 	- 不满足上述条件的群成员或者机器人，任何群信息都不能修改


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/update)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id`

### Method

`PUT`

## DeleteChat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.DeleteChat(ctx, &lark.DeleteChatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

解散群组

注意事项：
- 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 仅有 群主 或 创建群组且具备==更新应用所创建群的群信息==权限的机器人，可解散群


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/delete)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id`

### Method

`DELETE`

## GetChatListOfSelf

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.GetChatListOfSelf(ctx, &lark.GetChatListOfSelfReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取用户或者机器人所在群列表。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/list)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats`

### Method

`GET`

## SearchChat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.SearchChat(ctx, &lark.SearchChatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

搜索对用户或机器人可见的群列表，包括：用户或机器人所在的群、对用户或机器人公开的群。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/search)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/search`

### Method

`GET`

## GetChatMemberList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.GetChatMemberList(ctx, &lark.GetChatMemberListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

如果用户在群中，则返回该群的成员列表。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
 - 该接口不会返回群内的机器人成员


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/get)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members`

### Method

`GET`

## IsInChat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.IsInChat(ctx, &lark.IsInChatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

判断用户或者机器人是否在群里。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/is_in_chat](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/is_in_chat)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members/is_in_chat`

### Method

`GET`

## AddChatMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.AddChatMember(ctx, &lark.AddChatMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

将用户或机器人拉入群聊。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
 - 如需拉用户进群，需要机器人对用户有可见性
 - 在开启==仅群主可添加群成员==的设置时，仅有群主 或 创建群组且具备==更新应用所创建群的群信息==权限的机器人，可以拉用户或者机器人进群
 - 在未开启==仅群主可添加群成员==的设置时，所有群成员都可以拉用户或机器人进群
 - 每次请求，最多拉50个用户或者5个机器人，并且群组最多容纳15个机器人


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members`

### Method

`POST`

## DeleteChatMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.DeleteChatMember(ctx, &lark.DeleteChatMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

将用户或机器人移出群聊。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 用户或机器人在任何条件下均可移除自己出群（即主动退群）；但仅有群主 及 创建群组且具备==更新应用所创建群的群信息==权限的机器人，可以移除其他用户或者机器人
 - 每次请求，最多移除50个用户或者5个机器人


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/delete)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members`

### Method

`DELETE`

## JoinChat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.JoinChat(ctx, &lark.JoinChatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

用户或机器人主动加入群聊。

注意事项：
- 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
 - 目前仅支持加入公开群


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/me_join](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/me_join)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members/me_join`

### Method

`PATCH`

## GetChatAnnouncement

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.GetChatAnnouncement(ctx, &lark.GetChatAnnouncementReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取会话中的群公告信息，公告信息格式与[云文档](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)格式相同。

注意事项：
- 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/get)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/announcement`

### Method

`GET`

## UpdateChatAnnouncement

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.UpdateChatAnnouncement(ctx, &lark.UpdateChatAnnouncementReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新会话中的群公告信息，更新公告信息的格式和更新[云文档](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)格式相同。

注意事项：
- 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 当授权用户或机器人是群主时，可更新群公告信息
- 当授权用户或机器人非群主，但群主未设置 ==仅群主可编辑群信息== 时，可更新群公告信息
- 当授权用户或机器人非群主，但群主设置了 ==仅群主可编辑群信息== 时，无法更新公告信息


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/patch)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/announcement`

### Method

`PATCH`


# Bot

## AddBotToChat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bot.AddBotToChat(ctx, &lark.AddBotToChatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



拉机器人进群<br>

**权限说明** ：需要启用机器人能力；机器人的owner需要已经在群里

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYDO04iN4QjL2gDN](https://open.feishu.cn/document/ukTMukTMukTM/uYDO04iN4QjL2gDN)

### URL

`https://open.feishu.cn/open-apis/bot/v4/add`

### Method

`POST`


# Calendar

## CreateCalendarACL

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.CreateCalendarACL(ctx, &lark.CreateCalendarACLReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）给日历添加访问控制权限，即日历成员。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.acl&method=create)







当前身份需要有日历的 owner 权限，并且日历的类型只能为 primary 或 shared。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/create)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls`

### Method

`POST`

## DeleteCalendarACL

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.DeleteCalendarACL(ctx, &lark.DeleteCalendarACLReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）删除日历的控制权限，即日历成员。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.acl&method=delete)







当前身份需要有日历的 owner 权限，并且日历的类型只能为 primary 或 shared。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/delete)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id`

### Method

`DELETE`

## GetCalendarACLList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GetCalendarACLList(ctx, &lark.GetCalendarACLListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）获取日历的控制权限列表。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.acl&method=list)







当前身份需要有日历的 owner 权限，并且日历的类型只能为 primary 或 shared。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/list)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls`

### Method

`GET`

## SubscribeCalendarACL

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.SubscribeCalendarACL(ctx, &lark.SubscribeCalendarACLReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于以用户身份订阅指定日历下的日历成员变更事件。

用户必须对日历有访问权限。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/subscription](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/subscription)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls/subscription`

### Method

`POST`

## CreateCalendar

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.CreateCalendar(ctx, &lark.CreateCalendarReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于为当前身份（应用 / 用户）创建一个共享日历。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&method=create)










#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/create)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars`

### Method

`POST`

## DeleteCalendar

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.DeleteCalendar(ctx, &lark.DeleteCalendarReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）删除一个共享日历。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&method=delete)







当前身份必须对日历具有 owner 权限。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/delete)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id`

### Method

`DELETE`

## GetCalendar

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GetCalendar(ctx, &lark.GetCalendarReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）根据日历 ID 获取日历信息。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&method=get)







当前身份必须对日历有访问权限。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id`

### Method

`GET`

## GetCalendarList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GetCalendarList(ctx, &lark.GetCalendarListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于分页获得当前身份（应用 / 用户）的日历列表。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&method=list)







调用时首先使用 page_token 分页拉取存量数据，之后使用 sync_token 增量同步变更数据。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/list)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars`

### Method

`GET`

## UpdateCalendar

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.UpdateCalendar(ctx, &lark.UpdateCalendarReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）修改日历信息。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&method=patch)







当前身份对日历有 owner 权限时，可修改全局字段：summary, description, permission。

当前身份对日历不具有 owner 权限时，仅可修改对自己生效的字段：color, summary_alias。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/patch)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id`

### Method

`PATCH`

## SearchCalendar

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.SearchCalendar(ctx, &lark.SearchCalendarReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于通过关键字查询公共日历或用户主日历。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/search)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/search`

### Method

`POST`

## SubscribeCalendar

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.SubscribeCalendar(ctx, &lark.SubscribeCalendarReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）订阅某个日历。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&method=subscribe)







仅可订阅类型为 primary 或 shared 的公开日历。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/subscribe](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/subscribe)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/subscribe`

### Method

`POST`

## UnsubscribeCalendar

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.UnsubscribeCalendar(ctx, &lark.UnsubscribeCalendarReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）取消对某日历的订阅状态。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&method=unsubscribe)







仅可操作已经被当前身份订阅的日历。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/unsubscribe](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/unsubscribe)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe`

### Method

`POST`

## SubscribeCalendarChangeEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.SubscribeCalendarChangeEvent(ctx, &lark.SubscribeCalendarChangeEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于以用户身份订阅当前身份下日历列表中的所有日历变更。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/subscription](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/subscription)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/subscription`

### Method

`POST`

## CreateCalendarEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.CreateCalendarEvent(ctx, &lark.CreateCalendarEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）在日历上创建一个日程。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.event&method=create)







当前身份必须对日历有 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/create)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events`

### Method

`POST`

## DeleteCalendarEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.DeleteCalendarEvent(ctx, &lark.DeleteCalendarEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）删除日历上的一个日程。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.event&method=delete)







当前身份必须对日历有 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。

当前身份必须是日程的组织者。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/delete)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id`

### Method

`DELETE`

## GetCalendarEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GetCalendarEvent(ctx, &lark.GetCalendarEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于以当前身份（应用 / 用户）获取日历上的一个日程。

当前身份必须对日历有访问权限。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/get)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id`

### Method

`GET`

## GetCalendarEventList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GetCalendarEventList(ctx, &lark.GetCalendarEventListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于以当前身份（应用 / 用户）获取日历下的日程列表。

当前身份必须对日历有访问权限。

调用时首先使用 page_token 分页拉取存量数据，之后使用 sync_token 增量同步变更数据。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/list)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events`

### Method

`GET`

## UpdateCalendarEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.UpdateCalendarEvent(ctx, &lark.UpdateCalendarEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以当前身份（应用 / 用户）更新日历上的一个日程。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.event&method=patch)







当前身份必须对日历有 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。

当前身份为日程组织者时，可修改所有可编辑字段。

当前身份为日程参与者时，仅可编辑部分字段。（如：visibility, free_busy_status, color, reminders）




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/patch)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id`

### Method

`PATCH`

## SearchCalendarEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.SearchCalendarEvent(ctx, &lark.SearchCalendarEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于以用户身份搜索某日历下的相关日程。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.event&method=search)







当前身份必须对日历有访问权限。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/search)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/search`

### Method

`POST`

## SubscribeCalendarEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.SubscribeCalendarEvent(ctx, &lark.SubscribeCalendarEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于以用户身份订阅指定日历下的日程变更事件。

用户必须对日历有访问权限。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/subscription](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/subscription)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/subscription`

### Method

`POST`

## CreateCalendarEventAttendee

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.CreateCalendarEventAttendee(ctx, &lark.CreateCalendarEventAttendeeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

批量给日程添加参与人。

- 当前身份需要有日历的 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。

- 当前身份需要是日程的组织者，或日程设置了「参与人可邀请其它参与人」权限。

- 新添加的日程参与人必须与日程组织者在同一个企业内。

- 使用该接口添加会议室后，会议室会进入异步的预约流程，请求结束不代表会议室预约成功，需后续再查询预约状态。

- 每个日程最多只能有 3000 名参与人。







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/create)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees`

### Method

`POST`

## GetCalendarEventAttendeeList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GetCalendarEventAttendeeList(ctx, &lark.GetCalendarEventAttendeeListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取日程的参与人列表，若参与者列表中有群组，请使用 [获取参与人群成员列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee-chat_member/list) 。

- 当前身份必须对日历有访问权限。

- 当前身份必须有权限查看日程的参与人列表。





#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/list)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees`

### Method

`GET`

## DeleteCalendarEventAttendee

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.DeleteCalendarEventAttendee(ctx, &lark.DeleteCalendarEventAttendeeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

批量删除日程的参与人。

- 当前身份需要有日历的 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。

- 当前身份需要是日程的组织者。







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/batch_delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/batch_delete)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/batch_delete`

### Method

`POST`

## GetCalendarEventAttendeeChatMemberList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GetCalendarEventAttendeeChatMemberList(ctx, &lark.GetCalendarEventAttendeeChatMemberListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取日程的群参与人的群成员列表。

- 当前身份必须有权限查看日程的参与人列表。

- 当前身份必须在群聊中，或有权限查看群成员列表。





#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee-chat_member/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee-chat_member/list)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/:attendee_id/chat_members`

### Method

`GET`

## GetCalendarFreeBusyList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GetCalendarFreeBusyList(ctx, &lark.GetCalendarFreeBusyListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

查询用户主日历或会议室的忙闲信息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/freebusy/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/freebusy/list)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/freebusy/list`

### Method

`POST`

## CreateCalendarTimeoffEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.CreateCalendarTimeoffEvent(ctx, &lark.CreateCalendarTimeoffEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

为指定用户创建一个请假日程，可以是一个普通请假日程，也可以是一个全天日程。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/timeoff_event/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/timeoff_event/create)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/timeoff_events`

### Method

`POST`

## DeleteCalendarTimeoffEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.DeleteCalendarTimeoffEvent(ctx, &lark.DeleteCalendarTimeoffEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除一个指定的请假日程，请假日程删除，用户个人签名页的请假信息也会消失。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/timeoff_event/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/timeoff_event/delete)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/timeoff_events/:timeoff_event_id`

### Method

`DELETE`

## GenerateCaldavConf

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GenerateCaldavConf(ctx, &lark.GenerateCaldavConfReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

用于为当前用户生成一个CalDAV账号密码，用于将飞书日历信息同步到本地设备日历。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/setting/generate_caldav_conf](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/setting/generate_caldav_conf)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/settings/generate_caldav_conf`

### Method

`POST`


# Drive

## SearchDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.SearchDriveFile(ctx, &lark.SearchDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据搜索条件进行文档搜索。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ugDM4UjL4ADO14COwgTN](https://open.feishu.cn/document/ukTMukTMukTM/ugDM4UjL4ADO14COwgTN)

### URL

`https://open.feishu.cn/open-apis/suite/docs-api/search/object`

### Method

`POST`

## GetDriveFileMeta

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveFileMeta(ctx, &lark.GetDriveFileMetaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 token 获取各类文件的元数据。

请求用户需要拥有该文件的访问（读）权限

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMjN3UjLzYzN14yM2cTN](https://open.feishu.cn/document/ukTMukTMukTM/uMjN3UjLzYzN14yM2cTN)

### URL

`https://open.feishu.cn/open-apis/suite/docs-api/meta`

### Method

`POST`

## CreateDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateDriveFile(ctx, &lark.CreateDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于根据 folderToken 创建 Doc或 Sheet 。

若没有特定的文件夹用于承载创建的文档，可以先调用「获取文件夹元信息」文档中的「获取 root folder (我的空间) meta」接口，获得我的空间的 token，然后再使用此接口。创建的文档将会在「我的空间」的「归我所有」列表里。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=drive_explorer&version=v2&resource=file&method=create)


该接口不支持并发创建，且调用频率上限为 5QPS 且 10000次/天

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQTNzUjL0UzM14CN1MTN](https://open.feishu.cn/document/ukTMukTMukTM/uQTNzUjL0UzM14CN1MTN)

### URL

`https://open.feishu.cn/open-apis/drive/explorer/v2/file/:folderToken`

### Method

`POST`

## CopyDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CopyDriveFile(ctx, &lark.CopyDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据文件 token 复制 Doc 或 Sheet  到目标文件夹中。

若没有特定的文件夹用于承载创建的文档，可以先调用「获取文件夹元信息」文档中的「获取 root folder (我的空间) meta」接口，获得我的空间的 token，然后再使用此接口。复制的文档将会在「我的空间」的「归我所有」列表里。


该接口不支持并发创建，且调用频率上限为 5QPS 且 10000次/天

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYTNzUjL2UzM14iN1MTN](https://open.feishu.cn/document/ukTMukTMukTM/uYTNzUjL2UzM14iN1MTN)

### URL

`https://open.feishu.cn/open-apis/drive/explorer/v2/file/copy/files/:fileToken`

### Method

`POST`

## DeleteDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteDriveFile(ctx, &lark.DeleteDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 docToken 删除对应的 Docs 文档。

:::html

<md-alert type="warn">

文档只能被文档所有者删除，文档被删除后将会放到回收站里
</md-alert>

:::
该接口不支持并发调用，且调用频率上限为5QPS

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATM2UjLwEjN14CMxYTN](https://open.feishu.cn/document/ukTMukTMukTM/uATM2UjLwEjN14CMxYTN)

### URL

`https://open.feishu.cn/open-apis/drive/explorer/v2/file/docs/:docToken`

### Method

`DELETE`

## DeleteDriveSheetFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteDriveSheetFile(ctx, &lark.DeleteDriveSheetFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 删除对应的 sheet 文档。

:::html

<md-alert type="warn">

文档只能被文档所有者删除，文档被删除后将会放到回收站里
</md-alert>

:::

该接口不支持并发调用，且调用频率上限为5QPS

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUTNzUjL1UzM14SN1MTN/delete-sheet](https://open.feishu.cn/document/ukTMukTMukTM/uUTNzUjL1UzM14SN1MTN/delete-sheet)

### URL

`https://open.feishu.cn/open-apis/drive/explorer/v2/file/spreadsheets/:spreadsheetToken`

### Method

`DELETE`

## CreateDriveFolder

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateDriveFolder(ctx, &lark.CreateDriveFolderReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 folderToken 在该 folder 下创建文件夹。

该接口不支持并发创建，且调用频率上限为 5QPS 以及 10000次/天


#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukTNzUjL5UzM14SO1MTN](https://open.feishu.cn/document/ukTMukTMukTM/ukTNzUjL5UzM14SO1MTN)

### URL

`https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken`

### Method

`POST`

## GetDriveFolderMeta

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveFolderMeta(ctx, &lark.GetDriveFolderMetaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 folderToken 获取该文件夹的元信息。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uAjNzUjLwYzM14CM2MTN](https://open.feishu.cn/document/ukTMukTMukTM/uAjNzUjLwYzM14CM2MTN)

### URL

`https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/meta`

### Method

`GET`

## GetDriveRootFolderMeta

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveRootFolderMeta(ctx, &lark.GetDriveRootFolderMetaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于获取 "我的空间" 的元信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ugTNzUjL4UzM14CO1MTN/get-root-folder-meta](https://open.feishu.cn/document/ukTMukTMukTM/ugTNzUjL4UzM14CO1MTN/get-root-folder-meta)

### URL

`https://open.feishu.cn/open-apis/drive/explorer/v2/root_folder/meta`

### Method

`GET`

## GetDriveFolderChildren

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveFolderChildren(ctx, &lark.GetDriveFolderChildrenReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 folderToken 获取该文件夹的文档清单，如 doc、sheet、folder。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uEjNzUjLxYzM14SM2MTN](https://open.feishu.cn/document/ukTMukTMukTM/uEjNzUjLxYzM14SM2MTN)

### URL

`https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/children`

### Method

`GET`

## GetDriveFileStatistics

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveFileStatistics(ctx, &lark.GetDriveFileStatisticsReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

此接口用于获取文件统计信息，包括文档阅读人数、次数和点赞数。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-statistics/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-statistics/get)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/statistics`

### Method

`GET`

## DownloadDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DownloadDriveFile(ctx, &lark.DownloadDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

使用该接口可以下载在云空间目录下的文件（不含飞书文档/表格/思维导图等在线文档）。支持range下载。

该接口支持调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/download](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/download)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/download`

### Method

`GET`

## UploadDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UploadDriveFile(ctx, &lark.UploadDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

向云空间指定目录下上传一个文件

使用此方式上传可以快速传输小于等于20MB的文件


该接口支持调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_all](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_all)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/upload_all`

### Method

`POST`

## PrepareUploadDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.PrepareUploadDriveFile(ctx, &lark.PrepareUploadDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

发送初始化请求获取上传事务ID和分块策略，目前是以4MB大小进行定长分片。

您在24小时内可保存上传事务ID和上传进度，以便可以恢复上传


该接口不支持太高的并发，且调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_prepare](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_prepare)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/upload_prepare`

### Method

`POST`

## PartUploadDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.PartUploadDriveFile(ctx, &lark.PartUploadDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

上传对应的文件块。

该接口不支持太高的并发，且调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_part](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_part)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/upload_part`

### Method

`POST`

## FinishUploadDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.FinishUploadDriveFile(ctx, &lark.FinishUploadDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

触发完成上传。

该接口不支持太高的并发，且调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_finish](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_finish)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/upload_finish`

### Method

`POST`

## DownloadDriveMedia

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DownloadDriveMedia(ctx, &lark.DownloadDriveMediaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

使用该接口可以下载素材。素材表示在各种创作容器里的文件，如Doc文档内的图片，文件均属于素材。支持range下载。

该接口不支持太高的并发，且调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/download](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/download)

### URL

`https://open.feishu.cn/open-apis/drive/v1/medias/:file_token/download`

### Method

`GET`

## UploadDriveMedia

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UploadDriveMedia(ctx, &lark.UploadDriveMediaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

将云文档或其他飞书套件所需要的素材（图片/文件）上传到云空间。

该接口不支持太高的并发，且调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_all](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_all)

### URL

`https://open.feishu.cn/open-apis/drive/v1/medias/upload_all`

### Method

`POST`

## PrepareUploadDriveMedia

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.PrepareUploadDriveMedia(ctx, &lark.PrepareUploadDriveMediaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

发送初始化请求获取上传事务ID和分块策略，目前是以4MB大小进行定长分片。

您在24小时内可保存上传事务ID和上传进度，以便可以恢复上传


该接口不支持太高的并发，且调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_prepare](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_prepare)

### URL

`https://open.feishu.cn/open-apis/drive/v1/medias/upload_prepare`

### Method

`POST`

## PartUploadDriveMedia

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.PartUploadDriveMedia(ctx, &lark.PartUploadDriveMediaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

上传对应的文件块。

该接口不支持太高的并发，且调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_part](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_part)

### URL

`https://open.feishu.cn/open-apis/drive/v1/medias/upload_part`

### Method

`POST`

## FinishUploadDriveMedia

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.FinishUploadDriveMedia(ctx, &lark.FinishUploadDriveMediaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

触发完成上传。

该接口不支持太高的并发，且调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_finish](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_finish)

### URL

`https://open.feishu.cn/open-apis/drive/v1/medias/upload_finish`

### Method

`POST`

## CreateDriveMemberPermissionOld

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateDriveMemberPermissionOld(ctx, &lark.CreateDriveMemberPermissionOldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 filetoken 给用户增加文档的权限。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzNzUjLzczM14yM3MTN](https://open.feishu.cn/document/ukTMukTMukTM/uMzNzUjLzczM14yM3MTN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/member/create`

### Method

`POST`

## TransferDriveMemberPermission

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.TransferDriveMemberPermission(ctx, &lark.TransferDriveMemberPermissionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据文档信息和用户信息转移文档的所有者。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQzNzUjL0czM14CN3MTN](https://open.feishu.cn/document/ukTMukTMukTM/uQzNzUjL0czM14CN3MTN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/member/transfer`

### Method

`POST`

## GetDriveMemberPermissionList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveMemberPermissionList(ctx, &lark.GetDriveMemberPermissionListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 filetoken 查询协作者，目前包括人("user")和群("chat") 。

你能获取到协作者列表的前提是你对该文档有分享权限


#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATN3UjLwUzN14CM1cTN](https://open.feishu.cn/document/ukTMukTMukTM/uATN3UjLwUzN14CM1cTN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/member/list`

### Method

`POST`

## CreateDriveMemberPermission

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateDriveMemberPermission(ctx, &lark.CreateDriveMemberPermissionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 filetoken 给用户增加文档的权限。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/create)

### URL

`https://open.feishu.cn/open-apis/drive/v1/permissions/:token/members`

### Method

`POST`

## DeleteDriveMemberPermission

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteDriveMemberPermission(ctx, &lark.DeleteDriveMemberPermissionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 filetoken 移除文档协作者的权限。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYTN3UjL2UzN14iN1cTN](https://open.feishu.cn/document/ukTMukTMukTM/uYTN3UjL2UzN14iN1cTN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/member/delete`

### Method

`POST`

## UpdateDriveMemberPermission

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateDriveMemberPermission(ctx, &lark.UpdateDriveMemberPermissionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 filetoken 更新文档协作者的权限。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ucTN3UjL3UzN14yN1cTN](https://open.feishu.cn/document/ukTMukTMukTM/ucTN3UjL3UzN14yN1cTN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/member/update`

### Method

`POST`

## CheckDriveMemberPermission

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CheckDriveMemberPermission(ctx, &lark.CheckDriveMemberPermissionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 filetoken 判断当前登录用户是否具有某权限。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYzN3UjL2czN14iN3cTN](https://open.feishu.cn/document/ukTMukTMukTM/uYzN3UjL2czN14iN3cTN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/member/permitted`

### Method

`POST`

## UpdateDrivePublicPermissionV1Old

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateDrivePublicPermissionV1Old(ctx, &lark.UpdateDrivePublicPermissionV1OldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 filetoken 更新文档的公共设置。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukTM3UjL5EzN14SOxcTN](https://open.feishu.cn/document/ukTMukTMukTM/ukTM3UjL5EzN14SOxcTN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/public/update`

### Method

`POST`

## UpdateDrivePublicPermissionV2Old

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateDrivePublicPermissionV2Old(ctx, &lark.UpdateDrivePublicPermissionV2OldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 filetoken 更新文档的公共设置。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITN5UjLyUTO14iM1kTN](https://open.feishu.cn/document/ukTMukTMukTM/uITN5UjLyUTO14iM1kTN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/v2/public/update/`

### Method

`POST`

## GetDrivePublicPermissionV2

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDrivePublicPermissionV2(ctx, &lark.GetDrivePublicPermissionV2Req{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 filetoken 获取文档的公共设置。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITM3YjLyEzN24iMxcjN](https://open.feishu.cn/document/ukTMukTMukTM/uITM3YjLyEzN24iMxcjN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/v2/public/`

### Method

`POST`

## UpdateDrivePublicPermission

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateDrivePublicPermission(ctx, &lark.UpdateDrivePublicPermissionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 filetoken 更新文档的公共设置。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-public/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-public/patch)

### URL

`https://open.feishu.cn/open-apis/drive/v1/permissions/:token/public`

### Method

`PATCH`

## BatchGetDriveMediaTmpDownloadURL

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.BatchGetDriveMediaTmpDownloadURL(ctx, &lark.BatchGetDriveMediaTmpDownloadURLReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

通过file_token获取素材临时下载链接，链接时效性是24小时，过期失效。

该接口不支持太高的并发，且调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/batch_get_tmp_download_url](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/batch_get_tmp_download_url)

### URL

`https://open.feishu.cn/open-apis/drive/v1/medias/batch_get_tmp_download_url`

### Method

`GET`

## GetDriveCommentList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveCommentList(ctx, &lark.GetDriveCommentListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

通过分页方式获取云文档中的全文评论列表。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/list)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments`

### Method

`GET`

## GetDriveComment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveComment(ctx, &lark.GetDriveCommentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取云文档中的某条评论。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/get)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments/:comment_id`

### Method

`GET`

## CreateDriveComment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateDriveComment(ctx, &lark.CreateDriveCommentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

往云文档添加一条评论。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/create)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments`

### Method

`POST`

## UpdateDriveComment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateDriveComment(ctx, &lark.UpdateDriveCommentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新云文档中的某条回复。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment-reply/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment-reply/update)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies/:reply_id`

### Method

`PUT`

## DeleteDriveComment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteDriveComment(ctx, &lark.DeleteDriveCommentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除云文档中的某条回复。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment-reply/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment-reply/delete)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies/:reply_id`

### Method

`DELETE`

## UpdateDriveCommentPatch

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateDriveCommentPatch(ctx, &lark.UpdateDriveCommentPatchReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

解决或恢复云文档中的评论。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/patch)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments/:comment_id`

### Method

`PATCH`

## CreateDriveDoc

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateDriveDoc(ctx, &lark.CreateDriveDocReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



在使用此接口前，请仔细阅读[概述](/ssl:ttdoc/ukTMukTMukTM/ukjM5YjL5ITO24SOykjN)和[准备接入文档 API](/ssl:ttdoc/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)了解文档调用的规则和约束，确保你的文档数据不会丢失或出错。  
文档数据结构定义可参考：[文档数据结构概述](/ssl:ttdoc/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN) 

该接口用于创建并初始化文档。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=doc&version=v2&resource=doc&method=create)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ugDM2YjL4AjN24COwYjN](https://open.feishu.cn/document/ukTMukTMukTM/ugDM2YjL4AjN24COwYjN)

### URL

`https://open.feishu.cn/open-apis/doc/v2/create`

### Method

`POST`

## GetDriveDocContent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveDocContent(ctx, &lark.GetDriveDocContentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



在使用此接口前，请仔细阅读[概述](/ssl:ttdoc/ukTMukTMukTM/ukjM5YjL5ITO24SOykjN)和[准备接入文档 API](//ssl:ttdoc/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)了解文档调用的规则和约束，确保你的文档数据不会丢失或出错。  
文档数据结构定义可参考：[文档数据结构概述](/ssl:ttdoc/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN) 
该接口用于获取结构化的文档内容。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=doc&version=v2&resource=doc&method=content)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDM2YjL1AjN24SNwYjN](https://open.feishu.cn/document/ukTMukTMukTM/uUDM2YjL1AjN24SNwYjN)

### URL

`https://open.feishu.cn/open-apis/doc/v2/:docToken/content`

### Method

`GET`

## GetDriveDocRawContent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveDocRawContent(ctx, &lark.GetDriveDocRawContentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取文档的纯文本内容，不包含富文本格式信息。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukzNzUjL5czM14SO3MTN](https://open.feishu.cn/document/ukTMukTMukTM/ukzNzUjL5czM14SO3MTN)

### URL

`https://open.feishu.cn/open-apis/doc/v2/:docToken/raw_content`

### Method

`GET`

## GetDriveDocMeta

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveDocMeta(ctx, &lark.GetDriveDocMetaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 docToken 获取元数据。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uczN3UjL3czN14yN3cTN](https://open.feishu.cn/document/ukTMukTMukTM/uczN3UjL3czN14yN3cTN)

### URL

`https://open.feishu.cn/open-apis/doc/v2/meta/:docToken`

### Method

`GET`

## CreateSheet

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateSheet(ctx, &lark.CreateSheetReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

使用该接口可以在指定的目录下创建在线表格。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet/create](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet/create)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets`

### Method

`POST`

## GetSheetMeta

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetSheetMeta(ctx, &lark.GetSheetMetaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 获取表格元数据。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uETMzUjLxEzM14SMxMTN](https://open.feishu.cn/document/ukTMukTMukTM/uETMzUjLxEzM14SMxMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/metainfo`

### Method

`GET`

## UpdateSheetProperty

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateSheetProperty(ctx, &lark.UpdateSheetPropertyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 更新表格属性，如更新表格标题。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ucTMzUjL3EzM14yNxMTN](https://open.feishu.cn/document/ukTMukTMukTM/ucTMzUjL3EzM14yNxMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/properties`

### Method

`PUT`

## BatchUpdateSheet

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.BatchUpdateSheet(ctx, &lark.BatchUpdateSheetReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 操作表格，如增加工作表，复制工作表、删除工作表。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYTMzUjL2EzM14iNxMTN](https://open.feishu.cn/document/ukTMukTMukTM/uYTMzUjL2EzM14iNxMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/sheets_batch_update`

### Method

`POST`

## ImportSheet

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.ImportSheet(ctx, &lark.ImportSheetReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于将本地表格导入到云空间上。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATO2YjLwkjN24CM5YjN](https://open.feishu.cn/document/ukTMukTMukTM/uATO2YjLwkjN24CM5YjN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/import`

### Method

`POST`

## CreateDriveImportTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateDriveImportTask(ctx, &lark.CreateDriveImportTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

创建导入任务。支持导入为 doc、sheet、bitable，参考[导入用户指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/import-user-guide)

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/create)

### URL

`https://open.feishu.cn/open-apis/drive/v1/import_tasks`

### Method

`POST`

## GetDriveImportTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveImportTask(ctx, &lark.GetDriveImportTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据创建导入任务返回的 ticket 查询导入结果。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/get)

### URL

`https://open.feishu.cn/open-apis/drive/v1/import_tasks/:ticket`

### Method

`GET`

## MoveSheetDimension

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.MoveSheetDimension(ctx, &lark.MoveSheetDimensionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于移动行列，行列被移动到目标位置后，原本在目标位置的行列会对应右移或下移。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/move_dimension](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/move_dimension)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/move_dimension`

### Method

`POST`

## PrependSheetValue

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.PrependSheetValue(ctx, &lark.PrependSheetValueReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和 range 向范围之前增加相应数据的行和相应的数据，相当于数组的插入操作；单次写入不超过5000行，100列，每个格子大小为0.5M。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uIjMzUjLyIzM14iMyMTN](https://open.feishu.cn/document/ukTMukTMukTM/uIjMzUjLyIzM14iMyMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_prepend`

### Method

`POST`

## AppendSheetValue

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.AppendSheetValue(ctx, &lark.AppendSheetValueReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和 range 遇到空行则进行覆盖追加或新增行追加数据。 空行：默认该行第一个格子是空，则认为是空行；单次写入不超过5000行，100列，每个格子大小为0.5M。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMjMzUjLzIzM14yMyMTN](https://open.feishu.cn/document/ukTMukTMukTM/uMjMzUjLzIzM14yMyMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_append`

### Method

`POST`

## InsertSheetDimensionRange

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.InsertSheetDimensionRange(ctx, &lark.InsertSheetDimensionRangeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和维度信息 插入空行/列。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQjMzUjL0IzM14CNyMTN](https://open.feishu.cn/document/ukTMukTMukTM/uQjMzUjL0IzM14CNyMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/insert_dimension_range`

### Method

`POST`

## AddSheetDimensionRange

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.AddSheetDimensionRange(ctx, &lark.AddSheetDimensionRangeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和长度，在末尾增加空行/列；单次操作不超过5000行或列。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUjMzUjL1IzM14SNyMTN](https://open.feishu.cn/document/ukTMukTMukTM/uUjMzUjL1IzM14SNyMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range`

### Method

`POST`

## UpdateSheetDimensionRange

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateSheetDimensionRange(ctx, &lark.UpdateSheetDimensionRangeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和维度信息更新隐藏行列、单元格大小；单次操作不超过5000行或列。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYjMzUjL2IzM14iNyMTN](https://open.feishu.cn/document/ukTMukTMukTM/uYjMzUjL2IzM14iNyMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range`

### Method

`PUT`

## DeleteSheetDimensionRange

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteSheetDimensionRange(ctx, &lark.DeleteSheetDimensionRangeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和维度信息删除行/列 。单次删除最大5000行/列。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ucjMzUjL3IzM14yNyMTN](https://open.feishu.cn/document/ukTMukTMukTM/ucjMzUjL3IzM14yNyMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range`

### Method

`DELETE`

## GetSheetValue

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetSheetValue(ctx, &lark.GetSheetValueReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于根据 spreadsheetToken 和 range 读取表格单个范围的值，返回数据限制为10M。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ugTMzUjL4EzM14COxMTN](https://open.feishu.cn/document/ukTMukTMukTM/ugTMzUjL4EzM14COxMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values/:range`

### Method

`GET`

## BatchGetSheetValue

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.BatchGetSheetValue(ctx, &lark.BatchGetSheetValueReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于根据 spreadsheetToken 和 ranges 读取表格多个范围的值，返回数据限制为10M。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukTMzUjL5EzM14SOxMTN](https://open.feishu.cn/document/ukTMukTMukTM/ukTMzUjL5EzM14SOxMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_batch_get`

### Method

`GET`

## SetSheetValue

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.SetSheetValue(ctx, &lark.SetSheetValueReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和 range 向单个范围写入数据，若范围内有数据，将被更新覆盖；单次写入不超过5000行，100列，每个格子大小为0.5M。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uAjMzUjLwIzM14CMyMTN](https://open.feishu.cn/document/ukTMukTMukTM/uAjMzUjLwIzM14CMyMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values`

### Method

`PUT`

## BatchSetSheetValue

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.BatchSetSheetValue(ctx, &lark.BatchSetSheetValueReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和 range 向多个范围写入数据，若范围内有数据，将被更新覆盖；单次写入不超过5000行，100列，每个格子大小为0.5M。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uEjMzUjLxIzM14SMyMTN](https://open.feishu.cn/document/ukTMukTMukTM/uEjMzUjLxIzM14SMyMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_batch_update`

### Method

`POST`

## SetSheetStyle

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.SetSheetStyle(ctx, &lark.SetSheetStyleReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 、range 和样式信息更新单元格样式；单次写入不超过5000行，100列。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukjMzUjL5IzM14SOyMTN](https://open.feishu.cn/document/ukTMukTMukTM/ukjMzUjL5IzM14SOyMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/style`

### Method

`PUT`

## BatchSetSheetStyle

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.BatchSetSheetStyle(ctx, &lark.BatchSetSheetStyleReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于根据 spreadsheetToken 、range和样式信息 批量更新单元格样式；单次写入不超过5000行，100列。
#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uAzMzUjLwMzM14CMzMTN](https://open.feishu.cn/document/ukTMukTMukTM/uAzMzUjLwMzM14CMzMTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/styles_batch_update`

### Method

`PUT`

## MergeSheetCell

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.MergeSheetCell(ctx, &lark.MergeSheetCellReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和维度信息合并单元格；单次操作不超过5000行，100列。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukDNzUjL5QzM14SO0MTN](https://open.feishu.cn/document/ukTMukTMukTM/ukDNzUjL5QzM14SO0MTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/merge_cells`

### Method

`POST`

## UnmergeSheetCell

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UnmergeSheetCell(ctx, &lark.UnmergeSheetCellReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和维度信息拆分单元格；单次操作不超过5000行，100列。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATNzUjLwUzM14CM1MTN](https://open.feishu.cn/document/ukTMukTMukTM/uATNzUjLwUzM14CM1MTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/unmerge_cells`

### Method

`POST`

## SetSheetValueImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.SetSheetValueImage(ctx, &lark.SetSheetValueImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和 range 向单个格子写入图片。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDNxYjL1QTM24SN0EjN](https://open.feishu.cn/document/ukTMukTMukTM/uUDNxYjL1QTM24SN0EjN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_image`

### Method

`POST`

## FindSheet

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.FindSheet(ctx, &lark.FindSheetReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

按照指定的条件查找子表的某个范围内的数据符合条件的单元格位置。请求体中的 range 和 find 字段为必填。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/find](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/find)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/find`

### Method

`POST`

## ReplaceSheet

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.ReplaceSheet(ctx, &lark.ReplaceSheetReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

按照指定的条件查找子表的某个范围内的数据符合条件的单元格并替换值，返回替换成功的单元格位置。一次请求最多允许替换5000个单元格，如果超过请将range缩小范围再操作。请求体中的 range、find、replaccement 字段必填。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/replace](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/replace)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/replace`

### Method

`POST`

## CreateSheetConditionFormat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateSheetConditionFormat(ctx, &lark.CreateSheetConditionFormatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于创建新的条件格式，单次最多支持增加10个条件格式，每个条件格式的设置会返回成功或者失败，失败的情况包括各种参数的校验。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-set](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-set)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_create`

### Method

`POST`

## GetSheetConditionFormat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetSheetConditionFormat(ctx, &lark.GetSheetConditionFormatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据sheetId查询详细的条件格式信息，最多支持同时查询10个sheetId。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-get](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-get)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats`

### Method

`GET`

## UpdateSheetConditionFormat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateSheetConditionFormat(ctx, &lark.UpdateSheetConditionFormatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新已有的条件格式，单次最多支持更新10个条件格式，每个条件格式的更新会返回成功或者失败，失败的情况包括各种参数的校验。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-update](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-update)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_update`

### Method

`POST`

## DeleteSheetConditionFormat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteSheetConditionFormat(ctx, &lark.DeleteSheetConditionFormatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除已有的条件格式，单次最多支持删除10个条件格式，每个条件格式的删除会返回成功或者失败，失败的情况包括各种参数的校验。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-delete](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-delete)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_delete`

### Method

`DELETE`

## CreateSheetProtectedDimension

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateSheetProtectedDimension(ctx, &lark.CreateSheetProtectedDimensionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 spreadsheetToken 和维度信息增加多个保护范围；单次操作不超过5000行或列。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ugDNzUjL4QzM14CO0MTN](https://open.feishu.cn/document/ukTMukTMukTM/ugDNzUjL4QzM14CO0MTN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_dimension`

### Method

`POST`

## GetSheetProtectedDimension

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetSheetProtectedDimension(ctx, &lark.GetSheetProtectedDimensionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据保护范围ID查询详细的保护行列信息，最多支持同时查询5个ID。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQTM5YjL0ETO24CNxkjN](https://open.feishu.cn/document/ukTMukTMukTM/uQTM5YjL0ETO24CNxkjN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_get`

### Method

`GET`

## UpdateSheetProtectedDimension

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateSheetProtectedDimension(ctx, &lark.UpdateSheetProtectedDimensionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据保护范围ID修改保护范围，单次最多支持同时修改10个ID。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUTM5YjL1ETO24SNxkjN](https://open.feishu.cn/document/ukTMukTMukTM/uUTM5YjL1ETO24SNxkjN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_update`

### Method

`POST`

## DeleteSheetProtectedDimension

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteSheetProtectedDimension(ctx, &lark.DeleteSheetProtectedDimensionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据保护范围ID删除保护范围，最多支持同时删除10个ID。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYTM5YjL2ETO24iNxkjN](https://open.feishu.cn/document/ukTMukTMukTM/uYTM5YjL2ETO24iNxkjN)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_del`

### Method

`DELETE`

## CreateSheetDataValidationDropdown

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateSheetDataValidationDropdown(ctx, &lark.CreateSheetDataValidationDropdownReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口根据 spreadsheetToken 、range 和下拉列表属性给单元格设置下拉列表规则；单次设置范围不超过5000行，100列。当一个数据区域中已有数据，支持将有效数据直接转为选项。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/set-dropdown](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/set-dropdown)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation`

### Method

`POST`

## DeleteSheetDataValidationDropdown

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteSheetDataValidationDropdown(ctx, &lark.DeleteSheetDataValidationDropdownReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口根据 spreadsheetToken 、range 移除选定数据范围单元格的下拉列表设置，但保留选项文本。单个删除范围不超过5000单元格。单次请求range最大数量100个。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/delete-datavalidation](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/delete-datavalidation)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation`

### Method

`DELETE`

## UpdateSheetDataValidationDropdown

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateSheetDataValidationDropdown(ctx, &lark.UpdateSheetDataValidationDropdownReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口根据 spreadsheetToken 、sheetId、dataValidationId 更新下拉列表的属性。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/update-datavalidation](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/update-datavalidation)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation/:sheetId/:dataValidationId`

### Method

`PUT`

## GetSheetDataValidationDropdown

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetSheetDataValidationDropdown(ctx, &lark.GetSheetDataValidationDropdownReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口根据 spreadsheetToken 、range 查询range内的下拉列表设置信息；单次查询范围不超过5000行，100列。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/query-datavalidation](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/query-datavalidation)

### URL

`https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation`

### Method

`GET`

## CreateSheetFilter

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateSheetFilter(ctx, &lark.CreateSheetFilterReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

参数值可参考[筛选指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/filter-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/create](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/create)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter`

### Method

`POST`

## DeleteSheetFilter

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteSheetFilter(ctx, &lark.DeleteSheetFilterReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除子表的筛选

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/delete](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/delete)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter`

### Method

`DELETE`

## UpdateSheetFilter

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateSheetFilter(ctx, &lark.UpdateSheetFilterReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

参数值可参考[筛选指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/filter-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/update](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/update)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter`

### Method

`PUT`

## GetSheetFilter

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetSheetFilter(ctx, &lark.GetSheetFilterReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取子表的详细筛选信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/get](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/get)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter`

### Method

`GET`

## CreateSheetFilterView

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateSheetFilterView(ctx, &lark.CreateSheetFilterViewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

筛选范围的设置参考：[筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/create](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/create)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views`

### Method

`POST`

## DeleteSheetFilterView

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteSheetFilterView(ctx, &lark.DeleteSheetFilterViewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除指定 id 对应的筛选视图。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/delete](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/delete)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id`

### Method

`DELETE`

## UpdateSheetFilterView

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateSheetFilterView(ctx, &lark.UpdateSheetFilterViewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

筛选范围的设置参考：[筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/patch](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/patch)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id`

### Method

`PATCH`

## GetSheetFilterView

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetSheetFilterView(ctx, &lark.GetSheetFilterViewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取指定筛选视图 id 的名字和筛选范围。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/get](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/get)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id`

### Method

`GET`

## QuerySheetFilterView

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.QuerySheetFilterView(ctx, &lark.QuerySheetFilterViewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

查询子表内所有的筛选视图基本信息，包括 id、name 和 range

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/query](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/query)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/query`

### Method

`GET`

## CreateSheetFilterViewCondition

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateSheetFilterViewCondition(ctx, &lark.CreateSheetFilterViewConditionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

筛选条件参考 [筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/create](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/create)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions`

### Method

`POST`

## DeleteSheetFilterViewCondition

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteSheetFilterViewCondition(ctx, &lark.DeleteSheetFilterViewConditionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除筛选视图的筛选范围某一列的筛选条件。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/delete](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/delete)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id`

### Method

`DELETE`

## UpdateSheetFilterViewCondition

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateSheetFilterViewCondition(ctx, &lark.UpdateSheetFilterViewConditionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

筛选条件参数可参考 [筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/update](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/update)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id`

### Method

`PUT`

## GetSheetFilterViewCondition

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetSheetFilterViewCondition(ctx, &lark.GetSheetFilterViewConditionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

筛选条件含义参考 [筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/get](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/get)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id`

### Method

`GET`

## QuerySheetFilterViewCondition

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.QuerySheetFilterViewCondition(ctx, &lark.QuerySheetFilterViewConditionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

筛选条件含义可参考 [筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/query](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/query)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/query`

### Method

`GET`

## CreateSheetFloatImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateSheetFloatImage(ctx, &lark.CreateSheetFloatImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

浮动图片的设置参考：[浮动图片指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/create](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/create)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images`

### Method

`POST`

## DeleteSheetFloatImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteSheetFloatImage(ctx, &lark.DeleteSheetFloatImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除 float_image_id 对应的浮动图片。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/delete](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/delete)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id`

### Method

`DELETE`

## UpdateSheetFloatImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateSheetFloatImage(ctx, &lark.UpdateSheetFloatImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

浮动图片更新参考：[浮动图片指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/patch](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/patch)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id`

### Method

`PATCH`

## GetSheetFloatImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetSheetFloatImage(ctx, &lark.GetSheetFloatImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

浮动图片参考：[浮动图片指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/get](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/get)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id`

### Method

`GET`

## QuerySheetFloatImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.QuerySheetFloatImage(ctx, &lark.QuerySheetFloatImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

浮动图片参考：[浮动图片指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/query](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/query)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/query`

### Method

`GET`


# Bitable

## GetBitableViewList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.GetBitableViewList(ctx, &lark.GetBitableViewListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据 app_token 和 table_id，获取数据表的所有视图

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/list)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views`

### Method

`GET`

## CreateBitableView

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.CreateBitableView(ctx, &lark.CreateBitableViewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

在数据表中新增一个视图

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/create)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views`

### Method

`POST`

## DeleteBitableView

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.DeleteBitableView(ctx, &lark.DeleteBitableViewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除数据表中的视图

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/delete)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views/:view_id`

### Method

`DELETE`

## GetBitableRecordList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.GetBitableRecordList(ctx, &lark.GetBitableRecordListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于列出数据表中的现有记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/list)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records`

### Method

`GET`

## GetBitableRecord

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.GetBitableRecord(ctx, &lark.GetBitableRecordReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据 record_id 的值检索现有记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/get)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id`

### Method

`GET`

## CreateBitableRecord

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.CreateBitableRecord(ctx, &lark.CreateBitableRecordReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于在数据表中新增一条记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/create)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records`

### Method

`POST`

## BatchCreateBitableRecord

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.BatchCreateBitableRecord(ctx, &lark.BatchCreateBitableRecordReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于在数据表中新增多条记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_create)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_create`

### Method

`POST`

## UpdateBitableRecord

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.UpdateBitableRecord(ctx, &lark.UpdateBitableRecordReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新数据表中的一条记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/update)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id`

### Method

`PUT`

## BatchUpdateBitableRecord

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.BatchUpdateBitableRecord(ctx, &lark.BatchUpdateBitableRecordReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新数据表中的多条记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_update)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_update`

### Method

`POST`

## DeleteBitableRecord

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.DeleteBitableRecord(ctx, &lark.DeleteBitableRecordReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除数据表中的一条记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/delete)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id`

### Method

`DELETE`

## BatchDeleteBitableRecord

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.BatchDeleteBitableRecord(ctx, &lark.BatchDeleteBitableRecordReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除数据表中现有的多条记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_delete)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_delete`

### Method

`POST`

## GetBitableFieldList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.GetBitableFieldList(ctx, &lark.GetBitableFieldListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据 app_token 和 table_id，获取数据表的所有字段

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/list)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields`

### Method

`GET`

## CreateBitableField

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.CreateBitableField(ctx, &lark.CreateBitableFieldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于在数据表中新增一个字段

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/create)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields`

### Method

`POST`

## UpdateBitableField

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.UpdateBitableField(ctx, &lark.UpdateBitableFieldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于在数据表中更新一个字段

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/update)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id`

### Method

`PUT`

## DeleteBitableField

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.DeleteBitableField(ctx, &lark.DeleteBitableFieldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于在数据表中删除一个字段

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/delete)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id`

### Method

`DELETE`

## GetBitableTableList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.GetBitableTableList(ctx, &lark.GetBitableTableListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据  app_token，获取多维表格下的所有数据表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/list)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables`

### Method

`GET`

## CreateBitableTable

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.CreateBitableTable(ctx, &lark.CreateBitableTableReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

新增一个数据表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/create)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables`

### Method

`POST`

## BatchCreateBitableTable

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.BatchCreateBitableTable(ctx, &lark.BatchCreateBitableTableReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

新增多个数据表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/batch_create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/batch_create)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/batch_create`

### Method

`POST`

## DeleteBitableTable

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.DeleteBitableTable(ctx, &lark.DeleteBitableTableReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除一个数据表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/delete)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id`

### Method

`DELETE`

## BatchDeleteBitableTable

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.BatchDeleteBitableTable(ctx, &lark.BatchDeleteBitableTableReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除多个数据表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/batch_delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/batch_delete)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/batch_delete`

### Method

`POST`

## GetBitableMeta

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.GetBitableMeta(ctx, &lark.GetBitableMetaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

通过 app_token 获取多维表格元数据

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/get)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token`

### Method

`GET`


# MeetingRoom

## BatchGetMeetingRoomSummary

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.BatchGetMeetingRoomSummary(ctx, &lark.BatchGetMeetingRoomSummaryReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

通过日程的Uid和Original time，查询会议室日程主题。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uIjM5UjLyITO14iMykTN/](https://open.feishu.cn/document/ukTMukTMukTM/uIjM5UjLyITO14iMykTN/)

### URL

`https://open.feishu.cn/open-apis/meeting_room/summary/batch_get`

### Method

`POST`

## GetMeetingRoomBuildingList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.GetMeetingRoomBuildingList(ctx, &lark.GetMeetingRoomBuildingListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取本企业下的建筑物（办公大楼）。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ugzNyUjL4cjM14CO3ITN](https://open.feishu.cn/document/ukTMukTMukTM/ugzNyUjL4cjM14CO3ITN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/building/list`

### Method

`GET`

## BatchGetMeetingRoomBuilding

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.BatchGetMeetingRoomBuilding(ctx, &lark.BatchGetMeetingRoomBuildingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取指定建筑物的详细信息。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukzNyUjL5cjM14SO3ITN](https://open.feishu.cn/document/ukTMukTMukTM/ukzNyUjL5cjM14SO3ITN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/building/batch_get`

### Method

`GET`

## GetMeetingRoomRoomList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.GetMeetingRoomRoomList(ctx, &lark.GetMeetingRoomRoomListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取指定建筑下的会议室。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uADOyUjLwgjM14CM4ITN](https://open.feishu.cn/document/ukTMukTMukTM/uADOyUjLwgjM14CM4ITN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/room/list`

### Method

`GET`

## BatchGetMeetingRoomRoom

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.BatchGetMeetingRoomRoom(ctx, &lark.BatchGetMeetingRoomRoomReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取指定会议室的详细信息。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uEDOyUjLxgjM14SM4ITN](https://open.feishu.cn/document/ukTMukTMukTM/uEDOyUjLxgjM14SM4ITN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/room/batch_get`

### Method

`GET`

## BatchGetMeetingRoomFreebusy

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.BatchGetMeetingRoomFreebusy(ctx, &lark.BatchGetMeetingRoomFreebusyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取指定会议室的忙闲日程实例列表。非重复日程只有唯一实例；重复日程可能存在多个实例，依据重复规则和时间范围扩展。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uIDOyUjLygjM14iM4ITN](https://open.feishu.cn/document/ukTMukTMukTM/uIDOyUjLygjM14iM4ITN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/freebusy/batch_get`

### Method

`GET`

## ReplyMeetingRoomInstance

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.ReplyMeetingRoomInstance(ctx, &lark.ReplyMeetingRoomInstanceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于回复会议室日程实例，包括未签到释放和提前结束释放。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYzN4UjL2cDO14iN3gTN](https://open.feishu.cn/document/ukTMukTMukTM/uYzN4UjL2cDO14iN3gTN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/instance/reply`

### Method

`POST`

## CreateMeetingRoomBuilding

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.CreateMeetingRoomBuilding(ctx, &lark.CreateMeetingRoomBuildingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口对应管理后台的添加建筑，添加楼层的功能，可用于创建建筑物和建筑物的楼层信息。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATNwYjLwUDM24CM1AjN](https://open.feishu.cn/document/ukTMukTMukTM/uATNwYjLwUDM24CM1AjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/building/create`

### Method

`POST`

## UpdateMeetingRoomBuilding

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.UpdateMeetingRoomBuilding(ctx, &lark.UpdateMeetingRoomBuildingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于编辑建筑信息，添加楼层，删除楼层，编辑楼层信息。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uETNwYjLxUDM24SM1AjN](https://open.feishu.cn/document/ukTMukTMukTM/uETNwYjLxUDM24SM1AjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/building/update`

### Method

`POST`

## DeleteMeetingRoomBuilding

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.DeleteMeetingRoomBuilding(ctx, &lark.DeleteMeetingRoomBuildingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除建筑物（办公大楼）。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzMxYjLzMTM24yMzEjN](https://open.feishu.cn/document/ukTMukTMukTM/uMzMxYjLzMTM24yMzEjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/building/delete`

### Method

`POST`

## BatchGetMeetingRoomBuildingID

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.BatchGetMeetingRoomBuildingID(ctx, &lark.BatchGetMeetingRoomBuildingIDReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据租户自定义建筑 ID 查询建筑 ID。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQzMxYjL0MTM24CNzEjN](https://open.feishu.cn/document/ukTMukTMukTM/uQzMxYjL0MTM24CNzEjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/building/batch_get_id`

### Method

`GET`

## CreateMeetingRoomRoom

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.CreateMeetingRoomRoom(ctx, &lark.CreateMeetingRoomRoomReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于创建会议室。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITNwYjLyUDM24iM1AjN](https://open.feishu.cn/document/ukTMukTMukTM/uITNwYjLyUDM24iM1AjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/room/create`

### Method

`POST`

## UpdateMeetingRoomRoom

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.UpdateMeetingRoomRoom(ctx, &lark.UpdateMeetingRoomRoomReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新会议室。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMTNwYjLzUDM24yM1AjN](https://open.feishu.cn/document/ukTMukTMukTM/uMTNwYjLzUDM24yM1AjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/room/update`

### Method

`POST`

## DeleteMeetingRoomRoom

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.DeleteMeetingRoomRoom(ctx, &lark.DeleteMeetingRoomRoomReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除会议室。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUzMxYjL1MTM24SNzEjN](https://open.feishu.cn/document/ukTMukTMukTM/uUzMxYjL1MTM24SNzEjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/room/delete`

### Method

`POST`

## BatchGetMeetingRoomRoomID

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.BatchGetMeetingRoomRoomID(ctx, &lark.BatchGetMeetingRoomRoomIDReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于根据租户自定义会议室ID查询会议室ID。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYzMxYjL2MTM24iNzEjN](https://open.feishu.cn/document/ukTMukTMukTM/uYzMxYjL2MTM24iNzEjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/room/batch_get_id`

### Method

`GET`

## GetMeetingRoomCountryList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.GetMeetingRoomCountryList(ctx, &lark.GetMeetingRoomCountryListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

新建建筑时需要标明所处国家/地区，该接口用于获得系统预先提供的可供选择的国家 /地区列表。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQTNwYjL0UDM24CN1AjN](https://open.feishu.cn/document/ukTMukTMukTM/uQTNwYjL0UDM24CN1AjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/country/list`

### Method

`GET`

## GetMeetingRoomDistrictList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.GetMeetingRoomDistrictList(ctx, &lark.GetMeetingRoomDistrictListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

新建建筑时需要选择所处国家/地区，该接口用于获得系统预先提供的可供选择的城市列表。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUTNwYjL1UDM24SN1AjN](https://open.feishu.cn/document/ukTMukTMukTM/uUTNwYjL1UDM24SN1AjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/district/list`

### Method

`GET`


# VC

## ApplyVCReserve

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.ApplyVCReserve(ctx, &lark.ApplyVCReserveReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

创建一个会议预约。

支持预约最近30天内的会议（到期时间距离当前时间不超过30天），预约到期后会议号将被释放，如需继续使用可通过"更新预约"接口进行续期；预约会议时可配置参会人在会中的权限，以达到控制会议的目的







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/apply](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/apply)

### URL

`https://open.feishu.cn/open-apis/vc/v1/reserves/apply`

### Method

`POST`

## UpdateVCReserve

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.UpdateVCReserve(ctx, &lark.UpdateVCReserveReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新一个预约

只能更新归属于自己的预约，不需要更新的字段不传（如果传空则会被更新为空）；可用于续期操作，到期时间距离当前时间不超过30天







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/update)

### URL

`https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id`

### Method

`PUT`

## DeleteVCReserve

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.DeleteVCReserve(ctx, &lark.DeleteVCReserveReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除一个预约

只能删除归属于自己的预约；删除后数据不可恢复







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/delete)

### URL

`https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id`

### Method

`DELETE`

## GetVCReserve

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.GetVCReserve(ctx, &lark.GetVCReserveReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取一个预约的详情

只能获取归属于自己的预约







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/get)

### URL

`https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id`

### Method

`GET`

## GetVCReserveActiveMeeting

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.GetVCReserveActiveMeeting(ctx, &lark.GetVCReserveActiveMeetingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取一个预约的当前活跃会议

只能获取归属于自己的预约的活跃会议（一个预约最多有一个正在进行中的会议）







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/get_active_meeting](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/get_active_meeting)

### URL

`https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id/get_active_meeting`

### Method

`GET`

## GetVCMeeting

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.GetVCMeeting(ctx, &lark.GetVCMeetingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取一个会议的详细数据

只能获取归属于自己（或参与）的会议，支持查询最近90天内的会议







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/get)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id`

### Method

`GET`

## InviteVCMeeting

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.InviteVCMeeting(ctx, &lark.InviteVCMeetingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

邀请参会人进入会议

发起邀请的操作者必须具有相应的权限（如果操作者为用户，则必须在会中），如果会议被锁定、或参会人数如果达到上限，则会邀请失败







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/invite](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/invite)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/invite`

### Method

`PATCH`

## SetVCHostMeeting

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.SetVCHostMeeting(ctx, &lark.SetVCHostMeetingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

设置会议的主持人

发起设置主持人的操作者必须具有相应的权限（如果操作者为用户，必须是会中当前主持人）；该操作使用CAS并发安全机制，需传入会中当前主持人，如果操作失败可使用返回的最新数据重试







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/set_host](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/set_host)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/set_host`

### Method

`PATCH`

## EndVCMeeting

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.EndVCMeeting(ctx, &lark.EndVCMeetingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

结束一个进行中的会议

会议正在进行中，且操作者须具有相应的权限（如果操作者为用户，必须是会中当前主持人）







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/end](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/end)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/end`

### Method

`PATCH`

## StartVCMeetingRecording

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.StartVCMeetingRecording(ctx, &lark.StartVCMeetingRecordingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

在会议中开始录制。

会议正在进行中，且操作者具有相应权限（如果操作者为用户，必须是会中当前主持人）







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/start](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/start)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/start`

### Method

`PATCH`

## StopVCMeetingRecording

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.StopVCMeetingRecording(ctx, &lark.StopVCMeetingRecordingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

在会议中停止录制。

会议正在录制中，且操作者具有相应权限（如果操作者为用户，必须是会中当前主持人）







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/stop](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/stop)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/stop`

### Method

`PATCH`

## GetVCMeetingRecording

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.GetVCMeetingRecording(ctx, &lark.GetVCMeetingRecordingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取一个会议的录制文件。

会议结束后并且收到了"录制完成"的事件方可获取录制文件；只有会议owner（通过开放平台预约的会议即为预约人）有权限获取；录制时间太短(&lt;5s)有可能无法生成录制文件







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/get)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording`

### Method

`GET`

## SetVCPermissionMeetingRecording

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.SetVCPermissionMeetingRecording(ctx, &lark.SetVCPermissionMeetingRecordingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

将一个会议的录制文件授权给组织、用户或公开到公网

会议结束后并且收到了"录制完成"的事件方可进行授权；会议owner（通过开放平台预约的会议即为预约人）才有权限操作







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/set_permission](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/set_permission)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/set_permission`

### Method

`PATCH`

## GetVCDailyReport

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.GetVCDailyReport(ctx, &lark.GetVCDailyReportReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取一段时间内组织的每日会议使用报告。

支持最近90天内的数据查询







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/report/get_daily](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/report/get_daily)

### URL

`https://open.feishu.cn/open-apis/vc/v1/reports/get_daily`

### Method

`GET`

## GetVCTopUserReport

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.GetVCTopUserReport(ctx, &lark.GetVCTopUserReportReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取一段时间内组织内会议使用的top用户列表。

支持最近90天内的数据查询；默认返回前10位，最多可查询前100位





#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/report/get_top_user](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/report/get_top_user)

### URL

`https://open.feishu.cn/open-apis/vc/v1/reports/get_top_user`

### Method

`GET`

## GetVCRoomConfig

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.GetVCRoomConfig(ctx, &lark.GetVCRoomConfigReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

查询一个范围内的会议室配置。

根据查询范围传入对应的参数





#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_config/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_config/query)

### URL

`https://open.feishu.cn/open-apis/vc/v1/room_configs/query`

### Method

`GET`

## SetVCRoomConfig

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.SetVCRoomConfig(ctx, &lark.SetVCRoomConfigReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

设置一个范围内的会议室配置。

根据设置范围传入对应的参数





#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_config/set](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_config/set)

### URL

`https://open.feishu.cn/open-apis/vc/v1/room_configs/set`

### Method

`POST`


# Application

## IsApplicationUserAdmin

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.IsApplicationUserAdmin(ctx, &lark.IsApplicationUserAdminReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于查询用户是否为应用管理员。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITN1EjLyUTNx4iM1UTM](https://open.feishu.cn/document/ukTMukTMukTM/uITN1EjLyUTNx4iM1UTM)

### URL

`https://open.feishu.cn/open-apis/application/v3/is_user_admin`

### Method

`GET`

## GetApplicationUserAdminScope

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationUserAdminScope(ctx, &lark.GetApplicationUserAdminScopeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于获取应用管理员的管理范围，即该应用管理员能够管理哪些部门。  

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzN3QjLzczN04yM3cDN](https://open.feishu.cn/document/ukTMukTMukTM/uMzN3QjLzczN04yM3cDN)

### URL

`https://open.feishu.cn/open-apis/contact/v1/user/admin_scope/get`

### Method

`GET`

## GetApplicationAppVisibility

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationAppVisibility(ctx, &lark.GetApplicationAppVisibilityReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于查询应用在该企业内可以被使用的范围，只能被企业自建应用调用。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uIjM3UjLyIzN14iMycTN](https://open.feishu.cn/document/ukTMukTMukTM/uIjM3UjLyIzN14iMycTN)

### URL

`https://open.feishu.cn/open-apis/application/v2/app/visibility`

### Method

`GET`

## GetApplicationUserVisibleApp

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationUserVisibleApp(ctx, &lark.GetApplicationUserVisibleAppReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于查询用户可用的应用列表，只能被企业自建应用调用。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMjM3UjLzIzN14yMycTN](https://open.feishu.cn/document/ukTMukTMukTM/uMjM3UjLzIzN14yMycTN)

### URL

`https://open.feishu.cn/open-apis/application/v1/user/visible_apps`

### Method

`GET`

## GetApplicationAppList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationAppList(ctx, &lark.GetApplicationAppListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于查询企业安装的应用列表，只能被企业自建应用调用。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYDN3UjL2QzN14iN0cTN](https://open.feishu.cn/document/ukTMukTMukTM/uYDN3UjL2QzN14iN0cTN)

### URL

`https://open.feishu.cn/open-apis/application/v3/app/list`

### Method

`GET`

## UpdateApplicationAppVisibility

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.UpdateApplicationAppVisibility(ctx, &lark.UpdateApplicationAppVisibilityReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于增加或者删除指定应用被哪些人可用，只能被企业自建应用调用。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ucDN3UjL3QzN14yN0cTN](https://open.feishu.cn/document/ukTMukTMukTM/ucDN3UjL3QzN14yN0cTN)

### URL

`https://open.feishu.cn/open-apis/application/v3/app/update_visibility`

### Method

`POST`

## GetApplicationAppAdminUserList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationAppAdminUserList(ctx, &lark.GetApplicationAppAdminUserListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



查询应用管理员列表，返回应用的最新10个管理员账户id列表。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ucDOwYjL3gDM24yN4AjN](https://open.feishu.cn/document/ukTMukTMukTM/ucDOwYjL3gDM24yN4AjN)

### URL

`https://open.feishu.cn/open-apis/user/v4/app_admin_user/list`

### Method

`GET`

## CheckUserIsInApplicationPaidScope

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.CheckUserIsInApplicationPaidScope(ctx, &lark.CheckUserIsInApplicationPaidScopeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



当付费套餐是按人数收费 或者 限制最大使用人数时，开放平台会引导企业管理员设置“付费功能开通范围”。  但是受开通范围限制，部分用户就无法使用对应的付费功能。  可以通过此接口，在付费功能点入口判断是否允许某个用户进入使用。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uATNwUjLwUDM14CM1ATN](https://open.feishu.cn/document/ukTMukTMukTM/uATNwUjLwUDM14CM1ATN)

### URL

`https://open.feishu.cn/open-apis/pay/v1/paid_scope/check_user`

### Method

`GET`

## GetApplicationOrderList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationOrderList(ctx, &lark.GetApplicationOrderListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于分页查询应用租户下的已付费订单，每次购买对应一个唯一的订单，订单会记录购买的套餐的相关信息，业务方需要自行处理套餐的有效期和付费方案的升级。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uETNwUjLxUDM14SM1ATN](https://open.feishu.cn/document/ukTMukTMukTM/uETNwUjLxUDM14SM1ATN)

### URL

`https://open.feishu.cn/open-apis/pay/v1/order/list`

### Method

`GET`

## GetApplicationOrder

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationOrder(ctx, &lark.GetApplicationOrderReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于查询某个订单的具体信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITNwUjLyUDM14iM1ATN](https://open.feishu.cn/document/ukTMukTMukTM/uITNwUjLyUDM14iM1ATN)

### URL

`https://open.feishu.cn/open-apis/pay/v1/order/get`

### Method

`GET`

## GetApplicationUsageOverview

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationUsageOverview(ctx, &lark.GetApplicationUsageOverviewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



查询应用在指定时间段内企业员工的使用概览信息。
:::warning
此接口目前仅支持小程序的使用情况查询，不支持网页应用和机器人应用的使用情况查询；
仅支持查询自建应用，不支持查询商店应用
:::

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uETN0YjLxUDN24SM1QjN](https://open.feishu.cn/document/ukTMukTMukTM/uETN0YjLxUDN24SM1QjN)

### URL

`https://open.feishu.cn/open-apis/application/v1/app_usage_overview`

### Method

`POST`

## GetApplicationUsageTrend

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationUsageTrend(ctx, &lark.GetApplicationUsageTrendReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



查询应用在指定时间段内企业员工的使用趋势信息。
:::warning
此接口目前仅支持小程序的使用情况查询，不支持网页应用和机器人应用的使用情况查询;仅支持查询自建应用，不支持查询商店应用
:::

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITN0YjLyUDN24iM1QjN](https://open.feishu.cn/document/ukTMukTMukTM/uITN0YjLyUDN24iM1QjN)

### URL

`https://open.feishu.cn/open-apis/application/v1/app_usage_trend`

### Method

`POST`

## GetApplicationUsageDetail

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationUsageDetail(ctx, &lark.GetApplicationUsageDetailReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



查询应用在指定时间段内企业员工的使用详细信息。
:::warning
此接口目前仅支持小程序的使用情况查询，不支持网页应用和机器人应用的使用情况查询;仅支持查询自建应用，不支持查询商店应用
:::

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMTN0YjLzUDN24yM1QjN](https://open.feishu.cn/document/ukTMukTMukTM/uMTN0YjLzUDN24yM1QjN)

### URL

`https://open.feishu.cn/open-apis/application/v1/app_usage_detail`

### Method

`POST`

## GetApplicationMessageOverview

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationMessageOverview(ctx, &lark.GetApplicationMessageOverviewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



查询应用在指定时间段内机器人消息概览信息。

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQTN0YjL0UDN24CN1QjN](https://open.feishu.cn/document/ukTMukTMukTM/uQTN0YjL0UDN24CN1QjN)

### URL

`https://open.feishu.cn/open-apis/application/v1/app_message_overview`

### Method

`POST`

## GetApplicationMessageTrend

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationMessageTrend(ctx, &lark.GetApplicationMessageTrendReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



查询应用在指定时间段内机器人消息趋势信息。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUTN0YjL1UDN24SN1QjN](https://open.feishu.cn/document/ukTMukTMukTM/uUTN0YjL1UDN24SN1QjN)

### URL

`https://open.feishu.cn/open-apis/application/v1/app_message_trend`

### Method

`POST`

## GetApplicationMessageDetail

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationMessageDetail(ctx, &lark.GetApplicationMessageDetailReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



查询应用在指定时间段内机器人消息详细信息。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYTN0YjL2UDN24iN1QjN](https://open.feishu.cn/document/ukTMukTMukTM/uYTN0YjL2UDN24iN1QjN)

### URL

`https://open.feishu.cn/open-apis/application/v1/app_message_detail`

### Method

`POST`


# Mail

## CreateMailGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.CreateMailGroup(ctx, &lark.CreateMailGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

创建一个邮件组

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/create)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups`

### Method

`POST`

## GetMailGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetMailGroup(ctx, &lark.GetMailGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取特定邮件组信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/get)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id`

### Method

`GET`

## GetMailGroupList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetMailGroupList(ctx, &lark.GetMailGroupListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

分页批量获取邮件组

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/list)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups`

### Method

`GET`

## UpdateMailGroupPatch

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.UpdateMailGroupPatch(ctx, &lark.UpdateMailGroupPatchReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新邮件组部分字段，没有填写的字段不会被更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/patch)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id`

### Method

`PATCH`

## UpdateMailGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.UpdateMailGroup(ctx, &lark.UpdateMailGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新邮件组所有信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/update)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id`

### Method

`PUT`

## DeleteMailGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.DeleteMailGroup(ctx, &lark.DeleteMailGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除一个邮件组

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/delete)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id`

### Method

`DELETE`

## CreateMailGroupMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.CreateMailGroupMember(ctx, &lark.CreateMailGroupMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

向邮件组添加单个成员

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/create)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/members`

### Method

`POST`

## GetMailGroupMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetMailGroupMember(ctx, &lark.GetMailGroupMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取邮件组单个成员信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/get)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id`

### Method

`GET`

## GetMailGroupMemberList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetMailGroupMemberList(ctx, &lark.GetMailGroupMemberListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

分页批量获取邮件组成员列表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/list)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/members`

### Method

`GET`

## DeleteMailGroupMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.DeleteMailGroupMember(ctx, &lark.DeleteMailGroupMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除邮件组单个成员

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/delete)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id`

### Method

`DELETE`

## CreateMailGroupPermissionMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.CreateMailGroupPermissionMember(ctx, &lark.CreateMailGroupPermissionMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

向邮件组添加单个自定义权限成员，添加后该成员可发送邮件到该邮件组

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/create)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members`

### Method

`POST`

## GetMailGroupPermissionMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetMailGroupPermissionMember(ctx, &lark.GetMailGroupPermissionMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取邮件组单个权限成员信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/get)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/:permission_member_id`

### Method

`GET`

## GetMailGroupPermissionMemberList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetMailGroupPermissionMemberList(ctx, &lark.GetMailGroupPermissionMemberListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

分页批量获取邮件组权限成员列表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/list)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members`

### Method

`GET`

## DeleteMailGroupPermissionMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.DeleteMailGroupPermissionMember(ctx, &lark.DeleteMailGroupPermissionMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

从自定义成员中删除单个成员，删除后该成员无法发送邮件到该邮件组

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/delete)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/:permission_member_id`

### Method

`DELETE`

## CreatePublicMailbox

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.CreatePublicMailbox(ctx, &lark.CreatePublicMailboxReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

创建一个公共邮箱

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/create)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes`

### Method

`POST`

## GetPublicMailbox

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetPublicMailbox(ctx, &lark.GetPublicMailboxReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取公共邮箱信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/get)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id`

### Method

`GET`

## GetPublicMailboxList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetPublicMailboxList(ctx, &lark.GetPublicMailboxListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

分页批量获取公共邮箱列表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/list)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes`

### Method

`GET`

## UpdatePublicMailboxPatch

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.UpdatePublicMailboxPatch(ctx, &lark.UpdatePublicMailboxPatchReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新公共邮箱部分字段，没有填写的字段不会被更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/patch)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id`

### Method

`PATCH`

## UpdatePublicMailbox

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.UpdatePublicMailbox(ctx, &lark.UpdatePublicMailboxReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新公共邮箱所有信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/update)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id`

### Method

`PUT`

## CreatePublicMailboxMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.CreatePublicMailboxMember(ctx, &lark.CreatePublicMailboxMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

向公共邮箱添加单个成员

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/create)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members`

### Method

`POST`

## GetPublicMailboxMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetPublicMailboxMember(ctx, &lark.GetPublicMailboxMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取公共邮箱单个成员信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/get)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id`

### Method

`GET`

## GetPublicMailboxMemberList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetPublicMailboxMemberList(ctx, &lark.GetPublicMailboxMemberListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

分页批量获取公共邮箱成员列表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/list)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members`

### Method

`GET`

## DeletePublicMailboxMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.DeletePublicMailboxMember(ctx, &lark.DeletePublicMailboxMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除公共邮箱单个成员

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/delete)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id`

### Method

`DELETE`

## ClearPublicMailboxMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.ClearPublicMailboxMember(ctx, &lark.ClearPublicMailboxMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除公共邮箱所有成员

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/clear](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/clear)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/clear`

### Method

`POST`


# Approval

## GetApproval

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.GetApproval(ctx, &lark.GetApprovalReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



根据 Approval Code 获取某个审批定义的详情，用于构造创建审批实例的请求。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uADNyUjLwQjM14CM0ITN](https://open.feishu.cn/document/ukTMukTMukTM/uADNyUjLwQjM14CM0ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/approval/get`

### Method

`POST`

## GetApprovalInstanceList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.GetApprovalInstanceList(ctx, &lark.GetApprovalInstanceListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



根据 approval_code 批量获取审批实例的 instance_code，用于拉取租户下某个审批定义的全部审批实例。
默认以审批创建时间排序。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQDOyUjL0gjM14CN4ITN](https://open.feishu.cn/document/ukTMukTMukTM/uQDOyUjL0gjM14CN4ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/list`

### Method

`POST`

## GetApprovalInstance

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



通过审批实例 Instance Code  获取审批实例详情。Instance Code 由 [批量获取审批实例](/ssl:ttdoc/ukTMukTMukTM/uQDOyUjL0gjM14CN4ITN) 接口获取。


#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uEDNyUjLxQjM14SM0ITN](https://open.feishu.cn/document/ukTMukTMukTM/uEDNyUjLxQjM14SM0ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/get`

### Method

`POST`

## CreateApprovalInstance

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



创建一个审批实例，调用方需对审批定义的表单有详细了解，将按照定义的表单结构，将表单 Value 通过接口传入。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uIDNyUjLyQjM14iM0ITN](https://open.feishu.cn/document/ukTMukTMukTM/uIDNyUjLyQjM14iM0ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/create`

### Method

`POST`

## ApproveApprovalInstance

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.ApproveApprovalInstance(ctx, &lark.ApproveApprovalInstanceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



对于单个审批任务进行同意操作。同意后审批流程会流转到下一个审批人。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMDNyUjLzQjM14yM0ITN](https://open.feishu.cn/document/ukTMukTMukTM/uMDNyUjLzQjM14yM0ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/approve`

### Method

`POST`

## RejectApprovalInstance

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.RejectApprovalInstance(ctx, &lark.RejectApprovalInstanceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



对于单个审批任务进行拒绝操作。拒绝后审批流程结束。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQDNyUjL0QjM14CN0ITN](https://open.feishu.cn/document/ukTMukTMukTM/uQDNyUjL0QjM14CN0ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/reject`

### Method

`POST`

## TransferApprovalInstance

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.TransferApprovalInstance(ctx, &lark.TransferApprovalInstanceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



对于单个审批任务进行转交操作。转交后审批流程流转给被转交人。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDNyUjL1QjM14SN0ITN](https://open.feishu.cn/document/ukTMukTMukTM/uUDNyUjL1QjM14SN0ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/transfer`

### Method

`POST`

## CancelApprovalInstance

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.CancelApprovalInstance(ctx, &lark.CancelApprovalInstanceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



对于单个审批实例进行撤销操作。撤销后审批流程结束。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYDNyUjL2QjM14iN0ITN](https://open.feishu.cn/document/ukTMukTMukTM/uYDNyUjL2QjM14iN0ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/cancel`

### Method

`POST`

## UploadApprovalFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.UploadApprovalFile(ctx, &lark.UploadApprovalFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



当审批表单中有图片或附件控件时，开发者需在创建审批实例前通过审批上传文件接口将文件上传到审批系统。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDOyUjL1gjM14SN4ITN](https://open.feishu.cn/document/ukTMukTMukTM/uUDOyUjL1gjM14SN4ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/file/upload`

### Method

`POST`

## SearchApprovalInstance

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.SearchApprovalInstance(ctx, &lark.SearchApprovalInstanceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口通过不同条件查询审批系统中符合条件的审批实例列表。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQjMxYjL0ITM24CNyEjN](https://open.feishu.cn/document/ukTMukTMukTM/uQjMxYjL0ITM24CNyEjN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/search`

### Method

`POST`

## CreateApprovalCarbonCopy

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.CreateApprovalCarbonCopy(ctx, &lark.CreateApprovalCarbonCopyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



通过接口可以将当前审批实例抄送给其他人。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uADOzYjLwgzM24CM4MjN](https://open.feishu.cn/document/ukTMukTMukTM/uADOzYjLwgzM24CM4MjN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/cc`

### Method

`POST`


# Helpdesk

## StartHelpdeskService

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.StartHelpdeskService(ctx, &lark.StartHelpdeskServiceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于创建服务台对话。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/start_service](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/start_service)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/start_service`

### Method

`POST`

## GetHelpdeskTicket

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskTicket(ctx, &lark.GetHelpdeskTicketReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取单个服务台工单详情。仅支持自建应用。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/get)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id`

### Method

`GET`

## UpdateHelpdeskTicket

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.UpdateHelpdeskTicket(ctx, &lark.UpdateHelpdeskTicketReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新服务台工单详情。只会更新数据，不会触发相关操作。如修改工单状态到关单，不会关闭聊天页面。仅支持自建应用。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/update)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id`

### Method

`PUT`

## GetHelpdeskTicketList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskTicketList(ctx, &lark.GetHelpdeskTicketListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取全部工单详情。仅支持自建应用。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/list)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/tickets`

### Method

`GET`

## DownloadHelpdeskTicketImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.DownloadHelpdeskTicketImage(ctx, &lark.DownloadHelpdeskTicketImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取服务台工单消息图象。仅支持自建应用。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/ticket_image](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/ticket_image)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/ticket_images`

### Method

`GET`

## AnswerHelpdeskTicketUserQuery

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.AnswerHelpdeskTicketUserQuery(ctx, &lark.AnswerHelpdeskTicketUserQueryReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于回复用户提问结果至工单，需要工单仍处于进行中且未接入人工状态。仅支持自建应用。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/answer_user_query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/answer_user_query)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id/answer_user_query`

### Method

`POST`

## GetHelpdeskTicketMessageList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskTicketMessageList(ctx, &lark.GetHelpdeskTicketMessageListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取服务台工单消息详情。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket-message/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket-message/list)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id/messages`

### Method

`GET`

## SendHelpdeskTicketMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.SendHelpdeskTicketMessage(ctx, &lark.SendHelpdeskTicketMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于工单发送消息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket-message/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket-message/create)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id/messages`

### Method

`POST`

## GetHelpdeskTicketCustomizedFieldList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskTicketCustomizedFieldList(ctx, &lark.GetHelpdeskTicketCustomizedFieldListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于获取全部工单自定义字段。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/list-ticket-customized-fields](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/list-ticket-customized-fields)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields`

### Method

`GET`

## DeleteHelpdeskTicketCustomizedField

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.DeleteHelpdeskTicketCustomizedField(ctx, &lark.DeleteHelpdeskTicketCustomizedFieldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除工单自定义字段。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/delete)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields/:ticket_customized_field_id`

### Method

`DELETE`

## UpdateHelpdeskTicketCustomizedField

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.UpdateHelpdeskTicketCustomizedField(ctx, &lark.UpdateHelpdeskTicketCustomizedFieldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于更新自定义字段。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/update-ticket-customized-field](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/update-ticket-customized-field)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields/:ticket_customized_field_id`

### Method

`PATCH`

## CreateHelpdeskTicketCustomizedField

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.CreateHelpdeskTicketCustomizedField(ctx, &lark.CreateHelpdeskTicketCustomizedFieldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于创建自定义字段

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/create-ticket-customized-field](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/create-ticket-customized-field)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields`

### Method

`POST`

## GetHelpdeskTicketCustomizedField

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskTicketCustomizedField(ctx, &lark.GetHelpdeskTicketCustomizedFieldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于获取工单自定义字段详情。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/get-ticket-customized-field](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/get-ticket-customized-field)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields/:ticket_customized_field_id`

### Method

`GET`

## CreateHelpdeskCategory

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.CreateHelpdeskCategory(ctx, &lark.CreateHelpdeskCategoryReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于创建知识库分类。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/create)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/categories`

### Method

`POST`

## GetHelpdeskCategory

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskCategory(ctx, &lark.GetHelpdeskCategoryReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取知识库分类。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/get)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/categories/:id`

### Method

`GET`

## UpdateHelpdeskCategory

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.UpdateHelpdeskCategory(ctx, &lark.UpdateHelpdeskCategoryReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新知识库分类详情。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/patch)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/categories/:id`

### Method

`PATCH`

## DeleteHelpdeskCategory

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.DeleteHelpdeskCategory(ctx, &lark.DeleteHelpdeskCategoryReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除知识库分类详情。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/delete)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/categories/:id`

### Method

`DELETE`

## GetHelpdeskCategoryList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskCategoryList(ctx, &lark.GetHelpdeskCategoryListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



该接口用于获取服务台知识库所有分类。



#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/list-categories](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/list-categories)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/categories`

### Method

`GET`

## CreateHelpdeskFAQ

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.CreateHelpdeskFAQ(ctx, &lark.CreateHelpdeskFAQReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于创建知识库。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/create)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/faqs`

### Method

`POST`

## GetHelpdeskFAQ

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskFAQ(ctx, &lark.GetHelpdeskFAQReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取服务台知识库详情。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/get)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id`

### Method

`GET`

## UpdateHelpdeskFAQ

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.UpdateHelpdeskFAQ(ctx, &lark.UpdateHelpdeskFAQReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于修改知识库。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/patch)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id`

### Method

`PATCH`

## DeleteHelpdeskFAQ

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.DeleteHelpdeskFAQ(ctx, &lark.DeleteHelpdeskFAQReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除知识库。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/delete)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id`

### Method

`DELETE`

## GetHelpdeskFAQList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskFAQList(ctx, &lark.GetHelpdeskFAQListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取服务台知识库详情。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/list)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/faqs`

### Method

`GET`

## GetHelpdeskFAQImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskFAQImage(ctx, &lark.GetHelpdeskFAQImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取知识库图像。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/faq_image](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/faq_image)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id/image/:image_key`

### Method

`GET`

## SearchHelpdeskFAQ

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.SearchHelpdeskFAQ(ctx, &lark.SearchHelpdeskFAQReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于搜索服务台知识库。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/search)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/faqs/search`

### Method

`GET`

## UpdateHelpdeskAgent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.UpdateHelpdeskAgent(ctx, &lark.UpdateHelpdeskAgentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新客服状态等信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent/patch)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agents/:agent_id`

### Method

`PATCH`

## GetHelpdeskAgentEmail

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskAgentEmail(ctx, &lark.GetHelpdeskAgentEmailReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取客服邮箱地址

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent/agent_email](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent/agent_email)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agent_emails`

### Method

`GET`

## CreateHelpdeskAgentSchedule

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.CreateHelpdeskAgentSchedule(ctx, &lark.CreateHelpdeskAgentScheduleReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于创建客服

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_schedule/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_schedule/create)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agent_schedules`

### Method

`POST`

## DeleteHelpdeskAgentSchedule

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.DeleteHelpdeskAgentSchedule(ctx, &lark.DeleteHelpdeskAgentScheduleReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除客服

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent-schedules/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent-schedules/delete)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agents/:agent_id/schedules`

### Method

`DELETE`

## UpdateHelpdeskAgentSchedule

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.UpdateHelpdeskAgentSchedule(ctx, &lark.UpdateHelpdeskAgentScheduleReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新客服

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent-schedules/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent-schedules/patch)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agents/:agent_id/schedules`

### Method

`PATCH`

## GetHelpdeskAgentSchedule

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskAgentSchedule(ctx, &lark.GetHelpdeskAgentScheduleReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取客服信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent-schedules/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent-schedules/get)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agents/:agent_id/schedules`

### Method

`GET`

## GetHelpdeskAgentScheduleList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskAgentScheduleList(ctx, &lark.GetHelpdeskAgentScheduleListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取所有客服信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_schedule/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_schedule/list)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agent_schedules`

### Method

`GET`

## CreateHelpdeskAgentSkill

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.CreateHelpdeskAgentSkill(ctx, &lark.CreateHelpdeskAgentSkillReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于创建客服技能

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/create)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agent_skills`

### Method

`POST`

## GetHelpdeskAgentSkill

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskAgentSkill(ctx, &lark.GetHelpdeskAgentSkillReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取客服技能

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/get)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agent_skills/:agent_skill_id`

### Method

`GET`

## UpdateHelpdeskAgentSkill

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.UpdateHelpdeskAgentSkill(ctx, &lark.UpdateHelpdeskAgentSkillReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于更新客服技能

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/patch)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agent_skills/:agent_skill_id`

### Method

`PATCH`

## DeleteHelpdeskAgentSkill

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.DeleteHelpdeskAgentSkill(ctx, &lark.DeleteHelpdeskAgentSkillReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除客服技能

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/delete)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agent_skills/:agent_skill_id`

### Method

`DELETE`

## GetHelpdeskAgentSkillList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskAgentSkillList(ctx, &lark.GetHelpdeskAgentSkillListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取全部客服技能

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/list)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agent_skills`

### Method

`GET`

## GetHelpdeskAgentSkillRuleList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskAgentSkillRuleList(ctx, &lark.GetHelpdeskAgentSkillRuleListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取全部客服技能。仅支持自建应用。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill_rule/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill_rule/list)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/agent_skill_rules`

### Method

`GET`

## SubscribeHelpdeskEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.SubscribeHelpdeskEvent(ctx, &lark.SubscribeHelpdeskEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

用于订阅服务台事件

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/event/subscribe](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/event/subscribe)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/events/subscribe`

### Method

`POST`

## UnsubscribeHelpdeskEvent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.UnsubscribeHelpdeskEvent(ctx, &lark.UnsubscribeHelpdeskEventReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

用于取消订阅服务台事件

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/event/unsubscribe](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/event/unsubscribe)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/events/unsubscribe`

### Method

`POST`


# Admin

## GetAdminDeptStats

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Admin.GetAdminDeptStats(ctx, &lark.GetAdminDeptStatsReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取部门维度的用户活跃和功能使用数据，即IM（即时通讯）、日历、云文档、音视频会议功能的使用数据。

- 只有企业自建应用才有权限调用此接口

- 当天的数据会在第二天的凌晨五点产出（UTC+8）

- 部门维度的数据最多查询最近366天（包含366天）的数据







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/admin_dept_stat/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/admin_dept_stat/list)

### URL

`https://open.feishu.cn/open-apis/admin/v1/admin_dept_stats`

### Method

`GET`

## GetAdminUserStats

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Admin.GetAdminUserStats(ctx, &lark.GetAdminUserStatsReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

用于获取用户维度的用户活跃和功能使用数据，即IM（即时通讯）、日历、云文档、音视频会议功能的使用数据。

- 只有企业自建应用才有权限调用此接口

- 当天的数据会在第二天的凌晨五点产出（UTC+8）

- 用户维度的数据最多查询最近31天的数据（包含31天）的数据








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/admin_user_stat/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/admin_user_stat/list)

### URL

`https://open.feishu.cn/open-apis/admin/v1/admin_user_stats`

### Method

`GET`


# HumanAuth

## GetFaceVerifyAuthResult

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.HumanAuth.GetFaceVerifyAuthResult(ctx, &lark.GetFaceVerifyAuthResultReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



无源人脸比对接入需申请白名单，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。
无源人脸比对流程，开发者后台通过调用此接口请求飞书后台，对本次活体比对结果做校验。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/query-recognition-result](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/query-recognition-result)

### URL

`https://open.feishu.cn/open-apis/face_verify/v1/query_auth_result`

### Method

`GET`

## UploadFaceVerifyImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.HumanAuth.UploadFaceVerifyImage(ctx, &lark.UploadFaceVerifyImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



无源人脸比对接入需申请白名单，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。
无源人脸比对流程，开发者后台通过调用此接口将基准图片上传到飞书后台，做检测时的对比使用。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/upload-facial-reference-image](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/upload-facial-reference-image)

### URL

`https://open.feishu.cn/open-apis/face_verify/v1/upload_face_image`

### Method

`POST`

## CropFaceVerifyImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.HumanAuth.CropFaceVerifyImage(ctx, &lark.CropFaceVerifyImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



无源人脸比对接入需申请白名单，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。

无源人脸比对流程，开发者后台通过调用此接口对基准图片做规范校验及处理。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/facial-image-cropping](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/facial-image-cropping)

### URL

`https://open.feishu.cn/open-apis/face_verify/v1/crop_face_image`

### Method

`POST`

## CreateIdentity

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.HumanAuth.CreateIdentity(ctx, &lark.CreateIdentityReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于录入实名认证的身份信息，在唤起有源活体认证前，需要使用该接口进行实名认证。

实名认证接口会有计费管理，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/identity/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/identity/create)

### URL

`https://open.feishu.cn/open-apis/human_authentication/v1/identities`

### Method

`POST`


# AI

## RecognizeBasicImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.AI.RecognizeBasicImage(ctx, &lark.RecognizeBasicImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

基础图片识别接口，识别图片中的文字，按区域划分返回文本列表

单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/optical_char_recognition-v1/image/basic_recognize](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/optical_char_recognition-v1/image/basic_recognize)

### URL

`https://open.feishu.cn/open-apis/optical_char_recognition/v1/image/basic_recognize`

### Method

`POST`

## RecognizeSpeechStream

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.AI.RecognizeSpeechStream(ctx, &lark.RecognizeSpeechStreamReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

语音流式接口，将整个音频文件分片进行传入模型。能够实时返回数据。建议每个音频分片的大小为 100-200ms

单租户限流：20 路（一个 stream_id 称为一路会话），同租户下的应用没有限流，共享本租户的 20路限流




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/speech_to_text-v1/speech/stream_recognize](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/speech_to_text-v1/speech/stream_recognize)

### URL

`https://open.feishu.cn/open-apis/speech_to_text/v1/speech/stream_recognize`

### Method

`POST`

## RecognizeSpeechFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.AI.RecognizeSpeechFile(ctx, &lark.RecognizeSpeechFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

语音文件识别接口，上传整段语音文件进行一次性识别。接口适合 60 秒以内音频识别

单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/speech_to_text-v1/speech/file_recognize](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/speech_to_text-v1/speech/file_recognize)

### URL

`https://open.feishu.cn/open-apis/speech_to_text/v1/speech/file_recognize`

### Method

`POST`

## TranslateText

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.AI.TranslateText(ctx, &lark.TranslateTextReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

支持中日英（zh、ja、en）三语互译

单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/translate](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/translate)

### URL

`https://open.feishu.cn/open-apis/translation/v1/text/translate`

### Method

`POST`

## DetectTextLanguage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.AI.DetectTextLanguage(ctx, &lark.DetectTextLanguageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

支持 100 多种语言识别，返回符合 ISO 693-1 标准

单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/detect](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/detect)

### URL

`https://open.feishu.cn/open-apis/translation/v1/text/detect`

### Method

`POST`

## DetectFaceAttributes

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.AI.DetectFaceAttributes(ctx, &lark.DetectFaceAttributesReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

检测图片中的人脸属性和质量等信息

注意：返回值为 -1 表示该功能还暂未实现




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/face_detection-v1/image/detect_face_attributes](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/face_detection-v1/image/detect_face_attributes)

### URL

`https://open.feishu.cn/open-apis/face_detection/v1/image/detect_face_attributes`

### Method

`POST`


# Attendance

## DownloadAttendanceFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.DownloadAttendanceFile(ctx, &lark.DownloadAttendanceFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



通过文件 ID 下载指定的文件。
#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/download-file](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/download-file)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/files/:file_id/download`

### Method

`GET`

## UploadAttendanceFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



上传文件并获取文件 ID，可用于“修改用户设置”接口的 face_key 参数。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/file_upload](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/file_upload)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/files/upload`

### Method

`POST`

## QueryAttendanceUserSettings

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.QueryAttendanceUserSettings(ctx, &lark.QueryAttendanceUserSettingsReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



批量查询授权内员工的用户设置信息，包括人脸照片文件 ID、人脸照片更新时间。
#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/batch-query-user-settings](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/batch-query-user-settings)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_settings/query`

### Method

`GET`

## UpdateAttendanceUserSettings

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.UpdateAttendanceUserSettings(ctx, &lark.UpdateAttendanceUserSettingsReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



修改授权内员工的用户设置信息，包括人脸照片文件 ID。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/user-setting-modify](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/user-setting-modify)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_settings/modify`

### Method

`POST`

## CreateUpdateAttendanceGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.CreateUpdateAttendanceGroup(ctx, &lark.CreateUpdateAttendanceGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



考勤组，是对部门或者员工在某个特定场所及特定时间段内的出勤情况（包括上下班、迟到、早退、病假、婚假、丧假、公休、工作时间、加班情况等）的一种规则设定。

通过设置考勤组，可以从部门、员工两个维度，来设定考勤方式、考勤时间、考勤地点等考勤规则。
出于安全考虑，目前通过该接口只允许修改自己创建的考勤组。
### 考勤组负责人
考勤组负责人可修改该考勤组的排班，并查看该考勤组的考勤数据。

如果考勤组负责人同时被企业管理员赋予了考勤管理员的角色，则该考勤组负责人还拥有考勤管理员的权限，可以编辑及删除考勤规则。

### 考勤组人员
可按部门、员工两个维度，设置需要参加考勤或无需参加考勤的人员。
- 若是按部门维度添加的考勤人员，当有新员工加入该部门时，其会自动加入该考勤组。
- 若是按员工维度添加的考勤人员，当其上级部门被添加到其他考勤组时，该员工不会更换考勤组。

### 考勤组类型
提供 3 种不同的考勤类型：固定班制、排班制、自由班制。
- 固定班制：指考勤组内每位人员的上下班时间一致，适用于上下班时间固定或无需安排多个班次的考勤组。
- 排班制：指考勤组人员的上下班时间不完全一致，可自定义安排每位人员的上下班时间，适用于存在多个班次如早晚班的考勤组。
- 自由班制：指没有具体的班次，考勤组人员可以在打卡时段内自由打卡，按照打卡时段统计上班时长。

### 考勤班次
- 固定班制下，需设置周一到周日每天安排哪个班次，以及可针对特殊日期进行打卡设置。
- 排班制下，需对考勤组内每一位人员的每一天进行班次指定。
- 自由班制下，需设置一天中最早打卡时间和最晚打卡时间，以及一周中哪几天需要打卡。

### 考勤方式
支持 3 种考勤方式：GPS 打卡、Wi-Fi 打卡、考勤机打卡。
- GPS 打卡：需设置经纬度信息及考勤地点名称。
- Wi-Fi 打卡：需设置 Wi-Fi 名称及 Wi-Fi 的 MAC 地址。
- 考勤机打卡：需设置考勤机名称及考勤机序号。

### 考勤其他设置
- 规则设置：支持设置是否允许外勤打卡，是否允许补卡以及一个月补卡的次数，是否允许 PC 端打卡。
- 安全设置：支持设置是否开启人脸识别打卡，以及什么情况下开启人脸识别。
- 统计设置：支持设置考勤组人员是否可以查看到某些维度的统计数据。
- 加班设置：支持配置加班时间的计算规则。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//group_create_update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//group_create_update)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/groups`

### Method

`POST`

## DeleteAttendanceGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.DeleteAttendanceGroup(ctx, &lark.DeleteAttendanceGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



通过考勤组 ID 删除考勤组。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//group_delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//group_delete)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/groups/:group_id`

### Method

`DELETE`

## GetAttendanceGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceGroup(ctx, &lark.GetAttendanceGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



通过考勤组 ID 获取考勤组详情。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//group](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//group)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/groups/:group_id`

### Method

`GET`

## CreateAttendanceShift

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.CreateAttendanceShift(ctx, &lark.CreateAttendanceShiftReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



班次是描述一次考勤任务时间规则的统称，比如一天打多少次卡，每次卡的上下班时间，晚到多长时间算迟到，晚到多长时间算缺卡等。

- 创建一个考勤组前，必须先创建一个或者多个班次。
- 一个公司内的班次是共享的，你可以直接引用他人创建的班次，但是需要注意的是，若他人修改了班次，会影响到你的考勤组及其考勤结果。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_create)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/shifts`

### Method

`POST`

## DeleteAttendanceShift

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.DeleteAttendanceShift(ctx, &lark.DeleteAttendanceShiftReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



通过班次 ID 删除班次。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_delete)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id`

### Method

`DELETE`

## GetAttendanceShiftByID

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceShiftByID(ctx, &lark.GetAttendanceShiftByIDReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



通过班次 ID 获取班次详情。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_by_id](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_by_id)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id`

### Method

`GET`

## GetAttendanceShiftByName

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceShiftByName(ctx, &lark.GetAttendanceShiftByNameReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



通过班次的名称获取班次信息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_by_name](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_by_name)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/shifts/query`

### Method

`POST`

## GetAttendanceStatisticsData

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceStatisticsData(ctx, &lark.GetAttendanceStatisticsDataReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



查询日度统计或月度统计的统计数据。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-statistics-data](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-statistics-data)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_stats_datas/query`

### Method

`POST`

## GetAttendanceStatisticsHeader

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceStatisticsHeader(ctx, &lark.GetAttendanceStatisticsHeaderReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



查询日度统计或月度统计的统计表头。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-statistics-header](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-statistics-header)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_stats_fields/query`

### Method

`POST`

## UpdateAttendanceUserStatisticsSettings

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.UpdateAttendanceUserStatisticsSettings(ctx, &lark.UpdateAttendanceUserStatisticsSettingsReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



更新日度统计或月度统计的统计设置信息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/update-user-stats-settings](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/update-user-stats-settings)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_stats_views/:user_stats_view_id`

### Method

`PUT`

## GetAttendanceUserStatisticsSettings

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserStatisticsSettings(ctx, &lark.GetAttendanceUserStatisticsSettingsReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



查询日度统计或月度统计的统计设置信息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-user-statistics-settings](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-user-statistics-settings)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_stats_views/query`

### Method

`POST`

## GetAttendanceUserDailyShift

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserDailyShift(ctx, &lark.GetAttendanceUserDailyShiftReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



支持查询多个用户的排班情况，查询的时间跨度不能超过 30 天。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/GetScheduledShifts](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/GetScheduledShifts)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_daily_shifts/query`

### Method

`POST`

## GetAttendanceUserTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserTask(ctx, &lark.GetAttendanceUserTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



获取企业内员工的实际打卡结果，包括上班打卡结果和下班打卡结果。

* 如果企业给一个员工设定的班次是上午 9 点和下午 6 点各打一次上下班卡，即使员工在这期间打了多次卡，该接口也只会返回 1 条记录。
* 如果要获取打卡的详细数据，如打卡位置等信息，可使用“获取打卡流水记录”或“批量查询打卡流水记录”的接口。
#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//GetCheckinResults](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//GetCheckinResults)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_tasks/query`

### Method

`POST`

## GetAttendanceUserFlow

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserFlow(ctx, &lark.GetAttendanceUserFlowReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



通过打卡记录 ID 获取用户的打卡流水记录。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//GetCardSwipeHistory](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//GetCardSwipeHistory)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_flows/:user_flow_id`

### Method

`GET`

## BatchGetAttendanceUserFlow

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.BatchGetAttendanceUserFlow(ctx, &lark.BatchGetAttendanceUserFlowReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



批量查询授权内员工的实际打卡记录。例如，企业给一个员工设定的班次是上午 9 点和下午 6 点各打一次上下班卡，但是员工在这期间打了多次卡，该接口会把所有的打卡记录都返回。

如果只需获取打卡结果，而不需要打卡的详细数据，可使用“获取打卡结果”的接口。
#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//BatchQueryCheckinFlowHistory](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//BatchQueryCheckinFlowHistory)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_flows/query`

### Method

`POST`

## BatchCreateAttendanceUserFlow

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.BatchCreateAttendanceUserFlow(ctx, &lark.BatchCreateAttendanceUserFlowReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



导入授权内员工的打卡流水记录。导入后，会根据员工所在的考勤组班次规则，计算最终的打卡状态与结果。

适用于考勤机数据导入等场景。
#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//ImportAttendanceFlowRecords](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//ImportAttendanceFlowRecords)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_flows/batch_create`

### Method

`POST`

## GetAttendanceUserTaskRemedy

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserTaskRemedy(ctx, &lark.GetAttendanceUserTaskRemedyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



获取授权内员工的补卡记录。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//GetUsersRemedyRecords](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//GetUsersRemedyRecords)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_task_remedys/query`

### Method

`POST`

## CreateUpdateAttendanceUserDailyShift

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.CreateUpdateAttendanceUserDailyShift(ctx, &lark.CreateUpdateAttendanceUserDailyShiftReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



班表是用来描述考勤组内人员每天按哪个班次进行上班。目前班表支持按一个整月对一位或多位人员进行排班。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//CreateandEditShifts](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//CreateandEditShifts)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_daily_shifts/batch_create`

### Method

`POST`

## GetAttendanceUserApproval

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserApproval(ctx, &lark.GetAttendanceUserApprovalReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



获取员工在某段时间内的请假、加班、外出和出差四种审批的通过数据。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//RetrieveUserApprovals](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//RetrieveUserApprovals)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_approvals/query`

### Method

`POST`

## CreateAttendanceUserApproval

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.CreateAttendanceUserApproval(ctx, &lark.CreateAttendanceUserApprovalReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



由于部分企业使用的是自己的审批系统，而不是飞书审批系统，因此员工的请假、加班等数据无法流入到飞书考勤系统中，导致员工在请假时间段内依然收到打卡提醒，并且被记为缺卡。

对于这些只使用飞书考勤系统，而未使用飞书审批系统的企业，可以通过考勤开放接口的形式，将三方审批结果数据回写到飞书的考勤系统中。

目前支持加班、请假、出差和外出这四种审批结果的写入。
#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//AddApprovalsInLarkAttendance](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//AddApprovalsInLarkAttendance)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_approvals`

### Method

`POST`

## GetAttendanceUserAllowedRemedy

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserAllowedRemedy(ctx, &lark.GetAttendanceUserAllowedRemedyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



获取用户某天可以补第几次上/下班卡的时间。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-user-allowed-remedys](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-user-allowed-remedys)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_task_remedys/query_user_allowed_remedys`

### Method

`POST`

## InitAttendanceRemedyApproval

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.InitAttendanceRemedyApproval(ctx, &lark.InitAttendanceRemedyApprovalReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



对于只使用飞书考勤系统而未使用飞书审批系统的企业，可以通过该接口，将在三方审批系统中发起的补卡审批数据，写入到飞书的考勤系统中。
#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/notify-remedy-approval-initiation](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/notify-remedy-approval-initiation)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_task_remedys`

### Method

`POST`

## UpdateAttendanceRemedyApproval

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.UpdateAttendanceRemedyApproval(ctx, &lark.UpdateAttendanceRemedyApprovalReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



对于只使用飞书考勤系统而未使用飞书审批系统的企业，可以通过该接口更新写入飞书考勤系统中的三方系统审批状态，例如请假、加班、外出、出差、补卡等审批，状态包括通过、不通过、撤销等。
发起状态的审批才可以被更新为通过、不通过，已经通过的审批才可以被更新为撤销。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/notify-approval-status-update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/notify-approval-status-update)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/approval_infos/process`

### Method

`POST`


# File

## UploadImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.File.UploadImage(ctx, &lark.UploadImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

上传图片接口，可以上传 JPEG、PNG、WEBP 格式图片

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)

### URL

`https://open.feishu.cn/open-apis/im/v1/images`

### Method

`POST`

## DownloadImage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.File.DownloadImage(ctx, &lark.DownloadImageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

下载图片资源，只能下载应用自己上传且图片类型为message的图片

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 只能下载机器人自己上传且图片类型为message的图片，avatar类型暂不支持下载；


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/get)

### URL

`https://open.feishu.cn/open-apis/im/v1/images/:image_key`

### Method

`GET`

## UploadFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.File.UploadFile(ctx, &lark.UploadFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

上传文件，可以上传视频，音频和常见的文件类型

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create)

### URL

`https://open.feishu.cn/open-apis/im/v1/files`

### Method

`POST`

## DownloadFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.File.DownloadFile(ctx, &lark.DownloadFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

下载文件接口，只能下载应用自己上传的文件

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 只能下载机器人自己上传的文件


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/get)

### URL

`https://open.feishu.cn/open-apis/im/v1/files/:file_key`

### Method

`GET`


# OKR

## GetOKRPeriodList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.OKR.GetOKRPeriodList(ctx, &lark.GetOKRPeriodListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取OKR周期列表

使用tenant_access_token需要额外申请权限<md-perm 
href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">以应用身份访问OKR信息</md-perm>




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/period/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/period/list)

### URL

`https://open.feishu.cn/open-apis/okr/v1/periods`

### Method

`GET`

## BatchGetOKR

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.OKR.BatchGetOKR(ctx, &lark.BatchGetOKRReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据OKR id批量获取OKR

使用tenant_access_token需要额外申请权限<md-perm 
href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">以应用身份访问OKR信息</md-perm>




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/okr/batch_get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/okr/batch_get)

### URL

`https://open.feishu.cn/open-apis/okr/v1/okrs/batch_get`

### Method

`GET`

## GetUserOKRList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.OKR.GetUserOKRList(ctx, &lark.GetUserOKRListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据用户的id获取OKR列表

使用tenant_access_token需要额外申请权限<md-perm 
href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">以应用身份访问OKR信息</md-perm>




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/user-okr/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/user-okr/list)

### URL

`https://open.feishu.cn/open-apis/okr/v1/users/:user_id/okrs`

### Method

`GET`


# EHR

## GetEHREmployeeList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.EHR.GetEHREmployeeList(ctx, &lark.GetEHREmployeeListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据员工飞书用户 ID / 员工状态 / 雇员类型等搜索条件 ，批量获取员工花名册字段信息。字段包括「系统标准字段 / system_fields」和「自定义字段 / custom_fields」

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/employee/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/employee/list)

### URL

`https://open.feishu.cn/open-apis/ehr/v1/employees`

### Method

`GET`

## DownloadEHRAttachments

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.EHR.DownloadEHRAttachments(ctx, &lark.DownloadEHRAttachmentsReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



根据文件 token 下载文件。

调用「[批量获取员工花名册信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/employees)」接口的返回值中，「文件」类型的字段 id，即是文件 token

![image.png](//sf1-ttcdn-tos.pstatp.com/obj/open-platform-opendoc/bed391d2a8ce6ed2d5985ea69bf92850_9GY1mnuDXP.png){尝试一下}(url=/api/tools/api_explore/api_explore_config?project=ehr&version=v1&resource=attachment&method=get)












#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/attachment/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/attachment/get)

### URL

`https://open.feishu.cn/open-apis/ehr/v1/attachments/:token`

### Method

`GET`


# Tenant

## GetTenant

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Tenant.GetTenant(ctx, &lark.GetTenantReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取企业名称、企业编号等企业信息

如果ISV应用是预装的并且180天内企业未使用过此应用，则无法通过此接口获取到企业信息







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant/query)

### URL

`https://open.feishu.cn/open-apis/tenant/v2/tenant/query`

### Method

`GET`


# Search

## CreateSearchDataSourceItem

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Search.CreateSearchDataSourceItem(ctx, &lark.CreateSearchDataSourceItemReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

索引一条数据记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/create)

### URL

`https://open.feishu.cn/open-apis/search/v2/data_sources/:data_source_id/items`

### Method

`POST`

## GetSearchDataSourceItem

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Search.GetSearchDataSourceItem(ctx, &lark.GetSearchDataSourceItemReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取单个数据记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/get)

### URL

`https://open.feishu.cn/open-apis/search/v2/data_sources/:data_source_id/items/:item_id`

### Method

`GET`

## DeleteSearchDataSourceItem

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Search.DeleteSearchDataSourceItem(ctx, &lark.DeleteSearchDataSourceItemReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除数据记录

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/delete)

### URL

`https://open.feishu.cn/open-apis/search/v2/data_sources/:data_source_id/items/:item_id`

### Method

`DELETE`

## CreateSearchDataSource

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Search.CreateSearchDataSource(ctx, &lark.CreateSearchDataSourceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

创建一个数据源

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/create)

### URL

`https://open.feishu.cn/open-apis/search/v2/data_sources`

### Method

`POST`

## GetSearchDataSource

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Search.GetSearchDataSource(ctx, &lark.GetSearchDataSourceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取已经创建的数据源

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/get)

### URL

`https://open.feishu.cn/open-apis/search/v2/data_sources/:data_source_id`

### Method

`GET`

## UpdateSearchDataSource

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Search.UpdateSearchDataSource(ctx, &lark.UpdateSearchDataSourceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

更新一个已经存在的数据源

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/patch)

### URL

`https://open.feishu.cn/open-apis/search/v2/data_sources/:data_source_id`

### Method

`PATCH`

## GetSearchDataSourceList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Search.GetSearchDataSourceList(ctx, &lark.GetSearchDataSourceListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取创建的所有数据源信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/list)

### URL

`https://open.feishu.cn/open-apis/search/v2/data_sources`

### Method

`GET`

## DeleteSearchDataSource

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Search.DeleteSearchDataSource(ctx, &lark.DeleteSearchDataSourceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除一个已存在的数据源

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/delete)

### URL

`https://open.feishu.cn/open-apis/search/v2/data_sources/:data_source_id`

### Method

`DELETE`


# Hire

## GetHireJob

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireJob(ctx, &lark.GetHireJobReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据职位 ID 获取职位信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job/get](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job/get)

### URL

`https://open.feishu.cn/open-apis/hire/v1/jobs/:job_id`

### Method

`GET`

## GetHireJobManager

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireJobManager(ctx, &lark.GetHireJobManagerReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据职位 ID 获取职位上的招聘人员信息，如招聘负责人、用人经理

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job-manager/get](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job-manager/get)

### URL

`https://open.feishu.cn/open-apis/hire/v1/jobs/:job_id/managers/:manager_id`

### Method

`GET`

## GetHireTalent

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireTalent(ctx, &lark.GetHireTalentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据人才 ID 获取人才信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/talent/get](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/talent/get)

### URL

`https://open.feishu.cn/open-apis/hire/v1/talents/:talent_id`

### Method

`GET`

## GetHireAttachment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireAttachment(ctx, &lark.GetHireAttachmentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取招聘系统中附件的元信息，比如文件名、创建时间、文件url等

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/attachment/get](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/attachment/get)

### URL

`https://open.feishu.cn/open-apis/hire/v1/attachments/:attachment_id`

### Method

`GET`

## GetHireAttachmentPreview

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireAttachmentPreview(ctx, &lark.GetHireAttachmentPreviewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据附件 ID 获取附件预览信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/attachment/preview](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/attachment/preview)

### URL

`https://open.feishu.cn/open-apis/hire/v1/attachments/:attachment_id/preview`

### Method

`GET`

## GetHireResumeSource

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireResumeSource(ctx, &lark.GetHireResumeSourceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取简历来源列表

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/resume_source/list](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/resume_source/list)

### URL

`https://open.feishu.cn/open-apis/hire/v1/resume_sources`

### Method

`GET`

## CreateHireNote

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.CreateHireNote(ctx, &lark.CreateHireNoteReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

创建备注信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/note/create](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/note/create)

### URL

`https://open.feishu.cn/open-apis/hire/v1/notes`

### Method

`POST`

## UpdateHireNote

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.UpdateHireNote(ctx, &lark.UpdateHireNoteReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据备注 ID 更新备注信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/note/patch](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/note/patch)

### URL

`https://open.feishu.cn/open-apis/hire/v1/notes/:note_id`

### Method

`PATCH`

## GetHireNote

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireNote(ctx, &lark.GetHireNoteReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据备注 ID 获取备注信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/note/get](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/note/get)

### URL

`https://open.feishu.cn/open-apis/hire/v1/notes/:note_id`

### Method

`GET`

## GetHireNoteList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireNoteList(ctx, &lark.GetHireNoteListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取备注列表

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/note/list](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/note/list)

### URL

`https://open.feishu.cn/open-apis/hire/v1/notes`

### Method

`GET`

## GetHireReferralByApplication

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireReferralByApplication(ctx, &lark.GetHireReferralByApplicationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据投递 ID 获取内推信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/referral/get_by_application](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/referral/get_by_application)

### URL

`https://open.feishu.cn/open-apis/hire/v1/referrals/get_by_application`

### Method

`GET`

## GetHireJobProcessList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireJobProcessList(ctx, &lark.GetHireJobProcessListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

获取全部招聘流程信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job_process/list](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job_process/list)

### URL

`https://open.feishu.cn/open-apis/hire/v1/job_processes`

### Method

`GET`

## CreateHireApplication

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.CreateHireApplication(ctx, &lark.CreateHireApplicationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据人才 ID 和职位 ID 创建投递

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/create](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/create)

### URL

`https://open.feishu.cn/open-apis/hire/v1/applications`

### Method

`POST`

## TerminateHireApplication

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.TerminateHireApplication(ctx, &lark.TerminateHireApplicationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据投递 ID 修改投递状态为「已终止」

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/terminate](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/terminate)

### URL

`https://open.feishu.cn/open-apis/hire/v1/applications/:application_id/terminate`

### Method

`POST`

## GetHireApplication

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireApplication(ctx, &lark.GetHireApplicationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据投递 ID 获取单个投递信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/get](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/get)

### URL

`https://open.feishu.cn/open-apis/hire/v1/applications/:application_id`

### Method

`GET`

## GetHireApplicationList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireApplicationList(ctx, &lark.GetHireApplicationListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据限定条件获取投递列表信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/list](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/list)

### URL

`https://open.feishu.cn/open-apis/hire/v1/applications`

### Method

`GET`

## GetHireApplicationInterviewList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireApplicationInterviewList(ctx, &lark.GetHireApplicationInterviewListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据投递 ID 获取面试记录列表

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application-interview/list](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application-interview/list)

### URL

`https://open.feishu.cn/open-apis/hire/v1/applications/:application_id/interviews`

### Method

`GET`

## GetHireOfferByApplication

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireOfferByApplication(ctx, &lark.GetHireOfferByApplicationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据投递 ID 获取 Offer 信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/offer](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/offer)

### URL

`https://open.feishu.cn/open-apis/hire/v1/applications/:application_id/offer`

### Method

`GET`

## GetHireOfferSchema

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireOfferSchema(ctx, &lark.GetHireOfferSchemaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据 Offer 申请表 ID，获取 Offer 申请表的详细信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/offer_schema/get](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/offer_schema/get)

### URL

`https://open.feishu.cn/open-apis/hire/v1/offer_schemas/:offer_schema_id`

### Method

`GET`

## MakeHireTransferOnboardByApplication

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.MakeHireTransferOnboardByApplication(ctx, &lark.MakeHireTransferOnboardByApplicationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据投递 ID 操作候选人入职并创建员工，投递须处于「待入职」阶段

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/transfer_onboard](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/transfer_onboard)

### URL

`https://open.feishu.cn/open-apis/hire/v1/applications/:application_id/transfer_onboard`

### Method

`POST`

## UpdateHireEmployee

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.UpdateHireEmployee(ctx, &lark.UpdateHireEmployeeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

根据员工 ID 更新员工转正、离职状态

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/employee/patch](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/employee/patch)

### URL

`https://open.feishu.cn/open-apis/hire/v1/employees/:employee_id`

### Method

`PATCH`

## GetHireEmployeeByApplication

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireEmployeeByApplication(ctx, &lark.GetHireEmployeeByApplicationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

通过投递 ID 获取入职信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/employee/get_by_application](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/employee/get_by_application)

### URL

`https://open.feishu.cn/open-apis/hire/v1/employees/get_by_application`

### Method

`GET`

## GetHireEmployee

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Hire.GetHireEmployee(ctx, &lark.GetHireEmployeeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

通过员工 ID 获取入职信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/employee/get](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/employee/get)

### URL

`https://open.feishu.cn/open-apis/hire/v1/employees/:employee_id`

### Method

`GET`


# Task

## CreateTaskCollaborator

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.CreateTaskCollaborator(ctx, &lark.CreateTaskCollaboratorReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于新增任务协作者

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/create)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/collaborators`

### Method

`POST`

## GetTaskCollaboratorList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.GetTaskCollaboratorList(ctx, &lark.GetTaskCollaboratorListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于查询任务协作者列表，支持分页，最大值为50

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/list)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/collaborators`

### Method

`GET`

## DeleteTaskCollaborator

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.DeleteTaskCollaborator(ctx, &lark.DeleteTaskCollaboratorReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除任务协作者

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/delete)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/collaborators/:collaborator_id`

### Method

`DELETE`

## CreateTaskFollower

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.CreateTaskFollower(ctx, &lark.CreateTaskFollowerReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于创建任务关注者

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/create)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/followers`

### Method

`POST`

## GetTaskFollowerList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.GetTaskFollowerList(ctx, &lark.GetTaskFollowerListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于查询任务关注者列表，支持分页，最大值为50

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/list)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/followers`

### Method

`GET`

## DeleteTaskFollower

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.DeleteTaskFollower(ctx, &lark.DeleteTaskFollowerReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除任务关注者

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-follower/delete)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/followers/:follower_id`

### Method

`DELETE`

## CreateTaskReminder

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.CreateTaskReminder(ctx, &lark.CreateTaskReminderReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于创建任务的提醒时间

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/create)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/reminders`

### Method

`POST`

## GetTaskReminderList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.GetTaskReminderList(ctx, &lark.GetTaskReminderListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

返回提醒时间列表，支持分页，最大值为50

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/list)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/reminders`

### Method

`GET`

## DeleteTaskReminder

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.DeleteTaskReminder(ctx, &lark.DeleteTaskReminderReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

删除提醒时间，返回结果状态

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/delete)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/reminders/:reminder_id`

### Method

`DELETE`

## CreateTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.CreateTask(ctx, &lark.CreateTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口可以创建一个任务（基本信息），如果需要绑定协作者等需要调用别的资源管理接口。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/create)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks`

### Method

`POST`

## GetTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.GetTask(ctx, &lark.GetTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取任务详情，包括任务标题、描述、时间、来源等信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/get)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id`

### Method

`GET`

## DeleteTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.DeleteTask(ctx, &lark.DeleteTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于删除任务

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/delete)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id`

### Method

`DELETE`

## UpdateTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.UpdateTask(ctx, &lark.UpdateTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于修改任务的标题、描述、时间、来源等相关信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/patch)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id`

### Method

`PATCH`

## CompleteTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.CompleteTask(ctx, &lark.CompleteTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于将接任务状态修改为已完成

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/complete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/complete)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/complete`

### Method

`POST`

## UncompleteTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.UncompleteTask(ctx, &lark.UncompleteTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于取消任务的已完成状态

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/uncomplete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/uncomplete)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/uncomplete`

### Method

`POST`


# ACS

## GetACSAccessRecordPhoto

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.ACS.GetACSAccessRecordPhoto(ctx, &lark.GetACSAccessRecordPhotoReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



用户在门禁考勤机上成功开门或打卡后，智能门禁应用都会生成一条门禁记录，对于使用人脸识别方式进行开门的识别记录，还会有抓拍图。

可以用该接口下载开门时的人脸识别照片{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=acs&version=v1&resource=access_record.access_photo&method=get)












#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/access_record-access_photo/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/access_record-access_photo/get)

### URL

`https://open.feishu.cn/open-apis/acs/v1/access_records/:access_record_id/access_photo`

### Method

`GET`

## GetACSAccessRecordList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.ACS.GetACSAccessRecordList(ctx, &lark.GetACSAccessRecordListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```



用户在门禁考勤机上成功开门或打卡后，智能门禁应用都会生成一条门禁记录。

该接口返回满足查询参数的识别记录{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=acs&version=v1&resource=access_record&method=list)












#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/access_record/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/access_record/list)

### URL

`https://open.feishu.cn/open-apis/acs/v1/access_records`

### Method

`GET`

## GetACSDeviceList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.ACS.GetACSDeviceList(ctx, &lark.GetACSDeviceListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

使用该接口获取租户内所有设备

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/device/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/device/list)

### URL

`https://open.feishu.cn/open-apis/acs/v1/devices`

### Method

`GET`

## GetACSUserFace

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.ACS.GetACSUserFace(ctx, &lark.GetACSUserFaceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

对于已经录入人脸图片的用户，可以使用该接口下载用户人脸图片

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user-face/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user-face/get)

### URL

`https://open.feishu.cn/open-apis/acs/v1/users/:user_id/face`

### Method

`GET`

## UpdateACSUserFace

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.ACS.UpdateACSUserFace(ctx, &lark.UpdateACSUserFaceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

用户需要录入人脸图片才可以使用门禁考勤机。使用该 API 上传门禁用户的人脸图片。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user-face/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user-face/update)

### URL

`https://open.feishu.cn/open-apis/acs/v1/users/:user_id/face`

### Method

`PUT`

## GetACSUser

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.ACS.GetACSUser(ctx, &lark.GetACSUserReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

该接口用于获取智能门禁中单个用户的信息。

只能获取已加入智能门禁权限组的用户







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/get)

### URL

`https://open.feishu.cn/open-apis/acs/v1/users/:user_id`

### Method

`GET`

## UpdateACSUser

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.ACS.UpdateACSUser(ctx, &lark.UpdateACSUserReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

飞书智能门禁在人脸识别成功后会有韦根信号输出，输出用户的卡号。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/patch)

### URL

`https://open.feishu.cn/open-apis/acs/v1/users/:user_id`

### Method

`PATCH`

## GetACSUserList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.ACS.GetACSUserList(ctx, &lark.GetACSUserListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
todo
```

使用该接口获取智能门禁中所有用户信息

只能获取已加入智能门禁权限组的用户







#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/list)

### URL

`https://open.feishu.cn/open-apis/acs/v1/users`

### Method

`GET`


# EventCallback


# AppLink


