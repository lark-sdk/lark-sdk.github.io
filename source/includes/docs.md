
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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.auth.resend_app_ticket(pylark.ResendAppTicketReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

飞书每隔 1 小时会给应用推送一次最新的 app_ticket，应用也可以主动调用此接口，触发飞书进行及时的重新推送。（该接口并不能直接获取app_ticket，而是触发事件推送）

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_ticket_resend](https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_ticket_resend)

### URL

`https://open.feishu.cn/open-apis/auth/v3/app_ticket/resend`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.auth.get_access_token(pylark.GetAccessTokenReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取登录预授权码 code 对应的登录用户身份。

该接口仅适用于通过[第三方网站免登](/ssl:ttdoc/ukTMukTMukTM/uETOwYjLxkDM24SM5AjN)文档中的登录方式获取的预授权码，小程序登录中用户身份的获取，请使用[小程序 code2session 接口](/ssl:ttdoc/uYjL24iN/ukjM04SOyQjL5IDN)








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/authen/access_token](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/authen/access_token)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.auth.refresh_access_token(pylark.RefreshAccessTokenReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

user_access_token 具有一定的时效性，默认最长有效期为7200秒。该接口用于在 user_access_token 过期时用 refresh_token 重新获取 access_token。此时会返回新的 refresh_token，再次刷新 access_token 时需要使用新的 refresh_token。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/authen/refresh_access_token](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/authen/refresh_access_token)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.auth.get_user_info(pylark.GetUserInfoReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过 user_access_token 获取登录用户的信息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/authen/user_info](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/authen/user_info)

### URL

`https://open.feishu.cn/open-apis/authen/v1/user_info`

### Method

`GET`


# Contact

