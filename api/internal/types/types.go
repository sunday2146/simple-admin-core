// Code generated by goctl. DO NOT EDIT.
package types

// Create or update role information params | 创建或更新角色信息参数
// swagger:model RoleInfo
type RoleInfo struct {
	BaseInfo
	// Role translation | 角色名称国际化
	Title string `json:"title,optional"`
	// Role Name | 角色名
	// Required : true
	// Max length: 20
	Name string `json:"name" validate:"max=20"`
	// Role value | 角色值
	// Required : true
	// Max length: 10
	Value string `json:"value" validate:"max=10"`
	// Role's default page | 角色默认管理页面
	// Required : true
	// Max length: 20
	DefaultRouter string `json:"defaultRouter" validate:"max=50"`
	// Role status | 角色状态
	// Required : true
	// Maximum: 20
	Status uint32 `json:"status" validate:"number,max=20"`
	// Role remark | 角色备注
	// Required : true
	// Max length: 200
	Remark string `json:"remark,optional" validate:"omitempty,max=200"`
	// Role's sorting number | 角色排序
	// Required : true
	// Maximum: 1000
	Sort uint32 `json:"sort" validate:"number,max=1000"`
}

// The response data of role list | 角色列表返回数据
// swagger:model RoleListResp
type RoleListResp struct {
	BaseDataInfo
	// The role list data | 角色列表数据
	Data RoleListInfo `json:"data"`
}

// The data of role list | 角色列表数据
// swagger:model RoleListInfo
type RoleListInfo struct {
	BaseListInfo
	// The role list data | 角色列表数据
	Data []RoleInfo `json:"data"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseDataInfo
type BaseDataInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseListInfo
type BaseListInfo struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response without data | 基础不带数据信息
// swagger:model BaseMsgResp
type BaseMsgResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
}

// The simplest message | 最简单的信息
// swagger:response SimpleMsg
type SimpleMsg struct {
	// Message | 信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageInfo
type PageInfo struct {
	// Page number | 第几页
	// Required: true
	Page uint64 `json:"page" validate:"number"`
	// Page size | 单页数据行数
	// Required: true
	// Maximum: 100000
	PageSize uint64 `json:"pageSize" validate:"number,max=100000"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
}

// Basic IDs request | 基础ID数组参数请求
// swagger:model IDsReq
type IDsReq struct {
	// IDs
	// Required: true
	Ids []uint64 `json:"ids"`
}

// Basic ID request | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// ID
	// Required: true
	// Max length: 36
	Id string `json:"id" validate:"len=36"`
}

// Basic UUID array request | 基础UUID数组参数请求
// swagger:model UUIDsReq
type UUIDsReq struct {
	// Ids
	// Required: true
	// Max length: 36
	Ids []string `json:"ids"`
}

// The base response data | 基础信息
// swagger:model BaseInfo
type BaseInfo struct {
	// ID
	Id uint64 `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
}

// The base UUID response data | 基础信息
// swagger:model BaseUUIDInfo
type BaseUUIDInfo struct {
	// ID
	Id string `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
}

// The request params of setting status code | 设置状态参数
// swagger:model StatusCodeReq
type StatusCodeReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
	// Status code | 状态码
	// Required: true
	Status uint32 `json:"status" validate:"number"`
}

// The request params of setting status code by UUID | 根据UUID设置状态参数
// swagger:model StatusCodeUUIDReq
type StatusCodeUUIDReq struct {
	// ID
	// Required: true
	Id string `json:"id"`
	// Status code | 状态码
	// Required: true
	Status uint32 `json:"status" validate:"number"`
}

// login request | 登录参数
// swagger:model LoginReq
type LoginReq struct {
	// User Name | 用户名
	// Required: true
	// Max length: 20
	Username string `json:"username" validate:"alphanum,max=20"`
	// Password | 密码
	// Required: true
	// Min length: 6
	// Max length: 30
	Password string `json:"password" validate:"max=30,min=6"`
	// Captcha ID which store in redis | 验证码编号, 存在redis中
	// Required: true
	// Max length: 20
	CaptchaId string `json:"captchaId"  validate:"len=20"`
	// The Captcha which users input | 用户输入的验证码
	// Required: true
	// Max length: 5
	Captcha string `json:"captcha" validate:"len=5"`
}

// The log in response data | 登录返回数据
// swagger:model LoginResp
type LoginResp struct {
	BaseDataInfo
	// The log in information | 登陆返回的数据信息
	Data LoginInfo `json:"data"`
}

// The log in information | 登陆返回的数据信息
// swagger:model LoginInfo
type LoginInfo struct {
	// User's UUID | 用户的UUID
	UserId string `json:"userId"`
	// User's role information| 用户的角色信息
	// in: body
	Role RoleInfoSimple `json:"role"`
	// Token for authorization | 验证身份的token
	Token string `json:"token"`
	// Expire timestamp | 过期时间戳
	Expire uint64 `json:"expire"`
}

