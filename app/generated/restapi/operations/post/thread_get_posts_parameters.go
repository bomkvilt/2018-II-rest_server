// Code generated by go-swagger; DO NOT EDIT.

package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewThreadGetPostsParams creates a new ThreadGetPostsParams object
// with the default values initialized.
func NewThreadGetPostsParams() ThreadGetPostsParams {

	var (
		// initialize parameters with default values

		limitDefault = int32(100)

		sortDefault = string("flat")
	)

	return ThreadGetPostsParams{
		Limit: &limitDefault,

		Sort: &sortDefault,
	}
}

// ThreadGetPostsParams contains all the bound params for the thread get posts operation
// typically these are obtained from a http.Request
//
// swagger:parameters threadGetPosts
type ThreadGetPostsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Флаг сортировки по убыванию.

	  In: query
	*/
	Desc *bool
	/*Максимальное кол-во возвращаемых записей.
	  Maximum: 10000
	  Minimum: 1
	  In: query
	  Default: 100
	*/
	Limit *int32
	/*Идентификатор поста, после которого будут выводиться записи
	(пост с данным идентификатором в результат не попадает).

	  In: query
	*/
	Since *int64
	/*Идентификатор ветки обсуждения.
	  Required: true
	  In: path
	*/
	SlugOrID string
	/*Вид сортировки:

	 * flat - по дате, комментарии выводятся простым списком в порядке создания;
	 * tree - древовидный, комментарии выводятся отсортированные в дереве
	   по N штук;
	 * parent_tree - древовидные с пагинацией по родительским (parent_tree),
	   на странице N родительских комментов и все комментарии прикрепленные
	   к ним, в древвидном отображение.

	Подробности: https://park.mail.ru/blog/topic/view/1191/

	  In: query
	  Default: "flat"
	*/
	Sort *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewThreadGetPostsParams() beforehand.
func (o *ThreadGetPostsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDesc, qhkDesc, _ := qs.GetOK("desc")
	if err := o.bindDesc(qDesc, qhkDesc, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qSince, qhkSince, _ := qs.GetOK("since")
	if err := o.bindSince(qSince, qhkSince, route.Formats); err != nil {
		res = append(res, err)
	}

	rSlugOrID, rhkSlugOrID, _ := route.Params.GetOK("slug_or_id")
	if err := o.bindSlugOrID(rSlugOrID, rhkSlugOrID, route.Formats); err != nil {
		res = append(res, err)
	}

	qSort, qhkSort, _ := qs.GetOK("sort")
	if err := o.bindSort(qSort, qhkSort, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDesc binds and validates parameter Desc from query.
func (o *ThreadGetPostsParams) bindDesc(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("desc", "query", "bool", raw)
	}
	o.Desc = &value

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *ThreadGetPostsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewThreadGetPostsParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int32", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *ThreadGetPostsParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.Minimum("limit", "query", float64(*o.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.Maximum("limit", "query", float64(*o.Limit), 10000, false); err != nil {
		return err
	}

	return nil
}

// bindSince binds and validates parameter Since from query.
func (o *ThreadGetPostsParams) bindSince(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("since", "query", "int64", raw)
	}
	o.Since = &value

	return nil
}

// bindSlugOrID binds and validates parameter SlugOrID from path.
func (o *ThreadGetPostsParams) bindSlugOrID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.SlugOrID = raw

	return nil
}

// bindSort binds and validates parameter Sort from query.
func (o *ThreadGetPostsParams) bindSort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewThreadGetPostsParams()
		return nil
	}

	o.Sort = &raw

	if err := o.validateSort(formats); err != nil {
		return err
	}

	return nil
}

// validateSort carries on validations for parameter Sort
func (o *ThreadGetPostsParams) validateSort(formats strfmt.Registry) error {

	if err := validate.Enum("sort", "query", *o.Sort, []interface{}{"flat", "tree", "parent_tree"}); err != nil {
		return err
	}

	return nil
}