## SearchUserOld

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.SearchUserOld(ctx, &lark.SearchUserOldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.search_user_old(pylark.SearchUserOldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



以用户身份搜索其他用户的信息，无法搜索到外部企业或已离职的用户。<br>

调用该接口需要申请 `搜索用户` 权限。



#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMTM4UjLzEDO14yMxgTN](https://open.feishu.cn/document/ukTMukTMukTM/uMTM4UjLzEDO14yMxgTN)

### URL

`https://open.feishu.cn/open-apis/search/v1/user`

### Method

`GET`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.create_user(pylark.CreateUserReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

使用该接口向通讯录创建一个用户，可以理解为员工入职。创建用户后只返回有数据权限的数据。具体的数据权限的与字段的对应关系请参照[应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

- 新增用户的所有部门必须都在当前应用的通讯录授权范围内才允许新增用户，如果想要在根部门下新增用户，必须要有全员权限。
- 应用商店应用无权限调用此接口。
- 创建用户后，会给用户发送邀请短信/邮件，用户在操作同意后才可访问团队。








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.delete_user(pylark.DeleteUserReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口向通讯录删除一个用户信息，可以理解为员工离职。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

若用户归属部门A、部门B，应用的通讯录权限范围必须包括部门A和部门B才可以删除用户。应用商店应用无权限调用接口。用户可以在删除员工时设置删除员工数据的接收者，如果不设置则由其leader接收，如果该员工没有leader，则会将该员工的数据删除。








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_user(pylark.GetUserReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于获取通讯录中单个用户的信息。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users/:user_id`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_user_list(pylark.GetUserListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

基于部门ID获取，获取部门直属用户列表。

部门ID 必填，根部门的部门ID为0



- 使用 user_access_token 情况下根据个人组织架构的通讯录可见范围进行权限过滤，返回个人组织架构通讯录范围（[登陆企业管理后台进行权限配置](https://www.feishu.cn/admin/security/permission/visibility)）内可见的用户数据。
- 使用tenant_access_token，会根据应用通讯录的范围进行权限过滤。 如果请求的部门ID为0，则校验应用是否具有全员通讯录权限； 如果是非0的部门ID，则会验证应用是否具有该部门的通讯录权限。 无权限返回无权限错误码，有权限则返回对应部门下的直接用户列表。








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/find_by_department](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/find_by_department)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users/find_by_department`

### Method

`GET`

## GetUserListOld

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetUserListOld(ctx, &lark.GetUserListOldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_user_list_old(pylark.GetUserListOldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

基于部门ID获取部门下直属用户列表。

- 使用 user_access_token 情况下根据个人组织架构的通讯录可见范围进行权限过滤，返回个人组织架构通讯录范围（[登陆企业管理后台进行权限配置](https://www.feishu.cn/admin/security/permission/visibility)）内可见的用户数据。
-  tenant_access_token  基于应用通讯录范围进行权限鉴定。由于 department_id 是非必填参数，填与不填存在<b>两种数据权限校验与返回</b>情况：<br>1、请求设置了 department_id 
（根部门为0），会检验所带部门ID是否具有通讯录权限（如果带上 
 department_id=0 会校验是否有全员权限），有则返回部门下直属的成员列表, 否则提示无部门权限的错误码返回。<br>2、请求未带 
  department_id 参数，则会返回权限范围内的独立用户（权限范围直接包含了某用户，则该用户视为权限范围内的独立用户）。








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/list)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.batch_get_user(pylark.BatchGetUserReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



为了更好地提升该接口的安全性，我们对其进行了升级，请尽快迁移至[新版本>>](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get)


该接口用于批量获取用户详细信息。



#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uIzNz4iM3MjLyczM](https://open.feishu.cn/document/ukTMukTMukTM/uIzNz4iM3MjLyczM)

### URL

`https://open.feishu.cn/open-apis/contact/v1/user/batch_get`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.update_user_patch(pylark.UpdateUserPatchReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于更新通讯录中用户的字段，未传递的参数不会更新。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.update_user(pylark.UpdateUserReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于更新通讯录中用户的字段。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

应用需要拥有待更新用户的通讯录授权，如果涉及到用户部门变更，还需要同时拥有所有新部门的通讯录授权。应用商店应用无权限调用此接口。








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/update)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users/:user_id`

### Method

`PUT`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.batch_get_user_by_id(pylark.BatchGetUserByIDReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过该接口，可使用手机号/邮箱获取用户的 ID 信息，具体获取支持的 ID 类型包括 open_id、user_id、union_id，可通过查询参数指定。

如果查询的手机号、邮箱不存在，或者无权限查看对应的用户，则返回的open_id为空。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/batch_get_id](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/batch_get_id)

### URL

`https://open.feishu.cn/open-apis/contact/v3/users/batch_get_id`

### Method

`POST`

## BatchGetUserByIDOld

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.BatchGetUserByIDOld(ctx, &lark.BatchGetUserByIDOldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.batch_get_user_by_id_old(pylark.BatchGetUserByIDOldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



根据用户邮箱或手机号查询用户 open_id 和 user_id，支持批量查询。<br>

调用该接口需要申请 `通过手机号或邮箱获取用户 ID` 权限。<br>只能查询到应用可用性范围内的用户 ID，不在范围内的用户会表现为不存在。



#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUzMyUjL1MjM14SNzITN](https://open.feishu.cn/document/ukTMukTMukTM/uUzMyUjL1MjM14SNzITN)

### URL

`https://open.feishu.cn/open-apis/user/v1/batch_get_id`

### Method

`GET`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.create_department(pylark.CreateDepartmentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于向通讯录中创建部门。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

只可在应用的通讯录权限范围内的部门下创建部门。若需要在根部门下创建子部门，则应用通讯录权限范围需要设置为“全部成员”。应用商店应用无权限调用此接口。








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_department(pylark.GetDepartmentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于向通讯录获取单个部门信息。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

使用tenant_access_token时，应用需要拥有待查询部门的通讯录授权。如果需要获取根部门信息，则需要拥有全员权限。
使用user_access_token时，用户需要有待查询部门的可见性，如果需要获取根部门信息，则要求员工可见所有人。








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_department_list(pylark.GetDepartmentListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过部门ID获取部门的子部门列表。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

部门ID 必填，根部门的部门ID 为0



- 使用 user_access_token 时，返回该用户组织架构可见性范围（[登陆企业管理后台进行权限配置](https://www.feishu.cn/admin/security/permission/visibility)）内的所有可见部门。当进行递归查询时，只筛查最多1000个部门的可见性。

- 使用 
 tenant_access_token 则基于应用的通讯录权限范围进行权限校验与过滤。
如果部门ID为0，会检验应用是否有全员通讯录权限，如果是非0 部门ID，则会校验应用是否有该部门的通讯录权限。无部门权限返回无部门通讯录权限错误码，有权限则返回部门下子部门列表（根据fetch_child决定是否递归）。








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/children](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/children)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments/:department_id/children`

### Method

`GET`

## GetDepartmentListOld

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetDepartmentListOld(ctx, &lark.GetDepartmentListOldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_department_list_old(pylark.GetDepartmentListOldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于获取当前部门子部门列表。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

- 使用 user_access_token 时，返回该用户组织架构可见性范围（[登陆企业管理后台进行权限配置](https://www.feishu.cn/admin/security/permission/visibility)）内的所有可见部门。当进行递归查询时，只筛查最多1000个部门的可见性。

- 使用 
 tenant_access_token 则基于应用的通讯录权限范围进行权限校验与过滤。由于 
 parent_department_id 是非必填参数，填与不填存在<b>两种数据权限校验与返回</b>情况：
<br> <br>1、请求设置了 
 parent_department_id 为A（根部门0），会检验A是否在通讯录权限内，若在( parent_department_id=0 时会校验是否为全员权限），则返回部门下子部门列表（根据fetch_child决定是否递归），否则返回无部门通讯录权限错误码。
<br> <br>2、请求未带 
 parent_department_id 参数，如通讯录范围为全员权限，只返回根部门ID(部门ID为0)，否则返回根据通讯录范围配置的部门ID及子部门(根据 
 fetch_child 决定是否递归)。








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_parent_department(pylark.GetParentDepartmentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用来递归获取部门父部门的信息，并按照由子到父的顺序返回有权限的父部门信息列表。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

使用tenant_access_token时,该接口只返回可见性范围内的父部门信息

例如：A >>B>>C>>D四级部门，通讯录权限只到B，那么查询D部门的parent，会返回B和C两级部门。
使用user_access_token时,该接口只返回对于用户可见的父部门信息








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.search_department(pylark.SearchDepartmentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

搜索部门，用户通过关键词查询可见的部门数据，部门可见性需要管理员在后台配置。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.update_department_patch(pylark.UpdateDepartmentPatchReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于更新通讯录中部门的信息中的任一个字段。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

调用该接口需要具有该部门以及更新操作涉及的部门的通讯录权限。应用商店应用无权限调用此接口。








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.update_department(pylark.UpdateDepartmentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于更新当前部门所有信息。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

- 调用该接口需要具有该部门以及更新操作涉及的部门的通讯录权限。应用商店应用无权限调用此接口。

 - 没有填写的字段会被置为空值（order字段除外）。








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.delete_department(pylark.DeleteDepartmentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于向通讯录中删除部门。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。

应用需要同时拥有待删除部门及其父部门的通讯录授权。应用商店应用无权限调用该接口。








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/delete)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments/:department_id`

### Method

`DELETE`

## UnbindDepartmentChat

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.UnbindDepartmentChat(ctx, &lark.UnbindDepartmentChatReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.unbind_department_chat(pylark.UnbindDepartmentChatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过该接口将部门群转为普通群。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/unbind_department_chat](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/unbind_department_chat)

### URL

`https://open.feishu.cn/open-apis/contact/v3/departments/unbind_department_chat`

### Method

`POST`

## CreateContactGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.CreateContactGroup(ctx, &lark.CreateContactGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.create_contact_group(pylark.CreateContactGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

使用该接口创建用户组，请注意创建用户组时应用的通讯录权限范围需为“全部员工”，否则会创建失败，[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/create)

### URL

`https://open.feishu.cn/open-apis/contact/v3/group`

### Method

`POST`

## UpdateContactGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.UpdateContactGroup(ctx, &lark.UpdateContactGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.update_contact_group(pylark.UpdateContactGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

使用该接口更新用户组信息，请注意更新用户组时应用的通讯录权限范围需为“全部员工”，否则会更新失败。[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/patch)

### URL

`https://open.feishu.cn/open-apis/contact/v3/group/:group_id`

### Method

`PATCH`

## DeleteContactGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.DeleteContactGroup(ctx, &lark.DeleteContactGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.delete_contact_group(pylark.DeleteContactGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过该接口可删除企业中的用户组，请注意删除用户组时应用的通讯录权限范围需为“全部员工”，否则会删除失败，[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/delete)

### URL

`https://open.feishu.cn/open-apis/contact/v3/group/:group_id`

### Method

`DELETE`

## GetContactGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetContactGroup(ctx, &lark.GetContactGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_contact_group(pylark.GetContactGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据用户组 ID 查询某个用户组的基本信息，请确保应用的通讯录权限范围里包括该用户组或者是“全部员工”，[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/get)

### URL

`https://open.feishu.cn/open-apis/contact/v3/group/:group_id`

### Method

`GET`

## GetContactGroupList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetContactGroupList(ctx, &lark.GetContactGroupListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_contact_group_list(pylark.GetContactGroupListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过该接口可查询企业的用户组列表，如果应用的通讯录权限范围是“全部员工”，则可获取企业全部用户组列表。如果应用的通讯录权限范围不是“全部员工”，则仅可获取通讯录权限范围内的用户组。[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/simplelist](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/simplelist)

### URL

`https://open.feishu.cn/open-apis/contact/v3/group/simplelist`

### Method

`GET`

## AddContactGroupMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.AddContactGroupMember(ctx, &lark.AddContactGroupMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.add_contact_group_member(pylark.AddContactGroupMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

向用户组中添加成员(目前成员仅支持用户，未来会支持部门)，如果应用的通讯录权限范围是“全部员工”，则可将任何成员添加到任何用户组。如果应用的通讯录权限范围不是“全部员工”，则仅可将通讯录权限范围中的成员添加到通讯录权限范围的用户组中，[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/add](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/add)

### URL

`https://open.feishu.cn/open-apis/contact/v3/group/:group_id/member/add`

### Method

`POST`

## DeleteContactGroupMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.DeleteContactGroupMember(ctx, &lark.DeleteContactGroupMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.delete_contact_group_member(pylark.DeleteContactGroupMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

从用户组中移除成员 (目前成员仅支持用户，未来会支持部门)，如果应用的通讯录权限范围是“全部员工”，则可将任何成员移出任何用户组。如果应用的通讯录权限范围不是“全部员工”，则仅可将通讯录权限范围中的成员从通讯录权限范围的用户组中移除, [点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/remove](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/remove)

### URL

`https://open.feishu.cn/open-apis/contact/v3/group/:group_id/member/remove`

### Method

`POST`

## GetContactGroupMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetContactGroupMember(ctx, &lark.GetContactGroupMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_contact_group_member(pylark.GetContactGroupMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过该接口可查询某个用户组的成员(目前成员仅支持用户，未来会支持部门)列表，如果应用的通讯录权限范围是“全部员工”，则可查询企业内任何用户组的成员列表。如果应用的通讯录权限范围不是“全部员工”，则仅可查询通讯录权限范围中的用户组的成员列表，[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/simplelist](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/simplelist)

### URL

`https://open.feishu.cn/open-apis/contact/v3/group/:group_id/member/simplelist`

### Method

`GET`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_employee_type_enum_list(pylark.GetEmployeeTypeEnumListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.update_employee_type_enum_patch(pylark.UpdateEmployeeTypeEnumPatchReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.delete_employee_type_enum(pylark.DeleteEmployeeTypeEnumReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除自定义人员类型

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.create_employee_type_enum(pylark.CreateEmployeeTypeEnumReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_contact_custom_attr_list(pylark.GetContactCustomAttrListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取企业自定义的用户字段配置信息

调用该接口前，需要先确认[企业管理员](https://www.feishu.cn/hc/zh-CN/articles/360049067822)在[企业管理后台 - 组织架构 - 成员字段管理](http://www.feishu.cn/admin/contacts/employee-field-new/custom) 自定义字段管理栏开启了“允许开放平台API调用“。

![通讯录.gif](//sf3-cn.feishucdn.com/obj/open-platform-opendoc/544738c94f13ef0b9ebaff53a5133cc7_E9EGMkXyzX.gif)




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/custom_attr/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/custom_attr/list)

### URL

`https://open.feishu.cn/open-apis/contact/v3/custom_attrs`

### Method

`GET`

## CreateContactUnit

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.CreateContactUnit(ctx, &lark.CreateContactUnitReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.create_contact_unit(pylark.CreateContactUnitReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

使用该接口创建单位，需要有更新单位的权限。注意：单位功能属于旗舰版付费功能，企业需开通对应版本才可以创建单位。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/create)

### URL

`https://open.feishu.cn/open-apis/contact/v3/unit`

### Method

`POST`

## UpdateContactUnit

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.UpdateContactUnit(ctx, &lark.UpdateContactUnitReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.update_contact_unit(pylark.UpdateContactUnitReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

调用该接口，需要有更新单位的权限。注意：单位功能属于旗舰版付费功能，企业需开通对应版本才可以修改单位

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/patch)

### URL

`https://open.feishu.cn/open-apis/contact/v3/unit/:unit_id`

### Method

`PATCH`

## DeleteContactUnit

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.DeleteContactUnit(ctx, &lark.DeleteContactUnitReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.delete_contact_unit(pylark.DeleteContactUnitReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

使用该接口删除单位，需要有更新单位的权限。注意：如果单位的单位类型被其它的业务使用，不允许删除。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/delete)

### URL

`https://open.feishu.cn/open-apis/contact/v3/unit/:unit_id`

### Method

`DELETE`

## GetContactUnit

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetContactUnit(ctx, &lark.GetContactUnitReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_contact_unit(pylark.GetContactUnitReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

调用该接口获取单位信息，需有获取单位的权限

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/get)

### URL

`https://open.feishu.cn/open-apis/contact/v3/unit/:unit_id`

### Method

`GET`

## GetContactUnitList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetContactUnitList(ctx, &lark.GetContactUnitListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_contact_unit_list(pylark.GetContactUnitListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过该接口获取企业的单位列表，需获取单位的权限

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/list)

### URL

`https://open.feishu.cn/open-apis/contact/v3/unit`

### Method

`GET`

## BindContactUnitDepartment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.BindContactUnitDepartment(ctx, &lark.BindContactUnitDepartmentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.bind_contact_unit_department(pylark.BindContactUnitDepartmentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过该接口建立部门与单位的绑定关系，需更新单位的权限，需对应部门的通讯录权限。由于单位是旗舰版付费功能，企业需开通相关版本，否则会绑定失败

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/bind_department](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/bind_department)

### URL

`https://open.feishu.cn/open-apis/contact/v3/unit/bind_department`

### Method

`POST`

## UnbindContactUnitDepartment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.UnbindContactUnitDepartment(ctx, &lark.UnbindContactUnitDepartmentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.unbind_contact_unit_department(pylark.UnbindContactUnitDepartmentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过该接口解除部门与单位的绑定关系，需更新单位的权限，需对应部门的通讯录权限。由于单位是旗舰版付费功能，企业需开通相关功能，否则会解绑失败

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/unbind_department](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/unbind_department)

### URL

`https://open.feishu.cn/open-apis/contact/v3/unit/unbind_department`

### Method

`POST`

## GetContactUnitDepartmentList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetContactUnitDepartmentList(ctx, &lark.GetContactUnitDepartmentListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_contact_unit_department_list(pylark.GetContactUnitDepartmentListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过该接口获取单位绑定的部门列表，需具有获取单位的权限

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/list_department](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/list_department)

### URL

`https://open.feishu.cn/open-apis/contact/v3/unit/list_department`

### Method

`GET`

## GetContactScopeList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Contact.GetContactScopeList(ctx, &lark.GetContactScopeListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.contact.get_contact_scope_list(pylark.GetContactScopeListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于获取应用被授权可访问的通讯录范围，包括可访问的部门列表、用户列表和用户组列表。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/scope/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/scope/list)

### URL

`https://open.feishu.cn/open-apis/contact/v3/scopes`

### Method

`GET`


# Message

## SendEphemeralMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.SendEphemeralMessage(ctx, &lark.SendEphemeralMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.send_ephemeral_message(pylark.SendEphemeralMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



用于机器人在群会话中发送仅指定用户可见的消息卡片。<br>

## 使用场景
临时消息卡片多用于群聊中用户与机器人交互的中间态。例如在群聊中用户需要使用待办事项类bot创建一条提醒，bot 发送了可设置提醒日期和提醒内容的一张可交互的消息卡片，此卡片在没有设置为临时卡片的情况下为群内全员可见，即群内可看见该用户与 bot 交互的过程。而设置为临时卡片后，交互过程仅该用户可见，群内其他成员只会看到最终设置完成的提醒卡片。
<br><br>通过临时消息卡片，可以减少消息对群聊中不相关用户的打扰，有效降低群消息的噪声。


需要启用机器人能力；需要机器人在会话群里。



-  仅触发临时卡片的用户自己可见。
- 不支持转发。
- 只能在群聊使用。
- 仅在用户处于在线状态的飞书客户端上可见。
- 临时消息卡片的[呈现能力](/ssl:ttdoc/ukTMukTMukTM/uEjNwUjLxYDM14SM2ATN)、[交互能力](/ssl:ttdoc/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN)与消息卡片一致。


#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uETOyYjLxkjM24SM5IjN](https://open.feishu.cn/document/ukTMukTMukTM/uETOyYjLxkjM24SM5IjN)

### URL

`https://open.feishu.cn/open-apis/ephemeral/v1/send`

### Method

`POST`

## SendUrgentAppMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.SendUrgentAppMessage(ctx, &lark.SendUrgentAppMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.send_urgent_app_message(pylark.SendUrgentAppMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

对指定消息进行应用内加急。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 只能加急机器人自己发送的消息
- 加急时机器人仍需要在会话内




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/urgent_app](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/urgent_app)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id/urgent_app`

### Method

`PATCH`

## SendUrgentSmsMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.SendUrgentSmsMessage(ctx, &lark.SendUrgentSmsMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.send_urgent_sms_message(pylark.SendUrgentSmsMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

对指定消息进行应用内加急与短信加急。

特别说明：
- 通过接口产生的短信加急将消耗企业的加急额度，请慎重调用。
- 通过租户管理后台-费用中心-短信/电话加急 可以查看当前额度。
- 默认接口限流为50 QPS，请谨慎调用。



注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 只能加急机器人自己发送的消息
- 加急时机器人仍需要在会话内




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/urgent_sms](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/urgent_sms)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id/urgent_sms`

### Method

`PATCH`

## SendUrgentPhoneMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.SendUrgentPhoneMessage(ctx, &lark.SendUrgentPhoneMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.send_urgent_phone_message(pylark.SendUrgentPhoneMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

对指定消息进行应用内加急与电话加急

特别说明：
- 通过接口产生的电话加急将消耗企业的加急额度，请慎重调用。
- 通过租户管理后台-费用中心-短信/电话加急 可以查看当前额度。
- 默认接口限流为50 QPS，请谨慎调用。



注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 只能加急机器人自己发送的消息
- 加急时机器人仍需要在会话内




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/urgent_phone](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/urgent_phone)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id/urgent_phone`

### Method

`PATCH`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.send_raw_message(pylark.SendRawMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

给指定用户或者会话发送消息，支持文本、富文本、可交互的[消息卡片](https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN)、群名片、个人名片、图片、视频、音频、文件、表情包。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 给用户发送消息，需要机器人对用户有[可用性](/ssl:ttdoc/home/introduction-to-scope-and-authorization/availability)
- 给群组发送消息，需要机器人在群中
- 该接口不支持给部门成员发消息，请使用 [批量发送消息](/ssl:ttdoc/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM)
- 文本消息请求体最大不能超过150KB
- 卡片及富文本消息请求体最大不能超过30KB
- 消息卡片的 `update_multi`（是否为共享卡片）字段在卡片内容的`config`结构体中设置。详细参考文档[配置卡片属性](/ssl:ttdoc/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.send_raw_message_old(pylark.SendRawMessageOldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```





### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUjNz4SN2MjL1YzM](https://open.feishu.cn/document/ukTMukTMukTM/uUjNz4SN2MjL1YzM)

### URL

`https://open.feishu.cn/open-apis/message/v4/send/`

### Method

`POST`

## BatchSendOldRawMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.BatchSendOldRawMessage(ctx, &lark.BatchSendOldRawMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.batch_send_old_raw_message(pylark.BatchSendOldRawMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



给多个用户或者多个部门发送消息。

**注意事项：**
- 调用该接口需要注意
  - 应用需要启用机器人能力
  - 必须拥有**获取与发送单聊、群组消息**权限，或者**以应用的身份发消息**权限（历史版本）
  - 应用需要拥有批量发送消息权限
  	- 给用户发送需要拥有 **给多个用户批量发消息** 权限
  	- 给部门发送需要拥有 **给一个或多个部门的成员批量发消息** 权限
  - 应用需要拥有对所发送用户或部门的可见性
- 通过该接口发送的消息 **不支持更新以及回复等操作**
- 只能发送给用户，无法发送给群组
- 异步接口，会有一定延迟，每个应用待发送的消息按顺序处理，请合理安排批量发送范围和顺序。发送消息给单个用户的场景请使用[发送消息](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create)接口
- 单个应用每天通过该接口发送的总消息条数不超过50万

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM](https://open.feishu.cn/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM)

### URL

`https://open.feishu.cn/open-apis/message/v4/batch_send/`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.reply_raw_message(pylark.ReplyRawMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

回复指定消息，支持文本、富文本、卡片、群名片、个人名片、图片、视频、文件等多种消息类型。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 回复私聊消息，需要机器人对用户有可用性
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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.delete_message(pylark.DeleteMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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

## BatchDeleteMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.BatchDeleteMessage(ctx, &lark.BatchDeleteMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.batch_delete_message(pylark.BatchDeleteMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

批量撤回消息

注意事项：
- 只能撤回通过[批量发送消息](/ssl:ttdoc/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM)接口产生的消息，单条消息的撤回请使用[撤回消息](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/delete)接口
- 路径参数**batch_message_id**为[批量发送消息](/ssl:ttdoc/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM)接口返回值中的**message_id**字段，用于标识一次批量发送消息请求，格式为：**bm-xxx**
- 一次调用涉及大量消息，所以为异步接口，会有一定延迟。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/delete)

### URL

`https://open.feishu.cn/open-apis/im/v1/batch_messages/:batch_message_id`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.update_message(pylark.UpdateMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新应用已发送的消息卡片内容。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 当前仅支持更新 **卡片消息**
- **不支持更新批量消息**
- 只支持对所有人都更新的[「共享卡片」](ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN)，也即需要在卡片的`config`属性中，显式声明`"update_multi":true`。<br>如果你只想更新特定人的消息卡片，必须要用户在卡片操作交互后触发，开发文档参考[「独享卡片」](/ssl:ttdoc/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN#49904b71)








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/patch)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id`

### Method

`PATCH`

## UpdateMessageDelay

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.UpdateMessageDelay(ctx, &lark.UpdateMessageDelayReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.update_message_delay(pylark.UpdateMessageDelayReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```





### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMDO1YjLzgTN24yM4UjN](https://open.feishu.cn/document/ukTMukTMukTM/uMDO1YjLzgTN24yM4UjN)

### URL

`https://open.feishu.cn/open-apis/interactive/v1/card/update`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.get_message_read_user_list(pylark.GetMessageReadUserListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查询消息的已读信息。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 只能查询机器人自己发送，且发送时间不超过7天的消息
- 查询消息已读信息时机器人仍需要在会话内
- 本接口不支持查询批量消息




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/read_users](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/read_users)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id/read_users`

### Method

`GET`

## GetBatchSentMessageReadUser

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.GetBatchSentMessageReadUser(ctx, &lark.GetBatchSentMessageReadUserReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.get_batch_sent_message_read_user(pylark.GetBatchSentMessageReadUserReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查询批量消息推送和阅读人数

注意事项：
- 只能查询通过[批量发送消息](/ssl:ttdoc/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM)接口产生的消息
- 该接口返回的数据为查询时刻的快照数据。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/read_user](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/read_user)

### URL

`https://open.feishu.cn/open-apis/im/v1/batch_messages/:batch_message_id/read_user`

### Method

`GET`

## GetBatchSentMessageProgress

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.GetBatchSentMessageProgress(ctx, &lark.GetBatchSentMessageProgressReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.get_batch_sent_message_progress(pylark.GetBatchSentMessageProgressReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查询批量消息整体进度

注意事项:
* 该接口是[查询批量消息推送和阅读人数](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/read_user)接口的加强版
* 该接口返回的数据为查询时刻的快照数据




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/get_progress](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/batch_message/get_progress)

### URL

`https://open.feishu.cn/open-apis/im/v1/batch_messages/:batch_message_id/get_progress`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.get_message_list(pylark.GetMessageListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取会话（包括单聊、群组）的历史消息（聊天记录）。

接口级别权限默认只能获取单聊消息，如果需要获取群组消息，应用还必须拥有 ***获取群组中所有消息*** 权限



- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 获取消息时，机器人必须在群组中




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.get_message_file(pylark.GetMessageFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取消息中的资源文件，包括音频，视频，图片和文件，**暂不支持表情包资源下载**。当前仅支持 100M 以内的资源文件的下载。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 机器人和消息需要在同一会话中
- 请求的 file_key 和 message_id 需要匹配
- 暂不支持获取合并转发消息中的子消息的资源文件
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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.get_message(pylark.GetMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.delete_ephemeral_message(pylark.DeleteEphemeralMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



在群会话中删除指定用户可见的临时消息卡片<br>
临时卡片消息可以通过该接口进行显式删除，临时卡片消息删除后将不会在该设备上留下任何痕迹。

**权限说明** ：需要启用机器人能力；需要机器人在会话群里

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITOyYjLykjM24iM5IjN](https://open.feishu.cn/document/ukTMukTMukTM/uITOyYjLykjM24iM5IjN)

### URL

`https://open.feishu.cn/open-apis/ephemeral/v1/delete`

### Method

`POST`

## CreateMessageReaction

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.CreateMessageReaction(ctx, &lark.CreateMessageReactionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.create_message_reaction(pylark.CreateMessageReactionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

给指定消息添加指定类型的表情回复（reaction即表情回复，本说明文档统一用“reaction”代称）。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 待添加reaction的消息要真实存在，不能被撤回
- 给消息添加reaction，需要reaction的发送方（机器人或者用户）在消息所在的会话内




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/create)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id/reactions`

### Method

`POST`

## GetMessageReactionList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.GetMessageReactionList(ctx, &lark.GetMessageReactionListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.get_message_reaction_list(pylark.GetMessageReactionListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取指定消息的特定类型表情回复列表（reaction即表情回复，本说明文档统一用“reaction”代称）。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 待获取reaction信息的消息要真实存在，不能被撤回
- 获取消息的reaction，需要request的授权主体（机器人或者用户）在消息所在的会话内




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/list)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id/reactions`

### Method

`GET`

## DeleteMessageReaction

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Message.DeleteMessageReaction(ctx, &lark.DeleteMessageReactionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.message.delete_message_reaction(pylark.DeleteMessageReactionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除指定消息的表情回复（reaction即表情回复，本说明文档统一用“reaction”代称）。

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)  
- 只能删除真实存在的reaction，并且删除reaction请求的操作者必须是reaction的原始添加者




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/delete)

### URL

`https://open.feishu.cn/open-apis/im/v1/messages/:message_id/reactions/:reaction_id`

### Method

`DELETE`


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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.create_chat(pylark.CreateChatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

创建群并设置群头像、群名、群描述等。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 本接口只支持创建群，如果需要拉用户或者机器人入群参考 [将用户或机器人拉入群聊](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create)接口
- 每次请求，最多拉 50 个用户或者 5 个机器人，并且群组最多容纳 15 个机器人
 - 拉机器人入群请使用 ==app_id==




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.get_chat(pylark.GetChatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取群名称、群描述、群头像、群主 ID 等群基本信息。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
 - 机器人或授权用户必须在群里（否则只会返回群名称、群头像等基本信息）




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id`

### Method

`GET`

## GetChatOld

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.GetChatOld(ctx, &lark.GetChatOldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.get_chat_old(pylark.GetChatOldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



为了更好地提升该接口的安全性，我们对其进行了升级，请尽快迁移至[新版本>>](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get)


获取群名称、群主 ID、成员列表 ID 等群基本信息。  

- 需要启用机器人能力；机器人必须在群里


#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uMTO5QjLzkTO04yM5kDN](https://open.feishu.cn/document/ukTMukTMukTM/uMTO5QjLzkTO04yM5kDN)

### URL

`https://open.feishu.cn/open-apis/chat/v4/info`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.update_chat(pylark.UpdateChatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新群头像、群名称、群描述、群配置、转让群主等。

注意事项：
- 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 若群未开启 ==仅群主和群管理员可编辑群信息== 配置：
 	- 群主/群管理员 或 创建群组且具备==更新应用所创建群的群信息==权限的机器人，可更新所有信息
 	- 不满足上述条件的群成员或机器人，仅可更新群头像、群名称、群描述、群国际化名称信息 
- 若群开启了==仅群主和群管理员可编辑群信息==配置：
 	- 群主/群管理员 或 创建群组且具备==更新应用所创建群的群信息==权限的机器人，可更新所有信息
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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.delete_chat(pylark.DeleteChatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

解散群组

注意事项：
- 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 如果使用tenant_access_token，需要机器人是群的创建者且具备==更新应用所创建群的群信息==权限才可解散群
- 如果使用user_access_token，需要对应的用户是群主才可解散群


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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.get_chat_list_of_self(pylark.GetChatListOfSelfReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取用户或者机器人所在群列表。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 查询参数  **user_id_type** 用于控制响应体中 owner_id 的类型，如果是获取机器人所在群列表该值可以不填
- 请注意区分本接口和[获取群信息](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get)的请求 URL




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.search_chat(pylark.SearchChatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.get_chat_member_list(pylark.GetChatMemberListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

如果用户在群中，则返回该群的成员列表。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
 - 该接口不会返回群内的机器人成员
 - 由于返回的群成员列表会过滤掉机器人成员，因此返回的群成员个数可能会小于指定的page_size
 - 如果有同一时间加入群的群成员，会一次性返回，这会导致返回的群成员个数可能会大于指定的page_size




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.is_in_chat(pylark.IsInChatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据使用的access_token判断对应的用户或者机器人是否在群里。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/is_in_chat](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/is_in_chat)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members/is_in_chat`

### Method

`GET`

## CreateChatManager

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.CreateChatManager(ctx, &lark.CreateChatManagerReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.create_chat_manager(pylark.CreateChatManagerReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

将用户或机器人指定为群管理员。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 仅有群主可以指定群管理员
- 对于普通群，最多指定 10 个管理员
- 对于超大群，最多指定 20 个管理员
- 每次请求最多指定 50 个用户或者 5 个机器人
- 指定机器人类型的管理员请使用 ==app_id==




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-managers/add_managers](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-managers/add_managers)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/managers/add_managers`

### Method

`POST`

## DeleteChatManager

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.DeleteChatManager(ctx, &lark.DeleteChatManagerReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.delete_chat_manager(pylark.DeleteChatManagerReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除指定的群管理员（用户或机器人）

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 仅有群主可以删除群管理员
- 每次请求最多指定 50 个用户或者 5 个机器人
- 删除机器人类型的管理员请使用 ==app_id==




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-managers/delete_managers](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-managers/delete_managers)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/managers/delete_managers`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.add_chat_member(pylark.AddChatMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

将用户或机器人拉入群聊。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
 - 如需拉用户进群，需要机器人对用户有可见性
 - 在开启 ==仅群主和群管理员可添加群成员== 的设置时，仅有群主/管理员 或 创建群组且具备 ==更新应用所创建群的群信息== 权限的机器人，可以拉用户或者机器人进群
 - 在未开启 ==仅群主和群管理员可添加群成员== 的设置时，所有群成员都可以拉用户或机器人进群
 - 每次请求，最多拉50个用户或者5个机器人，并且群组最多容纳15个机器人
 - 拉机器人入群请使用 ==app_id==




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.delete_chat_member(pylark.DeleteChatMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

将用户或机器人移出群聊。

注意事项：
 - 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 用户或机器人在任何条件下均可移除自己出群（即主动退群）
- 仅有群主/管理员 或 创建群组并且具备 ==更新应用所创建群的群信息== 权限的机器人，可以移除其他用户或者机器人
 - 每次请求，最多移除50个用户或者5个机器人
- 移除机器人请使用 ==app_id==




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.join_chat(pylark.JoinChatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.get_chat_announcement(pylark.GetChatAnnouncementReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.update_chat_announcement(pylark.UpdateChatAnnouncementReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新会话中的群公告信息，更新公告信息的格式和更新[云文档](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)格式相同。

注意事项：
- 应用需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 若群开启了 ==仅群主和群管理员可编辑群信息== 配置，群主/群管理员 或 创建群组且具备 ==更新应用所创建群的群信息== 权限的机器人，可更新群公告
- 若群未开启 ==仅群主和群管理员可编辑群信息== 配置，所有成员可以更新群公告




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/patch)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/announcement`

### Method

`PATCH`

## GetChatModeration

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.GetChatModeration(ctx, &lark.GetChatModerationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.get_chat_moderation(pylark.GetChatModerationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取群发言模式、可发言用户名单等

注意事项：
 - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
 - 机器人 或 授权用户 必须在群里


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-moderation/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-moderation/get)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/moderation`

### Method

`GET`

## UpdateChatModeration

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Chat.UpdateChatModeration(ctx, &lark.UpdateChatModerationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.chat.update_chat_moderation(pylark.UpdateChatModerationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新群组的发言权限设置，可设置为全员可发言、仅管理员可发言  或 指定用户可发言。

注意事项：
 - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
- 若以用户授权调用接口，**当授权用户是群主**时，可更新群发言权限
- 若以租户授权调用接口(即以机器人身份调用接口)，当**机器人是群主** 或者 **机器人是创建群组、具备==更新应用所创建群的群信息==权限且仍在群内**时，可更新群发言权限


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-moderation/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-moderation/update)

### URL

`https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/moderation`

### Method

`PUT`


# Bot

## GetBotInfo

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bot.GetBotInfo(ctx, &lark.GetBotInfoReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bot.get_bot_info(pylark.GetBotInfoReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取机器人的基本信息。

需要启用机器人能力（前往[开发者后台](https://open.feishu.cn/app) - 选择你要获取信息的应用 - 导航栏点击应用功能 - 机器人，开启机器人能力并发布后即可。）








#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uAjMxEjLwITMx4CMyETM](https://open.feishu.cn/document/ukTMukTMukTM/uAjMxEjLwITMx4CMyETM)

### URL

`https://open.feishu.cn/open-apis/bot/v3/info`

### Method

`GET`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bot.add_bot_to_chat(pylark.AddBotToChatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



为了更好地提升该接口的安全性，我们对其进行了升级，请尽快迁移至[新版本>>](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create)


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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.create_calendar_acl(pylark.CreateCalendarACLReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.delete_calendar_acl(pylark.DeleteCalendarACLReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.get_calendar_acl_list(pylark.GetCalendarACLListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.subscribe_calendar_acl(pylark.SubscribeCalendarACLReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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

## GetPrimaryCalendar

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GetPrimaryCalendar(ctx, &lark.GetPrimaryCalendarReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.get_primary_calendar(pylark.GetPrimaryCalendarReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取当前身份的主日历信息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/primary](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/primary)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/calendars/primary`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.create_calendar(pylark.CreateCalendarReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.delete_calendar(pylark.DeleteCalendarReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.get_calendar(pylark.GetCalendarReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.get_calendar_list(pylark.GetCalendarListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.update_calendar(pylark.UpdateCalendarReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.search_calendar(pylark.SearchCalendarReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.subscribe_calendar(pylark.SubscribeCalendarReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于以当前身份（应用 / 用户）订阅某个日历。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&method=subscribe)










- 仅可订阅类型为 primary 或 shared 的公开日历。
- 可订阅日历数量上限为1000。




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.unsubscribe_calendar(pylark.UnsubscribeCalendarReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.subscribe_calendar_change_event(pylark.SubscribeCalendarChangeEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.create_calendar_event(pylark.CreateCalendarEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.delete_calendar_event(pylark.DeleteCalendarEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.get_calendar_event(pylark.GetCalendarEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于以当前身份（应用 / 用户）获取日历上的一个日程。

- 当前身份必须对日历有reader、writer或owner权限才会返回日程详细信息（调用[获取日历](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口，role字段可查看权限）。
- [例外日程](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/introduction#71c5ec78)可通过event_id的非0时间戳后缀，来获取修改的重复性日程的哪一天日程的时间信息。




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.get_calendar_event_list(pylark.GetCalendarEventListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于以当前身份（应用 / 用户）获取日历下的日程列表。

- 当前身份必须对日历有reader、writer或owner权限才会返回日程详细信息（调用[获取日历](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口，role字段可查看权限）。

- 仅支持primary、shared和resource类型的日历获取日程列表。

- 调用时首先使用 page_token 分页拉取存量数据，之后使用 sync_token 增量同步变更数据。

- 为了确保调用方日程同步数据的一致性，在使用sync_token时，不能同时使用start_time和end_time，否则可能造成日程数据缺失。




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.update_calendar_event(pylark.UpdateCalendarEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.search_calendar_event(pylark.SearchCalendarEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于以用户身份搜索某日历下的相关日程。

身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.event&method=search)










当前身份必须对日历有reader、writer或owner权限（调用[获取日历](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口，role字段可查看权限）。




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.subscribe_calendar_event(pylark.SubscribeCalendarEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于以用户身份订阅指定日历下的日程变更事件。

当前身份必须对日历有reader、writer或owner权限（调用[获取日历](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口，role字段可查看权限）。




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.create_calendar_event_attendee(pylark.CreateCalendarEventAttendeeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

批量给日程添加参与人。

- 当前身份需要有日历的 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。

- 当前身份需要是日程的组织者，或日程设置了「参与人可邀请其它参与人」权限。

- 新添加的日程参与人必须与日程组织者在同一个企业内。

- 使用该接口添加会议室后，会议室会进入异步的预约流程，请求结束不代表会议室预约成功，需后续再查询预约状态。

- 每个日程最多只能有 3000 名参与人。

- 开启管理员能力后预约会议室可不受会议室预约范围的限制（当前不支持用管理员身份给其他人的日程预约会议室）








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.get_calendar_event_attendee_list(pylark.GetCalendarEventAttendeeListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取日程的参与人列表，若参与者列表中有群组，请使用 [获取参与人群成员列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee-chat_member/list) 。

- 当前身份必须对日历有reader、writer或owner权限（调用[获取日历](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口，role字段可查看权限）。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.delete_calendar_event_attendee(pylark.DeleteCalendarEventAttendeeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.get_calendar_event_attendee_chat_member_list(pylark.GetCalendarEventAttendeeChatMemberListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.get_calendar_free_busy_list(pylark.GetCalendarFreeBusyListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.create_calendar_timeoff_event(pylark.CreateCalendarTimeoffEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.delete_calendar_timeoff_event(pylark.DeleteCalendarTimeoffEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.generate_caldav_conf(pylark.GenerateCaldavConfReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

用于为当前用户生成一个CalDAV账号密码，用于将飞书日历信息同步到本地设备日历。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/setting/generate_caldav_conf](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/setting/generate_caldav_conf)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/settings/generate_caldav_conf`

### Method

`POST`

## CreateCalendarExchangeBinding

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.CreateCalendarExchangeBinding(ctx, &lark.CreateCalendarExchangeBindingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.create_calendar_exchange_binding(pylark.CreateCalendarExchangeBindingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

本接口将Exchange账户绑定到飞书账户，进而支持Exchange日历的导入

操作用户需要是企业超级管理员




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/create)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/exchange_bindings`

### Method

`POST`

## GetCalendarExchangeBinding

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.GetCalendarExchangeBinding(ctx, &lark.GetCalendarExchangeBindingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.get_calendar_exchange_binding(pylark.GetCalendarExchangeBindingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

本接口获取Exchange账户的绑定状态，包括exchange日历是否同步完成。

操作用户需要是企业超级管理员




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/get)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id`

### Method

`GET`

## DeleteCalendarExchangeBinding

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Calendar.DeleteCalendarExchangeBinding(ctx, &lark.DeleteCalendarExchangeBindingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.calendar.delete_calendar_exchange_binding(pylark.DeleteCalendarExchangeBindingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

本接口解除Exchange账户和飞书账户的绑定关系，Exchange账户解除绑定后才能绑定其他飞书账户

操作用户需要是企业超级管理员




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/delete)

### URL

`https://open.feishu.cn/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id`

### Method

`DELETE`


# Drive

## SubscribeDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.SubscribeDriveFile(ctx, &lark.SubscribeDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.subscribe_drive_file(pylark.SubscribeDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于根据文件token和文件类型订阅 Doc/Docx/Sheet 的事件。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/subscribe](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/subscribe)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/subscribe`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.search_drive_file(pylark.SearchDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_file_meta(pylark.GetDriveFileMetaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_drive_file(pylark.CreateDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于根据 folderToken 创建 Doc、 Sheet 或 Bitable 。

若没有特定的文件夹用于承载创建的文档，可以先调用「获取文件夹元信息」文档中的「获取 root folder (我的空间) meta」接口，获得我的空间的 token，然后再使用此接口。创建的文档将会在「我的空间」的「归我所有」列表里。


该接口不支持并发创建，且调用频率上限为 5QPS 且 10000次/天

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQTNzUjL0UzM14CN1MTN](https://open.feishu.cn/document/ukTMukTMukTM/uQTNzUjL0UzM14CN1MTN)

### URL

`https://open.feishu.cn/open-apis/drive/explorer/v2/file/:folderToken`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_drive_file(pylark.DeleteDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除在云空间内的文件

文档只能被文档所有者删除(使用tenant_access_token 前，请确保该应用是文档的所有者，否则会报无权限错误)，文档被删除后将会放到回收站里



该接口不支持并发调用，且调用频率上限为5QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/delete)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_drive_sheet_file(pylark.DeleteDriveSheetFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于根据 spreadsheetToken 删除对应的 sheet 文档。

为了更好地提升该接口的安全性，我们对其进行了升级，请尽快迁移至
  [新版本>>](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/delete)

</md-alert>

<md-alert type="warn">

文档只能被文档所有者删除，文档被删除后将会放到回收站里


该接口不支持并发调用，且调用频率上限为5QPS

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUTNzUjL1UzM14SN1MTN/delete-sheet](https://open.feishu.cn/document/ukTMukTMukTM/uUTNzUjL1UzM14SN1MTN/delete-sheet)

### URL

`https://open.feishu.cn/open-apis/drive/explorer/v2/file/spreadsheets/:spreadsheetToken`

### Method

`DELETE`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_folder_meta(pylark.GetDriveFolderMetaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_root_folder_meta(pylark.GetDriveRootFolderMetaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_folder_children(pylark.GetDriveFolderChildrenReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于根据 folderToken 获取该文件夹的文档清单，如 doc、sheet、file、bitable、docx、folder。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_file_statistics(pylark.GetDriveFileStatisticsReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.download_drive_file(pylark.DownloadDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.copy_drive_file(pylark.CopyDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



将文件复制到用户云空间的其他文件夹中。不支持复制文件夹。

如果目标文件夹是我的空间，则复制的文件会在「**我的空间**」的「**归我所有**」列表里。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=drive&version=v1&resource=file&method=copy)






该接口不支持并发拷贝多个文件，且调用频率上限为 5QPS 且 10000次/天








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/copy](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/copy)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/copy`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_drive_folder(pylark.CreateDriveFolderReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

在用户云空间的指定文件夹中创建一个新的空文件夹。

该接口不支持并发创建，且调用频率上限为 5QPS 以及 10000次/天








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/create_folder](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/create_folder)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/create_folder`

### Method

`POST`

## MoveDriveFile

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.MoveDriveFile(ctx, &lark.MoveDriveFileReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.move_drive_file(pylark.MoveDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

将文件或者文件夹移动到用户云空间的其他位置。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/move](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/move)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/move`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.upload_drive_file(pylark.UploadDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.prepare_upload_drive_file(pylark.PrepareUploadDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

发送初始化请求获取上传事务ID和分块策略，目前是以4MB大小进行定长分片。

你在24小时内可保存上传事务ID和上传进度，以便可以恢复上传



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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.part_upload_drive_file(pylark.PartUploadDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.finish_upload_drive_file(pylark.FinishUploadDriveFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.download_drive_media(pylark.DownloadDriveMediaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.upload_drive_media(pylark.UploadDriveMediaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.prepare_upload_drive_media(pylark.PrepareUploadDriveMediaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.part_upload_drive_media(pylark.PartUploadDriveMediaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.finish_upload_drive_media(pylark.FinishUploadDriveMediaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_drive_member_permission_old(pylark.CreateDriveMemberPermissionOldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.transfer_drive_member_permission(pylark.TransferDriveMemberPermissionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_member_permission_list(pylark.GetDriveMemberPermissionListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_drive_member_permission(pylark.CreateDriveMemberPermissionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_drive_member_permission(pylark.DeleteDriveMemberPermissionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于根据 filetoken 移除文档协作者的权限。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/delete)

### URL

`https://open.feishu.cn/open-apis/drive/v1/permissions/:token/members/:member_id`

### Method

`DELETE`

## DeleteDriveMemberPermissionOld

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteDriveMemberPermissionOld(ctx, &lark.DeleteDriveMemberPermissionOldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_drive_member_permission_old(pylark.DeleteDriveMemberPermissionOldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于根据 filetoken 移除文档协作者的权限。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYTN3UjL2UzN14iN1cTN](https://open.feishu.cn/document/ukTMukTMukTM/uYTN3UjL2UzN14iN1cTN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/member/delete`

### Method

`POST`

## UpdateDriveMemberPermissionOld

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateDriveMemberPermissionOld(ctx, &lark.UpdateDriveMemberPermissionOldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_drive_member_permission_old(pylark.UpdateDriveMemberPermissionOldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于根据 filetoken 更新文档协作者的权限。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ucTN3UjL3UzN14yN1cTN](https://open.feishu.cn/document/ukTMukTMukTM/ucTN3UjL3UzN14yN1cTN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/member/update`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_drive_member_permission(pylark.UpdateDriveMemberPermissionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于根据 filetoken 更新文档协作者的权限。

该接口要求文档协作者已存在，如还未对文档协作者授权请先调用[「增加权限」 ](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/create)接口进行授权。








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/update)

### URL

`https://open.feishu.cn/open-apis/drive/v1/permissions/:token/members/:member_id`

### Method

`PUT`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.check_drive_member_permission(pylark.CheckDriveMemberPermissionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_drive_public_permission_v1_old(pylark.UpdateDrivePublicPermissionV1OldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_drive_public_permission_v2_old(pylark.UpdateDrivePublicPermissionV2OldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_public_permission_v2(pylark.GetDrivePublicPermissionV2Req(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于根据 filetoken 获取云文档的权限设置。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITM3YjLyEzN24iMxcjN](https://open.feishu.cn/document/ukTMukTMukTM/uITM3YjLyEzN24iMxcjN)

### URL

`https://open.feishu.cn/open-apis/drive/permission/v2/public/`

### Method

`POST`

## GetDrivePublicPermission

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDrivePublicPermission(ctx, &lark.GetDrivePublicPermissionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_public_permission(pylark.GetDrivePublicPermissionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于根据 filetoken 获取云文档的权限设置。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-public/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-public/get)

### URL

`https://open.feishu.cn/open-apis/drive/v1/permissions/:token/public`

### Method

`GET`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_drive_public_permission(pylark.UpdateDrivePublicPermissionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于根据 filetoken 更新云文档的权限设置。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.batch_get_drive_media_tmp_download_url(pylark.BatchGetDriveMediaTmpDownloadURLReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_comment_list(pylark.GetDriveCommentListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过分页方式获取云文档中的全文评论列表。

注意：该接口仅可获取在线文档的全文评论，不支持获取局部评论或者在线表格中的评论。








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_comment(pylark.GetDriveCommentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_drive_comment(pylark.CreateDriveCommentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

往云文档添加一条全局评论。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_drive_comment(pylark.UpdateDriveCommentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_drive_comment(pylark.DeleteDriveCommentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_drive_comment_patch(pylark.UpdateDriveCommentPatchReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

解决或恢复云文档中的评论。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/patch)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments/:comment_id`

### Method

`PATCH`

## CreateDriveFileSubscription

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateDriveFileSubscription(ctx, &lark.CreateDriveFileSubscriptionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_drive_file_subscription(pylark.CreateDriveFileSubscriptionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

订阅文档中的变更事件，当前支持文档评论订阅，订阅后文档评论更新会有“云文档助手”推送给订阅的用户

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-subscription/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-subscription/create)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/subscriptions`

### Method

`POST`

## GetDriveFileSubscription

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetDriveFileSubscription(ctx, &lark.GetDriveFileSubscriptionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_file_subscription(pylark.GetDriveFileSubscriptionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据订阅ID获取该订阅的状态

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-subscription/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-subscription/get)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/subscriptions/:subscription_id`

### Method

`GET`

## UpdateDriveFileSubscription

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateDriveFileSubscription(ctx, &lark.UpdateDriveFileSubscriptionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_drive_file_subscription(pylark.UpdateDriveFileSubscriptionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据订阅ID更新订阅状态

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-subscription/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-subscription/patch)

### URL

`https://open.feishu.cn/open-apis/drive/v1/files/:file_token/subscriptions/:subscription_id`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_drive_doc(pylark.CreateDriveDocReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



在使用此接口前，请仔细阅读[文档概述](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-doc-overview)和[准备接入文档 API](/ssl:ttdoc/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)了解文档调用的规则和约束，确保你的文档数据不会丢失或出错。  
文档数据结构定义可参考：[文档数据结构概述](/ssl:ttdoc/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN) 

该接口用于创建并初始化文档。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_doc_content(pylark.GetDriveDocContentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



在使用此接口前，请仔细阅读[文档概述](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-doc-overview)和[准备接入文档 API](/ssl:ttdoc/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)了解文档调用的规则和约束，确保你的文档数据不会丢失或出错。  
文档数据结构定义可参考：[文档数据结构概述](/ssl:ttdoc/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN) 
该接口用于获取结构化的文档内容。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_doc_raw_content(pylark.GetDriveDocRawContentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_doc_meta(pylark.GetDriveDocMetaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_sheet(pylark.CreateSheetReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_sheet_meta(pylark.GetSheetMetaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_sheet_property(pylark.UpdateSheetPropertyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.batch_update_sheet(pylark.BatchUpdateSheetReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于根据 spreadsheetToken 操作表格，如增加工作表，复制工作表、删除工作表。
该接口和 [更新工作表属性](/ssl:ttdoc/ukTMukTMukTM/ugjMzUjL4IzM14COyMTN) 的请求地址相同，但参数不同，调用前请仔细阅读文档。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.import_sheet(pylark.ImportSheetReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



>  为了更好地提升该接口的安全性，我们对其进行了升级，请尽快迁移至[新版本](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/import-user-guide)


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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_drive_import_task(pylark.CreateDriveImportTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_drive_import_task(pylark.GetDriveImportTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.move_sheet_dimension(pylark.MoveSheetDimensionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.prepend_sheet_value(pylark.PrependSheetValueReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于根据 spreadsheetToken 和 range 向范围之前增加相应数据的行和相应的数据，相当于数组的插入操作；单次写入不超过5000行，100列，每个格子不超过5万字符。
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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.append_sheet_value(pylark.AppendSheetValueReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于根据 spreadsheetToken 和 range 遇到空行则进行覆盖追加或新增行追加数据。 空行：默认该行第一个格子是空，则认为是空行；单次写入不超过5000行，100列，每个格子不超过5万字符。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.insert_sheet_dimension_range(pylark.InsertSheetDimensionRangeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于根据 spreadsheetToken 和维度信息 插入空行/列。  
如 startIndex=3， endIndex=7，则从第 4 行开始开始插入行列，一直到第 7 行，共插入 4 行；单次操作不超过5000行或列。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.add_sheet_dimension_range(pylark.AddSheetDimensionRangeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_sheet_dimension_range(pylark.UpdateSheetDimensionRangeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_sheet_dimension_range(pylark.DeleteSheetDimensionRangeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_sheet_value(pylark.GetSheetValueReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.batch_get_sheet_value(pylark.BatchGetSheetValueReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.set_sheet_value(pylark.SetSheetValueReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于根据 spreadsheetToken 和 range 向单个范围写入数据，若范围内有数据，将被更新覆盖；单次写入不超过5000行，100列，每个格子不超过5万字符。
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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.batch_set_sheet_value(pylark.BatchSetSheetValueReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于根据 spreadsheetToken 和 range 向多个范围写入数据，若范围内有数据，将被更新覆盖；单次写入不超过5000行，100列，每个格子不超过5万字符。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.set_sheet_style(pylark.SetSheetStyleReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.batch_set_sheet_style(pylark.BatchSetSheetStyleReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.merge_sheet_cell(pylark.MergeSheetCellReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.unmerge_sheet_cell(pylark.UnmergeSheetCellReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.set_sheet_value_image(pylark.SetSheetValueImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.find_sheet(pylark.FindSheetReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.replace_sheet(pylark.ReplaceSheetReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_sheet_condition_format(pylark.CreateSheetConditionFormatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_sheet_condition_format(pylark.GetSheetConditionFormatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_sheet_condition_format(pylark.UpdateSheetConditionFormatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_sheet_condition_format(pylark.DeleteSheetConditionFormatReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_sheet_protected_dimension(pylark.CreateSheetProtectedDimensionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_sheet_protected_dimension(pylark.GetSheetProtectedDimensionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_sheet_protected_dimension(pylark.UpdateSheetProtectedDimensionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_sheet_protected_dimension(pylark.DeleteSheetProtectedDimensionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_sheet_data_validation_dropdown(pylark.CreateSheetDataValidationDropdownReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_sheet_data_validation_dropdown(pylark.DeleteSheetDataValidationDropdownReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_sheet_data_validation_dropdown(pylark.UpdateSheetDataValidationDropdownReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_sheet_data_validation_dropdown(pylark.GetSheetDataValidationDropdownReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_sheet_filter(pylark.CreateSheetFilterReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

在子表内创建筛选。

参数值可参考[筛选指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/filter-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_sheet_filter(pylark.DeleteSheetFilterReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_sheet_filter(pylark.UpdateSheetFilterReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新子表筛选范围中的列筛选条件。

参数值可参考[筛选指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/filter-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_sheet_filter(pylark.GetSheetFilterReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_sheet_filter_view(pylark.CreateSheetFilterViewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据传入的参数创建一个筛选视图。Id 和 名字可选，不填的话会默认生成；range 必填。Id 长度为10，由 0-9、a-z、A-Z 组合生成。名字长度不超过100。单个子表内的筛选视图个数不超过 150。

筛选范围的设置参考：[筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_sheet_filter_view(pylark.DeleteSheetFilterViewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_sheet_filter_view(pylark.UpdateSheetFilterViewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新筛选视图的名字或者筛选范围。名字长度不超过100，不能重复即子表内唯一；筛选范围不超过子表的最大范围。

筛选范围的设置参考：[筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_sheet_filter_view(pylark.GetSheetFilterViewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.query_sheet_filter_view(pylark.QuerySheetFilterViewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_sheet_filter_view_condition(pylark.CreateSheetFilterViewConditionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

在筛选视图的筛选范围的某一列创建筛选条件。

筛选条件参考 [筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_sheet_filter_view_condition(pylark.DeleteSheetFilterViewConditionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_sheet_filter_view_condition(pylark.UpdateSheetFilterViewConditionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新筛选视图范围的某列的筛选条件，condition id 即为列的字母号。

筛选条件参数可参考 [筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_sheet_filter_view_condition(pylark.GetSheetFilterViewConditionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取筛选视图某列的筛选条件信息。

筛选条件含义参考 [筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.query_sheet_filter_view_condition(pylark.QuerySheetFilterViewConditionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查询一个筛选视图的所有筛选条件，返回筛选视图的筛选范围内的筛选条件。

筛选条件含义可参考 [筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_sheet_float_image(pylark.CreateSheetFloatImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据传入的参数创建一张浮动图片。Float_image_token （[上传图片至表格后得到](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_all)）和range（只支持一个单元格） 必填。Float_image_id 可选，不填的话会默认生成，长度为10，由 0-9、a-z、A-Z 组合生成。表格内不重复的图片（浮动图片+单元格图片）总数不超过4000。width 和 height 为图片展示的宽高，可选，不填的话会使用图片的真实宽高。offset_x 和 offset_y 为图片左上角距离所在单元格左上角的偏移，可选，默认为 0。

浮动图片的设置参考：[浮动图片指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_sheet_float_image(pylark.DeleteSheetFloatImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_sheet_float_image(pylark.UpdateSheetFloatImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新已有的浮动图片位置和宽高，包括 range、width、height、offset_x 和 offset_y，不包括 float_image_id 和 float_image_token。

浮动图片更新参考：[浮动图片指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_sheet_float_image(pylark.GetSheetFloatImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据 float_image_id 获取对应浮动图片的信息。

浮动图片参考：[浮动图片指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.query_sheet_float_image(pylark.QuerySheetFloatImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

返回子表内所有的浮动图片信息。

浮动图片参考：[浮动图片指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)




#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/query](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/query)

### URL

`https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/query`

### Method

`GET`

## CreateWikiSpace

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateWikiSpace(ctx, &lark.CreateWikiSpaceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_wiki_space(pylark.CreateWikiSpaceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

此接口用于创建知识空间

此接口不支持应用身份访问




#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/create](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/create)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces`

### Method

`POST`

## GetWikiSpaceList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetWikiSpaceList(ctx, &lark.GetWikiSpaceListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_wiki_space_list(pylark.GetWikiSpaceListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

此接口用于获取有权限访问的知识空间列表

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/list](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/list)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces`

### Method

`GET`

## GetWikiSpace

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetWikiSpace(ctx, &lark.GetWikiSpaceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_wiki_space(pylark.GetWikiSpaceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

此接口用于根据知识空间ID来查询知识空间的信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces/:space_id`

### Method

`GET`

## UpdateWikiSpaceSetting

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.UpdateWikiSpaceSetting(ctx, &lark.UpdateWikiSpaceSettingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.update_wiki_space_setting(pylark.UpdateWikiSpaceSettingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据space_id更新知识空间公共设置

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-setting/update](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-setting/update)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces/:space_id/setting`

### Method

`PUT`

## DeleteWikiSpaceMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.DeleteWikiSpaceMember(ctx, &lark.DeleteWikiSpaceMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.delete_wiki_space_member(pylark.DeleteWikiSpaceMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

此接口用于删除知识空间成员

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-member/delete](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-member/delete)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces/:space_id/members/:member_id`

### Method

`DELETE`

## AddWikiSpaceMember

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.AddWikiSpaceMember(ctx, &lark.AddWikiSpaceMemberReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.add_wiki_space_member(pylark.AddWikiSpaceMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

添加知识空间成员

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-member/create](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-member/create)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces/:space_id/members`

### Method

`POST`

## CreateWikiNode

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.CreateWikiNode(ctx, &lark.CreateWikiNodeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.create_wiki_node(pylark.CreateWikiNodeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

此接口用于在知识库里创建节点

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/create](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/create)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces/:space_id/nodes`

### Method

`POST`

## GetWikiNodeList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetWikiNodeList(ctx, &lark.GetWikiNodeListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_wiki_node_list(pylark.GetWikiNodeListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

此接口用于分页获取Wiki节点的子节点列表

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/list](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/list)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces/:space_id/nodes`

### Method

`GET`

## MoveWikiNode

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.MoveWikiNode(ctx, &lark.MoveWikiNodeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.move_wiki_node(pylark.MoveWikiNodeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

此方法用于在Wiki内移动节点，支持跨知识空间移动

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/move](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/move)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces/:space_id/nodes/:node_token/move`

### Method

`POST`

## GetWikiNode

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetWikiNode(ctx, &lark.GetWikiNodeReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_wiki_node(pylark.GetWikiNodeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取节点信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get_node](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get_node)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces/get_node`

### Method

`GET`

## MoveDocsToWiki

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.MoveDocsToWiki(ctx, &lark.MoveDocsToWikiReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.move_docs_to_wiki(pylark.MoveDocsToWikiReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口允许添加已有云文档至知识库，并挂载在指定父页面下

### 移动操作 ###
移动后，文档将从“我的空间”或“共享空间”转移至“知识库”，并将从以下功能入口消失：
- 云空间主页：最近访问、快速访问
- 我的空间
- 共享空间
- 收藏

### 权限变更 ###
移动后，文档会向所有可查看“页面树”的用户显示，默认继承父页面的权限设置。
</md-alert



仅支持文档所有者发起请求

此接口为异步接口。若移动已完成（或节点已在Wiki中），则直接返回结果（Wiki token）。若尚未完成，则返回task id。请使用[获取任务结果](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/task/get)接口进行查询。




#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/move_docs_to_wiki](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/move_docs_to_wiki)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/spaces/:space_id/nodes/move_docs_to_wiki`

### Method

`POST`

## GetWikiTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Drive.GetWikiTask(ctx, &lark.GetWikiTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.drive.get_wiki_task(pylark.GetWikiTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该方法用于获取wiki异步任务的结果

仅发起任务的用户（或应用）可以查询任务结果。否则会返回权限报错。




#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/task/get](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/task/get)

### URL

`https://open.feishu.cn/open-apis/wiki/v2/tasks/:task_id`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.get_bitable_view_list(pylark.GetBitableViewListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据 app_token 和 table_id，获取数据表的所有视图

该接口支持调用频率上限为 20 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.create_bitable_view(pylark.CreateBitableViewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

在数据表中新增一个视图

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.delete_bitable_view(pylark.DeleteBitableViewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除数据表中的视图

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.get_bitable_record_list(pylark.GetBitableRecordListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于列出数据表中的现有记录，单次最多列出 100 行记录，支持分页获取。

该接口支持调用频率上限为 1000 次/分钟




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.get_bitable_record(pylark.GetBitableRecordReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于根据 record_id 的值检索现有记录

该接口支持调用频率上限为 20 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.create_bitable_record(pylark.CreateBitableRecordReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于在数据表中新增一条记录

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.batch_create_bitable_record(pylark.BatchCreateBitableRecordReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于在数据表中新增多条记录

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.update_bitable_record(pylark.UpdateBitableRecordReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于更新数据表中的一条记录

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.batch_update_bitable_record(pylark.BatchUpdateBitableRecordReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于更新数据表中的多条记录

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.delete_bitable_record(pylark.DeleteBitableRecordReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于删除数据表中的一条记录

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.batch_delete_bitable_record(pylark.BatchDeleteBitableRecordReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于删除数据表中现有的多条记录

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.get_bitable_field_list(pylark.GetBitableFieldListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据 app_token 和 table_id，获取数据表的所有字段

该接口支持调用频率上限为 20 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.create_bitable_field(pylark.CreateBitableFieldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于在数据表中新增一个字段

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.update_bitable_field(pylark.UpdateBitableFieldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于在数据表中更新一个字段

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.delete_bitable_field(pylark.DeleteBitableFieldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于在数据表中删除一个字段

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.get_bitable_table_list(pylark.GetBitableTableListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据  app_token，获取多维表格下的所有数据表

该接口支持调用频率上限为 20 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.create_bitable_table(pylark.CreateBitableTableReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

新增一个数据表

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.batch_create_bitable_table(pylark.BatchCreateBitableTableReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

新增多个数据表

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.delete_bitable_table(pylark.DeleteBitableTableReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除一个数据表

该接口支持调用频率上限为 10 QPS




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.batch_delete_bitable_table(pylark.BatchDeleteBitableTableReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除多个数据表

该接口支持调用频率上限为 10 QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/batch_delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/batch_delete)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/batch_delete`

### Method

`POST`

## UpdateBitableMeta

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Bitable.UpdateBitableMeta(ctx, &lark.UpdateBitableMetaReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.update_bitable_meta(pylark.UpdateBitableMetaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过 app_token 更新多维表格名称

该接口支持调用频率上限为 10 QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/update)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token`

### Method

`PUT`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.bitable.get_bitable_meta(pylark.GetBitableMetaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过 app_token 获取多维表格元数据

该接口支持调用频率上限为 20 QPS




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/get)

### URL

`https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token`

### Method

`GET`


# MeetingRoom

## GetMeetingRoomCustomization

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MeetingRoom.GetMeetingRoomCustomization(ctx, &lark.GetMeetingRoomCustomizationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.get_meeting_room_customization(pylark.GetMeetingRoomCustomizationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于获取会议室个性化配置。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uIjM5UjLyITO14iMykTN/query-meeting-room-customization-setting](https://open.feishu.cn/document/ukTMukTMukTM/uIjM5UjLyITO14iMykTN/query-meeting-room-customization-setting)

### URL

`https://open.feishu.cn/open-apis/meeting_room/room/customization`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.batch_get_meeting_room_summary(pylark.BatchGetMeetingRoomSummaryReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.get_meeting_room_building_list(pylark.GetMeetingRoomBuildingListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.batch_get_meeting_room_building(pylark.BatchGetMeetingRoomBuildingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.get_meeting_room_room_list(pylark.GetMeetingRoomRoomListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.batch_get_meeting_room_room(pylark.BatchGetMeetingRoomRoomReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.batch_get_meeting_room_freebusy(pylark.BatchGetMeetingRoomFreebusyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.reply_meeting_room_instance(pylark.ReplyMeetingRoomInstanceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.create_meeting_room_building(pylark.CreateMeetingRoomBuildingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.update_meeting_room_building(pylark.UpdateMeetingRoomBuildingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.delete_meeting_room_building(pylark.DeleteMeetingRoomBuildingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.batch_get_meeting_room_building_id(pylark.BatchGetMeetingRoomBuildingIDReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.create_meeting_room_room(pylark.CreateMeetingRoomRoomReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.update_meeting_room_room(pylark.UpdateMeetingRoomRoomReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.delete_meeting_room_room(pylark.DeleteMeetingRoomRoomReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.batch_get_meeting_room_room_id(pylark.BatchGetMeetingRoomRoomIDReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.get_meeting_room_country_list(pylark.GetMeetingRoomCountryListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.meeting_room.get_meeting_room_district_list(pylark.GetMeetingRoomDistrictListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

新建建筑时需要选择所处国家/地区，该接口用于获得系统预先提供的可供选择的城市列表。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUTNwYjL1UDM24SN1AjN](https://open.feishu.cn/document/ukTMukTMukTM/uUTNwYjL1UDM24SN1AjN)

### URL

`https://open.feishu.cn/open-apis/meeting_room/district/list`

### Method

`GET`


# Jssdk

## GetJssdkTicket

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Jssdk.GetJssdkTicket(ctx, &lark.GetJssdkTicketReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.jssdk.get_jssdk_ticket(pylark.GetJssdkTicketReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



通过在你的网页中引入**飞书网页组件**，可在你的网页中直接使用飞书的功能。
- 网页组件只适用于自建应用，暂不支持商店应用。
- 网页组件适用于普通web网页，不建议在小程序webview中调用此组件

## 准备工作
- 在开发者后台[创建一个企业自建应用](/ssl:ttdoc/home/introduction-to-custom-app-development/self-built-application-development-process)。

- 引入组件库。在网页 html 中引入如下代码：
```html
<script src="https://lf1-cdn-tos.bytegoofy.com/goofy/locl/lark/external_js_sdk/h5-js-sdk-1.1.2.js"></script>
```

若要使用成员卡片组件，SDK需要在`<body>`加载后引入。


## 鉴权流程
### 1. 获取 access_token
不同的 token 代表了组件使用者的身份。user_access_token代表以用户身份鉴权，app_access_token代表以应用身份授权。
- 成员名片组件仅支持以用户身份(user_access_token)鉴权
- 云文档组件可以同时支持以用户身份(user_access_token)和应用身份(app_access_token)授权。但是以应用身份授权云文档时，访问量受 80 次/分钟限制，且组件不支持 “编辑”、“评论”、“点赞” 等功能


:::html
<md-td>
开发者需要通过以下两种方式之一获取 token，再通过接口生成 ticket。

- 方法一：获取用户身份。通过 [第三方网站免登](/ssl:ttdoc/ukTMukTMukTM/uETOwYjLxkDM24SM5AjN)获得 user_access_token

- 方法二：获取应用身份。通过[服务端API](/ssl:ttdoc/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_access_token_internal)获得 app_access_token
  
</md-td>
:::

### 2. 获取 jsapi_ticket

为了降低泄漏风险，这一步应当在你的服务端进行。

### Doc

[https://open.feishu.cn/document/uYjL24iN/uUDO3YjL1gzN24SN4cjN](https://open.feishu.cn/document/uYjL24iN/uUDO3YjL1gzN24SN4cjN)

### URL

`https://open.feishu.cn/open-apis/jssdk/ticket/get`

### Method

`POST`


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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.apply_vc_reserve(pylark.ApplyVCReserveReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.update_vc_reserve(pylark.UpdateVCReserveReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.delete_vc_reserve(pylark.DeleteVCReserveReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.get_vc_reserve(pylark.GetVCReserveReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.get_vc_reserve_active_meeting(pylark.GetVCReserveActiveMeetingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.get_vc_meeting(pylark.GetVCMeetingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取一个会议的详细数据

只能获取归属于自己的会议，支持查询最近90天内的会议








#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/get)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id`

### Method

`GET`

## ListVCMeetingByNo

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.ListVCMeetingByNo(ctx, &lark.ListVCMeetingByNoReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.list_vc_meeting_by_no(pylark.ListVCMeetingByNoReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取指定时间范围（90天内)会议号关联的会议简要信息列表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/list_by_no](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/list_by_no)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/list_by_no`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.invite_vc_meeting(pylark.InviteVCMeetingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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

## KickoutVCMeeting

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.VC.KickoutVCMeeting(ctx, &lark.KickoutVCMeetingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.kickout_vc_meeting(pylark.KickoutVCMeetingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

将参会人从会议中移除

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/kickout](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/kickout)

### URL

`https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/kickout`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.set_vc_host_meeting(pylark.SetVCHostMeetingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.end_vc_meeting(pylark.EndVCMeetingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.start_vc_meeting_recording(pylark.StartVCMeetingRecordingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.stop_vc_meeting_recording(pylark.StopVCMeetingRecordingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.get_vc_meeting_recording(pylark.GetVCMeetingRecordingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.set_vc_permission_meeting_recording(pylark.SetVCPermissionMeetingRecordingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.get_vc_daily_report(pylark.GetVCDailyReportReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.get_vc_top_user_report(pylark.GetVCTopUserReportReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.get_vc_room_config(pylark.GetVCRoomConfigReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.vc.set_vc_room_config(pylark.SetVCRoomConfigReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.is_application_user_admin(pylark.IsApplicationUserAdminReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于查询用户是否为应用管理员。
> 此处应用管理员是指可以进入企业管理后台对应用进行审核和管理的企业管理员，并不是应用的开发者。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_user_admin_scope(pylark.GetApplicationUserAdminScopeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_app_visibility(pylark.GetApplicationAppVisibilityReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_user_visible_app(pylark.GetApplicationUserVisibleAppReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_app_list(pylark.GetApplicationAppListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.update_application_app_visibility(pylark.UpdateApplicationAppVisibilityReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_app_admin_user_list(pylark.GetApplicationAppAdminUserListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



查询审核应用的管理员列表，返回最新10个管理员账户id列表。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.check_user_is_in_application_paid_scope(pylark.CheckUserIsInApplicationPaidScopeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_order_list(pylark.GetApplicationOrderListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_order(pylark.GetApplicationOrderReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口用于查询某个订单的具体信息

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITNwUjLyUDM14iM1ATN](https://open.feishu.cn/document/ukTMukTMukTM/uITNwUjLyUDM14iM1ATN)

### URL

`https://open.feishu.cn/open-apis/pay/v1/order/get`

### Method

`GET`

## GetApplicationUnderAuditList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationUnderAuditList(ctx, &lark.GetApplicationUnderAuditListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_under_audit_list(pylark.GetApplicationUnderAuditListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查看本企业下所有待审核的自建应用列表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/underauditlist](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/underauditlist)

### URL

`https://open.feishu.cn/open-apis/application/v6/applications/underauditlist`

### Method

`GET`

## GetApplication

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplication(ctx, &lark.GetApplicationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application(pylark.GetApplicationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据app_id获取应用的基础信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/get)

### URL

`https://open.feishu.cn/open-apis/application/v6/applications/:app_id`

### Method

`GET`

## GetApplicationVersion

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationVersion(ctx, &lark.GetApplicationVersionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_version(pylark.GetApplicationVersionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据 app_id，version_id 获取对应应用版本的信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/get)

### URL

`https://open.feishu.cn/open-apis/application/v6/applications/:app_id/app_versions/:version_id`

### Method

`GET`

## UpdateApplicationVersion

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.UpdateApplicationVersion(ctx, &lark.UpdateApplicationVersionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.update_application_version(pylark.UpdateApplicationVersionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过接口来更新应用版本的审核结果：通过后应用可以直接上架；拒绝后则开发者可以看到拒绝理由，并在修改后再次申请发布。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/patch)

### URL

`https://open.feishu.cn/open-apis/application/v6/applications/:app_id/app_versions/:version_id`

### Method

`PATCH`

## UpdateApplication

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.UpdateApplication(ctx, &lark.UpdateApplicationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.update_application(pylark.UpdateApplicationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新应用的分组信息（分组会影响应用在工作台中的分类情况，请谨慎更新）

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/patch)

### URL

`https://open.feishu.cn/open-apis/application/v6/applications/:app_id`

### Method

`PATCH`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_usage_overview(pylark.GetApplicationUsageOverviewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查看应用在某一天/某一周/某一个月的使用数据，可以查看租户整体对应用的使用情况，也可以分部门查看。

1. 仅支持企业版/旗舰版租户使用
2. 一般每天早上10点产出前一天的数据
3. 已经支持的指标包括：应用的活跃用户数、累计用户数、新增用户数
4. 数据从飞书3.46版本开始统计，使用飞书版本3.45及以下版本的用户数据不会被统计到
5. 按照部门查看数据时，会展示当前部门以及其子部门的整体使用情况




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_usage/overview](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_usage/overview)

### URL

`https://open.feishu.cn/open-apis/application/v6/applications/:app_id/app_usage/overview`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_usage_trend(pylark.GetApplicationUsageTrendReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



查询应用在指定时间段内企业员工的使用趋势信息。

此接口目前仅支持小程序的使用情况查询，不支持网页应用和机器人应用的使用情况查询;仅支持查询自建应用，不支持查询商店应用


#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uITN0YjLyUDN24iM1QjN](https://open.feishu.cn/document/ukTMukTMukTM/uITN0YjLyUDN24iM1QjN)

### URL

`https://open.feishu.cn/open-apis/application/v1/app_usage_trend`

### Method

`POST`

## UpdateApplicationFeedback

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.UpdateApplicationFeedback(ctx, &lark.UpdateApplicationFeedbackReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.update_application_feedback(pylark.UpdateApplicationFeedbackReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新应用的反馈数据

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-feedback/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-feedback/patch)

### URL

`https://open.feishu.cn/open-apis/application/v6/applications/:app_id/feedbacks/:feedback_id`

### Method

`PATCH`

## GetApplicationFeedbackList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Application.GetApplicationFeedbackList(ctx, &lark.GetApplicationFeedbackListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.application.get_application_feedback_list(pylark.GetApplicationFeedbackListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查询应用的反馈数据

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-feedback/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-feedback/list)

### URL

`https://open.feishu.cn/open-apis/application/v6/applications/:app_id/feedbacks`

### Method

`GET`


# Mail

## GetMailUser

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetMailUser(ctx, &lark.GetMailUserReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_mail_user(pylark.GetMailUserReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

使用邮箱状态查询接口，可以输入邮箱地址，查询出该邮箱地址对应的类型以及状态

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user/query)

### URL

`https://open.feishu.cn/open-apis/mail/v1/users/query`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.create_mailgroup(pylark.CreateMailGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_mailgroup(pylark.GetMailGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_mail_group_list(pylark.GetMailGroupListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.update_mail_group_patch(pylark.UpdateMailGroupPatchReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.update_mailgroup(pylark.UpdateMailGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.delete_mailgroup(pylark.DeleteMailGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.create_mail_group_member(pylark.CreateMailGroupMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_mail_group_member(pylark.GetMailGroupMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_mail_group_member_list(pylark.GetMailGroupMemberListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.delete_mail_group_member(pylark.DeleteMailGroupMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.create_mail_group_permission_member(pylark.CreateMailGroupPermissionMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_mail_group_permission_member(pylark.GetMailGroupPermissionMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_mail_group_permission_member_list(pylark.GetMailGroupPermissionMemberListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.delete_mail_group_permission_member(pylark.DeleteMailGroupPermissionMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

从自定义成员中删除单个成员，删除后该成员无法发送邮件到该邮件组

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/delete)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/:permission_member_id`

### Method

`DELETE`

## CreateMailGroupAlias

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.CreateMailGroupAlias(ctx, &lark.CreateMailGroupAliasReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.create_mail_group_alias(pylark.CreateMailGroupAliasReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

创建邮件组别名

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-alias/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-alias/create)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases`

### Method

`POST`

## GetMailGroupAliasList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetMailGroupAliasList(ctx, &lark.GetMailGroupAliasListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_mail_group_alias_list(pylark.GetMailGroupAliasListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取邮件组所有别名

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-alias/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-alias/list)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases`

### Method

`GET`

## DeleteMailGroupAlias

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.DeleteMailGroupAlias(ctx, &lark.DeleteMailGroupAliasReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.delete_mail_group_alias(pylark.DeleteMailGroupAliasReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除邮件组别名

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-alias/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-alias/delete)

### URL

`https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases/:alias_id`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.create_public_mailbox(pylark.CreatePublicMailboxReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_public_mailbox(pylark.GetPublicMailboxReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_public_mailbox_list(pylark.GetPublicMailboxListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.update_public_mailbox_patch(pylark.UpdatePublicMailboxPatchReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.update_public_mailbox(pylark.UpdatePublicMailboxReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新公共邮箱所有信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/update)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id`

### Method

`PUT`

## DeletePublicMailbox

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.DeletePublicMailbox(ctx, &lark.DeletePublicMailboxReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.delete_public_mailbox(pylark.DeletePublicMailboxReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口会永久删除公共邮箱地址。可用于释放邮箱回收站的公共邮箱地址，一旦删除，该邮箱地址将无法恢复。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/delete)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id`

### Method

`DELETE`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.create_public_mailbox_member(pylark.CreatePublicMailboxMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_public_mailbox_member(pylark.GetPublicMailboxMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_public_mailbox_member_list(pylark.GetPublicMailboxMemberListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.delete_public_mailbox_member(pylark.DeletePublicMailboxMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.clear_public_mailbox_member(pylark.ClearPublicMailboxMemberReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除公共邮箱所有成员

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/clear](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/clear)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/clear`

### Method

`POST`

## CreateMailPublicMailboxAlias

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.CreateMailPublicMailboxAlias(ctx, &lark.CreateMailPublicMailboxAliasReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.create_mail_public_mailbox_alias(pylark.CreateMailPublicMailboxAliasReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

创建公共邮箱别名

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-alias/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-alias/create)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases`

### Method

`POST`

## GetMailPublicMailboxAliasList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetMailPublicMailboxAliasList(ctx, &lark.GetMailPublicMailboxAliasListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_mail_public_mailbox_alias_list(pylark.GetMailPublicMailboxAliasListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取所有公共邮箱别名

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-alias/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-alias/list)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases`

### Method

`GET`

## DeleteMailPublicMailboxAlias

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.DeleteMailPublicMailboxAlias(ctx, &lark.DeleteMailPublicMailboxAliasReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.delete_mail_public_mailbox_alias(pylark.DeleteMailPublicMailboxAliasReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除公共邮箱别名

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-alias/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-alias/delete)

### URL

`https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases/:alias_id`

### Method

`DELETE`

## CreateMailUserMailboxAlias

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.CreateMailUserMailboxAlias(ctx, &lark.CreateMailUserMailboxAliasReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.create_mail_user_mailbox_alias(pylark.CreateMailUserMailboxAliasReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

创建用户邮箱别名

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/create)

### URL

`https://open.feishu.cn/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases`

### Method

`POST`

## DeleteMailUserMailboxAlias

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.DeleteMailUserMailboxAlias(ctx, &lark.DeleteMailUserMailboxAliasReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.delete_mail_user_mailbox_alias(pylark.DeleteMailUserMailboxAliasReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除用户邮箱别名

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/delete)

### URL

`https://open.feishu.cn/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases/:alias_id`

### Method

`DELETE`

## GetMailUserMailboxAliasList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.GetMailUserMailboxAliasList(ctx, &lark.GetMailUserMailboxAliasListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.get_mail_user_mailbox_alias_list(pylark.GetMailUserMailboxAliasListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取用户邮箱所有别名

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/list)

### URL

`https://open.feishu.cn/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases`

### Method

`GET`

## DeleteMailUserMailbox

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Mail.DeleteMailUserMailbox(ctx, &lark.DeleteMailUserMailboxReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mail.delete_mail_user_mailbox(pylark.DeleteMailUserMailboxReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口会永久删除用户邮箱地址。可用于删除位于邮箱回收站中的用户邮箱地址，一旦删除，将无法恢复。该接口支持邮件的转移，可以将被释放邮箱的邮件转移到另外一个可以使用的邮箱中。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox/delete)

### URL

`https://open.feishu.cn/open-apis/mail/v1/user_mailboxes/:user_mailbox_id`

### Method

`DELETE`


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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.get_approval(pylark.GetApprovalReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.get_approval_instance_list(pylark.GetApprovalInstanceListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.get_approval_instance(pylark.GetApprovalInstanceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.create_approval_instance(pylark.CreateApprovalInstanceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.approve_approval_instance(pylark.ApproveApprovalInstanceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.reject_approval_instance(pylark.RejectApprovalInstanceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.transfer_approval_instance(pylark.TransferApprovalInstanceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



对于单个审批任务进行转交操作。转交后审批流程流转给被转交人。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDNyUjL1QjM14SN0ITN](https://open.feishu.cn/document/ukTMukTMukTM/uUDNyUjL1QjM14SN0ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/transfer`

### Method

`POST`

## RollbackApprovalInstance

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.RollbackApprovalInstance(ctx, &lark.RollbackApprovalInstanceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.rollback_approval_instance(pylark.RollbackApprovalInstanceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



从当前审批任务，退回到已审批的一个或多个任务节点。退回后，已审批节点重新生成审批任务

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-task-return](https://open.feishu.cn/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-task-return)

### URL

`https://open.feishu.cn/open-apis/approval/v4/instances/specified_rollback`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.cancel_approval_instance(pylark.CancelApprovalInstanceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



对于状态为“审批中”的单个审批实例进行撤销操作，撤销后审批流程结束
#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYDNyUjL2QjM14iN0ITN](https://open.feishu.cn/document/ukTMukTMukTM/uYDNyUjL2QjM14iN0ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/cancel`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.search_approval_instance(pylark.SearchApprovalInstanceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口通过不同条件查询审批系统中符合条件的审批实例列表。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uQjMxYjL0ITM24CNyEjN](https://open.feishu.cn/document/ukTMukTMukTM/uQjMxYjL0ITM24CNyEjN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/search`

### Method

`POST`

## AddApprovalInstanceSign

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.AddApprovalInstanceSign(ctx, &lark.AddApprovalInstanceSignReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.add_approval_instance_sign(pylark.AddApprovalInstanceSignReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



对于单个审批任务进行加签操作。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-task-addsign](https://open.feishu.cn/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-task-addsign)

### URL

`https://open.feishu.cn/open-apis/approval/v4/instances/add_sign`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.upload_approval_file(pylark.UploadApprovalFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



当审批表单中有图片或附件控件时，开发者需在创建审批实例前通过审批上传文件接口将文件上传到审批系统，且附件上传大小限制为50M，图片上传大小为10M。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUDOyUjL1gjM14SN4ITN](https://open.feishu.cn/document/ukTMukTMukTM/uUDOyUjL1gjM14SN4ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/file/upload`

### Method

`POST`

## SearchApprovalTask

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.SearchApprovalTask(ctx, &lark.SearchApprovalTaskReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.search_approval_task(pylark.SearchApprovalTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口通过不同条件查询审批系统中符合条件的审批任务列表

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYjMxYjL2ITM24iNyEjN](https://open.feishu.cn/document/ukTMukTMukTM/uYjMxYjL2ITM24iNyEjN)

### URL

`https://www.feishu.cn/approval/openapi/v2/task/search`

### Method

`POST`

## GetApprovalUserTaskList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.GetApprovalUserTaskList(ctx, &lark.GetApprovalUserTaskListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.get_approval_user_task_list(pylark.GetApprovalUserTaskListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据用户和任务分组查询任务列表

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/query)

### URL

`https://open.feishu.cn/open-apis/approval/v4/tasks/query`

### Method

`GET`

## SearchApprovalCarbonCopy

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.SearchApprovalCarbonCopy(ctx, &lark.SearchApprovalCarbonCopyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.search_approval_carbon_copy(pylark.SearchApprovalCarbonCopyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



该接口通过不同条件查询审批系统中符合条件的审批抄送列表。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uUjMxYjL1ITM24SNyEjN](https://open.feishu.cn/document/ukTMukTMukTM/uUjMxYjL1ITM24SNyEjN)

### URL

`https://www.feishu.cn/approval/openapi/v2/cc/search`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.create_approval_carbon_copy(pylark.CreateApprovalCarbonCopyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



通过接口可以将当前审批实例抄送给其他人。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uADOzYjLwgzM24CM4MjN](https://open.feishu.cn/document/ukTMukTMukTM/uADOzYjLwgzM24CM4MjN)

### URL

`https://www.feishu.cn/approval/openapi/v2/instance/cc`

### Method

`POST`

## PreviewApprovalInstance

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.PreviewApprovalInstance(ctx, &lark.PreviewApprovalInstanceReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.preview_approval_instance(pylark.PreviewApprovalInstanceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



提交审批前，预览审批流程。或者发起审批后，在某一审批节点预览后续流程

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-preview](https://open.feishu.cn/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-preview)

### URL

`https://open.feishu.cn/open-apis/approval/v4/instances/preview`

### Method

`POST`

## UpdateApprovalMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.UpdateApprovalMessage(ctx, &lark.UpdateApprovalMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.update_approval_message(pylark.UpdateApprovalMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



此接口可以根据审批bot消息id及相应状态，更新相应的审批bot消息，只可用于更新待审批模板的bot消息。例如，给用户推送了审批待办消息，当用户处理该消息后，可以将之前推送的Bot消息更新为已审批。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uAjNyYjLwYjM24CM2IjN](https://open.feishu.cn/document/ukTMukTMukTM/uAjNyYjLwYjM24CM2IjN)

### URL

`https://www.feishu.cn/approval/openapi/v1/message/update`

### Method

`POST`

## SubscribeApprovalSubscription

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.SubscribeApprovalSubscription(ctx, &lark.SubscribeApprovalSubscriptionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.subscribe_approval_subscription(pylark.SubscribeApprovalSubscriptionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



应用订阅 approval_code 后，该应用就可以收到该审批定义对应实例的事件通知。同一应用只需要订阅一次，无需重复订阅。

当应用不希望再收到审批事件时，可以使用取消订阅接口进行取消，取消后将不再给应用推送消息。

订阅和取消订阅都是应用维度的，多个应用可以同时订阅同一个 approval_code，每个应用都能收到审批事件。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ucDOyUjL3gjM14yN4ITN](https://open.feishu.cn/document/ukTMukTMukTM/ucDOyUjL3gjM14yN4ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/subscription/subscribe`

### Method

`POST`

## UnsubscribeApprovalSubscription

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Approval.UnsubscribeApprovalSubscription(ctx, &lark.UnsubscribeApprovalSubscriptionReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.approval.unsubscribe_approval_subscription(pylark.UnsubscribeApprovalSubscriptionReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



取消订阅 approval_code 后，无法再收到该审批定义对应实例的事件通知。

#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/ugDOyUjL4gjM14CO4ITN](https://open.feishu.cn/document/ukTMukTMukTM/ugDOyUjL4gjM14CO4ITN)

### URL

`https://www.feishu.cn/approval/openapi/v2/subscription/unsubscribe`

### Method

`POST`


# Helpdesk

## CreateHelpdeskNotification

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.CreateHelpdeskNotification(ctx, &lark.CreateHelpdeskNotificationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.create_helpdesk_notification(pylark.CreateHelpdeskNotificationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

调用接口创建推送，创建成功后为草稿状态

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/create)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/notifications`

### Method

`POST`

## UpdateHelpdeskNotification

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.UpdateHelpdeskNotification(ctx, &lark.UpdateHelpdeskNotificationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.update_helpdesk_notification(pylark.UpdateHelpdeskNotificationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新推送信息，只有在草稿状态下才可以调用此接口进行更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/patch](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/patch)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/notifications/:notification_id`

### Method

`PATCH`

## GetHelpdeskNotification

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskNotification(ctx, &lark.GetHelpdeskNotificationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_notification(pylark.GetHelpdeskNotificationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查询推送详情

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/get)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/notifications/:notification_id`

### Method

`GET`

## PreviewHelpdeskNotification

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.PreviewHelpdeskNotification(ctx, &lark.PreviewHelpdeskNotificationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.preview_helpdesk_notification(pylark.PreviewHelpdeskNotificationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

在正式执行推送之前是可以调用此接口预览设置的推送内容

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/preview](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/preview)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/notifications/:notification_id/preview`

### Method

`POST`

## SubmitApproveHelpdeskNotification

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.SubmitApproveHelpdeskNotification(ctx, &lark.SubmitApproveHelpdeskNotificationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.submit_approve_helpdesk_notification(pylark.SubmitApproveHelpdeskNotificationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

正常情况下调用创建推送接口后，就可以调用提交审核接口，如果创建人是服务台owner则会自动审核通过，否则会通知服务台owner审核此推送信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/submit_approve](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/submit_approve)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/notifications/:notification_id/submit_approve`

### Method

`POST`

## CancelApproveHelpdeskNotification

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.CancelApproveHelpdeskNotification(ctx, &lark.CancelApproveHelpdeskNotificationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.cancel_approve_helpdesk_notification(pylark.CancelApproveHelpdeskNotificationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

提交审核后，如果需要取消审核，则调用此接口

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/cancel_approve](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/cancel_approve)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/notifications/:notification_id/cancel_approve`

### Method

`POST`

## ExecuteSendHelpdeskNotification

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.ExecuteSendHelpdeskNotification(ctx, &lark.ExecuteSendHelpdeskNotificationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.execute_send_helpdesk_notification(pylark.ExecuteSendHelpdeskNotificationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

审核通过后调用此接口设置推送时间，等待调度系统调度，发送消息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/execute_send](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/execute_send)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/notifications/:notification_id/execute_send`

### Method

`POST`

## CancelSendHelpdeskNotification

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.CancelSendHelpdeskNotification(ctx, &lark.CancelSendHelpdeskNotificationReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.cancel_send_helpdesk_notification(pylark.CancelSendHelpdeskNotificationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

取消推送接口，审核通过后待调度可以调用，发送过程中可以调用（会撤回已发送的消息），发送完成后可以需要推送（会撤回所有已发送的消息）

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/cancel_send](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/notification/cancel_send)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/notifications/:notification_id/cancel_send`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.start_helpdesk_service(pylark.StartHelpdeskServiceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_ticket(pylark.GetHelpdeskTicketReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.update_helpdesk_ticket(pylark.UpdateHelpdeskTicketReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于更新服务台工单详情。只会更新数据，不会触发相关操作。如修改工单状态到关单，不会关闭聊天页面。仅支持自建应用。要更新的工单字段必须至少输入一项。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_ticket_list(pylark.GetHelpdeskTicketListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.download_helpdesk_ticket_image(pylark.DownloadHelpdeskTicketImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.answer_helpdesk_ticket_user_query(pylark.AnswerHelpdeskTicketUserQueryReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于回复用户提问结果至工单，需要工单仍处于进行中且未接入人工状态。仅支持自建应用。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/answer_user_query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/answer_user_query)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id/answer_user_query`

### Method

`POST`

## GetHelpdeskTicketCustomizedFields

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.GetHelpdeskTicketCustomizedFields(ctx, &lark.GetHelpdeskTicketCustomizedFieldsReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_ticket_customized_fields(pylark.GetHelpdeskTicketCustomizedFieldsReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于获取服务台自定义字段详情。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/customized_fields](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/customized_fields)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/customized_fields`

### Method

`GET`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_ticket_message_list(pylark.GetHelpdeskTicketMessageListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.send_helpdesk_ticket_message(pylark.SendHelpdeskTicketMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于工单发送消息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket-message/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket-message/create)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id/messages`

### Method

`POST`

## SendHelpdeskMessage

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Helpdesk.SendHelpdeskMessage(ctx, &lark.SendHelpdeskMessageReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.send_helpdesk_message(pylark.SendHelpdeskMessageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过服务台机器人给指定用户的服务台专属群或私聊发送消息，支持文本、富文本、卡片、图片。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/bot-message/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/bot-message/create)

### URL

`https://open.feishu.cn/open-apis/helpdesk/v1/message`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_ticket_customized_field_list(pylark.GetHelpdeskTicketCustomizedFieldListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.delete_helpdesk_ticket_customized_field(pylark.DeleteHelpdeskTicketCustomizedFieldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.update_helpdesk_ticket_customized_field(pylark.UpdateHelpdeskTicketCustomizedFieldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.create_helpdesk_ticket_customized_field(pylark.CreateHelpdeskTicketCustomizedFieldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_ticket_customized_field(pylark.GetHelpdeskTicketCustomizedFieldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.create_helpdesk_category(pylark.CreateHelpdeskCategoryReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_category(pylark.GetHelpdeskCategoryReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.update_helpdesk_category(pylark.UpdateHelpdeskCategoryReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.delete_helpdesk_category(pylark.DeleteHelpdeskCategoryReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_category_list(pylark.GetHelpdeskCategoryListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.create_helpdesk_faq(pylark.CreateHelpdeskFAQReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_faq(pylark.GetHelpdeskFAQReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.update_helpdesk_faq(pylark.UpdateHelpdeskFAQReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.delete_helpdesk_faq(pylark.DeleteHelpdeskFAQReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_faq_list(pylark.GetHelpdeskFAQListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_faq_image(pylark.GetHelpdeskFAQImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.search_helpdesk_faq(pylark.SearchHelpdeskFAQReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.update_helpdesk_agent(pylark.UpdateHelpdeskAgentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_agent_email(pylark.GetHelpdeskAgentEmailReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.create_helpdesk_agent_schedule(pylark.CreateHelpdeskAgentScheduleReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.delete_helpdesk_agent_schedule(pylark.DeleteHelpdeskAgentScheduleReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.update_helpdesk_agent_schedule(pylark.UpdateHelpdeskAgentScheduleReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于更新客服的日程

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_agent_schedule(pylark.GetHelpdeskAgentScheduleReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_agent_schedule_list(pylark.GetHelpdeskAgentScheduleListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.create_helpdesk_agent_skill(pylark.CreateHelpdeskAgentSkillReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_agent_skill(pylark.GetHelpdeskAgentSkillReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.update_helpdesk_agent_skill(pylark.UpdateHelpdeskAgentSkillReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.delete_helpdesk_agent_skill(pylark.DeleteHelpdeskAgentSkillReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_agent_skill_list(pylark.GetHelpdeskAgentSkillListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.get_helpdesk_agent_skill_rule_list(pylark.GetHelpdeskAgentSkillRuleListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.subscribe_helpdesk_event(pylark.SubscribeHelpdeskEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.helpdesk.unsubscribe_helpdesk_event(pylark.UnsubscribeHelpdeskEventReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.admin.get_admin_dept_stats(pylark.GetAdminDeptStatsReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于获取部门维度的用户活跃和功能使用数据，即IM（即时通讯）、日历、云文档、音视频会议功能的使用数据。

- 只有企业自建应用才有权限调用此接口

- 当天的数据会在第二天的早上九点半产出（UTC+8）

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.admin.get_admin_user_stats(pylark.GetAdminUserStatsReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

用于获取用户维度的用户活跃和功能使用数据，即IM（即时通讯）、日历、云文档、音视频会议功能的使用数据。

- 只有企业自建应用才有权限调用此接口

- 当天的数据会在第二天的早上九点半产出（UTC+8）

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.human_auth.get_face_verify_auth_result(pylark.GetFaceVerifyAuthResultReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.human_auth.upload_face_verify_image(pylark.UploadFaceVerifyImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.human_auth.crop_face_verify_image(pylark.CropFaceVerifyImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.human_auth.create_identity(pylark.CreateIdentityReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于录入实名认证的身份信息，在唤起有源活体认证前，需要使用该接口进行实名认证。

实名认证接口会有计费管理，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。

仅通过计费申请的应用，才能在[开发者后台](https://open.feishu.cn/app)查找并申请该接口的权限。


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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.ai.recognize_basic_image(pylark.RecognizeBasicImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

可识别图片中的文字，按图片中的区域划分，分段返回文本列表

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.ai.recognize_speech_stream(pylark.RecognizeSpeechStreamReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.ai.recognize_speech_file(pylark.RecognizeSpeechFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.ai.translate_text(pylark.TranslateTextReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

机器翻译 (MT)，支持中日英（zh、ja、en）三语互译

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.ai.detect_text_language(pylark.DetectTextLanguageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

机器翻译 (MT)，支持 100 多种语言识别，返回符合 ISO 639-1 标准

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.ai.detect_face_attributes(pylark.DetectFaceAttributesReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

检测图片中的人脸属性和质量等信息



### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/face_detection-v1/image/detect_face_attributes](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/face_detection-v1/image/detect_face_attributes)

### URL

`https://open.feishu.cn/open-apis/face_detection/v1/image/detect_face_attributes`

### Method

`POST`


# Attendance

## CreateAttendanceGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.CreateAttendanceGroup(ctx, &lark.CreateAttendanceGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.create_attendance_group(pylark.CreateAttendanceGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



考勤组，是对部门或者员工在某个特定场所及特定时间段内的出勤情况（包括上下班、迟到、早退、病假、婚假、丧假、公休、工作时间、加班情况等）的一种规则设定。

通过设置考勤组，可以从部门、员工两个维度，来设定考勤方式、考勤时间、考勤地点等考勤规则。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=attendance&version=v1&resource=group&method=create)










出于安全考虑，目前通过该接口只允许修改自己创建的考勤组。


## 考勤组负责人
考勤组负责人可修改该考勤组的排班，并查看该考勤组的考勤数据。

如果考勤组负责人同时被企业管理员赋予了考勤管理员的角色，则该考勤组负责人还拥有考勤管理员的权限，可以编辑及删除考勤规则。

## 考勤组人员
可按部门、员工两个维度，设置需要参加考勤或无需参加考勤的人员。

- 若是按部门维度添加的考勤人员，当有新员工加入该部门时，其会自动加入该考勤组。
- 若是按员工维度添加的考勤人员，当其上级部门被添加到其他考勤组时，该员工不会更换考勤组。

## 考勤组类型
提供 3 种不同的考勤类型：固定班制、排班制、自由班制。

- 固定班制：指考勤组内每位人员的上下班时间一致，适用于上下班时间固定或无需安排多个班次的考勤组。
- 排班制：指考勤组人员的上下班时间不完全一致，可自定义安排每位人员的上下班时间，适用于存在多个班次如早晚班的考勤组。
- 自由班制：指没有具体的班次，考勤组人员可以在打卡时段内自由打卡，按照打卡时段统计上班时长。

## 考勤班次
- 固定班制下，需设置周一到周日每天安排哪个班次，以及可针对特殊日期进行打卡设置。
- 排班制下，需对考勤组内每一位人员的每一天进行班次指定。
- 自由班制下，需设置一天中最早打卡时间和最晚打卡时间，以及一周中哪几天需要打卡。

## 考勤方式
支持 3 种考勤方式：GPS 打卡、Wi-Fi 打卡、考勤机打卡。

- GPS 打卡：需设置经纬度信息及考勤地点名称。
- Wi-Fi 打卡：需设置 Wi-Fi 名称及 Wi-Fi 的 MAC 地址。
- 考勤机打卡：需设置考勤机名称及考勤机序号。

## 考勤其他设置
- 规则设置：支持设置是否允许外勤打卡，是否允许补卡以及一个月补卡的次数，是否允许 PC 端打卡。
- 安全设置：支持设置是否开启人脸识别打卡，以及什么情况下开启人脸识别。
- 统计设置：支持设置考勤组人员是否可以查看到某些维度的统计数据。
- 加班设置：支持配置加班时间的计算规则。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/create)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/groups`

### Method

`POST`

## SearchAttendanceGroup

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.SearchAttendanceGroup(ctx, &lark.SearchAttendanceGroupReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.search_attendance_group(pylark.SearchAttendanceGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

按考勤组名称查询考勤组摘要信息。查询条件支持名称精确匹配和模糊匹配两种方式。查询结果按考勤组修改时间 desc 排序，且最大记录数为 10 条。

该接口依赖的数据和考勤组主数据间存在数据同步延时（正常数据同步 2 秒以内），因此在使用该接口时需注意评估数据延迟潜在风险。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/search)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/groups/search`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_group(pylark.GetAttendanceGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过考勤组 ID 获取考勤组详情。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/get)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/groups/:group_id`

### Method

`GET`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.delete_attendance_group(pylark.DeleteAttendanceGroupReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过班次 ID 删除班次。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/delete)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/groups/:group_id`

### Method

`DELETE`

## GetAttendanceShift

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceShift(ctx, &lark.GetAttendanceShiftReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_shift(pylark.GetAttendanceShiftReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过班次的名称查询班次信息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/query)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/shifts/query`

### Method

`POST`

## GetAttendanceShiftDetail

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceShiftDetail(ctx, &lark.GetAttendanceShiftDetailReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_shift_detail(pylark.GetAttendanceShiftDetailReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过班次 ID 获取班次详情。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/get)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id`

### Method

`GET`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.delete_attendance_shift(pylark.DeleteAttendanceShiftReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过班次 ID 删除班次。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/delete)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id`

### Method

`DELETE`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.create_attendance_shift(pylark.CreateAttendanceShiftReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

班次是描述一次考勤任务时间规则的统称，比如一天打多少次卡，每次卡的上下班时间，晚到多长时间算迟到，晚到多长时间算缺卡等。

- 创建一个考勤组前，必须先创建一个或者多个班次。
- 一个公司内的班次是共享的，你可以直接引用他人创建的班次，但是需要注意的是，若他人修改了班次，会影响到你的考勤组及其考勤结果。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/create)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/shifts`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_user_daily_shift(pylark.GetAttendanceUserDailyShiftReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

支持查询多个用户的排班情况，查询的时间跨度不能超过 30 天。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_daily_shift/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_daily_shift/query)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_daily_shifts/query`

### Method

`POST`

## BatchCreateAttendanceUserDailyShift

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.BatchCreateAttendanceUserDailyShift(ctx, &lark.BatchCreateAttendanceUserDailyShiftReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.batch_create_attendance_user_daily_shift(pylark.BatchCreateAttendanceUserDailyShiftReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

班表是用来描述考勤组内人员每天按哪个班次进行上班。目前班表支持按一个整月对一位或多位人员进行排班。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_daily_shift/batch_create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_daily_shift/batch_create)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_daily_shifts/batch_create`

### Method

`POST`

## GetAttendanceUserStatsField

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserStatsField(ctx, &lark.GetAttendanceUserStatsFieldReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_user_stats_field(pylark.GetAttendanceUserStatsFieldReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查询考勤统计支持的日度统计或月度统计的统计表头。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_field/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_field/query)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_stats_fields/query`

### Method

`POST`

## GetAttendanceUserStatsView

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserStatsView(ctx, &lark.GetAttendanceUserStatsViewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_user_stats_view(pylark.GetAttendanceUserStatsViewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查询开发者定制的日度统计或月度统计的统计报表表头设置信息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/query)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_stats_views/query`

### Method

`POST`

## UpdateAttendanceUserStatsView

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.UpdateAttendanceUserStatsView(ctx, &lark.UpdateAttendanceUserStatsViewReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.update_attendance_user_stats_view(pylark.UpdateAttendanceUserStatsViewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新开发者定制的日度统计或月度统计的统计报表表头设置信息。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/update)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_stats_views/:user_stats_view_id`

### Method

`PUT`

## GetAttendanceUserStatsData

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserStatsData(ctx, &lark.GetAttendanceUserStatsDataReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_user_stats_data(pylark.GetAttendanceUserStatsDataReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

查询日度统计或月度统计的统计数据。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_data/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_data/query)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_stats_datas/query`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.batch_get_attendance_user_flow(pylark.BatchGetAttendanceUserFlowReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

批量查询授权内员工的实际打卡流水记录。例如，企业给一个员工设定的班次是上午 9 点和下午 6 点各打一次上下班卡，但是该员工在这期间打了多次卡，该接口会把所有的打卡记录都返回。

如果只需获取打卡结果，而不需要详细的打卡数据，可使用“获取打卡结果”的接口。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_flow/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_flow/query)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_flows/query`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_user_flow(pylark.GetAttendanceUserFlowReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过打卡记录 ID 获取用户的打卡流水记录。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_flow/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_flow/get)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_flows/:user_flow_id`

### Method

`GET`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_user_task(pylark.GetAttendanceUserTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取企业内员工的实际打卡结果，包括上班打卡结果和下班打卡结果。

- 如果企业给一个员工设定的班次是上午 9 点和下午 6 点各打一次上下班卡，即使员工在这期间打了多次卡，该接口也只会返回 1 条记录。
- 如果要获取打卡的详细数据，如打卡位置等信息，可使用“获取打卡流水记录”或“批量查询打卡流水记录”的接口。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task/query)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_tasks/query`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.batch_create_attendance_user_flow(pylark.BatchCreateAttendanceUserFlowReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

导入授权内员工的打卡流水记录。导入后，会根据员工所在的考勤组班次规则，计算最终的打卡状态与结果。

适用于考勤机数据导入等场景。




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_flow/batch_create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_flow/batch_create)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_flows/batch_create`

### Method

`POST`

## GetAttendanceUserTaskRemedyAllowedRemedyList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserTaskRemedyAllowedRemedyList(ctx, &lark.GetAttendanceUserTaskRemedyAllowedRemedyListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_user_task_remedy_allowed_remedy_list(pylark.GetAttendanceUserTaskRemedyAllowedRemedyListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取用户某天可以补的第几次上 / 下班卡的时间。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/query_user_allowed_remedys](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/query_user_allowed_remedys)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_task_remedys/query_user_allowed_remedys`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_user_task_remedy(pylark.GetAttendanceUserTaskRemedyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取授权内员工的补卡记录。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/query)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_task_remedys/query`

### Method

`POST`

## CreateAttendanceUserTaskRemedy

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.CreateAttendanceUserTaskRemedy(ctx, &lark.CreateAttendanceUserTaskRemedyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.create_attendance_user_task_remedy(pylark.CreateAttendanceUserTaskRemedyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

对于只使用飞书考勤系统而未使用飞书审批系统的企业，可以通过该接口，将在三方审批系统中发起的补卡审批数据，写入到飞书考勤系统中。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/create)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_task_remedys`

### Method

`POST`

## GetAttendanceUserSettingList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.GetAttendanceUserSettingList(ctx, &lark.GetAttendanceUserSettingListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_user_setting_list(pylark.GetAttendanceUserSettingListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

批量查询授权内员工的用户设置信息，包括人脸照片文件 ID、人脸照片更新时间。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_setting/query](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_setting/query)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_settings/query`

### Method

`GET`

## UpdateAttendanceUserSetting

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Attendance.UpdateAttendanceUserSetting(ctx, &lark.UpdateAttendanceUserSettingReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.update_attendance_user_setting(pylark.UpdateAttendanceUserSettingReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

修改授权内员工的用户设置信息，包括人脸照片文件 ID。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_setting/modify](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_setting/modify)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_settings/modify`

### Method

`POST`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.download_attendance_file(pylark.DownloadAttendanceFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过文件 ID 下载指定的文件。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/file/download](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/file/download)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.upload_attendance_file(pylark.UploadAttendanceFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

上传文件并获取文件 ID，可用于“修改用户设置”接口中的 face_key 参数。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/file/upload](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/file/upload)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/files/upload`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.get_attendance_user_approval(pylark.GetAttendanceUserApprovalReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



获取员工在某段时间内的请假、加班、外出和出差四种审批的通过数据。



请假的假期时长字段(interval)，暂未开放提供，待后续提供。


#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/get-user-attendance-data2](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/get-user-attendance-data2)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.create_attendance_user_approval(pylark.CreateAttendanceUserApprovalReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



由于部分企业使用的是自己的审批系统，而不是飞书审批系统，因此员工的请假、加班等数据无法流入到飞书考勤系统中，导致员工在请假时间段内依然收到打卡提醒，并且被记为缺卡。

对于这些只使用飞书考勤系统，而未使用飞书审批系统的企业，可以通过考勤开放接口的形式，将三方审批结果数据回写到飞书的考勤系统中。

目前支持加班、请假、出差和外出这四种审批结果的写入。
#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_approval/add-approved-data-in-feishu-attendance2](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_approval/add-approved-data-in-feishu-attendance2)

### URL

`https://open.feishu.cn/open-apis/attendance/v1/user_approvals`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.attendance.update_attendance_remedy_approval(pylark.UpdateAttendanceRemedyApprovalReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



对于只使用飞书考勤系统而未使用飞书审批系统的企业，可以通过该接口更新写入飞书考勤系统中的三方系统审批状态，例如请假、加班、外出、出差、补卡等审批，状态包括通过、不通过、撤销等。
发起状态的审批才可以被更新为通过、不通过，已经通过的审批才可以被更新为撤销。

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/notify-approval-status-update2](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/notify-approval-status-update2)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.file.upload_image(pylark.UploadImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

上传图片接口，可以上传 JPEG、PNG、WEBP、GIF、TIFF、BMP、ICO格式图片

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 上传的图片大小不能超过10MB




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.file.download_image(pylark.DownloadImageReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

下载图片资源，只能下载应用自己上传且图片类型为message的图片

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 只能下载机器人自己上传且图片类型为message的图片，avatar类型暂不支持下载；
- 下载用户发送的资源，请使用[获取消息中的资源文件](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-resource/get)接口




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.file.upload_file(pylark.UploadFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

上传文件，可以上传视频，音频和常见的文件类型

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 不允许上传空文件
- 示例代码中需要自行替换文件路径和鉴权Token




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.file.download_file(pylark.DownloadFileReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

下载文件接口，只能下载应用自己上传的文件

注意事项:
- 需要开启[机器人能力](/ssl:ttdoc/home/develop-a-bot-in-5-minutes/create-an-app)
- 只能下载机器人自己上传的文件
- 下载用户发送的资源，请使用[获取消息中的资源文件](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-resource/get)接口


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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.okr.get_okr_period_list(pylark.GetOKRPeriodListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

- 当前仅支持「飞书OKR 企业版」客户使用本接口。[了解更多](https://okr.feishu.cn/price)

使用tenant_access_token需要额外申请权限以应用身份访问OKR信息




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.okr.batch_get_okr(pylark.BatchGetOKRReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

- 当前仅支持「飞书OKR 企业版」客户使用本接口。[了解更多](https://okr.feishu.cn/price)

使用tenant_access_token需要额外申请权限以应用身份访问OKR信息




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.okr.get_user_okr_list(pylark.GetUserOKRListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

- 当前仅支持「飞书OKR 企业版」客户使用本接口。[了解更多](https://okr.feishu.cn/price)

使用tenant_access_token需要额外申请权限以应用身份访问OKR信息




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.ehr.get_ehr_employee_list(pylark.GetEHREmployeeListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.ehr.download_ehr_attachments(pylark.DownloadEHRAttachmentsReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```



根据文件 token 下载文件。

调用 「批量获取员工花名册信息」接口的返回值中，「文件」类型的字段 id，即是文件 token{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=ehr&version=v1&resource=attachment&method=get)










![image.png](//sf1-ttcdn-tos.pstatp.com/obj/open-platform-opendoc/bed391d2a8ce6ed2d5985ea69bf92850_9GY1mnuDXP.png)




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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.tenant.get_tenant(pylark.GetTenantReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取企业名称、企业编号等企业信息

如果ISV应用是企业创建时默认安装，并且180天内企业未打开或使用过此应用，则无法通过此接口获取到企业信息。








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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.search.create_search_data_source_item(pylark.CreateSearchDataSourceItemReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.search.get_search_data_source_item(pylark.GetSearchDataSourceItemReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.search.delete_search_data_source_item(pylark.DeleteSearchDataSourceItemReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.search.create_search_data_source(pylark.CreateSearchDataSourceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.search.get_search_data_source(pylark.GetSearchDataSourceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.search.update_search_data_source(pylark.UpdateSearchDataSourceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.search.get_search_data_source_list(pylark.GetSearchDataSourceListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.search.delete_search_data_source(pylark.DeleteSearchDataSourceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_job(pylark.GetHireJobReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_job_manager(pylark.GetHireJobManagerReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_talent(pylark.GetHireTalentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_attachment(pylark.GetHireAttachmentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_attachment_preview(pylark.GetHireAttachmentPreviewReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_resume_source(pylark.GetHireResumeSourceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.create_hire_note(pylark.CreateHireNoteReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.update_hire_note(pylark.UpdateHireNoteReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_note(pylark.GetHireNoteReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_note_list(pylark.GetHireNoteListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_referral_by_application(pylark.GetHireReferralByApplicationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_job_process_list(pylark.GetHireJobProcessListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.create_hire_application(pylark.CreateHireApplicationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.terminate_hire_application(pylark.TerminateHireApplicationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_application(pylark.GetHireApplicationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_application_list(pylark.GetHireApplicationListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_application_interview_list(pylark.GetHireApplicationInterviewListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_offer_by_application(pylark.GetHireOfferByApplicationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_offer_schema(pylark.GetHireOfferSchemaReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.make_hire_transfer_onboard_by_application(pylark.MakeHireTransferOnboardByApplicationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.update_hire_employee(pylark.UpdateHireEmployeeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_employee_by_application(pylark.GetHireEmployeeByApplicationReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.hire.get_hire_employee(pylark.GetHireEmployeeReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.create_task_collaborator(pylark.CreateTaskCollaboratorReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于新增任务执行者，一个任务最多添加50个执行者

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.get_task_collaborator_list(pylark.GetTaskCollaboratorListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于查询任务执行者列表，支持分页，最大值为50

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.delete_task_collaborator(pylark.DeleteTaskCollaboratorReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于删除任务执行者

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.create_task_follower(pylark.CreateTaskFollowerReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于创建任务关注者，一个任务最多添加50个关注者

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.get_task_follower_list(pylark.GetTaskFollowerListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.delete_task_follower(pylark.DeleteTaskFollowerReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.create_task_reminder(pylark.CreateTaskReminderReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于创建任务的提醒时间。提醒时间在截止时间基础上做偏移，但是偏移后的结果不能早于当前时间。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.get_task_reminder_list(pylark.GetTaskReminderListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.delete_task_reminder(pylark.DeleteTaskReminderReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.create_task(pylark.CreateTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口可以创建一个任务（基本信息），如果需要绑定协作者等需要调用别的资源管理接口。其中查询字段 user_id_type 是用于控制返回体中 creator_id 的类型，不传时默认返回 open_id。当使用 tenant_access_token 调用接口时不支持 user_id_type 为 user_id。

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.get_task(pylark.GetTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于获取任务详情，包括任务标题、描述、时间、来源等信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/get)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id`

### Method

`GET`

## GetTaskList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.GetTaskList(ctx, &lark.GetTaskListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.get_task_list(pylark.GetTaskListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口通过解析 header 中的  user_access_token 获取与该用户相关的所有任务

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/list)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks`

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.delete_task(pylark.DeleteTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.update_task(pylark.UpdateTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.complete_task(pylark.CompleteTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.uncomplete_task(pylark.UncompleteTaskReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于取消任务的已完成状态

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/uncomplete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/uncomplete)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/uncomplete`

### Method

`POST`

## CreateTaskComment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.CreateTaskComment(ctx, &lark.CreateTaskCommentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.create_task_comment(pylark.CreateTaskCommentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于创建和回复任务的评论。当parent_id字段为0时，为创建评论；当parent_id不为0时，为回复某条评论

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/create)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/comments`

### Method

`POST`

## GetTaskComment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.GetTaskComment(ctx, &lark.GetTaskCommentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.get_task_comment(pylark.GetTaskCommentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于通过评论ID获取评论详情

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/get)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/comments/:comment_id`

### Method

`GET`

## DeleteTaskComment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.DeleteTaskComment(ctx, &lark.DeleteTaskCommentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.delete_task_comment(pylark.DeleteTaskCommentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于通过评论ID删除评论

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/delete)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/comments/:comment_id`

### Method

`DELETE`

## UpdateTaskComment

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Task.UpdateTaskComment(ctx, &lark.UpdateTaskCommentReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.task.update_task_comment(pylark.UpdateTaskCommentReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

该接口用于更新评论内容

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/update)

### URL

`https://open.feishu.cn/open-apis/task/v1/tasks/:task_id/comments/:comment_id`

### Method

`PUT`


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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.acs.get_acs_access_record_photo(pylark.GetACSAccessRecordPhotoReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.acs.get_acs_access_record_list(pylark.GetACSAccessRecordListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.acs.get_acs_device_list(pylark.GetACSDeviceListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.acs.get_acs_user_face(pylark.GetACSUserFaceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.acs.update_acs_user_face(pylark.UpdateACSUserFaceReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.acs.get_acs_user(pylark.GetACSUserReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.acs.update_acs_user(pylark.UpdateACSUserReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.acs.get_acs_user_list(pylark.GetACSUserListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

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


# Baike

## CreateBaikeDraft

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Baike.CreateBaikeDraft(ctx, &lark.CreateBaikeDraftReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.baike.create_baike_draft(pylark.CreateBaikeDraftReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

草稿是单独的定义，并不和实体词进行绑定，创建草稿不等于创建实体词。创建草稿可能是创建实体词或更新实体词

需百科管理员审批通过后生效。当使用 user_access_token 访问时，会有 bot 通知对应用户审批结果。tenant_access_token 则不会触发通知



· 当 entity_id 字段不为空时表示更新该实体词<br>
· 当 entity_id 字段为空时表示创建新的实体词<br>
· 当 entity_id 字段为空且 outer_info 字段不为空时，会根据已有数据判断是创建还是更新实体词。原则为：当 outer_info 字段已经和已有实体词绑定，则会触发更新，若无绑定的实体词，则会创建一个新的和 outer_info 绑定的实体词




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/draft/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/draft/create)

### URL

`https://open.feishu.cn/open-apis/baike/v1/drafts`

### Method

`POST`

## CreateBaikeUpdate

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Baike.CreateBaikeUpdate(ctx, &lark.CreateBaikeUpdateReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.baike.create_baike_update(pylark.CreateBaikeUpdateReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据 draft_id 更新草稿内容，已审批的草稿无法编辑

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/draft/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/draft/update)

### URL

`https://open.feishu.cn/open-apis/baike/v1/drafts/:draft_id`

### Method

`PUT`

## GetBaikeEntity

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Baike.GetBaikeEntity(ctx, &lark.GetBaikeEntityReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.baike.get_baike_entity(pylark.GetBaikeEntityReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过词条 id 拉取对应的实体词详情信息

也支持通过 provider 和 outer_id 返回对应实体的详情数据。此时路径中的 entity_id 为固定的 enterprise_0




#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/get)

### URL

`https://open.feishu.cn/open-apis/baike/v1/entities/:entity_id`

### Method

`GET`

## GetBaikeEntityList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Baike.GetBaikeEntityList(ctx, &lark.GetBaikeEntityListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.baike.get_baike_entity_list(pylark.GetBaikeEntityListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过分页拉取词条数据，支持租户内的全部词条拉取

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/list)

### URL

`https://open.feishu.cn/open-apis/baike/v1/entities`

### Method

`GET`

## MatchBaikeEntity

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Baike.MatchBaikeEntity(ctx, &lark.MatchBaikeEntityReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.baike.match_baike_entity(pylark.MatchBaikeEntityReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

准确匹配词条的关键词、全名、别名属性，并召回对应的 ID

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/match](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/match)

### URL

`https://open.feishu.cn/open-apis/baike/v1/entities/match`

### Method

`POST`

## SearchBaikeEntity

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Baike.SearchBaikeEntity(ctx, &lark.SearchBaikeEntityReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.baike.search_baike_entity(pylark.SearchBaikeEntityReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

传入关键词，进行模糊匹配搜索相应的词条

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/search)

### URL

`https://open.feishu.cn/open-apis/baike/v1/entities/search`

### Method

`POST`

## HighlightBaikeEntity

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Baike.HighlightBaikeEntity(ctx, &lark.HighlightBaikeEntityReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.baike.highlight_baike_entity(pylark.HighlightBaikeEntityReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

用户传入一个句子，返回词条的位置和对应的 Id，通过 Get 接口可以获取词条的详情信息

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/highlight](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/highlight)

### URL

`https://open.feishu.cn/open-apis/baike/v1/entities/highlight`

### Method

`POST`


# MDM

## CreateMDMVendor

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.CreateMDMVendor(ctx, &lark.CreateMDMVendorReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.create_mdm_vendor(pylark.CreateMDMVendorReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

使用该接口创建一个交易方,字段是否必填是根据后台动态配置的，如想获取配置，主数据的配置开放文档里获取。参数均采用驼峰式

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/vendor/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/vendor/create)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/vendors`

### Method

`POST`

## UpdateMDMVendor

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.UpdateMDMVendor(ctx, &lark.UpdateMDMVendorReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.update_mdm_vendor(pylark.UpdateMDMVendorReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

使用该接口来根据id来更新交易方的全部字段，字段是否必填是根据后台动态配置的，如想获取配置，主数据的配置开放文档里获取。参数均采用驼峰式

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/vendor/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/vendor/update)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/vendors/:vendor_id`

### Method

`PUT`

## GetMDMVendorList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMVendorList(ctx, &lark.GetMDMVendorListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_vendor_list(pylark.GetMDMVendorListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据交易方编码来获取对应的交易方。参数均采用驼峰式

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/vendor/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/vendor/list)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/vendors`

### Method

`GET`

## GetMDMVendor

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMVendor(ctx, &lark.GetMDMVendorReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_vendor(pylark.GetMDMVendorReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据交易方id来获取交易方信息。参数均采用驼峰式

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/vendor/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/vendor/get)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/vendors/:vendor_id`

### Method

`GET`

## GetMDMVendorListAll

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMVendorListAll(ctx, &lark.GetMDMVendorListAllReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_vendor_list_all(pylark.GetMDMVendorListAllReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

交易方全量数据分页查询。参数均采用驼峰式

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/vendor/list_all](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/vendor/list_all)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/vendors/list_all`

### Method

`GET`

## CreateMDMLegalEntity

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.CreateMDMLegalEntity(ctx, &lark.CreateMDMLegalEntityReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.create_mdm_legal_entity(pylark.CreateMDMLegalEntityReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据该接口，来创建一个法人实体，字段是否必填是根据后台动态配置的，如想获取配置，主数据的配置开放文档里获取。参数均采用驼峰式

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/legal_entity/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/legal_entity/create)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/legal_entities`

### Method

`POST`

## UpdateMDMLegalEntity

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.UpdateMDMLegalEntity(ctx, &lark.UpdateMDMLegalEntityReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.update_mdm_legal_entity(pylark.UpdateMDMLegalEntityReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

使用该接口来根据id来更新法人实体的全部字段，字段是否必填是根据后台动态配置的，如想获取配置，主数据的配置开放文档里获取。参数均采用驼峰式

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/legal_entity/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/legal_entity/update)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/legal_entities/:legal_entity_id`

### Method

`PUT`

## GetMDMLegalEntityList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMLegalEntityList(ctx, &lark.GetMDMLegalEntityListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_legal_entity_list(pylark.GetMDMLegalEntityListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

根据法人实体编码，来获取对应的法人实体信息。参数均采用驼峰式

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/legal_entity/list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/legal_entity/list)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/legal_entities`

### Method

`GET`

## GetMDMLegalEntity

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMLegalEntity(ctx, &lark.GetMDMLegalEntityReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_legal_entity(pylark.GetMDMLegalEntityReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

通过该接口，凭借法人主体id 获取单个法人实体信息。参数均采用驼峰式

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/legal_entity/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/legal_entity/get)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/legal_entities/:legal_entity_id`

### Method

`GET`

## GetMDMLegalEntityListAll

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMLegalEntityListAll(ctx, &lark.GetMDMLegalEntityListAllReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_legal_entity_list_all(pylark.GetMDMLegalEntityListAllReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

法人实体全量数据分页查询。参数均采用驼峰式

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/legal_entity/list_all](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/legal_entity/list_all)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/legal_entities/list_all`

### Method

`GET`

## CreateMDMInternalOrder

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.CreateMDMInternalOrder(ctx, &lark.CreateMDMInternalOrderReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.create_mdm_internal_order(pylark.CreateMDMInternalOrderReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

新建单个内部订单

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/internal_order/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/internal_order/create)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/internal_orders`

### Method

`POST`

## UpdateMDMInternalOrder

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.UpdateMDMInternalOrder(ctx, &lark.UpdateMDMInternalOrderReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.update_mdm_internal_order(pylark.UpdateMDMInternalOrderReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新内部订单部分字段，没有填写的字段不会被更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/internal_order/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/internal_order/update)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/internal_orders/:internal_order_uid`

### Method

`PUT`

## DeleteMDMInternalOrder

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.DeleteMDMInternalOrder(ctx, &lark.DeleteMDMInternalOrderReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.delete_mdm_internal_order(pylark.DeleteMDMInternalOrderReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除单个内部订单

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/internal_order/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/internal_order/delete)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/internal_orders/:internal_order_uid`

### Method

`DELETE`

## GetMDMInternalOrder

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMInternalOrder(ctx, &lark.GetMDMInternalOrderReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_internal_order(pylark.GetMDMInternalOrderReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取单个内部订单

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/internal_order/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/internal_order/get)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/internal_orders/:internal_order_uid`

### Method

`GET`

## SearchMDMInternalOrder

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.SearchMDMInternalOrder(ctx, &lark.SearchMDMInternalOrderReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.search_mdm_internal_order(pylark.SearchMDMInternalOrderReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

分页搜索内部订单

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/internal_order/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/internal_order/search)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/internal_orders/search`

### Method

`POST`

## CreateMDMCostCenter

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.CreateMDMCostCenter(ctx, &lark.CreateMDMCostCenterReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.create_mdm_cost_center(pylark.CreateMDMCostCenterReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

新建单个成本中心

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/cost_center/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/cost_center/create)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/cost_centers`

### Method

`POST`

## UpdateMDMCostCenter

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.UpdateMDMCostCenter(ctx, &lark.UpdateMDMCostCenterReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.update_mdm_cost_center(pylark.UpdateMDMCostCenterReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新成本中心部分字段，没有填写的字段不会被更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/cost_center/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/cost_center/update)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/cost_centers/:cost_center_uid`

### Method

`PUT`

## DeleteMDMCostCenter

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.DeleteMDMCostCenter(ctx, &lark.DeleteMDMCostCenterReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.delete_mdm_cost_center(pylark.DeleteMDMCostCenterReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除单个成本中心

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/cost_center/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/cost_center/delete)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/cost_centers/:cost_center_uid`

### Method

`DELETE`

## GetMDMCostCenter

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMCostCenter(ctx, &lark.GetMDMCostCenterReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_cost_center(pylark.GetMDMCostCenterReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取单个成本中心

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/cost_center/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/cost_center/get)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/cost_centers/:cost_center_uid`

### Method

`GET`

## SearchMDMCostCenter

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.SearchMDMCostCenter(ctx, &lark.SearchMDMCostCenterReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.search_mdm_cost_center(pylark.SearchMDMCostCenterReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

分页搜索成本中心

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/cost_center/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/cost_center/search)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/cost_centers/search`

### Method

`POST`

## CreateMDMDepartmentCostCenterRelationship

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.CreateMDMDepartmentCostCenterRelationship(ctx, &lark.CreateMDMDepartmentCostCenterRelationshipReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.create_mdm_department_cost_center_relationship(pylark.CreateMDMDepartmentCostCenterRelationshipReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

创建单个部门与成本中心关系

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/department_cost_center_relationship/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/department_cost_center_relationship/create)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/department_cost_center_relationships`

### Method

`POST`

## UpdateMDMDepartmentCostCenterRelationship

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.UpdateMDMDepartmentCostCenterRelationship(ctx, &lark.UpdateMDMDepartmentCostCenterRelationshipReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.update_mdm_department_cost_center_relationship(pylark.UpdateMDMDepartmentCostCenterRelationshipReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新部门与成本中心关系部分字段，没有填写的字段不会被更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/department_cost_center_relationship/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/department_cost_center_relationship/update)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/department_cost_center_relationships/:department_cost_center_relationship_uid`

### Method

`PUT`

## DeleteMDMDepartmentCostCenterRelationship

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.DeleteMDMDepartmentCostCenterRelationship(ctx, &lark.DeleteMDMDepartmentCostCenterRelationshipReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.delete_mdm_department_cost_center_relationship(pylark.DeleteMDMDepartmentCostCenterRelationshipReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除单个部门与成本中心关系

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/department_cost_center_relationship/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/department_cost_center_relationship/delete)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/department_cost_center_relationships/:department_cost_center_relationship_uid`

### Method

`DELETE`

## GetMDMDepartmentCostCenterRelationship

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMDepartmentCostCenterRelationship(ctx, &lark.GetMDMDepartmentCostCenterRelationshipReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_department_cost_center_relationship(pylark.GetMDMDepartmentCostCenterRelationshipReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取单个部门与成本中心关系

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/department_cost_center_relationship/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/department_cost_center_relationship/get)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/department_cost_center_relationships/:department_cost_center_relationship_uid`

### Method

`GET`

## SearchMDMDepartmentCostCenterRelationship

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.SearchMDMDepartmentCostCenterRelationship(ctx, &lark.SearchMDMDepartmentCostCenterRelationshipReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.search_mdm_department_cost_center_relationship(pylark.SearchMDMDepartmentCostCenterRelationshipReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

分页搜索部门与成本中心关系

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/department_cost_center_relationship/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/department_cost_center_relationship/search)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/department_cost_center_relationships/search`

### Method

`POST`

## CreateMDMGlAccount

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.CreateMDMGlAccount(ctx, &lark.CreateMDMGlAccountReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.create_mdm_gl_account(pylark.CreateMDMGlAccountReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

创建单个会计科目

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account/create)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/gl_accounts`

### Method

`POST`

## UpdateMDMGlAccount

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.UpdateMDMGlAccount(ctx, &lark.UpdateMDMGlAccountReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.update_mdm_gl_account(pylark.UpdateMDMGlAccountReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新会计科目部分字段，没有填写的字段不会被更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account/update)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/gl_accounts/:gl_account_uid`

### Method

`PUT`

## DeleteMDMGlAccount

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.DeleteMDMGlAccount(ctx, &lark.DeleteMDMGlAccountReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.delete_mdm_gl_account(pylark.DeleteMDMGlAccountReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除单个会计科目

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account/delete)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/gl_accounts/:gl_account_uid`

### Method

`DELETE`

## GetMDMDepGlAccount

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMDepGlAccount(ctx, &lark.GetMDMDepGlAccountReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_dep_gl_account(pylark.GetMDMDepGlAccountReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取单个会计科目

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account/get)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/gl_accounts/:gl_account_uid`

### Method

`GET`

## SearchMDMGlAccount

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.SearchMDMGlAccount(ctx, &lark.SearchMDMGlAccountReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.search_mdm_gl_account(pylark.SearchMDMGlAccountReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

分页搜索会计科目

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account/search)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/gl_accounts/search`

### Method

`POST`

## CreateMDMCompany

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.CreateMDMCompany(ctx, &lark.CreateMDMCompanyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.create_mdm_company(pylark.CreateMDMCompanyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

创建单个公司

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company/create)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies`

### Method

`POST`

## UpdateMDMCompany

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.UpdateMDMCompany(ctx, &lark.UpdateMDMCompanyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.update_mdm_company(pylark.UpdateMDMCompanyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新公司部分字段，没有填写的字段不会被更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company/update)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid`

### Method

`PUT`

## DeleteMDMCompany

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.DeleteMDMCompany(ctx, &lark.DeleteMDMCompanyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.delete_mdm_company(pylark.DeleteMDMCompanyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除单个公司

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company/delete)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid`

### Method

`DELETE`

## GetMDMDepCompany

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMDepCompany(ctx, &lark.GetMDMDepCompanyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_dep_company(pylark.GetMDMDepCompanyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取单个公司

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company/get)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid`

### Method

`GET`

## SearchMDMCompany

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.SearchMDMCompany(ctx, &lark.SearchMDMCompanyReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.search_mdm_company(pylark.SearchMDMCompanyReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

分页搜索公司

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company/search)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/search`

### Method

`POST`

## CreateMDMGlAccountCompanyRelationship

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.CreateMDMGlAccountCompanyRelationship(ctx, &lark.CreateMDMGlAccountCompanyRelationshipReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.create_mdm_gl_account_company_relationship(pylark.CreateMDMGlAccountCompanyRelationshipReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

创建单个会计科目与公司关系

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account_company_relationship/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account_company_relationship/create)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/gl_account_company_relationships`

### Method

`POST`

## UpdateMDMGlAccountCompanyRelationship

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.UpdateMDMGlAccountCompanyRelationship(ctx, &lark.UpdateMDMGlAccountCompanyRelationshipReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.update_mdm_gl_account_company_relationship(pylark.UpdateMDMGlAccountCompanyRelationshipReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新会计科目与公司关系部分字段，没有填写的字段不会被更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account_company_relationship/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account_company_relationship/update)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/gl_account_company_relationships/:gl_account_company_relationship_uid`

### Method

`PUT`

## DeleteMDMGlAccountCompanyRelationship

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.DeleteMDMGlAccountCompanyRelationship(ctx, &lark.DeleteMDMGlAccountCompanyRelationshipReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.delete_mdm_gl_account_company_relationship(pylark.DeleteMDMGlAccountCompanyRelationshipReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除单个会计科目与公司关系

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account_company_relationship/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account_company_relationship/delete)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/gl_account_company_relationships/:gl_account_company_relationship_uid`

### Method

`DELETE`

## GetMDMDepGlAccountCompanyRelationship

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMDepGlAccountCompanyRelationship(ctx, &lark.GetMDMDepGlAccountCompanyRelationshipReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_dep_gl_account_company_relationship(pylark.GetMDMDepGlAccountCompanyRelationshipReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取单个会计科目与公司关系

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account_company_relationship/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account_company_relationship/get)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/gl_account_company_relationships/:gl_account_company_relationship_uid`

### Method

`GET`

## SearchMDMGlAccountCompanyRelationship

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.SearchMDMGlAccountCompanyRelationship(ctx, &lark.SearchMDMGlAccountCompanyRelationshipReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.search_mdm_gl_account_company_relationship(pylark.SearchMDMGlAccountCompanyRelationshipReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

分页搜索会计科目与公司关系

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account_company_relationship/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/gl_account_company_relationship/search)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/gl_account_company_relationships/search`

### Method

`POST`

## CreateMDMCompanyBankAccount

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.CreateMDMCompanyBankAccount(ctx, &lark.CreateMDMCompanyBankAccountReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.create_mdm_company_bank_account(pylark.CreateMDMCompanyBankAccountReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

创建单个公司银行账号

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-company_bank_account/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-company_bank_account/create)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid/company_bank_accounts`

### Method

`POST`

## UpdateMDMCompanyBankAccount

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.UpdateMDMCompanyBankAccount(ctx, &lark.UpdateMDMCompanyBankAccountReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.update_mdm_company_bank_account(pylark.UpdateMDMCompanyBankAccountReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新公司银行账号部分字段，没有填写的字段不会被更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-company_bank_account/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-company_bank_account/update)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid/company_bank_accounts/:company_bank_account_uid`

### Method

`PUT`

## DeleteMDMCompanyBankAccount

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.DeleteMDMCompanyBankAccount(ctx, &lark.DeleteMDMCompanyBankAccountReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.delete_mdm_company_bank_account(pylark.DeleteMDMCompanyBankAccountReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除单个公司银行账号

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-company_bank_account/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-company_bank_account/delete)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid/company_bank_accounts/:company_bank_account_uid`

### Method

`DELETE`

## GetMDMDepCompanyBankAccount

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMDepCompanyBankAccount(ctx, &lark.GetMDMDepCompanyBankAccountReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_dep_company_bank_account(pylark.GetMDMDepCompanyBankAccountReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取单个公司银行账号

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-company_bank_account/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-company_bank_account/get)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid/company_bank_accounts/:company_bank_account_uid`

### Method

`GET`

## SearchMDMCompanyBankAccount

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.SearchMDMCompanyBankAccount(ctx, &lark.SearchMDMCompanyBankAccountReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.search_mdm_company_bank_account(pylark.SearchMDMCompanyBankAccountReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

分页搜索公司银行账号

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-company_bank_account/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-company_bank_account/search)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid/company_bank_accounts/search`

### Method

`POST`

## CreateMDMCompanyAsset

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.CreateMDMCompanyAsset(ctx, &lark.CreateMDMCompanyAssetReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.create_mdm_company_asset(pylark.CreateMDMCompanyAssetReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

创建单个公司资产

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-asset/create](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-asset/create)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid/assets`

### Method

`POST`

## UpdateMDMCompanyAsset

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.UpdateMDMCompanyAsset(ctx, &lark.UpdateMDMCompanyAssetReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.update_mdm_company_asset(pylark.UpdateMDMCompanyAssetReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

更新公司资产部分字段，没有填写的字段不会被更新

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-asset/update](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-asset/update)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid/assets/:asset_uid`

### Method

`PUT`

## DeleteMDMCompanyAsset

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.DeleteMDMCompanyAsset(ctx, &lark.DeleteMDMCompanyAssetReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.delete_mdm_company_asset(pylark.DeleteMDMCompanyAssetReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

删除单个公司资产

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-asset/delete](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-asset/delete)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid/assets/:asset_uid`

### Method

`DELETE`

## GetMDMDepCompanyAsset

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMDepCompanyAsset(ctx, &lark.GetMDMDepCompanyAssetReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_dep_company_asset(pylark.GetMDMDepCompanyAssetReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

获取单个公司资产

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-asset/get](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-asset/get)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid/assets/:asset_uid`

### Method

`GET`

## SearchMDMCompanyAsset

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.SearchMDMCompanyAsset(ctx, &lark.SearchMDMCompanyAssetReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.search_mdm_company_asset(pylark.SearchMDMCompanyAssetReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

分页搜索公司资产

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-asset/search](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/company-asset/search)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/companies/:company_uid/assets/search`

### Method

`POST`

## GetMDMConfigList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.MDM.GetMDMConfigList(ctx, &lark.GetMDMConfigListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.mdm.get_mdm_config_list(pylark.GetMDMConfigListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

字段配置列表查询

#

### Doc

[https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/config/config_list](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/mdm-v1/config/config_list)

### URL

`https://open.feishu.cn/open-apis/mdm/v1/config/config_list`

### Method

`GET`


# Event

## GetEventOutboundIpList

```go
package example

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

func example(ctx context.Context, cli *lark.Lark) {
	res, response, err := cli.Event.GetEventOutboundIpList(ctx, &lark.GetEventOutboundIpListReq{
		...
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req-id:", response.RequestID)
	fmt.Println("res:", res)
}

```

```python
import pylark


def example(cli: pylark.Lark):
    try:
        res, response = cli.event.get_event_outbound_ip_list(pylark.GetEventOutboundIpListReq(
            ...
        ))
    except pylark.PyLarkError as e:
        # handle exception: e
        raise

    print('req-id: %s', response.request_id)
    print('res: %s', res)

```

飞书开放平台向应用配置的回调地址推送事件时，是通过特定的IP发送出去的。如果企业需要做防火墙配置，那么可以通过这个接口获取到所有相关的IP段。

IP段有变更可能，建议企业每隔6小时定时拉取IP段更新防火墙设置，这样因IP变更导致推送失败的事件还可以通过重试解决。




#

### Doc

[https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-v1/outbound_ip/list](https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-v1/outbound_ip/list)

### URL

`https://open.feishu.cn/open-apis/event/v1/outbound_ip`

### Method

`GET`


# EventCallback


# AppLink