// The profile information | 个人信息
// swagger:model ProfileInfo
type ProfileInfo struct {
	// user's nickname | 用户的昵称
	Nickname string `json:"nickname"`
	// The user's avatar path | 用户的头像路径
	Avatar string `json:"avatar"`
	// User's mobile phone number | 用户的手机号码
	Mobile string `json:"mobile"`
	// The user's email address | 用户的邮箱
	Email string `json:"email"`
}

// The profile response data | 个人信息返回数据
// swagger:model ProfileResp
type ProfileResp struct {
	BaseDataInfo
	// The profile information | 个人信息
	Data ProfileInfo `json:"data"`
}

// The profile request data | 个人信息请求参数
// swagger:model ProfileReq
type ProfileReq struct {
	// user's nickname | 用户的昵称
	// Max length: 10
	Nickname string `json:"nickname" validate:"omitempty,alphanumunicode,max=10"`
	// The user's avatar path | 用户的头像路径
	Avatar string `json:"avatar"`
	// User's mobile phone number | 用户的手机号码
	// Max length: 18
	Mobile string `json:"mobile" validate:"omitempty,numeric,max=18"`
	// The user's email address | 用户的邮箱
	// Max length: 100
	Email string `json:"email" validate:"omitempty,email,max=100"`
}

// The simple role data | 简单的角色数据
// swagger:model RoleInfoSimple
type RoleInfoSimple struct {
	// Role name | 角色名
	RoleName string `json:"roleName"`
	// Role value | 角色值
	Value string `json:"value"`
}

// register request | 注册参数
// swagger:model RegisterReq
type RegisterReq struct {
	// User Name | 用户名
	// Required: true
	// Max length: 20
	Username string `json:"username" validate:"alphanum,max=20"`
	// Password | 密码
	// Required: true
	// Min length: 6
	// Max length: 30
	Password string `json:"password" validate:"max=30,min=6"`
	// Captcha ID which store in redis | 验证码编号, 存在redis中
	// Required: true
	// Max length: 20
	CaptchaId string `json:"captchaId" validate:"len=20"`
	// The Captcha which users input | 用户输入的验证码
	// Required: true
	// Max length: 5
	Captcha string `json:"captcha" validate:"len=5"`
	// The user's email address | 用户的邮箱
	// Required: true
	// Max length: 100
	Email string `json:"email" validate:"email,max=100"`
}

// change user's password request | 修改密码请求参数
// swagger:model ChangePasswordReq
type ChangePasswordReq struct {
	// User's old password | 用户旧密码
	// Required: true
	// Max length: 30
	OldPassword string `json:"oldPassword" validate:"max=30"`
	// User's new password | 用户新密码
	// Required: true
	// Max length: 30
	NewPassword string `json:"newPassword" validate:"max=30"`
}

// The response data of user's information | 用户信息返回数据
// swagger:model UserInfoResp
type UserInfoResp struct {
	BaseUUIDInfo
	// User Name | 用户名
	Username string `json:"username"`
	// User's nickname | 用户的昵称
	Nickname string `json:"nickname"`
	// User's mobile phone number | 用户的手机号码
	Mobile string `json:"mobile"`
	// User's role id | 用户的角色ID
	RoleId uint64 `json:"roleId"`
	// The user's email address | 用户的邮箱
	Email string `json:"email"`
	// The user's avatar path | 用户的头像路径
	Avatar string `json:"avatar"`
	// The user's layout mode | 用户的布局
	SideMode string `json:"sideMode"`
	// The user's status | 用户状态
	// 1 normal, 2 ban | 1 正常 2 拉黑
	Status uint32 `json:"status"`
}

// The response data of user's basic information | 用户基本信息返回数据
// swagger:model GetUserInfoResp
type GetUserInfoResp struct {
	BaseDataInfo
	// The  data of user's basic information | 用户基本信息
	Data UserBaseInfo `json:"data"`
}

// The  data of user's basic information | 用户基本信息
// swagger:model UserBaseInfo
type UserBaseInfo struct {
	// User's UUID | 用户的UUID
	UUID string `json:"userId"`
	// User's name | 用户名
	Username string `json:"username"`
	// User's nickname | 用户的昵称
	Nickname string `json:"nickname"`
	// The user's avatar path | 用户的头像路径
	Avatar string `json:"avatar"`
	// User's role information| 用户的角色信息
	// in: body
	Roles GetUserRoleInfo `json:"roles"`
}

// The response data of user's basic role information | 用户角色信息数据
// swagger:model GetUserRoleInfo
type GetUserRoleInfo struct {
	// Role name | 角色名
	RoleName string `json:"roleName"`
	// Role value for permission control | 角色值用于前端页面组件显示权限
	Value string `json:"value"`
}

