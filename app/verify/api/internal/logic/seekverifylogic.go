package logic

import (
	"cointiger.com/verification/app/verify/model"
	"cointiger.com/verification/common/xerr"
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"cointiger.com/verification/app/verify/api/internal/svc"
	"cointiger.com/verification/app/verify/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeekverifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeekverifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeekverifyLogic {
	return &SeekverifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeekverifyLogic) Seekverify(req *types.SeekVerifyRequest, language string) (resp *types.SeekVerifyResponse,
	err error) {
	resp = &types.SeekVerifyResponse{
		VerifyList: make([]types.ListVerify, 0),
	}

	if len(req.VerifyInfo) > 1024 {
		return nil, xerr.NewParamsErr(language)
	}

	if len(req.VerifyType) > 50 {
		return nil, xerr.NewParamsErr(language)
	}

	var listVerify []*model.OfficialVerify
	switch req.VerifyType {
	case "Website":
		listVerify, err = l.website(req)

	case "Telegram Username":
		listVerify, err = l.telegramUsername(req)

	case "Social media":
		listVerify, err = l.common(req)

	case "Discord ID":
		listVerify, err = l.common(req)
	}

	if err != nil {
		return
	}

	for _, data := range listVerify {
		var Verify types.ListVerify
		Verify.VerifyType = data.VerifyType
		Verify.VerifyInfo = data.VerifyInfo
		Verify.SocialName = data.SocialName.String
		Verify.JobTiele = data.JobTiele.String
		Verify.IsPay = data.IsPay.String
		Verify.CreateTime = data.CreateTime.Format("2006-01-02 15:04:05")
		Verify.Creator = data.Creator.String
		Verify.Id = int(data.Id)
		resp.VerifyList = append(resp.VerifyList, Verify)
	}

	return
}

func (l *SeekverifyLogic) website(req *types.SeekVerifyRequest) (listVerify []*model.OfficialVerify, err error) {
	listVerify = make([]*model.OfficialVerify, 0)
	var domain string

	httpOk := strings.HasPrefix(req.VerifyInfo, "https://")
	httpsOk := strings.HasPrefix(req.VerifyInfo, "http://")

	if httpsOk || httpOk {
		u, _ := url.Parse(req.VerifyInfo)
		domain = u.Hostname()
	}

	if domain == "" {
		domain = strings.Split(req.VerifyInfo,"/")[0]
	}

    rootDomain := getRootDoman(domain)
    //fmt.Println("sssssssss",rootDomain)

	querySql := fmt.Sprintf("where verify_type = '%s' and verify_info ='%s' ", req.VerifyType, rootDomain)
	verifyList, err := l.svcCtx.OfficialVerify.CommonFind(l.ctx, querySql, "", "")

	if err != nil {
		l.Logger.Error(err)
		return
	}

	for _, verify := range verifyList {
		//verify.VerifyInfo = "https://" + verify.VerifyInfo
		verify.VerifyInfo = req.VerifyInfo
		listVerify = append(listVerify, verify)
	}

	return
}

func (l *SeekverifyLogic) common(req *types.SeekVerifyRequest) (listVerify []*model.OfficialVerify, err error) {
	listVerify = make([]*model.OfficialVerify, 0)

	verifyInfo := req.VerifyInfo

	verifyInfo = strings.Replace(verifyInfo,"http://www.","", -1)

	verifyInfo = strings.Replace(verifyInfo,"https://www.","", -1)


	querySql := fmt.Sprintf("where verify_type = '%s' and verify_info ='%s' ", req.VerifyType, verifyInfo)
	verifyList, err := l.svcCtx.OfficialVerify.CommonFind(l.ctx, querySql, "", "")

	if err != nil {
		l.Logger.Error(err)
		return
	}

	for _, verify := range verifyList {
		verify.VerifyInfo = "https://www." + verify.VerifyInfo
		listVerify = append(listVerify, verify)
	}

	return
}

func (l *SeekverifyLogic) telegramUsername(req *types.SeekVerifyRequest) (listVerify []*model.OfficialVerify,
	err error) {
	listVerify = make([]*model.OfficialVerify, 0)
	var verifyInfo string

	ok := strings.HasPrefix(req.VerifyInfo, "@")
	if ok {
		verifyInfo = req.VerifyInfo[1:]
	} else {
		verifyInfo = req.VerifyInfo
	}

	verifyInfo = strings.Replace(verifyInfo,"https://t.me/","",-1)

	verifyInfo = strings.Replace(verifyInfo,"http://t.me/","",-1)

	querySql := fmt.Sprintf("where verify_type = '%s' and verify_info ='%s' ", req.VerifyType, verifyInfo)
	verifyList, err := l.svcCtx.OfficialVerify.CommonFind(l.ctx, querySql, "", "")

	if err != nil {
		l.Logger.Error(err)
		return
	}

	for _, verify := range verifyList {
		listVerify = append(listVerify, verify)
	}

	return
}

func getRootDoman(str string) string {
	_, err := regexp.MatchString("(\\w*\\.?){1}\\.(com.cn|net.cn|gov.cn|org\\.nz|org.cn|com|net|org|gov|cc|biz|info|cn|co)$", str)
	if err != nil {
		fmt.Println("Match error: ",err.Error())
	}
	reg := regexp.MustCompile("(\\w*\\.?){1}\\.(com.cn|net.cn|gov.cn|org\\.nz|org.cn|com|net|org|gov|cc|biz|info|cn|co)$")
	data := reg.Find([]byte(str))
	return string(data)
}