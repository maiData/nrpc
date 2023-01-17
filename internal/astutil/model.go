package astutil

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

type Service struct {
	Name    string
	Methods []Method
}

func (svc Service) ClientStructName() string {
	return strcase.ToLowerCamel(svc.Name + "Client")
}

func (svc Service) ServerStructName() string {
	return strcase.ToLowerCamel(svc.Name + "Server")
}

func (svc Service) FileName() string {
	return strcase.ToSnake(svc.Name) + ".gen.go"
}

type Method struct {
	Name     string
	Request  *NamedField
	Response *NamedField
}

func (mth Method) NatsSubject(svc Service) string {
	return fmt.Sprintf("%s.%s", svc.Name, mth.Name)
}

func (mth Method) HandlerName() string {
	return fmt.Sprintf("handle%s", mth.Name)
}

func (mth Method) ResponseTypeName(svc Service) string {
	return fmt.Sprintf("%s%sResponse", strcase.ToLowerCamel(svc.Name), strcase.ToCamel(mth.Name))
}

func (mth Method) ErrorTypeName(svc Service) string {
	return fmt.Sprintf("%s%sError", strcase.ToCamel(svc.Name), strcase.ToCamel(mth.Name))
}

type NamedField struct {
	Pointer bool
	Name    *string
	Pkg     *string
	Type    string
}

func (nf NamedField) NameOrDefault(defaultName string) string {
	if nf.Name != nil {
		return *nf.Name
	}
	return defaultName
}