// The response data of user list | 用户列表返回数据
// swagger:model UserListResp
type UserListResp struct {
	BaseDataInfo
	// The user list data | 用户列表数据
	Data UserListInfo `json:"data"`
}

// The response data of user list | 用户列表数据
// swagger:model UserListInfo
type UserListInfo struct {
	BaseListInfo
	// The user list data | 用户列表数据
	Data []UserInfoResp `json:"data"`
}

// The permission code for front end permission control | 权限码： 用于前端权限控制
// swagger:model PermCodeResp
type PermCodeResp struct {
	BaseDataInfo
	// Permission code data | 权限码数据
	Data []string `json:"data"`
}

// Create or update user information request | 创建或更新用户信息
// swagger:model CreateOrUpdateUserReq
type CreateOrUpdateUserReq struct {
	// User's id | 用户ID
	// Required: true
	Id string `json:"id"`
	// User Name | 用户名
	// Required: true
	// Max length: 30
	Username string `json:"username" validate:"alphanum,max=30"`
	// User's nickname | 用户的昵称
	// Required: true
	// Max length: 30
	Nickname string `json:"nickname" validate:"alphanumunicode,max=30"`
	// Password | 密码
	// Required: true
	// Max length: 30
	// Min length: 6
	Password string `json:"password" validate:"omitempty,max=30,min=6"`
	// User's mobile phone number | 用户的手机号码
	// Required: true
	// Max length: 18
	Mobile string `json:"mobile,optional" validate:"numeric,max=18"`
	// User's role id | 用户的角色ID
	// Required: true
	// Maximum: 1000
	RoleId uint64 `json:"roleId" validate:"number,max=1000"`
	// The user's email address | 用户的邮箱
	// Required: true
	// Max length: 100
	Email string `json:"email" validate:"email,max=100"`
	// The user's avatar path | 用户的头像路径
	// Required: true
	// Example: https://localhost/static/1.png
	Avatar string `json:"avatar"`
	// The user's status | 用户状态
	// 1 normal, 2 ban | 1 正常 2 拉黑
	// Required: true
	// Maximum: 20
	Status uint32 `json:"status" validate:"number,max=20"`
}

// Get user list request | 获取用户列表请求参数
// swagger:model GetUserListReq
type GetUserListReq struct {
	// Page number | 第几页
	// Required: true
	Page uint64 `json:"page" validate:"number"`
	// Page size | 单页数据行数
	// Required: true
	// Maximum: 100000
	PageSize uint64 `json:"pageSize" validate:"number,max=100000"`
	// User Name | 用户名
	// Max length: 20
	Username string `json:"username,optional" validate:"omitempty,alphanum,max=20"`
	// User's nickname | 用户的昵称
	// Max length: 10
	Nickname string `json:"nickname,optional" validate:"omitempty,alphanumunicode,max=10"`
	// User's mobile phone number | 用户的手机号码
	// Max length: 18
	Mobile string `json:"mobile,optional" validate:"omitempty,numeric,max=18"`
	// The user's email address | 用户的邮箱
	// Max length: 100
	Email string `json:"email,optional" validate:"omitempty,email,max=100"`
	// User's role ID | 用户的角色ID
	// Maximum: 1000
	RoleId uint64 `json:"roleId,optional" validate:"omitempty,number,max=1000"`
}

// The response data of menu information | 菜单返回数据
// swagger:model MenuInfo
type MenuInfo struct {
	BaseInfo
	// Title translation | 标题翻译
	Trans string `json:"trans,optional"`
	// Menu type: directory or menu | 菜单类型: 目录或菜单
	// 0. directory group 1. menu | 0 目录 1 菜单
	MenuType uint32 `json:"type"`
	// Parent menu ID | 父级菜单ID
	ParentId uint64 `json:"parentId"`
	// The menu level | 菜单等级
	MenuLevel uint32 `json:"level"`
	// The path to visit menu | 菜单访问路径
	Path string `json:"path"`
	// Menu name | 菜单名
	Name string `json:"name"`
	// Redirect path | 跳转路径
	Redirect string `json:"redirect"`
	// The component path | 组件路径
	Component string `json:"component"`
	// The sorting number | 排序编号
	Sort uint32 `json:"sort"`
	// If disabled | 是否禁用菜单
	Disabled bool `json:"disabled"`
	Meta
}

