package logic

import (
	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"
	"cointiger.com/verification/app/verify/model"
	"cointiger.com/verification/common/convert"
	"cointiger.com/verification/common/instance"
	"cointiger.com/verification/common/parsetoken"
	"context"
	"fmt"
	"strconv"

	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type AddverifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddverifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddverifyLogic {
	return &AddverifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddverifyLogic) Addverify(req *types.AddVerifyRequest, token string) (resp *types.AddVerifyResponse, err error) {

	resp = &types.AddVerifyResponse{
		Repetition: make([]types.RepetitionVerifyData, 0),
	}
	strUserId, userName := GetUserInfo(l.svcCtx, token)
	userId, _ := strconv.ParseInt(strUserId, 10, 64)

	verifyType := req.VerifyType
	dataList := req.Data

	for i, info := range dataList {
		querySql := fmt.Sprintf("where verify_info = '%s'", info.VerifyInfo)
		arrys, _ := l.svcCtx.OfficialVerify.CommonFind(l.ctx, querySql, "", "limit 1")
		if len(arrys) > 0 {
			resp.Repetition = append(resp.Repetition, types.RepetitionVerifyData{
				Index:      int64(i),
				VerifyInfo: info.VerifyInfo,
				JobTitle:   info.JobTitle,
				IsPay:      info.IsPay,
			})
		}
	}

	if len(resp.Repetition) > 0 {
		return
	}

	for _, info := range dataList {
		officialVerify := &model.OfficialVerify{
			VerifyType: verifyType,
			VerifyInfo: info.VerifyInfo,
			JobTiele:   convert.StrToNullString(info.JobTitle),
			IsPay:      convert.StrToNullString(info.IsPay),
			Creator:    convert.StrToNullString(userName),
			CreatorId:  userId,
			SocialName: convert.StrToNullString(req.SocialName),
		}
		_, err := l.svcCtx.OfficialVerify.Insert(l.ctx, officialVerify)
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}

type userInfo struct {
	Type          string `json:"@type"`
	Browser       string `json:"browser"`
	ExpireTime    int64  `json:"expireTime"`
	Id            int    `json:"id"`
	Ipaddr        string `json:"ipaddr"`
	LoginLocation string `json:"loginLocation"`
	LoginTime     int64  `json:"loginTime"`
	Os            string `json:"os"`
	//Permissions   []string `json:"permissions"`
	Token string `json:"token"`
	User  struct {
		Avatar     string `json:"avatar"`
		CreateBy   string `json:"createBy"`
		CreateTime int64  `json:"createTime"`
		DelFlag    string `json:"delFlag"`
		Email      string `json:"email"`
		Id         int    `json:"id"`
		LoginDate  int64  `json:"loginDate"`
		LoginIp    string `json:"loginIp"`
		NickName   string `json:"nickName"`
		Params     struct {
			Type string `json:"@type"`
		} `json:"params"`
		Password string `json:"password"`
		Roles    []struct {
			Admin             bool   `json:"admin"`
			DataScope         string `json:"dataScope"`
			Flag              bool   `json:"flag"`
			Id                int    `json:"id"`
			MenuCheckStrictly bool   `json:"menuCheckStrictly"`
			Params            struct {
				Type string `json:"@type"`
			} `json:"params"`
			RoleKey  string `json:"roleKey"`
			RoleName string `json:"roleName"`
			RoleSort int    `json:"roleSort"`
			Status   string `json:"status"`
		} `json:"roles"`
		Sex      string `json:"sex"`
		Status   string `json:"status"`
		UserName string `json:"userName"`
	} `json:"user"`
	Username string `json:"username"`
}

func GetUserInfo(svc *svc.ServiceContext, token string) (userName string, loginUserId string) {
	token = strings.Replace(token, "Bearer ", "", -1)
	tokenMap := parsetoken.ParseJwt(token)
	loginUserId = tokenMap["login_user_id"].(string)
	loginUserKey := instance.GetRedisUserKey(loginUserId)
	value, err := svc.VeriflyRedis.Get(loginUserKey)
	value = strings.Replace(value, "Set", "", -1)

	if value != "" {
		user := new(userInfo)
		err = json.Unmarshal([]byte(value), user)
		if err != nil {
			fmt.Println(err)
		}
		userName = user.Username
	}
	return userName, loginUserId
}
