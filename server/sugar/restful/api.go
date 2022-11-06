package restful

import (
	"errors"
	"fmt"
	"path"

	"server/sugar/futils"
	"server/sugar/log"
	"server/sugar/plugs"

	"net/http"
	"server/sugar/service"

	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

var DefaultLanguage = "zh-CN"

// 断言中断器
type AssertInterrupter struct {
	Code  int
	Error error
}

type Api struct {
	Context *gin.Context
	Orm     *gorm.DB
	Errors  error
}

func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		log.Error(err)
		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}
}

// MakeContext 设置http上下文
func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	return e
}

func (e *Api) BindUris(c ...interface{}) error {
	for _, p := range c {
		err := e.Context.BindUri(p)
		if err != nil {
			return err
		}
	}
	return nil
}

// Bind 参数校验
func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = e.Context.ShouldBindUri(d)
		} else {
			err = e.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			log.Warn("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			e.AddError(err)
			break
		}
	}
	//vd.SetErrorFactory(func(failPath, msg string) error {
	//	return fmt.Errorf(`"validation failed: %s %s"`, failPath, msg)
	//})
	if err1 := vd.Validate(d); err1 != nil {
		e.AddError(err1)
	}
	return e
}

// GetOrm 获取Orm DB
func (e Api) GetOrm() (*gorm.DB, error) {
	var err error
	db, ok := e.Context.Get("db")
	if !ok {
		err = errors.New("数据库连接获取失败")
		log.Error(http.StatusInternalServerError, err, "")
		e.AddError(err)
	}
	e.Orm = db.(*gorm.DB)
	return e.Orm, err
}

func (e *Api) Make(c *gin.Context, s *service.Service) *Api {
	return e.MakeContext(c).MakeOrm().MakeService(s)
}

// MakeOrm 设置Orm DB
func (e *Api) MakeOrm() *Api {
	db, err := e.Context.Get("db")
	if !err {
		err := errors.New("数据库连接获取失败")
		log.Error(http.StatusInternalServerError, err, "")
		e.AddError(err)
	}
	e.Orm = db.(*gorm.DB)
	return e
}

func (e *Api) TraceID(c *service.Service) string {
	val, ok := e.Context.Get(plugs.TraceIdKey)
	if ok {
		return val.(string)
	}
	return ""
}

func (e *Api) MakeService(c *service.Service) *Api {
	c.Orm = e.Orm
	val, ok := e.Context.Get(plugs.TraceIdKey)
	if ok {
		c.RequestId = val.(string)
	}
	return e
}

// Error 通常错误数据处理
func (e Api) Error(code int, err error) {
	Error(e.Context, code, err)
}

// OK 通常成功数据处理
func (e Api) OK(data interface{}, msg string) {
	OK(e.Context, data, msg)
}

// AssertError 错误断言，中止后面所有行为
func (e Api) AssertError(err error, code int) {
	if err != nil {
		panic(AssertInterrupter{Code: code, Error: err})
	}
}

// Assert 条件断言，中止后面所有行为
func (e Api) Assert(assert bool, code int, message string) {
	if assert {
		panic(AssertInterrupter{Code: code, Error: errors.New(message)})
	}
}

// PageOK 分页数据处理
func (e Api) PageOK(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	PageOK(e.Context, result, count, pageIndex, pageSize, msg)
}

// Custom 兼容函数
func (e Api) Custom(data gin.H) {
	Custum(e.Context, data)
}

//
func (e *Api) SaveFileAs(filePath string) error {
	files, err := e.Context.FormFile("file")
	if err != nil {
		return err
	}
	// 上传文件至指定目录
	dir := path.Dir(filePath)
	err = futils.MkDirIfNotExist(dir)
	if err != nil {
		return err
	}
	return e.Context.SaveUploadedFile(files, filePath)
}