// The meta data of menu | 菜单的meta数据
// swagger:model Meta
type Meta struct {
	// Menu title show in page | 菜单显示名
	// Max length: 50
	Title string `json:"title" validate:"max=50"`
	// Menu Icon | 菜单图标
	// Max length: 50
	Icon string `json:"icon" validate:"max=50"`
	// Hide menu | 隐藏菜单
	HideMenu bool `json:"hideMenu" validate:"boolean"`
	// If hide the breadcrumb | 隐藏面包屑
	HideBreadcrumb bool `json:"hideBreadcrumb" validate:"boolean"`
	// Current active menu, if not nil, it will active the tab | 当前激活的菜单
	// Max length: 30
	CurrentActiveMenu string `json:"currentActiveMenu"`
	// Do not keep alive the tab | 不缓存Tab
	IgnoreKeepAlive bool `json:"ignoreKeepAlive" validate:"boolean"`
	// Hide the tab header | 当前路由不在标签页显示
	HideTab bool `json:"hideTab" validate:"boolean"`
	// Iframe path | 内嵌iframe的地址
	FrameSrc string `json:"frameSrc"`
	// The route carries parameters or not | 如果该路由会携带参数，且需要在tab页上面显示。则需要设置为true
	CarryParam bool `json:"carryParam" validate:"boolean"`
	// Hide children menu or not | 隐藏所有子菜单
	HideChildrenInMenu bool `json:"hideChildrenInMenu" validate:"boolean"`
	// Affix tab | 是否固定标签
	Affix bool `json:"affix" validate:"boolean"`
	// The maximum number of pages the router can open | 动态路由可打开Tab页数
	DynamicLevel uint32 `json:"dynamicLevel" validate:"number,lt=30"`
	// The real path of the route without dynamic part | 动态路由的实际Path, 即去除路由的动态部分
	RealPath string `json:"realPath"`
}

// The response data of menu list | 菜单列表返回数据
// swagger:model MenuListResp
type MenuListResp struct {
	BaseDataInfo
	// The menu list data | 菜单列表数据
	Data MenuListInfo `json:"data"`
}

// The  data of menu list | 菜单列表数据
// swagger:model MenuListInfo
type MenuListInfo struct {
	BaseListInfo
	// The menu list data | 菜单列表数据
	Data []*MenuInfo `json:"data"`
}

// The response data of role menu list data | 角色菜单列表数据
type GetMenuListBase struct {
	// ID
	// Required: true
	Id uint64 `json:"id"`
	// Menu type: directory or menu | 菜单类型: 目录或菜单
	MenuType uint32 `json:"type"`
	// Parent menu ID | 父级菜单ID
	ParentId uint64 `json:"parentId"`
	// The menu level | 菜单等级
	MenuLevel uint32 `json:"level"`
	// The path to visit menu | 菜单访问路径
	Path string `json:"path"`
	// Menu name | 菜单名
	Name string `json:"name"`
	// Redirect path | 跳转路径
	Redirect string `json:"redirect"`
	// The component path | 组件路径
	Component string `json:"component"`
	// The sorting number | 排序编号
	Sort uint32 `json:"sort"`
	// If disabled | 是否禁用菜单
	Disabled bool `json:"disabled"`
	Meta     Meta `json:"meta"`
}

// The response data of role menu list, show after user login | 角色菜单列表数据， 登录后自动获取
// swagger:model GetMenuListBaseResp
type GetMenuListBaseResp struct {
	BaseDataInfo
	// The data of role menu list data | 角色菜单列表数据
	Data GetMenuListBaseInfo `json:"data"`
}

// The data of role menu list, show after user login | 角色菜单列表数据
// swagger:model GetMenuListBaseInfo
type GetMenuListBaseInfo struct {
	BaseListInfo
	// The response data of role menu list data | 角色菜单列表数据
	Data []*GetMenuListBase `json:"data"`
}

// Create or update menu information request params | 创建或更新菜单信息参数
// swagger:model CreateOrUpdateMenuReq
type CreateOrUpdateMenuReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
	// Menu type: directory or menu | 菜单类型: 目录或菜单
	// Required: true
	// Maximum: 10
	MenuType uint32 `json:"type" validate:"number,max=10"`
	// Parent menu ID | 父级菜单ID
	// Required: true
	ParentId uint64 `json:"parentId" validate:"number"`
	// The path to visit menu | 菜单访问路径
	// Required: true
	// Max length: 30
	Path string `json:"path" validate:"max=200"`
	// Menu name | 菜单名
	// Required: true
	// Max length: 50
	Name string `json:"name" validate:"max=50"`
	// Redirect path | 跳转路径
	// Required: true
	// Max length: 100
	// Example: https://www.google.com
	Redirect string `json:"redirect" validate:"max=100"`
	// The component path | 组件路径
	// Required: true
	// Max length: 100
	// Example: /sys/menu/index
	Component string `json:"component" validate:"max=100"`
	// The sorting number | 排序编号
	// Required: true
	// Maximum: 1000
	Sort uint32 `json:"sort" validate:"number,max=1000"`
	// If disabled | 是否禁用菜单
	// Required: true
	Disabled bool `json:"disabled" validate:"boolean"`
	Meta
}

// Create or update menu extra parameters request params | 创建或更新菜单额外参数的请求参数
// swagger:model CreateOrUpdateMenuParamReq
type CreateOrUpdateMenuParamReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
	// Menu ID | 菜单ID
	// Required: true
	MenuId uint64 `json:"menuId" validate:"number"`
	// Data Type | 数据类型
	// Required: true
	// Min length: 1
	// Max length: 8
	DataType string `json:"dataType" validate:"min=1,max=8"`
	// Key | 键
	// Required: true
	// Min length: 1
	// Max length: 20
	Key string `json:"key" validate:"min=1,max=20"`
	// Value | 值
	// Required: true
	// Min length: 1
	// Max length: 100
	Value string `json:"value" validate:"min=1,max=100"`
}

// The response data of menu parameters  | 菜单参数返回数据
// swagger:model MenuParamResp
type MenuParamResp struct {
	BaseDataInfo
	// The information of menu parameter  | 菜单参数数据
	Data MenuParamInfo `json:"data"`
}

// The information of menu parameter  | 菜单参数数据
// swagger:model MenuParamInfo
type MenuParamInfo struct {
	BaseInfo
	// Data Type | 数据类型
	DataType string `json:"dataType"`
	// Key | 键
	Key string `json:"key"`
	// Value | 值
	Value string `json:"value"`
}

// The response data of menu parameters list which belong to some menu | 某个菜单的菜单参数列表数据
// swagger:model MenuParamListByMenuIdResp
type MenuParamListByMenuIdResp struct {
	BaseDataInfo
	// The menu list data | 菜单列表数据
	Data MenuParamListByMenuIdInfo `json:"data"`
}

// The response data of menu parameters list which belong to some menu | 某个菜单的菜单参数列表数据
// swagger:model MenuParamListByMenuIdInfo
type MenuParamListByMenuIdInfo struct {
	BaseListInfo
	// The menu list data | 菜单列表数据
	Data []MenuParamInfo `json:"data"`
}

// The information of captcha | 验证码数据
// swagger:model CaptchaInfo
type CaptchaInfo struct {
	CaptchaId string `json:"captchaId"`
	ImgPath   string `json:"imgPath"`
}

// The response data of captcha | 验证码返回数据
// swagger:model CaptchaResp
type CaptchaResp struct {
	BaseDataInfo
	// The menu authorization data | 菜单授权信息数据
	Data CaptchaInfo `json:"data"`
}

// The response data of API information | API信息
// swagger:model ApiInfo
type ApiInfo struct {
	BaseInfo
	// API path | API路径
	Path string `json:"path"`
	// Api translation | API 多语言翻译
	Title string `json:"title"`
	// API Description | API 描述
	Description string `json:"description"`
	// API group | API分组
	Group string `json:"group"`
	// API request method e.g. POST | API请求类型 如POST
	Method string `json:"method"`
}

// Create or update API information request | 创建或更新API信息
// swagger:model CreateOrUpdateApiReq
type CreateOrUpdateApiReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
	// API path | API路径
	// Required: true
	// Min length: 1
	// Max length: 50
	Path string `json:"path" validate:"min=1,max=50"`
	// API Description | API 描述
	// Required: true
	// Max length: 50
	Description string `json:"description" validate:"max=50"`
	// API group | API分组
	// Require: true
	// Min length: 1
	// Max length: 10
	Group string `json:"group" validate:"alphanum,min=1,max=10"`
	// API request method e.g. POST | API请求类型 如POST
	// Required: true
	// Min length: 3
	// Max length: 4
	Method string `json:"method" validate:"uppercase,min=3,max=4"`
}

// The response data of API list | API列表数据
// swagger:model ApiListResp
type ApiListResp struct {
	BaseDataInfo
	// API list data | API 列表数据
	Data ApiListInfo `json:"data"`
}

// API list data | API 列表数据
// swagger:model ApiListInfo
type ApiListInfo struct {
	BaseListInfo
	// The API list data | API列表数据
	Data []ApiInfo `json:"data"`
}

// Get API list request params | API列表请求参数
// swagger:model ApiListReq
type ApiListReq struct {
	PageInfo
	// API path | API路径
	// Max length: 100
	Path string `json:"path,optional" validate:"omitempty,max=100"`
	// API Description | API 描述
	// Max length: 50
	Description string `json:"description,optional" validate:"omitempty,max=50"`
	// API group | API分组
	// Max length: 10
	Group string `json:"group,optional" validate:"omitempty,alphanum,max=10"`
	// API request method e.g. POST | API请求类型 如POST
	// Max length: 4
	Method string `json:"method,optional" validate:"omitempty,uppercase,max=4"`
}

// The response data of api authorization | API授权数据
// swagger:model ApiAuthorityInfo
type ApiAuthorityInfo struct {
	// API path | API 路径
	Path string `json:"path"`
	// API method | API请求方法
	Method string `json:"method"`
}

// Create or update api authorization information request | 创建或更新API授权信息
// swagger:model CreateOrUpdateApiAuthorityReq
type CreateOrUpdateApiAuthorityReq struct {
	// Role ID | 角色ID
	// Required: true
	// Maximum: 1000
	RoleId uint64 `json:"roleId" validate:"number,max=1000"`
	// API authorization list | API授权列表数据
	// Required: true
	Data []ApiAuthorityInfo `json:"data"`
}

// The response data of api authorization list | API授权列表返回数据
// swagger:model ApiAuthorityListResp
type ApiAuthorityListResp struct {
	BaseDataInfo
	// The api authorization list data | API授权列表数据
	Data ApiAuthorityListInfo `json:"data"`
}

// The  data of api authorization list | API授权列表数据
// swagger:model ApiAuthorityListInfo
type ApiAuthorityListInfo struct {
	BaseListInfo
	// The api authorization list data | API授权列表数据
	Data []ApiAuthorityInfo `json:"data"`
}

// Create or update menu authorization information request params | 创建或更新菜单授权信息参数
// swagger:model MenuAuthorityInfoReq
type MenuAuthorityInfoReq struct {
	// role ID | 角色ID
	// Required: true
	// Maximum: 1000
	RoleId uint64 `json:"roleId" validate:"number,max=1000"`
	// menu ID array | 菜单ID数组
	// Required: true
	MenuIds []uint64 `json:"menuIds"`
}

// Menu authorization response data | 菜单授权信息数据
// swagger:model MenuAuthorityInfoResp
type MenuAuthorityInfoResp struct {
	BaseDataInfo
	// The menu authorization data | 菜单授权信息数据
	Data MenuAuthorityInfoReq `json:"data"`
}

// The response data of dictionary information | 字典信息
// swagger:model DictionaryInfo
type DictionaryInfo struct {
	BaseInfo
	// Dictionary title | 字典显示名称
	Title string `json:"title"`
	// Dictionary name | 字典名称
	Name string `json:"name"`
	// Dictionary status | 字典状态
	Status uint32 `json:"status"`
	// Dictionary description | 字典描述
	Description string `json:"description"`
}

// Create or update dictionary information request | 创建或更新字典信息请求
// swagger:model CreateOrUpdateDictionaryReq
type CreateOrUpdateDictionaryReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
	// Dictionary title | 字典显示名称
	// Required: true
	// Min length: 1
	// Max length: 50
	Title string `json:"title" validate:"min=1,max=50"`
	// Dictionary name | 字典名称
	// Required: true
	// Min length: 1
	// Max length: 50
	Name string `json:"name" validate:"min=1,max=50"`
	// Dictionary status | 字典状态
	// Required: true
	Status uint32 `json:"status" validate:"number"`
	// Dictionary description | 字典描述
	// Required: true
	// Max length: 50
	Description string `json:"description" validate:"max=50"`
}

// The response data of dictionary list | 字典列表数据
// swagger:model DictionaryListResp
type DictionaryListResp struct {
	BaseDataInfo
	// The dictionary list data | 字典列表数据
	Data DictionaryListInfo `json:"data"`
}

// The response data of dictionary list | 字典列表数据
// swagger:model DictionaryListInfo
type DictionaryListInfo struct {
	BaseListInfo
	// The dictionary list data | 字典列表数据
	Data []DictionaryInfo `json:"data"`
}

// Get dictionary list request params | 字典列表请求参数
// swagger:model DictionaryListReq
type DictionaryListReq struct {
	PageInfo
	// Dictionary title | 字典显示名称
	// Max length: 50
	Title string `json:"title,optional" validate:"max=50"`
	// Dictionary name | 字典名称
	// Max length: 50
	Name string `json:"name,optional" validate:"max=50"`
}

// The response data of dictionary information | 字典信息
// swagger:model DictionaryDetailInfo
type DictionaryDetailInfo struct {
	BaseInfo
	// Dictionary title | 字典显示名称
	Title string `json:"title"`
	// Key name | 键
	Key string `json:"key"`
	// Value | 值
	Value string `json:"value"`
	// Status | 状态
	Status uint32 `json:"status"`
}

// The response data of dictionary KV list | 字典值的列表数据
// swagger:model DictionaryDetailListResp
type DictionaryDetailListResp struct {
	BaseDataInfo
	// The dictionary list data | 字典列表数据
	Data DictionaryDetailListInfo `json:"data"`
}

// The data of dictionary KV list | 字典值的列表数据
// swagger:model DictionaryDetailListInfo
type DictionaryDetailListInfo struct {
	BaseListInfo
	// The dictionary list data | 字典列表数据
	Data []DictionaryDetailInfo `json:"data"`
}

// Create or update dictionary KV information request | 创建或更新字典键值信息请求
// swagger:model CreateOrUpdateDictionaryDetailReq
type CreateOrUpdateDictionaryDetailReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
	// Detail title | 字典值显示名称
	// Required: true
	// Min length: 1
	// Max length: 50
	Title string `json:"title" validate:"min=1,max=50"`
	// Detail key | 键
	// Required: true
	// Min length: 1
	// Max length: 50
	Key string `json:"key" validate:"min=1,max=50"`
	// Detail value | 值
	// Required: true
	Value string `json:"value"`
	// Status | 状态
	// Required: true
	Status uint32 `json:"status" validate:"number"`
	// Parent ID | 所属字典ID
	// Required: true
	ParentId uint64 `json:"parentId" validate:"number"`
}

// Get dictionary detail list by dictionary name request | 根据字典名称获取对应键值请求
// swagger:model DictionaryDetailReq
type DictionaryDetailReq struct {
	// Dictionary name | 字典名
	Name string `json:"name"`
}

// The response data of oauth provider information | 提供者信息
// swagger:model ProviderInfo
type ProviderInfo struct {
	BaseInfo
	// Provider name | 提供商名字
	Name string `json:"name"`
	// Client ID | 客户端ID
	ClientId string `json:"clientId"`
	// Client secret | 客户端密码
	ClientSecret string `json:"clientSecret"`
	// Redirect URL | 跳转URL
	RedirectURL string `json:"redirectURL"`
	// Scopes | 范围
	Scopes string `json:"scopes"`
	// Auth URL | 鉴权URL
	AuthURL string `json:"authURL"`
	// Token URL | 获取 Token 的网址
	TokenURL string `json:"tokenURL"`
	// Auth Style is specifies how the endpoint wants the client ID & client secret sent. The zero value means to auto-detect. | 鉴权方式, 0 表示自动检测
	AuthStyle int `json:"authStyle"`
	// Get User information URL | 获取用户信息地址
	InfoURL string `json:"infoURL"`
}

// Create or update provider information request | 创建或更新提供商信息
// swagger:model CreateOrUpdateProviderReq
type CreateOrUpdateProviderReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
	// Provider name | 提供商名字
	// Required: true
	// Min length: 1
	// Max length: 50
	Name string `json:"name" validate:"min=1,max=50"`
	// Client ID | 客户端ID
	// Required: true
	// Max length: 100
	ClientId string `json:"clientId" validate:"max=100"`
	// Client secret | 客户端密码
	// Require: true
	// Min length: 1
	// Max length: 100
	ClientSecret string `json:"clientSecret" validate:"min=1,max=100"`
	// Redirect URL | 跳转URL
	// Required: true
	// Max length: 200
	RedirectURL string `json:"redirectURL" validate:"max=200"`
	// Scopes | 范围
	// Required: true
	// Max length: 200
	Scopes string `json:"scopes" validate:"max=200"`
	// Auth URL | 鉴权URL
	// Required: true
	// Max length: 200
	AuthURL string `json:"authURL" validate:"max=200"`
	// Token URL | 获取 Token 的网址
	// Required: true
	// Max length: 200
	TokenURL string `json:"tokenURL" validate:"max=200"`
	// Auth Style is specifies how the endpoint wants the client ID & client secret sent. The zero value means to auto-detect. | 鉴权方式, 0 表示自动检测
	// 0 auto detect 1 client ID and client secret 2 username and password
	// Required: true
	// Example: 0
	AuthStyle int `json:"authStyle" validate:"number"`
	// Get User information URL | 获取用户信息地址
	// Required: true
	// Max length: 200
	InfoURL string `json:"infoURL" validate:"max=200"`
}

// The response data of provider list | 提供商列表返回数据
// swagger:model ProviderListResp
type ProviderListResp struct {
	BaseDataInfo
	// The provider list data | 提供商列表数据
	Data ProviderListInfo `json:"data"`
}

// The data of provider list | 提供商列表数据
// swagger:model ProviderListInfo
type ProviderListInfo struct {
	BaseListInfo
	// The provider list data | 提供商列表数据
	Data []ProviderInfo `json:"data"`
}

// Oauth log in request | Oauth 登录请求
// swagger:model OauthLoginReq
type OauthLoginReq struct {
	// State code to avoid hack | 状态码，请求前后相同避免安全问题
	// Required: true
	// Max Length: 30
	State string `json:"state" validate:"max=30"`
	// Provider name | 提供商名字
	// Required: true
	// Max Length: 40
	// Example: google
	Provider string `json:"provider" validate:"max=40"`
}

// Redirect response | 跳转网址返回信息
// swagger:model RedirectResp
type RedirectResp struct {
	BaseDataInfo
	// Redirect information | 跳转网址
	Data RedirectInfo `json:"data"`
}

// Redirect information | 跳转网址
// swagger:model RedirectInfo
type RedirectInfo struct {
	// Redirect URL | 跳转网址
	URL string `json:"URL"`
}

// The oauth callback response data | Oauth回调数据
// swagger:model CallbackResp
type CallbackResp struct {
	// User's UUID | 用户的UUID
	UserId string `json:"userId"`
	// User's role information| 用户的角色信息
	// in: body
	Role RoleInfoSimple `json:"role"`
	// Token for authorization | 验证身份的token
	Token string `json:"token"`
	// Expire timestamp | 过期时间戳
	Expire uint64 `json:"expire"`
}

// The response data of Token information | Token信息
// swagger:model TokenInfo
type TokenInfo struct {
	BaseUUIDInfo
	// User's UUID | 用户的UUID
	UUID string `json:"UUID"`
	// Token string | Token 字符串
	Token string `json:"token"`
	// Log in source such as github | Token 来源 （本地为core, 第三方如github等）
	Source string `json:"source"`
	// JWT status 0 ban 1 active | JWT状态， 0 禁止 1 正常
	Status uint32 `json:"status"`
	// Expire time | 过期时间
	ExpiredAt int64 `json:"expiredAt"`
}

// Create or update token information request | 创建或更新token信息
// swagger:model CreateOrUpdateTokenReq
type CreateOrUpdateTokenReq struct {
	// ID
	// Required: true
	Id string `json:"ID"`
	// User's UUID | 用户的UUID
	// Required: true
	// Max Length: 36
	UUID string `json:"UUID" validate:"len=36"`
	// Token string | Token 字符串
	// Required: true
	Token string `json:"token"`
	// Log in source such as github | Token 来源 （本地为core, 第三方如github等）
	// Required: true
	// Max Length: 50
	Source string `json:"source" validate:"max=50"`
	// JWT status 0 ban 1 active | JWT状态， 0 禁止 1 正常
	// Required: true
	Status uint32 `json:"status" validate:"number"`
	// Expire time | 过期时间
	// Required: true
	ExpiredAt int64 `json:"expiredAt" validate:"number"`
}

// The response data of Token list | Token列表返回数据
// swagger:model TokenListResp
type TokenListResp struct {
	BaseDataInfo
	// The token list data | Token列表数据
	Data TokenListInfo `json:"data"`
}

// The  data of Token list | Token列表数据
// swagger:model TokenListInfo
type TokenListInfo struct {
	BaseListInfo
	// The token list data | Token列表数据
	Data []TokenInfo `json:"data"`
}

// Get token list request params | token列表请求参数
// swagger:model TokenListReq
type TokenListReq struct {
	PageInfo
	// User's UUID | 用户的UUID
	// Max Length: 36
	UUID string `json:"UUID,optional" validate:"omitempty,len=36"`
	// user's nickname | 用户的昵称
	// Max length: 10
	Nickname string `json:"nickname,optional" validate:"omitempty,alphanumunicode,max=10"`
	// User Name | 用户名
	// Max length: 20
	Username string `json:"username,optional" validate:"omitempty,alphanum,max=20"`
	// The user's email address | 用户的邮箱
	// Max length: 100
	Email string `json:"email,optional" validate:"omitempty,email,max=100"`
}

// The response data of Department information | Department信息
// swagger:model DepartmentInfo
type DepartmentInfo struct {
	BaseInfo
	// Translated Name
	Trans string `json:"trans"`
	// Status
	Status uint32 `json:"status"`
	// Name
	Name string `json:"name"`
	// Ancestors
	Ancestors string `json:"ancestors"`
	// Leader
	Leader string `json:"leader"`
	// Phone
	Phone string `json:"phone"`
	// Email
	Email string `json:"email"`
	// Sort
	Sort uint32 `json:"sort"`
	// Remark
	Remark string `json:"remark"`
	// ParentId
	ParentId uint64 `json:"parentId"`
}

// Create or update Department information request | 创建或更新Department信息
// swagger:model CreateOrUpdateDepartmentReq
type CreateOrUpdateDepartmentReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id"`
	// Status
	Status uint32 `json:"status"`
	// Name
	Name string `json:"name"`
	// Ancestors
	Ancestors string `json:"ancestors,optional"`
	// Leader
	Leader string `json:"leader"`
	// Phone
	Phone string `json:"phone"`
	// Email
	Email string `json:"email"`
	// Sort
	Sort uint32 `json:"sort"`
	// Remark
	Remark string `json:"remark"`
	// ParentId
	ParentId uint64 `json:"parentId"`
}

// The response data of Department list | Department列表数据
// swagger:model DepartmentListResp
type DepartmentListResp struct {
	BaseDataInfo
	// Department list data | Department 列表数据
	Data DepartmentListInfo `json:"data"`
}

// Department list data | Department 列表数据
// swagger:model DepartmentListInfo
type DepartmentListInfo struct {
	BaseListInfo
	// The API list data | Department 列表数据
	Data []DepartmentInfo `json:"data"`
}

// Get department list request params | Department列表请求参数
// swagger:model DepartmentListReq
type DepartmentListReq struct {
	PageInfo
	// Name
	Name string `json:"name,optional"`
	// Leader
	Leader string `json:"leader,optional"`
}
